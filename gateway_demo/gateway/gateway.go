package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"consistent.com/m/consistent"
)

const (
	OperationAdd int = iota
	OperationRemove
	OperationSend
)

type Operation struct {
	Type int
	Node string
}

type Send struct {
	Id string
}

type ConsistentRouter struct {
	ring        *consistent.Consistent
	opChan      chan Operation
	sendChan    chan Send
	stopChan    chan struct{}
	workerAlive bool
}

var router *consistent.Consistent

func InitConsistentRouter(replicationFactor int) *ConsistentRouter {
	ring := consistent.NewRing(replicationFactor)
	return &ConsistentRouter{
		ring:        ring,
		opChan:      make(chan Operation, 100),
		sendChan:    make(chan Send, 100),
		stopChan:    make(chan struct{}),
		workerAlive: false,
	}
}

func main() {
	cr := InitConsistentRouter(3)
	go cr.run()
	router = cr.ring

	// Send
	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		if len(cr.ring.GetServerList()) == 0 {
			fmt.Println("[Router] No nodes available for sending")
			fmt.Fprintf(w, "No server exist\n")
		} else {
			cr.Send(id)
			fmt.Fprintf(w, "Node %s queued for sending\n", id)
		}

	})

	// ADD
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		node := r.URL.Query().Get("node")
		cr.AddNode(node)
		fmt.Fprintf(w, "Node %s queued for addition\n", node)
	})

	// REMOVE
	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		node := r.URL.Query().Get("node")
		cr.RemoveNode(node)
		fmt.Fprintf(w, "Node %s queued for removal\n", node)
	})
	fmt.Println("Gateway listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendUDP(addr string, msg string) error {
	fmt.Println("Preparing to send UDP to", addr, "with msg:", msg)

	remoteAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("resolve error:", err)
		return err
	}

	localAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 0, // 系统分配端口
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("dial error:", err)
		return err
	}
	defer conn.Close()

	n, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("write error:", err)
		return err
	}

	fmt.Printf("Sent %d bytes to %s\n", n, addr)
	return nil
}

func (cr *ConsistentRouter) run() {
	for {
		select {
		case op := <-cr.opChan:
			switch op.Type {
			case OperationAdd:
				cr.ring.AddServer(op.Node)
				fmt.Println("[Router] Added node:", op.Node)
			case OperationRemove:
				cr.ring.DelServer(op.Node)
				fmt.Println("[Router] Removed node:", op.Node)
			}
		case sendId := <-cr.sendChan:
			id_hashkey := consistent.Hashkey{Id: sendId.Id}
			var target string
			if _, ok := router.GetMapping()[id_hashkey]; !ok {
				fmt.Println("[Router] Adding route to router:", id_hashkey)
				router.AddKey(id_hashkey)
			}
			target = router.GetMapping()[id_hashkey]
			_ = sendUDP(target, "hello")

		case <-cr.stopChan:
			fmt.Println("[Router] Stopping router worker...")
			return
		}
	}
}

func (cr *ConsistentRouter) AddNode(node string) {
	cr.opChan <- Operation{Type: OperationAdd, Node: node}
}

func (cr *ConsistentRouter) RemoveNode(node string) {
	cr.opChan <- Operation{Type: OperationRemove, Node: node}
}

func (cr *ConsistentRouter) Send(id string) {
	cr.sendChan <- Send{Id: id}
}
