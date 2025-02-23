// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: server/v1/demo.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Demo_DemoCreateName_FullMethodName = "/server.v1.demo/DemoCreateName"
)

// DemoClient is the client API for Demo server.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting server definition.
type DemoClient interface {
	// Sends a greeting
	DemoCreateName(ctx context.Context, in *DemoCreateNameRequest, opts ...grpc.CallOption) (*DemoCreateNameResponse, error)
}

type demoClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoClient(cc grpc.ClientConnInterface) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) DemoCreateName(ctx context.Context, in *DemoCreateNameRequest, opts ...grpc.CallOption) (*DemoCreateNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoCreateNameResponse)
	err := c.cc.Invoke(ctx, Demo_DemoCreateName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo server.
// All implementations must embed UnimplementedDemoServer
// for forward compatibility.
//
// The greeting server definition.
type DemoServer interface {
	// Sends a greeting
	DemoCreateName(context.Context, *DemoCreateNameRequest) (*DemoCreateNameResponse, error)
	mustEmbedUnimplementedDemoServer()
}

// UnimplementedDemoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDemoServer struct{}

func (UnimplementedDemoServer) DemoCreateName(context.Context, *DemoCreateNameRequest) (*DemoCreateNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoCreateName not implemented")
}
func (UnimplementedDemoServer) mustEmbedUnimplementedDemoServer() {}
func (UnimplementedDemoServer) testEmbeddedByValue()              {}

// UnsafeDemoServer may be embedded to opt out of forward compatibility for this server.
// Use of this interface is not recommended, as added methods to DemoServer will
// result in compilation errors.
type UnsafeDemoServer interface {
	mustEmbedUnimplementedDemoServer()
}

func RegisterDemoServer(s grpc.ServiceRegistrar, srv DemoServer) {
	// If the following call pancis, it indicates UnimplementedDemoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Demo_ServiceDesc, srv)
}

func _Demo_DemoCreateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoCreateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).DemoCreateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Demo_DemoCreateName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).DemoCreateName(ctx, req.(*DemoCreateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Demo_ServiceDesc is the grpc.ServiceDesc for Demo server.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.v1.demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DemoCreateName",
			Handler:    _Demo_DemoCreateName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/v1/demo.proto",
}
