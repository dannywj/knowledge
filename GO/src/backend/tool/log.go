package tool

import (
	"fmt"
	"os"
	"time"
)

// 写日志
func WriteLog(logName string, str string, showTime bool) {
	// 获取当前时间
	now := time.Now()
	userFile := logName + now.Format("_20060102") + ".log"
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
