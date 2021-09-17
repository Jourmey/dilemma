package model

import (
	"database/sql/driver"
	"encoding/json"
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
		Count() (int, error)
	}

	defaultTaskModel struct {
		db *gorm.DB
	}

	Task struct {
		Id         int       `gorm:"id" json:"id"`                   // id
		Url        string    `gorm:"url" json:"url"`                 // 链接
		Signatures string    `gorm:"signatures" json:"signatures"`   // 特征码
		Tag        *TagList  `gorm:"tag;type:json" json:"tag"`       // 标签
		Status     int       `gorm:"status" json:"status"`           // 任务状态 0未处理 1处理中 2获取信息 3获取失败
		Title      string    `gorm:"title" json:"title"`             // 标题
		Site       string    `gorm:"site" json:"site"`               // 平台
		CreateTime time.Time `gorm:"create_time" json:"create_time"` // 创建时间
		UpdateTime time.Time `gorm:"update_time" json:"update_time"` // 修改时间
	}

	TagList []int
)

// Value 实现方法
func (p *TagList) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

// Scan 实现方法
func (p *TagList) Scan(input interface{}) error {
	i := input.([]byte)
	if len(i) == 0 {
		return nil
	}
	return json.Unmarshal(i, &p)
}

func NewTaskModel(db *gorm.DB) TaskModel {
	return &defaultTaskModel{
		db: db,
	}
}

func (*Task) TableName() string {
	return "task"
}

func (m *defaultTaskModel) findOne(where GormFunc) (*Task, error) {
	var v Task
	err := where(m.db).First(&v).Error
	return &v, err
}

func (m *defaultTaskModel) finds(where GormFunc) ([]*Task, error) {
	var v []*Task
	err := where(m.db).Table("task").Find(&v).Error
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
	return m.db.Where("id = ?", id).Delete(&Tag{}).Error
}

func (m *defaultTaskModel) Finds(limit, offset int) ([]*Task, error) {
	if limit <= 0 {
		limit = 20
	}
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

func (m *defaultTaskModel) Count() (int, error) {
	var count int64
	err := m.db.Model(Task{}).Count(&count).Error
	return int(count), err
}
