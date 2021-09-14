package logic

import (
	"context"
	"dilemma/internal/model"

	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TaskinfoLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	taskInfoDB model.TaskInfoModel
}

func NewTaskinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskinfoLogic {
	return TaskinfoLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		taskInfoDB: model.NewTaskInfoModel(svcCtx.DB),
	}
}

func (l *TaskinfoLogic) Taskinfo(req types.TaskInfoGetReq) ([]*model.TaskInfo, error) {
	// 按照id查询
	if req.Id != 0 {
		t, err := l.taskInfoDB.FindOne(req.Id)
		if err != nil {
			return nil, err
		} else {
			return []*model.TaskInfo{t}, nil
		}
	}

	if req.TaskId != 0 {
		return l.taskInfoDB.FindsByTaskID(req.TaskId)
	}

	// 分页查询
	return l.taskInfoDB.Finds(req.PageSize, req.PageNo)
}
