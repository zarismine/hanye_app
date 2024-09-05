package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDefaultAddressLogic {
	return &GetDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDefaultAddressLogic) GetDefaultAddress(in *pb.GetDefaultAddressReq) (*pb.GetDefaultAddressResp, error) {
	data, err := l.svcCtx.AddressModel.GetDefaultAddress(l.ctx, in.UserId)
	resp := &pb.GetDefaultAddressResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.SUCCESS
		return resp, nil
	}
	address := new(pb.Address)
	err = copier.Copy(address, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Address = address
	resp.Code = e.SUCCESS
	return resp, nil
}
