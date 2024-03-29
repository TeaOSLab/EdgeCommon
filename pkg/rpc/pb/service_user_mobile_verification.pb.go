// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_user_mobile_verification.proto

package pb

import (
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

// 认证手机号码
type VerifyUserMobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"` // 手机号
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`     // 激活码
}

func (x *VerifyUserMobileRequest) Reset() {
	*x = VerifyUserMobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserMobileRequest) ProtoMessage() {}

func (x *VerifyUserMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserMobileRequest.ProtoReflect.Descriptor instead.
func (*VerifyUserMobileRequest) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyUserMobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *VerifyUserMobileRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyUserMobileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`            // 手机号码对应的用户ID
	Mobile       string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`             // 手机号码
	ErrorCode    string `protobuf:"bytes,3,opt,name=errorCode,proto3" json:"errorCode,omitempty"`       // 错误代号，如果为空，说明没有错误
	ErrorMessage string `protobuf:"bytes,4,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"` // 错误信息
}

func (x *VerifyUserMobileResponse) Reset() {
	*x = VerifyUserMobileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserMobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserMobileResponse) ProtoMessage() {}

func (x *VerifyUserMobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserMobileResponse.ProtoReflect.Descriptor instead.
func (*VerifyUserMobileResponse) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyUserMobileResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VerifyUserMobileResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *VerifyUserMobileResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *VerifyUserMobileResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// 发送手机号码认证
type SendUserMobileVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"` // 待验证手机号码
}

func (x *SendUserMobileVerificationRequest) Reset() {
	*x = SendUserMobileVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserMobileVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserMobileVerificationRequest) ProtoMessage() {}

func (x *SendUserMobileVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserMobileVerificationRequest.ProtoReflect.Descriptor instead.
func (*SendUserMobileVerificationRequest) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{2}
}

