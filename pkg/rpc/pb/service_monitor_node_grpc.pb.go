// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_monitor_node.proto

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
	MonitorNodeService_CreateMonitorNode_FullMethodName           = "/pb.MonitorNodeService/createMonitorNode"
	MonitorNodeService_UpdateMonitorNode_FullMethodName           = "/pb.MonitorNodeService/updateMonitorNode"
	MonitorNodeService_DeleteMonitorNode_FullMethodName           = "/pb.MonitorNodeService/deleteMonitorNode"
	MonitorNodeService_FindAllEnabledMonitorNodes_FullMethodName  = "/pb.MonitorNodeService/findAllEnabledMonitorNodes"
	MonitorNodeService_CountAllEnabledMonitorNodes_FullMethodName = "/pb.MonitorNodeService/countAllEnabledMonitorNodes"
	MonitorNodeService_ListEnabledMonitorNodes_FullMethodName     = "/pb.MonitorNodeService/listEnabledMonitorNodes"
	MonitorNodeService_FindEnabledMonitorNode_FullMethodName      = "/pb.MonitorNodeService/findEnabledMonitorNode"
	MonitorNodeService_FindCurrentMonitorNode_FullMethodName      = "/pb.MonitorNodeService/findCurrentMonitorNode"
	MonitorNodeService_UpdateMonitorNodeStatus_FullMethodName     = "/pb.MonitorNodeService/updateMonitorNodeStatus"
)

