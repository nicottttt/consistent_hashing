# Consistent hash in gateway demo
## Start 3 server
Starting 3 udp server listening localhost 8081, 8082, 8083
```
go run server/server.go 127.0.0.1:8081
go run server/server.go 127.0.0.1:8082
go run server/server.go 127.0.0.1:8083
```

## Start gateway
```
go run gateway/gateway.go
```

## Send msg to server
```
curl "http://localhost:8080/route?id="client1""
```

## start etcd in docker:
```
docker run -d --name etcd \
  -p 2379:2379 -p 2380:2380 \
  quay.io/coreos/etcd:v3.5.7 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-state new
```

check etcd key-value:
```
docker exec -it etcd etcdctl --endpoints=http://localhost:2379 get /nodes/ --prefix
```