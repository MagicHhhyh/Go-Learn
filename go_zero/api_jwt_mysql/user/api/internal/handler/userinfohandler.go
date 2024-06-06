package handler

import (
	"GoLearn/common/response" // ①
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/logic"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/svc"
	"net/http"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err) //②

	}
}
