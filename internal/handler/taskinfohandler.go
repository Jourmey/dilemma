package handler

import (
	"net/http"

	"dilemma/internal/logic"
	"dilemma/internal/svc"
	"dilemma/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func taskinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskInfoGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskinfoLogic(r.Context(), ctx)
		resp, err := l.Taskinfo(req)
		httpx.OkJson(w, types.NewResultMsg(resp, err))
	}
}
