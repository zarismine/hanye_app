package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *pb.GetOrderByIdReq) (*pb.GetOrderByIdResp, error) {
	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
	resp := &pb.GetOrderByIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	if order.UserId != in.UserId {
		l.Logger.Error("无权限")
		resp.Code = e.ERROR_AUTH
		return resp, nil
	}
	orderDetails, err := l.svcCtx.OrderDetailModel.FindByOrderId(l.ctx, order.Id)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	orderResp := &pb.Order{}
	err = copier.Copy(orderResp, order)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	orderResp.OrderTime = order.OrderTime.Unix()
	orderResp.EstimatedDeliveryTime = order.EstimatedDeliveryTime.Unix()
	if order.CheckoutTime.Valid {
		orderResp.CheckoutTime = order.CheckoutTime.Time.Unix()
	}
	if order.CancelTime.Valid {
		orderResp.CancelTime = order.CancelTime.Time.Unix()
	}
	if order.DeliveryTime.Valid {
		orderResp.DeliveryTime = order.DeliveryTime.Time.Unix()
	}
	orderDetailResp := new([]*pb.OrderDetail)
	err = copier.Copy(orderDetailResp, orderDetails)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Order = orderResp
	resp.OrderDetailList = *orderDetailResp
	resp.Code = e.SUCCESS
	return resp, nil
}
