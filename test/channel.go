package test

import (
	"fmt"
	"time"
)

func WriteChannel()  {
	ch := make(chan int, 100)
	defer close(ch)
	go ReadChannel(ch)
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(i * 100) * time.Millisecond)
		ch <- i
	}
	time.Sleep(200 * time.Millisecond)
}

func ReadChannel(ch chan int)  {
	for {
		select {
		case val := <-ch:
			fmt.Printf("channel:%d\n", val)
		}
	}
}

func TestChannel()  {
	ch := make(chan int, 100)
	go func(chan int) {
		for {
			val := <-ch
			fmt.Printf("channel:%d\n", val)
		}
	}(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func TestChannel2()  {
	ch := make(chan int, 20)
	go goChanA(ch)
	// 往ch中写入数据值
	time.Sleep(1 * time.Second)
}

func goChanA(a <-chan int) {
	b := <-a
	fmt.Println("只能接收数据的channal[a]接收到的数据值为", b)
}

func TestChannel3()  {
	ch := make(chan string)

	for i := 0; i < 100; i++ {
		go func() {
			ch <- "写入"
		}()
		time.Sleep(time.Second * 1)
	}

	for {
		select {
		case res := <-ch:
			fmt.Println("res", res)

		}
	}

	go func() {
		
	}()

	time.Sleep(100*time.Second)
	fmt.Println("finished")
}

type Book struct {
	title string
	page int
	class string
}

func TestChannel4() {
	msgChan := make(chan int, 20)
	go func() {
		for i:=0; i<10; i++ {
			msgChan <- 1000
			time.Sleep(time.Second)
		}
		fmt.Println("10次消息结束")
	}()
	readMsg1(msgChan)
	//go readMsg2()
	time.Sleep(time.Second * 1000)
}

func readMsg1(msgChan chan int) {
	for {
		select {
		case msg1 := <- msgChan:
			fmt.Println("协程1收到消息", msg1)
		}
	}
}

func readMsg2(msgChan chan int) {
	msg2 := <- msgChan
	fmt.Println("协程2收到消息", msg2)
}

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func Test5()  {
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			goto end
		default:
			fmt.Println("no commucation")
		}
	}
end:
}