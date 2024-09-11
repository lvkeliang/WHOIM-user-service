package db

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
)

var session gocqlx.Session

// InitCassandra 初始化 Cassandra 连接
func InitCassandra() error {
	// 配置 Cassandra 集群
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "whoim"
	cluster.Consistency = gocql.Quorum

	// 创建 gocql.Session
	s, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	// 包装成 gocqlx.Session
	session, err = gocqlx.WrapSession(s, err)
	if err != nil {
		return err
	}

	log.Println("Cassandra connection established")
	return nil
}

// GetSession 获取 Cassandra Session
func GetSession() gocqlx.Session {
	// 确保 session 已初始化
	if session.Session == nil {
		log.Fatal("Cassandra session is not initialized. Call InitCassandra first.")
	}
	return session
}
