// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: gdpr.proto

package registered_v1

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

type DataGenerationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"` // namespace of user
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`       // user id
	UploadUrl string `protobuf:"bytes,3,opt,name=uploadUrl,proto3" json:"uploadUrl,omitempty"` // upload url for uploading the generated file into AccelByte GDPR Service storage
}

func (x *DataGenerationRequest) Reset() {
	*x = DataGenerationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gdpr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataGenerationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataGenerationRequest) ProtoMessage() {}

func (x *DataGenerationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gdpr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataGenerationRequest.ProtoReflect.Descriptor instead.
func (*DataGenerationRequest) Descriptor() ([]byte, []int) {
	return file_gdpr_proto_rawDescGZIP(), []int{0}
}

func (x *DataGenerationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DataGenerationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DataGenerationRequest) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

type DataGenerationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // indicate data generation was success
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // message from data generation process
}

func (x *DataGenerationResponse) Reset() {
	*x = DataGenerationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gdpr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataGenerationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataGenerationResponse) ProtoMessage() {}

func (x *DataGenerationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gdpr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataGenerationResponse.ProtoReflect.Descriptor instead.
func (*DataGenerationResponse) Descriptor() ([]byte, []int) {
	return file_gdpr_proto_rawDescGZIP(), []int{1}
}

func (x *DataGenerationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DataGenerationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DataDeletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"` // namespace of user
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`       // user id
}

func (x *DataDeletionRequest) Reset() {
	*x = DataDeletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gdpr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataDeletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDeletionRequest) ProtoMessage() {}

func (x *DataDeletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gdpr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDeletionRequest.ProtoReflect.Descriptor instead.
func (*DataDeletionRequest) Descriptor() ([]byte, []int) {
	return file_gdpr_proto_rawDescGZIP(), []int{2}
}

func (x *DataDeletionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DataDeletionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DataDeletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // indicate data deletion was success
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // message from data deletion process
}

func (x *DataDeletionResponse) Reset() {
	*x = DataDeletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gdpr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataDeletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataDeletionResponse) ProtoMessage() {}

func (x *DataDeletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gdpr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataDeletionResponse.ProtoReflect.Descriptor instead.
func (*DataDeletionResponse) Descriptor() ([]byte, []int) {
	return file_gdpr_proto_rawDescGZIP(), []int{3}
}

func (x *DataDeletionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DataDeletionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gdpr_proto protoreflect.FileDescriptor

var file_gdpr_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x63,
	0x63, 0x65, 0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x6b, 0x0a, 0x15, 0x44, 0x61,
	0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x4c, 0x0a, 0x16, 0x44, 0x61, 0x74, 0x61, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x4a, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfe,
	0x01, 0x0a, 0x04, 0x47, 0x44, 0x50, 0x52, 0x12, 0x7d, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x62, 0x79,
	0x74, 0x65, 0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x70, 0x0a, 0x20, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x62, 0x79, 0x74, 0x65,
	0x2e, 0x67, 0x64, 0x70, 0x72, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x67, 0x64, 0x70, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x20, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2f,
	0x67, 0x64, 0x70, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f,
	0x76, 0x31, 0xaa, 0x02, 0x1c, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x62, 0x79, 0x74, 0x65, 0x2e, 0x47,
	0x64, 0x70, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x2e, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gdpr_proto_rawDescOnce sync.Once
	file_gdpr_proto_rawDescData = file_gdpr_proto_rawDesc
)

func file_gdpr_proto_rawDescGZIP() []byte {
	file_gdpr_proto_rawDescOnce.Do(func() {
		file_gdpr_proto_rawDescData = protoimpl.X.CompressGZIP(file_gdpr_proto_rawDescData)
	})
	return file_gdpr_proto_rawDescData
}

var file_gdpr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gdpr_proto_goTypes = []interface{}{
	(*DataGenerationRequest)(nil),  // 0: accelbyte.gdpr.registered.v1.DataGenerationRequest
	(*DataGenerationResponse)(nil), // 1: accelbyte.gdpr.registered.v1.DataGenerationResponse
	(*DataDeletionRequest)(nil),    // 2: accelbyte.gdpr.registered.v1.DataDeletionRequest
	(*DataDeletionResponse)(nil),   // 3: accelbyte.gdpr.registered.v1.DataDeletionResponse
}
var file_gdpr_proto_depIdxs = []int32{
	0, // 0: accelbyte.gdpr.registered.v1.GDPR.DataGeneration:input_type -> accelbyte.gdpr.registered.v1.DataGenerationRequest
	2, // 1: accelbyte.gdpr.registered.v1.GDPR.DataDeletion:input_type -> accelbyte.gdpr.registered.v1.DataDeletionRequest
	1, // 2: accelbyte.gdpr.registered.v1.GDPR.DataGeneration:output_type -> accelbyte.gdpr.registered.v1.DataGenerationResponse
	3, // 3: accelbyte.gdpr.registered.v1.GDPR.DataDeletion:output_type -> accelbyte.gdpr.registered.v1.DataDeletionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gdpr_proto_init() }
func file_gdpr_proto_init() {
	if File_gdpr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gdpr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataGenerationRequest); i {
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
		file_gdpr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataGenerationResponse); i {
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
		file_gdpr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataDeletionRequest); i {
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
		file_gdpr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataDeletionResponse); i {
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
			RawDescriptor: file_gdpr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gdpr_proto_goTypes,
		DependencyIndexes: file_gdpr_proto_depIdxs,
		MessageInfos:      file_gdpr_proto_msgTypes,
	}.Build()
	File_gdpr_proto = out.File
	file_gdpr_proto_rawDesc = nil
	file_gdpr_proto_goTypes = nil
	file_gdpr_proto_depIdxs = nil
}