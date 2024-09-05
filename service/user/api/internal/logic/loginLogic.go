package logic

import (
	"context"
	"hanye/pkg/common"
	"hanye/pkg/e"
	"hanye/pkg/util"
	"hanye/service/user/rpc/pb"

	"hanye/service/user/api/internal/svc"
	"hanye/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	UserData, _ := l.svcCtx.UserRpcClient.AddUser(l.ctx, &pb.AddUserReq{
		Phone:  common.DefaultPhone,
		Openid: req.Code,
	})
	response := &types.LoginResp{}
	if UserData.Code != e.SUCCESS {
		response.Default = types.JsonError(int(UserData.Code))
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Data.Token, err = util.GenerateToken(req.Code, UserData.Id)
	if err != nil {
		response.Default = types.JsonError(e.ERROR_AUTH_TOKEN)
		return response, nil
	}
	response.Data.Id = UserData.Id
	response.Data.OpenId = req.Code
	return response, nil
}
