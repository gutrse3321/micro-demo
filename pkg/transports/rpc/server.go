package rpc

import (
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	rpcxServer "github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 11:49
 * @Title:
 * --- --- ---
 * @Desc:
 */
/**
从配置文件中获取 key: rpc
*/
type ServerOptions struct {
	BasePath       string
	UpdateInterval time.Duration
	EtcdAddress    []string
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	opt := &ServerOptions{}

	if err := v.UnmarshalKey("rpc", opt); err != nil {
		return nil, err
	}
	return opt, nil
}

type Server struct {
	opt      *ServerOptions
	app      string
	host     string
	port     int
	logger   *zap.Logger
	server   *rpcxServer.Server
	initFunc InitServers
}

type InitServers func(s *rpcxServer.Server)

/**
注入依赖项，返回实例
*/
func NewServer(opt *ServerOptions, logger *zap.Logger, init InitServers) (*Server, error) {
	var s *rpcxServer.Server

	logger = logger.With(zap.String("type", "rpcx"))
	{
		s = rpcxServer.NewServer()
	}

	return &Server{
		opt:      opt,
		logger:   logger,
		server:   s,
		initFunc: init,
	}, nil
}

func (s *Server) ApplicationName(name string) {
	s.app = name
}

func (s *Server) Start(ln net.Listener) error {
	s.logger.Info("rpc server starting ...")

	if err := s.register(); err != nil {
		return errors.Wrap(err, "register rpc server error")
	}

	//初始化rpc服务器的provider
	s.initFunc(s.server)

	go func() {
		if err := s.server.ServeListener("tcp", ln); err != nil {
			s.logger.Fatal("failed to serve rpc: %v", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) register() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		EtcdServers:    s.opt.EtcdAddress,
		BasePath:       s.opt.BasePath,
		UpdateInterval: s.opt.UpdateInterval * time.Millisecond,
	}

	if err := r.Start(); err != nil {
		return errors.Wrap(err, "register center error")
	}
	s.server.Plugins.Add(r)

	s.logger.Info("register rpc service success", zap.String("id", "tcp@"+addr))
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("rpc server stopping ... ")
	if err := s.server.Close(); err != nil {
		return errors.Wrap(err, "stop rpc server error")
	}
	return nil
}

var ServerProviderSet = wire.NewSet(NewServerOptions, NewServer)
