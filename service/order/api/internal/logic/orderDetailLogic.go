package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/dates"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailReq) (resp *types.OrderDetailResp, err error) {
	data, _ := l.svcCtx.OrderRpcClient.GetOrderById(l.ctx, &pb.GetOrderByIdReq{
		OrderId: req.Id,
		UserId:  req.UserId,
	})
	resp = &types.OrderDetailResp{}
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	err = copier.Copy(&resp.Orders, data.Order)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	resp.Orders.OrderTime = dates.Int642Time(data.Order.OrderTime)
	resp.Orders.EstimatedDeliveryTime = dates.Int642Time(data.Order.EstimatedDeliveryTime)
	if data.Order.CancelTime != 0 {
		resp.Orders.CancelTime = dates.Int642Time(data.Order.CancelTime)
	} else {
		resp.Orders.CancelTime = ""
	}
	if data.Order.DeliveryTime != 0 {
		resp.Orders.DeliveryTime = dates.Int642Time(data.Order.DeliveryTime)
	} else {
		resp.Orders.DeliveryTime = ""
	}
	if data.Order.CheckoutTime != 0 {
		resp.Orders.CheckoutTime = dates.Int642Time(data.Order.CheckoutTime)
	} else {
		resp.Orders.CheckoutTime = ""
	}
	err = copier.Copy(&resp.Orders.OrderDetailList, &data.OrderDetailList)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	return resp, nil
}
