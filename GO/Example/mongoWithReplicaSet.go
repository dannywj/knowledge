package main

// MongoDB with ReplicaSet 连接并获取数据
// https://docs.objectrocket.com/mongodb_go_examples.html
import (
	"fmt"
	"gopkg.in/mgo.v2" // 使用特定的mgo包
	"labix.org/v2/mgo/bson"
)

type Person struct {
	GUID   int
	ENERGY int `bson:"energy_score"` //表示mongodb数据库中对应的字段名称
}

// DB配置信息
const (
	// Username       = "YOUR_USERNAME"
	// Password       = "YOUR_PASSWORD"
	// Database       = "planting"
	ReplicaSetName = "iclient_02" // 副本集名称
)

func main() {
	fmt.Println("========begin task=========")
	// server list
	Host := []string{
		"10.80.22.148:10002",
		"10.80.23.148:10002",
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
		// Username: Username,
		// Password: Password,
		// Database:       Database,
		ReplicaSetName: ReplicaSetName,
	})

	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("planting") //数据库名称
	collection := db.C("user")   //如果该集合已经存在的话，则直接返回

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("total user count: ", countNum)

	//*****查询多条数据*******
	result := Person{}
	iter := collection.Find(bson.M{"guid": 97018491}).Limit(5).Iter()
	for iter.Next(&result) {
		fmt.Printf("guid: %v, energy:%v\n", result.GUID, result.ENERGY)
	}
	fmt.Println("========end task=========")
}
