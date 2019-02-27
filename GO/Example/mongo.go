package main

// MongoDB连接并获取数据
import (
	//"encoding/hex"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	//ID     []byte        `bson:"_id"` //也可以用这种方式获取,但是要用"encoding/hex" 解析
	IDNew  bson.ObjectId `bson:"_id"` //推荐的_id字段获取方式
	GUID   int
	ENERGY int `bson:"energy_score"` //表示mongodb数据库中对应的字段名称
}

const (
	URL = "10.21.6.39:27111"
)

func main() {
	fmt.Println("========begin task=========")
	session, err := mgo.Dial(URL) //连接数据库
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
	//iter := collection.Find(nil).Limit(20).Iter()
	iter := collection.Find(bson.M{"guid": 123}).Limit(20).Iter()

	for iter.Next(&result) {
		fmt.Printf("guid: %v, energy:%v\n", result.GUID, result.ENERGY)
		// id := hex.EncodeToString(result.ID)
		// fmt.Printf("_id: %v\n", id)
		fmt.Printf("_id new: %v\n", result.IDNew)        //ObjectIdHex("5ae442265142a3df768b4569") 类型的值
		hexid := fmt.Sprintf("%x", string(result.IDNew)) //ObjectIdHex 转string
		fmt.Println(hexid)
	}
	fmt.Println("========end task=========")
}
