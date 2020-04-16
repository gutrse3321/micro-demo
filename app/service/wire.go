// +build wireinject

package service

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 18:46
 * @Title:
 * --- --- ---
 * @Desc:
 */

func CreateDetailsService() (IUserService, error) {
	panic(wire.Build(ProviderSet))
}
