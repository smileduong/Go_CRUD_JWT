package users

import (
	"api-permission/common"
	"api-permission/internal/utils/response"
	"net/http"

	"api-permission/internal/logic/auth/users"
	"api-permission/internal/svc"
	"api-permission/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ViewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := users.NewViewLogic(r.Context(), svcCtx)
		resp, err := l.View(&req)

		token := r.Header.Get("Authorization")
		_, err = common.ChecktTokenRedis(svcCtx.Redis, resp.Username, token)
		if err != nil {
			response.Response(w, "token invalid ! ", err)
			return
		}
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
