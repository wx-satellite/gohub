package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	router := gin.Default()

	// 初始化路由
	bootstrap.SetupRoute(router)

	// 启动服务
	err := router.Run(":3000")
	if err != nil {
		fmt.Printf("server run fail：%s\n", err.Error())
		return
	}

}
