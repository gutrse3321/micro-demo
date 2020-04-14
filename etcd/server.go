package main

import (
	"context"
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/14 12:03
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	addr = flag.String("addr", "localhost:8972", "server address")
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

type Arith struct {
}

func (a *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.Register(new(Arith), "")

	s.Serve("tcp", *addr)
}

/**
	etcd 插件配置
 */
func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers: []string{*etcdAddr},
		BasePath: *basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	if err := r.Start(); err != nil {
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}
