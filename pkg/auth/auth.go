package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordEmpty = errors.New("request password field empty")

	secret = os.Getenv("SECRET")
)

// UserID checks the token and returns userID.
// TODO
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

	_ = token
	return "USERNAME", nil
}

// Login takes a User from the database and a password and returns a JWT.
// TODO
func Login(u *platform.User, pw string) (string, error) {

	if pw == "" {
		return "", ErrPasswordEmpty
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)); err != nil {
		return "", err
	}
	return "JWT", nil
}
