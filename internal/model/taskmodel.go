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
	taskFieldNames          = builderx.RawFieldNames(&Task{})
	taskRows                = strings.Join(taskFieldNames, ",")
	taskRowsExpectAutoSet   = strings.Join(stringx.Remove(taskFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	taskRowsWithPlaceHolder = strings.Join(stringx.Remove(taskFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	TaskModel interface {
		Insert(data Task) (sql.Result, error)
		FindOne(id int) (*Task, error)
		Update(data Task) error
		Delete(id int) error
		Finds(limit, offset int) ([]*Task, error)
	}

	defaultTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Task struct {
		Id         int       `db:"id"`          // id
		Url        string    `db:"url"`         // 链接
		Signatures string    `db:"signatures"`  // 特征码
		Tag        int       `db:"tag"`         // 标签
		Status     int       `db:"status"`      // 任务状态 0未处理 1获取信息
		Title      string    `db:"title"`       // 标题
		Site       string    `db:"site"`        // 平台
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
	}
)

func NewTaskModel(conn sqlx.SqlConn) TaskModel {
	return &defaultTaskModel{
		conn:  conn,
		table: "`task`",
	}
}

func (m *defaultTaskModel) Insert(data Task) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, taskRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Url, data.Signatures, data.Tag, data.Status, data.Title, data.Site)
	return ret, err
}

func (m *defaultTaskModel) FindOne(id int) (*Task, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskRows, m.table)
	var resp Task
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

func (m *defaultTaskModel) Update(data Task) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Url, data.Signatures, data.Tag, data.Status, data.Title, data.Site, data.Id)
	return err
}

func (m *defaultTaskModel) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultTaskModel) finds(other string) ([]*Task, error) {
	query := fmt.Sprintf("select %s from %s %s", taskRows, m.table, other)
	var resp []*Task
	err := m.conn.QueryRows(&resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultTaskModel) Finds(limit, offset int) ([]*Task, error) {
	o := fmt.Sprintf("limit %d offset %d", limit, offset)
	return m.finds(o)
}
