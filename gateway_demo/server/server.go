package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: server <ip:port>")
	}
	addr := os.Args[1]

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("Connect to etcd failed:", err)
	}
	defer cli.Close()

	// 注册服务到 etcd
	lease, _ := cli.Grant(context.TODO(), 10)
	_, err = cli.Put(context.TODO(), "/nodes/"+addr, "", clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatal("Put to etcd failed:", err)
	}

	// 自动续租
	ch, _ := cli.KeepAlive(context.TODO(), lease.ID)
	go func() {
		for range ch {
			// 保活中
		}
	}()

	// UDP 监听
	udpAddr, _ := net.ResolveUDPAddr("udp", addr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	fmt.Println("Listening UDP at", addr)

	buf := make([]byte, 1024)
	for {
		n, remote, _ := conn.ReadFromUDP(buf)
		fmt.Printf("From %s: %s\n", remote, string(buf[:n]))
	}
}
