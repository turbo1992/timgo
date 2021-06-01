package test

import (
	"fmt"
	"time"
)

func TestMutex()  {

	var count = 0
	go func() {
		for i := 0; i < 1000; i++ {
			count ++
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			count ++
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("count = ", count)

}