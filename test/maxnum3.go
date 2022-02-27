package test

import (
	"fmt"
	"math/rand"
	"time"
)

func myjob() {
	fmt.Println(rand.Int())
}

func setTaskPool(taskPool chan string) {
	for i:=0; i<5; i++ {
		// 发放任务
		taskPool <- fmt.Sprintf("task编号---------%d", i)
	}
}

func TestMaxNum3() {
	taskPool := make(chan string, 5)
	setTaskPool(taskPool)
	go func() {
		for {

			time.Sleep(time.Second*3)
			fmt.Println("发放了5个任务")
		}
	}()

	for {
		<-taskPool
		go myjob()
	}
}