package main

import (
	"context"
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
	"github.com/lvkeliang/WHOIM-user-service/services"
	"log"
)

type UserServiceImpl struct{}

// 实现 Login 方法，返回 string 和 error
func (s *UserServiceImpl) Login(ctx context.Context, username string, password string) (string, error) {
	token, err := services.Login(username, password)
	if err != nil {
		log.Println("Login failed:", err)
		return "", err
	}
	return token, nil
}

// 实现 Register 方法，返回 bool 和 error
func (s *UserServiceImpl) Register(ctx context.Context, username string, password string, email string) (bool, error) {
	err := services.Register(username, password, email)
	if err != nil {
		log.Println("Failed to register user:", err)
		return false, err
	}
	return true, nil
}

// 实现 GetUserInfo 方法，返回 *user.User 和 error
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, username string) (*user.User, error) {
	modelUser, err := services.GetUserByUsername(username)
	if err != nil {
		log.Println("Failed to get user info:", err)
		return nil, err
	}

	// 转换 models.User 为 user.User
	return &user.User{
		Id:       modelUser.ID.String(),
		Username: modelUser.Username,
		Email:    modelUser.Email,
		Status:   modelUser.Status,
	}, nil
}

// 实现 SetUserOnline 方法，返回 bool 和 error
func (s *UserServiceImpl) SetUserOnline(ctx context.Context, username string) (bool, error) {
	err := services.SetUserStatus(username, "online")
	if err != nil {
		log.Println("Failed to set user online:", err)
		return false, err
	}
	return true, nil
}

// 实现 SetUserOffline 方法，返回 bool 和 error
func (s *UserServiceImpl) SetUserOffline(ctx context.Context, username string) (bool, error) {
	err := services.SetUserStatus(username, "offline")
	if err != nil {
		log.Println("Failed to set user offline:", err)
		return false, err
	}
	return true, nil
}
