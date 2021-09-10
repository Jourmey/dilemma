package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	TaskModel interface {
		Insert(data *Task) (int, error)
		FindOne(id int) (*Task, error)
		Delete(id int) error
		Finds(limit, offset int) ([]*Task, error)
		UpdateStatus(id int, Status int, Title, Site string) error
		FindByIds(ids []int) ([]*Task, error)
	}

	defaultTaskModel struct {
		db    *gorm.DB
		table string
	}

	Task struct {
		Id         int       `db:"id"`          // id
		Url        string    `db:"url"`         // 链接
		Signatures string    `db:"signatures"`  // 特征码
		Tag        int       `db:"tag"`         // 标签
		Status     int       `db:"status"`      // 任务状态 0未处理 1处理中 2获取信息 3获取失败
		Title      string    `db:"title"`       // 标题
		Site       string    `db:"site"`        // 平台
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
	}
)

func NewTaskModel(db *gorm.DB) TaskModel {
	return &defaultTaskModel{
		db:    db,
		table: "`task`",
	}
}

func (Task) TableName() string {
	return "task"
}

func (m *defaultTaskModel) findOne(where GormFunc) (*Task, error) {
	var v Task
	err := where(m.db).First(&v).Error
	return &v, err
}

func (m *defaultTaskModel) finds(where GormFunc) ([]*Task, error) {
	var v []*Task
	err := where(m.db).Find(&v).Error
	return v, err
}

func (m *defaultTaskModel) Insert(data *Task) (int, error) {
	err := m.db.Omit("create_time", "update_time").Create(data).Error
	return data.Id, err
}

func (m *defaultTaskModel) FindOne(id int) (*Task, error) {
	return m.findOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
}

func (m *defaultTaskModel) Delete(id int) error {
	return m.db.Where("id = ?", id).Delete(&TaskInfo{}).Error
}

func (m *defaultTaskModel) Finds(limit, offset int) ([]*Task, error) {
	return m.finds(func(db *gorm.DB) *gorm.DB {
		return m.db.Limit(limit).Offset(offset)
	})
}

// 回写状态
func (m *defaultTaskModel) UpdateStatus(id int, Status int, Title, Site string) error {
	return m.db.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status": Status,
		"title":  Title,
		"site":   Site,
	}).Error
}

func (m *defaultTaskModel) FindByIds(ids []int) ([]*Task, error) {
	return m.finds(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})
}
