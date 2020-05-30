package utils

import (
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
		"exp": time.Now().AddDate(0, 0, 7).Unix(),
	})
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(token string) (bool, error) {
	tokenString := strings.Replace(token, "Bearer", "", 1)
	log.Println(tokenString)
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return signKey, nil
	})
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		fmt.Println(claims["foo"])
		return true, nil
	} else {
		fmt.Println(err)
		return false, err
	}
}
