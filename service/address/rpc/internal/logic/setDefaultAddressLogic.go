package logic

import (
	"context"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetDefaultAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetDefaultAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetDefaultAddressLogic {
	return &SetDefaultAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetDefaultAddressLogic) SetDefaultAddress(in *pb.SetDefaultAddressReq) (*pb.SetDefaultAddressResp, error) {
	err := l.svcCtx.AddressModel.SetDefaultAddress(l.ctx, in.UserId, in.AddressId)
	resp := &pb.SetDefaultAddressResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Code = e.SUCCESS
	return resp, nil
}
