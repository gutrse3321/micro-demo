package main

import (
	"demo/app/controller"
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 19:03
 * @Title:
 * --- --- ---
 * @Desc:
 */
var providerSet = wire.NewSet(
	controller.ProviderSet,
)
