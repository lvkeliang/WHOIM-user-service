package services

import (
	"github.com/lvkeliang/WHOIM-user-service/models"
	"log"
)

// SetUserStatus 设置用户在线或离线状态
func SetUserStatus(username, status string) error {
	// 查找用户
	user, err := models.GetUserByUsername(username)
	if err != nil {
		log.Println("Failed to find user:", err)
		return err
	}

	// 更新用户状态
	user.Status = status

	// 保存到数据库
	err = user.UpdateStatus(status)
	if err != nil {
		log.Println("Failed to update user status:", err)
		return err
	}

	return nil
}
