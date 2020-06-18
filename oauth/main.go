package oauth

import (
	"demo/oauth/auth"
	"demo/oauth/endpoints"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/app"
	akiHttp "github.com/gutrse3321/aki/pkg/transports/http"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewOptions(v *viper.Viper, logger *zap.Logger) (*app.Options, error) {
	var err error

	opt := &app.Options{}
	if err = v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app config error")
	}

	logger.Info("load application config success")
	return opt, err
}

func NewApp(opt *app.Options, logger *zap.Logger, httpServer *akiHttp.Server) (*app.Application, error) {
	application, err := app.New(opt, logger, app.HttpServerOption(httpServer))
	if err != nil {
		return nil, errors.Wrap(err, "new application error")
	}

	return application, nil
}

var WireSet = wire.NewSet(
	NewApp,
	NewOptions,
	auth.WireSet,
	endpoints.WireSet,
)
