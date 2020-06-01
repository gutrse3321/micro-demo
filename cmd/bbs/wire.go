//+build wireinject

package main

import (
	"demo/bbs"
	"demo/bbs/controller"
	"demo/bbs/provider"
	"demo/bbs/service"
	"github.com/google/wire"
	remote "github.com/gutrse3321/aki-remote"
	"github.com/gutrse3321/aki/pkg/app"
	"github.com/gutrse3321/aki/pkg/config"
	"github.com/gutrse3321/aki/pkg/log"
	"github.com/gutrse3321/aki/pkg/transports/http"
	"github.com/gutrse3321/aki/pkg/transports/rpc"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 17:58
 * @Title:
 * --- --- ---
 * @Desc:
 */
var wireSet = wire.NewSet(
	log.WireSet,
	config.WireSet,
	http.WireSet,
	rpc.WireServerSet,
	rpc.WireClientSet,
	provider.ProviderSet,
	remote.WireSet,
	bbs.ProviderSet,
	service.ProviderSet,
	controller.ProviderSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(wireSet))
}
