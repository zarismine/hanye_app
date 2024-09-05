package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReOrderLogic {
	return &ReOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReOrderLogic) ReOrder(req *types.ReOrderReq) (resp *types.ReOrderResp, err error) {
	resp = &types.ReOrderResp{}
	reqRpc := &pb.ReOrderReq{
		OrderId: req.Id,
		UserId:  req.UserId,
	}
	data, _ := l.svcCtx.OrderRpcClient.ReOrder(l.ctx, reqRpc)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	return resp, nil
}
