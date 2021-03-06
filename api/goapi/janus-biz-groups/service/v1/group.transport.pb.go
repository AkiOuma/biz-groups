// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: group.transport.proto

package v1

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

type FindGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Querier *GroupQuerier `protobuf:"bytes,1,opt,name=querier,proto3" json:"querier,omitempty"`
}

func (x *FindGroupsRequest) Reset() {
	*x = FindGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupsRequest) ProtoMessage() {}

func (x *FindGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupsRequest.ProtoReflect.Descriptor instead.
func (*FindGroupsRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{0}
}

func (x *FindGroupsRequest) GetQuerier() *GroupQuerier {
	if x != nil {
		return x.Querier
	}
	return nil
}

type FindGroupsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Data  []*Group `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindGroupsReply) Reset() {
	*x = FindGroupsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupsReply) ProtoMessage() {}

func (x *FindGroupsReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupsReply.ProtoReflect.Descriptor instead.
func (*FindGroupsReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{1}
}

func (x *FindGroupsReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FindGroupsReply) GetData() []*Group {
	if x != nil {
		return x.Data
	}
	return nil
}

type SaveGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Group `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveGroupsRequest) Reset() {
	*x = SaveGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveGroupsRequest) ProtoMessage() {}

func (x *SaveGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveGroupsRequest.ProtoReflect.Descriptor instead.
func (*SaveGroupsRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{2}
}

func (x *SaveGroupsRequest) GetData() []*Group {
	if x != nil {
		return x.Data
	}
	return nil
}

type SaveGroupsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Group `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveGroupsReply) Reset() {
	*x = SaveGroupsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveGroupsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveGroupsReply) ProtoMessage() {}

func (x *SaveGroupsReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveGroupsReply.ProtoReflect.Descriptor instead.
func (*SaveGroupsReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{3}
}

func (x *SaveGroupsReply) GetData() []*Group {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Querier *GroupQuerier `protobuf:"bytes,1,opt,name=querier,proto3" json:"querier,omitempty"`
}

func (x *RemoveGroupsRequest) Reset() {
	*x = RemoveGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGroupsRequest) ProtoMessage() {}

func (x *RemoveGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGroupsRequest.ProtoReflect.Descriptor instead.
func (*RemoveGroupsRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveGroupsRequest) GetQuerier() *GroupQuerier {
	if x != nil {
		return x.Querier
	}
	return nil
}

type RemoveGroupsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveGroupsReply) Reset() {
	*x = RemoveGroupsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGroupsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGroupsReply) ProtoMessage() {}

func (x *RemoveGroupsReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGroupsReply.ProtoReflect.Descriptor instead.
func (*RemoveGroupsReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{5}
}

type FindGroupMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int32 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *FindGroupMemberRequest) Reset() {
	*x = FindGroupMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupMemberRequest) ProtoMessage() {}

func (x *FindGroupMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupMemberRequest.ProtoReflect.Descriptor instead.
func (*FindGroupMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{6}
}

func (x *FindGroupMemberRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type FindGroupMemberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []int32 `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
}

func (x *FindGroupMemberReply) Reset() {
	*x = FindGroupMemberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupMemberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupMemberReply) ProtoMessage() {}

