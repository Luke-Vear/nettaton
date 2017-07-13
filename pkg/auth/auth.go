package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret = os.Getenv("SECRET")
)

// UserID checks the token and returns userId.
func UserID(bearer string) (string, error) {

	if bearer == "" {
		return "", nil
	}

	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return "", err
	}

	// TODO
	_ = token
	return "USERNAME", nil
}
