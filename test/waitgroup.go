package test

import (
	"sync"
	"time"
)

func TestWaitGroup()  {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		time.Sleep(1*time.Second)
		wg.Done()
		println("111111111")
	}()

	go func() {
		time.Sleep(2*time.Second)
		wg.Done()
		println("222222222")
	}()

	go func() {
		time.Sleep(10*time.Second)
		wg.Done()
		println("333333333")
	}()

	wg.Wait()

	println("44444444")
}