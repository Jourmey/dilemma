package logic

import (
	"context"
	"dilemma/internal/model"
	"dilemma/internal/svc"
	"dilemma/internal/types"
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
	c := model.GetSqlConn(svcCtx.Config.Mysql)
	return TaskLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		taskInfoDB: model.NewTaskInfoModel(c),
		taskDB:     model.NewTaskModel(c),
	}
}

func (l *TaskLogic) Task(req types.GetReq) ([]*model.Task, error) {
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

func (l *TaskLogic) Create(req types.TaskCreatReq) error {
	// 1.插入任务
	data := model.Task{
		Url: req.Url,
		Tag: req.Tag,
	}
	taskResult, err := l.taskDB.Insert(data)
	if err != nil {
		return err
	}
	id, err := taskResult.LastInsertId()
	if err != nil {
		return err
	}

	// 2.获取任务信息
	y := youget.NewYouGet()
	yInfo, err := y.Info(req.Url)
	if err != nil {
		return err
	}

	// 3.更新信息
	_, err = l.taskDB.UpdateStatus(int(id), 1, yInfo.Title, yInfo.Site)
	if err != nil {
		return err
	}

	// 4.插入info信息
	for format, stream := range yInfo.Streams {
		taskInfo := model.TaskInfo{
			TaskId:    int(id),
			Format:    format,
			Container: stream.Container,
			Quality:   stream.Quality,
			Size:      stream.Size,
		}
		_, err = l.taskInfoDB.Insert(taskInfo)
		if err != nil {
			continue
		}
	}

	return err
}
