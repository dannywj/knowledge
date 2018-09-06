package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	//"strconv"
	"web/model"
)

// 能量球
func GetBallInfoPage(c *gin.Context) {
	//模板文件的拼接
	_, err := template.ParseFiles("layout_ball.html", "head.tpl",
		"content_ball.html", "common.js", "scripts_ball.js")
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "layout_ball.html", gin.H{
		"title": "元宝树后台管理",
	})
}

func ResetBallAction(c *gin.Context) {
	guid := c.Query("guid")

	result := model.ResetUserBall(guid)

	Success(c, gin.H{
		"re": result,
	})
}
