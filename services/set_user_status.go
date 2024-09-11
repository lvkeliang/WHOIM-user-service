package services

import (
	"github.com/lvkeliang/WHOIM-user-service/models"
	"log"
)

// SetUserStatus 设置用户在线或离线状态，使用用户的 ID
func SetUserStatus(userID, status string) error {
	// 查找用户
	_, err := models.GetUserByID(userID)
	if err != nil {
		log.Println("Failed to find user:", err)
		return err
	}

	// 更新 Redis 中的用户在线状态
	err = models.SetUserStatus(userID, status)
	if err != nil {
		log.Println("Failed to set user status in Redis:", err)
		return err
	}

	return nil
}

// GetUserStatus 获取用户在线状态，使用用户的 ID
func GetUserStatus(userID string) (string, error) {
	// 从 Redis 中获取用户在线状态
	status, err := models.GetUserStatus(userID)
	if err != nil {
		log.Println("Failed to get user status from Redis:", err)
		return "", err
	}

	return status, nil
}
