package main

import (
	"consistent"
	"fmt"
)

func main() {
	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		server := c.MapKey(key)
		c.AddKey(key, server)
	}

	c.TraverseMapping()

	fmt.Println("After adding server 4:---------")
	c.AddServer("Server4")

	c.TraverseMapping()

}
