// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.5
// source: dynago-message.proto

package messages

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type DynagoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters map[string]*anypb.Any `protobuf:"bytes,1,rep,name=Parameters,proto3" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DynagoRequest) Reset() {
	*x = DynagoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynago_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynagoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynagoRequest) ProtoMessage() {}

func (x *DynagoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynago_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynagoRequest.ProtoReflect.Descriptor instead.
func (*DynagoRequest) Descriptor() ([]byte, []int) {
	return file_dynago_message_proto_rawDescGZIP(), []int{0}
}

func (x *DynagoRequest) GetParameters() map[string]*anypb.Any {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type DynagoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]*anypb.Any `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DynagoResponse) Reset() {
	*x = DynagoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynago_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynagoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynagoResponse) ProtoMessage() {}

func (x *DynagoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynago_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynagoResponse.ProtoReflect.Descriptor instead.
func (*DynagoResponse) Descriptor() ([]byte, []int) {
	return file_dynago_message_proto_rawDescGZIP(), []int{1}
}

func (x *DynagoResponse) GetResults() map[string]*anypb.Any {
	if x != nil {
		return x.Results
	}
	return nil
}

type ImportPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=Contents,proto3" json:"Contents,omitempty"`
}

func (x *ImportPluginRequest) Reset() {
	*x = ImportPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynago_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportPluginRequest) ProtoMessage() {}

func (x *ImportPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dynago_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportPluginRequest.ProtoReflect.Descriptor instead.
func (*ImportPluginRequest) Descriptor() ([]byte, []int) {
	return file_dynago_message_proto_rawDescGZIP(), []int{2}
}

func (x *ImportPluginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportPluginRequest) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type ImportPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *ImportPluginResponse) Reset() {
	*x = ImportPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dynago_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportPluginResponse) ProtoMessage() {}

func (x *ImportPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dynago_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportPluginResponse.ProtoReflect.Descriptor instead.
func (*ImportPluginResponse) Descriptor() ([]byte, []int) {
	return file_dynago_message_proto_rawDescGZIP(), []int{3}
}

func (x *ImportPluginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_dynago_message_proto protoreflect.FileDescriptor

var file_dynago_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x44, 0x79,
	0x6e, 0x61, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x53, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01, 0x0a, 0x0e, 0x44, 0x79, 0x6e, 0x61,
	0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x79,
	0x6e, 0x61, 0x67, 0x6f, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x13, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x49, 0x0a, 0x0d, 0x44, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x15, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x67, 0x6f, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x67, 0x6f, 0x2e, 0x44, 0x79, 0x6e,
	0x61, 0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5a, 0x0a, 0x13, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x67, 0x6f, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x67, 0x6f, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6d, 0x70, 0x30, 0x39, 0x31, 0x2f, 0x64, 0x79, 0x6e,
	0x61, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dynago_message_proto_rawDescOnce sync.Once
	file_dynago_message_proto_rawDescData = file_dynago_message_proto_rawDesc
)

func file_dynago_message_proto_rawDescGZIP() []byte {
	file_dynago_message_proto_rawDescOnce.Do(func() {
		file_dynago_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dynago_message_proto_rawDescData)
	})
	return file_dynago_message_proto_rawDescData
}

var file_dynago_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dynago_message_proto_goTypes = []interface{}{
	(*DynagoRequest)(nil),        // 0: dynago.DynagoRequest
	(*DynagoResponse)(nil),       // 1: dynago.DynagoResponse
	(*ImportPluginRequest)(nil),  // 2: dynago.ImportPluginRequest
	(*ImportPluginResponse)(nil), // 3: dynago.ImportPluginResponse
	nil,                          // 4: dynago.DynagoRequest.ParametersEntry
	nil,                          // 5: dynago.DynagoResponse.ResultsEntry
	(*anypb.Any)(nil),            // 6: google.protobuf.Any
}
var file_dynago_message_proto_depIdxs = []int32{
	4, // 0: dynago.DynagoRequest.Parameters:type_name -> dynago.DynagoRequest.ParametersEntry
	5, // 1: dynago.DynagoResponse.Results:type_name -> dynago.DynagoResponse.ResultsEntry
	6, // 2: dynago.DynagoRequest.ParametersEntry.value:type_name -> google.protobuf.Any
	6, // 3: dynago.DynagoResponse.ResultsEntry.value:type_name -> google.protobuf.Any
	0, // 4: dynago.DynagoService.Process:input_type -> dynago.DynagoRequest
	2, // 5: dynago.ImportPluginService.Import:input_type -> dynago.ImportPluginRequest
	1, // 6: dynago.DynagoService.Process:output_type -> dynago.DynagoResponse
	3, // 7: dynago.ImportPluginService.Import:output_type -> dynago.ImportPluginResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dynago_message_proto_init() }
func file_dynago_message_proto_init() {
	if File_dynago_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dynago_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynagoRequest); i {
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
		file_dynago_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynagoResponse); i {
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
		file_dynago_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportPluginRequest); i {
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
		file_dynago_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportPluginResponse); i {
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
			RawDescriptor: file_dynago_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_dynago_message_proto_goTypes,
		DependencyIndexes: file_dynago_message_proto_depIdxs,
		MessageInfos:      file_dynago_message_proto_msgTypes,
	}.Build()
	File_dynago_message_proto = out.File
	file_dynago_message_proto_rawDesc = nil
	file_dynago_message_proto_goTypes = nil
	file_dynago_message_proto_depIdxs = nil
}
