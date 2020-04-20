package ucenter

import (
	"demo/pkg/app"
	"demo/pkg/transports/rpc"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 17:43
 * @Title:
 * --- --- ---
 * @Desc:
 */
func NewOptions(v *viper.Viper, logger *zap.Logger) (*app.Options, error) {
	var err error

	opt := &app.Options{}
	if err = v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app config error")
	}

	logger.Info("load application config success")
	return opt, err
}

func NewApp(opt *app.Options, logger *zap.Logger, rpcServer *rpc.Server) (*app.Application, error) {
	application, err := app.New(opt, logger, app.RpcServerOption(rpcServer))
	if err != nil {
		return nil, errors.Wrap(err, "new application error")
	}

	return application, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)
