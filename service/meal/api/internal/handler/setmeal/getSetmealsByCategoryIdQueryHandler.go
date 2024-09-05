package setmeal

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/meal/api/internal/logic/setmeal"
	"hanye/service/meal/api/internal/svc"
	"hanye/service/meal/api/internal/types"
)

func GetSetmealsByCategoryIdQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSetmealsByCategoryIdQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := setmeal.NewGetSetmealsByCategoryIdQueryLogic(r.Context(), svcCtx)
		resp, err := l.GetSetmealsByCategoryIdQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
