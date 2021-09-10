package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type (
	Config struct {
		rest.RestConf
		Mysql      Mysql
		Staticfile Staticfile
	}
	Mysql struct {
		Host     string
		Port     int
		DbName   string
		User     string
		Password string
	}
	Staticfile struct {
		Host string
		Port int
		Root string
	}
)
