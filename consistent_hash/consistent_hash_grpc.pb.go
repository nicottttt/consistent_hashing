// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: consistent_hash.proto

package consistent_hash

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ConsistentHash_KeyMapServer_FullMethodName = "/consistent_hash.ConsistentHash/KeyMapServer"
	ConsistentHash_AddKey_FullMethodName       = "/consistent_hash.ConsistentHash/AddKey"
	ConsistentHash_RemoveServer_FullMethodName = "/consistent_hash.ConsistentHash/RemoveServer"
)

// ConsistentHashClient is the client API for ConsistentHash service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsistentHashClient interface {
	KeyMapServer(ctx context.Context, in *MapkeyRequest, opts ...grpc.CallOption) (*MapkeyResponse, error)
	AddKey(ctx context.Context, in *AddkeyRequest, opts ...grpc.CallOption) (*AddkeyResponse, error)
	RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*RemoveServerResponse, error)
}

type consistentHashClient struct {
	cc grpc.ClientConnInterface
}

func NewConsistentHashClient(cc grpc.ClientConnInterface) ConsistentHashClient {
	return &consistentHashClient{cc}
}

func (c *consistentHashClient) KeyMapServer(ctx context.Context, in *MapkeyRequest, opts ...grpc.CallOption) (*MapkeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MapkeyResponse)
	err := c.cc.Invoke(ctx, ConsistentHash_KeyMapServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentHashClient) AddKey(ctx context.Context, in *AddkeyRequest, opts ...grpc.CallOption) (*AddkeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddkeyResponse)
	err := c.cc.Invoke(ctx, ConsistentHash_AddKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consistentHashClient) RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*RemoveServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveServerResponse)
	err := c.cc.Invoke(ctx, ConsistentHash_RemoveServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsistentHashServer is the server API for ConsistentHash service.
// All implementations must embed UnimplementedConsistentHashServer
// for forward compatibility
type ConsistentHashServer interface {
	KeyMapServer(context.Context, *MapkeyRequest) (*MapkeyResponse, error)
	AddKey(context.Context, *AddkeyRequest) (*AddkeyResponse, error)
	RemoveServer(context.Context, *RemoveServerRequest) (*RemoveServerResponse, error)
	mustEmbedUnimplementedConsistentHashServer()
}

// UnimplementedConsistentHashServer must be embedded to have forward compatible implementations.
type UnimplementedConsistentHashServer struct {
}

func (UnimplementedConsistentHashServer) KeyMapServer(context.Context, *MapkeyRequest) (*MapkeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyMapServer not implemented")
}
func (UnimplementedConsistentHashServer) AddKey(context.Context, *AddkeyRequest) (*AddkeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKey not implemented")
}
func (UnimplementedConsistentHashServer) RemoveServer(context.Context, *RemoveServerRequest) (*RemoveServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServer not implemented")
}
func (UnimplementedConsistentHashServer) mustEmbedUnimplementedConsistentHashServer() {}

// UnsafeConsistentHashServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsistentHashServer will
// result in compilation errors.
type UnsafeConsistentHashServer interface {
	mustEmbedUnimplementedConsistentHashServer()
}

func RegisterConsistentHashServer(s grpc.ServiceRegistrar, srv ConsistentHashServer) {
	s.RegisterService(&ConsistentHash_ServiceDesc, srv)
}

func _ConsistentHash_KeyMapServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapkeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServer).KeyMapServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHash_KeyMapServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServer).KeyMapServer(ctx, req.(*MapkeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentHash_AddKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddkeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServer).AddKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHash_AddKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServer).AddKey(ctx, req.(*AddkeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsistentHash_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsistentHashServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsistentHash_RemoveServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsistentHashServer).RemoveServer(ctx, req.(*RemoveServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsistentHash_ServiceDesc is the grpc.ServiceDesc for ConsistentHash service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsistentHash_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consistent_hash.ConsistentHash",
	HandlerType: (*ConsistentHashServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyMapServer",
			Handler:    _ConsistentHash_KeyMapServer_Handler,
		},
		{
			MethodName: "AddKey",
			Handler:    _ConsistentHash_AddKey_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _ConsistentHash_RemoveServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consistent_hash.proto",
}
