package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("D:/goCode/goproject/src/gin_stu/gin_demo_17_LoadHTMLGlob/tem/*")
	// 上面的pattern需要加到项目的根路径才算完整路径
	// 如果相对路径有问题，就使用绝对路径
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})
	r.Run(":8000")
}