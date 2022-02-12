// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/models.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any "github.com/golang/protobuf/ptypes/any"
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

type DeletingType int32

const (
	DeletingType_CASCADE DeletingType = 0
	DeletingType_SET_NIL DeletingType = 1
)

// Enum value maps for DeletingType.
var (
	DeletingType_name = map[int32]string{
		0: "CASCADE",
		1: "SET_NIL",
	}
	DeletingType_value = map[string]int32{
		"CASCADE": 0,
		"SET_NIL": 1,
	}
)

func (x DeletingType) Enum() *DeletingType {
	p := new(DeletingType)
	*p = x
	return p
}

func (x DeletingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeletingType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_models_proto_enumTypes[0].Descriptor()
}

func (DeletingType) Type() protoreflect.EnumType {
	return &file_api_models_proto_enumTypes[0]
}

func (x DeletingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeletingType.Descriptor instead.
func (DeletingType) EnumDescriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{0}
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StringId        string   `protobuf:"bytes,2,opt,name=string_id,json=stringId,proto3" json:"string_id,omitempty"`
	ForeignObjectId string   `protobuf:"bytes,3,opt,name=foreign_object_id,json=foreignObjectId,proto3" json:"foreign_object_id,omitempty"`
	ActivityId      string   `protobuf:"bytes,4,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"` // only for reactions
	CreatedAt       int64    `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       int64    `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UserId          string   `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ActivityType    string   `protobuf:"bytes,8,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	ExtraData       *any.Any `protobuf:"bytes,9,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetStringId() string {
	if x != nil {
		return x.StringId
	}
	return ""
}

func (x *Activity) GetForeignObjectId() string {
	if x != nil {
		return x.ForeignObjectId
	}
	return ""
}

func (x *Activity) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *Activity) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Activity) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Activity) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Activity) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *Activity) GetExtraData() *any.Any {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ExtraData *any.Any `protobuf:"bytes,3,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{1}
}

func (x *Feed) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Feed) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Feed) GetExtraData() *any.Any {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type GroupingFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KeyFormat string   `protobuf:"bytes,3,opt,name=key_format,json=keyFormat,proto3" json:"key_format,omitempty"`
	ExtraData *any.Any `protobuf:"bytes,4,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (x *GroupingFeed) Reset() {
	*x = GroupingFeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupingFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupingFeed) ProtoMessage() {}

func (x *GroupingFeed) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupingFeed.ProtoReflect.Descriptor instead.
func (*GroupingFeed) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{2}
}

func (x *GroupingFeed) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupingFeed) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GroupingFeed) GetKeyFormat() string {
	if x != nil {
		return x.KeyFormat
	}
	return ""
}

func (x *GroupingFeed) GetExtraData() *any.Any {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeletingType DeletingType `protobuf:"varint,2,opt,name=deleting_type,json=deletingType,proto3,enum=syntoks_feed.DeletingType" json:"deleting_type,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{3}
}

func (x *Collection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Collection) GetDeletingType() DeletingType {
	if x != nil {
		return x.DeletingType
	}
	return DeletingType_CASCADE
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{4}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_models_proto protoreflect.FileDescriptor

var file_api_models_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x79, 0x6e, 0x74, 0x6f, 0x6b, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6d, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e,
	0x67, 0x46, 0x65, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0d, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x74, 0x6f, 0x6b, 0x73, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x2a, 0x28, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x53, 0x43, 0x41, 0x44, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x49, 0x4c, 0x10, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_models_proto_rawDescOnce sync.Once
	file_api_models_proto_rawDescData = file_api_models_proto_rawDesc
)

func file_api_models_proto_rawDescGZIP() []byte {
	file_api_models_proto_rawDescOnce.Do(func() {
		file_api_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_models_proto_rawDescData)
	})
	return file_api_models_proto_rawDescData
}

var file_api_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_models_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_models_proto_goTypes = []interface{}{
	(DeletingType)(0),    // 0: syntoks_feed.DeletingType
	(*Activity)(nil),     // 1: syntoks_feed.Activity
	(*Feed)(nil),         // 2: syntoks_feed.Feed
	(*GroupingFeed)(nil), // 3: syntoks_feed.GroupingFeed
	(*Collection)(nil),   // 4: syntoks_feed.Collection
	(*Object)(nil),       // 5: syntoks_feed.Object
	(*any.Any)(nil),      // 6: google.protobuf.Any
}
var file_api_models_proto_depIdxs = []int32{
	6, // 0: syntoks_feed.Activity.extra_data:type_name -> google.protobuf.Any
	6, // 1: syntoks_feed.Feed.extra_data:type_name -> google.protobuf.Any
	6, // 2: syntoks_feed.GroupingFeed.extra_data:type_name -> google.protobuf.Any
	0, // 3: syntoks_feed.Collection.deleting_type:type_name -> syntoks_feed.DeletingType
	6, // 4: syntoks_feed.Object.data:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_models_proto_init() }
func file_api_models_proto_init() {
	if File_api_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_api_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
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
		file_api_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupingFeed); i {
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
		file_api_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_api_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_api_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_models_proto_goTypes,
		DependencyIndexes: file_api_models_proto_depIdxs,
		EnumInfos:         file_api_models_proto_enumTypes,
		MessageInfos:      file_api_models_proto_msgTypes,
	}.Build()
	File_api_models_proto = out.File
	file_api_models_proto_rawDesc = nil
	file_api_models_proto_goTypes = nil
	file_api_models_proto_depIdxs = nil
}
