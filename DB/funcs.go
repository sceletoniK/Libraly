package DB

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sceletoniK/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

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
