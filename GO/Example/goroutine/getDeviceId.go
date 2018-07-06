package main

/*
根据用户id获取用户设备id
DannyWang 2018-07
*/
import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"runtime"
	"sync"
	"time"
)

// 静态常量
const (
	// 日志名称
	LOG_NAME = "device"
	// Mongo 连接配置串
	MONGO_URL = "10.21.6.39:27111"
	// 开启协程数
	NUMBER_GOROUTINES = 5
	// 通道缓冲数
	CHAN_TASK_BUFFER = 2000
)

// 用户结构体
type User struct {
	Guid   int
	Energy int `bson:"energy_score"`
}

// 用户设备信息结构体
type UserDevice struct {
	Guid       int
	DeviceId   string `bson:"deviceid"`
	UpdateTime string `bson:"update_time"`
}

// mongo连接
var GlobalMgoSession *mgo.Session

// wg 用来等待程序完成
var wg sync.WaitGroup

func main() {
	printLog("========begin task=========")
	t1 := time.Now()

	// 分配n个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	//创建任务通道,并配置消息缓冲区大小
	taskChan := make(chan interface{}, CHAN_TASK_BUFFER)

	//启动n个worker协程，异步完成任务
	wg.Add(NUMBER_GOROUTINES)
	for i := 1; i <= NUMBER_GOROUTINES; i++ {
		go worker(taskChan, i)
	}

	//增加任务(用户id)到任务通道
	getAllUserToChan(taskChan)

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	close(taskChan)

	// 等待所有工作完成
	wg.Wait()

	// 计算时间
	elapsed := time.Since(t1)
	printLog(fmt.Sprintf("total time:%v", elapsed))
	printLog("========end task=========")
}

// 获取用户列表to管道
func getAllUserToChan(taskChan chan interface{}) {
	session := cloneSession() //获得session
	defer session.Close()     //释放

	collection := session.DB("planting").C("user")
	iter := collection.Find(nil).Iter()
	//iter := collection.Find(nil).Limit(50).Iter()
	result := User{}
	for iter.Next(&result) {
		fmt.Printf("guid: %v \n", result.Guid)
		taskChan <- result.Guid
	}
}

// 任务处理器
func worker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		guid, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerID)
			GlobalMgoSession.Close()
			return
		}
		// 获取设备信息
		deviceId := getDeviceInfoByGuid(guid.(int))
		fmt.Printf("worker-%d 开始处理任务：[%v]-{%v}\n", workerID, guid, deviceId)
		writeLog(fmt.Sprintf("%v-%v", guid, deviceId), false)
	}
}

// 获取设备信息
func getDeviceInfoByGuid(guid int) string {
	collection := GlobalMgoSession.DB("iUser").C("user_device")
	result := UserDevice{}
	// 取最新的一条设备信息
	collection.Find(bson.M{"guid": guid}).Sort("-update_time").Limit(1).One(&result)
	return result.DeviceId
}

// 初始化,系统默认执行
func init() {
	globalMgoSession, err := mgo.DialWithTimeout(MONGO_URL, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
}

// 克隆mongo连接
func cloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

// 写日志
func writeLog(str string, showTime bool) {
	// 获取当前时间
	now := time.Now()
	userFile := LOG_NAME + now.Format("_20060102") + ".log"
	// 打开文件并追加内容（不存在则创建）
	fout, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x644)
	os.Chmod(userFile, 0666)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	// 当前时间格式化
	nowTime := now.Format("2006-01-02 15:04:05")
	var logInfo string
	if showTime == true {
		logInfo = "[INFO][" + nowTime + "] " + str
	} else {
		logInfo = str
	}

	fout.WriteString(logInfo + "\r\n")
}

// 打印&写日志
func printLog(str string) {
	fmt.Println(str)
	writeLog(str, false)
}
