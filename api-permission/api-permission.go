package main

import (
	"api-permission/internal/config"
	"api-permission/internal/handler"
	"api-permission/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-permission.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

/*
- redis là 1 db
- nó vừa lưu cả disk lẫn Ram cơ chế back up
- nó là key - value nosql
- như thằng Get Set
 + Set name Value : A
 + Get name = A
 map["name"] = "A"
 + Set Token

    err := rdb.Set(ctx, "token", "value", 0).Err()
    if err != nil {
        panic(err)
		return
    }

    val, err := rdb.Get(ctx, "token").Result()
    if err != nil {
        panic(err)
		return
    }
*/
