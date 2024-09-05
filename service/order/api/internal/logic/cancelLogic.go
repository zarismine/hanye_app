package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req *types.CancelOrderReq) (resp *types.CancelOrderResp, err error) {
	resp = &types.CancelOrderResp{}
	reqRpc := &pb.CancelReq{
		OrderId: req.Id,
		UserId:  req.UserId,
	}
	data, _ := l.svcCtx.OrderRpcClient.Cancel(l.ctx, reqRpc)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	return resp, nil
}
