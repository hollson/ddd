package main

import (
	"github.com/hollson/kendo/config"
	"github.com/hollson/kendo/domain"
	"github.com/hollson/kendo/infrastructure/db/mysql"
	"github.com/hollson/kendo/presentation"
	"golang.org/x/sync/errgroup"
)

func main() {
	StartUp()
	var eg errgroup.Group

	// 启动Http服务
	eg.Go(func() error {
		return presentation.NewHttpServer().Run(config.AppPort)
	})

	// 启动RPC服务
	eg.Go(func() error {
		return presentation.NewGrpcServer().Run(config.RpcPort)
	})

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}

// StartUp 加载配置文件、初始化数据库等
func StartUp() {
	mysql.InitDB()
	// redis.InitRedis()

	// 启动领域层
	domain.Start()
}
