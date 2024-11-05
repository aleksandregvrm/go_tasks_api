package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(SecretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected sign in method")
		}
		return SecretKey, nil
	})
	if err != nil {
		return errors.New("Could not parse the token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Invalid token")
	}
	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("Invalid token claims")
	}

	return nil
}
