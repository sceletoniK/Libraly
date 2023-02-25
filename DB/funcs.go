package DB

import (
	"context"
	"fmt"

	"github.com/sceletoniK/models"
)

func (db *DB) AddBook(newBook models.NewBook) error {
	tx, err := db.conn.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("(db) AddBook dont begin transaction: %w", err)
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(context.Background(),
		`insert into book (name, author, publisher) values ($1, $2, $3)`,
		newBook.Name, newBook.Author, newBook.Publisher); err != nil {
		return fmt.Errorf("(db) AddBook dont enter new book: %w", err)
	}

	var book models.Book
	if err := db.conn.GetContext(context.Background(), &book, `select * from book order by id desc limit 1`); err != nil {
		return fmt.Errorf("(db) AddBook dont select new book: %w", err)
	}
	for _, g := range newBook.Genres {
		if _, err := tx.ExecContext(context.Background(), `insert into bookgenre(bookId, genreId) values ($1, $2)`, book.Id, g); err != nil {
			return fmt.Errorf("(db) AddBook dont link genre: %w", err)
		}
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("(db) AddBook dont commit transaction: %w", err)
	}
	return nil
}

func (db *DB) GetFilteredBooks(filter models.Filter) ([]models.Book, error) {
	var books []models.Book

	query := `select * from book, bookgenre where book.id = bookgenre.bookId`
	and := " and "
	if filter.Name != "" {
		query += `name = ` + filter.Name
	}
	if filter.Author != "" {
		query += and + `author = ` + filter.Author
	}
	if len(filter.Genres) > 0 {
		query += and + `( `
		for i, g := range filter.Genres {
			if i > 0 {
				query += ` or `
			}
			query += fmt.Sprintf(`genre.genreId = %d`, g)
		}
		query += ` )`
	}
	if err := db.conn.SelectContext(context.Background(), &books, query); err != nil {
		return nil, err
	}
	return books, nil
}
