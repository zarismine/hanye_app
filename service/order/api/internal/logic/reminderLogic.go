package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"hanye/pkg/e"
	"hanye/service/order/rpc/pb"

	"hanye/service/order/api/internal/svc"
	"hanye/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReminderLogic {
	return &ReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReminderLogic) Reminder(req *types.ReminderReq) (resp *types.ReminderResp, err error) {
	resp = &types.ReminderResp{}
	reqRpc := &pb.ReOrderReq{
		OrderId: req.Id,
		UserId:  req.UserId,
	}
	data, _ := l.svcCtx.OrderRpcClient.ReOrder(l.ctx, reqRpc)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	wsResp := types.WebsocketResp{
		Type:    2,
		Content: "用户在催单",
		OrderId: req.Id,
	}
	jsonData, _ := jsonx.Marshal(wsResp)
	l.svcCtx.SendMsg(jsonData)
	resp.Default = types.JsonSuccess()
	return resp, nil
}
