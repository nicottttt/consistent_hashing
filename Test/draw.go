package main

import (
	"consistent/consistent"

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

	// key1 := "key222222"
	// key_server1 := c.MapKey(key1)
	// fmt.Println("Key1", key_server1)

	// key2 := "key2222"
	// key_server2 := c.MapKey(key2)
	// fmt.Println("Key2", key_server2)

	// key3 := "key222"
	// key_server3 := c.MapKey(key3)
	// fmt.Println("Key3", key_server3)

	consistent.DrawRing(c)
}
