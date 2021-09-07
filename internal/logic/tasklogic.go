package logic

import (
	"context"
	"dilemma/internal/model"

	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskLogic {
	return TaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogic) Task(req types.GetReq) ([]*model.Task, error) {
	m := model.NewTaskModel(model.GetSqlConn())
	// 按照id查询
	if req.Id != 0 {
		t, err := m.FindOne(req.Id)
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
	return m.Finds(req.PageSize, req.PageNo)
}

func (l *TaskLogic) Create(req types.TaskCreatReq) error {
	// todo: add your logic here and delete this line

	return nil
}
