package main

import (
	"fmt"
	"micro-memorandum/app/user/reposltory/db/dao"
	"micro-memorandum/app/user/service"
	"micro-memorandum/config"
	"micro-memorandum/idl/pb"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

// RPC 服务
func main() {
	config.Init()
	dao.InitDB()

	//etcd 注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	// new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	if err := microService.Run(); err != nil {
		panic(err)
	}
}
