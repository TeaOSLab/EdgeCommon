// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_plan.proto

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

// 创建套餐
type CreatePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClusterId          int64   `protobuf:"varint,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	TrafficLimitJSON   []byte  `protobuf:"bytes,3,opt,name=trafficLimitJSON,proto3" json:"trafficLimitJSON,omitempty"`
	FeaturesJSON       []byte  `protobuf:"bytes,4,opt,name=featuresJSON,proto3" json:"featuresJSON,omitempty"`
	PriceType          string  `protobuf:"bytes,5,opt,name=priceType,proto3" json:"priceType,omitempty"`
	TrafficPriceJSON   []byte  `protobuf:"bytes,6,opt,name=trafficPriceJSON,proto3" json:"trafficPriceJSON,omitempty"`
	BandwidthPriceJSON []byte  `protobuf:"bytes,10,opt,name=bandwidthPriceJSON,proto3" json:"bandwidthPriceJSON,omitempty"`
	MonthlyPrice       float32 `protobuf:"fixed32,7,opt,name=monthlyPrice,proto3" json:"monthlyPrice,omitempty"`
	SeasonallyPrice    float32 `protobuf:"fixed32,8,opt,name=seasonallyPrice,proto3" json:"seasonallyPrice,omitempty"`
	YearlyPrice        float32 `protobuf:"fixed32,9,opt,name=yearlyPrice,proto3" json:"yearlyPrice,omitempty"`
}

func (x *CreatePlanRequest) Reset() {
	*x = CreatePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanRequest) ProtoMessage() {}

func (x *CreatePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlanRequest.ProtoReflect.Descriptor instead.
func (*CreatePlanRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePlanRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *CreatePlanRequest) GetTrafficLimitJSON() []byte {
	if x != nil {
		return x.TrafficLimitJSON
	}
	return nil
}

func (x *CreatePlanRequest) GetFeaturesJSON() []byte {
	if x != nil {
		return x.FeaturesJSON
	}
	return nil
}

func (x *CreatePlanRequest) GetPriceType() string {
	if x != nil {
		return x.PriceType
	}
	return ""
}

func (x *CreatePlanRequest) GetTrafficPriceJSON() []byte {
	if x != nil {
		return x.TrafficPriceJSON
	}
	return nil
}

func (x *CreatePlanRequest) GetBandwidthPriceJSON() []byte {
	if x != nil {
		return x.BandwidthPriceJSON
	}
	return nil
}

func (x *CreatePlanRequest) GetMonthlyPrice() float32 {
	if x != nil {
		return x.MonthlyPrice
	}
	return 0
}

func (x *CreatePlanRequest) GetSeasonallyPrice() float32 {
	if x != nil {
		return x.SeasonallyPrice
	}
	return 0
}

func (x *CreatePlanRequest) GetYearlyPrice() float32 {
	if x != nil {
		return x.YearlyPrice
	}
	return 0
}

type CreatePlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId int64 `protobuf:"varint,1,opt,name=planId,proto3" json:"planId,omitempty"`
}

func (x *CreatePlanResponse) Reset() {
	*x = CreatePlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanResponse) ProtoMessage() {}

func (x *CreatePlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlanResponse.ProtoReflect.Descriptor instead.
func (*CreatePlanResponse) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlanResponse) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

// 修改套餐
type UpdatePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId             int64   `protobuf:"varint,1,opt,name=planId,proto3" json:"planId,omitempty"`
	Name               string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn               bool    `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
	ClusterId          int64   `protobuf:"varint,4,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	TrafficLimitJSON   []byte  `protobuf:"bytes,5,opt,name=trafficLimitJSON,proto3" json:"trafficLimitJSON,omitempty"`
	FeaturesJSON       []byte  `protobuf:"bytes,6,opt,name=featuresJSON,proto3" json:"featuresJSON,omitempty"`
	PriceType          string  `protobuf:"bytes,7,opt,name=priceType,proto3" json:"priceType,omitempty"`
	TrafficPriceJSON   []byte  `protobuf:"bytes,8,opt,name=trafficPriceJSON,proto3" json:"trafficPriceJSON,omitempty"`
	BandwidthPriceJSON []byte  `protobuf:"bytes,12,opt,name=bandwidthPriceJSON,proto3" json:"bandwidthPriceJSON,omitempty"`
	MonthlyPrice       float32 `protobuf:"fixed32,9,opt,name=monthlyPrice,proto3" json:"monthlyPrice,omitempty"`
	SeasonallyPrice    float32 `protobuf:"fixed32,10,opt,name=seasonallyPrice,proto3" json:"seasonallyPrice,omitempty"`
	YearlyPrice        float32 `protobuf:"fixed32,11,opt,name=yearlyPrice,proto3" json:"yearlyPrice,omitempty"`
}

func (x *UpdatePlanRequest) Reset() {
	*x = UpdatePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlanRequest) ProtoMessage() {}

func (x *UpdatePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlanRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlanRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePlanRequest) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *UpdatePlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePlanRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UpdatePlanRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *UpdatePlanRequest) GetTrafficLimitJSON() []byte {
	if x != nil {
		return x.TrafficLimitJSON
	}
	return nil
}

func (x *UpdatePlanRequest) GetFeaturesJSON() []byte {
	if x != nil {
		return x.FeaturesJSON
	}
	return nil
}

func (x *UpdatePlanRequest) GetPriceType() string {
	if x != nil {
		return x.PriceType
	}
	return ""
}

func (x *UpdatePlanRequest) GetTrafficPriceJSON() []byte {
	if x != nil {
		return x.TrafficPriceJSON
	}
	return nil
}

func (x *UpdatePlanRequest) GetBandwidthPriceJSON() []byte {
	if x != nil {
		return x.BandwidthPriceJSON
	}
	return nil
}

func (x *UpdatePlanRequest) GetMonthlyPrice() float32 {
	if x != nil {
		return x.MonthlyPrice
	}
	return 0
}

func (x *UpdatePlanRequest) GetSeasonallyPrice() float32 {
	if x != nil {
		return x.SeasonallyPrice
	}
	return 0
}

func (x *UpdatePlanRequest) GetYearlyPrice() float32 {
	if x != nil {
		return x.YearlyPrice
	}
	return 0
}

// 删除套餐
type DeletePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId int64 `protobuf:"varint,1,opt,name=planId,proto3" json:"planId,omitempty"`
}

func (x *DeletePlanRequest) Reset() {
	*x = DeletePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanRequest) ProtoMessage() {}

func (x *DeletePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlanRequest.ProtoReflect.Descriptor instead.
func (*DeletePlanRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePlanRequest) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

// 查找单个套餐
type FindEnabledPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId int64 `protobuf:"varint,1,opt,name=planId,proto3" json:"planId,omitempty"`
}

func (x *FindEnabledPlanRequest) Reset() {
	*x = FindEnabledPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledPlanRequest) ProtoMessage() {}

func (x *FindEnabledPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledPlanRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledPlanRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledPlanRequest) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type FindEnabledPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *FindEnabledPlanResponse) Reset() {
	*x = FindEnabledPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledPlanResponse) ProtoMessage() {}

func (x *FindEnabledPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledPlanResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledPlanResponse) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledPlanResponse) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

// 计算套餐数量
type CountAllEnabledPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountAllEnabledPlansRequest) Reset() {
	*x = CountAllEnabledPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledPlansRequest) ProtoMessage() {}

func (x *CountAllEnabledPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledPlansRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledPlansRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{6}
}

// 列出单页套餐
type ListEnabledPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListEnabledPlansRequest) Reset() {
	*x = ListEnabledPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledPlansRequest) ProtoMessage() {}

func (x *ListEnabledPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledPlansRequest.ProtoReflect.Descriptor instead.
func (*ListEnabledPlansRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{7}
}

func (x *ListEnabledPlansRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEnabledPlansRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListEnabledPlansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plans []*Plan `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
}

