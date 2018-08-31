package main

import (
	"web/database"
)

func main() {
	// 初始化mongo
	database.InitMongo()
	// 初始化路由配置
	router := initRouter()
	router.Run(":8088")
}

/*
运行: go run *.go
入口地址:http://localhost:8088/redbag/statistics/
*/
