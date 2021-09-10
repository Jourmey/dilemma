package logic

import (
	"context"
	"dilemma/internal/model"
	"dilemma/internal/svc"
	"dilemma/internal/types"
	"dilemma/tool"
	"dilemma/youget"
	"fmt"
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
			return l.buildVideo([]*model.Video{t}), nil
		}
	}

	// 分页查询
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	res, err := l.videoDB.Finds(req.PageSize, req.PageNo)
	if err != nil {
		return nil, err
	}

	return l.buildVideo(res), err

}

func (l *VideoLogic) buildVideo(res []*model.Video) []*model.Video {
	for i := 0; i < len(res); i++ {
		res[i].Path = fmt.Sprintf(fmt.Sprintf("%s:%d/%s", l.svcCtx.Config.Staticfile.Host, l.svcCtx.Config.Staticfile.Port, res[i].Path))
	}
	return res
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
		err = l.youGetDownload(task, info)
		if err != nil {
			l.Logger.Error("下载失败", err)
		}
	}()

	return err
}

func (l *VideoLogic) youGetDownload(task *model.Task, info *model.TaskInfo) error {
	l.Logger.Info("开始下载", tool.Json(task), tool.Json(info))

	y := youget.NewYouGet()
	site := "unknown"
	if task.Site != "" {
		site = task.Site
	}
	outputDir := fmt.Sprintf("%s/%d/%d", site, info.Id/100, info.Id%100)
	res, err := y.Download(task.Url, info.Format, fmt.Sprintf("/workspace/%s", outputDir))
	if err != nil {
		l.Logger.Error("youget", "下载失败", res, err)
		return err
	}

	video := model.Video{
		TaskInfoId: info.Id,
		Path:       outputDir,
	}

	_, err = l.videoDB.Insert(&video)
	if err != nil {
		l.Logger.Error("插入下载数据失败", err)
		return err
	}

	l.Logger.Info("下载成功")
	return nil
}
