package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"hanye/pkg/e"
	"hanye/service/cart/model"
	"hanye/service/cart/rpc/internal/svc"
	"hanye/service/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubCartLogic {
	return &AddSubCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSubCartLogic) AddSubCart(in *pb.AddSubCartReq) (*pb.AddSubCartResp, error) {
	response := &pb.AddSubCartResp{}
	cart := new(model.Cart)
	err := copier.Copy(cart, in)
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_COPY
		return response, nil
	}
	var data *model.Cart
	if in.SetmealId != 0 {
		data, err = l.svcCtx.CartModel.FindOneByUserIdSetmealIdFlavor(l.ctx, in.UserId, in.SetmealId, in.DishFlavor)
		if errors.Is(err, sqlc.ErrNotFound) {
			_, err = l.svcCtx.CartModel.Insert(l.ctx, cart)
			return &pb.AddSubCartResp{
				Code: e.SUCCESS,
			}, nil
		}
	} else if in.DishId != 0 {
		data, err = l.svcCtx.CartModel.FindOneByUserIdDishIdFlavor(l.ctx, in.UserId, in.DishId, in.DishFlavor)
		if errors.Is(err, sqlc.ErrNotFound) {
			_, err = l.svcCtx.CartModel.Insert(l.ctx, cart)
			return &pb.AddSubCartResp{
				Code: e.SUCCESS,
			}, nil
		}
	} else {
		response.Code = e.INVALID_PARAMS
		return response, nil
	}
	data.Number += in.Number
	if data.Number == 0 {
		err = l.svcCtx.CartModel.Delete(l.ctx, data.Id)
		if err != nil {
			l.Logger.Error(err)
			response.Code = e.ERROR_DATABASE
			return response, nil
		}
		return &pb.AddSubCartResp{
			Code: e.SUCCESS,
		}, nil
	}
	if data.Number < 0 {
		return &pb.AddSubCartResp{
			Code: e.ERROR,
		}, nil
	}
	data.Amount = in.Amount
	data.DishFlavor = in.DishFlavor
	err = l.svcCtx.CartModel.Update(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	return &pb.AddSubCartResp{
		Code: e.SUCCESS,
	}, nil
}
