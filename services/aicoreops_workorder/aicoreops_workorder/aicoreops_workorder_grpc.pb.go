// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_workorder.proto

package aicoreops_workorder

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
	AicoreopsWorkorder_Ping_FullMethodName = "/aicoreops_workorder.Aicoreops_workorder/Ping"
)

// AicoreopsWorkorderClient is the client API for AicoreopsWorkorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AicoreopsWorkorderClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type aicoreopsWorkorderClient struct {
	cc grpc.ClientConnInterface
}

func NewAicoreopsWorkorderClient(cc grpc.ClientConnInterface) AicoreopsWorkorderClient {
	return &aicoreopsWorkorderClient{cc}
}

func (c *aicoreopsWorkorderClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, AicoreopsWorkorder_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AicoreopsWorkorderServer is the server API for AicoreopsWorkorder service.
// All implementations must embed UnimplementedAicoreopsWorkorderServer
// for forward compatibility
type AicoreopsWorkorderServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAicoreopsWorkorderServer()
}

// UnimplementedAicoreopsWorkorderServer must be embedded to have forward compatible implementations.
type UnimplementedAicoreopsWorkorderServer struct {
}

func (UnimplementedAicoreopsWorkorderServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAicoreopsWorkorderServer) mustEmbedUnimplementedAicoreopsWorkorderServer() {}

// UnsafeAicoreopsWorkorderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AicoreopsWorkorderServer will
// result in compilation errors.
type UnsafeAicoreopsWorkorderServer interface {
	mustEmbedUnimplementedAicoreopsWorkorderServer()
}

func RegisterAicoreopsWorkorderServer(s grpc.ServiceRegistrar, srv AicoreopsWorkorderServer) {
	s.RegisterService(&AicoreopsWorkorder_ServiceDesc, srv)
}

func _AicoreopsWorkorder_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AicoreopsWorkorderServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AicoreopsWorkorder_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AicoreopsWorkorderServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AicoreopsWorkorder_ServiceDesc is the grpc.ServiceDesc for AicoreopsWorkorder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AicoreopsWorkorder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aicoreops_workorder.Aicoreops_workorder",
	HandlerType: (*AicoreopsWorkorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AicoreopsWorkorder_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_workorder.proto",
}