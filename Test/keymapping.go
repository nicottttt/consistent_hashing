package main

import (
	"consistent/consistent"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	c := consistent.NewRing(25)
	server1, _ := uuid.Parse("0fef7c29-d38b-43bf-b383-eaeefcc21bc5")
	server2, _ := uuid.Parse("5142f8f6-e676-4859-b1c5-03b41912747d")
	server3, _ := uuid.Parse("230a3a8b-f9b7-42f2-a99c-25cb6038dea2")
	server4, _ := uuid.Parse("ae9fb244-e985-4254-bc09-54a3aab47060")
	c.AddServer(server1.String(), 1) // Add server1 with 13 virtual nodes
	c.AddServer(server2.String(), 1)
	c.AddServer(server3.String(), 1)
	c.AddServer(server4.String(), 1)

	cnt_one := 0
	cnt_two := 0
	cnt_three := 0
	cnt_four := 0

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf(uuid.New().String())
		server := c.MapKey(key)
		switch server {
		case server1.String():
			cnt_one++
		case server2.String():
			cnt_two++
		case server3.String():
			cnt_three++
		case server4.String():
			cnt_four++
		}
	}

	fmt.Println("Server1's key:", cnt_one)
	fmt.Println("Server2's key:", cnt_two)
	fmt.Println("Server3's key:", cnt_three)
	fmt.Println("Server4's key:", cnt_four)

}
