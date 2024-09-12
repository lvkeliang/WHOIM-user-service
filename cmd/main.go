package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user/userservice"
	"github.com/lvkeliang/WHOIM-user-service/db"
	"github.com/lvkeliang/WHOIM-user-service/handler"
	"log" // 替换为你的包路径
	"net"
	"time"
)

func main() {
	// 初始化 Cassandra 连接
	if err := db.InitCassandra(); err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	// 初始化 Redis
	if err := db.InitRedis(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	// 初始化 etcd 注册中心
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // etcd 地址
	if err != nil {
		log.Fatalf("Failed to create etcd registry: %v", err)
	}

	// 实例化 UserService 实现
	svc := new(handler.UserServiceImpl)

	// 使用 net.ResolveTCPAddr 来创建监听地址
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %v", err)
	}

	// 启动 RPC 服务
	svr := userservice.NewServer(
		svc,
		server.WithRegistry(r),       // 使用 etcd 进行服务注册
		server.WithServiceAddr(addr), // 自定义服务监听地址
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "WHOIM.UserService",
			},
		),
		server.WithReadWriteTimeout(10*time.Second), // 设置读写超时时间
	)

	// 监听和运行服务
	err = svr.Run()
	if err != nil {
		log.Println("UserService failed to start:", err)
	} else {
		log.Println("UserService started successfully!")
	}
}
