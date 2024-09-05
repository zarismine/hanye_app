package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/cart/rpc/internal/svc"
	"hanye/service/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartListLogic) GetCartList(in *pb.GetCartListReq) (*pb.GetCartListResp, error) {
	data, err := l.svcCtx.CartModel.FindByUserId(l.ctx, in.UserId)
	response := &pb.GetCartListResp{}
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	cartList := new([]*pb.Cart)
	err = copier.Copy(cartList, data)
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_COPY
		return response, nil
	}
	response.Code = e.SUCCESS
	response.Carts = *cartList
	return response, nil
}
