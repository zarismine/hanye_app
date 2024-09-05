package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/user/rpc/pb"

	"hanye/service/user/api/internal/svc"
	"hanye/service/user/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserinfoReq) (resp *types.UserinfoResp, err error) {
	UserData, _ := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	response := &types.UserinfoResp{}
	if UserData.Code != e.SUCCESS {
		response.Default = types.JsonError(int(UserData.Code))
		return response, nil
	}
	response.Default = types.JsonSuccess()
	var user types.User
	err = copier.Copy(&user, UserData.User)
	if err != nil {
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.User = user
	return response, nil
}
