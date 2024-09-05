package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDishesByCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDishesByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishesByCategoryIdLogic {
	return &GetDishesByCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDishesByCategoryIdLogic) GetDishesByCategoryId(in *pb.GetDishesByCategoryIdReq) (*pb.GetDishesByCategoryIdResp, error) {
	data, err := l.svcCtx.DishModel.FindDishesByCategoryId(l.ctx, in.CategoryId)
	resp := &pb.GetDishesByCategoryIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var dishes []*pb.Dish
	err = copier.Copy(&dishes, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Dishes = dishes
	resp.Code = e.SUCCESS
	return resp, nil
}
