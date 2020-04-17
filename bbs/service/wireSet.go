package service

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 18:50
 * @Title:
 * --- --- ---
 * @Desc:
 */
var ProviderSet = wire.NewSet(NewUserService)
