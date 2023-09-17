package svc

import (
	"api-permission/internal/config"
	"api-permission/internal/middleware"
	"api-permission/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UsersModel
	RuleModel      model.RulesModel
	RoleModel      model.RolesModel
	UsersRoleModel model.UsersRoleRelationModel
	RoleRulesModel model.RoleRulesRelationModel
	Permissions    rest.Middleware
	Redis          *redis.Redis
	PermMenuAuth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	userModel := model.NewUsersModel(conn)
	ruleModel := model.NewRulesModel(conn)
	roleModel := model.NewRolesModel(conn)
	usersRoleModel := model.NewUsersRoleRelationModel(conn)
	roleRulesModel := model.NewRoleRulesRelationModel(conn)
	redisClient := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config: c,
		Redis:  redisClient,
		// PermMenuAuth:       middleware.NewPermMenuAuthMiddleware(redisClient).Handle,
		Permissions: middleware.NewPermissionsMiddleware(middleware.MustParams{
			UserModel:      userModel,
			RuleModel:      ruleModel,
			RoleModel:      roleModel,
			UsersRoleModel: usersRoleModel,
			RoleRulesModel: roleRulesModel,
		}).Handle,
		UserModel:      userModel,
		RuleModel:      ruleModel,
		RoleModel:      roleModel,
		UsersRoleModel: usersRoleModel,
		RoleRulesModel: roleRulesModel,
	}
}

func (l *ServiceContext) ChecktTokenRedis(secretKey string) (string, error) {

	return "", nil
}
