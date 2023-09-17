package users

import (
	"api-permission/internal/utils/response"
	"fmt"
	"net/http"

	"api-permission/internal/logic/auth/users"
	"api-permission/internal/svc"
	"api-permission/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("11111111111111111")
		var req types.CreateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		fmt.Println("111111111111111112")
		l := users.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			response.Response(w, err, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
