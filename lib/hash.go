package lib

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil

}

func CheckPassword(hashPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(dbPassword))
	if err == nil {
		return false
	}
	return true
}
