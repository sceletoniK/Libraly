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
	GetFilteredBooks(context.Context, models.Filter) ([]models.Book, error)
	AddBook(context.Context, models.NewBook) error
	EditBook(context.Context, models.NewBook) error
	DeleteBook(context.Context, models.Book) error

	GetGenres(context.Context) ([]models.Genre, error)
	AddGenre(context.Context, models.Genre) (models.Genre, error)
	ChangeGenre(context.Context, models.Genre) (models.Genre, error)

	RegisterUser(context.Context, models.User) (models.User, error)
	RegisterAdmin(context.Context, models.User) (models.User, error)
	AuthenticationUser(context.Context, models.User) (models.User, error)
	AddRefreshToken(models.User, string, context.Context, time.Duration) error
	GetRefreshToken(string, context.Context) (models.Session, error)
	GetUserById(int, context.Context) (models.User, error)

	GetCart(context.Context) ([]models.Book, error)
	AddToCart(context.Context, models.Cart) error
	DeleteFromCart(context.Context, models.Cart) error

	AddBookInstance(context.Context, models.BookInstance) (models.BookInstance, error)
	DeleteBookInstance(context.Context, models.BookInstance) error

	AddRentRequest(context.Context, models.BookInstance) (models.Rent, error)
	AcceptRent(context.Context, models.Rent, time.Duration) (models.Rent, error)
	GetUserRents(context.Context) ([]models.Rent, error)
	GetAllRents(context.Context, models.FilterRent) ([]models.Rent, error)
	CloseRent(context.Context, models.Rent) error

	GetUserHistory(context.Context) ([]models.RentHistory, error)
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

		rout.Get("/user", s.handlerGetUser)

		rout.Get("/cart", s.handlerGetCart)
		rout.Post("/cart", s.handlerAddToCart)
		rout.Delete("/cart", s.handlerDeleteFromCart)

		rout.Post("/rent", s.handlerAddRentRequest)
		rout.Get("/rent", s.handlerGetUserRents)

		rout.Get("/history", s.handlerGetHistory)

		rout.Group(func(rout chi.Router) {
			rout.Use(middleware.Admin(s.logger))

			rout.Post("/book", s.handlerNewBook)
			rout.Patch("/book", s.handlerEditBook)
			rout.Delete("/book", s.handlerDeleteBook)

			rout.Post("/genre", s.handlerAddGenre)
			rout.Put("/genre", s.handlerChangeGenre)

			rout.Post("/instance", s.handlerAddBookInstance)
			rout.Delete("/instance", s.handlerDeleteBookInstance)

			rout.Put("/rent", s.hadlerAcceptRent)
			rout.Delete("/rent", s.handlerCloseRent)
			rout.Get("/rents", s.handlerGetUsersRents)

			rout.Post("/admin", s.handlerRegisterAdmin)
		})

	})
	s.router.Get("/book", s.handlerFilterBooks)
	s.router.Get("/genre", s.handlerGetGenres)

	s.router.Post("/refresh", s.handlerRefreshToken)
	s.router.Post("/register", s.handlerRegisterUser)
	s.router.Post("/login", s.handlerAuthenticationUser)
	s.router.Options("/*", s.OptionResponce)
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

func (s *Server) OptionResponce(w http.ResponseWriter, r *http.Request) {
	s.responde(w, r, 200, "xyi")
}
