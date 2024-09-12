package main

import (
	"context"
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
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

// 验证 JWT 令牌
func (s *UserServiceImpl) ValidateToken(ctx context.Context, token string) (*user.User, error) {
	user, err := services.ValidateToken(token)
	if err != nil {
		log.Println("Failed to validate token:", err)
		return nil, err
	}
	return user, nil
}

// 获取用户信息
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, id string) (*user.User, error) {
	userInfo, err := services.GetUserInfo(ctx, id)
	if err != nil {
		log.Println("Failed to get user info:", err)
		return nil, err
	}
	return userInfo, nil
}

// 设置用户设备在线
func (s *UserServiceImpl) SetUserOnline(ctx context.Context, id string, deviceID string, serverAddress string) (bool, error) {
	err := services.SetUserStatus(id, deviceID, serverAddress, "online")
	if err != nil {
		log.Println("Failed to set user online:", err)
		return false, err
	}
	return true, nil
}

// 设置用户设备离线
func (s *UserServiceImpl) SetUserOffline(ctx context.Context, id string, deviceID string) (bool, error) {
	err := services.SetUserStatus(id, deviceID, "", "offline")
	if err != nil {
		log.Println("Failed to set user offline:", err)
		return false, err
	}
	return true, nil
}

// GetUserDevices 获取用户设备在线状态
func (s *UserServiceImpl) GetUserDevices(ctx context.Context, id string) (map[string]*user.UserStatus, error) {
	// 调用 services 层获取设备状态
	devices, err := services.GetUserStatus(id)
	if err != nil {
		log.Println("Failed to get user devices:", err)
		return nil, err
	}

	// 将 map[string]models.UserStatus 转换为 map[string]*user.UserStatus
	thriftDeviceStatuses := make(map[string]*user.UserStatus)
	for deviceID, deviceStatus := range devices {
		thriftDeviceStatuses[deviceID] = &user.UserStatus{
			DeviceID:      deviceStatus.DeviceID,
			ServerAddress: deviceStatus.ServerAddress,
		}
	}

	return thriftDeviceStatuses, nil
}
