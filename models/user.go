package models

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/lvkeliang/WHOIM-user-service/db"
	"github.com/lvkeliang/WHOIM-user-service/utils"
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

	// 使用 gocqlx.Session.Query 代替 gocqlx.Query
	queryx := session.Query(stmt, names).BindStruct(u)
	defer queryx.Release()

	return queryx.Exec()
}

func GetUserByUsername(username string) (*User, error) {
	session := db.GetSession()
	var user User
	stmt, names := qb.Select(userTable.Name).Where(qb.Eq("username")).Limit(1).ToCql()

	// 使用 gocqlx.Session.Query 代替 gocqlx.Query
	queryx := session.Query(stmt, names).BindMap(qb.M{"username": username})
	defer queryx.Release()

	err := queryx.Get(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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

// SetUserStatus 设置用户在线或离线状态，使用 Redis Bitmap
func SetUserStatus(userID string, status string) error {
	// 将 UUID 转换为整数，用于 Bitmap 的偏移量
	idInt, err := utils.UUIDToInt(userID)
	if err != nil {
		log.Println("Failed to convert userID:", err)
		return err
	}

	redisClient := db.GetRedisClient()

	// 设置用户的在线/离线状态到 Redis Bitmap
	if status == "online" {
		err = redisClient.SetBit(ctx, "user:status", idInt, 1).Err()
		if err != nil {
			log.Println("Failed to set user online status in Redis:", err)
			return err
		}
	} else {
		err = redisClient.SetBit(ctx, "user:status", idInt, 0).Err()
		if err != nil {
			log.Println("Failed to set user offline status in Redis:", err)
			return err
		}
	}

	log.Printf("User %s status set to %s", userID, status)
	return nil
}

// GetUserStatus 从 Redis 获取用户在线状态
func GetUserStatus(userID string) (string, error) {
	// 将 UUID 转换为整数
	idInt, err := utils.UUIDToInt(userID)
	if err != nil {
		log.Println("Failed to convert userID:", err)
		return "", err
	}

	fmt.Println("UUIDINT: ", idInt)

	redisClient := db.GetRedisClient()

	// 从 Redis Bitmap 中获取用户的在线状态
	status, err := redisClient.GetBit(ctx, "user:status", idInt).Result()
	if err != nil {
		log.Println("Failed to get user status from Redis:", err)
		return "", err
	}

	if status == 1 {
		return "online", nil
	}
	return "offline", nil
}
