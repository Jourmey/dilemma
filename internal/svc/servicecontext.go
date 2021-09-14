package svc

import (
	"dilemma/internal/config"
	"dilemma/internal/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     model.MustGetMysqlDB(c.Mysql),
		//DB:     model.MustGetMysqlDB(c.Mysql).Debug(),
	}
}
