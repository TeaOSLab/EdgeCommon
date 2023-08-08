// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_http_cache_task_key.proto

package pb

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

const (
	HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_FullMethodName     = "/pb.HTTPCacheTaskKeyService/validateHTTPCacheTaskKeys"
	HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_FullMethodName    = "/pb.HTTPCacheTaskKeyService/findDoingHTTPCacheTaskKeys"
	HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_FullMethodName = "/pb.HTTPCacheTaskKeyService/updateHTTPCacheTaskKeysStatus"
	HTTPCacheTaskKeyService_CountHTTPCacheTaskKeysWithDay_FullMethodName = "/pb.HTTPCacheTaskKeyService/countHTTPCacheTaskKeysWithDay"
)

// HTTPCacheTaskKeyServiceClient is the client API for HTTPCacheTaskKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPCacheTaskKeyServiceClient interface {
	// 校验缓存Key
	ValidateHTTPCacheTaskKeys(ctx context.Context, in *ValidateHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*ValidateHTTPCacheTaskKeysResponse, error)
	// 查找需要执行的Key
	FindDoingHTTPCacheTaskKeys(ctx context.Context, in *FindDoingHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*FindDoingHTTPCacheTaskKeysResponse, error)
	// 更新一组Key状态
	UpdateHTTPCacheTaskKeysStatus(ctx context.Context, in *UpdateHTTPCacheTaskKeysStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算当天已经清理的Key数量
	CountHTTPCacheTaskKeysWithDay(ctx context.Context, in *CountHTTPCacheTaskKeysWithDayRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
}

type hTTPCacheTaskKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPCacheTaskKeyServiceClient(cc grpc.ClientConnInterface) HTTPCacheTaskKeyServiceClient {
	return &hTTPCacheTaskKeyServiceClient{cc}
}

func (c *hTTPCacheTaskKeyServiceClient) ValidateHTTPCacheTaskKeys(ctx context.Context, in *ValidateHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*ValidateHTTPCacheTaskKeysResponse, error) {
	out := new(ValidateHTTPCacheTaskKeysResponse)
	err := c.cc.Invoke(ctx, HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCacheTaskKeyServiceClient) FindDoingHTTPCacheTaskKeys(ctx context.Context, in *FindDoingHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*FindDoingHTTPCacheTaskKeysResponse, error) {
	out := new(FindDoingHTTPCacheTaskKeysResponse)
	err := c.cc.Invoke(ctx, HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCacheTaskKeyServiceClient) UpdateHTTPCacheTaskKeysStatus(ctx context.Context, in *UpdateHTTPCacheTaskKeysStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCacheTaskKeyServiceClient) CountHTTPCacheTaskKeysWithDay(ctx context.Context, in *CountHTTPCacheTaskKeysWithDayRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, HTTPCacheTaskKeyService_CountHTTPCacheTaskKeysWithDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPCacheTaskKeyServiceServer is the server API for HTTPCacheTaskKeyService service.
// All implementations should embed UnimplementedHTTPCacheTaskKeyServiceServer
// for forward compatibility
type HTTPCacheTaskKeyServiceServer interface {
	// 校验缓存Key
	ValidateHTTPCacheTaskKeys(context.Context, *ValidateHTTPCacheTaskKeysRequest) (*ValidateHTTPCacheTaskKeysResponse, error)
	// 查找需要执行的Key
	FindDoingHTTPCacheTaskKeys(context.Context, *FindDoingHTTPCacheTaskKeysRequest) (*FindDoingHTTPCacheTaskKeysResponse, error)
	// 更新一组Key状态
	UpdateHTTPCacheTaskKeysStatus(context.Context, *UpdateHTTPCacheTaskKeysStatusRequest) (*RPCSuccess, error)
	// 计算当天已经清理的Key数量
	CountHTTPCacheTaskKeysWithDay(context.Context, *CountHTTPCacheTaskKeysWithDayRequest) (*RPCCountResponse, error)
}

// UnimplementedHTTPCacheTaskKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPCacheTaskKeyServiceServer struct {
}

func (UnimplementedHTTPCacheTaskKeyServiceServer) ValidateHTTPCacheTaskKeys(context.Context, *ValidateHTTPCacheTaskKeysRequest) (*ValidateHTTPCacheTaskKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateHTTPCacheTaskKeys not implemented")
}
func (UnimplementedHTTPCacheTaskKeyServiceServer) FindDoingHTTPCacheTaskKeys(context.Context, *FindDoingHTTPCacheTaskKeysRequest) (*FindDoingHTTPCacheTaskKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDoingHTTPCacheTaskKeys not implemented")
}
func (UnimplementedHTTPCacheTaskKeyServiceServer) UpdateHTTPCacheTaskKeysStatus(context.Context, *UpdateHTTPCacheTaskKeysStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPCacheTaskKeysStatus not implemented")
}
func (UnimplementedHTTPCacheTaskKeyServiceServer) CountHTTPCacheTaskKeysWithDay(context.Context, *CountHTTPCacheTaskKeysWithDayRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountHTTPCacheTaskKeysWithDay not implemented")
}

// UnsafeHTTPCacheTaskKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPCacheTaskKeyServiceServer will
// result in compilation errors.
type UnsafeHTTPCacheTaskKeyServiceServer interface {
	mustEmbedUnimplementedHTTPCacheTaskKeyServiceServer()
}

func RegisterHTTPCacheTaskKeyServiceServer(s grpc.ServiceRegistrar, srv HTTPCacheTaskKeyServiceServer) {
	s.RegisterService(&HTTPCacheTaskKeyService_ServiceDesc, srv)
}

func _HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateHTTPCacheTaskKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).ValidateHTTPCacheTaskKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).ValidateHTTPCacheTaskKeys(ctx, req.(*ValidateHTTPCacheTaskKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDoingHTTPCacheTaskKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).FindDoingHTTPCacheTaskKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).FindDoingHTTPCacheTaskKeys(ctx, req.(*FindDoingHTTPCacheTaskKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPCacheTaskKeysStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).UpdateHTTPCacheTaskKeysStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).UpdateHTTPCacheTaskKeysStatus(ctx, req.(*UpdateHTTPCacheTaskKeysStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCacheTaskKeyService_CountHTTPCacheTaskKeysWithDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountHTTPCacheTaskKeysWithDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).CountHTTPCacheTaskKeysWithDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPCacheTaskKeyService_CountHTTPCacheTaskKeysWithDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).CountHTTPCacheTaskKeysWithDay(ctx, req.(*CountHTTPCacheTaskKeysWithDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPCacheTaskKeyService_ServiceDesc is the grpc.ServiceDesc for HTTPCacheTaskKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPCacheTaskKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPCacheTaskKeyService",
	HandlerType: (*HTTPCacheTaskKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "validateHTTPCacheTaskKeys",
			Handler:    _HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_Handler,
		},
		{
			MethodName: "findDoingHTTPCacheTaskKeys",
			Handler:    _HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_Handler,
		},
		{
			MethodName: "updateHTTPCacheTaskKeysStatus",
			Handler:    _HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_Handler,
		},
		{
			MethodName: "countHTTPCacheTaskKeysWithDay",
			Handler:    _HTTPCacheTaskKeyService_CountHTTPCacheTaskKeysWithDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_cache_task_key.proto",
}