// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zero-shop/app/goods/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/goods/list",
				Handler: listHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/goods/detail",
				Handler: detailHandler(serverCtx),
			},
		},
	)
}