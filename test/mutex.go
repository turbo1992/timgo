package test

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func TestMutex()  {

	var count = 0
	go func() {
		for i := 0; i < 10000; i++ {
			mutex.Lock()
			count ++
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mutex.Lock()
			count ++
			mutex.Unlock()
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("count = ", count)

}