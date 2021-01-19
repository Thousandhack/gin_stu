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
        //http://localhost:8000/user 才会打印出来默认的值 为:  hello 小白 age is 18
        // http://localhost:8000/user?name=hsz 打印 hsz
        // http://localhost:8000/user?name=hsz&age=22 打印： hello hsz age is 22
        name := c.DefaultQuery("name", "小白")
        age := c.DefaultQuery("age", "18")
        addr := c.Query("addr")  // 这个如果addr没有相应的值，那么就是没有默认返回值
        c.String(http.StatusOK, fmt.Sprintf("hello %s age is %s addr %s", name, age, addr))
    })
    r.Run(":8000")
}
