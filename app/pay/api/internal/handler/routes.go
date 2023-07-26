// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zero-shop/app/pay/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/pay/notify",
				Handler: notifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create/third",
				Handler: createHandler(serverCtx),
			},
		},
	)
}
