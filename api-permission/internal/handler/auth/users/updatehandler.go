package users

import (
	"api-permission/common"
	"api-permission/internal/logic/auth/users"
	"api-permission/internal/svc"
	"api-permission/internal/types"
	"api-permission/internal/utils/errorsx"
	"api-permission/internal/utils/response"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("111111111111111111111111111111")
		var req types.UpdateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		token := r.Header.Get("Authorization")
		_, err := common.ChecktTokenRedis(svcCtx.Redis, req.Username, token)
		if err != nil {
			response.Response(w, "token invalid ! ", err)
			errorsx.NewDefaultError("ChecktTokenRedis")
			return
		}
		//if check //
		l := users.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		if err != nil {
			response.Response(w, nil, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
