package svc

import (
	"GoLearn/go_zero/api_jwt_mysql/model"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserModel(mysqlConn),
	}
}
