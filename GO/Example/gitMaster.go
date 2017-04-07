// git 日常发布分支生成
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	println("========== git release generate begin" + "========== ")
	/*
		git checkout master
		git pull
		git tag release-20161205 打昨天的tag
		git push origin --tag 推送昨天的tag
		git checkout -b release-20161206 打今天的新分支
		git push origin -u release-20161206 推送今天新分支
	*/
	now := time.Now()
	nowTime := now.Format("20060102")
	today := now.Weekday().String()
	yesterday := now.AddDate(0, 0, -1).Format("20060102")
	if today == "Monday" {
		yesterday = now.AddDate(0, 0, -3).Format("20060102")
	}

	println("[finish merge release to master]\n")
	println("git checkout master")
	println("git pull")
	println("git tag release-" + yesterday)
	println("git push origin --tag")
	println("git checkout -b release-" + nowTime)
	println("git push origin -u release-" + nowTime)
	println("\n========== finish git release generate" + "========== ")
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
