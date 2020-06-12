package provider

import (
	"context"
	"demo/ucenter/service"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/persit/remote"
	rpc "github.com/gutrse3321/aki/pkg/transports/rpc"
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
type UCenterProvider struct {
}

func (p *UCenterProvider) GetUserBaseInfo(ctx context.Context, args *string, codeRemote *remote.Remote) error {
	userBaseInfoDto := service.GetUserBaseInfo(args)
	remote.Init(codeRemote, userBaseInfoDto)
	return nil
}

func CreateInitRpcServersFn() rpc.InitServers {
	return func(s *server.Server) {
		s.Register(&UCenterProvider{}, "")
		log.Println("ucenter-provider create success ... ")
	}
}

var ProviderSet = wire.NewSet(CreateInitRpcServersFn)
