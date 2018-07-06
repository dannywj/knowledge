package main

// 连接redis cluster示例
// https://github.com/go-redis/redis
import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("========begin task=========")
	// 连接redis cluster
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"10.21.6.36:6666", "10.21.6.37:6666", "10.21.6.38:6666", "10.21.6.39:6666", "10.21.6.40:6666", "10.21.6.41:6666"},
		Password: "ifeng666",
	})
	// 测试连接是否成功
	pong, conn_err := client.Ping().Result()
	if conn_err != nil {
		panic(conn_err)
	} else {
		fmt.Println("redis cluster conn ok:", pong)
	}

	err := client.Set("key_test1", "val1", 0).Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("key_test1 set ok")
	}

	val, err := client.Get("key_test1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("get key_test1:", val)
	fmt.Println("========end task=========")
}
