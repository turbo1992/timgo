package test

import (
	"math/rand"
	"sync"
	"time"
)

func my_job()  {
	println(rand.Intn(100))
}

func TestRand()  {
	maxnum := make(chan struct{}, 5)
	go func() {
		for {
			for i:=0; i<5; i++ {
				maxnum <- struct{}{}
			}
			time.Sleep(time.Second*2)
			println("发放了5个任务")
		}
	}()

	for {
		<- maxnum
		go my_job()
	}
}

var wg = sync.WaitGroup{}

func pub_job()  {
	defer wg.Done()
	time.Sleep(time.Second*1)
}

func priv_job() {
	defer  wg.Done()
	time.Sleep(time.Second*2)
}

func TestAsyncJob()  {
	start := time.Now()
	wg.Add(2)
	go func() {
		pub_job()
	}()
	go func() {
		priv_job()
	}()
	wg.Wait()
	println("时间%d", time.Since(start).Milliseconds()/1000)
}