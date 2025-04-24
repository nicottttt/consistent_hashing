package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"consistent.com/m/consistent"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var ring *consistent.Consistent

func etcdWatcher(cli *clientv3.Client) {

	rch := cli.Watch(context.Background(), "/nodes/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			addr := string(ev.Kv.Key[len("/nodes/"):])
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Println("Added node:", addr)
				ring.AddServer(addr)
			case clientv3.EventTypeDelete:
				fmt.Println("Removed node:", addr)
				ring.DelServer(addr)
			}
		}
	}
}

func addExistingNode(cli *clientv3.Client) {
	resp, err := cli.Get(context.TODO(), "/nodes/", clientv3.WithPrefix())
	if err != nil {
		log.Println("Failed to list nodes:", err)
		return
	}

	for _, kv := range resp.Kvs {
		addr := string(kv.Key[len("/nodes/"):])
		fmt.Println("Added node:", addr)
		ring.AddServer(addr)
	}
}

func registerNode(cli *clientv3.Client, addr string) {
	lease, _ := cli.Grant(context.TODO(), 10)
	_, err := cli.Put(context.TODO(), "/nodes/"+addr, "", clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatal("Put to etcd failed:", err)
	}

	addExistingNode(cli)

	// 自动续租
	ch, _ := cli.KeepAlive(context.TODO(), lease.ID)
	go func() {
		for range ch {
			// 保活中
		}
	}()
}

func main() {
	ring = consistent.NewRing(3)
	if len(os.Args) != 3 {
		log.Fatal("Usage: server <ip:port> <etcd_port>")
	}
	addr := os.Args[1]

	// 注册服务到 etcd
	port := os.Args[2]
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:" + port},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("Connect to etcd failed:", err)
	}
	defer cli.Close()
	registerNode(cli, addr)

	// Etcd 监听
	go etcdWatcher(cli)

	// UDP 监听
	udpAddr, _ := net.ResolveUDPAddr("udp", addr)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	fmt.Println("Listening UDP at", addr)

	buf := make([]byte, 1024)
	for {
		n, remote, _ := conn.ReadFromUDP(buf)

		// Test consistent hash
		clientid := string(buf[:n])
		mapaddr := ring.MapKey(clientid)
		if mapaddr == addr {
			fmt.Println("Mapped correctly!")
		} else {
			fmt.Println("Wrong mapping!")
		}

		fmt.Printf("From %s: %s\n", remote, string(buf[:n]))
	}
}
