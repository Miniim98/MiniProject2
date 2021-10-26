// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: Chat/Chat.proto

package Chat

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

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeded  bool              `protobuf:"varint,1,opt,name=succeded,proto3" json:"succeded,omitempty"`
	Timestamp *LamportTimeStamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionResponse) GetSucceded() bool {
	if x != nil {
		return x.Succeded
	}
	return false
}

func (x *ConnectionResponse) GetTimestamp() *LamportTimeStamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *LamportTimeStamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BroadcastRequest) GetTimestamp() *LamportTimeStamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeded  bool              `protobuf:"varint,1,opt,name=succeded,proto3" json:"succeded,omitempty"`
	Timestamp *LamportTimeStamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastResponse) GetSucceded() bool {
	if x != nil {
		return x.Succeded
	}
	return false
}

func (x *BroadcastResponse) GetTimestamp() *LamportTimeStamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *LamportTimeStamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{4}
}

func (x *PublishRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PublishRequest) GetTimestamp() *LamportTimeStamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeded  bool              `protobuf:"varint,1,opt,name=succeded,proto3" json:"succeded,omitempty"`
	Timestamp *LamportTimeStamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{5}
}

func (x *PublishResponse) GetSucceded() bool {
	if x != nil {
		return x.Succeded
	}
	return false
}

func (x *PublishResponse) GetTimestamp() *LamportTimeStamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type LamportTimeStamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events int32 `protobuf:"varint,1,opt,name=events,proto3" json:"events,omitempty"`
}

func (x *LamportTimeStamp) Reset() {
	*x = LamportTimeStamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_Chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LamportTimeStamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LamportTimeStamp) ProtoMessage() {}

func (x *LamportTimeStamp) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_Chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LamportTimeStamp.ProtoReflect.Descriptor instead.
func (*LamportTimeStamp) Descriptor() ([]byte, []int) {
	return file_Chat_Chat_proto_rawDescGZIP(), []int{6}
}

func (x *LamportTimeStamp) GetEvents() int32 {
	if x != nil {
		return x.Events
	}
	return 0
}

var File_Chat_Chat_proto protoreflect.FileDescriptor

var file_Chat_Chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x43, 0x68, 0x61, 0x74, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x63, 0x63, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x73, 0x75, 0x63, 0x63, 0x65, 0x64, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x62, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x65, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x64, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x60, 0x0a, 0x0e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x63, 0x0a,
	0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x63, 0x63, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x73, 0x75, 0x63, 0x63, 0x65, 0x64, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x2a, 0x0a, 0x10, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc6,
	0x01, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x12, 0x3e, 0x0a,
	0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x64, 0x6b,
	0x2f, 0x6a, 0x61, 0x72, 0x64, 0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Chat_Chat_proto_rawDescOnce sync.Once
	file_Chat_Chat_proto_rawDescData = file_Chat_Chat_proto_rawDesc
)

func file_Chat_Chat_proto_rawDescGZIP() []byte {
	file_Chat_Chat_proto_rawDescOnce.Do(func() {
		file_Chat_Chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_Chat_Chat_proto_rawDescData)
	})
	return file_Chat_Chat_proto_rawDescData
}

var file_Chat_Chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Chat_Chat_proto_goTypes = []interface{}{
	(*ConnectionRequest)(nil),  // 0: Chat.ConnectionRequest
	(*ConnectionResponse)(nil), // 1: Chat.ConnectionResponse
	(*BroadcastRequest)(nil),   // 2: Chat.BroadcastRequest
	(*BroadcastResponse)(nil),  // 3: Chat.BroadcastResponse
	(*PublishRequest)(nil),     // 4: Chat.PublishRequest
	(*PublishResponse)(nil),    // 5: Chat.PublishResponse
	(*LamportTimeStamp)(nil),   // 6: Chat.LamportTimeStamp
}
var file_Chat_Chat_proto_depIdxs = []int32{
	6, // 0: Chat.ConnectionResponse.timestamp:type_name -> Chat.LamportTimeStamp
	6, // 1: Chat.BroadcastRequest.timestamp:type_name -> Chat.LamportTimeStamp
	6, // 2: Chat.BroadcastResponse.timestamp:type_name -> Chat.LamportTimeStamp
	6, // 3: Chat.PublishRequest.timestamp:type_name -> Chat.LamportTimeStamp
	6, // 4: Chat.PublishResponse.timestamp:type_name -> Chat.LamportTimeStamp
	2, // 5: Chat.Chittychat.Broadcast:input_type -> Chat.BroadcastRequest
	4, // 6: Chat.Chittychat.Publish:input_type -> Chat.PublishRequest
	0, // 7: Chat.Chittychat.Connect:input_type -> Chat.ConnectionRequest
	3, // 8: Chat.Chittychat.Broadcast:output_type -> Chat.BroadcastResponse
	5, // 9: Chat.Chittychat.Publish:output_type -> Chat.PublishResponse
	1, // 10: Chat.Chittychat.Connect:output_type -> Chat.ConnectionResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Chat_Chat_proto_init() }
func file_Chat_Chat_proto_init() {
	if File_Chat_Chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Chat_Chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_Chat_Chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
		file_Chat_Chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
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
		file_Chat_Chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
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
		file_Chat_Chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_Chat_Chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
		file_Chat_Chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LamportTimeStamp); i {
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
			RawDescriptor: file_Chat_Chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Chat_Chat_proto_goTypes,
		DependencyIndexes: file_Chat_Chat_proto_depIdxs,
		MessageInfos:      file_Chat_Chat_proto_msgTypes,
	}.Build()
	File_Chat_Chat_proto = out.File
	file_Chat_Chat_proto_rawDesc = nil
	file_Chat_Chat_proto_goTypes = nil
	file_Chat_Chat_proto_depIdxs = nil
}