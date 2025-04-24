#!/bin/bash
sudo ip netns add server1
sudo ip netns add server2
sudo ip netns add server3
ip link add veth-vpp1 type veth peer name veth-server1
ip link add veth-vpp2 type veth peer name veth-server2
ip link add veth-vpp3 type veth peer name veth-server3
ip link set veth-server1 netns server1
ip link set veth-server2 netns server2
ip link set veth-server3 netns server3
ip netns exec server1 ip addr add 10.0.1.2/24 dev veth-server1
ip netns exec server1 ip link set veth-server1 up

ip netns exec server2 ip addr add 10.0.2.2/24 dev veth-server2
ip netns exec server2 ip link set veth-server2 up

ip netns exec server3 ip addr add 10.0.3.2/24 dev veth-server3
ip netns exec server3 ip link set veth-server3 up

sudo ip netns exec server1 route add -net 10.240.165.0/24 gw 10.0.1.1
sudo ip netns exec server2 route add -net 10.240.165.0/24 gw 10.0.2.1
sudo ip netns exec server3 route add -net 10.240.165.0/24 gw 10.0.3.1
