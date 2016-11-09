// write log.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	println("========== write log task" + "========== ")
	println("begin log")
	writeLog("test wang info", false)
	writeLog("test wang error", true)
	for i := 0; i < 10; i++ {
		// 字符串与整型连接，需要将整型转换为字符串后再操作
		writeLog("test count info:"+strconv.Itoa(i+1), false)
	}
	println("finish log")
	println("========== finish write log task" + "========== ")
}

// 日志记录方法（不支持参数默认值）
func writeLog(str string, isError bool) {
	// 获取当前时间
	now := time.Now()
	userFile := now.Format("log_20060102") + ".txt"
	// 打开文件并追加内容（不存在则创建）
	fout, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0x644)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	// 当前时间格式化
	nowTime := now.Format("2006-01-02 15:04:05")
	var logInfo string
	if isError == true {
		logInfo = "[ERROR][" + nowTime + "] " + str
	} else {
		logInfo = "[INFO][" + nowTime + "] " + str
	}
	fout.WriteString(logInfo + "\r\n")
}
