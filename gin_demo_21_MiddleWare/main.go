package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间
// 一定返回 gin.HandlerFunc
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		time.Sleep(time.Millisecond*500)
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {}为了代码规范
	{
		r.GET("/midTest", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run(":8000")
	/*
	访问URL:http://127.0.0.1:8000/midTest
	返回打印：
		中间件开始执行了
		中间件执行完毕 200
		time: 0s
		request: 中间件

	返回如下的数据：
	{"request":"中间件"}
	*/
}