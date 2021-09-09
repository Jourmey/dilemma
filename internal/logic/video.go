package logic

import (
	"context"
	"dilemma/internal/model"
	"dilemma/youget"

	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	taskDB     model.TaskModel
	taskInfoDB model.TaskInfoModel
	videoDB    model.VideoModel
}

func NewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) VideoLogic {
	return VideoLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		taskDB:     model.NewTaskModel(svcCtx.DB),
		taskInfoDB: model.NewTaskInfoModel(svcCtx.DB),
		videoDB:    model.NewVideoModel(svcCtx.DB),
	}
}

func (l *VideoLogic) Video(req types.GetReq) ([]*model.Video, error) {
	// 按照id查询
	if req.Id != 0 {
		t, err := l.videoDB.FindOne(req.Id)
		if err != nil {
			return nil, err
		} else {
			return []*model.Video{t}, nil
		}
	}

	// 分页查询
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	return l.videoDB.Finds(req.PageSize, req.PageNo)
}

func (l *VideoLogic) VideoDownload(req types.VideoDownloadReq) error {
	info, err := l.taskInfoDB.FindOne(req.TaskInfoId)
	if err != nil {
		return err
	}

	task, err := l.taskDB.FindOne(info.TaskId)
	if err != nil {
		return err
	}

	go func() {
		err = l.YouGetDownload(task, info)
		if err != nil {
			l.Logger.Error("下载失败", err)
		}
	}()

	return err
}

func (l *VideoLogic) YouGetDownload(task *model.Task, info *model.TaskInfo) error {
	l.Logger.Info("开始下载", task, info)

	y := youget.NewYouGet()
	res, err := y.Download(task.Url, info.Format, "/workspace")
	if err != nil {
		l.Logger.Error("下载失败", res, err)
		return err
	}

	l.Logger.Info("下载成功", res)
	return nil
}
