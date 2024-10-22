package utils

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func ValidateToken(r *http.Request) error {
	jwtTokenString := extractJwtToken(r)

	jwtToken, err := jwt.Parse(jwtTokenString, VerifyJwtTokenSigningMethod)
	if err != nil {
		return err
	}

	if _, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return nil
	}

	return errors.New("invalid Token")
}

func extractJwtToken(r *http.Request) string {
	jwtToken := r.Header.Get("Authorization")

	if len(strings.Split(jwtToken, " ")) == 2 {
		return strings.Split(jwtToken, " ")[1]
	}

	return ""
}

func VerifyJwtTokenSigningMethod(jwtToken *jwt.Token) (interface{}, error) {
	if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong jwtToken signing method! %v", jwtToken.Header["alg"])
	}

	return config.JwtSecret, nil
}

func UserIdFromToken(r *http.Request) (uint64, error) {
	jwtTokenString := extractJwtToken(r)

	jwtToken, err := jwt.Parse(jwtTokenString, VerifyJwtTokenSigningMethod)
	if err != nil {
		return 0, err
	}

	if permissions, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, errors.New("invalid Token")
}
