package handler

import (
	"net/http"

	"dilemma/internal/logic"
	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func taskHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskLogic(r.Context(), ctx)
		resp, err := l.Task(req)

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

func taskCreatHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskCreatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskLogic(r.Context(), ctx)
		err := l.Create(req)

		var res types.ResultMsg
		if err != nil {
			res.Code = 1
			res.Msg = err.Error()
		} else {
			res.Result = nil
			res.Msg = "success"
		}
		httpx.OkJson(w, res)
	}
}
