package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 15:55
 * @Title: logger zap配置
 * --- --- ---
 * @Desc:
 */

type Options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error

	opt := &Options{}
	if err = v.UnmarshalKey("log", opt); err != nil {
		return nil, err
	}
	return opt, err
}

func New(opt *Options) (*zap.Logger, error) {
	var (
		err    error
		logger *zap.Logger
	)

	level := zap.NewAtomicLevel()

	err = level.UnmarshalText([]byte(opt.Level))
	if err != nil {
		return nil, err
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxBackups,
		MaxBackups: opt.MaxAge,
	})

	cw := zapcore.Lock(os.Stdout)

	//file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	//stdout core 采用consoleEncoder
	if opt.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)
	zap.ReplaceGlobals(logger)

	return logger, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
