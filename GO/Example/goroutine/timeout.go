package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=======begin task========")
	// 定义任务通道
	taskChan := make(chan int)
	// 超时通道,指定时间后会往该通道写入一个当前时间进去
	// 超时通道也可以自定义一个,在循环超时里更灵活,自定义示例:
	/*
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(time.Second * 2)
			timeout <- false
		}()
	*/
	//timeout := time.After(time.Second * 2)

	// 异步执行任务
	go doTask(taskChan)

	// 通道选择,没有default分支,程序会一直阻塞等待有任何一个case的路径中有值才会执行里面的逻辑
	// 一旦超时通道中有值了(到了超时时间),则会超时并退出

	select {
	case <-taskChan:
		fmt.Println("task finished!")
	case <-timeout:
		fmt.Println("task timeout!")
	}
}

// 执行任务,假设任务执行需要3s的时间
func doTask(ch chan int) {
	time.Sleep(time.Second * 3)
	// 任务执行完毕向通道写入标识
	ch <- 1
	fmt.Println("task ok!")
}
