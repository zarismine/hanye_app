package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/common"
	"hanye/pkg/e"
	"hanye/pkg/util"
	"hanye/service/order/model"
	"strconv"
	"time"

	"hanye/service/order/rpc/internal/svc"
	"hanye/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitLogic) Submit(in *pb.SubmitReq) (*pb.SubmitResp, error) {
	order := &model.Orders{}
	resp := &pb.SubmitResp{}
	err := copier.Copy(order, in)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	snowflake, err := util.NewSnowflake(common.UserApiMachineId)
	if err != nil {
		resp.Code = e.ERROR_SNOW
		return resp, nil
	}
	snowId := snowflake.Generate()
	order.Number = strconv.Itoa(int(snowId))
	order.OrderTime = time.Now()
	order.PayMethod = 0
	order.EstimatedDeliveryTime = time.Unix(in.EstimatedDeliveryTime, 0)
	ret, err := l.svcCtx.OrderModel.Insert(l.ctx, order)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	retId, _ := ret.LastInsertId()
	orderDetail := new([]*model.OrderDetail)
	err = copier.Copy(orderDetail, in.OrderDetailList)
	if err != nil {
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	for _, v := range *orderDetail {
		v.OrderId = retId
		_, err = l.svcCtx.OrderDetailModel.Insert(l.ctx, v)
		if err != nil {
			resp.Code = e.ERROR_DATABASE
			return resp, nil
		}
	}

	return &pb.SubmitResp{
		Code:        e.SUCCESS,
		Id:          retId,
		OrderTime:   order.OrderTime.Unix(),
		OrderAmount: order.Amount,
	}, nil
}
