// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pinecone

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SnapshotClientClient is the client API for SnapshotClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotClientClient interface {
	Call(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
}

type snapshotClientClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotClientClient(cc grpc.ClientConnInterface) SnapshotClientClient {
	return &snapshotClientClient{cc}
}

func (c *snapshotClientClient) Call(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := c.cc.Invoke(ctx, "/core.SnapshotClient/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotClientServer is the server API for SnapshotClient service.
// All implementations must embed UnimplementedSnapshotClientServer
// for forward compatibility
type SnapshotClientServer interface {
	Call(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
	mustEmbedUnimplementedSnapshotClientServer()
}

// UnimplementedSnapshotClientServer must be embedded to have forward compatible implementations.
type UnimplementedSnapshotClientServer struct {
}

func (UnimplementedSnapshotClientServer) Call(context.Context, *SnapshotRequest) (*SnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedSnapshotClientServer) mustEmbedUnimplementedSnapshotClientServer() {}

// UnsafeSnapshotClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotClientServer will
// result in compilation errors.
type UnsafeSnapshotClientServer interface {
	mustEmbedUnimplementedSnapshotClientServer()
}

func RegisterSnapshotClientServer(s grpc.ServiceRegistrar, srv SnapshotClientServer) {
	s.RegisterService(&SnapshotClient_ServiceDesc, srv)
}

func _SnapshotClient_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotClientServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.SnapshotClient/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotClientServer).Call(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnapshotClient_ServiceDesc is the grpc.ServiceDesc for SnapshotClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnapshotClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.SnapshotClient",
	HandlerType: (*SnapshotClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _SnapshotClient_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pinecone/core.proto",
}

// RPCClientClient is the client API for RPCClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClientClient interface {
	// Pass in a Request stream and completed Requests will be returned
	Call(ctx context.Context, opts ...grpc.CallOption) (RPCClient_CallClient, error)
	// Pass in a single request, and a completed request will be returned
	CallUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error)
}

type rPCClientClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClientClient(cc grpc.ClientConnInterface) RPCClientClient {
	return &rPCClientClient{cc}
}

func (c *rPCClientClient) Call(ctx context.Context, opts ...grpc.CallOption) (RPCClient_CallClient, error) {
	stream, err := c.cc.NewStream(ctx, &RPCClient_ServiceDesc.Streams[0], "/core.RPCClient/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCClientCallClient{stream}
	return x, nil
}

type RPCClient_CallClient interface {
	Send(*Request) error
	Recv() (*Request, error)
	grpc.ClientStream
}

type rPCClientCallClient struct {
	grpc.ClientStream
}

func (x *rPCClientCallClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCClientCallClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rPCClientClient) CallUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error) {
	out := new(Request)
	err := c.cc.Invoke(ctx, "/core.RPCClient/CallUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCClientServer is the server API for RPCClient service.
// All implementations must embed UnimplementedRPCClientServer
// for forward compatibility
type RPCClientServer interface {
	// Pass in a Request stream and completed Requests will be returned
	Call(RPCClient_CallServer) error
	// Pass in a single request, and a completed request will be returned
	CallUnary(context.Context, *Request) (*Request, error)
	mustEmbedUnimplementedRPCClientServer()
}

// UnimplementedRPCClientServer must be embedded to have forward compatible implementations.
type UnimplementedRPCClientServer struct {
}

func (UnimplementedRPCClientServer) Call(RPCClient_CallServer) error {
	return status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedRPCClientServer) CallUnary(context.Context, *Request) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallUnary not implemented")
}
func (UnimplementedRPCClientServer) mustEmbedUnimplementedRPCClientServer() {}

// UnsafeRPCClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCClientServer will
// result in compilation errors.
type UnsafeRPCClientServer interface {
	mustEmbedUnimplementedRPCClientServer()
}

func RegisterRPCClientServer(s grpc.ServiceRegistrar, srv RPCClientServer) {
	s.RegisterService(&RPCClient_ServiceDesc, srv)
}

func _RPCClient_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCClientServer).Call(&rPCClientCallServer{stream})
}

type RPCClient_CallServer interface {
	Send(*Request) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type rPCClientCallServer struct {
	grpc.ServerStream
}

func (x *rPCClientCallServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCClientCallServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RPCClient_CallUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCClientServer).CallUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RPCClient/CallUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCClientServer).CallUnary(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCClient_ServiceDesc is the grpc.ServiceDesc for RPCClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.RPCClient",
	HandlerType: (*RPCClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallUnary",
			Handler:    _RPCClient_CallUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _RPCClient_Call_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pinecone/core.proto",
}
