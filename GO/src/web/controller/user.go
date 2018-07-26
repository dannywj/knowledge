package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
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

	if endDate-beginDate > 10 {
		Error(c, 101, "只能查询10天以内的数据")
		return
	}
	result := make(map[string]interface{})

	for day := beginDate; day <= endDate; day++ {
		total, uniqueCount := model.GetActionCountByDate(day)
		result[strconv.Itoa(day)] = gin.H{
			"total":       total,
			"uniqueCount": uniqueCount,
		}
	}
	Success(c, result)
}

func GetActionStatisticsPage(c *gin.Context) {
	// c.HTML(http.StatusOK, "action_count.html", gin.H{
	// 	"title": "welcome action!",
	// })

	//模板文件的拼接
	_, err := template.ParseFiles("layout.html", "head.tpl",
		"content.html", "common.js", "scripts.js")
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title": "welcome action!",
	})

}
