// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.28.3
// source: network/proto/network.proto

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

// Send init messages to peers
type NodeInitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeInitRequest) Reset() {
	*x = NodeInitRequest{}
	mi := &file_network_proto_network_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInitRequest) ProtoMessage() {}

func (x *NodeInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_network_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInitRequest.ProtoReflect.Descriptor instead.
func (*NodeInitRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_network_proto_rawDescGZIP(), []int{0}
}

func (x *NodeInitRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Response to init message
type NodeInitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeInitResponse) Reset() {
	*x = NodeInitResponse{}
	mi := &file_network_proto_network_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInitResponse) ProtoMessage() {}

func (x *NodeInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_network_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInitResponse.ProtoReflect.Descriptor instead.
func (*NodeInitResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_network_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInitResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_network_proto_network_proto protoreflect.FileDescriptor

var file_network_proto_network_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x3c, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x31, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_network_proto_network_proto_rawDescOnce sync.Once
	file_network_proto_network_proto_rawDescData = file_network_proto_network_proto_rawDesc
)

func file_network_proto_network_proto_rawDescGZIP() []byte {
	file_network_proto_network_proto_rawDescOnce.Do(func() {
		file_network_proto_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_proto_network_proto_rawDescData)
	})
	return file_network_proto_network_proto_rawDescData
}

var file_network_proto_network_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_network_proto_network_proto_goTypes = []any{
	(*NodeInitRequest)(nil),  // 0: NodeInitRequest
	(*NodeInitResponse)(nil), // 1: NodeInitResponse
}
var file_network_proto_network_proto_depIdxs = []int32{
	0, // 0: Network.NodeInit:input_type -> NodeInitRequest
	1, // 1: Network.NodeInit:output_type -> NodeInitResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_network_proto_network_proto_init() }
func file_network_proto_network_proto_init() {
	if File_network_proto_network_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_proto_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_proto_network_proto_goTypes,
		DependencyIndexes: file_network_proto_network_proto_depIdxs,
		MessageInfos:      file_network_proto_network_proto_msgTypes,
	}.Build()
	File_network_proto_network_proto = out.File
	file_network_proto_network_proto_rawDesc = nil
	file_network_proto_network_proto_goTypes = nil
	file_network_proto_network_proto_depIdxs = nil
}
