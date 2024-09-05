package handler

import (
	"hanye/pkg/util"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/order/api/internal/logic"
	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"
)

func paymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PaymentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		token, _ := util.ParseToken(r.Header.Get("Authorization"))
		req.UserId = token.Id
		l := logic.NewPaymentLogic(r.Context(), svcCtx)
		resp, err := l.Payment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
