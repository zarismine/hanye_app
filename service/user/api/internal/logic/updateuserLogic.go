package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/user/api/internal/svc"
	"hanye/service/user/api/internal/types"
	"hanye/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateuserLogic {
	return &UpdateuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuserLogic) Updateuser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	idNumber := req.Id
	res, _ := l.svcCtx.UserRpcClient.UpdateUser(l.ctx, &pb.UpdateUserReq{
		IdNumber: idNumber,
		Name:     req.Name,
		Phone:    req.Phone,
		Gender:   req.Gender,
		Pic:      req.Pic,
	})
	response := &types.UpdateUserResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	response.Default = types.JsonSuccess()
	return response, nil
}
