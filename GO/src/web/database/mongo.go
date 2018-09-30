package database

import (
	"labix.org/v2/mgo"
	"time"
)

const (
	// 日志名称
	//LOG_NAME = "energy"
	// Mongo 连接配置串
	//MONGO_URL_PLANTING = "10.21.6.39:27111" //test
	// MONGO_URL_USER     = "10.21.6.39:27111" //test

	//MONGO_URL_PLANTING = "10.21.6.36:10011" //online
	MONGO_URL_PLANTING = "10.80.22.148:10002" //online
)

var GlobalMgoSessionPlanting *mgo.Session

// 初始化mongodb
func InitMongo() {
	// 初始化Planting Session
	globalMgoSessionPlanting, err := mgo.DialWithTimeout(MONGO_URL_PLANTING, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSessionPlanting = globalMgoSessionPlanting
	GlobalMgoSessionPlanting.SetMode(mgo.Monotonic, true)

	// // 初始化User Session
	// globalMgoSessionUser, err := mgo.DialWithTimeout(MONGO_URL_USER, 10*time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// GlobalMgoSessionUser = globalMgoSessionUser
	// GlobalMgoSessionUser.SetMode(mgo.Monotonic, true)
}
