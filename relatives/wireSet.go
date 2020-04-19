package relatives

import (
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 17:23
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(
	NewUCenterRemote,
)
