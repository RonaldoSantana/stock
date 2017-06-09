package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// SigningKey for token generation
const SigningKey = "MinhaStringSecreta"

// Hash given password
func HashPassword(password string) (hash string, err error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		hash = string(b)
	}
	return
}

// generates API token for JWT authentication
func GenerateToken(sessionID string) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        sessionID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SigningKey))
}
