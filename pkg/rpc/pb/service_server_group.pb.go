// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_server_group.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 创建分组
type CreateServerGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateServerGroupRequest) Reset() {
	*x = CreateServerGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerGroupRequest) ProtoMessage() {}

func (x *CreateServerGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateServerGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServerGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateServerGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroupId int64 `protobuf:"varint,1,opt,name=serverGroupId,proto3" json:"serverGroupId,omitempty"`
}

func (x *CreateServerGroupResponse) Reset() {
	*x = CreateServerGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerGroupResponse) ProtoMessage() {}

func (x *CreateServerGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateServerGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServerGroupResponse) GetServerGroupId() int64 {
	if x != nil {
		return x.ServerGroupId
	}
	return 0
}

// 修改分组
type UpdateServerGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroupId int64  `protobuf:"varint,1,opt,name=serverGroupId,proto3" json:"serverGroupId,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateServerGroupRequest) Reset() {
	*x = UpdateServerGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerGroupRequest) ProtoMessage() {}

func (x *UpdateServerGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateServerGroupRequest) GetServerGroupId() int64 {
	if x != nil {
		return x.ServerGroupId
	}
	return 0
}

func (x *UpdateServerGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 删除分组
type DeleteServerGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroupId int64 `protobuf:"varint,1,opt,name=serverGroupId,proto3" json:"serverGroupId,omitempty"`
}

func (x *DeleteServerGroupRequest) Reset() {
	*x = DeleteServerGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServerGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerGroupRequest) ProtoMessage() {}

func (x *DeleteServerGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteServerGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteServerGroupRequest) GetServerGroupId() int64 {
	if x != nil {
		return x.ServerGroupId
	}
	return 0
}

// 查询所有分组
type FindAllEnabledServerGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllEnabledServerGroupsRequest) Reset() {
	*x = FindAllEnabledServerGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledServerGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledServerGroupsRequest) ProtoMessage() {}

func (x *FindAllEnabledServerGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledServerGroupsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledServerGroupsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{4}
}

type FindAllEnabledServerGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroups []*ServerGroup `protobuf:"bytes,1,rep,name=serverGroups,proto3" json:"serverGroups,omitempty"`
}

func (x *FindAllEnabledServerGroupsResponse) Reset() {
	*x = FindAllEnabledServerGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledServerGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledServerGroupsResponse) ProtoMessage() {}

func (x *FindAllEnabledServerGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledServerGroupsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledServerGroupsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllEnabledServerGroupsResponse) GetServerGroups() []*ServerGroup {
	if x != nil {
		return x.ServerGroups
	}
	return nil
}

// 修改分组排序
type UpdateServerGroupOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroupIds []int64 `protobuf:"varint,1,rep,packed,name=serverGroupIds,proto3" json:"serverGroupIds,omitempty"`
}

func (x *UpdateServerGroupOrdersRequest) Reset() {
	*x = UpdateServerGroupOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServerGroupOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerGroupOrdersRequest) ProtoMessage() {}

func (x *UpdateServerGroupOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerGroupOrdersRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerGroupOrdersRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateServerGroupOrdersRequest) GetServerGroupIds() []int64 {
	if x != nil {
		return x.ServerGroupIds
	}
	return nil
}

// 查找单个分组信息
type FindEnabledServerGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroupId int64 `protobuf:"varint,1,opt,name=serverGroupId,proto3" json:"serverGroupId,omitempty"`
}

func (x *FindEnabledServerGroupRequest) Reset() {
	*x = FindEnabledServerGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledServerGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledServerGroupRequest) ProtoMessage() {}

func (x *FindEnabledServerGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledServerGroupRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledServerGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledServerGroupRequest) GetServerGroupId() int64 {
	if x != nil {
		return x.ServerGroupId
	}
	return 0
}

type FindEnabledServerGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroup *ServerGroup `protobuf:"bytes,1,opt,name=serverGroup,proto3" json:"serverGroup,omitempty"`
}

func (x *FindEnabledServerGroupResponse) Reset() {
	*x = FindEnabledServerGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledServerGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledServerGroupResponse) ProtoMessage() {}

func (x *FindEnabledServerGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledServerGroupResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledServerGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_server_group_proto_rawDescGZIP(), []int{8}
}

func (x *FindEnabledServerGroupResponse) GetServerGroup() *ServerGroup {
	if x != nil {
		return x.ServerGroup
	}
	return nil
}

var File_service_server_group_proto protoreflect.FileDescriptor

var file_service_server_group_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x54, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x22,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x48, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x73, 0x22, 0x45, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0x89, 0x04,
	0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6b, 0x0a, 0x1a,
	0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x17, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50,
	0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_group_proto_rawDescOnce sync.Once
	file_service_server_group_proto_rawDescData = file_service_server_group_proto_rawDesc
)

func file_service_server_group_proto_rawDescGZIP() []byte {
	file_service_server_group_proto_rawDescOnce.Do(func() {
		file_service_server_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_group_proto_rawDescData)
	})
	return file_service_server_group_proto_rawDescData
}

var file_service_server_group_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_server_group_proto_goTypes = []interface{}{
	(*CreateServerGroupRequest)(nil),           // 0: pb.CreateServerGroupRequest
	(*CreateServerGroupResponse)(nil),          // 1: pb.CreateServerGroupResponse
	(*UpdateServerGroupRequest)(nil),           // 2: pb.UpdateServerGroupRequest
	(*DeleteServerGroupRequest)(nil),           // 3: pb.DeleteServerGroupRequest
	(*FindAllEnabledServerGroupsRequest)(nil),  // 4: pb.FindAllEnabledServerGroupsRequest
	(*FindAllEnabledServerGroupsResponse)(nil), // 5: pb.FindAllEnabledServerGroupsResponse
	(*UpdateServerGroupOrdersRequest)(nil),     // 6: pb.UpdateServerGroupOrdersRequest
	(*FindEnabledServerGroupRequest)(nil),      // 7: pb.FindEnabledServerGroupRequest
	(*FindEnabledServerGroupResponse)(nil),     // 8: pb.FindEnabledServerGroupResponse
	(*ServerGroup)(nil),                        // 9: pb.ServerGroup
	(*RPCSuccess)(nil),                         // 10: pb.RPCSuccess
}
var file_service_server_group_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllEnabledServerGroupsResponse.serverGroups:type_name -> pb.ServerGroup
	9,  // 1: pb.FindEnabledServerGroupResponse.serverGroup:type_name -> pb.ServerGroup
	0,  // 2: pb.ServerGroupService.createServerGroup:input_type -> pb.CreateServerGroupRequest
	2,  // 3: pb.ServerGroupService.updateServerGroup:input_type -> pb.UpdateServerGroupRequest
	3,  // 4: pb.ServerGroupService.deleteServerGroup:input_type -> pb.DeleteServerGroupRequest
	4,  // 5: pb.ServerGroupService.findAllEnabledServerGroups:input_type -> pb.FindAllEnabledServerGroupsRequest
	6,  // 6: pb.ServerGroupService.updateServerGroupOrders:input_type -> pb.UpdateServerGroupOrdersRequest
	7,  // 7: pb.ServerGroupService.findEnabledServerGroup:input_type -> pb.FindEnabledServerGroupRequest
	1,  // 8: pb.ServerGroupService.createServerGroup:output_type -> pb.CreateServerGroupResponse
	10, // 9: pb.ServerGroupService.updateServerGroup:output_type -> pb.RPCSuccess
	10, // 10: pb.ServerGroupService.deleteServerGroup:output_type -> pb.RPCSuccess
	5,  // 11: pb.ServerGroupService.findAllEnabledServerGroups:output_type -> pb.FindAllEnabledServerGroupsResponse
	10, // 12: pb.ServerGroupService.updateServerGroupOrders:output_type -> pb.RPCSuccess
	8,  // 13: pb.ServerGroupService.findEnabledServerGroup:output_type -> pb.FindEnabledServerGroupResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_server_group_proto_init() }
func file_service_server_group_proto_init() {
	if File_service_server_group_proto != nil {
		return
	}
	file_models_model_server_group_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServerGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledServerGroupsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledServerGroupsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServerGroupOrdersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledServerGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_server_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledServerGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_server_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_group_proto_goTypes,
		DependencyIndexes: file_service_server_group_proto_depIdxs,
		MessageInfos:      file_service_server_group_proto_msgTypes,
	}.Build()
	File_service_server_group_proto = out.File
	file_service_server_group_proto_rawDesc = nil
	file_service_server_group_proto_goTypes = nil
	file_service_server_group_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerGroupServiceClient is the client API for ServerGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerGroupServiceClient interface {
	// 创建分组
	CreateServerGroup(ctx context.Context, in *CreateServerGroupRequest, opts ...grpc.CallOption) (*CreateServerGroupResponse, error)
	// 修改分组
	UpdateServerGroup(ctx context.Context, in *UpdateServerGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除分组
	DeleteServerGroup(ctx context.Context, in *DeleteServerGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledServerGroups(ctx context.Context, in *FindAllEnabledServerGroupsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerGroupsResponse, error)
	// 修改分组排序
	UpdateServerGroupOrders(ctx context.Context, in *UpdateServerGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledServerGroup(ctx context.Context, in *FindEnabledServerGroupRequest, opts ...grpc.CallOption) (*FindEnabledServerGroupResponse, error)
}

type serverGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerGroupServiceClient(cc grpc.ClientConnInterface) ServerGroupServiceClient {
	return &serverGroupServiceClient{cc}
}

func (c *serverGroupServiceClient) CreateServerGroup(ctx context.Context, in *CreateServerGroupRequest, opts ...grpc.CallOption) (*CreateServerGroupResponse, error) {
	out := new(CreateServerGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/createServerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverGroupServiceClient) UpdateServerGroup(ctx context.Context, in *UpdateServerGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/updateServerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverGroupServiceClient) DeleteServerGroup(ctx context.Context, in *DeleteServerGroupRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/deleteServerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverGroupServiceClient) FindAllEnabledServerGroups(ctx context.Context, in *FindAllEnabledServerGroupsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerGroupsResponse, error) {
	out := new(FindAllEnabledServerGroupsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/findAllEnabledServerGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverGroupServiceClient) UpdateServerGroupOrders(ctx context.Context, in *UpdateServerGroupOrdersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/updateServerGroupOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverGroupServiceClient) FindEnabledServerGroup(ctx context.Context, in *FindEnabledServerGroupRequest, opts ...grpc.CallOption) (*FindEnabledServerGroupResponse, error) {
	out := new(FindEnabledServerGroupResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerGroupService/findEnabledServerGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerGroupServiceServer is the server API for ServerGroupService service.
type ServerGroupServiceServer interface {
	// 创建分组
	CreateServerGroup(context.Context, *CreateServerGroupRequest) (*CreateServerGroupResponse, error)
	// 修改分组
	UpdateServerGroup(context.Context, *UpdateServerGroupRequest) (*RPCSuccess, error)
	// 删除分组
	DeleteServerGroup(context.Context, *DeleteServerGroupRequest) (*RPCSuccess, error)
	// 查询所有分组
	FindAllEnabledServerGroups(context.Context, *FindAllEnabledServerGroupsRequest) (*FindAllEnabledServerGroupsResponse, error)
	// 修改分组排序
	UpdateServerGroupOrders(context.Context, *UpdateServerGroupOrdersRequest) (*RPCSuccess, error)
	// 查找单个分组信息
	FindEnabledServerGroup(context.Context, *FindEnabledServerGroupRequest) (*FindEnabledServerGroupResponse, error)
}

// UnimplementedServerGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerGroupServiceServer struct {
}

func (*UnimplementedServerGroupServiceServer) CreateServerGroup(context.Context, *CreateServerGroupRequest) (*CreateServerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServerGroup not implemented")
}
func (*UnimplementedServerGroupServiceServer) UpdateServerGroup(context.Context, *UpdateServerGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServerGroup not implemented")
}
func (*UnimplementedServerGroupServiceServer) DeleteServerGroup(context.Context, *DeleteServerGroupRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServerGroup not implemented")
}
func (*UnimplementedServerGroupServiceServer) FindAllEnabledServerGroups(context.Context, *FindAllEnabledServerGroupsRequest) (*FindAllEnabledServerGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledServerGroups not implemented")
}
func (*UnimplementedServerGroupServiceServer) UpdateServerGroupOrders(context.Context, *UpdateServerGroupOrdersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServerGroupOrders not implemented")
}
func (*UnimplementedServerGroupServiceServer) FindEnabledServerGroup(context.Context, *FindEnabledServerGroupRequest) (*FindEnabledServerGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledServerGroup not implemented")
}

func RegisterServerGroupServiceServer(s *grpc.Server, srv ServerGroupServiceServer) {
	s.RegisterService(&_ServerGroupService_serviceDesc, srv)
}

func _ServerGroupService_CreateServerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).CreateServerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/CreateServerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).CreateServerGroup(ctx, req.(*CreateServerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerGroupService_UpdateServerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).UpdateServerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/UpdateServerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).UpdateServerGroup(ctx, req.(*UpdateServerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerGroupService_DeleteServerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).DeleteServerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/DeleteServerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).DeleteServerGroup(ctx, req.(*DeleteServerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerGroupService_FindAllEnabledServerGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledServerGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).FindAllEnabledServerGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/FindAllEnabledServerGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).FindAllEnabledServerGroups(ctx, req.(*FindAllEnabledServerGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerGroupService_UpdateServerGroupOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerGroupOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).UpdateServerGroupOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/UpdateServerGroupOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).UpdateServerGroupOrders(ctx, req.(*UpdateServerGroupOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerGroupService_FindEnabledServerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledServerGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerGroupServiceServer).FindEnabledServerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerGroupService/FindEnabledServerGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerGroupServiceServer).FindEnabledServerGroup(ctx, req.(*FindEnabledServerGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerGroupService",
	HandlerType: (*ServerGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createServerGroup",
			Handler:    _ServerGroupService_CreateServerGroup_Handler,
		},
		{
			MethodName: "updateServerGroup",
			Handler:    _ServerGroupService_UpdateServerGroup_Handler,
		},
		{
			MethodName: "deleteServerGroup",
			Handler:    _ServerGroupService_DeleteServerGroup_Handler,
		},
		{
			MethodName: "findAllEnabledServerGroups",
			Handler:    _ServerGroupService_FindAllEnabledServerGroups_Handler,
		},
		{
			MethodName: "updateServerGroupOrders",
			Handler:    _ServerGroupService_UpdateServerGroupOrders_Handler,
		},
		{
			MethodName: "findEnabledServerGroup",
			Handler:    _ServerGroupService_FindEnabledServerGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_group.proto",
}
