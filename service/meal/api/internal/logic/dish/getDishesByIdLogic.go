package dish

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/meal/rpc/pb"

	"hanye/service/meal/api/internal/svc"
	"hanye/service/meal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDishesByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDishesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishesByIdLogic {
	return &GetDishesByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDishesByIdLogic) GetDishesById(req *types.GetDishesByIdReq) (resp *types.GetDishesByIdResp, err error) {
	data, _ := l.svcCtx.MealRpcClient.GetDishesById(l.ctx, &pb.GetDishesByIdReq{
		DishId: req.DishId,
	})
	response := &types.GetDishesByIdResp{}
	if data.Code != e.SUCCESS {
		l.Logger.Error(err)
		response.Default = types.JsonError(int(data.Code))
		return response, nil
	}
	dish := new(types.Dish)
	err = copier.Copy(dish, &data.Dish)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Dishs = *dish
	return response, nil
}
