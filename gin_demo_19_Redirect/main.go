package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 访问 http://127.0.0.1:8000/index 后直接重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.cnblogs.com/hszstudypy/articles/12153575.html")
	})
	r.Run(":8000")
}