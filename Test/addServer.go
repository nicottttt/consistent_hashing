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

	for i := 0; i < 100; i++ {
		hashkey := consistent.Hashkey{
			SrcIP: fmt.Sprintf("key%d", i),
			DstIP: fmt.Sprintf("key%d", i),
		}
		c.AddKey(hashkey)
	}

	//c.TraverseMapping()

	fmt.Println("After adding server 4:---------")
	c.AddServer("Server4", 15)

	c.TraverseMapping()

	elapsed := time.Since(start)
	fmt.Printf("The code took %s to execute.\n", elapsed)

}
