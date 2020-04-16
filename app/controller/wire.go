// +build wireinject

package controller

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 19:00
 * @Title:
 * --- --- ---
 * @Desc:
 */
func CreateIndexController() (Controller, error) {
	panic(wire.Build(ProviderSet))
}
