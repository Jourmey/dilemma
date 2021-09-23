package handler

import (
	"net/http"

	"dilemma/internal/logic"
	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func videoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideoLogic(r.Context(), ctx)
		resp, err := l.Video(req)
		httpx.OkJson(w, types.NewResultMsg(resp, err))
	}
}

func videoDownloadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoDownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideoLogic(r.Context(), ctx)
		err := l.VideoDownload(req)
		httpx.OkJson(w, types.NewResultMsg(nil, err))
	}
}


// 清理视频
func deleteVideoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideoLogic(r.Context(), ctx)
		err := l.DeleteVideo(req)
		httpx.OkJson(w, types.NewResultMsg(nil, err))
	}
}


func homepageInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVideoLogic(r.Context(), ctx)
		resp, err := l.HomepageInfo()
		httpx.OkJson(w, types.NewResultMsg(resp, err))
	}
}
