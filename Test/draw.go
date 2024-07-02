package main

import (
	"consistent/consistent"
	"fmt"
)

func main() {
	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")
	c.AddServer("Server4")
	key := "key222222"
	key_server := c.MapKey(key)
	fmt.Println("Key", key_server)

	consistent.DrawRing(c, key)
}
