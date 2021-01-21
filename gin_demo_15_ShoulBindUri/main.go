package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello https://golang.org !",
	})
}

// 定义接收数据的结构体
type Login struct {
	// binding:"required" 表示这个参数是必传的
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	// 声明接收的变量
	var login Login
	if err := c.ShouldBindUri(&login); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 判断用户名密码是否正常
	if login.User != "root" || login.Password != "123456" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "304",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "200",
		"uri":   "ok",
	})
}

func main() {
	r := gin.Default()
	r.GET("/login/:user/:password", LoginHandler)
	// 如果run启动失败打印下面的语句
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	/*
	   GET 方法
	   http://127.0.0.1:8080/login/root/123456
	   {
	       "status": "200",
	       "uri": "ok"
	       }

	*/
}
