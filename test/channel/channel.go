package channel

import (
	"fmt"
	"strconv"
	"time"
)

type TestChannelStruct struct {

}

func (*TestChannelStruct)TestChannel()  {
	//TestChannel1()
	//TestChannel2()
	TestChannel3()

	c := make(chan struct{})
	<-c
}

func TestChannel1()  {
	var ch chan int
	ch = make(chan int)

	go getChVal(ch)

	ch <- 10
	//val := <- ch
	//fmt.Println("ch val is: ", val)
}

func getChVal(ch chan int)  {
	x := <- ch
	fmt.Println("ch val is ", x)
}

type FabricClient struct {
	clientCh chan string
	address string
	owner string
	close chan bool
}

func chainsqlJob(c FabricClient) {
	for i:=0; i<10; i++ {
		c.clientCh <- fmt.Sprintf("hash-%s", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	c.close <- true
}

func logChainsqlJobHash(c FabricClient) {
	for {
		select {
		case hashVal := <- c.clientCh:
			fmt.Println("hash is ", hashVal)
		case isClose := <- c.close:
			if isClose {
				fmt.Println("任务完成，3s后关闭管道")
				goto end
			}
		}
	}
end:
}

func TestChannel2()  {
	client := FabricClient{
		clientCh: make(chan string, 10),
		address: "8080",
		owner: "null",
		close: make(chan bool),
	}

	// 开启协程上链并写入hash
	go func(c FabricClient) {
		chainsqlJob(c)
	}(client)

	// 从channel中读取上链hash
	logChainsqlJobHash(client)
}


func TestChannel3() {
	ch := make(chan int, 1)
	for i:=0; i<10; i++ {
		select {
		case x := <- ch:
			fmt.Println("val is ", x)
		case ch <- i:
			fmt.Println("insert ", i)
		}
	}
}