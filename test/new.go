package test

import "fmt"

type Student struct {
	Name string
	Class string
}

func main()  {
	fmt.Println("main...")

}

func init()  {
	//fmt.Println("init new...")
}

func TestNew()  {
	//var s = new(Student)
	var s = &Student{}
	s.Name = "golang"
	s.Class = "c++"
	//s = &Student{
	//	Name: "1",
	//	Class: "2",
	//}
	fmt.Println("s", s)
	main()
}