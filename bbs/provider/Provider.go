package provider

import (
	"context"
	"demo/bbs/provider/param"
	"demo/pkg/transports/rpc"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/persit/remote"
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
type BBSProvider struct {
}

func (p *BBSProvider) GetNotThing(ctx context.Context, args *param.GetNoThingArgs, codeRemote *remote.Remote) error {
	remote.Init(codeRemote, "nothing")
	return nil
}

func CreateInitRpcServersFn() rpc.InitServers {
	return func(s *server.Server) {
		s.Register(&BBSProvider{}, "")
		log.Println("bbs-provider create success ... ")
	}
}

var ProviderSet = wire.NewSet(CreateInitRpcServersFn)
