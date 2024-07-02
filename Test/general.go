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

	c.TraverseHashRing()
	//c.TraverseSortedSet()

	c.RemoveServer("Server1")
	c.RemoveServer("Server2")

	fmt.Println("After removing Server1 and Server2:")

	c.TraverseHashRing()
	//c.TraverseSortedSet()
}
