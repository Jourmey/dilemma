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
	info, err := l.taskInfoDB.FindDetailsById(req.TaskInfoIds)
	if err != nil {
		return err
	}

	go func() {
		err = l.youGetDownload(info, l.svcCtx.Config.Staticfile.Root)
		if err != nil {
			l.Logger.Error("下载失败", err)
		}
	}()

	return err
}

func (l *VideoLogic) youGetDownload(info map[int]*model.TaskAndInfo, root string) error {
	l.Logger.Info("下载视频", tool.Start)

	y := youget.NewYouGet(l.ctx)
	// 任务维度
	for _, t := range info {
		site := "unknown"
		if t.Task.Site != "" {
			site = t.Task.Site
		}

		// 信息维度
		for _, i := range t.Info {
			err := l.taskInfoDB.UpdateStatus(i.Id, model.StatusRunning)
			if err != nil {
				return err
			}

			outputDir := fmt.Sprintf("%s/%d", site, i.Id)
			_, err = y.Download(t.Task.Url, i.Format, fmt.Sprintf("%s/%s", root, outputDir))

			status := model.StatusFailed
			if err == nil { // 成功
				_, err = l.videoDB.Insert(&model.Video{
					TaskInfoId: i.Id,
					Path:       outputDir,
				})
				if err != nil {
					return err
				}
				status = model.StatusSuccess
			}

			err = l.taskInfoDB.UpdateStatus(i.Id, status)
			if err != nil {
				return err
			}
		}
	}

	l.Logger.Info(tool.Success)
	return nil
}
