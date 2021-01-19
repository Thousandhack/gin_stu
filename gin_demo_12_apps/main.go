package main

import (
	"fmt"
	"gin_project/gin_demo_12_apps/app/blog"
	"gin_project/gin_demo_12_apps/app/shop"
	"gin_project/gin_demo_12_apps/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
