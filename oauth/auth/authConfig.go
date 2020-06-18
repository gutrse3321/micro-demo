package auth

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 15:57
 * @Title:
 * --- --- ---
 * @Desc:
 */

type OAuthOptions struct {
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
	RedisOptions    RedisOpts
	Clients         []OAuthClients
}

type RedisOpts struct {
	Addr     string
	Password string
	Db       int
}

type OAuthClients struct {
	Id     string
	Secret string
}

func NewOAuthOptions(v *viper.Viper, logger *zap.Logger) (*OAuthOptions, error) {
	var err error

	opt := &OAuthOptions{}
	if err = v.UnmarshalKey("oauth", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal redis config error")
	}

	logger.Info("load oauth server config success")
	return opt, err
}
