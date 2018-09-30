package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/gomail.v2"
	"strconv"
	"time"
)

var GlobalRedisClient *redis.ClusterClient

func main() {
	fmt.Println("========begin task=========")
	dbPlantingQueue := "planting:task:queue"       // 监控的redis队列key
	dbAlarmPlantingCount := "alarm_planting_count" // 报警计数器key
	messageStoreLimit := 50                        // 消息堆积上限
	InitRedis()
	itemCount, _ := GlobalRedisClient.LLen(dbPlantingQueue).Result()
	fmt.Println("planting_task_queue item count:", itemCount)
	if int(itemCount) > messageStoreLimit {
		alarmCount, _ := GlobalRedisClient.Get(dbAlarmPlantingCount).Result()
		alarmCountNum, _ := strconv.Atoi(alarmCount)
		if alarmCountNum <= 5 { // 每天最多发5次报警邮件
			sendAlarmEmail(itemCount)
			GlobalRedisClient.Incr(dbAlarmPlantingCount)
			GlobalRedisClient.Expire(dbAlarmPlantingCount, time.Duration(3600*12)*time.Second)
		} else {
			fmt.Println("send email limit")
		}
	} else {
		fmt.Println("planting_task_queue status OK")
	}
	fmt.Println("========end task=========")
}

func sendAlarmEmail(itemCount int64) {
	m := gomail.NewMessage()
	m.SetHeader("From", "dannywj@live.cn")
	m.SetHeader("To", "wangjue1@ifeng.com")
	m.SetAddressHeader("Cc", "danny__wj@163.com", "Danny")
	m.SetHeader("Subject", "Alarm-Planting Queue Store!")
	m.SetBody("text/html", "Planting Tree queue store item:"+strconv.FormatInt(itemCount, 10))
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.office365.com", 587, "dannywj@live.cn", "juewang137")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("send success")
}

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
