package controller

import (
	"fmt"
	"tim-go/client"
)

func Test()  {
	c := &client.Client{}
	c.SetHeader("http://").Connect("192.168.0.1")
	fmt.Println("c:", c.PreHeader, c.Ip)
}

