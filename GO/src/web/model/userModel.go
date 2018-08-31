package model

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"runtime"
	"sort"
	"strconv"
	"strings"
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

func GetActionLogByDate(beginDate int, endDate int) []string {
	fmt.Println("========begin task=========")
	// 分配n个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	//创建任务通道
	taskChan := make(chan string, chanTaskBuffer)
	//创建结果通道
	resultChan := make(chan string, chanTaskBuffer)

	// 启动n个worker，异步完成任务
	i := 1
	for day := beginDate; day <= endDate; day++ {
		go actionWorker(taskChan, i, resultChan)
		i++
	}

	wg.Add(i - 1)
	//增加任务到任务通道
	for day := beginDate; day <= endDate; day++ {
		taskChan <- strconv.Itoa(day)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	close(taskChan)

	// 等待所有工作完成
	wg.Wait()

	fmt.Println("--finish data--")

	// 任务执行完毕,关闭结果通道
	close(resultChan)

	// 处理结果管道的数据
	// 保存结果管道里的信息到切片
	var list []string
	for {
		//使用两个变量接受返回值，如果ok为false，则m为零值，但是不会报错。
		re, ok := <-resultChan
		if !ok {
			//通道为空,处理结束
			break
		}
		list = append(list, re)
	}

	// 结果排序
	sort.Strings(list)

	// 循环结果数据,元素字符串拆分,返回结果
	for _, value := range list {
		listItemArr := strings.Split(value, "_")
		fmt.Printf("date:%v fee:%v\n", listItemArr[0], listItemArr[1])
	}
	fmt.Println("========end task=========")
	return list
}

func actionWorker(taskChan chan string, workerId int, resultChan chan string) {
	// 通知函数已经返回
	defer wg.Done()
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则m为零值，但是不会报错。
		task, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerId)
			return
		}
		fmt.Printf("worker-%d 开始处理任务：[%s]\n", workerId, task)

		// 处理任务
		collection := db.GlobalMgoSessionPlanting.DB("planting").C("user_action201808")

		countNum, _ := collection.Count()
		fmt.Println(countNum)
		ymd_date, _ := strconv.Atoi(task)
		iter := collection.Find(bson.M{"ymd_date": ymd_date}).Iter()
		fmt.Println(task)
		result := UserAction{}
		var list []string
		for iter.Next(&result) {
			fmt.Printf("for user:%v\n", result.User)
			list = append(list, strconv.Itoa(result.User))
		}
		fmt.Println(list)
		total := len(list)
		ret := utils.SliceUnique(list)
		fmt.Println(ret)
		uniqueCount := len(ret)

		// 数据结果push到结果通道
		resultChan <- fmt.Sprintf("%v_%v_%v", task, total, uniqueCount)
		// 任务处理完成
		fmt.Printf("worker-%d 完成任务：[%s]\n", workerId, task)

		// result := Redbag{}
		// cond := bson.M{
		// 	"channel":  "TREE_PLANTING_REWARD",
		// 	"day_aggr": task,
		// }
		// iter := collection.Find(cond).Iter()
		// var totalFee float32
		// for iter.Next(&result) {
		// 	// 输出处理信息
		// 	fmt.Printf("guid: %v, fee:%v\n", result.GUID, result.Fee)
		// 	totalFee += result.Fee
		// }
		// //fmt.Printf("%v-total fee: %v \n", task, totalFee/100)

		// // 数据结果push到结果通道
		// resultChan <- fmt.Sprintf("%v_%v", task, totalFee/100)
		// // 任务处理完成
		// fmt.Printf("worker-%d 完成任务：[%s]\n", workerId, task)
	}
}
