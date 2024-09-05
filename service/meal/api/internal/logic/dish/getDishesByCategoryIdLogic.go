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

type GetDishesByCategoryIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDishesByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDishesByCategoryIdLogic {
	return &GetDishesByCategoryIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDishesByCategoryIdLogic) GetDishesByCategoryId(req *types.GetDishesByCategoryIdReq) (resp *types.GetDishesByCategoryIdResp, err error) {
	data, _ := l.svcCtx.MealRpcClient.GetDishesByCategoryId(l.ctx, &pb.GetDishesByCategoryIdReq{
		CategoryId: req.CategoryId,
	})
	response := &types.GetDishesByCategoryIdResp{}
	if data.Code != e.SUCCESS {
		l.Logger.Error(err)
		response.Default = types.JsonError(int(data.Code))
		return response, nil
	}
	dishes := new([]*types.Dish)
	err = copier.Copy(dishes, &data.Dishes)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Dishes = *dishes
	return response, nil
}
