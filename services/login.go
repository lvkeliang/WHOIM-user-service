package services

import (
	"github.com/lvkeliang/WHOIM-user-service/auth"
	"github.com/lvkeliang/WHOIM-user-service/models"
	"golang.org/x/crypto/bcrypt"
)

// Login 使用用户名和密码进行登录，并返回带有用户 ID 和用户名的 JWT
func Login(username, password string) (string, error) {
	// 根据用户名查找用户
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	// 比较密码哈希
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	// 生成包含用户 ID 和用户名的 JWT
	token, err := auth.GenerateJWT(user.ID.String(), user.Username)
	if err != nil {
		return "", err
	}

	// 返回生成的 JWT
	return token, nil
}
