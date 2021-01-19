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

func main() {
	r := gin.Default()
	r.GET("/togopher", helloHandler)
	// 如果run启动失败打印下面的语句
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}