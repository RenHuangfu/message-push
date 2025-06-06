// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: server/v1/message.proto

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

type ProcessMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	SendTime  int64 `protobuf:"varint,2,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	SendCount int32 `protobuf:"varint,3,opt,name=send_count,json=sendCount,proto3" json:"send_count,omitempty"`
	Delay     int64 `protobuf:"varint,4,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *ProcessMessageRequest) Reset() {
	*x = ProcessMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessMessageRequest) ProtoMessage() {}

func (x *ProcessMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessMessageRequest.ProtoReflect.Descriptor instead.
func (*ProcessMessageRequest) Descriptor() ([]byte, []int) {
	return file_server_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessMessageRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ProcessMessageRequest) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *ProcessMessageRequest) GetSendCount() int32 {
	if x != nil {
		return x.SendCount
	}
	return 0
}

func (x *ProcessMessageRequest) GetDelay() int64 {
	if x != nil {
		return x.Delay
	}
	return 0
}

type ProcessMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProcessMessageResponse) Reset() {
	*x = ProcessMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessMessageResponse) ProtoMessage() {}

func (x *ProcessMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessMessageResponse.ProtoReflect.Descriptor instead.
func (*ProcessMessageResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessMessageResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type DirectSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId    int32   `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ClientId []int32 `protobuf:"varint,2,rep,packed,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Content  string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DirectSendRequest) Reset() {
	*x = DirectSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectSendRequest) ProtoMessage() {}

func (x *DirectSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectSendRequest.ProtoReflect.Descriptor instead.
func (*DirectSendRequest) Descriptor() ([]byte, []int) {
	return file_server_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *DirectSendRequest) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DirectSendRequest) GetClientId() []int32 {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *DirectSendRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DirectSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DirectSendResponse) Reset() {
	*x = DirectSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectSendResponse) ProtoMessage() {}

func (x *DirectSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectSendResponse.ProtoReflect.Descriptor instead.
func (*DirectSendResponse) Descriptor() ([]byte, []int) {
	return file_server_v1_message_proto_rawDescGZIP(), []int{3}
}

var File_server_v1_message_proto protoreflect.FileDescriptor

var file_server_v1_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x80, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x30, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x61, 0x0a, 0x11, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x65, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x55, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5e, 0x0a, 0x11, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x0a, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_v1_message_proto_rawDescOnce sync.Once
	file_server_v1_message_proto_rawDescData = file_server_v1_message_proto_rawDesc
)

func file_server_v1_message_proto_rawDescGZIP() []byte {
	file_server_v1_message_proto_rawDescOnce.Do(func() {
		file_server_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_v1_message_proto_rawDescData)
	})
	return file_server_v1_message_proto_rawDescData
}

var file_server_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_v1_message_proto_goTypes = []interface{}{
	(*ProcessMessageRequest)(nil),  // 0: server.v1.ProcessMessageRequest
	(*ProcessMessageResponse)(nil), // 1: server.v1.ProcessMessageResponse
	(*DirectSendRequest)(nil),      // 2: server.v1.DirectSendRequest
	(*DirectSendResponse)(nil),     // 3: server.v1.DirectSendResponse
}
var file_server_v1_message_proto_depIdxs = []int32{
	0, // 0: server.v1.TriggerEvent.ProcessMessage:input_type -> server.v1.ProcessMessageRequest
	2, // 1: server.v1.DirectSendService.DirectSend:input_type -> server.v1.DirectSendRequest
	1, // 2: server.v1.TriggerEvent.ProcessMessage:output_type -> server.v1.ProcessMessageResponse
	3, // 3: server.v1.DirectSendService.DirectSend:output_type -> server.v1.DirectSendResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_v1_message_proto_init() }
func file_server_v1_message_proto_init() {
	if File_server_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessMessageRequest); i {
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
		file_server_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessMessageResponse); i {
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
		file_server_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectSendRequest); i {
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
		file_server_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectSendResponse); i {
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
			RawDescriptor: file_server_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_server_v1_message_proto_goTypes,
		DependencyIndexes: file_server_v1_message_proto_depIdxs,
		MessageInfos:      file_server_v1_message_proto_msgTypes,
	}.Build()
	File_server_v1_message_proto = out.File
	file_server_v1_message_proto_rawDesc = nil
	file_server_v1_message_proto_goTypes = nil
	file_server_v1_message_proto_depIdxs = nil
}
