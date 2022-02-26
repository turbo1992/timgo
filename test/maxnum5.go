package test

import "sync"

var ch chan struct{}

func job1() {
	println("job1")
	ch <- struct{}{}
	ch <- struct{}{}
}

func job2()  {
	<-ch
	println("job2")
}

func job3() {
	<-ch
	println("job3")
}

func do(fs ... func()) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	//for v := range fs {
	//	go func() {
	//		v()
	//	}()
	//	println(v)
	//}
	return &wg
}

func TestMaxNum4() {
	ch = make(chan struct{}, 2)
	wg := do(job1, job2, job3)
	wg.Wait()
}




