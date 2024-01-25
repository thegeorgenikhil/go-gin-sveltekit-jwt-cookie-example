package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwtSecret = "aZaROAuB5xOnc8WYYPpmZBDESkLWulmpltdc7m7NY+Q="
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Claims contains the email and the standard claims for the jwt
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken creates a new token by taking in the user email and secret
func GenerateToken(email string, t time.Duration) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Local().Add(t).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(
		jwtSecret,
	))

	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken returns the Claims struct if the given token is a valid token, error if not
func ValidateToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(
				jwtSecret,
			), nil
		},
	)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
