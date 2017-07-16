package auth

import (
	"fmt"
	"os"
	"time"

	"strings"

	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret = os.Getenv("SECRET")
)

// Login takes a User from the database and a password and returns a JWT.
func Login(u *platform.User, pw string) (string, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)); err != nil {
		return "", err
	}

	// Create a new token object, specifying signing method and claims.
	// Standard claims: https://tools.ietf.org/html/rfc7519#section-4.1
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   u.UserID,
		ExpiresAt: time.Now().Add(time.Hour * 150).Unix(),
		NotBefore: time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// UserID checks the token and returns userID.
func UserID(bearer string) (string, error) {

	token, err := jwt.Parse(strings.Split(bearer, " ")[1],

		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

	if err != nil {
		return "", err
	}

	// Passed validation above so will have `sub`
	return token.Claims.(jwt.MapClaims)["sub"].(string), nil
}
