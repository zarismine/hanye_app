package handler

import (
	"hanye/pkg/util"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/address/api/internal/logic"
	"hanye/service/address/api/internal/svc"
	"hanye/service/address/api/internal/types"
)

func deleteAddressByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteAddressByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		token, _ := util.ParseToken(r.Header.Get("Authorization"))
		req.UserId = token.Id
		l := logic.NewDeleteAddressByIdLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAddressById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
