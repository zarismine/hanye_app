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

type UpdateAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressByIdLogic {
	return &UpdateAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAddressByIdLogic) UpdateAddressById(req *types.UpdateAddressByIdReq) (resp *types.UpdateAddressByIdResp, err error) {
	address := new(pb.Address)
	response := &types.UpdateAddressByIdResp{}
	err = copier.Copy(address, &req)
	//l.Logger.Info(address, req)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	res, _ := l.svcCtx.AddressRpcClient.UpdateAddress(l.ctx, &pb.UpdateAddressReq{
		Address: address,
		UserId:  req.UserId,
	})
	if res.Code != e.SUCCESS {
		response.Default = types.JsonError(int(res.Code))
		return response, nil
	}
	return response, nil
}
