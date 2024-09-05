package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressLogic {
	return &GetAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressLogic) GetAddress(in *pb.GetAddressReq) (*pb.GetAddressResp, error) {
	l.Logger.Info(in)
	data, err := l.svcCtx.AddressModel.FindOne(l.ctx, in.AddressId)
	resp := &pb.GetAddressResp{}
	if data == nil && errors.Is(err, sqlx.ErrNotFound) {
		//l.Logger.Error(err)
		resp.Code = e.SUCCESS
		return resp, nil
	}
	if data == nil && !errors.Is(err, sqlx.ErrNotFound) {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	if data.UserId != in.UserId {
		resp.Code = e.ERROR_AUTH
		return resp, nil
	}
	address := new(pb.Address)
	err = copier.Copy(address, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Code = e.SUCCESS
	resp.Address = address
	return resp, nil
}
