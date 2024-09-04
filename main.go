package main

import (
	"consistent/consistent"
	consistent_hash "consistent/consistent_hash"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	consistent_hash.UnimplementedConsistentHashServer
	consistent *consistent.Consistent
}

func (s *server) KeyMapServer(ctx context.Context, req *consistent_hash.MapkeyRequest) (*consistent_hash.MapkeyResponse, error) {
	result := s.consistent.MapKey(req.Server)
	return &consistent_hash.MapkeyResponse{Result: result}, nil
}

func (s *server) AddKey(ctx context.Context, req *consistent_hash.AddkeyRequest) (*consistent_hash.AddkeyResponse, error) {
	s.consistent.AddKey(req.Key)
	result := "Successfully added key to server"
	return &consistent_hash.AddkeyResponse{Result: result}, nil
}

func (s *server) RemoveServer(ctx context.Context, req *consistent_hash.RemoveServerRequest) (*consistent_hash.RemoveServerResponse, error) {
	s.consistent.DelServer(req.Server)
	result := s.consistent.GetMapping()
	return &consistent_hash.RemoveServerResponse{Result: result}, nil
}

func main() {
	c := consistent.NewRing(15)
	c.AddServer("Server1")
	c.AddServer("Server2")
	c.AddServer("Server3")
	c.AddServer("Server4")

	// c.TraverseServerList()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	consistent_hash.RegisterConsistentHashServer(s, &server{consistent: c})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
