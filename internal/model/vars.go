package model

import (
	"dilemma/internal/config"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

func GetSqlConn(config config.Mysql) sqlx.SqlConn {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DbName)
	return sqlx.NewMysql(dataSource)
}
