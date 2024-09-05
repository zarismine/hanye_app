package handler

import (
	"hanye/pkg/util"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/cart/api/internal/logic"
	"hanye/service/cart/api/internal/svc"
	"hanye/service/cart/api/internal/types"
)

func subCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddSubCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		token, _ := util.ParseToken(r.Header.Get("Authorization"))
		req.UserId = token.Id
		l := logic.NewSubCartLogic(r.Context(), svcCtx)
		resp, err := l.SubCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
