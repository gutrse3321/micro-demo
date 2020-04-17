package controller

import (
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 18:43
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(
	NewIndexController,
	CreateInitControllersFn,
)
