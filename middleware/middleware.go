package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Auth(key []byte, lg *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Middleware: Authorization")
			bearer := strings.Split(header, " ")
			if len(bearer) != 2 {
				lg.Info("Middleware: 401")
				w.WriteHeader(401)
				return
			}

			token := bearer[1]
			b, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				lg.Info("Middleware: ", err)
				w.WriteHeader(401)
				return
			}
			if !b.Valid {
				lg.Info("Middleware: not valid token")
				w.WriteHeader(401)
				return
			}

			_, ok := b.Claims.(*AuthClaims)
			if !ok {
				lg.Info("Middleware: claims - ", ok)
				w.WriteHeader(500)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
