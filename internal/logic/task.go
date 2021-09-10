package logic

import (
	"context"
	"dilemma/internal/model"
	"dilemma/internal/svc"
	"dilemma/internal/types"
	"dilemma/tool"
	"dilemma/youget"
	"github.com/tal-tech/go-zero/core/logx"
)

type TaskLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	taskDB     model.TaskModel
	taskInfoDB model.TaskInfoModel
}

func NewTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskLogic {
	return TaskLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		taskInfoDB: model.NewTaskInfoModel(svcCtx.DB),
		taskDB:     model.NewTaskModel(svcCtx.DB),
	}
}

func (l *TaskLogic) Task(req types.GetReq) ([]*model.Task, error) {
	l.Logger.Info("获取任务信息,请求参数为:", req, "开始处理")
	// 按照id查询
	if req.Id != 0 {
		t, err := l.taskDB.FindOne(req.Id)
		if err != nil {
			return nil, err
		} else {
			return []*model.Task{t}, nil
		}
	}
	// 分页查询
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	return l.taskDB.Finds(req.PageSize, req.PageNo)
}

// 插入任务
func (l *TaskLogic) Create(req types.TaskCreatReq) error {
	data := model.Task{
		Url: req.Url,
		Tag: req.Tag,
	}
	id, err := l.taskDB.Insert(&data)
	if err != nil {
		l.Logger.Error("创建任务失败", err)
		return err
	}
	l.Logger.Info("创建任务", "任务id:", id)

	go func() {
		err = l.YouGetInfo(data)
		if err != nil {
			l.Logger.Error("处理任务失败", err)
		}
	}()
	return nil
}

// 获取任务信息
func (l *TaskLogic) YouGetInfo(task model.Task) error {
	l.Logger.Info("获取任务信息", tool.Start, "参数:", task.Url)

	// 更新信息
	err := l.taskDB.UpdateStatus(task.Id, model.StatusRunning, "", "")
	if err != nil {
		return err
	}

	y := youget.NewYouGet(l.ctx)
	yInfo, err := y.Info(task.Url)
	if err != nil {
		// 更新信息
		err = l.taskDB.UpdateStatus(task.Id, model.StatusFailed, "", "")
		if err != nil {
			return err
		}
		l.Logger.Error(tool.Failed, err)
		return err
	}

	// 更新信息
	err = l.taskDB.UpdateStatus(task.Id, model.StatusSuccess, yInfo.Title, yInfo.Site)
	if err != nil {
		return err
	}

	// 插入info信息
	for format, stream := range yInfo.Streams {
		taskInfo := model.TaskInfo{
			TaskId:    task.Id,
			Format:    format,
			Container: stream.Container,
			Quality:   stream.Quality,
			Size:      stream.Size,
		}
		_, err = l.taskInfoDB.Insert(&taskInfo)
		if err != nil {
			continue
		}
	}

	l.Logger.Info(tool.Success)
	return nil
}
