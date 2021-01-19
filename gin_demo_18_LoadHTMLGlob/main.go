package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("gin_project/gin_demo_18_LoadHTMLGlob/tem/**/*")
	// 如果你需要引入静态文件需要定义一个静态文件目录
	//r.Static("/assets", "./assets")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "The_test", "address": "www.baidu.com"})
	})
	r.Run(":8000")
}
