package model

import (
	"dilemma/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustGetMysqlDB(config config.Mysql) *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dataSource,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

type GormFunc func(*gorm.DB) *gorm.DB
