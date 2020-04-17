package database

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 15:34
 * @Title: 数据库配置
 * --- --- ---
 * @Desc:
 */

type Options struct {
	User         string
	Password     string
	Host         string
	Db           string
	MaxIdleConns int
	MaxOpenConns int
	Debug        bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("database", opt); err != nil {
		return nil, errors.Wrap(err, "Unmarshal database config error")
	}

	logger.Info("load database config success", zap.String("host", opt.Host))

	return opt, err
}

/**
初始化连接数据库(mySql)
*/
func New(opt *Options) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Db))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open database connection error")
	}
	if opt.Debug {
		db = db.Debug()
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(opt.MaxIdleConns)
	db.DB().SetMaxOpenConns(opt.MaxOpenConns)

	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
