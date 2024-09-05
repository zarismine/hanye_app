package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/order/api/internal/logic"
	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"
)

func websocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWebsocketLogic(r.Context(), svcCtx)
		err := l.Websocket(w, r, req.UserId)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
