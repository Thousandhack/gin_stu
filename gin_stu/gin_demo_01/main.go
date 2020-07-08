package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	// 1.创建路由
	// 默认使用两个中间件   Logger(), Recovery()
	r := gin.Default()
	// 2.绑定路由规则，指定的函数
	// *gin.Context   封装了request和response
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"hello, world!")
	})

	r.POST("/userPost",userPost)
	// 3.监听端口
	// // Run("里面不指定端口号默认为8080")
	r.Run(":8000") // 自定义
}


func userPost(c *gin.Context){

	c.String(http.StatusOK,"hello, world!")
}