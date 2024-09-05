package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/user/api/internal/logic"
	"hanye/service/user/api/internal/svc"
	"hanye/service/user/api/internal/types"
)

func updateuserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateuserLogic(r.Context(), svcCtx)
		resp, err := l.Updateuser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
