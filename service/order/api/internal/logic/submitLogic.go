package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"hanye/pkg/dates"
	"hanye/pkg/e"
	addresspb "hanye/service/address/rpc/pb"
	cartpb "hanye/service/cart/rpc/pb"
	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"
	"hanye/service/order/rpc/pb"
	userpb "hanye/service/user/rpc/pb"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req *types.SubmitReq) (resp *types.SubmitResp, err error) {
	resp = &types.SubmitResp{}
	order := &pb.SubmitReq{}
	err = copier.Copy(order, req)
	l.Logger.Info(order.UserId)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	user, _ := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &userpb.GetUserInfoReq{
		Id: req.UserId,
	})
	if user.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(user.Code))
		return resp, nil
	}
	order.UserName = user.User.Name
	address, _ := l.svcCtx.AddressRpcClient.GetAddress(l.ctx, &addresspb.GetAddressReq{
		AddressId: req.AddressId,
		UserId:    req.UserId,
	})
	if address.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(address.Code))
		return resp, nil
	}
	order.Address = addresspb.Address2string(address.Address)
	order.AddressBookId = address.Address.Id
	order.Consignee = address.Address.Consignee
	order.Phone = address.Address.Phone
	order.Status = 1
	order.EstimatedDeliveryTime = dates.Time2Int64(req.EstimatedDeliveryTime)
	carts, _ := l.svcCtx.CartRpcClient.GetCartList(l.ctx, &cartpb.GetCartListReq{
		UserId: req.UserId,
	})
	if carts.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(carts.Code))
		return resp, nil
	}
	orderDetail := new([]*pb.OrderDetail)
	err = copier.Copy(orderDetail, &carts.Carts)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	order.OrderDetailList = *orderDetail
	data, _ := l.svcCtx.OrderRpcClient.Submit(l.ctx, order)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	err = copier.Copy(&resp.Data, data)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	resp.Data.OrderTime = dates.Int642Time(data.OrderTime)
	return resp, nil
}
