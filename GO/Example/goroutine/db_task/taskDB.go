package main

/*
	使用协程并发完成数据表的数据迁移
*/
import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "github.com/Go-SQL-Driver/MYSQL"
)

//声明静态变量
const (
	// 开启协程数
	numberGoroutines = 700

	// 通道缓冲数
	chanTaskBuffer = 2000
)

// wg 用来等待程序完成
var wg sync.WaitGroup

var db *sql.DB

// 定义用户地址结构体，方便从DB获取数据后赋值
type userAddress struct {
	user_id int
	address string
}

// userAddress对象实现String方法，方便输出展示
func (s userAddress) String() string {
	return fmt.Sprintf("id:%s address:%s", strconv.Itoa(s.user_id), s.address)
}

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
	db, _ = sql.Open("mysql", "write_user:write_pwd@tcp(172.16.104.207:3307)/mia_test2?charset=utf8")
	/*
			SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
		    SetMaxIdleConns用于设置闲置的连接数。
			设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	*/
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)

	if err := db.Ping(); err != nil {
		fmt.Printf("error ping database: %s", err.Error())
		return
	}
}

func main() {
	fmt.Println("========begin task=========")
	t1 := time.Now()

	// 分配n个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	//创建任务通道
	taskChan := make(chan interface{}, chanTaskBuffer)

	//启动n个worker，异步完成任务
	wg.Add(numberGoroutines)
	for i := 1; i <= numberGoroutines; i++ {
		go worker(taskChan, i)
	}

	//增加任务到任务通道
	getDataToChan(taskChan)

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	close(taskChan)

	// 等待所有工作完成
	wg.Wait()

	elapsed := time.Since(t1)
	fmt.Println("time: ", elapsed)
	fmt.Println("========end task=========")
}

func worker(taskChan chan interface{}, workerID int) {
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
			return
		}
		// fmt.Println(reflect.TypeOf(data)) 检测对象类型
		// 将chan里的data转换为目标类型userAddress
		userAddressObj := data.(userAddress)
		fmt.Printf("worker-%d 开始处理任务：[%s]\n", workerID, userAddressObj)
		insertID := insertDataToTable(userAddressObj)
		if insertID > 0 {
			// 任务处理完成
			fmt.Printf("worker-%d 完成任务：[%s]-[newid:%d]\n", workerID, data, insertID)
		} else {
			fmt.Printf("worker-%d 任务失败：[%s]\n", workerID, data)
		}
	}
}

func getDataToChan(taskChan chan interface{}) {
	rows, _ := db.Query("SELECT user_id, address from user_address order by id limit 100")
	for rows.Next() {
		defer rows.Close()

		var address userAddress

		rows.Scan(&address.user_id, &address.address)
		//rows.Scan()
		fmt.Println(address)
		taskChan <- address //fmt.Sprintf("%s", strconv.Itoa(address.user_id)+address.address)
	}
}

func insertDataToTable(addressInfo userAddress) int64 {
	result, _ := db.Exec(
		"INSERT INTO member_task_final_log_copy (user_id, task_id,processed_reward_name) VALUES (?,?,?)",
		addressInfo.user_id,
		5,
		addressInfo.address,
	)
	insertID, _ := result.LastInsertId()
	return insertID
}
