package main

import (
	"fmt"
	"gin_stu/gin_demo_10_router/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}