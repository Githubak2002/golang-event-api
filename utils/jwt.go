package utils

import (
	"errors"
	"fmt"
	// "log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(emial string, userId int64) (string, error) {

	// getting JWT_secret from env
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET not set in environment")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  emial,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
		// "exp":    time.Now().Add(time.Minute * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func ValidToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// getting JWT_secret from env
		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			return "", fmt.Errorf("JWT_SECRET not set in environment")
		}
		
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected SignIn Method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// log.Println("Error: ",err.Error())
		return 0, errors.New("could not Parse token")
	}



	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
