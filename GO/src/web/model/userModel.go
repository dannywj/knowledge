package model

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"strconv"
	db "web/database"
	"web/utils"
)

// User 用户结构体
type User struct {
	Guid   int
	Energy int `bson:"energy_score"`
}

// UserAction 用户行为结构体
type UserAction struct {
	Guid int
	User int
}

// GetUserInfoByGuid 获取用户信息
func GetUserInfoByGuid(guid int) int {
	collection := db.GlobalMgoSessionPlanting.DB("planting").C("user")
	result := User{}
	collection.Find(bson.M{"guid": guid}).One(&result)
	return result.Energy
}

// GetActionCountByDate 根据日期统计行为日志数
func GetActionCountByDate(ymdDate int) (int, int) {
	collection := db.GlobalMgoSessionPlanting.DB("planting").C("user_action201807")
	iter := collection.Find(bson.M{"ymd_date": ymdDate}).Limit(50).Iter()
	result := UserAction{}
	var list []string
	for iter.Next(&result) {
		list = append(list, strconv.Itoa(result.User))
	}
	fmt.Println(list)
	total := len(list)
	ret := utils.SliceUnique(list)
	fmt.Println(ret)
	uniqueCount := len(ret)
	return total, uniqueCount
}
