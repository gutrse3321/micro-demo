package rpc

import (
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	rpcxServer "github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	Ip             string
	Port           int
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
	provider struct{}
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
		init(s)
	}

	return &Server{
		opt:    opt,
		logger: logger,
		server: s,
	}, nil
}

func (s *Server) ApplicationName(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.opt.Port
	if s.port == 0 {
		return errors.New("missing port: 0")
	}

	s.host = s.opt.Ip
	if s.host == "" {
		return errors.New("missing server ip: \"\"")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.logger.Info("rpc server starting ...", zap.String("addr", addr))

	if err := s.register(); err != nil {
		return errors.Wrap(err, "register rpc server error")
	}

	go func() {
		if err := s.server.Serve("tcp", addr); err != nil {
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
