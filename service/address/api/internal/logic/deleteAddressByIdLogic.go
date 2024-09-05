package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/address/rpc/pb"

	"hanye/service/address/api/internal/svc"
	"hanye/service/address/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressByIdLogic {
	return &DeleteAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAddressByIdLogic) DeleteAddressById(req *types.DeleteAddressByIdReq) (resp *types.DeleteAddressByIdResp, err error) {
	res, _ := l.svcCtx.AddressRpcClient.DeleteAddress(l.ctx, &pb.DeleteAddressReq{
		AddressId: req.Id,
	})
	response := &types.DeleteAddressByIdResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	response.Default = types.JsonSuccess()
	return response, nil
}
