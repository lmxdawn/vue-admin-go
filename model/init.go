package model

import (
	"fmt"
	"github.com/lmxdawn/vue-admin-go/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DB ...
var DB *gorm.DB

// Database ...
func Database(mysqlConfig config.MySQLConfig) error {
	log.Info().Msgf("mysql %v", mysqlConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Db)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		log.Error().Msgf("conect db err:%v", err)
		return err
	}
	//设置连接池
	//空闲
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Msgf("get sqlDB err:%v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	//打开
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConfig.ConnMaxLifeTime))

	DB = db

	return nil
}
