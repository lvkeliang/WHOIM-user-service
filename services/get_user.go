package services

import (
	"github.com/lvkeliang/WHOIM-user-service/models"
	"log"
)

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(username string) (*models.User, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println("Failed to get user by username:", err)
		return nil, err
	}
	return user, nil
}

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(userID string) (*models.User, error) {
	user, err := models.GetUserByID(userID)
	if err != nil {
		log.Println("Failed to get user by ID:", err)
		return nil, err
	}
	return user, nil
}
