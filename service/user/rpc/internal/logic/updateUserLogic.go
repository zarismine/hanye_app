package logic

import (
	"context"
	"database/sql"
	"hanye/pkg/e"
	"hanye/service/user/rpc/internal/svc"
	"hanye/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	l.Logger.Info("UpdateUser方法请求参数：", in)
	phone := sql.NullString{
		String: in.Phone,
		Valid:  true,
	}

	name := sql.NullString{
		String: in.Name,
		Valid:  true,
	}
	pic := sql.NullString{
		String: in.Pic,
		Valid:  true,
	}
	user, err := l.svcCtx.UserinfoModel.FindOneById(l.ctx, in.IdNumber)
	response := &pb.UpdateUserResp{}
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	user.Name = name
	user.Pic = pic
	user.Phone = phone
	user.Gender = sql.NullInt64{Int64: in.Gender, Valid: true}
	err = l.svcCtx.UserinfoModel.UpdateUnique(l.ctx, user)
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	response.Code = e.SUCCESS
	return response, nil
}
