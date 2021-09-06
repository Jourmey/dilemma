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
		FindOne(id int64) (*TaskInfo, error)
		Update(data TaskInfo) error
		Delete(id int64) error
	}

	defaultTaskInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TaskInfo struct {
		Id         int64     `db:"id"`          // id
		TaskId     int64     `db:"task_id"`     // 关联任务
		Format     string    `db:"format"`      // 链接:dash-flv360
		Container  string    `db:"container"`   // 类型:mp4
		Quality    string    `db:"quality"`     // 质量:流畅 360P
		Size       int64     `db:"size"`        // 任务大小
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

func (m *defaultTaskInfoModel) FindOne(id int64) (*TaskInfo, error) {
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

func (m *defaultTaskInfoModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
