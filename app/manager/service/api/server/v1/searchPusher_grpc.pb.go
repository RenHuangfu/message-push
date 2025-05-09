// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: server/v1/searchPusher.proto

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
	SearchPusher_PushSearch_FullMethodName = "/server.v1.SearchPusher/PushSearch"
)

// SearchPusherClient is the client API for SearchPusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchPusherClient interface {
	PushSearch(ctx context.Context, in *PushSearchRequest, opts ...grpc.CallOption) (*PushSearchResponse, error)
}

type searchPusherClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchPusherClient(cc grpc.ClientConnInterface) SearchPusherClient {
	return &searchPusherClient{cc}
}

func (c *searchPusherClient) PushSearch(ctx context.Context, in *PushSearchRequest, opts ...grpc.CallOption) (*PushSearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushSearchResponse)
	err := c.cc.Invoke(ctx, SearchPusher_PushSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchPusherServer is the server API for SearchPusher service.
// All implementations must embed UnimplementedSearchPusherServer
// for forward compatibility.
type SearchPusherServer interface {
	PushSearch(context.Context, *PushSearchRequest) (*PushSearchResponse, error)
	mustEmbedUnimplementedSearchPusherServer()
}

// UnimplementedSearchPusherServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSearchPusherServer struct{}

func (UnimplementedSearchPusherServer) PushSearch(context.Context, *PushSearchRequest) (*PushSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushSearch not implemented")
}
func (UnimplementedSearchPusherServer) mustEmbedUnimplementedSearchPusherServer() {}
func (UnimplementedSearchPusherServer) testEmbeddedByValue()                      {}

// UnsafeSearchPusherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchPusherServer will
// result in compilation errors.
type UnsafeSearchPusherServer interface {
	mustEmbedUnimplementedSearchPusherServer()
}

func RegisterSearchPusherServer(s grpc.ServiceRegistrar, srv SearchPusherServer) {
	// If the following call pancis, it indicates UnimplementedSearchPusherServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchPusher_ServiceDesc, srv)
}

func _SearchPusher_PushSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchPusherServer).PushSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchPusher_PushSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchPusherServer).PushSearch(ctx, req.(*PushSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchPusher_ServiceDesc is the grpc.ServiceDesc for SearchPusher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchPusher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.v1.SearchPusher",
	HandlerType: (*SearchPusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushSearch",
			Handler:    _SearchPusher_PushSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/v1/searchPusher.proto",
}
