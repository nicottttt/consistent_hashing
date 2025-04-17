package main

import (
	"consistent/consistent"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := consistent.NewRing(15)
	c.AddServer("Server1", 15)
	c.AddServer("Server2", 15)
	c.AddServer("Server3", 15)
	c.AddServer("Server4", 15)

	for i := 0; i < 1000000; i++ {
		hashkey := consistent.Hashkey{
			SrcIP: fmt.Sprintf("key%d", i),
			DstIP: fmt.Sprintf("key%d", i),
		}
		c.AddKey(hashkey)
	}

	c.TraverseMapping()

	fmt.Println("After deleting Server1 :---------")
	c.DelServer("Server5")

	c.TraverseMapping()
	elapsed := time.Since(start)
	fmt.Printf("The code took %s to execute.\n", elapsed)

}
