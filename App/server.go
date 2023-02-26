package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sceletoniK/models"
	"github.com/sirupsen/logrus"
)

type Libraly interface {
	AddBook(context.Context, models.NewBook) error
	GetFilteredBooks(context.Context, models.Filter) ([]models.Book, error)
	RegisterUser(context.Context, models.User) (models.User, error)
}

type Server struct {
	router     *chi.Mux
	db         Libraly
	httpServer *http.Server
	logger     *logrus.Logger
}

func NewServer(l Libraly) *Server {
	s := &Server{
		router: chi.NewRouter(),
		db:     l,
		logger: logrus.New(),
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	s.router.Post("/newbook", s.handlerNewBook)
	s.router.Get("/filterbook", s.handlerFilterBooks)
	s.router.Post("/register", s.handlerRegisterUser)

}

func (s *Server) responde(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) httpError(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.responde(w, r, code, err.Error())
}

func (s *Server) bodyParse(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}
	return nil
}
