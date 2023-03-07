package middleware

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sceletoniK/models"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID    int  `json:"Id"`
	Admin bool `json:"Role"`
}

func GetAccessToken(d time.Duration, user models.User, key []byte) (string, error) {
	log.Println(user.Id)
	data := AuthClaims{
		ID:    user.Id,
		Admin: user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	st, err := token.SignedString(key)
	return st, err
}

func GetRefreshToken(d time.Duration, key []byte) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(d)),
	})
	st, err := token.SignedString(key)
	return st, err
}
