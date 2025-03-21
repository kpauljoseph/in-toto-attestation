// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: in_toto_attestation/v1/resource_descriptor.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Proto representation of the in-toto v1 ResourceDescriptor.
// https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
// Validation of all fields is left to the users of this proto.
type ResourceDescriptor struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Name             string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uri              string                 `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest           map[string]string      `protobuf:"bytes,3,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Content          []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	DownloadLocation string                 `protobuf:"bytes,5,opt,name=download_location,json=downloadLocation,proto3" json:"download_location,omitempty"`
	MediaType        string                 `protobuf:"bytes,6,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// Per the Struct protobuf spec, this type corresponds to
	// a JSON Object, which is truly a map<string, Value> under the hood.
	// So, the Struct a) is still consistent with our specification for
	// the `annotations` field, and b) has native support in some language
	// bindings making their use easier in implementations.
	// See: https://pkg.go.dev/google.golang.org/protobuf/types/known/structpb#Struct
	Annotations   *structpb.Struct `protobuf:"bytes,7,opt,name=annotations,proto3" json:"annotations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceDescriptor) Reset() {
	*x = ResourceDescriptor{}
	mi := &file_in_toto_attestation_v1_resource_descriptor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDescriptor) ProtoMessage() {}

func (x *ResourceDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_v1_resource_descriptor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDescriptor.ProtoReflect.Descriptor instead.
func (*ResourceDescriptor) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_v1_resource_descriptor_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceDescriptor) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *ResourceDescriptor) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *ResourceDescriptor) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ResourceDescriptor) GetDownloadLocation() string {
	if x != nil {
		return x.DownloadLocation
	}
	return ""
}

func (x *ResourceDescriptor) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ResourceDescriptor) GetAnnotations() *structpb.Struct {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_in_toto_attestation_v1_resource_descriptor_proto protoreflect.FileDescriptor

var file_in_toto_attestation_v1_resource_descriptor_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x4e, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x47, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69,
	0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_in_toto_attestation_v1_resource_descriptor_proto_rawDescOnce sync.Once
	file_in_toto_attestation_v1_resource_descriptor_proto_rawDescData []byte
)

func file_in_toto_attestation_v1_resource_descriptor_proto_rawDescGZIP() []byte {
	file_in_toto_attestation_v1_resource_descriptor_proto_rawDescOnce.Do(func() {
		file_in_toto_attestation_v1_resource_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_in_toto_attestation_v1_resource_descriptor_proto_rawDesc), len(file_in_toto_attestation_v1_resource_descriptor_proto_rawDesc)))
	})
	return file_in_toto_attestation_v1_resource_descriptor_proto_rawDescData
}

var file_in_toto_attestation_v1_resource_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_in_toto_attestation_v1_resource_descriptor_proto_goTypes = []any{
	(*ResourceDescriptor)(nil), // 0: in_toto_attestation.v1.ResourceDescriptor
	nil,                        // 1: in_toto_attestation.v1.ResourceDescriptor.DigestEntry
	(*structpb.Struct)(nil),    // 2: google.protobuf.Struct
}
var file_in_toto_attestation_v1_resource_descriptor_proto_depIdxs = []int32{
	1, // 0: in_toto_attestation.v1.ResourceDescriptor.digest:type_name -> in_toto_attestation.v1.ResourceDescriptor.DigestEntry
	2, // 1: in_toto_attestation.v1.ResourceDescriptor.annotations:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_in_toto_attestation_v1_resource_descriptor_proto_init() }
func file_in_toto_attestation_v1_resource_descriptor_proto_init() {
	if File_in_toto_attestation_v1_resource_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_in_toto_attestation_v1_resource_descriptor_proto_rawDesc), len(file_in_toto_attestation_v1_resource_descriptor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_toto_attestation_v1_resource_descriptor_proto_goTypes,
		DependencyIndexes: file_in_toto_attestation_v1_resource_descriptor_proto_depIdxs,
		MessageInfos:      file_in_toto_attestation_v1_resource_descriptor_proto_msgTypes,
	}.Build()
	File_in_toto_attestation_v1_resource_descriptor_proto = out.File
	file_in_toto_attestation_v1_resource_descriptor_proto_goTypes = nil
	file_in_toto_attestation_v1_resource_descriptor_proto_depIdxs = nil
}
