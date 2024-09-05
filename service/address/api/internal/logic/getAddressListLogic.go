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

type GetAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressListLogic {
	return &GetAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressListLogic) GetAddressList(req *types.GetAddressListReq) (resp *types.GetAddressListResp, err error) {
	res, _ := l.svcCtx.AddressRpcClient.GetAddressList(l.ctx, &pb.GetAddressListReq{
		UserId: req.UserId,
	})
	response := &types.GetAddressListResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	addresses := new([]*types.AddressWithId)
	err = copier.Copy(addresses, &res.Addresses)
	if err != nil {
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Addresses = *addresses
	return response, nil
}
