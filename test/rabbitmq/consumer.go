package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func consumeMessage()  {
	uri := "amqp://guest:guest@localhost:5672/"
	exchange := "project"
	queue := "pj_event"

	err := Use_mq(uri, exchange, queue)
	if err != nil {
		fmt.Println("Consumer Msg Error:", err)
	}
}

// 消费者
func Use_mq(uri, exchange, queue string) error {
	// 建立连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer conn.Close()
	// 启动一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}

	// 声明一个队列
	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue:", err.Error())
		return err
	}
	// 注册消费者
	msgs, err := ch.Consume(
		q.Name,    // queue
		exchange, // 标签
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:", err.Error())
		return err
	}

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			log.Println(d.Type)
			log.Println(d.MessageId)
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
