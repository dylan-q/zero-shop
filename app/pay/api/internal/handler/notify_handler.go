package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-shop/app/pay/api/internal/logic"
	"zero-shop/app/pay/api/internal/svc"
	"zero-shop/app/pay/api/internal/types"
)

func notifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNotifyLogic(r.Context(), svcCtx)
		resp, err := l.Notify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
