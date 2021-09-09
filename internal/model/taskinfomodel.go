package model

import (
	"gorm.io/gorm"
)

type (
	TaskInfoModel interface {
		Insert(data TaskInfo) (int, error)
		FindOne(id int) (*TaskInfo, error)
		Delete(id int) error
		Finds(limit, offset int) ([]*TaskInfo, error)
	}

	defaultTaskInfoModel struct {
		db    *gorm.DB
		table string
	}

	TaskInfo struct {
		gorm.Model
		TaskId    int    `db:"task_id"`   // 关联任务
		Format    string `db:"format"`    // 链接:dash-flv360
		Container string `db:"container"` // 类型:mp4
		Quality   string `db:"quality"`   // 质量:流畅 360P
		Size      int    `db:"size"`      // 任务大小
	}
)

func NewTaskInfoModel(db *gorm.DB) TaskInfoModel {
	return &defaultTaskInfoModel{
		db:    db,
		table: "`task_info`",
	}
}

func (TaskInfo) TableName() string {
	return "task_info"
}

func (d *defaultTaskInfoModel) findOne(where GormFunc) (*TaskInfo, error) {
	var v TaskInfo
	err := where(d.db).First(&v).Error
	return &v, err
}

func (d *defaultTaskInfoModel) finds(where GormFunc) ([]*TaskInfo, error) {
	var v []*TaskInfo
	err := where(d.db).Find(&v).Error
	return v, err
}

func (d *defaultTaskInfoModel) Insert(data TaskInfo) (int, error) {
	err := d.db.Omit("create_time", "update_time").Create(&data).Error
	return int(data.ID), err
}

func (d *defaultTaskInfoModel) FindOne(id int) (*TaskInfo, error) {
	return d.findOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})
}

func (d *defaultTaskInfoModel) Delete(id int) error {
	return d.db.Where("id = ?", id).Delete(&TaskInfo{}).Error
}

func (d *defaultTaskInfoModel) Finds(limit, offset int) ([]*TaskInfo, error) {
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return d.db.Limit(limit).Offset(offset)
	})
}
