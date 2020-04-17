package main

import "flag"

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 17:57
 * @Title:
 * --- --- ---
 * @Desc:
 */

var configFile = flag.String("f", "configs/user.yml", "set config file which viper will loading")

func main() {
	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
