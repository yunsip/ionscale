// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ionscale.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_ionscale_proto protoreflect.FileDescriptor

var file_api_ionscale_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xc3, 0x0c, 0x0a, 0x08, 0x49, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x12,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73,
	0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x43,
	0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x0d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65, 0x62, 0x65, 0x6e, 0x73, 0x2f,
	0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_ionscale_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),        // 0: api.GetVersionRequest
	(*GetDERPMapRequest)(nil),        // 1: api.GetDERPMapRequest
	(*SetDERPMapRequest)(nil),        // 2: api.SetDERPMapRequest
	(*CreateAuthMethodRequest)(nil),  // 3: api.CreateAuthMethodRequest
	(*ListAuthMethodsRequest)(nil),   // 4: api.ListAuthMethodsRequest
	(*CreateTailnetRequest)(nil),     // 5: api.CreateTailnetRequest
	(*GetTailnetRequest)(nil),        // 6: api.GetTailnetRequest
	(*ListTailnetRequest)(nil),       // 7: api.ListTailnetRequest
	(*DeleteTailnetRequest)(nil),     // 8: api.DeleteTailnetRequest
	(*GetDNSConfigRequest)(nil),      // 9: api.GetDNSConfigRequest
	(*SetDNSConfigRequest)(nil),      // 10: api.SetDNSConfigRequest
	(*GetACLPolicyRequest)(nil),      // 11: api.GetACLPolicyRequest
	(*SetACLPolicyRequest)(nil),      // 12: api.SetACLPolicyRequest
	(*GetAuthKeyRequest)(nil),        // 13: api.GetAuthKeyRequest
	(*CreateAuthKeyRequest)(nil),     // 14: api.CreateAuthKeyRequest
	(*DeleteAuthKeyRequest)(nil),     // 15: api.DeleteAuthKeyRequest
	(*ListAuthKeysRequest)(nil),      // 16: api.ListAuthKeysRequest
	(*ListMachinesRequest)(nil),      // 17: api.ListMachinesRequest
	(*ExpireMachineRequest)(nil),     // 18: api.ExpireMachineRequest
	(*DeleteMachineRequest)(nil),     // 19: api.DeleteMachineRequest
	(*GetMachineRoutesRequest)(nil),  // 20: api.GetMachineRoutesRequest
	(*SetMachineRoutesRequest)(nil),  // 21: api.SetMachineRoutesRequest
	(*GetVersionResponse)(nil),       // 22: api.GetVersionResponse
	(*GetDERPMapResponse)(nil),       // 23: api.GetDERPMapResponse
	(*SetDERPMapResponse)(nil),       // 24: api.SetDERPMapResponse
	(*CreateAuthMethodResponse)(nil), // 25: api.CreateAuthMethodResponse
	(*ListAuthMethodsResponse)(nil),  // 26: api.ListAuthMethodsResponse
	(*CreateTailnetResponse)(nil),    // 27: api.CreateTailnetResponse
	(*GetTailnetResponse)(nil),       // 28: api.GetTailnetResponse
	(*ListTailnetResponse)(nil),      // 29: api.ListTailnetResponse
	(*DeleteTailnetResponse)(nil),    // 30: api.DeleteTailnetResponse
	(*GetDNSConfigResponse)(nil),     // 31: api.GetDNSConfigResponse
	(*SetDNSConfigResponse)(nil),     // 32: api.SetDNSConfigResponse
	(*GetACLPolicyResponse)(nil),     // 33: api.GetACLPolicyResponse
	(*SetACLPolicyResponse)(nil),     // 34: api.SetACLPolicyResponse
	(*GetAuthKeyResponse)(nil),       // 35: api.GetAuthKeyResponse
	(*CreateAuthKeyResponse)(nil),    // 36: api.CreateAuthKeyResponse
	(*DeleteAuthKeyResponse)(nil),    // 37: api.DeleteAuthKeyResponse
	(*ListAuthKeysResponse)(nil),     // 38: api.ListAuthKeysResponse
	(*ListMachinesResponse)(nil),     // 39: api.ListMachinesResponse
	(*ExpireMachineResponse)(nil),    // 40: api.ExpireMachineResponse
	(*DeleteMachineResponse)(nil),    // 41: api.DeleteMachineResponse
	(*GetMachineRoutesResponse)(nil), // 42: api.GetMachineRoutesResponse
}
var file_api_ionscale_proto_depIdxs = []int32{
	0,  // 0: api.Ionscale.GetVersion:input_type -> api.GetVersionRequest
	1,  // 1: api.Ionscale.GetDERPMap:input_type -> api.GetDERPMapRequest
	2,  // 2: api.Ionscale.SetDERPMap:input_type -> api.SetDERPMapRequest
	3,  // 3: api.Ionscale.CreateAuthMethod:input_type -> api.CreateAuthMethodRequest
	4,  // 4: api.Ionscale.ListAuthMethods:input_type -> api.ListAuthMethodsRequest
	5,  // 5: api.Ionscale.CreateTailnet:input_type -> api.CreateTailnetRequest
	6,  // 6: api.Ionscale.GetTailnet:input_type -> api.GetTailnetRequest
	7,  // 7: api.Ionscale.ListTailnets:input_type -> api.ListTailnetRequest
	8,  // 8: api.Ionscale.DeleteTailnet:input_type -> api.DeleteTailnetRequest
	9,  // 9: api.Ionscale.GetDNSConfig:input_type -> api.GetDNSConfigRequest
	10, // 10: api.Ionscale.SetDNSConfig:input_type -> api.SetDNSConfigRequest
	11, // 11: api.Ionscale.GetACLPolicy:input_type -> api.GetACLPolicyRequest
	12, // 12: api.Ionscale.SetACLPolicy:input_type -> api.SetACLPolicyRequest
	13, // 13: api.Ionscale.GetAuthKey:input_type -> api.GetAuthKeyRequest
	14, // 14: api.Ionscale.CreateAuthKey:input_type -> api.CreateAuthKeyRequest
	15, // 15: api.Ionscale.DeleteAuthKey:input_type -> api.DeleteAuthKeyRequest
	16, // 16: api.Ionscale.ListAuthKeys:input_type -> api.ListAuthKeysRequest
	17, // 17: api.Ionscale.ListMachines:input_type -> api.ListMachinesRequest
	18, // 18: api.Ionscale.ExpireMachine:input_type -> api.ExpireMachineRequest
	19, // 19: api.Ionscale.DeleteMachine:input_type -> api.DeleteMachineRequest
	20, // 20: api.Ionscale.GetMachineRoutes:input_type -> api.GetMachineRoutesRequest
	21, // 21: api.Ionscale.SetMachineRoutes:input_type -> api.SetMachineRoutesRequest
	22, // 22: api.Ionscale.GetVersion:output_type -> api.GetVersionResponse
	23, // 23: api.Ionscale.GetDERPMap:output_type -> api.GetDERPMapResponse
	24, // 24: api.Ionscale.SetDERPMap:output_type -> api.SetDERPMapResponse
	25, // 25: api.Ionscale.CreateAuthMethod:output_type -> api.CreateAuthMethodResponse
	26, // 26: api.Ionscale.ListAuthMethods:output_type -> api.ListAuthMethodsResponse
	27, // 27: api.Ionscale.CreateTailnet:output_type -> api.CreateTailnetResponse
	28, // 28: api.Ionscale.GetTailnet:output_type -> api.GetTailnetResponse
	29, // 29: api.Ionscale.ListTailnets:output_type -> api.ListTailnetResponse
	30, // 30: api.Ionscale.DeleteTailnet:output_type -> api.DeleteTailnetResponse
	31, // 31: api.Ionscale.GetDNSConfig:output_type -> api.GetDNSConfigResponse
	32, // 32: api.Ionscale.SetDNSConfig:output_type -> api.SetDNSConfigResponse
	33, // 33: api.Ionscale.GetACLPolicy:output_type -> api.GetACLPolicyResponse
	34, // 34: api.Ionscale.SetACLPolicy:output_type -> api.SetACLPolicyResponse
	35, // 35: api.Ionscale.GetAuthKey:output_type -> api.GetAuthKeyResponse
	36, // 36: api.Ionscale.CreateAuthKey:output_type -> api.CreateAuthKeyResponse
	37, // 37: api.Ionscale.DeleteAuthKey:output_type -> api.DeleteAuthKeyResponse
	38, // 38: api.Ionscale.ListAuthKeys:output_type -> api.ListAuthKeysResponse
	39, // 39: api.Ionscale.ListMachines:output_type -> api.ListMachinesResponse
	40, // 40: api.Ionscale.ExpireMachine:output_type -> api.ExpireMachineResponse
	41, // 41: api.Ionscale.DeleteMachine:output_type -> api.DeleteMachineResponse
	42, // 42: api.Ionscale.GetMachineRoutes:output_type -> api.GetMachineRoutesResponse
	42, // 43: api.Ionscale.SetMachineRoutes:output_type -> api.GetMachineRoutesResponse
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_ionscale_proto_init() }
func file_api_ionscale_proto_init() {
	if File_api_ionscale_proto != nil {
		return
	}
	file_api_version_proto_init()
	file_api_tailnets_proto_init()
	file_api_auth_methods_proto_init()
	file_api_auth_keys_proto_init()
	file_api_machines_proto_init()
	file_api_routes_proto_init()
	file_api_dns_proto_init()
	file_api_acl_proto_init()
	file_api_derp_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ionscale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ionscale_proto_goTypes,
		DependencyIndexes: file_api_ionscale_proto_depIdxs,
	}.Build()
	File_api_ionscale_proto = out.File
	file_api_ionscale_proto_rawDesc = nil
	file_api_ionscale_proto_goTypes = nil
	file_api_ionscale_proto_depIdxs = nil
}
