package provider

import (
	"context"
	"demo/pkg/transports/rpc"
	"demo/ucenter/provider/param"
	"demo/ucenter/service"
	"github.com/google/wire"
	"github.com/gutrse3321/aki-persit/remote"
	"github.com/smallnest/rpcx/server"
	"log"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 17:50
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Provider struct {
}

func (p *Provider) GetUserBaseInfo(ctx context.Context, args *param.GetUserBaseInfoArgs, codeRemote *remote.Remote) error {
	userBaseInfoDto := service.GetUserBaseInfo()
	remote.Init(codeRemote, userBaseInfoDto)
	return nil
}

func CreateInitRpcServersFn() rpc.InitServers {
	return func(s *server.Server) {
		s.RegisterName("ucenter-provider", &Provider{}, "")
		log.Println("ucenter-provider create success ... ")
	}
}

var ProviderSet = wire.NewSet(CreateInitRpcServersFn)
