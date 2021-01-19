package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" like "+action)
	})
	//默认为监听8080端口
	r.Run(":8000")
}


/*
	使用：http://127.0.0.1:8000/user/hsz/zero   GET方法进行访问
	返回：
		hsz like zero

*/