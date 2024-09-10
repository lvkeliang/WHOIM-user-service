package models

import (
	"github.com/gocql/gocql"
	"github.com/lvkeliang/WHOIM-user-service/db"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"time"
)

var userTable = table.Metadata{
	Name:    "users",
	Columns: []string{"id", "username", "password_hash", "email", "status", "created_at", "updated_at"},
	PartKey: []string{"id"},
}

var users = table.New(userTable)

type User struct {
	ID           gocql.UUID
	Username     string
	PasswordHash string
	Email        string
	Status       string
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

// UpdateStatus 更新用户状态
func (u *User) UpdateStatus(status string) error {
	session := db.GetSession()

	u.Status = status
	u.UpdatedAt = time.Now()

	stmt, names := qb.Update(userTable.Name).
		Set("status", "updated_at").
		Where(qb.Eq("username")).ToCql()

	// 执行更新操作
	queryx := session.Query(stmt, names).BindStruct(u)
	defer queryx.Release()

	return queryx.Exec()
}
