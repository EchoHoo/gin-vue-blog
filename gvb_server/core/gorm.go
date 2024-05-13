package core

import (
	"fmt"
	"gvb_server/gvb_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	return MysqlConnect()
}
func MysqlConnect() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置mysql")
		//log.Println("未配置mysql")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info) //打印所有的日志
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //打印错误的日志
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info) //默认

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   mysqlLogger, //log的配置
		DisableForeignKeyConstraintWhenMigrating: true,        //禁用外键约束
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
		//log.Fatalf(fmt.Sprintf("[%s] mysql 连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用事件，不能超过mysql的wait_timeout
	return db
}
