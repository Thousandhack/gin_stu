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
	var form Login
	// Bind() 默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

    // 判断用户名密码是否正常
    if form.User != "root" || form.Password != "123456" {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "304",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status": "200",
        "form":"ok",
    })

    //// 声明接收的变量
    //var data Login
    //// 将request的body中的数据自动按照json格式解析到结构体
    //if err := c.ShouldBindJSON(&data); err != nil {
    //    // 返回错误信息
    //    // gin.H 封装了生成json数据的工具
    //    c.JSON(http.StatusBadRequest, gin.H{
    //        "err": err.Error(),
    //    })
    //    return
    //}
    //// 判断用户名密码是否正常
    //if data.User != "root" || data.Password != "123456" {
    //    c.JSON(http.StatusBadRequest, gin.H{
    //        "status": "304",
    //    })
    //    return
    //}
    //c.JSON(http.StatusOK, gin.H{
    //    "status": "200",
    //})

    /*
    	http://127.0.0.1:8080/login

    {
	    "user":"root",
	    "password":"123456"
			}
	返回：
    	{
	    "status": "200"
		}

    */

}

func main() {
    r := gin.Default()
    // 最简单的路由树
    r.GET("/gopher", helloHandler)
    r.GET("/blog", helloHandler)
    r.GET("/website", helloHandler)
    r.POST("/login", LoginHandler)
    // 如果run启动失败打印下面的语句
    if err := r.Run(); err != nil {
        fmt.Println("startup service failed, err:%v\n", err)
    }
    /*
    		http://127.0.0.1:8080/blog
    		http://127.0.0.1:8080/gopher
    		http://127.0.0.1:8080/website

    		返回：
    			{
    	    "message": "Hello https://golang.org !"
    			}


    */
}
