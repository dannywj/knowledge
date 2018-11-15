package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	re1 := GetDataFromApi() // 使用chan计数
	//re2 := GetDataFromApi2() //使用wg计数
	fmt.Println(re1)
}
func GetDataFromApi() []string {
	chs := make(chan int)
	count := 0
	var result []string
	fmt.Println("begin GetDataFromApi")
	go func() {
		re := test1()
		result = append(result, re)
		count++
		chs <- count
	}()
	go func() {
		re := test2()
		result = append(result, re)
		count++
		chs <- count
	}()

	for c := range chs {
		if c == 2 {
			close(chs)
		}
	}
	fmt.Println("result:", result)
	fmt.Println("finish GetDataFromApi")
	return (result)
}

func test1() string {
	fmt.Println("test 1 begin")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("test 1 sleep 5s")
	return "test 1 ok"
}

func test2() string {
	fmt.Println("test 2 begin")
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("test 2 sleep 2s")
	return "test 2 ok"
}

func GetDataFromApi2() []string {
	var wg sync.WaitGroup
	//添加2个任务
	wg.Add(2)

	var result []string
	fmt.Println("begin GetDataFromApi2")
	go func() {
		// 完成任务
		defer wg.Done()
		re := test1()
		result = append(result, re)
	}()
	go func() {
		defer wg.Done()
		re := test2()
		result = append(result, re)
	}()
	// 等待全部完成
	wg.Wait()

	fmt.Println("result:", result)
	fmt.Println("finish GetDataFromApi2")
	return (result)
}
