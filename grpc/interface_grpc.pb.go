// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc/interface.proto

package request

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

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	Request(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) Request(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error) {
	out := new(Request)
	err := c.cc.Invoke(ctx, "/request.RequestService/request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	Request(context.Context, *Request) (*Request, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) Request(context.Context, *Request) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.RequestService/request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).Request(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "request.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "request",
			Handler:    _RequestService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}

// JoinServiceClient is the client API for JoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JoinServiceClient interface {
	Join(ctx context.Context, in *Join, opts ...grpc.CallOption) (*Join, error)
}

type joinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJoinServiceClient(cc grpc.ClientConnInterface) JoinServiceClient {
	return &joinServiceClient{cc}
}

func (c *joinServiceClient) Join(ctx context.Context, in *Join, opts ...grpc.CallOption) (*Join, error) {
	out := new(Join)
	err := c.cc.Invoke(ctx, "/request.JoinService/join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JoinServiceServer is the server API for JoinService service.
// All implementations must embed UnimplementedJoinServiceServer
// for forward compatibility
type JoinServiceServer interface {
	Join(context.Context, *Join) (*Join, error)
	mustEmbedUnimplementedJoinServiceServer()
}

// UnimplementedJoinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJoinServiceServer struct {
}

func (UnimplementedJoinServiceServer) Join(context.Context, *Join) (*Join, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedJoinServiceServer) mustEmbedUnimplementedJoinServiceServer() {}

// UnsafeJoinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JoinServiceServer will
// result in compilation errors.
type UnsafeJoinServiceServer interface {
	mustEmbedUnimplementedJoinServiceServer()
}

func RegisterJoinServiceServer(s grpc.ServiceRegistrar, srv JoinServiceServer) {
	s.RegisterService(&JoinService_ServiceDesc, srv)
}

func _JoinService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Join)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JoinServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/request.JoinService/join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JoinServiceServer).Join(ctx, req.(*Join))
	}
	return interceptor(ctx, in, info, handler)
}

// JoinService_ServiceDesc is the grpc.ServiceDesc for JoinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JoinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "request.JoinService",
	HandlerType: (*JoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "join",
			Handler:    _JoinService_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}