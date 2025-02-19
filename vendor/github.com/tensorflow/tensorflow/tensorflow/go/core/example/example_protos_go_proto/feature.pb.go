// Protocol messages for describing features for machine learning model
// training or inference.
//
// There are three base Feature types:
//   - bytes
//   - float
//   - int64
//
// A Feature contains Lists which may hold zero or more values.  These
// lists are the base values BytesList, FloatList, Int64List.
//
// Features are organized into categories by name.  The Features message
// contains the mapping from name to Feature.
//
// Example Features for a movie recommendation application:
//   feature {
//     key: "age"
//     value { float_list {
//       value: 29.0
//     }}
//   }
//   feature {
//     key: "movie"
//     value { bytes_list {
//       value: "The Shawshank Redemption"
//       value: "Fight Club"
//     }}
//   }
//   feature {
//     key: "movie_ratings"
//     value { float_list {
//       value: 9.0
//       value: 9.7
//     }}
//   }
//   feature {
//     key: "suggestion"
//     value { bytes_list {
//       value: "Inception"
//     }}
//   }
//   feature {
//     key: "suggestion_purchased"
//     value { int64_list {
//       value: 1
//     }}
//   }
//   feature {
//     key: "purchase_price"
//     value { float_list {
//       value: 9.99
//     }}
//   }
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.8.0
// source: tensorflow/core/example/feature.proto

package example_protos_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

// Containers to hold repeated fundamental values.
type BytesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *BytesList) Reset() {
	*x = BytesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesList) ProtoMessage() {}

func (x *BytesList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesList.ProtoReflect.Descriptor instead.
func (*BytesList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{0}
}

func (x *BytesList) GetValue() [][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type FloatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *FloatList) Reset() {
	*x = FloatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatList) ProtoMessage() {}

func (x *FloatList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatList.ProtoReflect.Descriptor instead.
func (*FloatList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{1}
}

func (x *FloatList) GetValue() []float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int64List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int64 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Int64List) Reset() {
	*x = Int64List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64List) ProtoMessage() {}

func (x *Int64List) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64List.ProtoReflect.Descriptor instead.
func (*Int64List) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{2}
}

func (x *Int64List) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

// Containers for non-sequential data.
type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each feature can be exactly one kind.
	//
	// Types that are assignable to Kind:
	//	*Feature_BytesList
	//	*Feature_FloatList
	//	*Feature_Int64List
	Kind isFeature_Kind `protobuf_oneof:"kind"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{3}
}

func (m *Feature) GetKind() isFeature_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Feature) GetBytesList() *BytesList {
	if x, ok := x.GetKind().(*Feature_BytesList); ok {
		return x.BytesList
	}
	return nil
}

func (x *Feature) GetFloatList() *FloatList {
	if x, ok := x.GetKind().(*Feature_FloatList); ok {
		return x.FloatList
	}
	return nil
}

func (x *Feature) GetInt64List() *Int64List {
	if x, ok := x.GetKind().(*Feature_Int64List); ok {
		return x.Int64List
	}
	return nil
}

type isFeature_Kind interface {
	isFeature_Kind()
}

type Feature_BytesList struct {
	BytesList *BytesList `protobuf:"bytes,1,opt,name=bytes_list,json=bytesList,proto3,oneof"`
}

type Feature_FloatList struct {
	FloatList *FloatList `protobuf:"bytes,2,opt,name=float_list,json=floatList,proto3,oneof"`
}

type Feature_Int64List struct {
	Int64List *Int64List `protobuf:"bytes,3,opt,name=int64_list,json=int64List,proto3,oneof"`
}

func (*Feature_BytesList) isFeature_Kind() {}

func (*Feature_FloatList) isFeature_Kind() {}

func (*Feature_Int64List) isFeature_Kind() {}

type Features struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from feature name to feature.
	Feature map[string]*Feature `protobuf:"bytes,1,rep,name=feature,proto3" json:"feature,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Features) Reset() {
	*x = Features{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Features) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Features) ProtoMessage() {}

func (x *Features) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Features.ProtoReflect.Descriptor instead.
func (*Features) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{4}
}

func (x *Features) GetFeature() map[string]*Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

// Containers for sequential data.
//
// A FeatureList contains lists of Features.  These may hold zero or more
// Feature values.
//
// FeatureLists are organized into categories by name.  The FeatureLists message
// contains the mapping from name to FeatureList.
//
type FeatureList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature []*Feature `protobuf:"bytes,1,rep,name=feature,proto3" json:"feature,omitempty"`
}

