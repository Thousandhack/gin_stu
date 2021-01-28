package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		/*
			通过：http://localhost:8000/cookie  访问两次的打印如下：
			cookie的值是： NotSet
			[GIN] 2021/01/20 - 15:38:02 |?[97;42m 200 ?[0m|       997.1µs |             ::1 |?[97;44m GET     ?[0m "/cookie"
			[GIN] 2021/01/20 - 15:38:02 |?[90;43m 404 ?[0m|            0s |             ::1 |?[97;44m GET     ?[0m "/favicon.ico"
			cookie的值是： value_cookie

		*/
		fmt.Printf("cookie的值是： %s\n", cookie)
	})
	r.Run(":8000")
}