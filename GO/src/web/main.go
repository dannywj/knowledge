package main

import (
	"web/database"
)

func main() {
	// 初始化mongo
	database.InitMongo()
	database.InitRedis()
	// 初始化路由配置
	router := initRouter()
	router.Run(":9999")
}

/*
运行: go run *.go
入口地址:http://localhost:8088/redbag/statistics/
*/
