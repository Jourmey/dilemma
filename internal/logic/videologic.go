package logic

import (
	"context"
	"dilemma/internal/model"

	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) VideoLogic {
	return VideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoLogic) Video(req types.GetReq) ([]model.Video, error) {
	// todo: add your logic here and delete this line

	return []model.Video{}, nil
}

func (l *VideoLogic) VideoDownload(req types.VideoDownloadReq) error {
	// todo: add your logic here and delete this line

	return nil
}
