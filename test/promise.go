package test

import (
	"fmt"
	"sync"
	"time"
)

type Resolve func(resp interface{})
type Reject func(err error)
type PromiseFunc func(resolve Resolve, reject Reject)
type Promise struct {
	function PromiseFunc
	resolve Resolve
	reject Reject
	wg sync.WaitGroup
}

func NewPromise(function PromiseFunc) *Promise  {
	return &Promise{function: function}
}

func (this *Promise) Then(resolve Resolve) *Promise {
	this.resolve = resolve
	return this
}

func (this *Promise) Catch(reject Reject) *Promise {
	this.reject = reject
	return this
}

func (this *Promise) Done() {
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()
		this.function(this.resolve, this.reject)
	}()
	this.wg.Wait()
}
func TestPromise()  {
	NewPromise(func(resolve Resolve, reject Reject) {
		time.Sleep(time.Second)
		if time.Now().Unix()%2 == 0 {
			resolve("OK")
		} else {
			reject(fmt.Errorf("my error"))
		}
	}).Then(func(resp interface{}) {
		fmt.Println(resp)
	}).Catch(func(err error) {
		fmt.Println(err)
	}).Done()
}