package ucenterRemote

import (
	"context"
	"demo/ucenter/provider/param"
	"flag"
	"github.com/gutrse3321/aki-persit/remote"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:06
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	EtcdAddr = flag.String("etcdAddr", "175.24.110.208:2380", "etcd address")
	BasePath = flag.String("base", "/rpcx_test", "prefix path")
)

func GetUserBaseInfoRemote(args param.GetUserBaseInfoArgs) interface{} {
	xClient := createUcenterRemote()
	defer xClient.Close()

	codeRemote := &remote.Remote{}
	err := xClient.Call(context.Background(), "GetUserBaseInfo", args, codeRemote)
	if err != nil {
		log.Printf("GetUserBaseInfo failed to call: %v\n", err)
	}

	res, err := remote.ResolveRemote(*codeRemote)
	if err != nil {
		log.Println(err)
	}

	return res.(interface{})
}

func createUcenterRemote() client.XClient {
	d := client.NewEtcdV3Discovery(*BasePath, "ucenter-remote", []string{*EtcdAddr}, nil)
	opt := client.DefaultOption
	opt.Heartbeat = true
	opt.HeartbeatInterval = time.Second * 10

	xclient := client.NewXClient("ucenter-remote", client.Failbackup, client.WeightedRoundRobin, d, opt)

	return xclient
}
