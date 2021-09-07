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
		taskInfoDB: model.NewTaskInfoModel(model.GetSqlConn(svcCtx.Config.Mysql)),
	}
}

func (l *TaskinfoLogic) Taskinfo(req types.GetReq) ([]*model.TaskInfo, error) {
	// 按照id查询
	if req.Id != 0 {
		t, err := l.taskInfoDB.FindOne(req.Id)
		if err != nil {
			return nil, err
		} else {
			return []*model.TaskInfo{t}, nil
		}
	}

	// 分页查询
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	return l.taskInfoDB.Finds(req.PageSize, req.PageNo)
}
