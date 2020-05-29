//+build wireinject

package main

import (
	"demo/ucenter"
	"demo/ucenter/provider"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/app"
	"github.com/gutrse3321/aki/pkg/config"
	"github.com/gutrse3321/aki/pkg/log"
	"github.com/gutrse3321/aki/pkg/transports/rpc"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 17:58
 * @Title:
 * --- --- ---
 * @Desc:
 */
var providerSet = wire.NewSet(
	log.WireSet,
	config.WireSet,
	rpc.WireServerSet,
	ucenter.ProviderSet,
	provider.ProviderSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
