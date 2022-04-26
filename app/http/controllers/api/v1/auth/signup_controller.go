package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseController
}

// IsPhoneExist 检测手机号是否被注册
func (c *SignupController) IsPhoneExist(ctx *gin.Context) {
	var in struct {
		Phone string `json:"phone"`
	}

	// 解析 json 请求
	if err := ctx.ShouldBindJSON(&in); err != nil {
		// 422 用来表示校验错误
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(in.Phone)})
}