func (x *ListEnabledPlansResponse) Reset() {
	*x = ListEnabledPlansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledPlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledPlansResponse) ProtoMessage() {}

func (x *ListEnabledPlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledPlansResponse.ProtoReflect.Descriptor instead.
func (*ListEnabledPlansResponse) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{8}
}

func (x *ListEnabledPlansResponse) GetPlans() []*Plan {
	if x != nil {
		return x.Plans
	}
	return nil
}

// 对套餐进行排序
type SortPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanIds []int64 `protobuf:"varint,1,rep,packed,name=planIds,proto3" json:"planIds,omitempty"`
}

func (x *SortPlansRequest) Reset() {
	*x = SortPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_plan_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortPlansRequest) ProtoMessage() {}

func (x *SortPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_plan_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortPlansRequest.ProtoReflect.Descriptor instead.
func (*SortPlansRequest) Descriptor() ([]byte, []int) {
	return file_service_plan_proto_rawDescGZIP(), []int{9}
}

func (x *SortPlansRequest) GetPlanIds() []int64 {
	if x != nil {
		return x.PlanIds
	}
	return nil
}

var File_service_plan_proto protoreflect.FileDescriptor

var file_service_plan_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x02, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2e,
	0x0a, 0x12, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22,
	0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2c,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0xab, 0x03, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73,
	0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x2e, 0x0a, 0x12, 0x62, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x79, 0x65, 0x61, 0x72,
	0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x79,
	0x65, 0x61, 0x72, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x17, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c,
	0x61, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x45, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x05, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x49,
	0x64, 0x73, 0x32, 0xd1, 0x03, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4a, 0x0a, 0x0f, 0x66, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x14, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1f, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_plan_proto_rawDescOnce sync.Once
	file_service_plan_proto_rawDescData = file_service_plan_proto_rawDesc
)

