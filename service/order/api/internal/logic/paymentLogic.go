package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/jsonx"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentLogic) Payment(req *types.PaymentReq) (resp *types.PaymentResp, err error) {
	resp = &types.PaymentResp{}
	reqRpc := &pb.PaymentReq{}
	err = copier.Copy(reqRpc, req)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	data, _ := l.svcCtx.OrderRpcClient.Payment(l.ctx, reqRpc)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	wsResp := types.WebsocketResp{
		Type:    1,
		Content: "您有新的订单",
		OrderId: data.Id,
	}
	jsonData, _ := jsonx.Marshal(wsResp)
	l.svcCtx.SendMsg(jsonData)
	resp.Default = types.JsonSuccess()
	return resp, nil
}
