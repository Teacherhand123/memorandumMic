# 备忘录

## 项目简介
这是一个基于go的微服务项目，包含RabbitMQ消息队列，Gin框架，MySQL等数据库等常用组件

## 目录结构

```text
- app
  - gateway
    - cmd: Gin 网关服务入口
    - http: 连接后端 RPC
    - middleware: 中间件
    - router: 路由
    - rpc: 初始化 RPC 客户端
    - warppers: 暂未使用
  - task
    - cmd: task 服务入口
    - repository
      - db
        - dao: 数据库 CURD
        - model: 表模型
      - mq
        - task: 消费者监听
        - InitRabbitMQ: 初始化 RabbitMQ
      - script
        - rabbitMq: 脚本长时间监听
      - service: 业务逻辑
  - user
    - cmd: user 服务入口
    - repository
      - db
        - dao: 数据库 CURD
        - model: 表模型
      - service: 业务逻辑
- config: 配置文件
- consts: 全局常量
- idl
  - pb: Proto 生成代码
  - proto: Proto 定义
- pkg
  - ctl
    - user_info: JWT 与 Context
  - e: 错误码
  - jwt: JWT 鉴权
- types: JSON 定义
- docker-compose.yml: 一键启动
```

```go
## 依赖库
go-micro.dev/v4 v4.11.0
google.golang.org/protobuf v1.36.6
gopkg.in/ini.v1 v1.67.0
gorm.io/driver/mysql v1.6.0
gorm.io/gorm v1.30.0
```

### 如何通过proto生成代码
```cmd
protoc -I ./idl 文件名.proto --micro_out ./idl/pb --go_out=./idl/pb
protoc-go-inject-tag -input ./idl/pb/文件名.pb.go # 将注解写入 文件名.pb.go文件
```

### 为什么要保证Srv只创建一次？
TaskSrv 有可能放到 HTTP 接口里也用、或者测试也用、或者脚本也用，多个 goroutine 并发去拿 GetTaskSrv()。

### 如何启动一个 rpcService 服务，然后把它注册到 etcd，别的服务就能通过 etcd 找到它并调用它暴露的接口
ex:
```go
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
```

### 如何初始化好所有要远程调用的 RPC 客户端代理，帮你自动连接到已注册的远程服务。
前端 (浏览器, Postman) --> API 网关 (HTTP) --> RPC 客户端 (InitRPC) --> 后端 RPC 服务

ex: 
```go
package rpc

import (
	"micro-memorandum/idl/pb"

	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())

	taskMicroService := micro.NewService(micro.Name("taskService.client"))
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())

	UserService = userService
	TaskService = taskService
}
```