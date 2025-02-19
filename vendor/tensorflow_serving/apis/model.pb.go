// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.8.0
// source: tensorflow_serving/apis/model.proto

package tensorflow_serving

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Metadata for an inference request such as the model name and version.
type ModelSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required servable name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional choice of which version of the model to use.
	//
	// Recommended to be left unset in the common case. Should be specified only
	// when there is a strong version consistency requirement.
	//
	// When left unspecified, the system will serve the best available version.
	// This is typically the latest version, though during version transitions,
	// notably when serving on a fleet of instances, may be either the previous or
	// new version.
	//
	// Types that are assignable to VersionChoice:
	//	*ModelSpec_Version
	//	*ModelSpec_VersionLabel
	VersionChoice isModelSpec_VersionChoice `protobuf_oneof:"version_choice"`
	// A named signature to evaluate. If unspecified, the default signature will
	// be used.
	SignatureName string `protobuf:"bytes,3,opt,name=signature_name,json=signatureName,proto3" json:"signature_name,omitempty"`
}

func (x *ModelSpec) Reset() {
	*x = ModelSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelSpec) ProtoMessage() {}

func (x *ModelSpec) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelSpec.ProtoReflect.Descriptor instead.
func (*ModelSpec) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_model_proto_rawDescGZIP(), []int{0}
}

func (x *ModelSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ModelSpec) GetVersionChoice() isModelSpec_VersionChoice {
	if m != nil {
		return m.VersionChoice
	}
	return nil
}

func (x *ModelSpec) GetVersion() *wrappers.Int64Value {
	if x, ok := x.GetVersionChoice().(*ModelSpec_Version); ok {
		return x.Version
	}
	return nil
}

func (x *ModelSpec) GetVersionLabel() string {
	if x, ok := x.GetVersionChoice().(*ModelSpec_VersionLabel); ok {
		return x.VersionLabel
	}
	return ""
}

func (x *ModelSpec) GetSignatureName() string {
	if x != nil {
		return x.SignatureName
	}
	return ""
}

type isModelSpec_VersionChoice interface {
	isModelSpec_VersionChoice()
}

type ModelSpec_Version struct {
	// Use this specific version number.
	Version *wrappers.Int64Value `protobuf:"bytes,2,opt,name=version,proto3,oneof"`
}

type ModelSpec_VersionLabel struct {
	// Use the version associated with the given label.
	VersionLabel string `protobuf:"bytes,4,opt,name=version_label,json=versionLabel,proto3,oneof"`
}

func (*ModelSpec_Version) isModelSpec_VersionChoice() {}

func (*ModelSpec_VersionLabel) isModelSpec_VersionChoice() {}

var File_tensorflow_serving_apis_model_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_model_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x09, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x42, 0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_serving_apis_model_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_model_proto_rawDescData = file_tensorflow_serving_apis_model_proto_rawDesc
)

func file_tensorflow_serving_apis_model_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_model_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_model_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_model_proto_rawDescData
}

var file_tensorflow_serving_apis_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_apis_model_proto_goTypes = []interface{}{
	(*ModelSpec)(nil),           // 0: tensorflow.serving.ModelSpec
	(*wrappers.Int64Value)(nil), // 1: google.protobuf.Int64Value
}
var file_tensorflow_serving_apis_model_proto_depIdxs = []int32{
	1, // 0: tensorflow.serving.ModelSpec.version:type_name -> google.protobuf.Int64Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_model_proto_init() }
func file_tensorflow_serving_apis_model_proto_init() {
	if File_tensorflow_serving_apis_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelSpec); i {
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
	file_tensorflow_serving_apis_model_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ModelSpec_Version)(nil),
		(*ModelSpec_VersionLabel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_model_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_model_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_model_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_model_proto = out.File
	file_tensorflow_serving_apis_model_proto_rawDesc = nil
	file_tensorflow_serving_apis_model_proto_goTypes = nil
	file_tensorflow_serving_apis_model_proto_depIdxs = nil
}
