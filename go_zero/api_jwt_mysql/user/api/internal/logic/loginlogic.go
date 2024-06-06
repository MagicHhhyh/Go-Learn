package logic

import (
	"GoLearn/common/jwt"
	"GoLearn/go_zero/api_jwt_mysql/model"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/svc"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	if U, _ := l.svcCtx.UsersModel.FindOneByUsername(context.Background(), req.UserName); U == nil {
		l.svcCtx.UsersModel.Insert(context.Background(), &model.User{
			Username: req.UserName,
			Password: req.Password,
		})
	} else if U.Password != req.Password {
		return "wrong password", nil
	}
	token, err := jwt.GenToken(
		jwt.JwtPayLoad{
			UserID:   1,
			Username: req.UserName,
			Role:     1,
		},
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
	)
	if err != nil {
		return "token ERROR", err
	}
	return token, nil
}
