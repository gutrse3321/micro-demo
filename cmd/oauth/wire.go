//+build wireinject

package main

import (
	"demo/oauth"
	"github.com/google/wire"
	remote "github.com/gutrse3321/aki-remote"
	"github.com/gutrse3321/aki/pkg/app"
	"github.com/gutrse3321/aki/pkg/config"
	"github.com/gutrse3321/aki/pkg/log"
	"github.com/gutrse3321/aki/pkg/transports/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 16:14
 * @Title:
 * --- --- ---
 * @Desc:
 */
var wireSet = wire.NewSet(
	log.WireSet,
	config.WireSet,
	http.WireSet,
	remote.WireSet,
	oauth.WireSet,
)

func CreateApp(configPath string) (*app.Application, error) {
	panic(wire.Build(wireSet))
}
