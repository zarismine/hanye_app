package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/address/model"
	"hanye/service/address/rpc/internal/svc"
	"hanye/service/address/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAddressLogic) AddAddress(in *pb.AddAddressReq) (*pb.AddAddressResp, error) {
	l.Logger.Info("AddAddress方法请求参数：", in)
	resp := &pb.AddAddressResp{}
	_, err := l.svcCtx.AddressModel.Insert(l.ctx, &model.AddressBook{
		//Id:           snowId,
		UserId:       in.Address.UserId,
		Consignee:    in.Address.Consignee,
		Gender:       in.Address.Gender,
		Phone:        in.Address.Phone,
		ProvinceCode: in.Address.ProvinceCode,
		ProvinceName: in.Address.ProvinceName,
		CityCode:     in.Address.CityCode,
		DistrictCode: in.Address.DistrictCode,
		CityName:     in.Address.CityName,
		DistrictName: in.Address.DistrictName,
		Detail:       in.Address.Detail,
		Label:        in.Address.Label,
		IsDefault:    0,
	})
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Code = e.SUCCESS
	return resp, nil
}
