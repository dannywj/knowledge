package main

// MongoDB连接并group by获取数据
// 大数据量的表会出现性能问题,查询卡住,原因是排序语句写在前被优先执行了
// 实际放在最后即可解决
import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	//URL = "10.21.6.36:10011" //online
	URL        = "10.21.6.39:27111"
	f_datetime = "2006-01-02 15:04:05"
)

type Money struct {
	Id    string `bson:"_id"`
	Total uint   `bson:"total"`
}

var (
	ress = []*Money{}
)

func main() {
	fmt.Println("========begin task=========")
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Eventual, true)

	db := session.DB("iPayment")  //数据库名称
	collection := db.C("journal") //如果该集合已经存在的话，则直接返回

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("total count: ", countNum)
	//https://www.jankl.com/info/golang%20mgo%20%E4%BD%BF%E7%94%A8%E6%A6%82%E8%A7%88

	aggregate := []bson.M{
		bson.M{
			"$match": bson.M{
				"ctime": bson.M{
					"$gt": StrToTime("2018-07-17 00:00:00"),
					"$lt": StrToTime("2018-07-30 00:00:00"),
				},
				"channel": "TREE_PLANTING_REWARD",
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$day_aggr",
				"total": bson.M{"$sum": "$fee"},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":   1,
				"total": "$total",
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": 1,
			},
		},
	}
	fmt.Println("begin groupby")
	if err = collection.Pipe(aggregate).All(&ress); err != nil {
		//log.Error("c.Pipe() failed(%s)", err)
		fmt.Println("groupby err")
		fmt.Println(err)
		return
	}
	//fmt.Println(ress[0])
	for _, val := range ress {
		fmt.Println(fmt.Sprintf("%v-%v", val.Id, val.Total))
	}
	fmt.Println("========end task=========")
}

//将str转换为时间格式
func StrToTime(st string) time.Time {
	t, _ := time.ParseInLocation(f_datetime, st, time.Local)
	return t
}

/*
	db.getCollection("journal").aggregate(
[
  {
    $match: {
      ctime: {
        $gt: ISODate("2018-07-18T00:00:00.953+08:00"),
        $lt: ISODate("2018-07-19T23:59:59.953+08:00")
      },
      "channel": "TREE_PLANTING_REWARD"
    }
  },
  {
    $group: {
      _id: "$day_aggr",
      total: {
        $sum: "$fee"
      }
    }
  },
  {
    "$sort": {
      "_id": NumberInt(1)
    }
  }
]
)
*/