// MonitorNodeServiceClient is the client API for MonitorNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorNodeServiceClient interface {
	// 创建监控节点
	CreateMonitorNode(ctx context.Context, in *CreateMonitorNodeRequest, opts ...grpc.CallOption) (*CreateMonitorNodeResponse, error)
	// 修改监控节点
	UpdateMonitorNode(ctx context.Context, in *UpdateMonitorNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除监控节点
	DeleteMonitorNode(ctx context.Context, in *DeleteMonitorNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 列出所有可用监控节点
	FindAllEnabledMonitorNodes(ctx context.Context, in *FindAllEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledMonitorNodesResponse, error)
	// 计算监控节点数量
	CountAllEnabledMonitorNodes(ctx context.Context, in *CountAllEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的监控节点
	ListEnabledMonitorNodes(ctx context.Context, in *ListEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*ListEnabledMonitorNodesResponse, error)
	// 根据ID查找节点
	FindEnabledMonitorNode(ctx context.Context, in *FindEnabledMonitorNodeRequest, opts ...grpc.CallOption) (*FindEnabledMonitorNodeResponse, error)
	// 获取当前监控节点信息
	FindCurrentMonitorNode(ctx context.Context, in *FindCurrentMonitorNodeRequest, opts ...grpc.CallOption) (*FindCurrentMonitorNodeResponse, error)
	// 更新节点状态
	UpdateMonitorNodeStatus(ctx context.Context, in *UpdateMonitorNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type monitorNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorNodeServiceClient(cc grpc.ClientConnInterface) MonitorNodeServiceClient {
	return &monitorNodeServiceClient{cc}
}

func (c *monitorNodeServiceClient) CreateMonitorNode(ctx context.Context, in *CreateMonitorNodeRequest, opts ...grpc.CallOption) (*CreateMonitorNodeResponse, error) {
	out := new(CreateMonitorNodeResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_CreateMonitorNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) UpdateMonitorNode(ctx context.Context, in *UpdateMonitorNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, MonitorNodeService_UpdateMonitorNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) DeleteMonitorNode(ctx context.Context, in *DeleteMonitorNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, MonitorNodeService_DeleteMonitorNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) FindAllEnabledMonitorNodes(ctx context.Context, in *FindAllEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledMonitorNodesResponse, error) {
	out := new(FindAllEnabledMonitorNodesResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_FindAllEnabledMonitorNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) CountAllEnabledMonitorNodes(ctx context.Context, in *CountAllEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_CountAllEnabledMonitorNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) ListEnabledMonitorNodes(ctx context.Context, in *ListEnabledMonitorNodesRequest, opts ...grpc.CallOption) (*ListEnabledMonitorNodesResponse, error) {
	out := new(ListEnabledMonitorNodesResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_ListEnabledMonitorNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) FindEnabledMonitorNode(ctx context.Context, in *FindEnabledMonitorNodeRequest, opts ...grpc.CallOption) (*FindEnabledMonitorNodeResponse, error) {
	out := new(FindEnabledMonitorNodeResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_FindEnabledMonitorNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) FindCurrentMonitorNode(ctx context.Context, in *FindCurrentMonitorNodeRequest, opts ...grpc.CallOption) (*FindCurrentMonitorNodeResponse, error) {
	out := new(FindCurrentMonitorNodeResponse)
	err := c.cc.Invoke(ctx, MonitorNodeService_FindCurrentMonitorNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorNodeServiceClient) UpdateMonitorNodeStatus(ctx context.Context, in *UpdateMonitorNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, MonitorNodeService_UpdateMonitorNodeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorNodeServiceServer is the server API for MonitorNodeService service.
// All implementations should embed UnimplementedMonitorNodeServiceServer
// for forward compatibility
type MonitorNodeServiceServer interface {
	// 创建监控节点
	CreateMonitorNode(context.Context, *CreateMonitorNodeRequest) (*CreateMonitorNodeResponse, error)
	// 修改监控节点
	UpdateMonitorNode(context.Context, *UpdateMonitorNodeRequest) (*RPCSuccess, error)
	// 删除监控节点
	DeleteMonitorNode(context.Context, *DeleteMonitorNodeRequest) (*RPCSuccess, error)
	// 列出所有可用监控节点
	FindAllEnabledMonitorNodes(context.Context, *FindAllEnabledMonitorNodesRequest) (*FindAllEnabledMonitorNodesResponse, error)
	// 计算监控节点数量
	CountAllEnabledMonitorNodes(context.Context, *CountAllEnabledMonitorNodesRequest) (*RPCCountResponse, error)
	// 列出单页的监控节点
	ListEnabledMonitorNodes(context.Context, *ListEnabledMonitorNodesRequest) (*ListEnabledMonitorNodesResponse, error)
	// 根据ID查找节点
	FindEnabledMonitorNode(context.Context, *FindEnabledMonitorNodeRequest) (*FindEnabledMonitorNodeResponse, error)
	// 获取当前监控节点信息
	FindCurrentMonitorNode(context.Context, *FindCurrentMonitorNodeRequest) (*FindCurrentMonitorNodeResponse, error)
	// 更新节点状态
	UpdateMonitorNodeStatus(context.Context, *UpdateMonitorNodeStatusRequest) (*RPCSuccess, error)
}

// UnimplementedMonitorNodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMonitorNodeServiceServer struct {
}

func (UnimplementedMonitorNodeServiceServer) CreateMonitorNode(context.Context, *CreateMonitorNodeRequest) (*CreateMonitorNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonitorNode not implemented")
}
func (UnimplementedMonitorNodeServiceServer) UpdateMonitorNode(context.Context, *UpdateMonitorNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonitorNode not implemented")
}
func (UnimplementedMonitorNodeServiceServer) DeleteMonitorNode(context.Context, *DeleteMonitorNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMonitorNode not implemented")
}
func (UnimplementedMonitorNodeServiceServer) FindAllEnabledMonitorNodes(context.Context, *FindAllEnabledMonitorNodesRequest) (*FindAllEnabledMonitorNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledMonitorNodes not implemented")
}
func (UnimplementedMonitorNodeServiceServer) CountAllEnabledMonitorNodes(context.Context, *CountAllEnabledMonitorNodesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledMonitorNodes not implemented")
}
func (UnimplementedMonitorNodeServiceServer) ListEnabledMonitorNodes(context.Context, *ListEnabledMonitorNodesRequest) (*ListEnabledMonitorNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledMonitorNodes not implemented")
}
func (UnimplementedMonitorNodeServiceServer) FindEnabledMonitorNode(context.Context, *FindEnabledMonitorNodeRequest) (*FindEnabledMonitorNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledMonitorNode not implemented")
}
func (UnimplementedMonitorNodeServiceServer) FindCurrentMonitorNode(context.Context, *FindCurrentMonitorNodeRequest) (*FindCurrentMonitorNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCurrentMonitorNode not implemented")
}
func (UnimplementedMonitorNodeServiceServer) UpdateMonitorNodeStatus(context.Context, *UpdateMonitorNodeStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMonitorNodeStatus not implemented")
}

// UnsafeMonitorNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorNodeServiceServer will
// result in compilation errors.
type UnsafeMonitorNodeServiceServer interface {
	mustEmbedUnimplementedMonitorNodeServiceServer()
}

func RegisterMonitorNodeServiceServer(s grpc.ServiceRegistrar, srv MonitorNodeServiceServer) {
	s.RegisterService(&MonitorNodeService_ServiceDesc, srv)
}

func _MonitorNodeService_CreateMonitorNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMonitorNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).CreateMonitorNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_CreateMonitorNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).CreateMonitorNode(ctx, req.(*CreateMonitorNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_UpdateMonitorNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMonitorNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).UpdateMonitorNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_UpdateMonitorNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).UpdateMonitorNode(ctx, req.(*UpdateMonitorNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_DeleteMonitorNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMonitorNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).DeleteMonitorNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_DeleteMonitorNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).DeleteMonitorNode(ctx, req.(*DeleteMonitorNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_FindAllEnabledMonitorNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledMonitorNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).FindAllEnabledMonitorNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_FindAllEnabledMonitorNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).FindAllEnabledMonitorNodes(ctx, req.(*FindAllEnabledMonitorNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_CountAllEnabledMonitorNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledMonitorNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).CountAllEnabledMonitorNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_CountAllEnabledMonitorNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).CountAllEnabledMonitorNodes(ctx, req.(*CountAllEnabledMonitorNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_ListEnabledMonitorNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledMonitorNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).ListEnabledMonitorNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_ListEnabledMonitorNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).ListEnabledMonitorNodes(ctx, req.(*ListEnabledMonitorNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_FindEnabledMonitorNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledMonitorNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).FindEnabledMonitorNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_FindEnabledMonitorNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).FindEnabledMonitorNode(ctx, req.(*FindEnabledMonitorNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_FindCurrentMonitorNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCurrentMonitorNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).FindCurrentMonitorNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_FindCurrentMonitorNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).FindCurrentMonitorNode(ctx, req.(*FindCurrentMonitorNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorNodeService_UpdateMonitorNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMonitorNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorNodeServiceServer).UpdateMonitorNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonitorNodeService_UpdateMonitorNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorNodeServiceServer).UpdateMonitorNodeStatus(ctx, req.(*UpdateMonitorNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MonitorNodeService_ServiceDesc is the grpc.ServiceDesc for MonitorNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonitorNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MonitorNodeService",
	HandlerType: (*MonitorNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMonitorNode",
			Handler:    _MonitorNodeService_CreateMonitorNode_Handler,
		},
		{
			MethodName: "updateMonitorNode",
			Handler:    _MonitorNodeService_UpdateMonitorNode_Handler,
		},
		{
			MethodName: "deleteMonitorNode",
			Handler:    _MonitorNodeService_DeleteMonitorNode_Handler,
		},
		{
			MethodName: "findAllEnabledMonitorNodes",
			Handler:    _MonitorNodeService_FindAllEnabledMonitorNodes_Handler,
		},
		{
			MethodName: "countAllEnabledMonitorNodes",
			Handler:    _MonitorNodeService_CountAllEnabledMonitorNodes_Handler,
		},
		{
			MethodName: "listEnabledMonitorNodes",
			Handler:    _MonitorNodeService_ListEnabledMonitorNodes_Handler,
		},
		{
			MethodName: "findEnabledMonitorNode",
			Handler:    _MonitorNodeService_FindEnabledMonitorNode_Handler,
		},
		{
			MethodName: "findCurrentMonitorNode",
			Handler:    _MonitorNodeService_FindCurrentMonitorNode_Handler,
		},
		{
			MethodName: "updateMonitorNodeStatus",
			Handler:    _MonitorNodeService_UpdateMonitorNodeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_monitor_node.proto",
}
