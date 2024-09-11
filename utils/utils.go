package utils

import (
	"crypto/md5"
	"math/big"
)

// UUIDToInt 将 UUID 转换为整数偏移量
func UUIDToInt(uuidStr string) (int64, error) {
	hasher := md5.New()
	hasher.Write([]byte(uuidStr))
	hashBytes := hasher.Sum(nil)

	// 取前 8 字节作为整数
	intValue := new(big.Int).SetBytes(hashBytes[:4]).Int64()
	return intValue, nil
}
