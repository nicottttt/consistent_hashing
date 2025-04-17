# Consistent hash in gateway demo
## Start 3 server
Starting 3 udp server listening localhost 8081, 8082, 8083
```
go run server/server.go 8081
go run server/server.go 8082
go run server/server.go 8083
```

## Start gateway
```
go run gateway/gateway.go
```

## Add server into gateway
```
curl "http://localhost:8080/add?node="127.0.0.1:8081""
curl "http://localhost:8080/add?node="127.0.0.1:8082""
curl "http://localhost:8080/add?node="127.0.0.1:8083""
```

## Send msg to server
```
curl "http://localhost:8080/route?id="client1""
```