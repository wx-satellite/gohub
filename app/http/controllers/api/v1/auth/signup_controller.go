package auth

import (
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

	if ok := requests.Validate(ctx, &in, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(in.Phone)})
}

// IsEmailExist 检测手机号是否被注册
func (c *SignupController) IsEmailExist(ctx *gin.Context) {
	var in requests.SignupEmailExistRequest

	if ok := requests.Validate(ctx, &in, requests.ValidateSignupEmailExist); !ok {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exist": user.IsEmailExist(in.Email)})
}
