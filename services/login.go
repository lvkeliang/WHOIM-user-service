package services

import (
	"github.com/lvkeliang/WHOIM-user-service/auth"
	"github.com/lvkeliang/WHOIM-user-service/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (string, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
