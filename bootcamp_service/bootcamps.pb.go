// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: bootcamp_service/bootcamps.proto

package v1

import (
	request "github.com/dipanshuchaubey/protos-package/bootcamp_service/request"
	response "github.com/dipanshuchaubey/protos-package/bootcamp_service/response"
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

var File_bootcamp_service_bootcamps_proto protoreflect.FileDescriptor

var file_bootcamp_service_bootcamps_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x62, 0x6f, 0x6f,
	0x74, 0x63, 0x61, 0x6d, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8d, 0x02, 0x0a,
	0x0f, 0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x83, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x12, 0x2f, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61,
	0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x61, 0x6e,
	0x73, 0x68, 0x75, 0x63, 0x68, 0x61, 0x75, 0x62, 0x65, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61,
	0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bootcamp_service_bootcamps_proto_goTypes = []any{
	(*request.GetBootcampsDetailsRequest)(nil),   // 0: bootcamp_service.request.GetBootcampsDetailsRequest
	(*request.CreateBootcampRequest)(nil),        // 1: bootcamp_service.request.CreateBootcampRequest
	(*response.GetBootcampsDetailsResponse)(nil), // 2: bootcamp_service.response.GetBootcampsDetailsResponse
	(*response.CreateBootcampResponse)(nil),      // 3: bootcamp_service.response.CreateBootcampResponse
}
var file_bootcamp_service_bootcamps_proto_depIdxs = []int32{
	0, // 0: bootcamp_service.v1.BootcampService.GetBootcampsDetails:input_type -> bootcamp_service.request.GetBootcampsDetailsRequest
	1, // 1: bootcamp_service.v1.BootcampService.CreateBootcamp:input_type -> bootcamp_service.request.CreateBootcampRequest
	2, // 2: bootcamp_service.v1.BootcampService.GetBootcampsDetails:output_type -> bootcamp_service.response.GetBootcampsDetailsResponse
	3, // 3: bootcamp_service.v1.BootcampService.CreateBootcamp:output_type -> bootcamp_service.response.CreateBootcampResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bootcamp_service_bootcamps_proto_init() }
func file_bootcamp_service_bootcamps_proto_init() {
	if File_bootcamp_service_bootcamps_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bootcamp_service_bootcamps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bootcamp_service_bootcamps_proto_goTypes,
		DependencyIndexes: file_bootcamp_service_bootcamps_proto_depIdxs,
	}.Build()
	File_bootcamp_service_bootcamps_proto = out.File
	file_bootcamp_service_bootcamps_proto_rawDesc = nil
	file_bootcamp_service_bootcamps_proto_goTypes = nil
	file_bootcamp_service_bootcamps_proto_depIdxs = nil
}
