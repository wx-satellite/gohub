package user

import "gohub/app/models"

type User struct {
	models.BaseModel
	models.CommonTimestampsField

	Name string `json:"name,omitempty"`

	// json:"-" 指示JSON解析器忽略字段
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
}
