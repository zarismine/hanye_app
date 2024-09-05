package logic

import (
	"context"
	"errors"
	"hanye/pkg/e"
	"hanye/service/order/model"
	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReOrderLogic {
	return &ReOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReOrderLogic) ReOrder(in *pb.ReOrderReq) (*pb.ReOrderResp, error) {
	resp := &pb.ReOrderResp{}
	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
	if errors.Is(err, model.ErrNotFound) {
		resp.Code = e.ERROR
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
	resp.Code = e.SUCCESS
	return resp, nil
}
