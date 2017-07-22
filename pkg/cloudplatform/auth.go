package cloudplatform

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// genPwHash takes a cleartext password and returns a hashed one.
func genPwHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// login takes a User from the database and a password and returns a JWT.
func login(u *User) (string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(u.ClearTextPassword)); err != nil {
		return "", err
	}
	return generateTokenString(u)
}

// Create a new token, specifying signing method and standard claims.
// Standard claims: https://tools.ietf.org/html/rfc7519#section-4.1
// Get the complete encoded token as a string using the secret.
func generateTokenString(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   u.ID,
		ExpiresAt: time.Now().Add(time.Hour * 150).Unix(),
		NotBefore: time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// parseJWT parses and validates jwt, returns claim field requested by want.
// Ensures to verify alg claims, won't accept a jwt with no/wrong alg.
func parseJWT(bearer, want string) (string, error) {
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

	if _, ok := token.Claims.(jwt.MapClaims)[want].(string); !ok {
		return "", ErrClaimNotFoundInJWT
	}
	return token.Claims.(jwt.MapClaims)[want].(string), nil
}

// IDFromToken checks the token and returns id.
func IDFromToken(bearer string) (string, error) {
	return parseJWT(bearer, "sub")
}
