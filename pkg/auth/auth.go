package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Luke-Vear/nettaton/pkg/do"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret = os.Getenv("SECRET")
)

// Login takes a User from the database and a password and returns a JWT.
func Login(u *do.User, pw string) (string, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)); err != nil {
		return "", err
	}

	return generateTokenString(u)
}

// GenPasswordHash takes a User from the database and a password and returns a JWT.
func GenPasswordHash(u *do.User, pw string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return nil
}

// Create a new token, specifying signing method and claims.
// Standard claims: https://tools.ietf.org/html/rfc7519#section-4.1
func generateTokenString(u *do.User) (string, error) {

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
			// verify alg claim
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
