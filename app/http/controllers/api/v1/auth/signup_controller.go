package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseController
}

// IsPhoneExist 检测手机号是否被注册
func (c *SignupController) IsPhoneExist(ctx *gin.Context) {
	var in requests.SignupPhoneExistRequest

	// 解析 json 请求
	if err := ctx.ShouldBindJSON(&in); err != nil {
		// 422 用来表示校验错误
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&in, ctx)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(in.Phone)})
}
