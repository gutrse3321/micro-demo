package main

import (
	"demo/ucenter/provider"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 17:43
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	addr = flag.String("addr", "localhost:3001", "server address")
	etcdAddr = flag.String("etcdAddr", "175.24.110.208:2380", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	go createApiServer()
	go createRpcServer()
	select {}
	//v1 := r.Group("/")
	//{
	//	//v1.POST("/getUserInfo")
	//}
}

func createApiServer() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Run(":3000")
}

func createRpcServer() {
	flag.Parse()
	rpcServer := server.NewServer()
	addRegistryPlugin(rpcServer)
	rpcServer.RegisterName("ucenter-remote", &provider.Provider{}, "")
	rpcServer.Serve("tcp", *addr)
}

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
