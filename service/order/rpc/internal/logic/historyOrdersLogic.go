package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/order/model"

	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHistoryOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryOrdersLogic {
	return &HistoryOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HistoryOrdersLogic) HistoryOrders(in *pb.HistoryOrdersReq) (*pb.HistoryOrdersResp, error) {
	resp := &pb.HistoryOrdersResp{}
	req := &model.PageReq{}
	err := copier.Copy(req, in)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	orders, err := l.svcCtx.OrderModel.QueryCondition(l.ctx, req)
	if errors.Is(err, model.ErrNotFound) {
		resp.Code = e.SUCCESS
		return resp, nil
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var ordersResp []*pb.OrderWithDetail
	err = copier.Copy(&ordersResp, orders)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	for _, v := range ordersResp {
		orderDetails := new([]*pb.OrderDetail)
		orderDetail, _ := l.svcCtx.OrderDetailModel.FindByOrderId(l.ctx, v.Id)
		err = copier.Copy(orderDetails, orderDetail)
		if err != nil {
			resp.Code = e.ERROR_COPY
			return resp, nil
		}
		v.OrderDetailList = *orderDetails
	}
	resp.Code = e.SUCCESS
	resp.OrderWithDetails = ordersResp
	return resp, nil
}
