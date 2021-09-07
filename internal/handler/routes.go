// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"dilemma/internal/svc"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/task",
				Handler: taskHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/task/creat",
				Handler: taskCreatHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/taskinfo",
				Handler: taskinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video",
				Handler: videoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/video/download",
				Handler: videoDownloadHandler(serverCtx),
			},
		},
	)
}