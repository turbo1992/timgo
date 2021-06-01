package test

import (
	"fmt"
)

type Client struct {
	PreHeader string
	Ip string
}

func init()  {
	//fmt.Println("init reactive...")
}

func (c *Client) SetHeader (header string) *Client  {
	c.PreHeader = header
	return c
}

func (c *Client) Connect (ip string) *Client  {
	c.Ip = ip
	return c
}

func TestReactive()  {
	c := &Client{}
	c.SetHeader("http://").Connect("192.168.0.1")
	fmt.Println("c:", c.PreHeader, c.Ip)
}
