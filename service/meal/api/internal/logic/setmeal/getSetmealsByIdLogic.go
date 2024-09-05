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

type GetSetmealsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSetmealsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSetmealsByIdLogic {
	return &GetSetmealsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSetmealsByIdLogic) GetSetmealsById(req *types.GetSetmealsByIdReq) (resp *types.GetSetmealsByIdResp, err error) {
	data, _ := l.svcCtx.MealRpcClient.GetSetmealsById(l.ctx, &pb.GetSetmealsByIdReq{
		SetmealId: req.Id,
	})
	response := &types.GetSetmealsByIdResp{}
	if data.Code != e.SUCCESS {
		l.Logger.Error(err)
		response.Default = types.JsonError(int(data.Code))
		return response, nil
	}
	setmeals := new(types.SetmealWithDishes)
	err = copier.Copy(setmeals, &data.Setmeal)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Setmeals = setmeals
	return response, nil
}
