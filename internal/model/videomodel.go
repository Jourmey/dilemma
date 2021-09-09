package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	VideoModel interface {
		Insert(data *Video) (int, error)
		FindOne(id int) (*Video, error)
		Delete(id int) error
		Finds(limit, offset int) ([]*Video, error)
	}

	defaultVideoModel struct {
		db    *gorm.DB
		table string
	}

	Video struct {
		Id         int       `db:"id"`           // id
		TaskInfoId int       `db:"task_info_id"` // 关联任务信息
		Path       string    `db:"path"`         // 路径
		CreateTime time.Time `db:"create_time"`  // 创建时间
		UpdateTime time.Time `db:"update_time"`  // 修改时间
	}
)

func NewVideoModel(db *gorm.DB) VideoModel {
	return &defaultVideoModel{
		db:    db,
		table: "`video`",
	}
}
func (*Video) TableName() string {
	return "video"
}

func (m *defaultVideoModel) findOne(where GormFunc) (*Video, error) {
	var v Video
	err := where(m.db).First(&v).Error
	return &v, err
}

func (m *defaultVideoModel) finds(where GormFunc) ([]*Video, error) {
	var v []*Video
	err := where(m.db).Find(&v).Error
	return v, err
}

func (m *defaultVideoModel) Insert(data *Video) (int, error) {
	err := m.db.Omit("create_time", "update_time").Create(data).Error
	return data.Id, err
}

func (m *defaultVideoModel) FindOne(id int) (*Video, error) {
	return m.findOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
}

func (m *defaultVideoModel) Delete(id int) error {
	return m.db.Where("id = ?", id).Delete(&TaskInfo{}).Error
}

func (m *defaultVideoModel) Finds(limit, offset int) ([]*Video, error) {
	return m.finds(func(db *gorm.DB) *gorm.DB {
		return m.db.Limit(limit).Offset(offset)
	})
}
