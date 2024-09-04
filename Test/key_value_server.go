package main

import (
	"consistent/consistent"
	"fmt"

	"github.com/cespare/xxhash"
)

func main() {
	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")
	c.AddServer("Server4")

	key1 := "key2222"
	key2 := "key222222"
	key3 := "key22"
	key4 := "key2"

	fmt.Printf("Key|%s|'s value is %d, and it is mapped to:%s\n", key1,
		xxhash.Sum64([]byte(key1))%1024, c.MapKey(key1))

	fmt.Printf("Key|%s|'s value is %d, and it is mapped to:%s\n", key2,
		xxhash.Sum64([]byte(key2))%1024, c.MapKey(key2))

	fmt.Printf("Key|%s|'s value is %d, and it is mapped to:%s\n", key3,
		xxhash.Sum64([]byte(key3))%1024, c.MapKey(key3))

	fmt.Printf("Key|%s|'s value is %d, and it is mapped to:%s\n", key4,
		xxhash.Sum64([]byte(key4))%1024, c.MapKey(key4))
}
