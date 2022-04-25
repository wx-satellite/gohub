package config

import (
	"github.com/spf13/cast"
	viperLib "github.com/spf13/viper"
	"os"
)

var viper *viperLib.Viper

type Func func() map[string]interface{}

var C map[string]Func

func init() {

	// 初始化 viper 库
	viper = viperLib.New()

	// 设置配置文件类型
	// 支持 "json", "toml", "yaml", "yml", "properties","props", "prop", "env", "dotenv"
	viper.SetConfigType("env")

	// 设置查找配置文件的路径，相对于 main.go
	viper.AddConfigPath(".")

	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")

	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	C = make(map[string]Func)
}

func InitConfig(env string) {

	// 加在环境配置
	loadEnv(env)

	// 注册配置信息
	loadConfig()
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if envSuffix != "" {
		filePath := envPath + "." + envSuffix
		if _, err := os.Stat(filePath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filePath
		}
	}

	viper.SetConfigName(envPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控环境配置文件
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range C {
		viper.Set(name, fn())
	}
}

// Env 获取环境变量的值，name 的格式：database.port 或者 port
func Env(name string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(name, defaultValue[0])
	}
	return internalGet(name)
}

// Add 新增配置项
func Add(name string, f Func) {
	C[name] = f
}

// internalGet 获取指定key的配置信息，底层方法
func internalGet(name string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(name) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(name)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
