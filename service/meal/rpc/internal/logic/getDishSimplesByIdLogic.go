package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDishSimplesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDishSimplesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishSimplesByIdLogic {
	return &GetDishSimplesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDishSimplesByIdLogic) GetDishSimplesById(in *pb.GetDishSimplesByIdReq) (*pb.GetDishSimplesByIdResp, error) {
	data, err := l.svcCtx.SetmealDishModel.FindBySetmealId(l.ctx, in.SetmealId)
	resp := &pb.GetDishSimplesByIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var dishSimples []*pb.DishSimple
	for _, v := range *data {
		detail, pic, _ := l.svcCtx.DishModel.FindDetailAndPicById(l.ctx, v.DishId)
		dishSimple := &pb.DishSimple{
			Name:   v.Name,
			Pic:    pic,
			Detail: detail,
			Copies: v.Copies,
		}
		dishSimples = append(dishSimples, dishSimple)
	}
	resp.Dishes = dishSimples
	resp.Code = e.SUCCESS
	return resp, nil
}
