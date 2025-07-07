package main

import (
	"fmt"
	"micro-memorandum/app/gateway/router"
	"micro-memorandum/app/gateway/rpc"
	"micro-memorandum/config"

	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

// API 网关 或 HTTP 服务
func main() {
	config.Init()
	rpc.InitRPC()
	//etcd 注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	// new 一个微服务实例
	webService := web.NewService(
		web.Name("httpService"),
		web.Address("localhost:4000"),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()), // 来自router文件夹的routers
		web.Metadata(map[string]string{"protocl": "http"}),
	)

	webService.Init()

	_ = webService.Run()
}
