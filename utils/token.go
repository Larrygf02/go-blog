package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"
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
		"exp": time.Now().Add(time.Minute * 3).Unix(),
	})
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(token string) (bool, error) {
	tokenSplit := strings.Split(token, "Bearer ")
	if len(tokenSplit) > 1 {
		tokenString := tokenSplit[1]
		_token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return signKey, nil
		})
		if _token.Valid {
			return true, nil
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				_error := errors.New("Token malformed")
				return false, _error
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				log.Println("ERROR TOKEN EXPIRED")
				_error := errors.New("Token has expired")
				return false, _error
			} else {
				_error := errors.New("Could not handle this token")
				return false, _error
			}
		} else {
			return false, errors.New("Opp Error")
		}
	} else {
		return false, errors.New("No token")
	}
}
