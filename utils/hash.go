package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(bytes), err
}

func CheckHashPassword(passowrd, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passowrd))
	return err == nil
}