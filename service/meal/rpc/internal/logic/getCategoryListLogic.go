package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"hanye/pkg/e"

	"hanye/service/meal/rpc/internal/svc"
	"hanye/service/meal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryListLogic) GetCategoryList(in *pb.GetCategoryListReq) (*pb.GetCategoryListResp, error) {
	data, err := l.svcCtx.CategoryModel.GetCategories(l.ctx)
	resp := &pb.GetCategoryListResp{}
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_DATABASE
		return resp, nil
	}
	var categories []*pb.Category
	err = copier.Copy(&categories, data)
	if err != nil {
		l.Logger.Error(err)
		resp.Code = e.ERROR_COPY
		return resp, nil
	}
	resp.Categories = categories
	resp.Code = e.SUCCESS
	return resp, nil
}
