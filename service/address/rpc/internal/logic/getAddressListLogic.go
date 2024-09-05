package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressListLogic {
	return &GetAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressListLogic) GetAddressList(in *pb.GetAddressListReq) (*pb.GetAddressListResp, error) {
	data, err := l.svcCtx.AddressModel.FindAddressesByUserId(l.ctx, in.UserId)
	response := &pb.GetAddressListResp{}
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	var resp []*pb.Address
	err = copier.Copy(&resp, data)
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_COPY
		return response, nil
	}
	response.Code = e.SUCCESS
	response.Addresses = resp
	return response, nil
}
