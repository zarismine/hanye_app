package logic

import (
	"context"
	"hanye/pkg/e"

	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnPayOrderCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnPayOrderCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnPayOrderCountLogic {
	return &UnPayOrderCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnPayOrderCountLogic) UnPayOrderCount(in *pb.UnPayOrderCountReq) (*pb.UnPayOrderCountResp, error) {
	count, err := l.svcCtx.OrderModel.GetCountByUserId(l.ctx, in.UserId)
	resp := &pb.UnPayOrderCountResp{}
	if err != nil {
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	resp.Count = count
	resp.Code = e.SUCCESS
	return resp, nil
}
