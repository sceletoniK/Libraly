package DB

import (
	"context"
	"errors"
	"fmt"

	"github.com/sceletoniK/models"
)

func (db *DB) AddBook(ctx context.Context, newBook models.NewBook) error {
	tx, err := db.conn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("(db) AddBook dont begin transaction: %w", err)
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(ctx,
		`insert into book (name, author, publisher) values ($1, $2, $3)`,
		newBook.Name, newBook.Author, newBook.Publisher); err != nil {
		return fmt.Errorf("(db) AddBook dont enter new book: %w", err)
	}

	var book models.Book
	if err := db.conn.GetContext(ctx, &book, `select * from book order by id desc limit 1`); err != nil {
		return fmt.Errorf("(db) AddBook dont select new book: %w", err)
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
	fmt.Println(query)
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
	if err := tx.GetContext(ctx, &addedUser, "insert into client (login, password, admin) values ($1,$2,$3) returning *", newUser.Login, newUser.Password, 0); err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser cant add new user: %w", err)
	}
	addedUser.Password = ""
	if err = tx.Commit(); err != nil {
		return addedUser, fmt.Errorf("(db) RegisterUser dont commit transaction: %w", err)
	}
	return addedUser, nil
}
