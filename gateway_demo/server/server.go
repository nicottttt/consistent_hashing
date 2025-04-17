package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run server.go [port]")
		return
	}
	port := os.Args[1]
	portNum, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	addr := net.UDPAddr{
		Port: portNum,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on port %d\n", portNum)

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		fmt.Printf("Connection establish\n")

		msg := string(buf[:n])
		fmt.Printf("Received from %s: %s\n", remoteAddr, msg)
	}
}
