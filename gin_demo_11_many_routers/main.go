package main

import (
	"fmt"
	"gin_project/gin_demo_11_many_routers/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}