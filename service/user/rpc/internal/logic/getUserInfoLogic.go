package logic

import (
	"context"
	"hanye/pkg/e"
	"hanye/service/user/rpc/internal/svc"
	"hanye/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l.Logger.Info(in)
	idNumber := in.Id
	user, err := l.svcCtx.UserinfoModel.FindOneById(l.ctx, idNumber)
	response := &pb.GetUserInfoResp{}
	if err != nil {
		l.Logger.Error(err)
		response.Code = e.ERROR_DATABASE
		return response, nil
	}
	respUser := &pb.User{
		Id:         user.Id,
		Phone:      user.Phone.String,
		Name:       user.Name.String,
		Gender:     user.Gender.Int64,
		IdNumber:   user.IdNumber.String,
		Pic:        user.Pic.String,
		Openid:     user.Openid,
		CreateTime: user.CreateTime.String(),
	}
	response.Code = e.SUCCESS
	response.User = respUser
	return response, nil
}
