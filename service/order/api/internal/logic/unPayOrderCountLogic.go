package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnPayOrderCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnPayOrderCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnPayOrderCountLogic {
	return &UnPayOrderCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnPayOrderCountLogic) UnPayOrderCount(req *types.UnPayOrderCountReq) (resp *types.UnPayOrderCountResp, err error) {
	resp = &types.UnPayOrderCountResp{}
	data, _ := l.svcCtx.OrderRpcClient.UnPayOrderCount(l.ctx, &pb.UnPayOrderCountReq{
		UserId: req.UserId,
	})
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	resp.Count = data.Count
	return resp, nil
}
