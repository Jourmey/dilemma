package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type (
	Config struct {
		rest.RestConf
		Mysql Mysql
	}
	Mysql struct {
		Host     string
		Port     int
		DbName   string
		User     string
		Password string
	}
)
