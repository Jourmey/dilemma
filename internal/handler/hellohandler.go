package handler

import (
	"net/http"

	"dilemma/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func helloHandler(_ *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, "hi.dilemma!")
	}
}
