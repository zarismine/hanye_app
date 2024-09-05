package setmeal

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/meal/rpc/pb"

	"hanye/service/meal/api/internal/svc"
	"hanye/service/meal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDishSimplesByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDishSimplesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishSimplesByIdLogic {
	return &GetDishSimplesByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDishSimplesByIdLogic) GetDishSimplesById(req *types.GetDishSimplesByIdReq) (resp *types.GetDishSimplesByIdResp, err error) {
	l.Logger.Info("参数", req)
	data, _ := l.svcCtx.MealRpcClient.GetDishSimplesById(l.ctx, &pb.GetDishSimplesByIdReq{
		SetmealId: req.Id,
	})
	response := &types.GetDishSimplesByIdResp{}
	if data.Code != e.SUCCESS {
		l.Logger.Error(err)
		response.Default = types.JsonError(int(data.Code))
		return response, nil
	}
	dishSimples := new([]*types.DishSimple)
	err = copier.Copy(dishSimples, &data.Dishes)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.DishSimples = *dishSimples
	return response, nil
}
