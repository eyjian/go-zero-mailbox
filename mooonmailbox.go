// Written by yijian on 2024/01/20
package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mooon-mailbox/model"
	"os"

	"mooon-mailbox/internal/config"
	"mooon-mailbox/internal/server"
	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon-mailbox"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	help       = flag.Bool("h", false, "Display a help message and exit")
	configFile = flag.String("f", "etc/mooonmailbox.yaml", "Config file")

	// 格式：
	// "dbuser:dbpassword@tcp(dbhost:dbport)/dbname?charset=utf8mb3&parseTime=True&loc=Local"
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

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	dbConn := sqlx.NewMysql(*dsn)
	ctx.MailboxModel = model.NewTMooonMailboxModel(dbConn, c.CacheConf)
	ctx.CachedConn = sqlc.NewConn(dbConn, c.CacheConf)

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
		if *configFile == "" {
			fmt.Println("Parameter[-f] is not set")
			break
		}

		// dsn
		if *dsn == "" {
			fmt.Println("Parameter[-dsn] is not set")
			break
		}

		return true
	}

	return false
}
