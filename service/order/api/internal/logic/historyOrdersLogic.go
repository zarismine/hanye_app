package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"
	"slices"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryOrdersLogic {
	return &HistoryOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryOrdersLogic) HistoryOrders(req *types.HistoryOrdersReq) (resp *types.HistoryOrdersResp, err error) {
	resp = &types.HistoryOrdersResp{}
	reqRpc := &pb.HistoryOrdersReq{}
	_ = copier.Copy(reqRpc, req)
	data, _ := l.svcCtx.OrderRpcClient.HistoryOrders(l.ctx, reqRpc)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	slices.Reverse(data.OrderWithDetails)
	err = copier.Copy(&resp.Orders.Records, data.OrderWithDetails)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	for i := 0; i < len(resp.Orders.Records); i++ {
		err = copier.Copy(&resp.Orders.Records[i].OrderDetailList, &data.OrderWithDetails[i].OrderDetailList)
		if err != nil {
			resp.Default = types.JsonError(e.ERROR_COPY)
			return resp, nil
		}
	}
	resp.Default = types.JsonSuccess()
	resp.Orders.Total = int64(len(data.OrderWithDetails))
	return resp, nil
}
