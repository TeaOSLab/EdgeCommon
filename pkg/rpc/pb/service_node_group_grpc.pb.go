// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_node_group.proto

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
	NodeGroupService_CreateNodeGroup_FullMethodName                           = "/pb.NodeGroupService/createNodeGroup"
	NodeGroupService_UpdateNodeGroup_FullMethodName                           = "/pb.NodeGroupService/updateNodeGroup"
	NodeGroupService_DeleteNodeGroup_FullMethodName                           = "/pb.NodeGroupService/deleteNodeGroup"
	NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_FullMethodName = "/pb.NodeGroupService/findAllEnabledNodeGroupsWithNodeClusterId"
	NodeGroupService_UpdateNodeGroupOrders_FullMethodName                     = "/pb.NodeGroupService/updateNodeGroupOrders"
	NodeGroupService_FindEnabledNodeGroup_FullMethodName                      = "/pb.NodeGroupService/findEnabledNodeGroup"
)

// NodeGroupServiceClient is the client API for NodeGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeGroupServiceClient interface {
	// 创建分组
	CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*CreateNodeGroupResponse, error)
	// 修改分组
	UpdateNodeGroup(ctx context.Context, in *UpdateNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除分组
	DeleteNodeGroup(ctx context.Context, in *DeleteNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledNodeGroupsWithNodeClusterId(ctx context.Context, in *FindAllEnabledNodeGroupsWithNodeClusterIdRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error)
	// 修改分组排序
	UpdateNodeGroupOrders(ctx context.Context, in *UpdateNodeGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledNodeGroup(ctx context.Context, in *FindEnabledNodeGroupRequest, opts ...grpc.CallOption) (*FindEnabledNodeGroupResponse, error)
}

type nodeGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeGroupServiceClient(cc grpc.ClientConnInterface) NodeGroupServiceClient {
	return &nodeGroupServiceClient{cc}
}

func (c *nodeGroupServiceClient) CreateNodeGroup(ctx context.Context, in *CreateNodeGroupRequest, opts ...grpc.CallOption) (*CreateNodeGroupResponse, error) {
	out := new(CreateNodeGroupResponse)
	err := c.cc.Invoke(ctx, NodeGroupService_CreateNodeGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) UpdateNodeGroup(ctx context.Context, in *UpdateNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodeGroupService_UpdateNodeGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) DeleteNodeGroup(ctx context.Context, in *DeleteNodeGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodeGroupService_DeleteNodeGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) FindAllEnabledNodeGroupsWithNodeClusterId(ctx context.Context, in *FindAllEnabledNodeGroupsWithNodeClusterIdRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error) {
	out := new(FindAllEnabledNodeGroupsWithNodeClusterIdResponse)
	err := c.cc.Invoke(ctx, NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) UpdateNodeGroupOrders(ctx context.Context, in *UpdateNodeGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, NodeGroupService_UpdateNodeGroupOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGroupServiceClient) FindEnabledNodeGroup(ctx context.Context, in *FindEnabledNodeGroupRequest, opts ...grpc.CallOption) (*FindEnabledNodeGroupResponse, error) {
	out := new(FindEnabledNodeGroupResponse)
	err := c.cc.Invoke(ctx, NodeGroupService_FindEnabledNodeGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeGroupServiceServer is the server API for NodeGroupService service.
// All implementations should embed UnimplementedNodeGroupServiceServer
// for forward compatibility
type NodeGroupServiceServer interface {
	// 创建分组
	CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*CreateNodeGroupResponse, error)
	// 修改分组
	UpdateNodeGroup(context.Context, *UpdateNodeGroupRequest) (*RPCSuccess, error)
	// 删除分组
	DeleteNodeGroup(context.Context, *DeleteNodeGroupRequest) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledNodeGroupsWithNodeClusterId(context.Context, *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error)
	// 修改分组排序
	UpdateNodeGroupOrders(context.Context, *UpdateNodeGroupOrdersRequest) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledNodeGroup(context.Context, *FindEnabledNodeGroupRequest) (*FindEnabledNodeGroupResponse, error)
}

// UnimplementedNodeGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNodeGroupServiceServer struct {
}

func (UnimplementedNodeGroupServiceServer) CreateNodeGroup(context.Context, *CreateNodeGroupRequest) (*CreateNodeGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeGroup not implemented")
}
func (UnimplementedNodeGroupServiceServer) UpdateNodeGroup(context.Context, *UpdateNodeGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeGroup not implemented")
}
func (UnimplementedNodeGroupServiceServer) DeleteNodeGroup(context.Context, *DeleteNodeGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeGroup not implemented")
}
func (UnimplementedNodeGroupServiceServer) FindAllEnabledNodeGroupsWithNodeClusterId(context.Context, *FindAllEnabledNodeGroupsWithNodeClusterIdRequest) (*FindAllEnabledNodeGroupsWithNodeClusterIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledNodeGroupsWithNodeClusterId not implemented")
}
func (UnimplementedNodeGroupServiceServer) UpdateNodeGroupOrders(context.Context, *UpdateNodeGroupOrdersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeGroupOrders not implemented")
}
func (UnimplementedNodeGroupServiceServer) FindEnabledNodeGroup(context.Context, *FindEnabledNodeGroupRequest) (*FindEnabledNodeGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledNodeGroup not implemented")
}

// UnsafeNodeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeGroupServiceServer will
// result in compilation errors.
type UnsafeNodeGroupServiceServer interface {
	mustEmbedUnimplementedNodeGroupServiceServer()
}

func RegisterNodeGroupServiceServer(s grpc.ServiceRegistrar, srv NodeGroupServiceServer) {
	s.RegisterService(&NodeGroupService_ServiceDesc, srv)
}

func _NodeGroupService_CreateNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).CreateNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_CreateNodeGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).CreateNodeGroup(ctx, req.(*CreateNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_UpdateNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).UpdateNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_UpdateNodeGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).UpdateNodeGroup(ctx, req.(*UpdateNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_DeleteNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).DeleteNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_DeleteNodeGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).DeleteNodeGroup(ctx, req.(*DeleteNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledNodeGroupsWithNodeClusterIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).FindAllEnabledNodeGroupsWithNodeClusterId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).FindAllEnabledNodeGroupsWithNodeClusterId(ctx, req.(*FindAllEnabledNodeGroupsWithNodeClusterIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_UpdateNodeGroupOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeGroupOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).UpdateNodeGroupOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_UpdateNodeGroupOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).UpdateNodeGroupOrders(ctx, req.(*UpdateNodeGroupOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGroupService_FindEnabledNodeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledNodeGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGroupServiceServer).FindEnabledNodeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeGroupService_FindEnabledNodeGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGroupServiceServer).FindEnabledNodeGroup(ctx, req.(*FindEnabledNodeGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeGroupService_ServiceDesc is the grpc.ServiceDesc for NodeGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeGroupService",
	HandlerType: (*NodeGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeGroup",
			Handler:    _NodeGroupService_CreateNodeGroup_Handler,
		},
		{
			MethodName: "updateNodeGroup",
			Handler:    _NodeGroupService_UpdateNodeGroup_Handler,
		},
		{
			MethodName: "deleteNodeGroup",
			Handler:    _NodeGroupService_DeleteNodeGroup_Handler,
		},
		{
			MethodName: "findAllEnabledNodeGroupsWithNodeClusterId",
			Handler:    _NodeGroupService_FindAllEnabledNodeGroupsWithNodeClusterId_Handler,
		},
		{
			MethodName: "updateNodeGroupOrders",
			Handler:    _NodeGroupService_UpdateNodeGroupOrders_Handler,
		},
		{
			MethodName: "findEnabledNodeGroup",
			Handler:    _NodeGroupService_FindEnabledNodeGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_group.proto",
}
