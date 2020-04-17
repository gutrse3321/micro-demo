package main

import (
	"demo/bbs"
	"demo/bbs/controller"
	"demo/bbs/service"
	"demo/pkg/app"
	"demo/pkg/config"
	"demo/pkg/log"
	"demo/pkg/transports/http"
	"demo/pkg/transports/rpc"
	"demo/relatives"
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 17:58
 * @Title:
 * --- --- ---
 * @Desc:
 */
var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	rpc.ClientProviderSet,
	relatives.ProviderSet,
	bbs.ProviderSet,
	service.ProviderSet,
	controller.ProviderSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
