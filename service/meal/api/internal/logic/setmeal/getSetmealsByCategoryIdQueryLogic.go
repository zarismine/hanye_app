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

type GetSetmealsByCategoryIdQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSetmealsByCategoryIdQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSetmealsByCategoryIdQueryLogic {
	return &GetSetmealsByCategoryIdQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSetmealsByCategoryIdQueryLogic) GetSetmealsByCategoryIdQuery(req *types.GetSetmealsByCategoryIdQueryReq) (resp *types.GetSetmealsByCategoryIdResp, err error) {
	data, _ := l.svcCtx.MealRpcClient.GetSetmealsByCategoryId(l.ctx, &pb.GetSetmealsByCategoryIdReq{
		CategoryId: req.CategoryId,
	})
	response := &types.GetSetmealsByCategoryIdResp{}
	if data.Code != e.SUCCESS {
		l.Logger.Error(err)
		response.Default = types.JsonError(int(data.Code))
		return response, nil
	}
	setmeals := new([]*types.Setmeal)
	err = copier.Copy(setmeals, &data.Setmeals)
	if err != nil {
		l.Logger.Error(err)
		response.Default = types.JsonError(e.ERROR_COPY)
		return response, nil
	}
	response.Default = types.JsonSuccess()
	response.Setmeals = *setmeals
	return response, nil
}
