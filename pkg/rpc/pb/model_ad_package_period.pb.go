// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_ad_package_period.proto

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

// 高防实例有效期
type ADPackagePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn   bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Count  int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Unit   string `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Months int32  `protobuf:"varint,5,opt,name=months,proto3" json:"months,omitempty"`
}

func (x *ADPackagePeriod) Reset() {
	*x = ADPackagePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ad_package_period_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADPackagePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADPackagePeriod) ProtoMessage() {}

func (x *ADPackagePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ad_package_period_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADPackagePeriod.ProtoReflect.Descriptor instead.
func (*ADPackagePeriod) Descriptor() ([]byte, []int) {
	return file_models_model_ad_package_period_proto_rawDescGZIP(), []int{0}
}

func (x *ADPackagePeriod) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ADPackagePeriod) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *ADPackagePeriod) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ADPackagePeriod) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ADPackagePeriod) GetMonths() int32 {
	if x != nil {
		return x.Months
	}
	return 0
}

var File_models_model_ad_package_period_proto protoreflect.FileDescriptor

var file_models_model_ad_package_period_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x77, 0x0a, 0x0f, 0x41, 0x44,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_models_model_ad_package_period_proto_rawDescOnce sync.Once
	file_models_model_ad_package_period_proto_rawDescData = file_models_model_ad_package_period_proto_rawDesc
)

func file_models_model_ad_package_period_proto_rawDescGZIP() []byte {
	file_models_model_ad_package_period_proto_rawDescOnce.Do(func() {
		file_models_model_ad_package_period_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ad_package_period_proto_rawDescData)
	})
	return file_models_model_ad_package_period_proto_rawDescData
}

var file_models_model_ad_package_period_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ad_package_period_proto_goTypes = []interface{}{
	(*ADPackagePeriod)(nil), // 0: pb.ADPackagePeriod
}
var file_models_model_ad_package_period_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_ad_package_period_proto_init() }
func file_models_model_ad_package_period_proto_init() {
	if File_models_model_ad_package_period_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_ad_package_period_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADPackagePeriod); i {
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
			RawDescriptor: file_models_model_ad_package_period_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ad_package_period_proto_goTypes,
		DependencyIndexes: file_models_model_ad_package_period_proto_depIdxs,
		MessageInfos:      file_models_model_ad_package_period_proto_msgTypes,
	}.Build()
	File_models_model_ad_package_period_proto = out.File
	file_models_model_ad_package_period_proto_rawDesc = nil
	file_models_model_ad_package_period_proto_goTypes = nil
	file_models_model_ad_package_period_proto_depIdxs = nil
}
