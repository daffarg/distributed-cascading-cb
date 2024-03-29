// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: circuitbreaker.proto

package protobuf

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

// CircuitBreakerClient is the client API for CircuitBreaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CircuitBreakerClient interface {
	GeneralRequest(ctx context.Context, in *GeneralRequestInput, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetInput, opts ...grpc.CallOption) (*Response, error)
	Post(ctx context.Context, in *PostInput, opts ...grpc.CallOption) (*Response, error)
	Put(ctx context.Context, in *PutInput, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteInput, opts ...grpc.CallOption) (*Response, error)
}

type circuitBreakerClient struct {
	cc grpc.ClientConnInterface
}

func NewCircuitBreakerClient(cc grpc.ClientConnInterface) CircuitBreakerClient {
	return &circuitBreakerClient{cc}
}

func (c *circuitBreakerClient) GeneralRequest(ctx context.Context, in *GeneralRequestInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protobuf.CircuitBreaker/GeneralRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circuitBreakerClient) Get(ctx context.Context, in *GetInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protobuf.CircuitBreaker/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circuitBreakerClient) Post(ctx context.Context, in *PostInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protobuf.CircuitBreaker/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circuitBreakerClient) Put(ctx context.Context, in *PutInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protobuf.CircuitBreaker/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circuitBreakerClient) Delete(ctx context.Context, in *DeleteInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protobuf.CircuitBreaker/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CircuitBreakerServer is the server API for CircuitBreaker service.
// All implementations must embed UnimplementedCircuitBreakerServer
// for forward compatibility
type CircuitBreakerServer interface {
	GeneralRequest(context.Context, *GeneralRequestInput) (*Response, error)
	Get(context.Context, *GetInput) (*Response, error)
	Post(context.Context, *PostInput) (*Response, error)
	Put(context.Context, *PutInput) (*Response, error)
	Delete(context.Context, *DeleteInput) (*Response, error)
	mustEmbedUnimplementedCircuitBreakerServer()
}

// UnimplementedCircuitBreakerServer must be embedded to have forward compatible implementations.
type UnimplementedCircuitBreakerServer struct {
}

func (UnimplementedCircuitBreakerServer) GeneralRequest(context.Context, *GeneralRequestInput) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneralRequest not implemented")
}
func (UnimplementedCircuitBreakerServer) Get(context.Context, *GetInput) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCircuitBreakerServer) Post(context.Context, *PostInput) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedCircuitBreakerServer) Put(context.Context, *PutInput) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedCircuitBreakerServer) Delete(context.Context, *DeleteInput) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCircuitBreakerServer) mustEmbedUnimplementedCircuitBreakerServer() {}

// UnsafeCircuitBreakerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CircuitBreakerServer will
// result in compilation errors.
type UnsafeCircuitBreakerServer interface {
	mustEmbedUnimplementedCircuitBreakerServer()
}

func RegisterCircuitBreakerServer(s grpc.ServiceRegistrar, srv CircuitBreakerServer) {
	s.RegisterService(&CircuitBreaker_ServiceDesc, srv)
}

func _CircuitBreaker_GeneralRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralRequestInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircuitBreakerServer).GeneralRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CircuitBreaker/GeneralRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircuitBreakerServer).GeneralRequest(ctx, req.(*GeneralRequestInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CircuitBreaker_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircuitBreakerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CircuitBreaker/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircuitBreakerServer).Get(ctx, req.(*GetInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CircuitBreaker_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircuitBreakerServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CircuitBreaker/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircuitBreakerServer).Post(ctx, req.(*PostInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CircuitBreaker_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircuitBreakerServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CircuitBreaker/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircuitBreakerServer).Put(ctx, req.(*PutInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CircuitBreaker_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CircuitBreakerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CircuitBreaker/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CircuitBreakerServer).Delete(ctx, req.(*DeleteInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CircuitBreaker_ServiceDesc is the grpc.ServiceDesc for CircuitBreaker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CircuitBreaker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.CircuitBreaker",
	HandlerType: (*CircuitBreakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneralRequest",
			Handler:    _CircuitBreaker_GeneralRequest_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CircuitBreaker_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _CircuitBreaker_Post_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _CircuitBreaker_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CircuitBreaker_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "circuitbreaker.proto",
}
