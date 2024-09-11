package services

import (
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
	"github.com/lvkeliang/WHOIM-user-service/auth"
	"log"
)

// ValidateToken 验证 JWT 令牌并直接从令牌中返回用户信息
func ValidateToken(token string) (*user.User, error) {
	// 验证并解析 JWT
	claims, err := auth.ValidateJWT(token)
	if err != nil {
		log.Println("Failed to validate token:", err)
		return nil, err
	}

	// 直接从 JWT 中返回用户信息
	return &user.User{
		Id:       claims.UserID,
		Username: claims.Username,
	}, nil
}
