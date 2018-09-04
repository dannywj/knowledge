package utils

import (
	"fmt"
	"time"
)

//StrToTime 将str转换为时间格式("2006-01-02 15:04:05")
func StrToTime(format string, st string) time.Time {
	t, _ := time.ParseInLocation(format, st, time.Local)
	return t
}

// DayAggrToTime 将数字时间转成带分隔符的时间字符串  20180521  -> 2018-05-21
func DayAggrToTime(str string) string {
	return fmt.Sprintf("%v-%v-%v", str[0:4], str[4:6], str[6:8])
}
