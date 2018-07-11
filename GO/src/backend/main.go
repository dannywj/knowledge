package main

import (
	"backend/business"
)

/*
任务列表:
	get_device_from_db-种树用户设备id
	get_device_from_file-文件用户设备id
*/
func main() {
	business.PrintLog("========begin task=========")
	business.Run("get_device_from_db")
	business.PrintLog("========end task=========")
}
