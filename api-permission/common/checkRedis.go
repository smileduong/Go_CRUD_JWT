package common

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func ChecktTokenRedis(redis *redis.Redis, secretKey string, tokeCilent string) (string, error) {

	key := fmt.Sprintf("token_%s", secretKey)
	token, err := redis.Get(key)
	// fmt.Println(key)

	fmt.Println(token)

	fmt.Println(tokeCilent)

	if err != nil {
		panic(err)
	}

	if tokeCilent != token {
		return "", fmt.Errorf("token invalid !")
	}
	return token, nil
}
