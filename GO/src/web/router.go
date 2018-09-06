package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	. "web/controller"
)

// 初始路由
func initRouter() *gin.Engine {
	router := gin.Default()
	// 设置中间件过滤器
	router.Use(Logger())
	// 设置页面模板路径
	router.LoadHTMLGlob("view/*")
	// 绑定接口
	router.GET("/", IndexAction)
	router.GET("/user/:guid", GetUserAction)
	router.GET("/action/count/", GetActionCountAction)
	router.GET("/action/statistics/", GetActionStatisticsPage)
	router.GET("/redbag/statistics/", GetRedbagStatisticsPage)
	router.GET("/redbag/info/", GetRedbagInfoAction)
	router.GET("/base/info/", GetBaseInfoAction)
	router.GET("/withdraw/card/", GetWithdrawCardPage)
	router.GET("/withdraw/cardAdd/", GetWithdrawCardAction)
	return router
}

// 路由中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set example variable
		//c.Set("example", "12345")

		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(fmt.Sprintf("api call time:%v", latency))
		// access the status we are sending
		status := c.Writer.Status()
		log.Println(fmt.Sprintf("http status:%v", status))
	}
}
