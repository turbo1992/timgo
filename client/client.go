package client

type Client struct {
	PreHeader string
	Ip string
}

func (c *Client) SetHeader (header string) *Client  {
	c.PreHeader = header
	return c
}

func (c *Client) Connect (ip string) *Client  {
	c.Ip = ip
	return c
}