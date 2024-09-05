package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/address/rpc/pb"

	"hanye/service/address/api/internal/svc"
	"hanye/service/address/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultAddressLogic {
	return &SetDefaultAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetDefaultAddressLogic) SetDefaultAddress(req *types.SetDefaultAddressReq) (resp *types.SetDefaultAddressResp, err error) {
	res, _ := l.svcCtx.AddressRpcClient.SetDefaultAddress(l.ctx, &pb.SetDefaultAddressReq{
		AddressId: req.Id,
		UserId:    req.UserId,
	})
	response := &types.SetDefaultAddressResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	response.Default = types.JsonSuccess()
	return response, nil
}
