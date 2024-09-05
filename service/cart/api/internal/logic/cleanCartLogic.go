package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/cart/rpc/pb"

	"hanye/service/cart/api/internal/svc"
	"hanye/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCleanCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanCartLogic {
	return &CleanCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanCartLogic) CleanCart(req *types.CleanCartReq) (resp *types.CleanCartResp, err error) {
	data, _ := l.svcCtx.CartRpcClient.CleanCart(l.ctx, &pb.CleanCartReq{
		UserId: req.UserId,
	})
	resp = &types.CleanCartResp{}
	l.Logger.Info(data)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	return resp, nil
}