func (x *SendUserMobileVerificationRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SendUserMobileVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk      bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`          // 是否发送成功
	ErrorCode string `protobuf:"bytes,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"` // 错误代号
}

func (x *SendUserMobileVerificationResponse) Reset() {
	*x = SendUserMobileVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserMobileVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserMobileVerificationResponse) ProtoMessage() {}

func (x *SendUserMobileVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserMobileVerificationResponse.ProtoReflect.Descriptor instead.
func (*SendUserMobileVerificationResponse) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{3}
}

func (x *SendUserMobileVerificationResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *SendUserMobileVerificationResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

// 查找用户正在等待激活的认证
type FindLatestUserMobileVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindLatestUserMobileVerificationRequest) Reset() {
	*x = FindLatestUserMobileVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLatestUserMobileVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLatestUserMobileVerificationRequest) ProtoMessage() {}

func (x *FindLatestUserMobileVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLatestUserMobileVerificationRequest.ProtoReflect.Descriptor instead.
func (*FindLatestUserMobileVerificationRequest) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{4}
}

type FindLatestUserMobileVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMobileVerification *UserMobileVerification `protobuf:"bytes,1,opt,name=userMobileVerification,proto3" json:"userMobileVerification,omitempty"`
}

func (x *FindLatestUserMobileVerificationResponse) Reset() {
	*x = FindLatestUserMobileVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_mobile_verification_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLatestUserMobileVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLatestUserMobileVerificationResponse) ProtoMessage() {}

func (x *FindLatestUserMobileVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_mobile_verification_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLatestUserMobileVerificationResponse.ProtoReflect.Descriptor instead.
func (*FindLatestUserMobileVerificationResponse) Descriptor() ([]byte, []int) {
	return file_service_user_mobile_verification_proto_rawDescGZIP(), []int{5}
}

func (x *FindLatestUserMobileVerificationResponse) GetUserMobileVerification() *UserMobileVerification {
	if x != nil {
		return x.UserMobileVerification
	}
	return nil
}

var File_service_user_mobile_verification_proto protoreflect.FileDescriptor

var file_service_user_mobile_verification_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x2b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x17, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x8c, 0x01, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3b, 0x0a, 0x21, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x56, 0x0a, 0x22,
	0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x7e, 0x0a, 0x28, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x16, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xda, 0x02, 0x0a, 0x1d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6b, 0x0a, 0x1a, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a,
	0x20, 0x66, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_mobile_verification_proto_rawDescOnce sync.Once
	file_service_user_mobile_verification_proto_rawDescData = file_service_user_mobile_verification_proto_rawDesc
)

func file_service_user_mobile_verification_proto_rawDescGZIP() []byte {
	file_service_user_mobile_verification_proto_rawDescOnce.Do(func() {
		file_service_user_mobile_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_mobile_verification_proto_rawDescData)
	})
	return file_service_user_mobile_verification_proto_rawDescData
}

var file_service_user_mobile_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_user_mobile_verification_proto_goTypes = []interface{}{
	(*VerifyUserMobileRequest)(nil),                  // 0: pb.VerifyUserMobileRequest
	(*VerifyUserMobileResponse)(nil),                 // 1: pb.VerifyUserMobileResponse
	(*SendUserMobileVerificationRequest)(nil),        // 2: pb.SendUserMobileVerificationRequest
	(*SendUserMobileVerificationResponse)(nil),       // 3: pb.SendUserMobileVerificationResponse
	(*FindLatestUserMobileVerificationRequest)(nil),  // 4: pb.FindLatestUserMobileVerificationRequest
	(*FindLatestUserMobileVerificationResponse)(nil), // 5: pb.FindLatestUserMobileVerificationResponse
	(*UserMobileVerification)(nil),                   // 6: pb.UserMobileVerification
}
var file_service_user_mobile_verification_proto_depIdxs = []int32{
	6, // 0: pb.FindLatestUserMobileVerificationResponse.userMobileVerification:type_name -> pb.UserMobileVerification
	0, // 1: pb.UserMobileVerificationService.verifyUserMobile:input_type -> pb.VerifyUserMobileRequest
	2, // 2: pb.UserMobileVerificationService.sendUserMobileVerification:input_type -> pb.SendUserMobileVerificationRequest
	4, // 3: pb.UserMobileVerificationService.findLatestUserMobileVerification:input_type -> pb.FindLatestUserMobileVerificationRequest
	1, // 4: pb.UserMobileVerificationService.verifyUserMobile:output_type -> pb.VerifyUserMobileResponse
	3, // 5: pb.UserMobileVerificationService.sendUserMobileVerification:output_type -> pb.SendUserMobileVerificationResponse
	5, // 6: pb.UserMobileVerificationService.findLatestUserMobileVerification:output_type -> pb.FindLatestUserMobileVerificationResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_mobile_verification_proto_init() }
func file_service_user_mobile_verification_proto_init() {
	if File_service_user_mobile_verification_proto != nil {
		return
	}
	file_models_model_user_mobile_verification_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_mobile_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserMobileRequest); i {
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
		file_service_user_mobile_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserMobileResponse); i {
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
		file_service_user_mobile_verification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserMobileVerificationRequest); i {
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
		file_service_user_mobile_verification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserMobileVerificationResponse); i {
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
		file_service_user_mobile_verification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLatestUserMobileVerificationRequest); i {
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
		file_service_user_mobile_verification_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLatestUserMobileVerificationResponse); i {
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
			RawDescriptor: file_service_user_mobile_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_mobile_verification_proto_goTypes,
		DependencyIndexes: file_service_user_mobile_verification_proto_depIdxs,
		MessageInfos:      file_service_user_mobile_verification_proto_msgTypes,
	}.Build()
	File_service_user_mobile_verification_proto = out.File
	file_service_user_mobile_verification_proto_rawDesc = nil
	file_service_user_mobile_verification_proto_goTypes = nil
	file_service_user_mobile_verification_proto_depIdxs = nil
}
