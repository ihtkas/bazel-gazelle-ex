// protolint:disable MAX_LINE_LENGTH

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: protos/pinger/service.proto

package pingerpb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	At *timestamp.Timestamp `protobuf:"bytes,1,opt,name=at,proto3" json:"at,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pinger_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pinger_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_protos_pinger_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetAt() *timestamp.Timestamp {
	if x != nil {
		return x.At
	}
	return nil
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pinger_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pinger_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_protos_pinger_service_proto_rawDescGZIP(), []int{1}
}

var File_protos_pinger_service_proto protoreflect.FileDescriptor

var file_protos_pinger_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x43, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x02, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x02, 0x61, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x44, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e,
	0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x68, 0x74, 0x6b, 0x61, 0x73, 0x2f,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2d, 0x67, 0x61, 0x7a, 0x65, 0x6c, 0x6c, 0x65, 0x2d, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x3b, 0x70, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_pinger_service_proto_rawDescOnce sync.Once
	file_protos_pinger_service_proto_rawDescData = file_protos_pinger_service_proto_rawDesc
)

func file_protos_pinger_service_proto_rawDescGZIP() []byte {
	file_protos_pinger_service_proto_rawDescOnce.Do(func() {
		file_protos_pinger_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_pinger_service_proto_rawDescData)
	})
	return file_protos_pinger_service_proto_rawDescData
}

var file_protos_pinger_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_pinger_service_proto_goTypes = []interface{}{
	(*PingRequest)(nil),         // 0: pinger.PingRequest
	(*PingResponse)(nil),        // 1: pinger.PingResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_protos_pinger_service_proto_depIdxs = []int32{
	2, // 0: pinger.PingRequest.at:type_name -> google.protobuf.Timestamp
	0, // 1: pinger.PingerService.Ping:input_type -> pinger.PingRequest
	1, // 2: pinger.PingerService.Ping:output_type -> pinger.PingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_pinger_service_proto_init() }
func file_protos_pinger_service_proto_init() {
	if File_protos_pinger_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_pinger_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_protos_pinger_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_protos_pinger_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_pinger_service_proto_goTypes,
		DependencyIndexes: file_protos_pinger_service_proto_depIdxs,
		MessageInfos:      file_protos_pinger_service_proto_msgTypes,
	}.Build()
	File_protos_pinger_service_proto = out.File
	file_protos_pinger_service_proto_rawDesc = nil
	file_protos_pinger_service_proto_goTypes = nil
	file_protos_pinger_service_proto_depIdxs = nil
}
