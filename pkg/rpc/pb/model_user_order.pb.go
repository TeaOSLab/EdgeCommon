// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_user_order.proto

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

// 用户订单
type UserOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64        `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Code          string       `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type          string       `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	OrderMethodId int64        `protobuf:"varint,5,opt,name=orderMethodId,proto3" json:"orderMethodId,omitempty"`
	Status        string       `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Amount        float32      `protobuf:"fixed32,7,opt,name=amount,proto3" json:"amount,omitempty"`
	ParamsJSON    []byte       `protobuf:"bytes,8,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	CreatedAt     int64        `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	CancelledAt   int64        `protobuf:"varint,10,opt,name=cancelledAt,proto3" json:"cancelledAt,omitempty"`
	FinishedAt    int64        `protobuf:"varint,11,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	IsExpired     bool         `protobuf:"varint,12,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
	User          *User        `protobuf:"bytes,30,opt,name=user,proto3" json:"user,omitempty"`
	OrderMethod   *OrderMethod `protobuf:"bytes,31,opt,name=orderMethod,proto3" json:"orderMethod,omitempty"`
	CanPay        bool         `protobuf:"varint,32,opt,name=canPay,proto3" json:"canPay,omitempty"` // 是否可以支付
	PayURL        string       `protobuf:"bytes,33,opt,name=payURL,proto3" json:"payURL,omitempty"`  // 构造后的支付URL
}

func (x *UserOrder) Reset() {
	*x = UserOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrder) ProtoMessage() {}

func (x *UserOrder) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrder.ProtoReflect.Descriptor instead.
func (*UserOrder) Descriptor() ([]byte, []int) {
	return file_models_model_user_order_proto_rawDescGZIP(), []int{0}
}

func (x *UserOrder) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserOrder) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserOrder) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserOrder) GetOrderMethodId() int64 {
	if x != nil {
		return x.OrderMethodId
	}
	return 0
}

func (x *UserOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserOrder) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UserOrder) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *UserOrder) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserOrder) GetCancelledAt() int64 {
	if x != nil {
		return x.CancelledAt
	}
	return 0
}

func (x *UserOrder) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *UserOrder) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

func (x *UserOrder) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserOrder) GetOrderMethod() *OrderMethod {
	if x != nil {
		return x.OrderMethod
	}
	return nil
}

func (x *UserOrder) GetCanPay() bool {
	if x != nil {
		return x.CanPay
	}
	return false
}

func (x *UserOrder) GetPayURL() string {
	if x != nil {
		return x.PayURL
	}
	return ""
}

var File_models_model_user_order_proto protoreflect.FileDescriptor

var file_models_model_user_order_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x03,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x79, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x79, 0x55,
	0x52, 0x4c, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x79, 0x55, 0x52, 0x4c,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_order_proto_rawDescOnce sync.Once
	file_models_model_user_order_proto_rawDescData = file_models_model_user_order_proto_rawDesc
)

func file_models_model_user_order_proto_rawDescGZIP() []byte {
	file_models_model_user_order_proto_rawDescOnce.Do(func() {
		file_models_model_user_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_order_proto_rawDescData)
	})
	return file_models_model_user_order_proto_rawDescData
}

var file_models_model_user_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_order_proto_goTypes = []interface{}{
	(*UserOrder)(nil),   // 0: pb.UserOrder
	(*User)(nil),        // 1: pb.User
	(*OrderMethod)(nil), // 2: pb.OrderMethod
}
var file_models_model_user_order_proto_depIdxs = []int32{
	1, // 0: pb.UserOrder.user:type_name -> pb.User
	2, // 1: pb.UserOrder.orderMethod:type_name -> pb.OrderMethod
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_user_order_proto_init() }
func file_models_model_user_order_proto_init() {
	if File_models_model_user_order_proto != nil {
		return
	}
	file_models_model_user_proto_init()
	file_models_model_order_method_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrder); i {
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
			RawDescriptor: file_models_model_user_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_order_proto_goTypes,
		DependencyIndexes: file_models_model_user_order_proto_depIdxs,
		MessageInfos:      file_models_model_user_order_proto_msgTypes,
	}.Build()
	File_models_model_user_order_proto = out.File
	file_models_model_user_order_proto_rawDesc = nil
	file_models_model_user_order_proto_goTypes = nil
	file_models_model_user_order_proto_depIdxs = nil
}
