// Written by yijian on 2024/01/20
package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mooon-mailbox/model"
	"os"

	"mooon-mailbox/internal/config"
	"mooon-mailbox/internal/server"
	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon_mailbox"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	help       = flag.Bool("h", false, "Display a help message and exit")
	configFile = flag.String("f", "etc/mooon_mailbox.yaml", "Config file")

	// 格式：
	// "dbuser:dbpassword@tcp(dbhost:dbport)/dbname?charset=utf8mb3&parseTime=True&loc=Local"
	// 优先级高于配置中的 DSN
	dsn = flag.String("dsn", "", "MySQL data source name, example: --dsn='dbuser:dbpassword@tcp(dbhost:dbport)/dbname?charset=utf8mb3&parseTime=True&loc=Local'")
)

func main() {
	flag.Parse()
	if *help {
		usage()
		os.Exit(1)
	}
	if !checkParams() {
		usage()
		os.Exit(1)
	}

	var dbConn sqlx.SqlConn
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	if len(*dsn) == 0 && len(c.DB.DSN) == 0 {
		logx.Error("Both parameter[--dsn] and config[DB.DSN] are not set\n")
		os.Exit(1)
	}
	if len(*dsn) > 0 {
		dbConn = sqlx.NewMysql(*dsn)
	} else {
		dbConn = sqlx.NewMysql(c.DB.DSN)
	}
	ctx.MailboxModel = model.NewTMooonMailboxModel(dbConn, c.Cache)
	ctx.CachedConn = sqlc.NewConn(dbConn, c.Cache)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		mooon_mailbox.RegisterMooonMailboxServer(grpcServer, server.NewMooonMailboxServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting mailbox server at %s...\n", c.ListenOn)
	s.Start()
}

func usage() {
	flag.Usage()
}

func checkParams() bool {
	for {
		// configFile
		if len(*configFile) == 0 {
			fmt.Println("Parameter[-f] is not set")
			break
		}

		return true
	}

	return false
}
