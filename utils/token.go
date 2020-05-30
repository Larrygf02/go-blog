package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh_token"`
}

var signKey = []byte("SECRET")

func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().AddDate(0, 0, 7).Unix(),
	})
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
