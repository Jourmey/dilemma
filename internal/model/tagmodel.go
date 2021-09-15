package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	TagModel interface {
		Insert(data *Tag) (int, error)
		FindOne(id int) (*Tag, error)
		Finds(limit, offset int) ([]*Tag, error)
	}

	defaultTag struct {
		db *gorm.DB
	}

	Tag struct {
		Id         int       `gorm:"id" json:"id"`                   // id
		Tag        string    `gorm:"tag" json:"tag"`                 // tag
		CreateTime time.Time `gorm:"create_time" json:"create_time"` // 创建时间
		UpdateTime time.Time `gorm:"update_time" json:"update_time"` // 修改时间
	}
)

func NewTagModel(db *gorm.DB) TagModel {
	return &defaultTag{
		db: db,
	}
}

func (Tag) TableName() string {
	return "tag"
}

func (d *defaultTag) findOne(where GormFunc) (*Tag, error) {
	var v Tag
	err := where(d.db).First(&v).Error
	return &v, err
}

func (d *defaultTag) finds(where GormFunc) ([]*Tag, error) {
	var v []*Tag
	err := where(d.db).Find(&v).Error
	return v, err
}

func (d *defaultTag) Insert(data *Tag) (int, error) {
	err := d.db.Omit("create_time", "update_time").Create(data).Error
	return data.Id, err
}

func (d *defaultTag) FindOne(id int) (*Tag, error) {
	return d.findOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
}

func (d *defaultTag) Finds(limit, offset int) ([]*Tag, error) {
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	})
}
