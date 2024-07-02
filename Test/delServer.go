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
	c.AddServer("Server4")

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		server := c.MapKey(key)
		c.AddKey(key, server)
	}

	c.TraverseMapping()

	fmt.Println("After deleting Server1 and Server2:---------")
	c.DelServer("Server1")
	c.DelServer("Server2")

	c.TraverseMapping()

}
