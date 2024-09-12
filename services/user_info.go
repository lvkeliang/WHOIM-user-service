package services

import (
	"context"
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
	"github.com/lvkeliang/WHOIM-user-service/models"
	"log"
)

// SetUserStatus 设置用户设备在线或离线状态
func SetUserStatus(userID, deviceID, serverAddress, status string) error {
	if status == "online" {
		// 设置设备在线状态
		err := models.SetUserDeviceOnline(userID, deviceID, serverAddress)
		if err != nil {
			log.Println("Failed to set user device online:", err)
			return err
		}
	} else if status == "offline" {
		// 设置设备离线状态
		err := models.RemoveUserDevice(userID, deviceID)
		if err != nil {
			log.Println("Failed to set user device offline:", err)
			return err
		}
	}

	return nil
}

// GetUserStatus 获取用户所有在线设备及其连接的服务器
func GetUserStatus(userID string) (map[string]models.UserStatus, error) {
	devices, err := models.GetUserDevices(userID)
	if err != nil {
		log.Println("Failed to get user devices:", err)
		return nil, err
	}

	if len(devices) == 0 {
		log.Printf("User %s is offline", userID)
	} else {
		log.Printf("User %s is online on devices: %v", userID, devices)
	}

	return devices, nil
}

// 获取用户信息
func GetUserInfo(ctx context.Context, id string) (*user.User, error) {
	// 获取用户基本信息
	modelUser, err := models.GetUserByID(id)
	if err != nil {
		log.Println("Failed to get user info:", err)
		return nil, err
	}

	// 获取用户设备状态
	deviceStatuses, err := GetUserStatus(id)
	if err != nil {
		log.Println("Failed to get user status:", err)
		return nil, err
	}

	// 转换 map[string]models.UserStatus 为 map[string]*user.UserStatus
	thriftDeviceStatuses := make(map[string]*user.UserStatus)
	for deviceID, deviceStatus := range deviceStatuses {
		thriftDeviceStatuses[deviceID] = &user.UserStatus{
			DeviceID:      deviceStatus.DeviceID,
			ServerAddress: deviceStatus.ServerAddress,
		}
	}

	// 将 models.User 转换为 user.User，并添加设备状态信息
	return &user.User{
		Id:       modelUser.ID.String(),
		Username: modelUser.Username,
		Email:    modelUser.Email,
		Status:   thriftDeviceStatuses, // 将状态信息添加到返回的用户结构中
	}, nil
}
