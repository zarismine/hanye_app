package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/address/rpc/pb"

	"hanye/service/address/api/internal/svc"
	"hanye/service/address/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDefaultAddressLogic {
	return &GetDefaultAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDefaultAddressLogic) GetDefaultAddress(req *types.GetDefaultAddressReq) (resp *types.GetDefaultAddressResp, err error) {
	res, _ := l.svcCtx.AddressRpcClient.GetDefaultAddress(l.ctx, &pb.GetDefaultAddressReq{
		UserId: req.UserId,
	})
	response := &types.GetDefaultAddressResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	if res.Address == nil {
		response.Default = types.JsonError(e.ERROR_NOT_EXIST)
		return response, nil
	}
	address := new(types.AddressWithId)
	err = copier.Copy(address, &res.Address)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.AddressInfo = address
	return response, nil
}
