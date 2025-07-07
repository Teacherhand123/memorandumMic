package main

import (
	"context"
	"fmt"
	"micro-memorandum/app/task/reposltory/db/dao"
	"micro-memorandum/app/task/reposltory/mq"
	"micro-memorandum/app/task/script"
	"micro-memorandum/app/task/service"
	"micro-memorandum/config"
	"micro-memorandum/idl/pb"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

// RPC 服务
func main() {
	config.Init()
	dao.InitDB()
	mq.InitRabbitMQ()
	loadingSrcipt()

	//etcd 注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	// new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.TaskServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	// RegisterTaskServiceHandler内的函数写完才能注册成功
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	if err := microService.Run(); err != nil {
		panic(err)
	}
}

func loadingSrcipt() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
