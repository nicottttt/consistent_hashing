package main

import (
	"fmt"
	"net"
	"os"

	"go.fd.io/govpp"
	"go.fd.io/govpp/api"

	"mygit.com/consistent_hashing/vpp_demo/vppbinapi/fib_types"
	"mygit.com/consistent_hashing/vpp_demo/vppbinapi/ip_session_redirect"
	"mygit.com/consistent_hashing/vpp_demo/vppbinapi/ip_types"
)

func addSourceBasedRouting(ch api.Channel, srcIP, nextHop string, tableIndex uint32) error {
	// IP address validation
	srcAddr := net.ParseIP(srcIP)
	if srcAddr == nil {
		return fmt.Errorf("invalid source IP address: %s", srcIP)
	}
	nhAddr := net.ParseIP(nextHop)
	if nhAddr == nil {
		return fmt.Errorf("invalid next-hop IP address: %s", nextHop)
	}

	match := make([]byte, 4)
	copy(match, srcAddr.To4())

	var nhAddrUnion ip_types.AddressUnion
	copy(nhAddrUnion.XXX_UnionData[:4], nhAddr.To4())

	fibPath := fib_types.FibPath{
		SwIfIndex:  2,          // Interface index (example value)
		TableID:    ^uint32(0), // Default FIB table
		RpfID:      0,
		Weight:     1,
		Preference: 0,
		Type:       fib_types.FIB_API_PATH_TYPE_NORMAL,
		Flags:      fib_types.FIB_API_PATH_FLAG_NONE,
		Proto:      fib_types.FIB_API_PATH_NH_PROTO_IP4,
		Nh: fib_types.FibPathNh{
			Address: nhAddrUnion,
		},
	}
	request := &ip_session_redirect.IPSessionRedirectAddV2{
		TableIndex:  tableIndex,
		OpaqueIndex: 0, // Default value
		Proto:       fib_types.FIB_API_PATH_NH_PROTO_IP4,
		IsPunt:      false, // Forward traffic, not punt
		MatchLen:    4,     // Length of IP address in bytes
		Match:       match,
		NPaths:      1,
		Paths:       []fib_types.FibPath{fibPath},
	}
	response := &ip_session_redirect.IPSessionRedirectAddV2Reply{}

	if err := ch.SendRequest(request).ReceiveReply(response); err != nil {
		return fmt.Errorf("VPP request error: %w", err)
	}
	if response.Retval != 0 {
		return fmt.Errorf("VPP API returned error: %d", response.Retval)
	}

	return nil
}

func main() {
	// Connect to VPP
	conn, err := govpp.Connect("/run/vpp/api.sock")
	if err != nil {
		fmt.Printf("Could not connect: %s\n", err)
		os.Exit(1)
	}
	defer conn.Disconnect()

	// Open channel
	ch, err := conn.NewAPIChannel()
	if err != nil {
		fmt.Printf("Could not open API channel: %s\n", err)
		os.Exit(1)
	}
	defer ch.Close()
	if err != nil {
		fmt.Printf("Could not open API channel: %s\n", err)
		os.Exit(1)
	}

	err = addSourceBasedRouting(ch, "10.240.165.1", "10.0.0.2", 0)
	if err != nil {
		fmt.Printf("Could not add source-based routing: %s\n", err)
		os.Exit(1)
	}

}
