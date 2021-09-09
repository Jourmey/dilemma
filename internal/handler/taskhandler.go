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
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskLogic(r.Context(), ctx)
		resp, err := l.Task(req)

		httpx.OkJson(w, types.NewResultMsg(resp, err))
	}
}

func taskCreatHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskCreatReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskLogic(r.Context(), ctx)
		err := l.Create(req)

		httpx.OkJson(w, types.NewResultMsg(nil, err))
	}
}
