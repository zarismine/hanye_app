package logic

import (
	"context"
	"database/sql"
	"hanye/pkg/common"
	"hanye/pkg/e"
	"hanye/pkg/util"
	"hanye/service/user/model"
	"strconv"

	"hanye/service/user/rpc/internal/svc"
	"hanye/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	l.Logger.Info("SaveUser方法请求参数：", in)
	snowflake, err := util.NewSnowflake(common.UserApiMachineId)
	if err != nil {
		return &pb.AddUserResp{
			Code: e.ERROR_SNOW,
		}, nil
	}
	snowId := snowflake.Generate()
	idNumber := sql.NullString{
		String: strconv.FormatInt(snowId, 10),
		Valid:  true,
	}
	phone := sql.NullString{
		String: in.Phone,
		Valid:  true,
	}

	Name := sql.NullString{
		String: common.NamePrefix + in.Phone,
		Valid:  true,
	}
	defaultPic := sql.NullString{
		String: l.svcCtx.DefaultPic,
		Valid:  true,
	}
	user, _ := l.svcCtx.UserinfoModel.FindOneByOpenId(l.ctx, in.Openid)
	if user != nil {
		return &pb.AddUserResp{
			Code: e.SUCCESS,
			Id:   user.IdNumber.String,
		}, nil
	}
	_, err = l.svcCtx.UserinfoModel.Insert(l.ctx, &model.User{
		Name:     Name,
		IdNumber: idNumber,
		Openid:   in.Openid,
		Pic:      defaultPic,
		Phone:    phone,
		Gender:   sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		return &pb.AddUserResp{
			Code: e.ERROR_ADD_USER,
		}, nil
	}
	return &pb.AddUserResp{
		Code: e.SUCCESS,
		Id:   idNumber.String,
	}, nil
}
