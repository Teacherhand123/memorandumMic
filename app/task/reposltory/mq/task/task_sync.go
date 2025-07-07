package task

import (
	"context"
	"encoding/json"
	"micro-memorandum/app/task/reposltory/mq"
	"micro-memorandum/app/task/service"
	"micro-memorandum/consts"
	"micro-memorandum/idl/pb"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMqTaskQueue // 队列名
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}

	var forever chan struct{}
	go func() {
		// range 可以自动从 channel 中一条一条读数据。
		// 当发送方（写端）关闭了 channel，for range 会自动退出循环。
		for d := range msgs {
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}
			err = service.TaskMQ2DB(ctx, req)
			if err != nil {
				return
			}

			// 告诉rabbitmq，已完成
			// multiple == false：
			// → 只确认当前这一条消息。
			// multiple == true：
			// → 一次确认当前消息及其之前所有未确认的消息。
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}
