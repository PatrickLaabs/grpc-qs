// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: helloworld/helloworld.proto

package helloworld

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations should embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer should be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}

// ByeClient is the client API for Bye service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ByeClient interface {
	// sends bye message
	SayBye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeReply, error)
}

type byeClient struct {
	cc grpc.ClientConnInterface
}

func NewByeClient(cc grpc.ClientConnInterface) ByeClient {
	return &byeClient{cc}
}

func (c *byeClient) SayBye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeReply, error) {
	out := new(ByeReply)
	err := c.cc.Invoke(ctx, "/helloworld.Bye/SayBye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ByeServer is the server API for Bye service.
// All implementations should embed UnimplementedByeServer
// for forward compatibility
type ByeServer interface {
	// sends bye message
	SayBye(context.Context, *ByeRequest) (*ByeReply, error)
}

// UnimplementedByeServer should be embedded to have forward compatible implementations.
type UnimplementedByeServer struct {
}

func (UnimplementedByeServer) SayBye(context.Context, *ByeRequest) (*ByeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayBye not implemented")
}

// UnsafeByeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ByeServer will
// result in compilation errors.
type UnsafeByeServer interface {
	mustEmbedUnimplementedByeServer()
}

func RegisterByeServer(s grpc.ServiceRegistrar, srv ByeServer) {
	s.RegisterService(&Bye_ServiceDesc, srv)
}

func _Bye_SayBye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ByeServer).SayBye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Bye/SayBye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ByeServer).SayBye(ctx, req.(*ByeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bye_ServiceDesc is the grpc.ServiceDesc for Bye service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bye_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Bye",
	HandlerType: (*ByeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayBye",
			Handler:    _Bye_SayBye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}

// VersionClient is the client API for Version service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionClient interface {
	SendVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
}

type versionClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionClient(cc grpc.ClientConnInterface) VersionClient {
	return &versionClient{cc}
}

func (c *versionClient) SendVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := c.cc.Invoke(ctx, "/helloworld.Version/SendVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServer is the server API for Version service.
// All implementations should embed UnimplementedVersionServer
// for forward compatibility
type VersionServer interface {
	SendVersion(context.Context, *VersionRequest) (*VersionReply, error)
}

// UnimplementedVersionServer should be embedded to have forward compatible implementations.
type UnimplementedVersionServer struct {
}

func (UnimplementedVersionServer) SendVersion(context.Context, *VersionRequest) (*VersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVersion not implemented")
}

// UnsafeVersionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServer will
// result in compilation errors.
type UnsafeVersionServer interface {
	mustEmbedUnimplementedVersionServer()
}

func RegisterVersionServer(s grpc.ServiceRegistrar, srv VersionServer) {
	s.RegisterService(&Version_ServiceDesc, srv)
}

func _Version_SendVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).SendVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Version/SendVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).SendVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Version_ServiceDesc is the grpc.ServiceDesc for Version service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Version_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Version",
	HandlerType: (*VersionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVersion",
			Handler:    _Version_SendVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}

// EitcoClient is the client API for Eitco service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EitcoClient interface {
	SayEitco(ctx context.Context, in *HelloEitcoRequest, opts ...grpc.CallOption) (*HelloEitcoReply, error)
}

type eitcoClient struct {
	cc grpc.ClientConnInterface
}

func NewEitcoClient(cc grpc.ClientConnInterface) EitcoClient {
	return &eitcoClient{cc}
}

func (c *eitcoClient) SayEitco(ctx context.Context, in *HelloEitcoRequest, opts ...grpc.CallOption) (*HelloEitcoReply, error) {
	out := new(HelloEitcoReply)
	err := c.cc.Invoke(ctx, "/helloworld.Eitco/SayEitco", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EitcoServer is the server API for Eitco service.
// All implementations should embed UnimplementedEitcoServer
// for forward compatibility
type EitcoServer interface {
	SayEitco(context.Context, *HelloEitcoRequest) (*HelloEitcoReply, error)
}

// UnimplementedEitcoServer should be embedded to have forward compatible implementations.
type UnimplementedEitcoServer struct {
}

func (UnimplementedEitcoServer) SayEitco(context.Context, *HelloEitcoRequest) (*HelloEitcoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayEitco not implemented")
}

// UnsafeEitcoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EitcoServer will
// result in compilation errors.
type UnsafeEitcoServer interface {
	mustEmbedUnimplementedEitcoServer()
}

func RegisterEitcoServer(s grpc.ServiceRegistrar, srv EitcoServer) {
	s.RegisterService(&Eitco_ServiceDesc, srv)
}

func _Eitco_SayEitco_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloEitcoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EitcoServer).SayEitco(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Eitco/SayEitco",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EitcoServer).SayEitco(ctx, req.(*HelloEitcoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Eitco_ServiceDesc is the grpc.ServiceDesc for Eitco service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Eitco_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Eitco",
	HandlerType: (*EitcoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayEitco",
			Handler:    _Eitco_SayEitco_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
