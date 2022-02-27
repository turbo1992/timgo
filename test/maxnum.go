package test

import (
	"fmt"
	"time"
)

func do_job(ch chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}

func TestMaxNum() {
	taskPool := make(chan string, 5)
	go func() {
		for {
			for i:=0; i<5; i++ {
				// 发放任务
				taskPool <- fmt.Sprintf("task编号---------%d", i)
			}
			time.Sleep(time.Second*3)
			fmt.Println("发放了5个任务")
		}
	}()

	for {
		go do_job(taskPool)
	}
}

func TestMyChan()  {
	pool := make(chan string, 5)
	go read(pool)
	for {
		for i:=0; i<5; i++ {
			pool <- fmt.Sprintf("index:%d", i)
			time.Sleep(time.Second)
		}
	}
}

func read(pool chan string)  {
	for v := range pool {
		fmt.Println(v)
	}
}
