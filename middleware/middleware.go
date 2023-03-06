package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(key []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			bearer := strings.Split(header, " ")
			if len(bearer) != 2 {
				log.Println("Middleware: 401")
				w.WriteHeader(401)
				return
			}

			token := bearer[1]
			b, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				log.Println(err)
				w.WriteHeader(401)
				return
			}
			if !b.Valid {
				log.Default().Println("not valid token")
				w.WriteHeader(401)
				return
			}

			_, ok := b.Claims.(*AuthClaims)
			if !ok {
				log.Default().Println("claims", ok)
				w.WriteHeader(500)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
