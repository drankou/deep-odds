// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.8.0
// source: tensorflow_serving/apis/prediction_service.proto

package tensorflow_serving

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_tensorflow_serving_apis_prediction_service_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_prediction_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x67, 0x0a,
	0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x22, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_tensorflow_serving_apis_prediction_service_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),  // 0: tensorflow.serving.PredictRequest
	(*PredictResponse)(nil), // 1: tensorflow.serving.PredictResponse
}
var file_tensorflow_serving_apis_prediction_service_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.PredictionService.Predict:input_type -> tensorflow.serving.PredictRequest
	1, // 1: tensorflow.serving.PredictionService.Predict:output_type -> tensorflow.serving.PredictResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_prediction_service_proto_init() }
func file_tensorflow_serving_apis_prediction_service_proto_init() {
	if File_tensorflow_serving_apis_prediction_service_proto != nil {
		return
	}
	file_tensorflow_serving_apis_predict_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_prediction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_serving_apis_prediction_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_prediction_service_proto_depIdxs,
	}.Build()
	File_tensorflow_serving_apis_prediction_service_proto = out.File
	file_tensorflow_serving_apis_prediction_service_proto_rawDesc = nil
	file_tensorflow_serving_apis_prediction_service_proto_goTypes = nil
	file_tensorflow_serving_apis_prediction_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Predict -- provides access to loaded TensorFlow model.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Predict -- provides access to loaded TensorFlow model.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/prediction_service.proto",
}
