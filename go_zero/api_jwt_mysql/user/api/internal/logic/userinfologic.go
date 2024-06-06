package logic

import (
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/svc"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/types"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	uid, _ := userId.Int64()
	U, _ := l.svcCtx.UsersModel.FindOne(
		context.Background(),
		uid,
	)
	return &types.UserInfoResponse{
		Id:       uint(U.Id),
		UserName: U.Username,
		Password: U.Password,
	}, nil
}
