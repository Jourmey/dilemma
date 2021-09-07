package logic

import (
	"context"

	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TaskinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) TaskinfoLogic {
	return TaskinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskinfoLogic) Taskinfo(req types.GetReq) (types.ResultMsg, error) {
	// todo: add your logic here and delete this line

	return types.ResultMsg{}, nil
}
