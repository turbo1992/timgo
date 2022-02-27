package rabbitmq

import (
	"fmt"
)

type TestRabbitMQStruct struct {

}

func (*TestRabbitMQStruct)TestRabbitMQ()  {
	fmt.Println("开始使用RabbitMQ...")
	TestRabbitMQ1()

	ch := make(chan struct{})
	<-ch
}

func TestRabbitMQ1()  {
	// 生产消息
	go productMessage()
	// 消费消息
	go consumeMessage()
}
