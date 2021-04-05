// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_message_task_log.proto

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

// 计算日志数量
type CountMessageTaskLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountMessageTaskLogsRequest) Reset() {
	*x = CountMessageTaskLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountMessageTaskLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountMessageTaskLogsRequest) ProtoMessage() {}

func (x *CountMessageTaskLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountMessageTaskLogsRequest.ProtoReflect.Descriptor instead.
func (*CountMessageTaskLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_log_proto_rawDescGZIP(), []int{0}
}

// 列出当页日志
type ListMessageTaskLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListMessageTaskLogsRequest) Reset() {
	*x = ListMessageTaskLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessageTaskLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessageTaskLogsRequest) ProtoMessage() {}

func (x *ListMessageTaskLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessageTaskLogsRequest.ProtoReflect.Descriptor instead.
func (*ListMessageTaskLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_log_proto_rawDescGZIP(), []int{1}
}

func (x *ListMessageTaskLogsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListMessageTaskLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListMessageTaskLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTaskLogs []*MessageTaskLog `protobuf:"bytes,1,rep,name=messageTaskLogs,proto3" json:"messageTaskLogs,omitempty"`
}

func (x *ListMessageTaskLogsResponse) Reset() {
	*x = ListMessageTaskLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessageTaskLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessageTaskLogsResponse) ProtoMessage() {}

func (x *ListMessageTaskLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessageTaskLogsResponse.ProtoReflect.Descriptor instead.
func (*ListMessageTaskLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_message_task_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListMessageTaskLogsResponse) GetMessageTaskLogs() []*MessageTaskLog {
	if x != nil {
		return x.MessageTaskLogs
	}
	return nil
}

var File_service_message_task_log_proto protoreflect.FileDescriptor

var file_service_message_task_log_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x1b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x5b, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x6f, 0x67, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x6f, 0x67, 0x73, 0x32, 0xbe, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x13, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_task_log_proto_rawDescOnce sync.Once
	file_service_message_task_log_proto_rawDescData = file_service_message_task_log_proto_rawDesc
)

func file_service_message_task_log_proto_rawDescGZIP() []byte {
	file_service_message_task_log_proto_rawDescOnce.Do(func() {
		file_service_message_task_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_task_log_proto_rawDescData)
	})
	return file_service_message_task_log_proto_rawDescData
}

var file_service_message_task_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_message_task_log_proto_goTypes = []interface{}{
	(*CountMessageTaskLogsRequest)(nil), // 0: pb.CountMessageTaskLogsRequest
	(*ListMessageTaskLogsRequest)(nil),  // 1: pb.ListMessageTaskLogsRequest
	(*ListMessageTaskLogsResponse)(nil), // 2: pb.ListMessageTaskLogsResponse
	(*MessageTaskLog)(nil),              // 3: pb.MessageTaskLog
	(*RPCCountResponse)(nil),            // 4: pb.RPCCountResponse
}
var file_service_message_task_log_proto_depIdxs = []int32{
	3, // 0: pb.ListMessageTaskLogsResponse.messageTaskLogs:type_name -> pb.MessageTaskLog
	0, // 1: pb.MessageTaskLogService.countMessageTaskLogs:input_type -> pb.CountMessageTaskLogsRequest
	1, // 2: pb.MessageTaskLogService.listMessageTaskLogs:input_type -> pb.ListMessageTaskLogsRequest
	4, // 3: pb.MessageTaskLogService.countMessageTaskLogs:output_type -> pb.RPCCountResponse
	2, // 4: pb.MessageTaskLogService.listMessageTaskLogs:output_type -> pb.ListMessageTaskLogsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_message_task_log_proto_init() }
func file_service_message_task_log_proto_init() {
	if File_service_message_task_log_proto != nil {
		return
	}
	file_models_message_task_log_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_task_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountMessageTaskLogsRequest); i {
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
		file_service_message_task_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessageTaskLogsRequest); i {
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
		file_service_message_task_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessageTaskLogsResponse); i {
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
			RawDescriptor: file_service_message_task_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_task_log_proto_goTypes,
		DependencyIndexes: file_service_message_task_log_proto_depIdxs,
		MessageInfos:      file_service_message_task_log_proto_msgTypes,
	}.Build()
	File_service_message_task_log_proto = out.File
	file_service_message_task_log_proto_rawDesc = nil
	file_service_message_task_log_proto_goTypes = nil
	file_service_message_task_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageTaskLogServiceClient is the client API for MessageTaskLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageTaskLogServiceClient interface {
	// 计算日志数量
	CountMessageTaskLogs(ctx context.Context, in *CountMessageTaskLogsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出当页日志
	ListMessageTaskLogs(ctx context.Context, in *ListMessageTaskLogsRequest, opts ...grpc.CallOption) (*ListMessageTaskLogsResponse, error)
}

type messageTaskLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageTaskLogServiceClient(cc grpc.ClientConnInterface) MessageTaskLogServiceClient {
	return &messageTaskLogServiceClient{cc}
}

func (c *messageTaskLogServiceClient) CountMessageTaskLogs(ctx context.Context, in *CountMessageTaskLogsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskLogService/countMessageTaskLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageTaskLogServiceClient) ListMessageTaskLogs(ctx context.Context, in *ListMessageTaskLogsRequest, opts ...grpc.CallOption) (*ListMessageTaskLogsResponse, error) {
	out := new(ListMessageTaskLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskLogService/listMessageTaskLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageTaskLogServiceServer is the server API for MessageTaskLogService service.
type MessageTaskLogServiceServer interface {
	// 计算日志数量
	CountMessageTaskLogs(context.Context, *CountMessageTaskLogsRequest) (*RPCCountResponse, error)
	// 列出当页日志
	ListMessageTaskLogs(context.Context, *ListMessageTaskLogsRequest) (*ListMessageTaskLogsResponse, error)
}

// UnimplementedMessageTaskLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageTaskLogServiceServer struct {
}

func (*UnimplementedMessageTaskLogServiceServer) CountMessageTaskLogs(context.Context, *CountMessageTaskLogsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountMessageTaskLogs not implemented")
}
func (*UnimplementedMessageTaskLogServiceServer) ListMessageTaskLogs(context.Context, *ListMessageTaskLogsRequest) (*ListMessageTaskLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageTaskLogs not implemented")
}

func RegisterMessageTaskLogServiceServer(s *grpc.Server, srv MessageTaskLogServiceServer) {
	s.RegisterService(&_MessageTaskLogService_serviceDesc, srv)
}

func _MessageTaskLogService_CountMessageTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountMessageTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskLogServiceServer).CountMessageTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskLogService/CountMessageTaskLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskLogServiceServer).CountMessageTaskLogs(ctx, req.(*CountMessageTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageTaskLogService_ListMessageTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskLogServiceServer).ListMessageTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskLogService/ListMessageTaskLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskLogServiceServer).ListMessageTaskLogs(ctx, req.(*ListMessageTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageTaskLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageTaskLogService",
	HandlerType: (*MessageTaskLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "countMessageTaskLogs",
			Handler:    _MessageTaskLogService_CountMessageTaskLogs_Handler,
		},
		{
			MethodName: "listMessageTaskLogs",
			Handler:    _MessageTaskLogService_ListMessageTaskLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_message_task_log.proto",
}
