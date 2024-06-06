package handler

import (
	"GoLearn/common/response" // ①
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/logic"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/svc"
	"GoLearn/go_zero/api_jwt_mysql/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err) //②
	}
}
