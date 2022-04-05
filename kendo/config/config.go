// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

// 可以从 yaml、json、toml等格式的配置文件读取
const (
	AppMode  = "dev"
	AppPort  = ":8080"
	RpcPort  = ":8081"
	FileRoot = "./upload/"

	// 数据库配置
	MysqlSource  = "root:123456@tcp(127.0.0.1:3306)/kendo?charset=utf8mb4&loc=Local&parseTime=true&timeout=3s"
	MysqlMaxOpen = 15
	MysqlMaxIdle = 3
	MysqlShowSQL = false

	// Redis配置
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDb       = 0
)
