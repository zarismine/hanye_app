package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressLogic) UpdateAddress(in *pb.UpdateAddressReq) (*pb.UpdateAddressResp, error) {
	address, err := l.svcCtx.AddressModel.FindOne(l.ctx, in.Address.Id)
	resp := &pb.UpdateAddressResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	if address.UserId != in.UserId {
		//l.Logger.Error(address.UserId, in.UserId)
		resp.Code = e.ERROR_AUTH
		return resp, nil
	}
	isDefault := address.IsDefault
	err = copier.Copy(address, in.Address)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	address.IsDefault = isDefault
	err = l.svcCtx.AddressModel.Update(l.ctx, address)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Code = e.SUCCESS
	return resp, nil
}
