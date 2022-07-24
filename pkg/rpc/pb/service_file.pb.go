// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_file.proto

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

// 查找文件
type FindEnabledFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId int64 `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *FindEnabledFileRequest) Reset() {
	*x = FindEnabledFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledFileRequest) ProtoMessage() {}

func (x *FindEnabledFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledFileRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledFileRequest) Descriptor() ([]byte, []int) {
	return file_service_file_proto_rawDescGZIP(), []int{0}
}

func (x *FindEnabledFileRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

type FindEnabledFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *File `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FindEnabledFileResponse) Reset() {
	*x = FindEnabledFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledFileResponse) ProtoMessage() {}

func (x *FindEnabledFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledFileResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledFileResponse) Descriptor() ([]byte, []int) {
	return file_service_file_proto_rawDescGZIP(), []int{1}
}

func (x *FindEnabledFileResponse) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

// 创建文件
type CreateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	IsPublic bool   `protobuf:"varint,3,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	MimeType string `protobuf:"bytes,4,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	Type     string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_service_file_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateFileRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateFileRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *CreateFileRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CreateFileRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CreateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId int64 `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *CreateFileResponse) Reset() {
	*x = CreateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileResponse) ProtoMessage() {}

func (x *CreateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileResponse.ProtoReflect.Descriptor instead.
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return file_service_file_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFileResponse) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

// 将文件置为已完成
type UpdateFileFinishedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId int64 `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *UpdateFileFinishedRequest) Reset() {
	*x = UpdateFileFinishedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileFinishedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileFinishedRequest) ProtoMessage() {}

func (x *UpdateFileFinishedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileFinishedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileFinishedRequest) Descriptor() ([]byte, []int) {
	return file_service_file_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFileFinishedRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

var File_service_file_proto protoreflect.FileDescriptor

var file_service_file_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x16,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x37,
	0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x32, 0xdb, 0x01, 0x0a,
	0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_file_proto_rawDescOnce sync.Once
	file_service_file_proto_rawDescData = file_service_file_proto_rawDesc
)

func file_service_file_proto_rawDescGZIP() []byte {
	file_service_file_proto_rawDescOnce.Do(func() {
		file_service_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_file_proto_rawDescData)
	})
	return file_service_file_proto_rawDescData
}

var file_service_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_file_proto_goTypes = []interface{}{
	(*FindEnabledFileRequest)(nil),    // 0: pb.FindEnabledFileRequest
	(*FindEnabledFileResponse)(nil),   // 1: pb.FindEnabledFileResponse
	(*CreateFileRequest)(nil),         // 2: pb.CreateFileRequest
	(*CreateFileResponse)(nil),        // 3: pb.CreateFileResponse
	(*UpdateFileFinishedRequest)(nil), // 4: pb.UpdateFileFinishedRequest
	(*File)(nil),                      // 5: pb.File
	(*RPCSuccess)(nil),                // 6: pb.RPCSuccess
}
var file_service_file_proto_depIdxs = []int32{
	5, // 0: pb.FindEnabledFileResponse.file:type_name -> pb.File
	0, // 1: pb.FileService.findEnabledFile:input_type -> pb.FindEnabledFileRequest
	2, // 2: pb.FileService.createFile:input_type -> pb.CreateFileRequest
	4, // 3: pb.FileService.updateFileFinished:input_type -> pb.UpdateFileFinishedRequest
	1, // 4: pb.FileService.findEnabledFile:output_type -> pb.FindEnabledFileResponse
	3, // 5: pb.FileService.createFile:output_type -> pb.CreateFileResponse
	6, // 6: pb.FileService.updateFileFinished:output_type -> pb.RPCSuccess
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_file_proto_init() }
func file_service_file_proto_init() {
	if File_service_file_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledFileRequest); i {
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
		file_service_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledFileResponse); i {
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
		file_service_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileRequest); i {
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
		file_service_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileResponse); i {
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
		file_service_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileFinishedRequest); i {
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
			RawDescriptor: file_service_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_file_proto_goTypes,
		DependencyIndexes: file_service_file_proto_depIdxs,
		MessageInfos:      file_service_file_proto_msgTypes,
	}.Build()
	File_service_file_proto = out.File
	file_service_file_proto_rawDesc = nil
	file_service_file_proto_goTypes = nil
	file_service_file_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	// 查找文件
	FindEnabledFile(ctx context.Context, in *FindEnabledFileRequest, opts ...grpc.CallOption) (*FindEnabledFileResponse, error)
	// 创建文件
	CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error)
	// 将文件置为已完成
	UpdateFileFinished(ctx context.Context, in *UpdateFileFinishedRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) FindEnabledFile(ctx context.Context, in *FindEnabledFileRequest, opts ...grpc.CallOption) (*FindEnabledFileResponse, error) {
	out := new(FindEnabledFileResponse)
	err := c.cc.Invoke(ctx, "/pb.FileService/findEnabledFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error) {
	out := new(CreateFileResponse)
	err := c.cc.Invoke(ctx, "/pb.FileService/createFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFileFinished(ctx context.Context, in *UpdateFileFinishedRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.FileService/updateFileFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	// 查找文件
	FindEnabledFile(context.Context, *FindEnabledFileRequest) (*FindEnabledFileResponse, error)
	// 创建文件
	CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error)
	// 将文件置为已完成
	UpdateFileFinished(context.Context, *UpdateFileFinishedRequest) (*RPCSuccess, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) FindEnabledFile(context.Context, *FindEnabledFileRequest) (*FindEnabledFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledFile not implemented")
}
func (*UnimplementedFileServiceServer) CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (*UnimplementedFileServiceServer) UpdateFileFinished(context.Context, *UpdateFileFinishedRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFileFinished not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_FindEnabledFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).FindEnabledFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileService/FindEnabledFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).FindEnabledFile(ctx, req.(*FindEnabledFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFile(ctx, req.(*CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFileFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileFinishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFileFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileService/UpdateFileFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFileFinished(ctx, req.(*UpdateFileFinishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findEnabledFile",
			Handler:    _FileService_FindEnabledFile_Handler,
		},
		{
			MethodName: "createFile",
			Handler:    _FileService_CreateFile_Handler,
		},
		{
			MethodName: "updateFileFinished",
			Handler:    _FileService_UpdateFileFinished_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_file.proto",
}
