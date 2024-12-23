// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: bootcamp_service/request/bootcamps.proto

package request

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

type GetBootcampsDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BootcampIds []string `protobuf:"bytes,1,rep,name=bootcamp_ids,json=bootcampIds,proto3" json:"bootcamp_ids,omitempty"`
}

func (x *GetBootcampsDetailsRequest) Reset() {
	*x = GetBootcampsDetailsRequest{}
	mi := &file_bootcamp_service_request_bootcamps_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBootcampsDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBootcampsDetailsRequest) ProtoMessage() {}

func (x *GetBootcampsDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_service_request_bootcamps_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBootcampsDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetBootcampsDetailsRequest) Descriptor() ([]byte, []int) {
	return file_bootcamp_service_request_bootcamps_proto_rawDescGZIP(), []int{0}
}

func (x *GetBootcampsDetailsRequest) GetBootcampIds() []string {
	if x != nil {
		return x.BootcampIds
	}
	return nil
}

type CreateBootcampRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Website     string   `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	Phone       string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Careers     []string `protobuf:"bytes,6,rep,name=careers,proto3" json:"careers,omitempty"`
	Address     string   `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateBootcampRequest) Reset() {
	*x = CreateBootcampRequest{}
	mi := &file_bootcamp_service_request_bootcamps_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBootcampRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBootcampRequest) ProtoMessage() {}

func (x *CreateBootcampRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_service_request_bootcamps_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBootcampRequest.ProtoReflect.Descriptor instead.
func (*CreateBootcampRequest) Descriptor() ([]byte, []int) {
	return file_bootcamp_service_request_bootcamps_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBootcampRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBootcampRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBootcampRequest) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *CreateBootcampRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateBootcampRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateBootcampRequest) GetCareers() []string {
	if x != nil {
		return x.Careers
	}
	return nil
}

func (x *CreateBootcampRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_bootcamp_service_request_bootcamps_proto protoreflect.FileDescriptor

var file_bootcamp_service_request_bootcamps_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x63,
	0x61, 0x6d, 0x70, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61,
	0x6d, 0x70, 0x49, 0x64, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x70, 0x61, 0x6e, 0x73, 0x68, 0x75, 0x63, 0x68, 0x61, 0x75, 0x62, 0x65, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x62,
	0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootcamp_service_request_bootcamps_proto_rawDescOnce sync.Once
	file_bootcamp_service_request_bootcamps_proto_rawDescData = file_bootcamp_service_request_bootcamps_proto_rawDesc
)

func file_bootcamp_service_request_bootcamps_proto_rawDescGZIP() []byte {
	file_bootcamp_service_request_bootcamps_proto_rawDescOnce.Do(func() {
		file_bootcamp_service_request_bootcamps_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootcamp_service_request_bootcamps_proto_rawDescData)
	})
	return file_bootcamp_service_request_bootcamps_proto_rawDescData
}

var file_bootcamp_service_request_bootcamps_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bootcamp_service_request_bootcamps_proto_goTypes = []any{
	(*GetBootcampsDetailsRequest)(nil), // 0: bootcamp_service.request.GetBootcampsDetailsRequest
	(*CreateBootcampRequest)(nil),      // 1: bootcamp_service.request.CreateBootcampRequest
}
var file_bootcamp_service_request_bootcamps_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bootcamp_service_request_bootcamps_proto_init() }
func file_bootcamp_service_request_bootcamps_proto_init() {
	if File_bootcamp_service_request_bootcamps_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bootcamp_service_request_bootcamps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bootcamp_service_request_bootcamps_proto_goTypes,
		DependencyIndexes: file_bootcamp_service_request_bootcamps_proto_depIdxs,
		MessageInfos:      file_bootcamp_service_request_bootcamps_proto_msgTypes,
	}.Build()
	File_bootcamp_service_request_bootcamps_proto = out.File
	file_bootcamp_service_request_bootcamps_proto_rawDesc = nil
	file_bootcamp_service_request_bootcamps_proto_goTypes = nil
	file_bootcamp_service_request_bootcamps_proto_depIdxs = nil
}
