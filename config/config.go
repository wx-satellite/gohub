package config

import (
	"gohub/pkg/config"
)

// Initialize 触发 config 包的所有 init 函数并加载配置文件
func Initialize(env string) {
	config.InitConfig(env)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return config.GetString(path, defaultValue...)
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return config.GetInt(path, defaultValue...)
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return config.GetFloat64(path, defaultValue...)
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return config.GetInt64(path, defaultValue...)
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return config.GetUint(path, defaultValue...)
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return config.GetBool(path, defaultValue...)
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return config.GetStringMapString(path)
}
