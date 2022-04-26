package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	"gohub/config"
)

func main() {

	// 初始化配置文件
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.Initialize(env)

	// 初始化 gin
	router := gin.Default()

	// 初始化DB
	bootstrap.SetupDB()
	// 初始化路由
	bootstrap.SetupRoute(router)

	// 启动服务
	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		fmt.Printf("server run fail：%s\n", err.Error())
		return
	}

}
