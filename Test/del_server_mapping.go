package main

import (
	"consistent/consistent"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	c := consistent.NewRing(15)

	uuidStr := "d986d6f2-d8bd-11ed-bef7-fbe3cc6ffb9e"
	server, _ := uuid.Parse(uuidStr)
	server1 := server.String()
	// server1 := "Server1"
	c.AddServer(server1)

	uuidStr2 := "caac6d5e-1291-480b-9ce5-c3fa1ad8efad"
	server, _ = uuid.Parse(uuidStr2)
	server2 := server.String()
	// server2 := "Server2"
	c.AddServer(server2)

	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf(uuid.New().String())
		c.AddKey(key)
	}

	fmt.Printf("Len set: %d, ring len:%d\n", len(c.GetSortedset()), len(c.GetRing()))

	// c.TraverseSortedSet()

	c.DelServer(server1)

	fmt.Printf("Len set: %d, ring len:%d\n", len(c.GetSortedset()), len(c.GetRing()))

	// c.DelServer(server1)

	// fmt.Printf("Len set: %d, ring len:%d\n", len(c.GetSortedset()), len(c.GetRing()))
	//c.TraverseMapping()

}
