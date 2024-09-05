package dish

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/meal/api/internal/logic/dish"
	"hanye/service/meal/api/internal/svc"
	"hanye/service/meal/api/internal/types"
)

func GetDishesByCategoryIdQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDishesByCategoryIdQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dish.NewGetDishesByCategoryIdQueryLogic(r.Context(), svcCtx)
		resp, err := l.GetDishesByCategoryIdQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
