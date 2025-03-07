// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: server/v1/demo.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type DemoCreateNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DemoCreateNameRequest) Reset() {
	*x = DemoCreateNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoCreateNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoCreateNameRequest) ProtoMessage() {}

func (x *DemoCreateNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoCreateNameRequest.ProtoReflect.Descriptor instead.
func (*DemoCreateNameRequest) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{0}
}

func (x *DemoCreateNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type DemoCreateNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DemoCreateNameResponse) Reset() {
	*x = DemoCreateNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoCreateNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoCreateNameResponse) ProtoMessage() {}

func (x *DemoCreateNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoCreateNameResponse.ProtoReflect.Descriptor instead.
func (*DemoCreateNameResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{1}
}

func (x *DemoCreateNameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DemoSearchNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DemoSearchNameRequest) Reset() {
	*x = DemoSearchNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoSearchNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoSearchNameRequest) ProtoMessage() {}

func (x *DemoSearchNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoSearchNameRequest.ProtoReflect.Descriptor instead.
func (*DemoSearchNameRequest) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{2}
}

func (x *DemoSearchNameRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DemoSearchNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool   `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DemoSearchNameResponse) Reset() {
	*x = DemoSearchNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoSearchNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoSearchNameResponse) ProtoMessage() {}

func (x *DemoSearchNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoSearchNameResponse.ProtoReflect.Descriptor instead.
func (*DemoSearchNameResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{3}
}

func (x *DemoSearchNameResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *DemoSearchNameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DemoMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DemoMessageRequest) Reset() {
	*x = DemoMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoMessageRequest) ProtoMessage() {}

func (x *DemoMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoMessageRequest.ProtoReflect.Descriptor instead.
func (*DemoMessageRequest) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{4}
}

func (x *DemoMessageRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DemoMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DemoMessageResponse) Reset() {
	*x = DemoMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoMessageResponse) ProtoMessage() {}

func (x *DemoMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoMessageResponse.ProtoReflect.Descriptor instead.
func (*DemoMessageResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{5}
}

func (x *DemoMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DemoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DemoMessage) Reset() {
	*x = DemoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoMessage) ProtoMessage() {}

func (x *DemoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoMessage.ProtoReflect.Descriptor instead.
func (*DemoMessage) Descriptor() ([]byte, []int) {
	return file_server_v1_demo_proto_rawDescGZIP(), []int{6}
}

func (x *DemoMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_server_v1_demo_proto protoreflect.FileDescriptor

var file_server_v1_demo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x15, 0x44, 0x65, 0x6d, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16,
	0x44, 0x65, 0x6d, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x44, 0x65, 0x6d,
	0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd5, 0x02,
	0x0a, 0x04, 0x64, 0x65, 0x6d, 0x6f, 0x12, 0x73, 0x0a, 0x0e, 0x44, 0x65, 0x6d, 0x6f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x44,
	0x65, 0x6d, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x63, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x64, 0x65, 0x6d, 0x6f,
	0x5f, 0x70, 0x75, 0x73, 0x68, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_v1_demo_proto_rawDescOnce sync.Once
	file_server_v1_demo_proto_rawDescData = file_server_v1_demo_proto_rawDesc
)

func file_server_v1_demo_proto_rawDescGZIP() []byte {
	file_server_v1_demo_proto_rawDescOnce.Do(func() {
		file_server_v1_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_v1_demo_proto_rawDescData)
	})
	return file_server_v1_demo_proto_rawDescData
}

var file_server_v1_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_server_v1_demo_proto_goTypes = []interface{}{
	(*DemoCreateNameRequest)(nil),  // 0: server.v1.DemoCreateNameRequest
	(*DemoCreateNameResponse)(nil), // 1: server.v1.DemoCreateNameResponse
	(*DemoSearchNameRequest)(nil),  // 2: server.v1.DemoSearchNameRequest
	(*DemoSearchNameResponse)(nil), // 3: server.v1.DemoSearchNameResponse
	(*DemoMessageRequest)(nil),     // 4: server.v1.DemoMessageRequest
	(*DemoMessageResponse)(nil),    // 5: server.v1.DemoMessageResponse
	(*DemoMessage)(nil),            // 6: server.v1.DemoMessage
}
var file_server_v1_demo_proto_depIdxs = []int32{
	0, // 0: server.v1.demo.DemoCreateName:input_type -> server.v1.DemoCreateNameRequest
	2, // 1: server.v1.demo.DemoSearchName:input_type -> server.v1.DemoSearchNameRequest
	4, // 2: server.v1.demo.DemoMessage:input_type -> server.v1.DemoMessageRequest
	1, // 3: server.v1.demo.DemoCreateName:output_type -> server.v1.DemoCreateNameResponse
	3, // 4: server.v1.demo.DemoSearchName:output_type -> server.v1.DemoSearchNameResponse
	5, // 5: server.v1.demo.DemoMessage:output_type -> server.v1.DemoMessageResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_v1_demo_proto_init() }
func file_server_v1_demo_proto_init() {
	if File_server_v1_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_v1_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoCreateNameRequest); i {
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
		file_server_v1_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoCreateNameResponse); i {
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
		file_server_v1_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoSearchNameRequest); i {
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
		file_server_v1_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoSearchNameResponse); i {
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
		file_server_v1_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoMessageRequest); i {
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
		file_server_v1_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoMessageResponse); i {
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
		file_server_v1_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoMessage); i {
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
			RawDescriptor: file_server_v1_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_v1_demo_proto_goTypes,
		DependencyIndexes: file_server_v1_demo_proto_depIdxs,
		MessageInfos:      file_server_v1_demo_proto_msgTypes,
	}.Build()
	File_server_v1_demo_proto = out.File
	file_server_v1_demo_proto_rawDesc = nil
	file_server_v1_demo_proto_goTypes = nil
	file_server_v1_demo_proto_depIdxs = nil
}
