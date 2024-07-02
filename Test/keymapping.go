package main

import (
	"consistent"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")
	c.AddServer("Server4")

	cnt_one := 0
	cnt_two := 0
	cnt_three := 0
	cnt_four := 0

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf(uuid.New().String())
		server := c.MapKey(key)
		switch server {
		case "Server1":
			cnt_one++
		case "Server2":
			cnt_two++
		case "Server3":
			cnt_three++
		case "Server4":
			cnt_four++
		}
	}

	fmt.Println("Server1's key:", cnt_one)
	fmt.Println("Server2's key:", cnt_two)
	fmt.Println("Server3's key:", cnt_three)
	fmt.Println("Server4's key:", cnt_four)
}