func (x *FindGroupMemberReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupMemberReply.ProtoReflect.Descriptor instead.
func (*FindGroupMemberReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{7}
}

func (x *FindGroupMemberReply) GetData() []int32 {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddGroupMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GroupMember `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddGroupMemberRequest) Reset() {
	*x = AddGroupMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupMemberRequest) ProtoMessage() {}

func (x *AddGroupMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupMemberRequest.ProtoReflect.Descriptor instead.
func (*AddGroupMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{8}
}

func (x *AddGroupMemberRequest) GetData() *GroupMember {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddGroupMemberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GroupMember `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddGroupMemberReply) Reset() {
	*x = AddGroupMemberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupMemberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupMemberReply) ProtoMessage() {}

func (x *AddGroupMemberReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupMemberReply.ProtoReflect.Descriptor instead.
func (*AddGroupMemberReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{9}
}

func (x *AddGroupMemberReply) GetData() *GroupMember {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveGroupMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GroupMember `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RemoveGroupMemberRequest) Reset() {
	*x = RemoveGroupMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGroupMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGroupMemberRequest) ProtoMessage() {}

func (x *RemoveGroupMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGroupMemberRequest.ProtoReflect.Descriptor instead.
func (*RemoveGroupMemberRequest) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveGroupMemberRequest) GetData() *GroupMember {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveGroupMemberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GroupMember `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RemoveGroupMemberReply) Reset() {
	*x = RemoveGroupMemberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_transport_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGroupMemberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGroupMemberReply) ProtoMessage() {}

func (x *RemoveGroupMemberReply) ProtoReflect() protoreflect.Message {
	mi := &file_group_transport_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGroupMemberReply.ProtoReflect.Descriptor instead.
func (*RemoveGroupMemberReply) Descriptor() ([]byte, []int) {
	return file_group_transport_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveGroupMemberReply) GetData() *GroupMember {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_group_transport_proto protoreflect.FileDescriptor

var file_group_transport_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72,
	0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x0f, 0x46, 0x69, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x11, 0x53, 0x61, 0x76,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x49, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x32, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x44, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x18, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x45, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x89, 0x04, 0x0a, 0x09, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_transport_proto_rawDescOnce sync.Once
	file_group_transport_proto_rawDescData = file_group_transport_proto_rawDesc
)

func file_group_transport_proto_rawDescGZIP() []byte {
	file_group_transport_proto_rawDescOnce.Do(func() {
		file_group_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_transport_proto_rawDescData)
	})
	return file_group_transport_proto_rawDescData
}

var file_group_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_group_transport_proto_goTypes = []interface{}{
	(*FindGroupsRequest)(nil),        // 0: service.v1.FindGroupsRequest
	(*FindGroupsReply)(nil),          // 1: service.v1.FindGroupsReply
	(*SaveGroupsRequest)(nil),        // 2: service.v1.SaveGroupsRequest
	(*SaveGroupsReply)(nil),          // 3: service.v1.SaveGroupsReply
	(*RemoveGroupsRequest)(nil),      // 4: service.v1.RemoveGroupsRequest
	(*RemoveGroupsReply)(nil),        // 5: service.v1.RemoveGroupsReply
	(*FindGroupMemberRequest)(nil),   // 6: service.v1.FindGroupMemberRequest
	(*FindGroupMemberReply)(nil),     // 7: service.v1.FindGroupMemberReply
	(*AddGroupMemberRequest)(nil),    // 8: service.v1.AddGroupMemberRequest
	(*AddGroupMemberReply)(nil),      // 9: service.v1.AddGroupMemberReply
	(*RemoveGroupMemberRequest)(nil), // 10: service.v1.RemoveGroupMemberRequest
	(*RemoveGroupMemberReply)(nil),   // 11: service.v1.RemoveGroupMemberReply
	(*GroupQuerier)(nil),             // 12: service.v1.GroupQuerier
	(*Group)(nil),                    // 13: service.v1.Group
	(*GroupMember)(nil),              // 14: service.v1.GroupMember
}
var file_group_transport_proto_depIdxs = []int32{
	12, // 0: service.v1.FindGroupsRequest.querier:type_name -> service.v1.GroupQuerier
	13, // 1: service.v1.FindGroupsReply.data:type_name -> service.v1.Group
	13, // 2: service.v1.SaveGroupsRequest.data:type_name -> service.v1.Group
	13, // 3: service.v1.SaveGroupsReply.data:type_name -> service.v1.Group
	12, // 4: service.v1.RemoveGroupsRequest.querier:type_name -> service.v1.GroupQuerier
	14, // 5: service.v1.AddGroupMemberRequest.data:type_name -> service.v1.GroupMember
	14, // 6: service.v1.AddGroupMemberReply.data:type_name -> service.v1.GroupMember
	14, // 7: service.v1.RemoveGroupMemberRequest.data:type_name -> service.v1.GroupMember
	14, // 8: service.v1.RemoveGroupMemberReply.data:type_name -> service.v1.GroupMember
	0,  // 9: service.v1.Transport.FindGroups:input_type -> service.v1.FindGroupsRequest
	2,  // 10: service.v1.Transport.SaveGroups:input_type -> service.v1.SaveGroupsRequest
	4,  // 11: service.v1.Transport.RemoveGroups:input_type -> service.v1.RemoveGroupsRequest
	6,  // 12: service.v1.Transport.FindGroupMember:input_type -> service.v1.FindGroupMemberRequest
	8,  // 13: service.v1.Transport.AddGroupMember:input_type -> service.v1.AddGroupMemberRequest
	10, // 14: service.v1.Transport.RemoveGroupMember:input_type -> service.v1.RemoveGroupMemberRequest
	1,  // 15: service.v1.Transport.FindGroups:output_type -> service.v1.FindGroupsReply
	3,  // 16: service.v1.Transport.SaveGroups:output_type -> service.v1.SaveGroupsReply
	5,  // 17: service.v1.Transport.RemoveGroups:output_type -> service.v1.RemoveGroupsReply
	7,  // 18: service.v1.Transport.FindGroupMember:output_type -> service.v1.FindGroupMemberReply
	9,  // 19: service.v1.Transport.AddGroupMember:output_type -> service.v1.AddGroupMemberReply
	11, // 20: service.v1.Transport.RemoveGroupMember:output_type -> service.v1.RemoveGroupMemberReply
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_group_transport_proto_init() }
func file_group_transport_proto_init() {
	if File_group_transport_proto != nil {
		return
	}
	file_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupsRequest); i {
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
		file_group_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupsReply); i {
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
		file_group_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveGroupsRequest); i {
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
		file_group_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveGroupsReply); i {
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
		file_group_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGroupsRequest); i {
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
		file_group_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGroupsReply); i {
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
		file_group_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupMemberRequest); i {
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
		file_group_transport_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupMemberReply); i {
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
		file_group_transport_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupMemberRequest); i {
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
		file_group_transport_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupMemberReply); i {
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
		file_group_transport_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGroupMemberRequest); i {
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
		file_group_transport_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGroupMemberReply); i {
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
			RawDescriptor: file_group_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_transport_proto_goTypes,
		DependencyIndexes: file_group_transport_proto_depIdxs,
		MessageInfos:      file_group_transport_proto_msgTypes,
	}.Build()
	File_group_transport_proto = out.File
	file_group_transport_proto_rawDesc = nil
	file_group_transport_proto_goTypes = nil
	file_group_transport_proto_depIdxs = nil
}
