package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	TaskInfoModel interface {
		Insert(data *TaskInfo) (int, error)
		FindOne(id int) (*TaskInfo, error)
		Delete(id int) error
		FindByIds(id []int) ([]*TaskInfo, error)
		Finds(limit, offset int) ([]*TaskInfo, error)
		FindsByTaskID(taskId int) ([]*TaskInfo, error)
		FindDetailsById(ids []int) (map[int]*TaskAndInfo, error)
		UpdateStatus(id int, Status int) error
		Count() (int, error)
	}

	defaultTaskInfoModel struct {
		db *gorm.DB
	}

	TaskInfo struct {
		Id         int       `gorm:"id" json:"id"`                   // id
		TaskId     int       `gorm:"task_id" json:"task_id"`         // 关联任务
		Format     string    `gorm:"format" json:"format"`           // 链接:dash-flv360
		Container  string    `gorm:"container" json:"container"`     // 类型:mp4
		Quality    string    `gorm:"quality" json:"quality"`         // 质量:流畅 360P
		Size       int       `gorm:"size" json:"size"`               // 任务大小
		Status     int       `gorm:"status" json:"status"`           // 任务状态 0未处理 1处理中 2获取信息 3获取失败
		CreateTime time.Time `gorm:"create_time" json:"create_time"` // 创建时间
		UpdateTime time.Time `gorm:"update_time" json:"update_time"` // 修改时间
	}
)

func NewTaskInfoModel(db *gorm.DB) TaskInfoModel {
	return &defaultTaskInfoModel{
		db: db,
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

func (d *defaultTaskInfoModel) Insert(data *TaskInfo) (int, error) {
	err := d.db.Omit("create_time", "update_time").Create(data).Error
	return data.Id, err
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
	if limit <= 0 {
		limit = 20
	}

	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	})
}

func (d *defaultTaskInfoModel) FindByIds(ids []int) ([]*TaskInfo, error) {
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	})
}

func (d *defaultTaskInfoModel) FindsByTaskID(taskId int) ([]*TaskInfo, error) {
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Where("task_id = ?", taskId)
	})
}

type TaskAndInfo struct {
	Task *Task
	Info map[int]*TaskInfo
}

// 查询info和关联的task
func (d *defaultTaskInfoModel) FindDetailsById(ids []int) (map[int]*TaskAndInfo, error) {
	infos, err := d.FindByIds(ids)
	if err != nil {
		return nil, err
	}
	taskIdM := make(map[int]struct{})
	taskIds := make([]int, 0, len(taskIdM))
	for _, info := range infos {
		if _, ok := taskIdM[info.TaskId]; !ok {
			taskIdM[info.TaskId] = struct{}{}
			taskIds = append(taskIds, info.TaskId)
		}

	}

	tasks, err := NewTaskModel(d.db).FindByIds(taskIds)
	if err != nil {
		return nil, err
	}

	res := make(map[int]*TaskAndInfo, 0)
	for _, task := range tasks {
		var t TaskAndInfo
		t.Task = task
		t.Info = make(map[int]*TaskInfo, 0)
		for _, info := range infos {
			if info.TaskId == task.Id {
				t.Info[info.Id] = info
			}
		}
		res[task.Id] = &t
	}
	return res, nil
}

// 回写状态
func (d *defaultTaskInfoModel) UpdateStatus(id int, Status int) error {
	return d.db.Model(&TaskInfo{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status": Status,
	}).Error
}

func (d *defaultTaskInfoModel) Count() (int, error) {
	var count int64
	err := d.db.Model(TaskInfo{}).Count(&count).Error
	return int(count), err
}
