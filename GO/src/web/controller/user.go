package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"web/model"
)

// IndexAction 默认接口
func IndexAction(c *gin.Context) {
	// c.String(http.StatusOK, "welcome planting!")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "welcome planting!",
	})
}

// GetUserAction 获取用户信息 by guid
func GetUserAction(c *gin.Context) {
	guid, _ := strconv.Atoi(c.Param("guid"))
	energy := model.GetUserInfoByGuid(guid)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"guid": energy,
	})
}

// GetActionCountAction 查询指定时间段的操作日志数
func GetActionCountAction(c *gin.Context) {
	beginDate, _ := strconv.Atoi(c.Query("begin_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))

	if endDate-beginDate > 5 {
		Error(c, 101, "只能查询5天以内的数据")
		return
	}
	result := make(map[string]interface{})
	resultList := model.GetActionLogByDate(beginDate, endDate)
	for _, value := range resultList {
		listItemArr := strings.Split(value, "_")
		fmt.Printf("date:%v total:%v unique:%v\n", listItemArr[0], listItemArr[1], listItemArr[2])
		result[listItemArr[0]] = gin.H{
			"total":       listItemArr[1],
			"uniqueCount": listItemArr[2],
		}
	}
	Success(c, result)

	// 非协程循环获取方式
	// result := make(map[string]interface{})
	// for day := beginDate; day <= endDate; day++ {
	// 	total, uniqueCount := model.GetActionCountByDate(day)
	// 	result[strconv.Itoa(day)] = gin.H{
	// 		"total":       total,
	// 		"uniqueCount": uniqueCount,
	// 	}
	// }
	// Success(c, result)
}

func GetActionStatisticsPage(c *gin.Context) {
	//模板文件的拼接
	_, err := template.ParseFiles("layout_action.html", "head.tpl",
		"content.html", "common.js", "scripts_action.js")
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "layout_action.html", gin.H{
		"title": "welcome action!",
	})
}

// 红包页
func GetRedbagStatisticsPage(c *gin.Context) {
	//模板文件的拼接
	_, err := template.ParseFiles("layout_redbag.html", "head.tpl",
		"content_redbag.html", "common.js", "scripts_redbag.js")
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "layout_redbag.html", gin.H{
		"title": "welcome action!",
	})
}

// 红包信息
func GetRedbagInfoAction(c *gin.Context) {
	beginDate, _ := strconv.Atoi(c.Query("begin_date"))
	endDate, _ := strconv.Atoi(c.Query("end_date"))

	if endDate-beginDate > 10 {
		Error(c, 101, "只能查询10天以内的数据")
		return
	}
	result := make(map[string]interface{})

	resultList := model.GetMoneyByDate(beginDate, endDate)

	for _, value := range resultList {
		listItemArr := strings.Split(value, "_")
		fmt.Printf("date:%v fee:%v\n", listItemArr[0], listItemArr[1])
		result[listItemArr[0]] = gin.H{
			"total":       listItemArr[1],
			"uniqueCount": 0,
		}
	}
	Success(c, result)
}
