// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: models/model_node.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                   string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StatusJSON             []byte             `protobuf:"bytes,3,opt,name=statusJSON,proto3" json:"statusJSON,omitempty"`
	InstallDir             string             `protobuf:"bytes,4,opt,name=installDir,proto3" json:"installDir,omitempty"`
	IsInstalled            bool               `protobuf:"varint,5,opt,name=isInstalled,proto3" json:"isInstalled,omitempty"`
	Code                   string             `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	UniqueId               string             `protobuf:"bytes,7,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret                 string             `protobuf:"bytes,8,opt,name=secret,proto3" json:"secret,omitempty"`
	Version                int64              `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty"`
	LatestVersion          int64              `protobuf:"varint,10,opt,name=latestVersion,proto3" json:"latestVersion,omitempty"`
	ConnectedAPINodeIds    []int64            `protobuf:"varint,11,rep,packed,name=connectedAPINodeIds,proto3" json:"connectedAPINodeIds,omitempty"`
	MaxCPU                 int32              `protobuf:"varint,12,opt,name=maxCPU,proto3" json:"maxCPU,omitempty"`
	IsOn                   bool               `protobuf:"varint,13,opt,name=isOn,proto3" json:"isOn,omitempty"`
	IsUp                   bool               `protobuf:"varint,14,opt,name=isUp,proto3" json:"isUp,omitempty"`
	DnsRoutes              []*DNSRoute        `protobuf:"bytes,15,rep,name=dnsRoutes,proto3" json:"dnsRoutes,omitempty"`
	IsActive               bool               `protobuf:"varint,16,opt,name=isActive,proto3" json:"isActive,omitempty"`
	MaxCacheDiskCapacity   *SizeCapacity      `protobuf:"bytes,17,opt,name=maxCacheDiskCapacity,proto3" json:"maxCacheDiskCapacity,omitempty"`
	MaxCacheMemoryCapacity *SizeCapacity      `protobuf:"bytes,18,opt,name=maxCacheMemoryCapacity,proto3" json:"maxCacheMemoryCapacity,omitempty"`
	NodeCluster            *NodeCluster       `protobuf:"bytes,32,opt,name=nodeCluster,proto3" json:"nodeCluster,omitempty"`
	Login                  *NodeLogin         `protobuf:"bytes,33,opt,name=login,proto3" json:"login,omitempty"`
	InstallStatus          *NodeInstallStatus `protobuf:"bytes,34,opt,name=installStatus,proto3" json:"installStatus,omitempty"`
	IpAddresses            []*NodeIPAddress   `protobuf:"bytes,35,rep,name=ipAddresses,proto3" json:"ipAddresses,omitempty"`
	NodeGroup              *NodeGroup         `protobuf:"bytes,36,opt,name=nodeGroup,proto3" json:"nodeGroup,omitempty"`
	NodeRegion             *NodeRegion        `protobuf:"bytes,37,opt,name=nodeRegion,proto3" json:"nodeRegion,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_models_model_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetStatusJSON() []byte {
	if x != nil {
		return x.StatusJSON
	}
	return nil
}

func (x *Node) GetInstallDir() string {
	if x != nil {
		return x.InstallDir
	}
	return ""
}

func (x *Node) GetIsInstalled() bool {
	if x != nil {
		return x.IsInstalled
	}
	return false
}

func (x *Node) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Node) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *Node) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Node) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Node) GetLatestVersion() int64 {
	if x != nil {
		return x.LatestVersion
	}
	return 0
}

func (x *Node) GetConnectedAPINodeIds() []int64 {
	if x != nil {
		return x.ConnectedAPINodeIds
	}
	return nil
}

func (x *Node) GetMaxCPU() int32 {
	if x != nil {
		return x.MaxCPU
	}
	return 0
}

func (x *Node) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *Node) GetIsUp() bool {
	if x != nil {
		return x.IsUp
	}
	return false
}

func (x *Node) GetDnsRoutes() []*DNSRoute {
	if x != nil {
		return x.DnsRoutes
	}
	return nil
}

func (x *Node) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Node) GetMaxCacheDiskCapacity() *SizeCapacity {
	if x != nil {
		return x.MaxCacheDiskCapacity
	}
	return nil
}

func (x *Node) GetMaxCacheMemoryCapacity() *SizeCapacity {
	if x != nil {
		return x.MaxCacheMemoryCapacity
	}
	return nil
}

func (x *Node) GetNodeCluster() *NodeCluster {
	if x != nil {
		return x.NodeCluster
	}
	return nil
}

func (x *Node) GetLogin() *NodeLogin {
	if x != nil {
		return x.Login
	}
	return nil
}

func (x *Node) GetInstallStatus() *NodeInstallStatus {
	if x != nil {
		return x.InstallStatus
	}
	return nil
}

func (x *Node) GetIpAddresses() []*NodeIPAddress {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *Node) GetNodeGroup() *NodeGroup {
	if x != nil {
		return x.NodeGroup
	}
	return nil
}

func (x *Node) GetNodeRegion() *NodeRegion {
	if x != nil {
		return x.NodeRegion
	}
	return nil
}

var File_models_model_node_proto protoreflect.FileDescriptor

var file_models_model_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x07, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x44, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x44, 0x69, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x50, 0x49, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x03, 0x52, 0x13, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x50, 0x49, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x50, 0x55, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x50, 0x55, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x55, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x55,
	0x70, 0x12, 0x2a, 0x0a, 0x09, 0x64, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x44, 0x0a, 0x14, 0x6d, 0x61, 0x78,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a,
	0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x48, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x3b, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33,
	0x0a, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x23, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x24, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_node_proto_rawDescOnce sync.Once
	file_models_model_node_proto_rawDescData = file_models_model_node_proto_rawDesc
)

func file_models_model_node_proto_rawDescGZIP() []byte {
	file_models_model_node_proto_rawDescOnce.Do(func() {
		file_models_model_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_node_proto_rawDescData)
	})
	return file_models_model_node_proto_rawDescData
}

var file_models_model_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_node_proto_goTypes = []interface{}{
	(*Node)(nil),              // 0: pb.Node
	(*DNSRoute)(nil),          // 1: pb.DNSRoute
	(*SizeCapacity)(nil),      // 2: pb.SizeCapacity
	(*NodeCluster)(nil),       // 3: pb.NodeCluster
	(*NodeLogin)(nil),         // 4: pb.NodeLogin
	(*NodeInstallStatus)(nil), // 5: pb.NodeInstallStatus
	(*NodeIPAddress)(nil),     // 6: pb.NodeIPAddress
	(*NodeGroup)(nil),         // 7: pb.NodeGroup
	(*NodeRegion)(nil),        // 8: pb.NodeRegion
}
var file_models_model_node_proto_depIdxs = []int32{
	1, // 0: pb.Node.dnsRoutes:type_name -> pb.DNSRoute
	2, // 1: pb.Node.maxCacheDiskCapacity:type_name -> pb.SizeCapacity
	2, // 2: pb.Node.maxCacheMemoryCapacity:type_name -> pb.SizeCapacity
	3, // 3: pb.Node.nodeCluster:type_name -> pb.NodeCluster
	4, // 4: pb.Node.login:type_name -> pb.NodeLogin
	5, // 5: pb.Node.installStatus:type_name -> pb.NodeInstallStatus
	6, // 6: pb.Node.ipAddresses:type_name -> pb.NodeIPAddress
	7, // 7: pb.Node.nodeGroup:type_name -> pb.NodeGroup
	8, // 8: pb.Node.nodeRegion:type_name -> pb.NodeRegion
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_models_model_node_proto_init() }
func file_models_model_node_proto_init() {
	if File_models_model_node_proto != nil {
		return
	}
	file_models_model_node_cluster_proto_init()
	file_models_model_node_login_proto_init()
	file_models_model_node_install_status_proto_init()
	file_models_model_node_ip_address_proto_init()
	file_models_model_node_group_proto_init()
	file_models_model_node_region_proto_init()
	file_models_model_dns_route_proto_init()
	file_models_model_size_capacity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_models_model_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_node_proto_goTypes,
		DependencyIndexes: file_models_model_node_proto_depIdxs,
		MessageInfos:      file_models_model_node_proto_msgTypes,
	}.Build()
	File_models_model_node_proto = out.File
	file_models_model_node_proto_rawDesc = nil
	file_models_model_node_proto_goTypes = nil
	file_models_model_node_proto_depIdxs = nil
}
