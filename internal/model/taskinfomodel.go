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
	taskInfoFieldNames          = builderx.RawFieldNames(&TaskInfo{})
	taskInfoRows                = strings.Join(taskInfoFieldNames, ",")
	taskInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(taskInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	taskInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(taskInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	TaskInfoModel interface {
		Insert(data TaskInfo) (sql.Result, error)
		FindOne(id int) (*TaskInfo, error)
		Update(data TaskInfo) error
		Delete(id int) error
		Finds(limit, offset int) ([]*TaskInfo, error)
	}

	defaultTaskInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TaskInfo struct {
		Id         int       `db:"id"`          // id
		TaskId     int       `db:"task_id"`     // 关联任务
		Format     string    `db:"format"`      // 链接:dash-flv360
		Container  string    `db:"container"`   // 类型:mp4
		Quality    string    `db:"quality"`     // 质量:流畅 360P
		Size       int       `db:"size"`        // 任务大小
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
	}
)

func NewTaskInfoModel(conn sqlx.SqlConn) TaskInfoModel {
	return &defaultTaskInfoModel{
		conn:  conn,
		table: "`task_info`",
	}
}

func (m *defaultTaskInfoModel) Insert(data TaskInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, taskInfoRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.TaskId, data.Format, data.Container, data.Quality, data.Size)
	return ret, err
}

func (m *defaultTaskInfoModel) FindOne(id int) (*TaskInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskInfoRows, m.table)
	var resp TaskInfo
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

func (m *defaultTaskInfoModel) Update(data TaskInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskInfoRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.TaskId, data.Format, data.Container, data.Quality, data.Size, data.Id)
	return err
}

func (m *defaultTaskInfoModel) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultTaskInfoModel) finds(other string) ([]*TaskInfo, error) {
	query := fmt.Sprintf("select %s from %s %s", taskRows, m.table, other)
	var resp []*TaskInfo
	err := m.conn.QueryRows(&resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultTaskInfoModel) Finds(limit, offset int) ([]*TaskInfo, error) {
	o := fmt.Sprintf("limit %d offset %d", limit, offset)
	return m.finds(o)
}
