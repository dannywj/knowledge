package business

/*
根据用户id获取用户设备id
DannyWang 2018-07
重构修改版
*/
import (
	"backend/tool"
	"bufio"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 静态常量
const (
	// 日志名称
	LOG_NAME = "device"
	// Mongo 连接配置串
	// MONGO_URL_PLANTING = "10.21.6.39:27111" //test
	// MONGO_URL_USER     = "10.21.6.39:27111" //test

	MONGO_URL_PLANTING = "10.21.6.36:10011" //online
	MONGO_URL_USER     = "10.21.6.36:10010" //online

	// 开启协程数
	NUMBER_GOROUTINES = 50
	// 通道缓冲数
	CHAN_TASK_BUFFER = 20000
	// 用户id文件路径
	USER_ID_FILE_NAME = "test.txt"
)

// 方法路由
var funcs = map[string]interface{}{
	"get_device_from_db":        getAllUserToChan, //数据源
	"get_device_from_db_worker": deviceInfoWorker, //worker方法

	"get_device_from_file":        getAllUserFromFileToChan,
	"get_device_from_file_worker": deviceInfoWorker,
}

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
var GlobalMgoSessionPlanting *mgo.Session
var GlobalMgoSessionUser *mgo.Session

// wg 用来等待程序完成
var wg sync.WaitGroup

// 脚本主入口,根据任务名称执行不同任务
func Run(taskName string) {
	t1 := time.Now()

	// 分配n个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	//创建任务通道,并配置消息缓冲区大小
	taskChan := make(chan interface{}, CHAN_TASK_BUFFER)

	wg.Add(NUMBER_GOROUTINES)

	// 不同任务,不同处理
	PrintLog(fmt.Sprintf("task name:%v", taskName))

	// 验证任务合法性
	if funcs[taskName] == nil {
		PrintLog("invalid task name")
		return
	}

	// 开启协程
	for i := 1; i <= NUMBER_GOROUTINES; i++ {
		go Worker(taskName+"_worker", taskChan, i)
	}

	// 数据源调用
	tool.Call(funcs, taskName, taskChan)

	// 当所有工作都处理完时关闭通道,以便所有 goroutine 退出
	close(taskChan)

	// 等待所有工作完成
	wg.Wait()

	// 计算时间
	elapsed := time.Since(t1)
	PrintLog(fmt.Sprintf("total time:%v", elapsed))
}

// 获取种树用户列表to管道
func getAllUserToChan(taskChan chan interface{}) {
	session := GlobalMgoSessionPlanting //获得session
	defer session.Close()               //释放

	collection := session.DB("planting").C("user")
	//iter := collection.Find(nil).Iter()
	iter := collection.Find(nil).Limit(10000).Iter()
	result := User{}
	i := 1
	for iter.Next(&result) {
		fmt.Printf("guid: %v \n", result.Guid)

		fmt.Printf("================= DB queue->%v =================\n", i)
		i++
		taskChan <- result.Guid
	}
}

// 从文件获取用户列表to管道
func getAllUserFromFileToChan(taskChan chan interface{}) {
	f, _ := os.Open(USER_ID_FILE_NAME)
	defer f.Close()
	r := bufio.NewReader(f)
	i := 1
	for {
		str, err := tool.ReadLineBufio(r)
		if err != nil {
			break
		}
		guid, _ := strconv.Atoi(str)
		fmt.Printf("guid: %v \n", guid)

		fmt.Printf("================= file queue->%v =================\n", i)
		i++
		taskChan <- guid

	}
}

// device任务处理器
func deviceInfoWorker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()
	// 计数器
	i := 1
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		guid, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerID)
			GlobalMgoSessionPlanting.Close()
			return
		}
		// 获取设备信息
		deviceId := getDeviceInfoByGuid(guid.(int))
		fmt.Printf("worker-%d 处理任务：[%v]-{%v}\n", workerID, guid, deviceId)
		tool.WriteLog(LOG_NAME, fmt.Sprintf("%v-%v", guid, deviceId), false)
		// 计数,2000条暂停1s
		i++
		if i%2000 == 0 {
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

// 协程worker转发
func Worker(taskName string, taskChan chan interface{}, workerID int) {
	tool.Call(funcs, taskName, taskChan, workerID)
}

// 默认任务处理器
func defaultWorker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()

	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		data, ok := <-taskChan
		if !ok {
			//通道为空
			//fmt.Printf("worker-%d 所有任务完成\n", workerID)
			return
		}
		fmt.Printf("worker-%d 处理任务：[%v]\n", workerID, data)
	}
}

// 获取设备信息
func getDeviceInfoByGuid(guid int) string {
	collection := GlobalMgoSessionUser.DB("iUser").C("user_device")
	result := UserDevice{}
	// 取最新的一条设备信息
	collection.Find(bson.M{"guid": guid}).Sort("-update_time").Limit(1).One(&result)
	return result.DeviceId
}

// 初始化,系统默认执行
func init() {
	// 初始化Planting Session
	globalMgoSessionPlanting, err := mgo.DialWithTimeout(MONGO_URL_PLANTING, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSessionPlanting = globalMgoSessionPlanting
	GlobalMgoSessionPlanting.SetMode(mgo.Monotonic, true)

	// 初始化User Session
	globalMgoSessionUser, err := mgo.DialWithTimeout(MONGO_URL_USER, 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSessionUser = globalMgoSessionUser
	GlobalMgoSessionUser.SetMode(mgo.Monotonic, true)
}

// 打印&写日志
func PrintLog(str string) {
	fmt.Println(str)
	tool.WriteLog(LOG_NAME, str, false)
}
