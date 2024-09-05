package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"
	"hanye/service/cart/api/internal/svc"
	"hanye/service/cart/api/internal/types"
	"hanye/service/cart/rpc/pb"
	pb1 "hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartLogic) AddCart(req *types.AddSubCartReq) (resp *types.AddSubCartResp, err error) {
	cart := &pb.AddSubCartReq{}
	cart.Number = 1
	cart.DishFlavor = req.DishFlavor
	resp = &types.AddSubCartResp{}
	if req.DishId != 0 {
		res, _ := l.svcCtx.MealRpcClient.GetDishesById(l.ctx, &pb1.GetDishesByIdReq{
			DishId: req.DishId,
		})
		if res.Code != e.SUCCESS {
			resp.Default = types.JsonError(int(res.Code))
			return resp, nil
		}
		err = copier.Copy(cart, res.Dish)
		if err != nil {
			resp.Default = types.JsonError(e.ERROR_COPY)
			return resp, nil
		}
		cart.Amount = res.Dish.Price
		l.Logger.Info(cart)
	} else {
		res, _ := l.svcCtx.MealRpcClient.GetSetmealsById(l.ctx, &pb1.GetSetmealsByIdReq{
			SetmealId: req.SetmealId,
		})
		if res.Code != e.SUCCESS {
			resp.Default = types.JsonError(int(res.Code))
			return resp, nil
		}
		err = copier.Copy(cart, res.Setmeal)
		if err != nil {
			resp.Default = types.JsonError(e.ERROR_COPY)
			return resp, nil
		}
		cart.Amount = res.Setmeal.Price
		l.Logger.Info(cart)
	}
	cart.SetmealId = req.SetmealId
	cart.DishId = req.DishId
	cart.UserId = req.UserId
	data, _ := l.svcCtx.CartRpcClient.AddSubCart(l.ctx, cart)
	if data.Code != e.SUCCESS {
		resp.Default = types.JsonError(int(data.Code))
		return resp, nil
	}
	resp.Default = types.JsonSuccess()
	return resp, nil
}
