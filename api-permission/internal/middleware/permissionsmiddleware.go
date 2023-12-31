package middleware

import (
	"encoding/json"
	"api-permission/internal/utils/errorsx"
	"api-permission/internal/utils/response"
	"api-permission/model"
	"net/http"
	

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type PermMenuAuthMiddleware struct {
	Redis *redis.Redis
}

func NewPermMenuAuthMiddleware(r *redis.Redis) *PermMenuAuthMiddleware {
	return &PermMenuAuthMiddleware{
		Redis: r,
	}
}

type MustParams struct {
	UserModel  		model.UsersModel
	RuleModel  		model.RulesModel
	RoleModel  		model.RolesModel
	UsersRoleModel  model.UsersRoleRelationModel
	RoleRulesModel  model.RoleRulesRelationModel
}

type PermissionsMiddleware struct {
	param MustParams
}

func NewPermissionsMiddleware(param MustParams) *PermissionsMiddleware {
	return &PermissionsMiddleware{param: param}
}

func (m *PermissionsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.String()
		userIdNumber := r.Context().Value("userId").(json.Number)
		userIdInt, _ := userIdNumber.Int64()
		if userIdInt <= 0 {
			response.Response(w,nil, errorsx.NewDefaultError("must param user_id"))
			return
		}
		roleIds, err := m.param.UsersRoleModel.FindAllByUserId(r.Context(), userIdInt)
		if err != nil {
			response.Response(w,nil, errorsx.NewDefaultError("user role does not exist" + err.Error()))
			return
		}
		roles, err := m.param.RoleModel.FindAll(r.Context(), roleIds)
		if err != nil {
			response.Response(w,nil, errorsx.NewDefaultError("user role does not exist"))
			return
		}
		if !IsAdministrator(roles) {
			ruleIds, err := m.param.RoleRulesModel.FindAllByRoleIds(r.Context(), roleIds)
			if err != nil {
				response.Response(w,nil, errorsx.NewDefaultError("no permission"))
				return
			}
			rules, err := m.param.RuleModel.FindAllByRules(r.Context(), ruleIds, uri)
			if err != nil || len(rules) <= 0 {
				response.Response(w,nil, errorsx.NewDefaultError("no permission api: " + uri))
				return
			}
		}
		next(w, r)
	}

}

func IsAdministrator(roles []model.Roles) bool {
	for _, role := range roles {
		if role.Slug == "administrator" {
			return true
		}
	}
	return false
}
// func (m *PermMenuAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if len(r.Header.Get("Authorization")) > 0 {
// 			userId := utils.GetUserId(r.Context())
// 			online, err := m.Redis.Get(globalkey.SysOnlineUserCachePrefix + strconv.FormatInt(userId, 10))
// 			if err != nil || online == "" {
// 				httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
// 				var erring any
// 				erring = "Auth fail"
// 				panic(erring)
// 			}

// 			uri := strings.Split(r.RequestURI, "?")
// 			is, err := m.Redis.Sismember(globalkey.SysPermMenuCachePrefix+strconv.FormatInt(userId, 10), uri[0])
// 			if err != nil || is != true {
// 				httpx.Error(w, errorx.NewDefaultError(errorx.NotPermMenuErrorCode))
// 			} else {
// 				next(w, r)
// 			}
// 		}
// 	}
// }