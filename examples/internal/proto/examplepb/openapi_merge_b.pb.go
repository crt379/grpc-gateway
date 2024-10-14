// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: examples/internal/proto/examplepb/openapi_merge_b.proto

// Merging Services
//
// This is an example of merging two proto files.

package examplepb

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

// InMessageB represents a message to ServiceB.
type InMessageB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Here is the explanation about InMessageB.values
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InMessageB) Reset() {
	*x = InMessageB{}
	mi := &file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InMessageB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMessageB) ProtoMessage() {}

func (x *InMessageB) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMessageB.ProtoReflect.Descriptor instead.
func (*InMessageB) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescGZIP(), []int{0}
}

func (x *InMessageB) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// OutMessageB represents a message returned from ServiceB.
type OutMessageB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Here is the explanation about OutMessageB.value
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *OutMessageB) Reset() {
	*x = OutMessageB{}
	mi := &file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutMessageB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutMessageB) ProtoMessage() {}

func (x *OutMessageB) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutMessageB.ProtoReflect.Descriptor instead.
func (*OutMessageB) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescGZIP(), []int{1}
}

func (x *OutMessageB) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_examples_internal_proto_examplepb_openapi_merge_b_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x5f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xb8, 0x02, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x12, 0x94, 0x01, 0x0a, 0x09, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x35, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x2f, 0x31,
	0x12, 0x94, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x77, 0x6f, 0x12, 0x35,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x34, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x49, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x62, 0x2f, 0x32, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescData = file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDesc
)

func file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDescData
}

var file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_internal_proto_examplepb_openapi_merge_b_proto_goTypes = []any{
	(*InMessageB)(nil),  // 0: grpc.gateway.examples.internal.examplepb.InMessageB
	(*OutMessageB)(nil), // 1: grpc.gateway.examples.internal.examplepb.OutMessageB
}
var file_examples_internal_proto_examplepb_openapi_merge_b_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.examples.internal.examplepb.ServiceB.MethodOne:input_type -> grpc.gateway.examples.internal.examplepb.InMessageB
	1, // 1: grpc.gateway.examples.internal.examplepb.ServiceB.MethodTwo:input_type -> grpc.gateway.examples.internal.examplepb.OutMessageB
	1, // 2: grpc.gateway.examples.internal.examplepb.ServiceB.MethodOne:output_type -> grpc.gateway.examples.internal.examplepb.OutMessageB
	0, // 3: grpc.gateway.examples.internal.examplepb.ServiceB.MethodTwo:output_type -> grpc.gateway.examples.internal.examplepb.InMessageB
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_openapi_merge_b_proto_init() }
func file_examples_internal_proto_examplepb_openapi_merge_b_proto_init() {
	if File_examples_internal_proto_examplepb_openapi_merge_b_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_openapi_merge_b_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_openapi_merge_b_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_openapi_merge_b_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_openapi_merge_b_proto = out.File
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_goTypes = nil
	file_examples_internal_proto_examplepb_openapi_merge_b_proto_depIdxs = nil
}
