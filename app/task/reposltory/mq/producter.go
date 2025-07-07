package mq

import (
	"fmt"
	"micro-memorandum/consts"

	"github.com/streadway/amqp"
)

func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	// 队列名称。指定要声明的队列名。
	// true durable（持久化）。为 true 时，队列会持久化到磁盘，RabbitMQ 重启后队列依然存在。
	// false autoDelete（自动删除）。为 true 时，如果没有消费者连接，队列会自动删除。这里为 false，表示不会自动删除。
	// false exclusive（排他队列）。为 true 时，只能被当前连接访问，连接断开后队列自动删除。这里为 false，表示不是排他队列。
	// false oWait（不等待）。为 true 时，不等待服务器响应，直接返回。这里为 false，表示要等待服务器确认队列已声明。
	// nil rgs（额外参数）。可以设置一些扩展参数，这里为 nil，表示没有额外参数。
	q, _ := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})

	if err != nil {
		return
	}

	fmt.Println("发送MQ成功")
	return
}
