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
		var res types.ResultMsg
		if err != nil {
			res.Code = 1
			res.Msg = err.Error()
		} else {
			res.Result = resp
			res.Msg = "success"
		}
		httpx.OkJson(w, resp)
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
		resp, err := l.VideoDownload(req)
		var res types.ResultMsg
		if err != nil {
			res.Code = 1
			res.Msg = err.Error()
		} else {
			res.Result = resp
			res.Msg = "success"
		}
		httpx.OkJson(w, resp)
	}
}

