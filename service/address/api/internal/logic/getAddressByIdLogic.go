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

type GetAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressByIdLogic {
	return &GetAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAddressByIdLogic) GetAddressById(req *types.GetAddressByIdReq) (resp *types.GetAddressByIdResp, err error) {
	res, _ := l.svcCtx.AddressRpcClient.GetAddress(l.ctx, &pb.GetAddressReq{
		AddressId: req.Id,
		UserId:    req.UserId,
	})
	response := &types.GetAddressByIdResp{}
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	addressInfo := new(types.AddressWithId)
	err = copier.Copy(addressInfo, res.Address)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.AddressInfo = addressInfo
	return response, nil
}
