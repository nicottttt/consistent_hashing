# Implement consistent hashing in vpp
## Configure network architecture in linux network namespace and VPP
Create namespace first:
```
sudo ip netns add server1
sudo ip netns add server2
sudo ip netns add server3
```

create veth-pair for each namespace:
```
ip link add veth-vpp1 type veth peer name veth-server1
ip link add veth-vpp2 type veth peer name veth-server2
ip link add veth-vpp3 type veth peer name veth-server3
```

set veth-server{n} to namespace:
```
ip link set veth-server1 netns server1
ip link set veth-server2 netns server2
ip link set veth-server3 netns server3
```

set up the interface in each namespace:
```
ip netns exec server1 ip addr add 10.0.1.2/24 dev veth-server1
ip netns exec server1 ip link set veth-server1 up

ip netns exec server2 ip addr add 10.0.2.2/24 dev veth-server2
ip netns exec server2 ip link set veth-server2 up

ip netns exec server3 ip addr add 10.0.3.2/24 dev veth-server3
ip netns exec server3 ip link set veth-server3 up
```

create related routes in namespace:
```
sudo ip netns exec server1 ip route add 10.0.1.0/24 via 10.0.1.2
sudo ip netns exec server2 ip route add 10.0.2.0/24 via 10.0.2.2
sudo ip netns exec server3 ip route add 10.0.3.0/24 via 10.0.3.2
```

link and set up related interface in VPP:
```
create host-interface name veth-vpp1
set interface ip address host-veth-vpp1 10.0.1.1/24
set interface state host-veth-vpp1 up
ip route add 10.0.1.2/32 via host-veth-vpp1

create host-interface name veth-vpp2
set interface ip address host-veth-vpp1 10.0.2.1/24
set interface state host-veth-vpp2 up
ip route add 10.0.2.2/32 via host-veth-vpp2

create host-interface name veth-vpp3
set interface ip address host-veth-vpp1 10.0.3.1/24
set interface state host-veth-vpp3 up
ip route add 10.0.3.2/32 via host-veth-vpp3
```
## TODO