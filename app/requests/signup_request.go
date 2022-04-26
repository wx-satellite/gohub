package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type ValidateFunc func(ctx *gin.Context, data interface{}) map[string][]string

func Validate(ctx *gin.Context, data interface{}, valid ValidateFunc) bool {
	if err := ctx.ShouldBindJSON(data); err != nil {
		// 422 用来表示校验错误
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	errs := valid(ctx, data)
	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}
	return true
}

type SignupPhoneExistRequest struct {
	Phone string `json:"phone" valid:"phone"`
}

func ValidateSignupPhoneExist(ctx *gin.Context, data interface{}) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义错误信息
	messages := govalidator.MapData{
		"phone": []string{
			"required:Phone 为必填项",
			"digits:手机号长度必须是 11 位",
		},
	}

	return validate(data, rules, messages)
}

type SignupEmailExistRequest struct {
	Email string `json:"email" valid:"email"`
}

func ValidateSignupEmailExist(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		// min 和 max 是闭区间
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需要大于4",
			"max:Email 长度需要小于30",
			"email:Email 格式不正确",
		},
	}
	return validate(data, rules, messages)
}

// validate 验证底层逻辑
func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opt := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
		// 默认情况下除非使用了required，否则只有这个key有值了才会被校验
		// RequiredDefault: true, //  所有的验证规则都会校验，不管前端是否传递了key
	}
	return govalidator.New(opt).ValidateStruct()
}
