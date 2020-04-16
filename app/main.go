package main

import (
	"demo/app/controller"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:20
 * @Title:
 * --- --- ---
 * @Desc:
 */
//var (
//	addr = flag.String("addr", "localhost:4001", "server address")
//	etcdAddr = flag.String("etcdAddr", "175.24.110.208:2380", "etcd address")
//	basePath = flag.String("base", "/rpcx_test", "prefix path")
//)

func main() {
	createApiServer()
	//go createRpcServer()
	//select {}
}

func createApiServer() {
	r := gin.New()
	r.Use(gin.Logger())
	index := controller.CreateIndexController

	v1 := r.Group("/")
	{
		v1.POST("/getUserInfo", index.GetUserBaseInfo)
	}

	r.Run(":4000")
}

//func createRpcServer() {
//	flag.Parse()
//	rpcServer := server.NewServer()
//	addRegistryPlugin(rpcServer)
//	rpcServer.RegisterName("ucenter-remote", &provider.Provider{}, "")
//	rpcServer.Serve("tcp", *addr)
//}
//
//func addRegistryPlugin(s *server.Server) {
//	r := &serverplugin.EtcdV3RegisterPlugin{
//		ServiceAddress: "tcp@" + *addr,
//		EtcdServers: []string{*etcdAddr},
//		BasePath: *basePath,
//		Metrics: metrics.NewRegistry(),
//		UpdateInterval: time.Minute,
//	}
//
//	if err := r.Start(); err != nil {
//		log.Fatal(err)
//	}
//
//	s.Plugins.Add(r)
//}
