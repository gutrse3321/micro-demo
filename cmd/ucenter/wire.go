//+build wireinject

package main

import (
	"demo/pkg/app"
	"demo/pkg/config"
	"demo/pkg/log"
	"demo/pkg/transports/rpc"
	"demo/ucenter"
	"demo/ucenter/provider"
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
	rpc.ServerProviderSet,
	ucenter.ProviderSet,
	provider.ProviderSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
