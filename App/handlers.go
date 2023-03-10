package app

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sceletoniK/middleware"
	"github.com/sceletoniK/models"
)

func (s *Server) handlerGetRent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to get cart")

	var rents []models.Rent
	var err error
	if rents, err = s.db.GetUserRents(r.Context()); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, rents)
	s.logger.Info("Success")
}

func (s *Server) hadlerAcceptRent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to accept cart")

	var rent models.Rent
	err := s.bodyParse(r, &rent)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}

	dur, err := time.ParseDuration(s.config.RentDuration)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}

	if rent, err = s.db.AcceptRent(r.Context(), rent, dur); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, rent)
	s.logger.Info("Success")
}

func (s *Server) handlerGetCart(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to get cart")

	var books []models.Book
	var err error
	if books, err = s.db.GetCart(r.Context()); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, books)
	s.logger.Info("Success")
}

func (s *Server) handlerAddRentRequest(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to add rent request")
	var book models.BookInstance
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	var rent models.Rent
	if rent, err = s.db.AddRentRequest(r.Context(), book); err != nil {
		s.responde(w, r, 400, "No free books")
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, rent)
	s.logger.Info("Success")
}

func (s *Server) handlerDeleteBookInstance(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to delete instance book")
	var book models.BookInstance
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if err := s.db.DeleteBookInstance(r.Context(), book); err != nil {
		s.responde(w, r, 400, "?????????????????? ????????????????????????")
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "?????????????????? ????????????")
	s.logger.Info("Success")
}

func (s *Server) handlerAddBookInstance(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to add instance book")
	var book models.BookInstance
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if book, err = s.db.AddBookInstance(r.Context(), book); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, book)
	s.logger.Info("Success")
}

func (s *Server) handlerAddToCart(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to add to cart")
	var cart models.Cart
	err := s.bodyParse(r, &cart)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	cart.ClientId = r.Context().Value(middleware.Key{K: "id"}).(models.User).Id
	if err := s.db.AddToCart(r.Context(), cart); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "?????????? ?????????????????? ?? ??????????????")
	s.logger.Info("Success")
}

func (s *Server) handlerDeleteFromCart(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to delete from cart")
	var cart models.Cart
	err := s.bodyParse(r, &cart)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if err := s.db.DeleteFromCart(r.Context(), cart); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "?????????? ?????????????? ???? ??????????????")
	s.logger.Info("Success")
}

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
	s.responde(w, r, 200, "?????????? ??????????????????")
	s.logger.Info("Success")
}

func (s *Server) handlerEditBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to edit book")
	var book models.NewBook
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if err = s.db.EditBook(r.Context(), book); err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "?????????? ????????????????")
	s.logger.Info("Success")
}

func (s *Server) handlerDeleteBook(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to delete book")
	var book models.Book
	err := s.bodyParse(r, &book)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if err = s.db.DeleteBook(r.Context(), book); err != nil {
		s.responde(w, r, 400, "?????????? ????????????????????????")
		s.logger.Error(err)
		return
	}
	s.responde(w, r, 200, "?????????? ??????????????")
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
	acst, _, err := s.CreateTokenPair(user, w, r)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}

	s.responde(w, r, 200, acst)
	s.logger.Info("Success")
}

func (s *Server) CreateTokenPair(user models.User, w http.ResponseWriter, r *http.Request) (string, string, error) {
	var accessToken string
	var refreshToken string

	dur, err := time.ParseDuration(s.config.AccessTime)
	if err != nil {
		return accessToken, refreshToken, err
	}

	accessToken, err = middleware.GetAccessToken(dur, user, []byte(s.config.Key))
	if err != nil {
		return accessToken, refreshToken, err
	}
	dur, err = time.ParseDuration(s.config.RefreshTime)
	if err != nil {
		return accessToken, refreshToken, err
	}

	refreshToken, err = middleware.GetRefreshToken(dur, user, []byte(s.config.Key))
	if err != nil {
		return accessToken, refreshToken, err
	}
	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/refresh",
		HttpOnly: true,
		Expires:  time.Now().Add(dur),
	}
	s.db.AddRefreshToken(user, refreshToken, r.Context(), dur)
	http.SetCookie(w, cookie)
	return accessToken, refreshToken, nil
}

func (s *Server) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Try to refresh token")
	rf, err := r.Cookie("refreshToken")
	if err != nil {
		s.logger.Info("Cookie not found")
		s.httpError(w, r, 401, err)
		return
	}
	s.logger.Info("cookie was founded")
	token := rf.Value
	s.logger.Info(token)
	b, err := jwt.ParseWithClaims(token, &middleware.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.Key), nil
	})
	if err != nil {
		s.logger.Info("handlerRefreshToken: ", err)
		s.httpError(w, r, 401, err)
		return
	}
	if !b.Valid {
		s.logger.Info("handlerRefreshToken: not valid token")
		s.httpError(w, r, 401, err)
		return
	}
	claim, ok := b.Claims.(*middleware.AuthClaims)
	if !ok {
		s.logger.Error("handlerRefreshToken: claims - ", ok)
		s.httpError(w, r, 401, err)
		return
	}
	sess, err := s.db.GetRefreshToken(token, r.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			s.httpError(w, r, 401, err)
			s.logger.Info("RefreshToken not found")
			return
		}
		s.httpError(w, r, 500, err)
		s.logger.Error(err)
		return
	}
	if sess.RefreshToken != token {
		s.httpError(w, r, 401, err)
		s.logger.Info("RefreshToken dont compare")
		return
	}
	user := models.User{
		Id:    claim.ID,
		Admin: claim.Admin,
	}
	acst, _, err := s.CreateTokenPair(user, w, r)
	if err != nil {
		s.httpError(w, r, 500, err)
		s.logger.Info("handlerRefreshToken: cant create token pair")
		return
	}
	s.responde(w, r, 200, acst)
	s.logger.Info("Success")
}
