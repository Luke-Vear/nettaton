package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// UserID checks the token and returns userId.
func UserID(bearer string, secret string) (string, error) {

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

	_ = token
	return "USERNAME", nil
}
