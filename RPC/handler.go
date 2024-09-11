package main

import (
	"context"
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
	"github.com/lvkeliang/WHOIM-user-service/models"
	"github.com/lvkeliang/WHOIM-user-service/services"
	"log"
)

type UserServiceImpl struct{}

// NewUserServiceImpl 创建并返回 UserServiceImpl 的实例
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

// 注册
func (s *UserServiceImpl) Register(ctx context.Context, username, password, email string) (bool, error) {
	err := services.Register(username, password, email)
	if err != nil {
		log.Println("Failed to register user:", err)
		return false, err
	}
	return true, nil
}

// 登录
func (s *UserServiceImpl) Login(ctx context.Context, username, password string) (string, error) {
	token, err := services.Login(username, password)
	if err != nil {
		log.Println("Failed to login:", err)
		return "", err
	}
	return token, nil
}

// 获取用户信息
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, id string) (*user.User, error) {
	modelUser, err := models.GetUserByID(id)
	if err != nil {
		log.Println("Failed to get user info:", err)
		return nil, err
	}

	// 获取用户状态
	status, err := services.GetUserStatus(id)
	if err != nil {
		log.Println("Failed to get user status:", err)
		return nil, err
	}

	// 将 models.User 转换为 user.User，并添加状态信息
	return &user.User{
		Id:       modelUser.ID.String(),
		Username: modelUser.Username,
		Email:    modelUser.Email,
		Status:   status,
	}, nil
}

// 设置用户在线
func (s *UserServiceImpl) SetUserOnline(ctx context.Context, id string) (bool, error) {
	err := services.SetUserStatus(id, "online")
	if err != nil {
		log.Println("Failed to set user online:", err)
		return false, err
	}
	return true, nil
}

// 设置用户离线
func (s *UserServiceImpl) SetUserOffline(ctx context.Context, id string) (bool, error) {
	err := services.SetUserStatus(id, "offline")
	if err != nil {
		log.Println("Failed to set user offline:", err)
		return false, err
	}
	return true, nil
}

// GetUserStatus implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserStatus(ctx context.Context, id string) (resp string, err error) {
	// 从 Redis 获取用户状态
	status, err := models.GetUserStatus(id)
	if err != nil {
		log.Println("Failed to get user status:", err)
		return "", err
	}

	return status, nil
}
