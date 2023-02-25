package app

import (
	"net/http"

	"github.com/sceletoniK/models"
)

func (s *Server) handlerNewBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to insert new Book")
	var book models.NewBook
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if err = s.db.AddBook(book); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "Книга добавлена")
	s.logger.Info("Success")
}

func (s *Server) handlerFilterBooks(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to filter books")
	var filter models.Filter
	err := s.bodyParse(r, &filter)
	if err != nil {
		s.httpError(w, r, 501, err)
		s.logger.Error(err)
		return
	}
	var books []models.Book
	if books, err = s.db.GetFilteredBooks(filter); err != nil {
		s.httpError(w, r, 502, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, books)
	s.logger.Info("Success")
}
