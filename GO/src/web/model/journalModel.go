package model

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"strconv"
	db "web/database"
	"web/utils"
)

type Money struct {
	Id    string  `bson:"_id"`
	Total float32 `bson:"total"` //需要除100转换 因此定义成浮点型
}

var (
	ress = []*Money{}
)

func GetJournalGroupByDate(beginDate int, endDate int) []string {

	aggregate := []bson.M{
		bson.M{
			"$match": bson.M{
				"ctime": bson.M{
					"$gt": utils.StrToTime("2006-01-02 15:04:05", fmt.Sprintf("%v 00:00:00", utils.DayAggrToTime(strconv.Itoa(beginDate)))),
					"$lt": utils.StrToTime("2006-01-02 15:04:05", fmt.Sprintf("%v 23:59:59", utils.DayAggrToTime(strconv.Itoa(endDate)))),
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
	var list []string
	collection := db.GlobalMgoSessionPlanting.DB("iPayment").C("journal")
	var err error
	if err = collection.Pipe(aggregate).All(&ress); err != nil {
		fmt.Println("groupby err")
		fmt.Println(err)
		return list
	}

	for _, val := range ress {
		fmt.Println(fmt.Sprintf("%v-%v", val.Id, val.Total))
		list = append(list, fmt.Sprintf("%v_%v", utils.DayAggrToTime(val.Id), val.Total/100))
	}
	return list
}
