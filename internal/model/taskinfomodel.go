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
		FindDetailsById(ids []int) (map[int]*TaskAndInfo, error)
		UpdateStatus(id int, Status int) error
	}

	defaultTaskInfoModel struct {
		db    *gorm.DB
		table string
	}

	TaskInfo struct {
		Id         int       `db:"id"`          // id
		TaskId     int       `db:"task_id"`     // 关联任务
		Format     string    `db:"format"`      // 链接:dash-flv360
		Container  string    `db:"container"`   // 类型:mp4
		Quality    string    `db:"quality"`     // 质量:流畅 360P
		Size       int       `db:"size"`        // 任务大小
		Status     int       `db:"status"`      // 任务状态 0未处理 1处理中 2获取信息 3获取失败
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
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
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	})
}

func (d *defaultTaskInfoModel) FindByIds(ids []int) ([]*TaskInfo, error) {
	return d.finds(func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
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
