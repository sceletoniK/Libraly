package DB

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sceletoniK/middleware"
	"github.com/sceletoniK/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

func (db *DB) GetUserRents(ctx context.Context) ([]models.Rent, error) {
	var rents []models.Rent
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return rents, fmt.Errorf("(db)  GetUserRents dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	clientId := ctx.Value(middleware.Key{K: "id"}).(models.User).Id

	if err := tx.SelectContext(ctx, &rents, "select clientId, instanceBookId, requestDate, coalesce(startRentDate, timestamp '2000-01-01 00:00:00') as startRentDate, coalesce(deadline, timestamp '2000-01-01 00:00:00') as deadline from bookrent where clientId = $1", clientId); err != nil {
		return rents, fmt.Errorf("(db)  GetUserRents dont select rents: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return rents, fmt.Errorf("(db)  GetUserRents dont commit transaction: %w", err)
	}
	return rents, nil
}

func (db *DB) AcceptRent(ctx context.Context, rent models.Rent, dur time.Duration) (models.Rent, error) {

	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return rent, fmt.Errorf("(db)  AcceptRent dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := tx.GetContext(ctx, &rent, "update bookrent set startRentDate = $1, deadline = $2 where clientid = $3 and instanceBookId = $4 returning *", time.Now(), time.Now().Add(dur), rent.ClientId, rent.InstanceId); err != nil {
		return rent, fmt.Errorf("(db)  AcceptRent cant accept rent: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return rent, fmt.Errorf("(db)  AcceptRent dont commit transaction: %w", err)
	}
	return rent, nil
}

func (db *DB) GetCart(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return books, fmt.Errorf("(db) GetCart dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	clientId := ctx.Value(middleware.Key{K: "id"}).(models.User).Id

	if err := tx.SelectContext(ctx, &books, "select book.id, book.name, book.author, book.publisher from book, bookselect where bookselect.clientId = $1 and book.id = bookselect.BookId", clientId); err != nil {
		return books, fmt.Errorf("(db) GetCart dont select cart books: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return books, fmt.Errorf("(db) GetCart dont commit transaction: %w", err)
	}
	return books, nil
}

func (db *DB) AddRentRequest(ctx context.Context, book models.BookInstance) (models.Rent, error) {
	var rent models.Rent
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return rent, fmt.Errorf("(db) AddRequestToRent dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err = tx.GetContext(ctx, &book, "select * from bookinstance WHERE originalId = $1 AND (SELECT count(*) from bookrent where instanceBookId = bookinstance.id) = 0 limit 1", book.BookId); err != nil {
		return rent, fmt.Errorf("(db) AddRequestToRent dont find free instance book: %w", err)
	}

	clientId := ctx.Value(middleware.Key{K: "id"}).(models.User).Id

	if err = tx.GetContext(ctx, &rent, "insert into bookrent(clientId, instancebookid, requestdate, startrentdate, deadline) values ($1,$2,$3,null,null) returning clientId, instancebookid, requestdate", clientId, book.InstanceId, time.Now()); err != nil {
		return rent, fmt.Errorf("(db) AddRequestToRent dont insert rent: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return rent, fmt.Errorf("(db) AddRequestToRent dont commit transaction: %w", err)
	}
	return rent, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (db *DB) AddBook(ctx context.Context, newBook models.NewBook) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) AddBook dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	var book models.Book
	if err := tx.GetContext(ctx, &book,
		`insert into book (name, author, publisher) values ($1, $2, $3) returning *`,
		newBook.Name, newBook.Author, newBook.Publisher); err != nil {
		return fmt.Errorf("(db) AddBook dont enter new book: %w", err)
	}

	for _, g := range newBook.Genres {
		if _, err := tx.ExecContext(ctx, `insert into bookgenre(bookId, genreId) values ($1, $2)`, book.Id, g); err != nil {
			return fmt.Errorf("(db) AddBook dont link genre: %w", err)
		}
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) AddBook dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) EditBook(ctx context.Context, book models.NewBook) error {

	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) EditBook dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	request := ""

	if book.Name != "" {
		request += "name = '" + book.Name + "'"
	}
	if book.Author != "" {
		if len(request) > 0 {
			request += ","
		}
		request += "author = '" + book.Author + "'"
	}
	if book.Publisher != "" {
		if len(request) > 0 {
			request += ","
		}
		request += "publisher = '" + book.Publisher + "'"
	}

	if _, err := tx.ExecContext(ctx, "update book set "+request+" where id = $1", book.Id); err != nil {
		return fmt.Errorf("(db) EditBook dont update book: %w", err)
	}

	var genres []int

	if err := tx.SelectContext(ctx, &genres, "select genreId from bookgenre where bookId = $1", book.Id); err != nil {
		return fmt.Errorf("(db) EditBook dont select genres: %w", err)
	}

	for _, g := range book.Genres {
		var j int
		if j = slices.IndexFunc(genres, func(i int) bool { return i == g }); j == -1 {
			if _, err := tx.ExecContext(ctx, "insert into bookgenre(bookId, genreId) values ($1, $2)", book.Id, g); err != nil {
				fmt.Println(book.Id, g)
				return fmt.Errorf("(db) EditBook dont insert genre: %w", err)
			}
		} else {
			genres = slices.Delete(genres, j, j)
		}
	}

	for g := range genres {
		if _, err := tx.ExecContext(ctx, "delete from bookgenre where bookId = $1 and genreId = $2", book.Id, g); err != nil {
			return fmt.Errorf("(db) EditBook dont delete genre: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) EditBook dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) DeleteBook(ctx context.Context, book models.Book) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) DelBook dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "delete from bookgenre where bookId = $1", book.Id); err != nil {
		return fmt.Errorf("(db) DelBook dont delete genres: %w", err)
	}

	if _, err := tx.ExecContext(ctx, "delete from book where id = $1", book.Id); err != nil {
		return fmt.Errorf("(db) DelBook dont delete book: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) DelBook dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) AddToCart(ctx context.Context, cart models.Cart) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) AddToCart dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "insert into bookselect(clientId, bookId) values ($1, $2)", cart.ClientId, cart.BookId); err != nil {
		return fmt.Errorf("(db) AddToCart dont insert to bookselect: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) AddToCart dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) DeleteFromCart(ctx context.Context, cart models.Cart) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) DeleteFromCart dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "delete from bookselect where clientId = $1 and bookId = $2", cart.ClientId, cart.BookId); err != nil {
		return fmt.Errorf("(db) DeleteFromCart dont insert to bookselect: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) DeleteFromCart dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) AddBookInstance(ctx context.Context, book models.BookInstance) (models.BookInstance, error) {

	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return book, fmt.Errorf("(db) AddBookInstance dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := tx.GetContext(ctx, &book, "insert into bookinstance (originalId) values ($1) returning *", book.BookId); err != nil {
		return book, fmt.Errorf("(db) AddBookInstance dont insert to bookinstance: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return book, fmt.Errorf("(db) AddBookInstance dont commit transaction: %w", err)
	}
	return book, nil
}

func (db *DB) DeleteBookInstance(ctx context.Context, book models.BookInstance) error {

	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) DeleteBookInstance dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "delete from bookinstance where id = $1", book.InstanceId); err != nil {
		return fmt.Errorf("(db) DeleteBookInstance dont insert to bookinstance: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) DeleteBookInstance dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) GetFilteredBooks(ctx context.Context, filter models.Filter) ([]models.Book, error) {
	var books []models.Book

	query := `select book.id, book.name, book.author, book.publisher from book left join bookgenre on book.id = bookgenre.bookId`
	and := " where "
	if filter.Name != "" {
		query += and + `name = '` + filter.Name + `'`
		and = " and "
	}
	if filter.Author != "" {
		query += and + `author = '` + filter.Author + `'`
		and = " and "
	}
	if len(filter.Genres) > 0 {
		query += and + `( `
		i := 0
		for _, g := range filter.Genres {
			if i > 0 {
				query += ` or `
			}
			query += fmt.Sprintf(`bookgenre.genreId = %d`, g)
			i++
		}
		query += ` )`
	}
	query += ` group by book.id, book.name, book.author, book.publisher`
	if err := db.conn.SelectContext(ctx, &books, query); err != nil {
		return nil, fmt.Errorf("(db) GetFilterBooks cant select books: %w", err)
	}
	return books, nil
}

func (db *DB) RegisterUser(ctx context.Context, newUser models.User) (models.User, error) {
	var addedUser models.User
	var othersUsers []models.User
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser dont begin transaction: %w", err)
	}
	defer tx.Rollback()
	if err := tx.SelectContext(ctx, &othersUsers, "select * from client where login = '"+newUser.Login+"'"); err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser cant select users: %w", err)
	}
	if len(othersUsers) > 0 {
		return addedUser, errors.New("(db) RegisterUser: new login isn't unique")
	}

	newUser.Password, err = HashPassword(newUser.Password)

	if err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser cant hash password: %w", err)
	}

	if err := tx.GetContext(ctx, &addedUser, "insert into client (login, password, admin) values ($1,$2,$3) returning *", newUser.Login, newUser.Password, 0); err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser cant add new user: %w", err)
	}
	addedUser.Password = ""
	if err = tx.Commit(); err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser dont commit transaction: %w", err)
	}
	return addedUser, nil
}

func (db *DB) AuthenticationUser(ctx context.Context, user models.User) (models.User, error) {
	var dbUser models.User
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return user, fmt.Errorf("(db) AuthenticationUser dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := tx.GetContext(ctx, &dbUser, "select * from client where login = $1", user.Login); err != nil {
		return user, fmt.Errorf("(db) AuthenticationUser dont select or find user: %w", err)
	}

	if !CheckPasswordHash(user.Password, dbUser.Password) {
		return user, fmt.Errorf("(db) AuthenticationUser password dont compare")
	}

	if err = tx.Commit(); err != nil {
		return user, fmt.Errorf("(db) AuthenticationUser dont commit transaction: %w", err)
	}
	return dbUser, nil
}

func (db *DB) AddRefreshToken(user models.User, token string, ctx context.Context, dur time.Duration) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) AddRefreshToken dont begin transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "delete from sessions where clientId = $1", user.Id); err != nil {
		return fmt.Errorf("(db) AddRefreshToken dont delete old refresh: %w", err)
	}

	if _, err := tx.ExecContext(ctx, "insert into sessions (refreshToken, clientId, expiresAt) values ($1, $2, $3)", token, user.Id, time.Now().Add(dur)); err != nil {
		return fmt.Errorf("(db) AddRefreshToken dont add new refresh: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) AddRefreshToken dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) GetRefreshToken(token string, ctx context.Context) (models.Session, error) {
	var s models.Session
	if err := db.conn.GetContext(ctx, &s, "select * from sessions where refreshToken = $1", token); err != nil {
		return s, fmt.Errorf("(db) GetRefreshToken dont get refresh token: %w", err)
	}
	return s, nil
}

func (db *DB) GetUserById(id int, ctx context.Context) (models.User, error) {
	var user models.User

	if err := db.conn.GetContext(ctx, &user, "select refreshToken, clientId, expiresAt from client where clientId = $1", id); err != nil {
		return user, fmt.Errorf("(db) GetUserById dont get user: %w", err)
	}

	return user, nil
}
