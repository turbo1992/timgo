package test

import (
	"fmt"
	"sync"
)

func GoFunc()  {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(message string) {
		defer wg.Done()
		fmt.Println("message", message)
	}("123456")
	wg.Wait()
}