// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ionscale/v1/ref.proto

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

type Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Ref) Reset() {
	*x = Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ionscale_v1_ref_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ref) ProtoMessage() {}

func (x *Ref) ProtoReflect() protoreflect.Message {
	mi := &file_ionscale_v1_ref_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ref.ProtoReflect.Descriptor instead.
func (*Ref) Descriptor() ([]byte, []int) {
	return file_ionscale_v1_ref_proto_rawDescGZIP(), []int{0}
}

func (x *Ref) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ref) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_ionscale_v1_ref_proto protoreflect.FileDescriptor

var file_ionscale_v1_ref_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73,
	0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ionscale_v1_ref_proto_rawDescOnce sync.Once
	file_ionscale_v1_ref_proto_rawDescData = file_ionscale_v1_ref_proto_rawDesc
)

func file_ionscale_v1_ref_proto_rawDescGZIP() []byte {
	file_ionscale_v1_ref_proto_rawDescOnce.Do(func() {
		file_ionscale_v1_ref_proto_rawDescData = protoimpl.X.CompressGZIP(file_ionscale_v1_ref_proto_rawDescData)
	})
	return file_ionscale_v1_ref_proto_rawDescData
}

var file_ionscale_v1_ref_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ionscale_v1_ref_proto_goTypes = []interface{}{
	(*Ref)(nil), // 0: ionscale.v1.Ref
}
var file_ionscale_v1_ref_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ionscale_v1_ref_proto_init() }
func file_ionscale_v1_ref_proto_init() {
	if File_ionscale_v1_ref_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ionscale_v1_ref_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ref); i {
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
			RawDescriptor: file_ionscale_v1_ref_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ionscale_v1_ref_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_ref_proto_depIdxs,
		MessageInfos:      file_ionscale_v1_ref_proto_msgTypes,
	}.Build()
	File_ionscale_v1_ref_proto = out.File
	file_ionscale_v1_ref_proto_rawDesc = nil
	file_ionscale_v1_ref_proto_goTypes = nil
	file_ionscale_v1_ref_proto_depIdxs = nil
}
