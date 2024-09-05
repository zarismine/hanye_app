package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/cart/rpc/pb"

	"hanye/service/cart/api/internal/svc"
	"hanye/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartListLogic) GetCartList(req *types.GetCartListReq) (resp *types.GetCartListResp, err error) {
	data, _ := l.svcCtx.CartRpcClient.GetCartList(l.ctx, &pb.GetCartListReq{
		UserId: req.UserId,
	})
	resp = &types.GetCartListResp{}
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	err = copier.Copy(&resp.Carts, &data.Carts)
	if err != nil {
		resp.Default = types.JsonError(e.ERROR_COPY)
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	return resp, nil
}
