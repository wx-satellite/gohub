// bootstrap 处理程序的初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册 api 路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(ctx *gin.Context) {
		// 获取请求头的Accept
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 json
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
