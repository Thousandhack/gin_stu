package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8000/user 才会打印出来默认的值 为小白
		name := c.DefaultQuery("name", "小白")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
		// http://localhost:8000/user?name=hsz
		// 返回结果为：
		// hello hsz
	})
	r.Run(":8000")
}
