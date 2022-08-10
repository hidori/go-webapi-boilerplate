package main

import (
	"fmt"
	"os"

	"github.com/hidori/go-webapi-boilerplate/go/internal/config"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/api"
	"github.com/hidori/go-webapi-boilerplate/go/internal/infrastructure/log"
)

var logger = log.GetLogger()

// main は、API サーバのエントリポイントです。
func main() {
	// 構成情報の新規インスタンスを取得する
	config, err := config.NewConfig(os.Getenv)
	if err != nil {
		logger.Panicf("fail to config.Config(): err=%v", err)
	}

	// API サーバの新規インスタンスを取得する
	server, err := api.NewServer(config)
	if err != nil {
		logger.Panicf("fail to api.NewServer(): err=%v", err)
	}

	// API サーバを開始する
	err = server.Start(fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		logger.Panicf("fail to server.Start(): err=%v", err)
	}
}
