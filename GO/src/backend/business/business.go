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
	"github.com/go-redis/redis"
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
	LOG_NAME = "energy"
	// Mongo 连接配置串
	// MONGO_URL_PLANTING = "10.21.6.39:27111" //test
	// MONGO_URL_USER     = "10.21.6.39:27111" //test

	MONGO_URL_PLANTING = "10.21.6.36:10011" //online
	MONGO_URL_USER     = "10.21.6.36:10010" //online

	// 开启协程数
	NUMBER_GOROUTINES = 30
	// 通道缓冲数
	CHAN_TASK_BUFFER = 50000
	// 用户id文件路径
	USER_ID_FILE_NAME = "alluser_20180712.log"
)

// redis-test
// var RedisServerList = []string{"10.21.6.36:6666", "10.21.6.37:6666", "10.21.6.38:6666", "10.21.6.39:6666", "10.21.6.40:6666", "10.21.6.41:6666"}
// var RedisPassword="ifeng666"
// redis-cluster2
var RedisServerList = []string{"10.80.17.178:6379", "10.80.18.178:6379", "10.80.19.178:6379", "10.80.20.178:6379", "10.80.21.178:6379", "10.80.22.178:6379", "10.80.23.178:6379", "10.80.24.178:6379"}
var RedisPassword = "tv3nIQJgjaSd-"

// 方法路由
var funcs = map[string]interface{}{
	"get_device_from_db":        getAllUserToChan, //数据源
	"get_device_from_db_worker": deviceInfoWorker, //worker方法

	"get_device_from_file":        getAllUserFromFileToChan,
	"get_device_from_file_worker": deviceInfoWorker,

	"check_user_energy":        getAllUserFromFileToChan,
	"check_user_energy_worker": userEnergyWorker,

	"gen_user_list_file":        getAllUserToChan,
	"gen_user_list_file_worker": genUserListWorker,
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

// redis连接
var RedisClient *redis.ClusterClient

// wg 用来等待程序完成
var wg sync.WaitGroup

// 执行限速器
var lr tool.LimitRate

// 脚本主入口,根据任务名称执行不同任务
func Run(taskName string) {
	// 计时器
	beginTime := time.Now()

	// 每秒限速配置
	lr.SetRate(-1) // 每秒执行任务数配置, -1 不限速

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
	PrintLog(fmt.Sprintf("total time:%v", time.Since(beginTime)))
}

// 获取种树用户列表to管道
func getAllUserToChan(taskChan chan interface{}) {
	session := GlobalMgoSessionPlanting //获得session
	defer session.Close()               //释放

	collection := session.DB("planting").C("user")
	iter := collection.Find(nil).Iter()
	//iter := collection.Find(nil).Limit(20000).Iter()
	result := User{}
	i := 1
	for iter.Next(&result) {
		fmt.Printf("================= DB queue->%v =======guid: %v==========\n", i, result.Guid)
		i++
		taskChan <- result
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
		// 构造user对象
		userObj := User{}
		userObj.Guid = guid
		fmt.Printf("================= file queue->%v =================\n", i)
		i++
		taskChan <- userObj

	}
}

// device任务处理器
func deviceInfoWorker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		data, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerID)
			GlobalMgoSessionPlanting.Close()
			return
		}
		// 限速执行
		if lr.Limit() {
			// 获取设备信息
			userObj := data.(User)
			deviceId := getDeviceInfoByGuid(userObj.Guid)
			fmt.Printf("worker-%d 处理任务：[%v]-{%v}\n", workerID, userObj.Guid, deviceId)
			fmt.Printf("=======================↓↓↓↓↓↓↓[%v]↓↓↓↓↓↓↓=======================\n", len(taskChan)) //该长度值的计算在并发执行过程中不准确,会有重叠的现象发生,只能作为大概的参考
			tool.WriteLog(LOG_NAME, fmt.Sprintf("%v-%v", userObj.Guid, deviceId), false)
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

// 获取用户信息
func getUserInfoByGuid(guid int) int {
	collection := GlobalMgoSessionPlanting.DB("planting").C("user")
	result := User{}
	collection.Find(bson.M{"guid": guid}).One(&result)
	return result.Energy
}

// 初始化,系统默认执行
func init() {
	initMongo()
	initRedis()
}

// 初始化mongodb
func initMongo() {
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

// 初始化redis
func initRedis() {
	RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    RedisServerList,
		Password: RedisPassword,
	})
}

// 打印&写日志
func PrintLog(str string) {
	fmt.Println(str)
	tool.WriteLog(LOG_NAME, str, false)
}

// 获取用户能量(from redis)
func getUserEnergyFromRedis(guid int) int {
	key := fmt.Sprintf("planting:user:info:%v", guid)
	val, err := RedisClient.HGetAll(key).Result()
	if err != nil {
		panic(err)
	}
	if len(val) > 0 {
		energy_score, _ := strconv.Atoi(val["energy_score"])
		return energy_score
	}
	return 0
}

// 能量比较worker
func userEnergyWorker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()
	// 计数器
	i := 1
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		data, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerID)
			return
		}
		// 获取设备信息
		userObj := data.(User)
		energyRedis := getUserEnergyFromRedis(userObj.Guid)
		energyMongo := getUserInfoByGuid(userObj.Guid)
		fmt.Printf("worker-%d 处理任务：[%v]-{redis:%v  mongo:%v}\n", workerID, userObj.Guid, energyRedis, energyMongo)
		fmt.Printf("=======================↓↓↓↓↓↓↓[%v]↓↓↓↓↓↓↓=======================\n", len(taskChan))
		if tool.CalcIntAbs(energyRedis-energyMongo) != 0 {
			tool.WriteLog(LOG_NAME, fmt.Sprintf("%v:%v-%v", userObj.Guid, energyRedis, energyMongo), false)
		}

		// 计数,2000条暂停1s
		i++
		if i%2000 == 0 {
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

// 生成用户列表文件worker
func genUserListWorker(taskChan chan interface{}, workerID int) {
	// 通知函数已经返回
	defer wg.Done()
	// 计数器
	i := 1
	// 循环接收工作
	for {
		//使用两个变量接受返回值，如果ok为false，则taskChan为零值，但是不会报错。
		//用于检测通道是否关闭或为空
		data, ok := <-taskChan
		if !ok {
			//通道为空
			fmt.Printf("worker-%d 所有任务完成\n", workerID)
			GlobalMgoSessionPlanting.Close()
			return
		}
		userObj := data.(User)
		tool.WriteLog(LOG_NAME, fmt.Sprintf("%v", userObj.Guid), false)
		// 计数,2000条暂停1s
		i++
		if i%2000 == 0 {
			time.Sleep(time.Millisecond * 1000)
		}
	}
}
