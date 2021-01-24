package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件 传进去参数是另外一种形式
	r.Use(myTime)
	// {}为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(200, gin.H{"time": 5})
}

func shopHomeHandler(c *gin.Context) {
	start := time.Now()
	time.Sleep(3 * time.Second)
	res  := time.Since(start)
	c.JSON(200, gin.H{"time": res,"place":"home"})
}

/*
	http://127.0.0.1:8000/shopping/index
	命令行打印：程序用时： 5.0007853s

	返回：
		{
    "time": 5
		}

	http://127.0.0.1:8000/shopping/home
	命令行打印： 程序用时： 3.0009863s
	返回：
		{
    "place": "home",
    "time": 3000978700
		}

*/