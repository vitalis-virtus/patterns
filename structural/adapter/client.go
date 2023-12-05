package adapter

import "fmt"

type Client struct{}

func (c *Client) InsertLigtingConnectorIntoComputer(comp Computer) {
	fmt.Println("Client inserts Lighting connector into computer")
	comp.InsertIntoLightingPort()
}
