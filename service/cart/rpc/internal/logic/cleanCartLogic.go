package logic

import (
	"context"
	"hanye/pkg/e"

	"hanye/service/cart/rpc/internal/svc"
	"hanye/service/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanCartLogic {
	return &CleanCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CleanCartLogic) CleanCart(in *pb.CleanCartReq) (*pb.CleanCartResp, error) {
	err := l.svcCtx.CartModel.DeleteByUserId(l.ctx, in.UserId)
	response := &pb.CleanCartResp{}
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	response.Code = e.SUCCESS
	return response, nil
}
