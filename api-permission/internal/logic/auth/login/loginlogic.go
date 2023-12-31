package login

import (
	"api-permission/internal/utils/hashx"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"api-permission/internal/svc"
	"api-permission/internal/types"

	// "api-permission/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginUsernameReq) (resp *types.LoginUsernameResp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil && err != sqlc.ErrNotFound {
		fmt.Println("error find user")
		return nil, err
	}
	if !hashx.BcryptCheck(req.Password, user.Password) {
		return nil, errors.New("password or username error")
	}
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		user.Id)
	if err != nil {
		return nil, err
	}

	go func() {
		key := fmt.Sprintf("token_%s", user.Username)
		err = l.svcCtx.Redis.Setex(key, token, 60)
		fmt.Println("ssssssssss")
		if err != nil {
			panic(err)
		}
	}()

	return &types.LoginUsernameResp{
		Token: token,
	}, err
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// https://github.com/bxcodec/go-clean-archs
