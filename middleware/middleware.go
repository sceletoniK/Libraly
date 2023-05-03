package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sceletoniK/models"
	"github.com/sirupsen/logrus"
)

type Key struct {
	K string
}

func Auth(key []byte, lg *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lg.Info("Auth Middleware: autorization started")
			header := r.Header.Get("Authorization")
			bearer := strings.Split(header, " ")
			if len(bearer) != 2 {
				lg.Info("Auth Middleware: auth header dont found")
				w.WriteHeader(401)
				return
			}

			token := bearer[1]
			b, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				if errors.As(jwt.ErrTokenExpired, &err) {
					lg.Info("Auth Middleware: token is expired")
					w.WriteHeader(901)
					return
				}
				lg.Info("Auth Middleware: ", err)
				w.WriteHeader(401)
				return
			}
			if !b.Valid {
				lg.Info("Auth Middleware: not valid token")
				w.WriteHeader(901)
				return
			}

			claim, ok := b.Claims.(*AuthClaims)
			if !ok {
				lg.Info("Auth Middleware: claims - ", ok)
				w.WriteHeader(500)
				return
			}
			user := models.User{
				Id:    claim.ID,
				Admin: claim.Admin,
			}
			lg.Info("Auth Middleware: autorization success")
			r = r.WithContext(context.WithValue(r.Context(), Key{K: "id"}, user))
			next.ServeHTTP(w, r)
		})
	}
}

func Admin(lg *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lg.Info("Admin Middleware: autorization started")
			user := r.Context().Value(Key{K: "id"}).(models.User)
			if !user.Admin {
				lg.Info("Admin Middleware: forbidden")
				w.WriteHeader(403)
				return
			}
			lg.Info("Admin Middleware: autorization success")
			next.ServeHTTP(w, r)
		})
	}
}
