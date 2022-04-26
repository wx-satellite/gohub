package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

func RegisterAPIRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		// 登陆注册
		authGroup := v1.Group("/auth")
		{
			sc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", sc.IsPhoneExist)
		}
	}
}
