// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ionscale/v1/derp.proto

package ionscalev1

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

type GetDERPMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDERPMapRequest) Reset() {
	*x = GetDERPMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_derp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDERPMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDERPMapRequest) ProtoMessage() {}

func (x *GetDERPMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_derp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDERPMapRequest.ProtoReflect.Descriptor instead.
func (*GetDERPMapRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_derp_proto_rawDescGZIP(), []int{0}
}

type GetDERPMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetDERPMapResponse) Reset() {
	*x = GetDERPMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_derp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDERPMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDERPMapResponse) ProtoMessage() {}

func (x *GetDERPMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_derp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDERPMapResponse.ProtoReflect.Descriptor instead.
func (*GetDERPMapResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_derp_proto_rawDescGZIP(), []int{1}
}

func (x *GetDERPMapResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetDERPMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetDERPMapRequest) Reset() {
	*x = SetDERPMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_derp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDERPMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDERPMapRequest) ProtoMessage() {}

func (x *SetDERPMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_derp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDERPMapRequest.ProtoReflect.Descriptor instead.
func (*SetDERPMapRequest) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_derp_proto_rawDescGZIP(), []int{2}
}

func (x *SetDERPMapRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetDERPMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetDERPMapResponse) Reset() {
	*x = SetDERPMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_derp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDERPMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDERPMapResponse) ProtoMessage() {}

func (x *SetDERPMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_derp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDERPMapResponse.ProtoReflect.Descriptor instead.
func (*SetDERPMapResponse) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_derp_proto_rawDescGZIP(), []int{3}
}

func (x *SetDERPMapResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_ionscale_v1_derp_proto protoreflect.FileDescriptor

var file_ionscale_v1_derp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52,
	0x50, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65,
	0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ionscale_v1_derp_proto_rawDescOnce sync.Once
	file_ionscale_v1_derp_proto_rawDescData = file_ionscale_v1_derp_proto_rawDesc
)

func file_ionscale_v1_derp_proto_rawDescGZIP() []byte {
	file_ionscale_v1_derp_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_derp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ionscale_v1_derp_proto_rawDescData)
	})
	return file_ionscale_v1_derp_proto_rawDescData
}

var file_ionscale_v1_derp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ionscale_v1_derp_proto_goTypes = []interface{}{
	(*GetDERPMapRequest)(nil),  // 0: ionscale.v1.GetDERPMapRequest
	(*GetDERPMapResponse)(nil), // 1: ionscale.v1.GetDERPMapResponse
	(*SetDERPMapRequest)(nil),  // 2: ionscale.v1.SetDERPMapRequest
	(*SetDERPMapResponse)(nil), // 3: ionscale.v1.SetDERPMapResponse
}
var file_ionscale_v1_derp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ionscale_v1_derp_proto_init() }
func file_ionscale_v1_derp_proto_init() {
	if File_ionscale_v1_derp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ionscale_v1_derp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDERPMapRequest); i {
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
		file_ionscale_v1_derp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDERPMapResponse); i {
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
		file_ionscale_v1_derp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDERPMapRequest); i {
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
		file_ionscale_v1_derp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDERPMapResponse); i {
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
			RawDescriptor: file_ionscale_v1_derp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_derp_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_derp_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_derp_proto_msgTypes,
	}.Build()
	File_ionscale_v1_derp_proto = out.File
	file_ionscale_v1_derp_proto_rawDesc = nil
	file_ionscale_v1_derp_proto_goTypes = nil
	file_ionscale_v1_derp_proto_depIdxs = nil
}
