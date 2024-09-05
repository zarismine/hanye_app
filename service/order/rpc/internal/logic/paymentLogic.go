package logic

import (
	"context"
	"database/sql"
	"errors"
	"hanye/pkg/e"
	"hanye/service/order/model"
	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaymentLogic) Payment(in *pb.PaymentReq) (*pb.PaymentResp, error) {
	resp := &pb.PaymentResp{}
	order, err := l.svcCtx.OrderModel.FindOneByNumber(l.ctx, in.OrderNumber)
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
	order.Status = 2
	order.PayMethod = in.PayMethod
	order.CheckoutTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	order.PayStatus = 1
	err = l.svcCtx.OrderModel.Update(l.ctx, order)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Id = order.Id
	resp.Code = e.SUCCESS
	return resp, nil
}
