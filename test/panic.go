package test

import "fmt"

func TestPanic()  {
	// 先defer的后执行，recover后输出panic中的信息
	defer func(){
		if err := recover(); err != nil{
			fmt.Print(err)
		} else {
			fmt.Print("no")
		}
	}()
	//defer func(){
	//	panic("11111111111")
	//}()
	println("qqqqqqqq")
	panic("22222222222")

}

func TestDefer()  {
	defer func() {
		println("a")
	}()
	defer func() {
		println("b")
	}()
	defer func() {
		println("c")
	}()
	println("d")
}
