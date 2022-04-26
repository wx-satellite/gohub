package bootstrap

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/config"
	"gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func SetupDB() {
	var dbConfig gorm.Dialector

	switch config.GetString("database.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		db := config.GetString("database.sqlite.database")
		dbConfig = sqlite.Open(db)
	default:
		panic("database connection not supported")
	}

	// 连接数据库，并设置日志格式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 空闲时间
	database.SQLDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	// 空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	// 自动迁移
	autoMigrate()
}

// gorm 自带自动迁移，文档地址：https://gorm.io/zh_CN/docs/migration.html
func autoMigrate() {
	_ = database.DB.AutoMigrate(new(user.User))
}
