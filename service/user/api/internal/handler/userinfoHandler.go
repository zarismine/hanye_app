package handler

import (
	"errors"
	"hanye/pkg/util"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hanye/service/user/api/internal/logic"
	"hanye/service/user/api/internal/svc"
	"hanye/service/user/api/internal/types"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserinfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		token, _ := util.ParseToken(r.Header.Get("Authorization"))
		if token.Id != req.UserId {
			httpx.ErrorCtx(r.Context(), w, errors.New("无权限"))
			return
		}
		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
