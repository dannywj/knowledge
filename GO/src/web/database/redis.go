package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var GlobalRedisClient *redis.ClusterClient

func InitRedis() {
	// 连接redis cluster
	client := redis.NewClusterClient(&redis.ClusterOptions{
		//Addrs:    []string{"10.21.6.36:6666", "10.21.6.37:6666", "10.21.6.38:6666", "10.21.6.39:6666", "10.21.6.40:6666", "10.21.6.41:6666"}, //test
		Addrs: []string{"10.80.17.178:6379", "10.80.18.178:6379", "10.80.19.178:6379", "10.80.20.178:6379", "10.80.21.178:6379", "10.80.22.178:6379", "10.80.23.178:6379", "10.80.24.178:6379"},
		//Password: "ifeng666",
		Password: "tv3nIQJgjaSd-",
	})
	// 测试连接是否成功
	pong, conn_err := client.Ping().Result()
	if conn_err != nil {
		panic(conn_err)
	} else {
		fmt.Println("redis cluster conn ok:", pong)
	}
	GlobalRedisClient = client
}
