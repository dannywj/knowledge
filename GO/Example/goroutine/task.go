package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//声明静态变量
const (
	// 开启协程数
	numberGoroutines = 4
	// 总任务数
	totalTask = 10
	// 通道缓冲数
	chanTaskBuffer = 10
)

// wg 用来等待程序完成
var wg sync.WaitGroup

func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("========begin task=========")

	//创建任务通道
	taskChan := make(chan string, chanTaskBuffer)

	//启动n个worker，异步完成任务
	wg.Add(numberGoroutines)
	for i := 1; i <= numberGoroutines; i++ {
		go worker(taskChan, i)
	}

	//增加任务到任务通道
	for post := 1; post <= totalTask; post++ {
		taskChan <- fmt.Sprintf("task %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	close(taskChan)

	// 等待所有工作完成
	wg.Wait()
}
func worker(taskChan chan string, workerId int) {
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

		// 模拟处理任务
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 任务处理完成
		fmt.Printf("worker-%d 完成任务：[%s]\n", workerId, task)
	}
}
