package services

import (
	"github.com/gocql/gocql"
	"github.com/lvkeliang/WHOIM-user-service/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(username, password, email string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		ID:           gocql.TimeUUID(),
		Username:     username,
		PasswordHash: string(passwordHash),
		Email:        email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return user.Create()
}
