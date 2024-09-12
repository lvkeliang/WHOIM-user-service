package models

import (
	"context"
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/lvkeliang/WHOIM-user-service/db"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"log"
	"time"
)

var ctx = context.Background()

var userTable = table.Metadata{
	Name:    "users",
	Columns: []string{"id", "username", "password_hash", "email", "created_at", "updated_at"},
	PartKey: []string{"id"},
}

var users = table.New(userTable)

// UserStatus 表示用户设备的连接状态
type UserStatus struct {
	DeviceID      string // 设备 ID
	ServerAddress string // 设备连接的服务器地址
}

// User 表示用户信息
type User struct {
	ID           gocql.UUID
	Username     string
	PasswordHash string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 插入新用户
func (u *User) Create() error {
	session := db.GetSession()
	stmt, names := qb.Insert(userTable.Name).Columns(userTable.Columns...).ToCql()

	queryx := session.Query(stmt, names).BindStruct(u)
	defer queryx.Release()

	return queryx.Exec()
}

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(userID string) (*User, error) {
	session := db.GetSession()
	var user User

	// 将用户 ID 转换为 UUID
	id, err := gocql.ParseUUID(userID)
	if err != nil {
		return nil, err
	}

	stmt, names := qb.Select(userTable.Name).Where(qb.Eq("id")).Limit(1).ToCql()
	queryx := session.Query(stmt, names).BindMap(qb.M{"id": id})
	defer queryx.Release()

	err = queryx.GetRelease(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(username string) (*User, error) {
	session := db.GetSession()
	var user User

	stmt, names := qb.Select(userTable.Name).Where(qb.Eq("username")).Limit(1).ToCql()
	queryx := session.Query(stmt, names).BindMap(qb.M{"username": username})
	defer queryx.Release()

	err := queryx.Get(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SetUserDeviceOnline 设置用户设备在线状态，记录设备连接的服务器
func SetUserDeviceOnline(userID, deviceID, serverAddress string) error {
	redisClient := db.GetRedisClient()

	userStatus := UserStatus{
		DeviceID:      deviceID,
		ServerAddress: serverAddress,
	}

	// 手动编码为 JSON 字符串
	statusJSON, err := json.Marshal(userStatus)
	if err != nil {
		log.Println("Failed to marshal user status:", err)
		return err
	}

	// 存储用户的设备与服务器信息到 Redis，使用 Redis 哈希表
	err = redisClient.HSet(ctx, "user:"+userID+":devices", deviceID, statusJSON).Err()
	if err != nil {
		log.Println("Failed to set user device online:", err)
		return err
	}

	log.Printf("User %s's device %s connected to server %s", userID, deviceID, serverAddress)
	return nil
}

// GetUserDevices 获取用户所有在线设备及其连接的服务器
func GetUserDevices(userID string) (map[string]UserStatus, error) {
	redisClient := db.GetRedisClient()

	// 获取用户的所有设备与服务器连接信息
	devicesMap, err := redisClient.HGetAll(ctx, "user:"+userID+":devices").Result()
	if err != nil {
		log.Println("Failed to get user devices:", err)
		return nil, err
	}

	// 将获取的结果转换为 UserStatus 类型的映射
	devices := make(map[string]UserStatus)
	for deviceID, data := range devicesMap {
		var status UserStatus
		// 手动解码 JSON 字符串为 UserStatus
		err = json.Unmarshal([]byte(data), &status)
		if err != nil {
			log.Println("Failed to unmarshal user device status:", err)
			return nil, err
		}
		devices[deviceID] = status
	}

	return devices, nil
}

// RemoveUserDevice 设置用户设备离线状态
func RemoveUserDevice(userID, deviceID string) error {
	redisClient := db.GetRedisClient()

	// 移除设备的在线状态
	err := redisClient.HDel(ctx, "user:"+userID+":devices", deviceID).Err()
	if err != nil {
		log.Println("Failed to remove user device:", err)
		return err
	}

	log.Printf("User %s's device %s disconnected", userID, deviceID)
	return nil
}
