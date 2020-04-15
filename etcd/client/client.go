package main

import (
	"context"
	"flag"
	"github.com/gutrse3321/aki-persit/remote"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/14 12:10
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	etcdAddr = flag.String("etcdAddr", "175.24.110.208:2380", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

type Args struct {
	A int
	B int
}

func main() {
	flag.Parse()

	go createRemoteClient("air1", "%d * %d = %d")
	go createRemoteClient("air2", "%d + %d = %d")

	select {}
}

func createRemoteClient(serviceName string, logContent string) {
	d := client.NewEtcdV3Discovery(*basePath, serviceName, []string{*etcdAddr}, nil)
	opt := client.DefaultOption
	opt.Heartbeat = true
	opt.HeartbeatInterval = time.Second * 10

	xclient := client.NewXClient(serviceName, client.Failbackup, client.WeightedRoundRobin, d, opt)

	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &remote.Remote{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		res, err := remote.ResolveRemote(*reply)
		if err != nil {
			log.Fatal("remote", err)
		}

		log.Printf(logContent, args.A, args.B, res.(int64))

		time.Sleep(5 * time.Second)
	}
}
