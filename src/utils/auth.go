package utils

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashString(str string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
}

func ValidateHash(hashedStr, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
}

func CreateTokenJWT(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["expire"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return jwtToken.SignedString([]byte(config.JwtSecret))
}
