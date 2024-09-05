package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSetmealsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSetmealsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSetmealsByIdLogic {
	return &GetSetmealsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSetmealsByIdLogic) GetSetmealsById(in *pb.GetSetmealsByIdReq) (*pb.GetSetmealsByIdResp, error) {
	data, err := l.svcCtx.SetmealModel.FindOne(l.ctx, in.SetmealId)
	resp := &pb.GetSetmealsByIdResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	dishes, err := l.svcCtx.SetmealDishModel.FindBySetmealId(l.ctx, in.SetmealId)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var setmealdishes []*pb.SetmealDishes
	err = copier.Copy(&setmealdishes, dishes)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	setmealWithDishes := &pb.SetmealWithDishes{
		Id:     data.Id,
		Name:   data.Name,
		Pic:    data.Pic,
		Detail: data.Detail,
		Status: data.Status,
		//UpdateTime:    data.UpdateTime,
		Price:         data.Price,
		CategoryId:    data.CategoryId,
		SetmealDishes: setmealdishes,
	}
	resp.Code = e.SUCCESS
	resp.Setmeal = setmealWithDishes
	return resp, nil
}
