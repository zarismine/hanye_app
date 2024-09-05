package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDishesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDishesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishesByIdLogic {
	return &GetDishesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDishesByIdLogic) GetDishesById(in *pb.GetDishesByIdReq) (*pb.GetDishesByIdResp, error) {
	data, err := l.svcCtx.DishModel.FindOne(l.ctx, in.DishId)
	resp := &pb.GetDishesByIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var dish pb.Dish
	err = copier.Copy(&dish, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	flavorData, err := l.svcCtx.DishFlavorModel.FindByDishId(l.ctx, in.DishId)
	var flavors []*pb.Flavor
	err = copier.Copy(&flavors, flavorData)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	dish.Flavors = flavors
	resp.Dish = &dish
	resp.Code = e.SUCCESS
	return resp, nil
}
