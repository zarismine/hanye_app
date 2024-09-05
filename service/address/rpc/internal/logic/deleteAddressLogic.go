package logic

import (
	"context"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressLogic {
	return &DeleteAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAddressLogic) DeleteAddress(in *pb.DeleteAddressReq) (*pb.DeleteAddressResp, error) {
	resp := &pb.DeleteAddressResp{}
	err := l.svcCtx.AddressModel.Delete(l.ctx, in.AddressId)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Code = e.SUCCESS
	return resp, nil
}
