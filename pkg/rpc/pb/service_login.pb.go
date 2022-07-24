// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_login.proto

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

// 查找认证
type FindEnabledLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId int64  `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	UserId  int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *FindEnabledLoginRequest) Reset() {
	*x = FindEnabledLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledLoginRequest) ProtoMessage() {}

func (x *FindEnabledLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledLoginRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledLoginRequest) Descriptor() ([]byte, []int) {
	return file_service_login_proto_rawDescGZIP(), []int{0}
}

func (x *FindEnabledLoginRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *FindEnabledLoginRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindEnabledLoginRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type FindEnabledLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login *Login `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *FindEnabledLoginResponse) Reset() {
	*x = FindEnabledLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledLoginResponse) ProtoMessage() {}

func (x *FindEnabledLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledLoginResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledLoginResponse) Descriptor() ([]byte, []int) {
	return file_service_login_proto_rawDescGZIP(), []int{1}
}

func (x *FindEnabledLoginResponse) GetLogin() *Login {
	if x != nil {
		return x.Login
	}
	return nil
}

// 修改认证
type UpdateLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login *Login `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *UpdateLoginRequest) Reset() {
	*x = UpdateLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLoginRequest) ProtoMessage() {}

func (x *UpdateLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLoginRequest.ProtoReflect.Descriptor instead.
func (*UpdateLoginRequest) Descriptor() ([]byte, []int) {
	return file_service_login_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateLoginRequest) GetLogin() *Login {
	if x != nil {
		return x.Login
	}
	return nil
}

var File_service_login_proto protoreflect.FileDescriptor

var file_service_login_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3b, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x35, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x32, 0x94, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_login_proto_rawDescOnce sync.Once
	file_service_login_proto_rawDescData = file_service_login_proto_rawDesc
)

func file_service_login_proto_rawDescGZIP() []byte {
	file_service_login_proto_rawDescOnce.Do(func() {
		file_service_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_login_proto_rawDescData)
	})
	return file_service_login_proto_rawDescData
}

var file_service_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_login_proto_goTypes = []interface{}{
	(*FindEnabledLoginRequest)(nil),  // 0: pb.FindEnabledLoginRequest
	(*FindEnabledLoginResponse)(nil), // 1: pb.FindEnabledLoginResponse
	(*UpdateLoginRequest)(nil),       // 2: pb.UpdateLoginRequest
	(*Login)(nil),                    // 3: pb.Login
	(*RPCSuccess)(nil),               // 4: pb.RPCSuccess
}
var file_service_login_proto_depIdxs = []int32{
	3, // 0: pb.FindEnabledLoginResponse.login:type_name -> pb.Login
	3, // 1: pb.UpdateLoginRequest.login:type_name -> pb.Login
	0, // 2: pb.LoginService.findEnabledLogin:input_type -> pb.FindEnabledLoginRequest
	2, // 3: pb.LoginService.updateLogin:input_type -> pb.UpdateLoginRequest
	1, // 4: pb.LoginService.findEnabledLogin:output_type -> pb.FindEnabledLoginResponse
	4, // 5: pb.LoginService.updateLogin:output_type -> pb.RPCSuccess
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_login_proto_init() }
func file_service_login_proto_init() {
	if File_service_login_proto != nil {
		return
	}
	file_models_model_login_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledLoginRequest); i {
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
		file_service_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledLoginResponse); i {
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
		file_service_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLoginRequest); i {
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
			RawDescriptor: file_service_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_login_proto_goTypes,
		DependencyIndexes: file_service_login_proto_depIdxs,
		MessageInfos:      file_service_login_proto_msgTypes,
	}.Build()
	File_service_login_proto = out.File
	file_service_login_proto_rawDesc = nil
	file_service_login_proto_goTypes = nil
	file_service_login_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	// 查找认证
	FindEnabledLogin(ctx context.Context, in *FindEnabledLoginRequest, opts ...grpc.CallOption) (*FindEnabledLoginResponse, error)
	// 修改认证
	UpdateLogin(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) FindEnabledLogin(ctx context.Context, in *FindEnabledLoginRequest, opts ...grpc.CallOption) (*FindEnabledLoginResponse, error) {
	out := new(FindEnabledLoginResponse)
	err := c.cc.Invoke(ctx, "/pb.LoginService/findEnabledLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) UpdateLogin(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.LoginService/updateLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	// 查找认证
	FindEnabledLogin(context.Context, *FindEnabledLoginRequest) (*FindEnabledLoginResponse, error)
	// 修改认证
	UpdateLogin(context.Context, *UpdateLoginRequest) (*RPCSuccess, error)
}

// UnimplementedLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (*UnimplementedLoginServiceServer) FindEnabledLogin(context.Context, *FindEnabledLoginRequest) (*FindEnabledLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledLogin not implemented")
}
func (*UnimplementedLoginServiceServer) UpdateLogin(context.Context, *UpdateLoginRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLogin not implemented")
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_FindEnabledLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).FindEnabledLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoginService/FindEnabledLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).FindEnabledLogin(ctx, req.(*FindEnabledLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_UpdateLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).UpdateLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoginService/UpdateLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).UpdateLogin(ctx, req.(*UpdateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findEnabledLogin",
			Handler:    _LoginService_FindEnabledLogin_Handler,
		},
		{
			MethodName: "updateLogin",
			Handler:    _LoginService_UpdateLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_login.proto",
}
