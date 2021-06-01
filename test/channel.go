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
	var res string
	ch := make(chan string)
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second * 1)
		ch <- "写入"
	}()

	go func() {
		res = <-ch
		fmt.Println("res", res)
	}()

	fmt.Println("finished")
}