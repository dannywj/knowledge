package utils

import (
	"time"
)

//StrToTime 将str转换为时间格式("2006-01-02 15:04:05")
func StrToTime(format string, st string) time.Time {
	t, _ := time.ParseInLocation(format, st, time.Local)
	return t
}
