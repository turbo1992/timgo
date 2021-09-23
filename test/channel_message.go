package test

import (
	"fmt"
	"time"
)

type MessageHandler struct {
	Message chan string
}

func (handler *MessageHandler) Send(message string) {
	handler.Message <- message
}

func (handler *MessageHandler) Listen() {
	go func() {
		for {
			select {
			case message := <-handler.Message:
				fmt.Println("new message: ", message)
			}
		}
	}()
}

func MySend(ch chan string, message string)  {
	ch <- message
}

func MyListen(ch chan string) {
	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println("received:", msg)
			}
		}
	}()
}

func TestChannelMessage()  {
	handler := &MessageHandler{Message: make(chan string)}

	// 监听消息
	handler.Listen()

	handler.Send("1")

	time.Sleep(time.Second)

	handler.Send("2")

	time.Sleep(time.Second)

	handler.Send("3")

	time.Sleep(time.Second)

	handler.Send("4")

	time.Sleep(time.Second)

	ch := make(chan string)
	MyListen(ch)
	MySend(ch, "process end")
}