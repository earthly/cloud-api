// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: api/public/logs/api.proto

package logs

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

type UploadLogBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogID       string `protobuf:"bytes,1,opt,name=logID,proto3" json:"logID,omitempty"`
	ManifestURL string `protobuf:"bytes,2,opt,name=manifestURL,proto3" json:"manifestURL,omitempty"`
	ViewURL     string `protobuf:"bytes,3,opt,name=viewURL,proto3" json:"viewURL,omitempty"`
}

func (x *UploadLogBundleResponse) Reset() {
	*x = UploadLogBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_logs_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLogBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLogBundleResponse) ProtoMessage() {}

func (x *UploadLogBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_logs_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadLogBundleResponse.ProtoReflect.Descriptor instead.
func (*UploadLogBundleResponse) Descriptor() ([]byte, []int) {
	return file_api_public_logs_api_proto_rawDescGZIP(), []int{0}
}

func (x *UploadLogBundleResponse) GetLogID() string {
	if x != nil {
		return x.LogID
	}
	return ""
}

func (x *UploadLogBundleResponse) GetManifestURL() string {
	if x != nil {
		return x.ManifestURL
	}
	return ""
}

func (x *UploadLogBundleResponse) GetViewURL() string {
	if x != nil {
		return x.ViewURL
	}
	return ""
}

type ListAvaliableLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogIDs []string `protobuf:"bytes,1,rep,name=logIDs,proto3" json:"logIDs,omitempty"`
}

func (x *ListAvaliableLogsResponse) Reset() {
	*x = ListAvaliableLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_logs_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAvaliableLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAvaliableLogsResponse) ProtoMessage() {}

func (x *ListAvaliableLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_logs_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAvaliableLogsResponse.ProtoReflect.Descriptor instead.
func (*ListAvaliableLogsResponse) Descriptor() ([]byte, []int) {
	return file_api_public_logs_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListAvaliableLogsResponse) GetLogIDs() []string {
	if x != nil {
		return x.LogIDs
	}
	return nil
}

var File_api_public_logs_api_proto protoreflect.FileDescriptor

var file_api_public_logs_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x6b, 0x0a, 0x17,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c, 0x22, 0x33, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x73, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_public_logs_api_proto_rawDescOnce sync.Once
	file_api_public_logs_api_proto_rawDescData = file_api_public_logs_api_proto_rawDesc
)

func file_api_public_logs_api_proto_rawDescGZIP() []byte {
	file_api_public_logs_api_proto_rawDescOnce.Do(func() {
		file_api_public_logs_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_public_logs_api_proto_rawDescData)
	})
	return file_api_public_logs_api_proto_rawDescData
}

var file_api_public_logs_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_public_logs_api_proto_goTypes = []interface{}{
	(*UploadLogBundleResponse)(nil),   // 0: api.public.logs.UploadLogBundleResponse
	(*ListAvaliableLogsResponse)(nil), // 1: api.public.logs.ListAvaliableLogsResponse
}
var file_api_public_logs_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_public_logs_api_proto_init() }
func file_api_public_logs_api_proto_init() {
	if File_api_public_logs_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_public_logs_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadLogBundleResponse); i {
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
		file_api_public_logs_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAvaliableLogsResponse); i {
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
			RawDescriptor: file_api_public_logs_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_public_logs_api_proto_goTypes,
		DependencyIndexes: file_api_public_logs_api_proto_depIdxs,
		MessageInfos:      file_api_public_logs_api_proto_msgTypes,
	}.Build()
	File_api_public_logs_api_proto = out.File
	file_api_public_logs_api_proto_rawDesc = nil
	file_api_public_logs_api_proto_goTypes = nil
	file_api_public_logs_api_proto_depIdxs = nil
}
