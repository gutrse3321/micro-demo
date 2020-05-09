package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 15:29
 * @Title: 配置处理初始化
 * --- --- ---
 * @Desc:
 */

/**
初始化viper
*/
func New(path string) (*viper.Viper, error) {
	var err error

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err = v.ReadInConfig(); err != nil {
		fmt.Println("use config file ->", v.ConfigFileUsed())
		return nil, err
	}

	return v, err
}

var ProviderSet = wire.NewSet(New)
