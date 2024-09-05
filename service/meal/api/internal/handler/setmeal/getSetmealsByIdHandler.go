package setmeal

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/meal/api/internal/logic/setmeal"
	"hanye/service/meal/api/internal/svc"
	"hanye/service/meal/api/internal/types"
)

func GetSetmealsByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSetmealsByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := setmeal.NewGetSetmealsByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetSetmealsById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
