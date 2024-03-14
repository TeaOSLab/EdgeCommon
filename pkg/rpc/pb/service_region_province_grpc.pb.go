// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_region_province.proto

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
	RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_FullMethodName = "/pb.RegionProvinceService/findAllEnabledRegionProvincesWithCountryId"
	RegionProvinceService_FindEnabledRegionProvince_FullMethodName                  = "/pb.RegionProvinceService/findEnabledRegionProvince"
	RegionProvinceService_FindAllRegionProvincesWithRegionCountryId_FullMethodName  = "/pb.RegionProvinceService/findAllRegionProvincesWithRegionCountryId"
	RegionProvinceService_FindAllRegionProvinces_FullMethodName                     = "/pb.RegionProvinceService/findAllRegionProvinces"
	RegionProvinceService_FindRegionProvince_FullMethodName                         = "/pb.RegionProvinceService/findRegionProvince"
	RegionProvinceService_UpdateRegionProvinceCustom_FullMethodName                 = "/pb.RegionProvinceService/updateRegionProvinceCustom"
)

// RegionProvinceServiceClient is the client API for RegionProvinceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionProvinceServiceClient interface {
	// Deprecated: Do not use.
	// 根据国家|地区ID查找所有省份
	FindAllEnabledRegionProvincesWithCountryId(ctx context.Context, in *FindAllEnabledRegionProvincesWithCountryIdRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error)
	// Deprecated: Do not use.
	// 查找单个省份信息
	FindEnabledRegionProvince(ctx context.Context, in *FindEnabledRegionProvinceRequest, opts ...grpc.CallOption) (*FindEnabledRegionProvinceResponse, error)
	// 根据国家|地区ID查找所有省份
	FindAllRegionProvincesWithRegionCountryId(ctx context.Context, in *FindAllRegionProvincesWithRegionCountryIdRequest, opts ...grpc.CallOption) (*FindAllRegionProvincesWithRegionCountryIdResponse, error)
	// 查找所有国家|地区的所有省份
	FindAllRegionProvinces(ctx context.Context, in *FindAllRegionProvincesRequest, opts ...grpc.CallOption) (*FindAllRegionProvincesResponse, error)
	// 查找单个省份信息
	FindRegionProvince(ctx context.Context, in *FindRegionProvinceRequest, opts ...grpc.CallOption) (*FindRegionProvinceResponse, error)
	// 修改国家/地区定制信息
	UpdateRegionProvinceCustom(ctx context.Context, in *UpdateRegionProvinceCustomRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type regionProvinceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionProvinceServiceClient(cc grpc.ClientConnInterface) RegionProvinceServiceClient {
	return &regionProvinceServiceClient{cc}
}

// Deprecated: Do not use.
func (c *regionProvinceServiceClient) FindAllEnabledRegionProvincesWithCountryId(ctx context.Context, in *FindAllEnabledRegionProvincesWithCountryIdRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error) {
	out := new(FindAllEnabledRegionProvincesWithCountryIdResponse)
	err := c.cc.Invoke(ctx, RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *regionProvinceServiceClient) FindEnabledRegionProvince(ctx context.Context, in *FindEnabledRegionProvinceRequest, opts ...grpc.CallOption) (*FindEnabledRegionProvinceResponse, error) {
	out := new(FindEnabledRegionProvinceResponse)
	err := c.cc.Invoke(ctx, RegionProvinceService_FindEnabledRegionProvince_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionProvinceServiceClient) FindAllRegionProvincesWithRegionCountryId(ctx context.Context, in *FindAllRegionProvincesWithRegionCountryIdRequest, opts ...grpc.CallOption) (*FindAllRegionProvincesWithRegionCountryIdResponse, error) {
	out := new(FindAllRegionProvincesWithRegionCountryIdResponse)
	err := c.cc.Invoke(ctx, RegionProvinceService_FindAllRegionProvincesWithRegionCountryId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionProvinceServiceClient) FindAllRegionProvinces(ctx context.Context, in *FindAllRegionProvincesRequest, opts ...grpc.CallOption) (*FindAllRegionProvincesResponse, error) {
	out := new(FindAllRegionProvincesResponse)
	err := c.cc.Invoke(ctx, RegionProvinceService_FindAllRegionProvinces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionProvinceServiceClient) FindRegionProvince(ctx context.Context, in *FindRegionProvinceRequest, opts ...grpc.CallOption) (*FindRegionProvinceResponse, error) {
	out := new(FindRegionProvinceResponse)
	err := c.cc.Invoke(ctx, RegionProvinceService_FindRegionProvince_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionProvinceServiceClient) UpdateRegionProvinceCustom(ctx context.Context, in *UpdateRegionProvinceCustomRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, RegionProvinceService_UpdateRegionProvinceCustom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionProvinceServiceServer is the server API for RegionProvinceService service.
// All implementations should embed UnimplementedRegionProvinceServiceServer
// for forward compatibility
type RegionProvinceServiceServer interface {
	// Deprecated: Do not use.
	// 根据国家|地区ID查找所有省份
	FindAllEnabledRegionProvincesWithCountryId(context.Context, *FindAllEnabledRegionProvincesWithCountryIdRequest) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error)
	// Deprecated: Do not use.
	// 查找单个省份信息
	FindEnabledRegionProvince(context.Context, *FindEnabledRegionProvinceRequest) (*FindEnabledRegionProvinceResponse, error)
	// 根据国家|地区ID查找所有省份
	FindAllRegionProvincesWithRegionCountryId(context.Context, *FindAllRegionProvincesWithRegionCountryIdRequest) (*FindAllRegionProvincesWithRegionCountryIdResponse, error)
	// 查找所有国家|地区的所有省份
	FindAllRegionProvinces(context.Context, *FindAllRegionProvincesRequest) (*FindAllRegionProvincesResponse, error)
	// 查找单个省份信息
	FindRegionProvince(context.Context, *FindRegionProvinceRequest) (*FindRegionProvinceResponse, error)
	// 修改国家/地区定制信息
	UpdateRegionProvinceCustom(context.Context, *UpdateRegionProvinceCustomRequest) (*RPCSuccess, error)
}

// UnimplementedRegionProvinceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRegionProvinceServiceServer struct {
}

func (UnimplementedRegionProvinceServiceServer) FindAllEnabledRegionProvincesWithCountryId(context.Context, *FindAllEnabledRegionProvincesWithCountryIdRequest) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledRegionProvincesWithCountryId not implemented")
}
func (UnimplementedRegionProvinceServiceServer) FindEnabledRegionProvince(context.Context, *FindEnabledRegionProvinceRequest) (*FindEnabledRegionProvinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledRegionProvince not implemented")
}
func (UnimplementedRegionProvinceServiceServer) FindAllRegionProvincesWithRegionCountryId(context.Context, *FindAllRegionProvincesWithRegionCountryIdRequest) (*FindAllRegionProvincesWithRegionCountryIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllRegionProvincesWithRegionCountryId not implemented")
}
func (UnimplementedRegionProvinceServiceServer) FindAllRegionProvinces(context.Context, *FindAllRegionProvincesRequest) (*FindAllRegionProvincesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllRegionProvinces not implemented")
}
func (UnimplementedRegionProvinceServiceServer) FindRegionProvince(context.Context, *FindRegionProvinceRequest) (*FindRegionProvinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRegionProvince not implemented")
}
func (UnimplementedRegionProvinceServiceServer) UpdateRegionProvinceCustom(context.Context, *UpdateRegionProvinceCustomRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegionProvinceCustom not implemented")
}

// UnsafeRegionProvinceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionProvinceServiceServer will
// result in compilation errors.
type UnsafeRegionProvinceServiceServer interface {
	mustEmbedUnimplementedRegionProvinceServiceServer()
}

func RegisterRegionProvinceServiceServer(s grpc.ServiceRegistrar, srv RegionProvinceServiceServer) {
	s.RegisterService(&RegionProvinceService_ServiceDesc, srv)
}

func _RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledRegionProvincesWithCountryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindAllEnabledRegionProvincesWithCountryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindAllEnabledRegionProvincesWithCountryId(ctx, req.(*FindAllEnabledRegionProvincesWithCountryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_FindEnabledRegionProvince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledRegionProvinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindEnabledRegionProvince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_FindEnabledRegionProvince_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindEnabledRegionProvince(ctx, req.(*FindEnabledRegionProvinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_FindAllRegionProvincesWithRegionCountryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRegionProvincesWithRegionCountryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindAllRegionProvincesWithRegionCountryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_FindAllRegionProvincesWithRegionCountryId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindAllRegionProvincesWithRegionCountryId(ctx, req.(*FindAllRegionProvincesWithRegionCountryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_FindAllRegionProvinces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRegionProvincesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindAllRegionProvinces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_FindAllRegionProvinces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindAllRegionProvinces(ctx, req.(*FindAllRegionProvincesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_FindRegionProvince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRegionProvinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindRegionProvince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_FindRegionProvince_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindRegionProvince(ctx, req.(*FindRegionProvinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_UpdateRegionProvinceCustom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionProvinceCustomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).UpdateRegionProvinceCustom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionProvinceService_UpdateRegionProvinceCustom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).UpdateRegionProvinceCustom(ctx, req.(*UpdateRegionProvinceCustomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionProvinceService_ServiceDesc is the grpc.ServiceDesc for RegionProvinceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionProvinceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegionProvinceService",
	HandlerType: (*RegionProvinceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledRegionProvincesWithCountryId",
			Handler:    _RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_Handler,
		},
		{
			MethodName: "findEnabledRegionProvince",
			Handler:    _RegionProvinceService_FindEnabledRegionProvince_Handler,
		},
		{
			MethodName: "findAllRegionProvincesWithRegionCountryId",
			Handler:    _RegionProvinceService_FindAllRegionProvincesWithRegionCountryId_Handler,
		},
		{
			MethodName: "findAllRegionProvinces",
			Handler:    _RegionProvinceService_FindAllRegionProvinces_Handler,
		},
		{
			MethodName: "findRegionProvince",
			Handler:    _RegionProvinceService_FindRegionProvince_Handler,
		},
		{
			MethodName: "updateRegionProvinceCustom",
			Handler:    _RegionProvinceService_UpdateRegionProvinceCustom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_region_province.proto",
}
