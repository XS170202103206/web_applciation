package main

import (
	"gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化路由路径
	router.InitRouter(r)

	//r.MaxMultipartMemory = 8 << 20 //8MiB

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
