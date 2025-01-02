// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: messages/proto/messages.proto

package __

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

// Message to be shared with the neighbours
type MessageSharing struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageSharing) Reset() {
	*x = MessageSharing{}
	mi := &file_messages_proto_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSharing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSharing) ProtoMessage() {}

func (x *MessageSharing) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSharing.ProtoReflect.Descriptor instead.
func (*MessageSharing) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSharing) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Response message
type MessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_messages_proto_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// Message to be shared with the neighbours
type MessageCommitSharing struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageCommitSharing) Reset() {
	*x = MessageCommitSharing{}
	mi := &file_messages_proto_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageCommitSharing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCommitSharing) ProtoMessage() {}

func (x *MessageCommitSharing) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCommitSharing.ProtoReflect.Descriptor instead.
func (*MessageCommitSharing) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{2}
}

func (x *MessageCommitSharing) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Response message
type MessageCommitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageCommitResponse) Reset() {
	*x = MessageCommitResponse{}
	mi := &file_messages_proto_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageCommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCommitResponse) ProtoMessage() {}

func (x *MessageCommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCommitResponse.ProtoReflect.Descriptor instead.
func (*MessageCommitResponse) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{3}
}

func (x *MessageCommitResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_messages_proto_messages_proto protoreflect.FileDescriptor

var file_messages_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2a, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x8d, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x10, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x16, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_messages_proto_rawDescOnce sync.Once
	file_messages_proto_messages_proto_rawDescData = file_messages_proto_messages_proto_rawDesc
)

func file_messages_proto_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_messages_proto_rawDescData)
	})
	return file_messages_proto_messages_proto_rawDescData
}

var file_messages_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_messages_proto_goTypes = []any{
	(*MessageSharing)(nil),        // 0: MessageSharing
	(*MessageResponse)(nil),       // 1: MessageResponse
	(*MessageCommitSharing)(nil),  // 2: MessageCommitSharing
	(*MessageCommitResponse)(nil), // 3: MessageCommitResponse
}
var file_messages_proto_messages_proto_depIdxs = []int32{
	0, // 0: Message.MessageMulticast:input_type -> MessageSharing
	2, // 1: Message.MessageCommitMulticast:input_type -> MessageCommitSharing
	1, // 2: Message.MessageMulticast:output_type -> MessageResponse
	3, // 3: Message.MessageCommitMulticast:output_type -> MessageCommitResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_messages_proto_init() }
func file_messages_proto_messages_proto_init() {
	if File_messages_proto_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_messages_proto_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_messages_proto_msgTypes,
	}.Build()
	File_messages_proto_messages_proto = out.File
	file_messages_proto_messages_proto_rawDesc = nil
	file_messages_proto_messages_proto_goTypes = nil
	file_messages_proto_messages_proto_depIdxs = nil
}