func file_service_plan_proto_rawDescGZIP() []byte {
	file_service_plan_proto_rawDescOnce.Do(func() {
		file_service_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_plan_proto_rawDescData)
	})
	return file_service_plan_proto_rawDescData
}

var file_service_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_plan_proto_goTypes = []interface{}{
	(*CreatePlanRequest)(nil),           // 0: pb.CreatePlanRequest
	(*CreatePlanResponse)(nil),          // 1: pb.CreatePlanResponse
	(*UpdatePlanRequest)(nil),           // 2: pb.UpdatePlanRequest
	(*DeletePlanRequest)(nil),           // 3: pb.DeletePlanRequest
	(*FindEnabledPlanRequest)(nil),      // 4: pb.FindEnabledPlanRequest
	(*FindEnabledPlanResponse)(nil),     // 5: pb.FindEnabledPlanResponse
	(*CountAllEnabledPlansRequest)(nil), // 6: pb.CountAllEnabledPlansRequest
	(*ListEnabledPlansRequest)(nil),     // 7: pb.ListEnabledPlansRequest
	(*ListEnabledPlansResponse)(nil),    // 8: pb.ListEnabledPlansResponse
	(*SortPlansRequest)(nil),            // 9: pb.SortPlansRequest
	(*Plan)(nil),                        // 10: pb.Plan
	(*RPCSuccess)(nil),                  // 11: pb.RPCSuccess
	(*RPCCountResponse)(nil),            // 12: pb.RPCCountResponse
}
var file_service_plan_proto_depIdxs = []int32{
	10, // 0: pb.FindEnabledPlanResponse.plan:type_name -> pb.Plan
	10, // 1: pb.ListEnabledPlansResponse.plans:type_name -> pb.Plan
	0,  // 2: pb.PlanService.createPlan:input_type -> pb.CreatePlanRequest
	2,  // 3: pb.PlanService.updatePlan:input_type -> pb.UpdatePlanRequest
	3,  // 4: pb.PlanService.deletePlan:input_type -> pb.DeletePlanRequest
	4,  // 5: pb.PlanService.findEnabledPlan:input_type -> pb.FindEnabledPlanRequest
	6,  // 6: pb.PlanService.countAllEnabledPlans:input_type -> pb.CountAllEnabledPlansRequest
	7,  // 7: pb.PlanService.listEnabledPlans:input_type -> pb.ListEnabledPlansRequest
	9,  // 8: pb.PlanService.sortPlans:input_type -> pb.SortPlansRequest
	1,  // 9: pb.PlanService.createPlan:output_type -> pb.CreatePlanResponse
	11, // 10: pb.PlanService.updatePlan:output_type -> pb.RPCSuccess
	11, // 11: pb.PlanService.deletePlan:output_type -> pb.RPCSuccess
	5,  // 12: pb.PlanService.findEnabledPlan:output_type -> pb.FindEnabledPlanResponse
	12, // 13: pb.PlanService.countAllEnabledPlans:output_type -> pb.RPCCountResponse
	8,  // 14: pb.PlanService.listEnabledPlans:output_type -> pb.ListEnabledPlansResponse
	11, // 15: pb.PlanService.sortPlans:output_type -> pb.RPCSuccess
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_plan_proto_init() }
func file_service_plan_proto_init() {
	if File_service_plan_proto != nil {
		return
	}
	file_models_model_plan_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlanRequest); i {
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
		file_service_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlanResponse); i {
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
		file_service_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlanRequest); i {
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
		file_service_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlanRequest); i {
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
		file_service_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledPlanRequest); i {
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
		file_service_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledPlanResponse); i {
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
		file_service_plan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledPlansRequest); i {
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
		file_service_plan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledPlansRequest); i {
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
		file_service_plan_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledPlansResponse); i {
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
		file_service_plan_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortPlansRequest); i {
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
			RawDescriptor: file_service_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_plan_proto_goTypes,
		DependencyIndexes: file_service_plan_proto_depIdxs,
		MessageInfos:      file_service_plan_proto_msgTypes,
	}.Build()
	File_service_plan_proto = out.File
	file_service_plan_proto_rawDesc = nil
	file_service_plan_proto_goTypes = nil
	file_service_plan_proto_depIdxs = nil
}
