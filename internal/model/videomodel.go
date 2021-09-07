package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	videoFieldNames          = builderx.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	VideoModel interface {
		Insert(data Video) (sql.Result, error)
		FindOne(id int64) (*Video, error)
		Update(data Video) error
		Delete(id int64) error
	}

	defaultVideoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Video struct {
		Id         int64     `db:"id"`           // id
		TaskInfoId int64     `db:"task_info_id"` // 关联任务信息
		Path       string    `db:"path"`         // 路径
		CreateTime time.Time `db:"create_time"`  // 创建时间
		UpdateTime time.Time `db:"update_time"`  // 修改时间
	}
)

func NewVideoModel(conn sqlx.SqlConn) VideoModel {
	return &defaultVideoModel{
		conn:  conn,
		table: "`video`",
	}
}

func (m *defaultVideoModel) Insert(data Video) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, videoRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.TaskInfoId, data.Path)
	return ret, err
}

func (m *defaultVideoModel) FindOne(id int64) (*Video, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	var resp Video
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Update(data Video) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.TaskInfoId, data.Path, data.Id)
	return err
}

func (m *defaultVideoModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
