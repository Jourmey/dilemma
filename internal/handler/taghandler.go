package handler

import (
	"dilemma/internal/logic"
	"dilemma/internal/svc"
	"dilemma/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func tagHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTagLogic(r.Context(), ctx)
		resp, err := l.Tag(req)

		httpx.OkJson(w, types.NewResultMsg(resp, err))
	}
}
