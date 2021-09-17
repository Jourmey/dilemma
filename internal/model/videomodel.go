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
		Count() (int, error)
	}

	defaultVideoModel struct {
		db *gorm.DB
	}

	Video struct {
		Id         int       `gorm:"id" json:"id"`                     // id
		TaskInfoId int       `gorm:"task_info_id" json:"task_info_id"` // 关联任务信息
		Path       string    `gorm:"path" json:"path"`                 // 路径
		Title      string    `gorm:"title" json:"title"`               // 标题
		CreateTime time.Time `gorm:"create_time" json:"create_time"`   // 创建时间
		UpdateTime time.Time `gorm:"update_time" json:"update_time"`   // 修改时间
	}
)

func NewVideoModel(db *gorm.DB) VideoModel {
	return &defaultVideoModel{
		db: db,
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
	return m.db.Where("id = ?", id).Delete(&Tag{}).Error
}

func (m *defaultVideoModel) Finds(limit, offset int) ([]*Video, error) {
	if limit <= 0 {
		limit = 20
	}
	return m.finds(func(db *gorm.DB) *gorm.DB {
		return m.db.Limit(limit).Offset(offset)
	})
}

func (m *defaultVideoModel) Count() (int, error) {
	var count int64
	err := m.db.Model(Video{}).Count(&count).Error
	return int(count), err
}
