package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "JWTSuperSecretKeyByAk"

func GenerateToken(emial string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : emial,
		"userId" : userId,
		"exp": time.Now().Add(time.Hour*2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}


func ValidToken (token string) (int64, error) {
	parsedToken, err :=	jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_,ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected SignIn Method")
		}
		return []byte(secretKey), nil
	})
	if err != nil{
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