package user

import "gohub/pkg/database"

// IsEmailExist 判断邮箱是否已经被注册
// 因为包名是 user，所以调用的形式是 user.IsEmailExist，因此不需要在函数名中加入 user
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号是否已经被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
