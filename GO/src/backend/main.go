package main

import (
	"backend/business"
)

/*
任务列表:
	get_device_from_db-种树用户设备id
	get_device_from_file-文件用户设备id
	check_user_energy-检查用户能量等级
	gen_user_list_file-生成种树用户id文件
*/
func main() {
	business.PrintLog("========begin task=========")
	business.Run("check_user_energy")
	business.PrintLog("========end task=========")
}
