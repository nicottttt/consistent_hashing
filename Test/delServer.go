package main

import (
	"consistent/consistent"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")
	c.AddServer("Server4")

	for i := 0; i < 1000000; i++ {
		c.AddKey(fmt.Sprintf("key%d", i))
	}

	//c.TraverseMapping()

	fmt.Println("After deleting Server1 :---------")
	c.DelServer("Server1")

	//c.TraverseMapping()
	elapsed := time.Since(start)
	fmt.Printf("The code took %s to execute.\n", elapsed)

}
