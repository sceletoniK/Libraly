package app

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sceletoniK/config"
	"github.com/sceletoniK/middleware"
	"github.com/sceletoniK/models"
	"github.com/sirupsen/logrus"
)

type Libraly interface {
	AddBook(context.Context, models.NewBook) error
	GetFilteredBooks(context.Context, models.Filter) ([]models.Book, error)
	RegisterUser(context.Context, models.User) (models.User, error)
	AuthenticationUser(context.Context, models.User) (models.User, error)
	AddRefreshToken(models.User, string, context.Context, time.Duration) error
	GetRefreshToken(string, context.Context) (models.Session, error)
	GetUserById(int, context.Context) (models.User, error)
}

type Server struct {
	router     *chi.Mux
	db         Libraly
	httpServer *http.Server
	logger     *logrus.Logger
	config     *config.Config
}

func NewServer(l Libraly, cfg *config.Config) *Server {
	s := &Server{
		router: chi.NewRouter(),
		db:     l,
		logger: logrus.New(),
		config: cfg,
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {

	s.router.Group(func(rout chi.Router) {
		rout.Use(middleware.Auth([]byte(s.config.Key), s.logger))

		rout.Get("/filterbook", s.handlerFilterBooks)

		rout.Group(func(rout chi.Router) {
			rout.Use(middleware.Admin(s.logger))

			rout.Post("/newbook", s.handlerNewBook)
		})

	})
	s.router.Post("/refresh", s.handlerRefreshToken)
	s.router.Post("/register", s.handlerRegisterUser)
	s.router.Post("/login", s.handlerAuthenticationUser)
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
