package login

import (
	"api-permission/internal/utils/response"
	"net/http"

	"api-permission/internal/logic/auth/login"
	"api-permission/internal/svc"
	"api-permission/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginUsernameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			a := "error login"
			// set redis
			response.Response(w, a, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
