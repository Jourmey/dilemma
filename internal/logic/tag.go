package logic

import (
	"context"
	"dilemma/internal/model"
	"dilemma/internal/svc"
	"dilemma/internal/types"
	"github.com/tal-tech/go-zero/core/logx"
)

type TagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	tagDB  model.TagModel
}

func NewTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) TagLogic {
	return TagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		tagDB:  model.NewTagModel(svcCtx.DB),
	}
}

func (l *TagLogic) Tag(req types.GetReq) ([]*model.Tag, error) {
	// 按照id查询
	if req.Id != 0 {
		t, err := l.tagDB.FindOne(req.Id)
		if err != nil {
			return nil, err
		} else {
			return []*model.Tag{t}, nil
		}
	}
	// 分页查询
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	return l.tagDB.Finds(req.PageSize, req.PageNo)
}