func (x *FeatureList) Reset() {
	*x = FeatureList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureList) ProtoMessage() {}

func (x *FeatureList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureList.ProtoReflect.Descriptor instead.
func (*FeatureList) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{5}
}

func (x *FeatureList) GetFeature() []*Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type FeatureLists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from feature name to feature list.
	FeatureList map[string]*FeatureList `protobuf:"bytes,1,rep,name=feature_list,json=featureList,proto3" json:"feature_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FeatureLists) Reset() {
	*x = FeatureLists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_feature_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureLists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureLists) ProtoMessage() {}

func (x *FeatureLists) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_feature_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureLists.ProtoReflect.Descriptor instead.
func (*FeatureLists) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_feature_proto_rawDescGZIP(), []int{6}
}

func (x *FeatureLists) GetFeatureList() map[string]*FeatureList {
	if x != nil {
		return x.FeatureList
	}
	return nil
}

var File_tensorflow_core_example_feature_proto protoreflect.FileDescriptor

var file_tensorflow_core_example_feature_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x21, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x02, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a,
	0x09, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x22, 0x98, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x4f, 0x0a, 0x0c, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3c, 0x0a, 0x0b, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x57, 0x0a, 0x10, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x81, 0x01, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0d, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_example_feature_proto_rawDescOnce sync.Once
	file_tensorflow_core_example_feature_proto_rawDescData = file_tensorflow_core_example_feature_proto_rawDesc
)

func file_tensorflow_core_example_feature_proto_rawDescGZIP() []byte {
	file_tensorflow_core_example_feature_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_example_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_example_feature_proto_rawDescData)
	})
	return file_tensorflow_core_example_feature_proto_rawDescData
}

var file_tensorflow_core_example_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tensorflow_core_example_feature_proto_goTypes = []interface{}{
	(*BytesList)(nil),    // 0: tensorflow.BytesList
	(*FloatList)(nil),    // 1: tensorflow.FloatList
	(*Int64List)(nil),    // 2: tensorflow.Int64List
	(*Feature)(nil),      // 3: tensorflow.Feature
	(*Features)(nil),     // 4: tensorflow.Features
	(*FeatureList)(nil),  // 5: tensorflow.FeatureList
	(*FeatureLists)(nil), // 6: tensorflow.FeatureLists
	nil,                  // 7: tensorflow.Features.FeatureEntry
	nil,                  // 8: tensorflow.FeatureLists.FeatureListEntry
}
var file_tensorflow_core_example_feature_proto_depIdxs = []int32{
	0, // 0: tensorflow.Feature.bytes_list:type_name -> tensorflow.BytesList
	1, // 1: tensorflow.Feature.float_list:type_name -> tensorflow.FloatList
	2, // 2: tensorflow.Feature.int64_list:type_name -> tensorflow.Int64List
	7, // 3: tensorflow.Features.feature:type_name -> tensorflow.Features.FeatureEntry
	3, // 4: tensorflow.FeatureList.feature:type_name -> tensorflow.Feature
	8, // 5: tensorflow.FeatureLists.feature_list:type_name -> tensorflow.FeatureLists.FeatureListEntry
	3, // 6: tensorflow.Features.FeatureEntry.value:type_name -> tensorflow.Feature
	5, // 7: tensorflow.FeatureLists.FeatureListEntry.value:type_name -> tensorflow.FeatureList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_core_example_feature_proto_init() }
func file_tensorflow_core_example_feature_proto_init() {
	if File_tensorflow_core_example_feature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_example_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesList); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatList); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64List); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Features); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureList); i {
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
		file_tensorflow_core_example_feature_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureLists); i {
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
	file_tensorflow_core_example_feature_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Feature_BytesList)(nil),
		(*Feature_FloatList)(nil),
		(*Feature_Int64List)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_example_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_example_feature_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_example_feature_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_example_feature_proto_msgTypes,
	}.Build()
	File_tensorflow_core_example_feature_proto = out.File
	file_tensorflow_core_example_feature_proto_rawDesc = nil
	file_tensorflow_core_example_feature_proto_goTypes = nil
	file_tensorflow_core_example_feature_proto_depIdxs = nil
}
