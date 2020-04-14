package main

import (
	"context"
	"flag"
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

type Reply struct {
	C int
}

func main() {
	flag.Parse()

	d := client.NewEtcdV3Discovery(*basePath, "Arith", []string{*etcdAddr}, nil)
	opt := client.DefaultOption
	opt.Heartbeat = true
	opt.HeartbeatInterval = time.Second * 10

	xclient := client.NewXClient("Arith", client.Failbackup, client.WeightedRoundRobin, d, opt)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(5 * time.Second)
	}
}
