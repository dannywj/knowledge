package main

// 连接redis服务器
import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("========begin task=========")
	client := redis.NewClient(&redis.Options{
		Addr:     "10.21.6.37:6666",
		Password: "ifeng666", // no password set
		DB:       0,          // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key_test", "222", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key_test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
