package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 成功返回
func Success(c *gin.Context, data map[string]interface{}) {
	result := make(map[string]interface{})
	result["msg"] = "success"
	result["code"] = 0
	result["data"] = data
	c.JSON(http.StatusOK, result)
}

// Error 失败返回
func Error(c *gin.Context, code int, msg string) {
	result := make(map[string]interface{})
	result["msg"] = msg
	result["code"] = code
	c.JSON(http.StatusOK, result)
}
