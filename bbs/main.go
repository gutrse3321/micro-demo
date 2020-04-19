package bbs

import (
	"demo/pkg/app"
	"demo/pkg/transports/http"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:20
 * @Title:
 * --- --- ---
 * @Desc:
 */

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app config error")
	}

	logger.Info("load application config success")
	return opt, err
}

func NewApp(opt *Options, logger *zap.Logger, httpServer *http.Server) (*app.Application, error) {
	application, err := app.New(opt.Name, logger, app.HttpServerOption(httpServer))
	if err != nil {
		return nil, errors.Wrap(err, "new application error")
	}

	return application, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)

//func main() {
//	createApiServer()
//	//go createRpcServer()
//	//select {}
//}
//
//func createApiServer() {
//	r := gin.New()
//	r.Use(gin.Logger())
//	//index := controller.
//
//	//v1 := r.Group("/")
//	//{
//	//	v1.POST("/getUserInfo", index.GetUserBaseInfo)
//	//}
//
//	r.Run(":4000")
//}
