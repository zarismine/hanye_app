package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSetmealsByCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSetmealsByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSetmealsByCategoryIdLogic {
	return &GetSetmealsByCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSetmealsByCategoryIdLogic) GetSetmealsByCategoryId(in *pb.GetSetmealsByCategoryIdReq) (*pb.GetSetmealsByCategoryIdResp, error) {
	data, err := l.svcCtx.SetmealModel.FindSetmealsByCategoryId(l.ctx, in.CategoryId)
	resp := &pb.GetSetmealsByCategoryIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var setmeals []*pb.Setmeal
	err = copier.Copy(&setmeals, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Setmeals = setmeals
	resp.Code = e.SUCCESS
	return resp, nil
}
