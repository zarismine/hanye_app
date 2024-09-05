package logic

import (
	"context"
	"database/sql"
	"errors"
	"hanye/pkg/e"
	"hanye/service/order/model"
	"time"

	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLogic) Cancel(in *pb.CancelReq) (*pb.CancelResp, error) {
	resp := &pb.CancelResp{}
	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
	if errors.Is(err, model.ErrNotFound) {
		resp.Code = e.SUCCESS
		return resp, nil
	}
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	if order.UserId != in.UserId {
		resp.Code = e.ERROR_AUTH
		return resp, nil
	}
	if order.Status == 1 {
		order.Status = 6
	}
	if order.Status == 2 {
		order.Status = 6
	}
	order.PayStatus = 2
	order.CancelTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	order.CancelReason = "用户取消订单"
	err = l.svcCtx.OrderModel.Update(l.ctx, order)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Code = e.SUCCESS
	return resp, nil
}
