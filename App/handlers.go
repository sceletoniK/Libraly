package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/sceletoniK/middleware"
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
	if err = s.db.AddBook(r.Context(), book); err != nil {
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
	if books, err = s.db.GetFilteredBooks(r.Context(), filter); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, books)
	s.logger.Info("Success")
}

func (s *Server) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to register user")
	var newUser models.User
	err := s.bodyParse(r, &newUser)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if newUser, err = s.db.RegisterUser(r.Context(), newUser); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, newUser)
	s.logger.Info("Success")
}

func (s *Server) handlerAuthenticationUser(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to login user")
	var user models.User
	err := s.bodyParse(r, &user)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if user, err = s.db.AuthenticationUser(r.Context(), user); err != nil {
		if err == sql.ErrNoRows {
			s.responde(w, r, 401, "Unauthorized")
			s.logger.Info("User not found")
			return
		}
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	dur, err := time.ParseDuration(s.config.AccessTime)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	acst, err := middleware.GetAccessToken(dur, user, []byte(s.config.Key))
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	rfh, err := middleware.GetRefreshToken([]byte(s.config.Key))
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	dur, err = time.ParseDuration(s.config.RefreshTime)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    rfh,
		Path:     "/ref",
		HttpOnly: true,
		Expires:  time.Now().Add(dur),
	}
	s.db.AddRefreshToken(user, rfh, r.Context(), dur)
	r.AddCookie(cookie)

	s.responde(w, r, 200, acst)
	s.logger.Info("Success")
}
