// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_api.proto

package types

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
	ApiService_CreateApi_FullMethodName = "/aicoreops_api.ApiService/CreateApi"
	ApiService_GetApi_FullMethodName    = "/aicoreops_api.ApiService/GetApi"
	ApiService_UpdateApi_FullMethodName = "/aicoreops_api.ApiService/UpdateApi"
	ApiService_DeleteApi_FullMethodName = "/aicoreops_api.ApiService/DeleteApi"
	ApiService_ListApis_FullMethodName  = "/aicoreops_api.ApiService/ListApis"
)

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// API 服务定义
type ApiServiceClient interface {
	// 创建 API
	CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiResponse, error)
	// 获取 API 详情
	GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiResponse, error)
	// 更新 API
	UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiResponse, error)
	// 删除 API
	DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiResponse, error)
	// 列出 APIs
	ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateApiResponse)
	err := c.cc.Invoke(ctx, ApiService_CreateApi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApiResponse)
	err := c.cc.Invoke(ctx, ApiService_GetApi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateApiResponse)
	err := c.cc.Invoke(ctx, ApiService_UpdateApi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteApiResponse)
	err := c.cc.Invoke(ctx, ApiService_DeleteApi_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ListApis(ctx context.Context, in *ListApisRequest, opts ...grpc.CallOption) (*ListApisResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApisResponse)
	err := c.cc.Invoke(ctx, ApiService_ListApis_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
//
// API 服务定义
type ApiServiceServer interface {
	// 创建 API
	CreateApi(context.Context, *CreateApiRequest) (*CreateApiResponse, error)
	// 获取 API 详情
	GetApi(context.Context, *GetApiRequest) (*GetApiResponse, error)
	// 更新 API
	UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiResponse, error)
	// 删除 API
	DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiResponse, error)
	// 列出 APIs
	ListApis(context.Context, *ListApisRequest) (*ListApisResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) CreateApi(context.Context, *CreateApiRequest) (*CreateApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedApiServiceServer) GetApi(context.Context, *GetApiRequest) (*GetApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApi not implemented")
}
func (UnimplementedApiServiceServer) UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedApiServiceServer) DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedApiServiceServer) ListApis(context.Context, *ListApisRequest) (*ListApisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApis not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_CreateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateApi(ctx, req.(*CreateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetApi(ctx, req.(*GetApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_UpdateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).UpdateApi(ctx, req.(*UpdateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_DeleteApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteApi(ctx, req.(*DeleteApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ListApis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ListApis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_ListApis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ListApis(ctx, req.(*ListApisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aicoreops_api.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApi",
			Handler:    _ApiService_CreateApi_Handler,
		},
		{
			MethodName: "GetApi",
			Handler:    _ApiService_GetApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _ApiService_UpdateApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _ApiService_DeleteApi_Handler,
		},
		{
			MethodName: "ListApis",
			Handler:    _ApiService_ListApis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_api.proto",
}
