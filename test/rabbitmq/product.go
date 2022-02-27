package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func productMessage() {
	uri := "amqp://guest:guest@localhost:5672/"
	exchange := "project"
	queue := "pj_event"
	routing_key := "pj_event"

	for i:=0; i<10; i++ {
		content := map[string]interface{}{
			"name": fmt.Sprintf("turbo-%d", i),
		}
		err := Pub_mq(uri, exchange, queue, routing_key, content)
		if err != nil {
			fmt.Println("Publish Msg Error:", err)
		}
		time.Sleep(time.Second)
	}
}

// 生产者
func Pub_mq(uri, exchange, queue, routing_key string, content map[string]interface{}) error {
	// 建立连接
	connection, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer connection.Close()
	// 创建一个Channel
	channel, err := connection.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}
	defer channel.Close()

	// 声明exchange
	if err := channel.ExchangeDeclare(
		 exchange, //name
		"direct", //exchangeType
		true,     //durable
		false,    //auto-deleted
		false,    //internal
		false,    //noWait
		nil,      //arguments
	); err != nil {
		log.Println("Failed to declare a exchange:", err.Error())
		return err
	}
	// 声明一个queue
	if _, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	); err != nil {
		log.Println("Failed to declare a queue:", err.Error())
		return err
	}
	// exchange 绑定 queue
	channel.QueueBind(queue, routing_key, exchange, false, nil)

	// 发送
	messageBody := MapToJson(content)
	if err = channel.Publish(
		exchange,    // exchange
		routing_key, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(messageBody),
			//Expiration:      "60000", // 消息过期时间
		},
	); err != nil {
		log.Println("Failed to publish a message:", err.Error())
		return err
	}
	return nil
}

func MapToJson(param map[string]interface{}) string{
	dataType , _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}