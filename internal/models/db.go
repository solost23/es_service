package models

import (
	"es_service/configs"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMysql(config *configs.ServerConfig) (db *gorm.DB, err error) {
	mySQLConf := config.MySQLConfig

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               mySQLConf.DataSourceName,
		DefaultStringSize: 100,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Panic(err)
	}
	sqlDB.SetMaxOpenConns(mySQLConf.MaxOpenConn)
	sqlDB.SetMaxIdleConns(mySQLConf.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(mySQLConf.MaxConnLifeTime) * time.Second)

	if err != nil {
		return nil, err
	}
	return db, nil
}
