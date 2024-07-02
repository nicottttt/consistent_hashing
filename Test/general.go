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

	c.TraverseHashRing()
	//c.TraverseSortedSet()

	c.DelServer("Server1")
	c.DelServer("Server2")

	fmt.Println("After removing Server1 and Server2:")

	c.TraverseHashRing()
	//c.TraverseSortedSet()
}
