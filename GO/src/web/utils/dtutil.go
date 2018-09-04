// https://github.com/xcltapestry/xclpkg/blob/master/utils/dtutil.go
// 使用示例:https://blog.csdn.net/xcl168/article/details/42710537
package utils

import (
	"fmt"
	"strings"
	"time"
)

func GoStdTime() string {
	return "2006-01-02 15:04:05"
}

func GoStdUnixDate() string {
	return "Mon Jan _2 15:04:05 MST 2006"
}

func GoStdRubyDate() string {
	return "Mon Jan 02 15:04:05 -0700 2006"
}

func GetTmStr(tm time.Time, format string) string {
	patterns := []string{
		"y", "2006",
		"m", "01",
		"d", "02",

		"Y", "2006",
		"M", "01",
		"D", "02",

		"h", "03", //12小时制
		"H", "15", //24小时制

		"i", "04",
		"s", "05",

		"t", "pm",
		"T", "PM",
	}
	return convStr(tm, format, patterns)
}

func GetTmShortStr(tm time.Time, format string) string {
	patterns := []string{
		"y", "06",
		"m", "1",
		"d", "2",

		"Y", "06",
		"M", "1",
		"D", "2",

		"h", "3", //12小时制
		"H", "15", //24小时制

		"i", "4",
		"s", "5",

		"t", "pm",
		"T", "PM",
	}

	return convStr(tm, format, patterns)
}

func convStr(tm time.Time, format string, patterns []string) string {
	replacer := strings.NewReplacer(patterns...)
	strfmt := replacer.Replace(format)
	return tm.Format(strfmt)
}

func GetLocaltimeStr() string {
	now := time.Now().Local()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	zone, _ := now.Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

func GetGmtimeStr() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
}

func GetUnixTimeStr(ut int64, format string) string {
	t := time.Unix(ut, 0)
	return GetTmStr(t, format)
}

func GetUnixTimeShortStr(ut int64, format string) string {
	t := time.Unix(ut, 0)
	return GetTmShortStr(t, format)
}

func Greatest(arr []time.Time) time.Time {
	var temp time.Time
	for _, at := range arr {
		if temp.Before(at) {
			temp = at
		}
	}
	return temp
}

type TimeSlice []time.Time

func (s TimeSlice) Len() int {
	return len(s)
}

func (s TimeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s TimeSlice) Less(i, j int) bool {
	if s[i].IsZero() {
		return false
	}
	if s[j].IsZero() {
		return true
	}
	return s[i].Before(s[j])
}

/*
package main

//author:xcl
//2014-1-14

import (
	"fmt"
	"time"
    "github.com/xclpkg/utils"
    "sort"
)

func main(){

    t := time.Now();
    //alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';
    //select to_date('2014-06-09 18:04:06','yyyy-MM-dd HH24:mi:ss') as dt from dual;
    fmt.Println("\n演示时间 => ",utils.GetTmShortStr(t,"y-m-d H:i:s a"))

    //返回当前是一年中的第几天
    //select to_char(sysdate,'ddd'),sysdate from dual;
    yd := t.YearDay();
    fmt.Println("一年中的第几天: ",yd)

    //一年中的第几周
    year,week := t.ISOWeek()
    fmt.Println("一年中的第几周: ",year," | ",week)

    //当前是周几
    //select to_char(sysdate,'day') from dual;
    //select to_char(sysdate,'day','NLS_DATE_LANGUAGE = American') from dual;
    fmt.Println("当前是周几: ",t.Weekday().String())

    //字符串转成time.Time
    //alter session set nls_date_format='yyyy-mm-dd hh:mi:ss';
    //select to_date('14-06-09 6:04:06','yy-MM-dd hh:mi:ss') as dt from dual;
    tt,er := time.Parse(utils.GoStdTime(),"2014-06-09 16:04:06")
    if(er != nil){
        fmt.Println("字符串转时间: parse error!")
    }else{
        fmt.Println("字符串转时间: ",tt.String())
    }


    fmt.Println("\n演示时间 => ",utils.GetTmStr(t,"y-m-d h:i:s"))

    ta := t.AddDate(1,0,0)
    fmt.Println("增加一年 => ",utils.GetTmStr(ta,"y-m-d"))

    ta = t.AddDate(0,1,0)
    fmt.Println("增加一月 => ",utils.GetTmStr(ta,"y-m-d"))

    //select sysdate,sysdate + interval '1' day from dual;
    ta = t.AddDate(0,0,1) //18
    fmt.Println("增加一天 => ",utils.GetTmStr(ta,"y-m-d"))

    durdm,_ := time.ParseDuration("432h")
    ta = t.Add(durdm)
    fmt.Println("增加18天(18*24=432h) => ",utils.GetTmStr(ta,"y-m-d"))

    //select sysdate,sysdate - interval '7' hour from dual;
    dur,_ := time.ParseDuration("-2h")
    ta = t.Add(dur)
    fmt.Println("减去二小时 => ",utils.GetTmStr(ta,"y-m-d h:i:s"))

    //select sysdate,sysdate - interval '7' MINUTE from dual;
    durmi,_ := time.ParseDuration("-7m")
    ta = t.Add(durmi)
    fmt.Println("减去7分钟 => ",utils.GetTmStr(ta,"y-m-d h:i:s"))

    //select sysdate,sysdate - interval '10' second from dual;
    durs,_ := time.ParseDuration("-10s")
    ta = t.Add(durs)
    fmt.Println("减去10秒 => ",utils.GetTmStr(ta,"y-m-d h:i:s"))

    ttr,er := time.Parse(utils.GoStdTime(),"2014-06-09 16:58:06")
    if(er != nil){
        fmt.Println("字符串转时间: 转换失败!")
    }else{
        fmt.Println("字符串转时间: ",ttr.String())
    }

    //alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';
    //select trunc(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;
    // SQL => 2014-06-09 16:58:00
    // Truncate =>  2014-06-09 16:50:00
    durtr,_ := time.ParseDuration("10m")
    ta = ttr.Truncate(durtr)
    fmt.Println("Truncate => ",utils.GetTmStr(ta,"y-m-d H:i:s"))

    //select round(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;
    // SQL => 2014-06-09 16:58:00
    // Round =>  2014-06-09 17:00:00
    ta = ttr.Round(durtr)
    fmt.Println("Round => ",utils.GetTmStr(ta,"y-m-d H:i:s"))

    //日期比较
    tar1,_ := time.Parse(utils.GoStdTime(),"2014-06-09 19:38:36")
    tar2,_ := time.Parse(utils.GoStdTime(),"2015-01-14 17:08:26")
    if tar1.After(tar2) {
        fmt.Println("tar1 > tar2")
    }else if tar1.Before(tar2) {
        fmt.Println("tar1 < tar2")
    }else{
        fmt.Println("tar1 = tar2")
    }
    tar3,_ := time.Parse(utils.GoStdTime(),"2000-07-19 15:58:16")

    //日期列表中最晚日期
    // select greatest('2014-06-09','2015-01-14','2000-07-19') from dual;
    var arr utils.TimeSlice
    arr = []time.Time{tar1,tar2,tar3}
    temp := utils.Greatest(arr)
    fmt.Println("日期列表中最晚日期 => ",utils.GetTmStr(temp,"y-m-d"))

    //日期数组从早至晚排序
    fmt.Println("\n日期数组从早至晚排序")
    sort.Sort(arr)
    for _,at := range arr {
         fmt.Println("Sort => ",utils.GetTmStr(at,"y-m-d H:i:s"))
    }

}

*/
