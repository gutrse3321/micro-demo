package endpoints

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 16:29
 * @Title:
 * --- --- ---
 * @Desc:
 */

var WireSet = wire.NewSet(
	NewLoginEndpoint,
	CreateInitControllersFn,
)
