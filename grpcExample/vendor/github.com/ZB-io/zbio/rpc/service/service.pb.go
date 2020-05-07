// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: service.proto

package service

import (
	context "context"
	common "github.com/ZB-io/zbio/rpc/common"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x0a, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x1f, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1f, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72,
	0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x25, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72,
	0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x27, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x7a,
	0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x7a,
	0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x0a, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e,
	0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0b, 0x50, 0x65, 0x65,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e,
	0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x65, 0x0a, 0x10,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x12, 0x2a, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x7a,
	0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x7a, 0x62, 0x69,
	0x6f, 0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x7a, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f,
	0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x7a, 0x62, 0x69, 0x6f,
	0x2e, 0x72, 0x6f, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x42, 0x2d, 0x69, 0x6f, 0x2f, 0x7a, 0x62,
	0x69, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*common.TestRequest)(nil),               // 0: zbio.roost.common.TestRequest
	(*common.TopicRequest)(nil),              // 1: zbio.roost.common.TopicRequest
	(*common.DeleteTopicRequest)(nil),        // 2: zbio.roost.common.DeleteTopicRequest
	(*empty.Empty)(nil),                      // 3: google.protobuf.Empty
	(*common.DescribeTopicRequest)(nil),      // 4: zbio.roost.common.DescribeTopicRequest
	(*common.NewMessageRequest)(nil),         // 5: zbio.roost.common.NewMessageRequest
	(*common.MessageRequest)(nil),            // 6: zbio.roost.common.MessageRequest
	(*common.PeekMessageRequest)(nil),        // 7: zbio.roost.common.PeekMessageRequest
	(*common.ConsumerRequest)(nil),           // 8: zbio.roost.common.ConsumerRequest
	(*common.DescribeConsumerRequest)(nil),   // 9: zbio.roost.common.DescribeConsumerRequest
	(*common.TestResponse)(nil),              // 10: zbio.roost.common.TestResponse
	(*common.TopicResponse)(nil),             // 11: zbio.roost.common.TopicResponse
	(*common.UpdateTopicResponse)(nil),       // 12: zbio.roost.common.UpdateTopicResponse
	(*common.DeleteTopicResponse)(nil),       // 13: zbio.roost.common.DeleteTopicResponse
	(*common.DescribeTopicResponse)(nil),     // 14: zbio.roost.common.DescribeTopicResponse
	(*common.NewMessageResponse)(nil),        // 15: zbio.roost.common.NewMessageResponse
	(*common.GetMessageResponse)(nil),        // 16: zbio.roost.common.GetMessageResponse
	(*common.AckMessageResponse)(nil),        // 17: zbio.roost.common.AckMessageResponse
	(*common.ConsumerResponse)(nil),          // 18: zbio.roost.common.ConsumerResponse
	(*common.GetConsumerOffsetResponse)(nil), // 19: zbio.roost.common.GetConsumerOffsetResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: zbio.roost.service.Service.TestMessage:input_type -> zbio.roost.common.TestRequest
	1,  // 1: zbio.roost.service.Service.CreateTopic:input_type -> zbio.roost.common.TopicRequest
	1,  // 2: zbio.roost.service.Service.UpdateTopic:input_type -> zbio.roost.common.TopicRequest
	2,  // 3: zbio.roost.service.Service.DeleteTopic:input_type -> zbio.roost.common.DeleteTopicRequest
	3,  // 4: zbio.roost.service.Service.ListTopic:input_type -> google.protobuf.Empty
	4,  // 5: zbio.roost.service.Service.DescribeTopic:input_type -> zbio.roost.common.DescribeTopicRequest
	5,  // 6: zbio.roost.service.Service.NewMessage:input_type -> zbio.roost.common.NewMessageRequest
	6,  // 7: zbio.roost.service.Service.GetMessage:input_type -> zbio.roost.common.MessageRequest
	6,  // 8: zbio.roost.service.Service.AckMessage:input_type -> zbio.roost.common.MessageRequest
	7,  // 9: zbio.roost.service.Service.PeekMessage:input_type -> zbio.roost.common.PeekMessageRequest
	8,  // 10: zbio.roost.service.Service.SubscribeConsumer:input_type -> zbio.roost.common.ConsumerRequest
	9,  // 11: zbio.roost.service.Service.DescribeConsumer:input_type -> zbio.roost.common.DescribeConsumerRequest
	3,  // 12: zbio.roost.service.Service.GetConsumer:input_type -> google.protobuf.Empty
	6,  // 13: zbio.roost.service.Service.GetConsumerOffset:input_type -> zbio.roost.common.MessageRequest
	10, // 14: zbio.roost.service.Service.TestMessage:output_type -> zbio.roost.common.TestResponse
	11, // 15: zbio.roost.service.Service.CreateTopic:output_type -> zbio.roost.common.TopicResponse
	12, // 16: zbio.roost.service.Service.UpdateTopic:output_type -> zbio.roost.common.UpdateTopicResponse
	13, // 17: zbio.roost.service.Service.DeleteTopic:output_type -> zbio.roost.common.DeleteTopicResponse
	14, // 18: zbio.roost.service.Service.ListTopic:output_type -> zbio.roost.common.DescribeTopicResponse
	14, // 19: zbio.roost.service.Service.DescribeTopic:output_type -> zbio.roost.common.DescribeTopicResponse
	15, // 20: zbio.roost.service.Service.NewMessage:output_type -> zbio.roost.common.NewMessageResponse
	16, // 21: zbio.roost.service.Service.GetMessage:output_type -> zbio.roost.common.GetMessageResponse
	17, // 22: zbio.roost.service.Service.AckMessage:output_type -> zbio.roost.common.AckMessageResponse
	16, // 23: zbio.roost.service.Service.PeekMessage:output_type -> zbio.roost.common.GetMessageResponse
	18, // 24: zbio.roost.service.Service.SubscribeConsumer:output_type -> zbio.roost.common.ConsumerResponse
	18, // 25: zbio.roost.service.Service.DescribeConsumer:output_type -> zbio.roost.common.ConsumerResponse
	18, // 26: zbio.roost.service.Service.GetConsumer:output_type -> zbio.roost.common.ConsumerResponse
	19, // 27: zbio.roost.service.Service.GetConsumerOffset:output_type -> zbio.roost.common.GetConsumerOffsetResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	TestMessage(ctx context.Context, in *common.TestRequest, opts ...grpc.CallOption) (*common.TestResponse, error)
	// Topic is responsible for creating new Topic(s) and filtered Topic data is returned.
	CreateTopic(ctx context.Context, in *common.TopicRequest, opts ...grpc.CallOption) (*common.TopicResponse, error)
	UpdateTopic(ctx context.Context, in *common.TopicRequest, opts ...grpc.CallOption) (*common.UpdateTopicResponse, error)
	DeleteTopic(ctx context.Context, in *common.DeleteTopicRequest, opts ...grpc.CallOption) (*common.DeleteTopicResponse, error)
	ListTopic(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.DescribeTopicResponse, error)
	// Describes topic showing information about name, partition, owner_broker, replicas, producers, consumers
	DescribeTopic(ctx context.Context, in *common.DescribeTopicRequest, opts ...grpc.CallOption) (*common.DescribeTopicResponse, error)
	// Produce a message
	NewMessage(ctx context.Context, in *common.NewMessageRequest, opts ...grpc.CallOption) (*common.NewMessageResponse, error)
	// Retrieve a Message
	GetMessage(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.GetMessageResponse, error)
	// Acknowledge message has been received
	AckMessage(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.AckMessageResponse, error)
	// Peek message
	PeekMessage(ctx context.Context, in *common.PeekMessageRequest, opts ...grpc.CallOption) (*common.GetMessageResponse, error)
	// Consumer related
	SubscribeConsumer(ctx context.Context, opts ...grpc.CallOption) (Service_SubscribeConsumerClient, error)
	DescribeConsumer(ctx context.Context, in *common.DescribeConsumerRequest, opts ...grpc.CallOption) (*common.ConsumerResponse, error)
	GetConsumer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.ConsumerResponse, error)
	// Get consumer bookmark for a topic, consumer group
	GetConsumerOffset(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.GetConsumerOffsetResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) TestMessage(ctx context.Context, in *common.TestRequest, opts ...grpc.CallOption) (*common.TestResponse, error) {
	out := new(common.TestResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/TestMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateTopic(ctx context.Context, in *common.TopicRequest, opts ...grpc.CallOption) (*common.TopicResponse, error) {
	out := new(common.TopicResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateTopic(ctx context.Context, in *common.TopicRequest, opts ...grpc.CallOption) (*common.UpdateTopicResponse, error) {
	out := new(common.UpdateTopicResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/UpdateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteTopic(ctx context.Context, in *common.DeleteTopicRequest, opts ...grpc.CallOption) (*common.DeleteTopicResponse, error) {
	out := new(common.DeleteTopicResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListTopic(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.DescribeTopicResponse, error) {
	out := new(common.DescribeTopicResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/ListTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeTopic(ctx context.Context, in *common.DescribeTopicRequest, opts ...grpc.CallOption) (*common.DescribeTopicResponse, error) {
	out := new(common.DescribeTopicResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/DescribeTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NewMessage(ctx context.Context, in *common.NewMessageRequest, opts ...grpc.CallOption) (*common.NewMessageResponse, error) {
	out := new(common.NewMessageResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/NewMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetMessage(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.GetMessageResponse, error) {
	out := new(common.GetMessageResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AckMessage(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.AckMessageResponse, error) {
	out := new(common.AckMessageResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/AckMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PeekMessage(ctx context.Context, in *common.PeekMessageRequest, opts ...grpc.CallOption) (*common.GetMessageResponse, error) {
	out := new(common.GetMessageResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/PeekMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SubscribeConsumer(ctx context.Context, opts ...grpc.CallOption) (Service_SubscribeConsumerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/zbio.roost.service.Service/SubscribeConsumer", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceSubscribeConsumerClient{stream}
	return x, nil
}

type Service_SubscribeConsumerClient interface {
	Send(*common.ConsumerRequest) error
	Recv() (*common.ConsumerResponse, error)
	grpc.ClientStream
}

type serviceSubscribeConsumerClient struct {
	grpc.ClientStream
}

func (x *serviceSubscribeConsumerClient) Send(m *common.ConsumerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceSubscribeConsumerClient) Recv() (*common.ConsumerResponse, error) {
	m := new(common.ConsumerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) DescribeConsumer(ctx context.Context, in *common.DescribeConsumerRequest, opts ...grpc.CallOption) (*common.ConsumerResponse, error) {
	out := new(common.ConsumerResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/DescribeConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetConsumer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.ConsumerResponse, error) {
	out := new(common.ConsumerResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/GetConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetConsumerOffset(ctx context.Context, in *common.MessageRequest, opts ...grpc.CallOption) (*common.GetConsumerOffsetResponse, error) {
	out := new(common.GetConsumerOffsetResponse)
	err := c.cc.Invoke(ctx, "/zbio.roost.service.Service/GetConsumerOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	TestMessage(context.Context, *common.TestRequest) (*common.TestResponse, error)
	// Topic is responsible for creating new Topic(s) and filtered Topic data is returned.
	CreateTopic(context.Context, *common.TopicRequest) (*common.TopicResponse, error)
	UpdateTopic(context.Context, *common.TopicRequest) (*common.UpdateTopicResponse, error)
	DeleteTopic(context.Context, *common.DeleteTopicRequest) (*common.DeleteTopicResponse, error)
	ListTopic(context.Context, *empty.Empty) (*common.DescribeTopicResponse, error)
	// Describes topic showing information about name, partition, owner_broker, replicas, producers, consumers
	DescribeTopic(context.Context, *common.DescribeTopicRequest) (*common.DescribeTopicResponse, error)
	// Produce a message
	NewMessage(context.Context, *common.NewMessageRequest) (*common.NewMessageResponse, error)
	// Retrieve a Message
	GetMessage(context.Context, *common.MessageRequest) (*common.GetMessageResponse, error)
	// Acknowledge message has been received
	AckMessage(context.Context, *common.MessageRequest) (*common.AckMessageResponse, error)
	// Peek message
	PeekMessage(context.Context, *common.PeekMessageRequest) (*common.GetMessageResponse, error)
	// Consumer related
	SubscribeConsumer(Service_SubscribeConsumerServer) error
	DescribeConsumer(context.Context, *common.DescribeConsumerRequest) (*common.ConsumerResponse, error)
	GetConsumer(context.Context, *empty.Empty) (*common.ConsumerResponse, error)
	// Get consumer bookmark for a topic, consumer group
	GetConsumerOffset(context.Context, *common.MessageRequest) (*common.GetConsumerOffsetResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) TestMessage(context.Context, *common.TestRequest) (*common.TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMessage not implemented")
}
func (*UnimplementedServiceServer) CreateTopic(context.Context, *common.TopicRequest) (*common.TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (*UnimplementedServiceServer) UpdateTopic(context.Context, *common.TopicRequest) (*common.UpdateTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (*UnimplementedServiceServer) DeleteTopic(context.Context, *common.DeleteTopicRequest) (*common.DeleteTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (*UnimplementedServiceServer) ListTopic(context.Context, *empty.Empty) (*common.DescribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopic not implemented")
}
func (*UnimplementedServiceServer) DescribeTopic(context.Context, *common.DescribeTopicRequest) (*common.DescribeTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTopic not implemented")
}
func (*UnimplementedServiceServer) NewMessage(context.Context, *common.NewMessageRequest) (*common.NewMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMessage not implemented")
}
func (*UnimplementedServiceServer) GetMessage(context.Context, *common.MessageRequest) (*common.GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (*UnimplementedServiceServer) AckMessage(context.Context, *common.MessageRequest) (*common.AckMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckMessage not implemented")
}
func (*UnimplementedServiceServer) PeekMessage(context.Context, *common.PeekMessageRequest) (*common.GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeekMessage not implemented")
}
func (*UnimplementedServiceServer) SubscribeConsumer(Service_SubscribeConsumerServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeConsumer not implemented")
}
func (*UnimplementedServiceServer) DescribeConsumer(context.Context, *common.DescribeConsumerRequest) (*common.ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeConsumer not implemented")
}
func (*UnimplementedServiceServer) GetConsumer(context.Context, *empty.Empty) (*common.ConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumer not implemented")
}
func (*UnimplementedServiceServer) GetConsumerOffset(context.Context, *common.MessageRequest) (*common.GetConsumerOffsetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerOffset not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_TestMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TestMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/TestMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TestMessage(ctx, req.(*common.TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateTopic(ctx, req.(*common.TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateTopic(ctx, req.(*common.TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteTopic(ctx, req.(*common.DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/ListTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListTopic(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.DescribeTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/DescribeTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeTopic(ctx, req.(*common.DescribeTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_NewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.NewMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).NewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/NewMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).NewMessage(ctx, req.(*common.NewMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetMessage(ctx, req.(*common.MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AckMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AckMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/AckMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AckMessage(ctx, req.(*common.MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PeekMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PeekMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PeekMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/PeekMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PeekMessage(ctx, req.(*common.PeekMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_SubscribeConsumer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).SubscribeConsumer(&serviceSubscribeConsumerServer{stream})
}

type Service_SubscribeConsumerServer interface {
	Send(*common.ConsumerResponse) error
	Recv() (*common.ConsumerRequest, error)
	grpc.ServerStream
}

type serviceSubscribeConsumerServer struct {
	grpc.ServerStream
}

func (x *serviceSubscribeConsumerServer) Send(m *common.ConsumerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceSubscribeConsumerServer) Recv() (*common.ConsumerRequest, error) {
	m := new(common.ConsumerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_DescribeConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.DescribeConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/DescribeConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeConsumer(ctx, req.(*common.DescribeConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/GetConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetConsumer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetConsumerOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetConsumerOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zbio.roost.service.Service/GetConsumerOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetConsumerOffset(ctx, req.(*common.MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zbio.roost.service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMessage",
			Handler:    _Service_TestMessage_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _Service_CreateTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _Service_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Service_DeleteTopic_Handler,
		},
		{
			MethodName: "ListTopic",
			Handler:    _Service_ListTopic_Handler,
		},
		{
			MethodName: "DescribeTopic",
			Handler:    _Service_DescribeTopic_Handler,
		},
		{
			MethodName: "NewMessage",
			Handler:    _Service_NewMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Service_GetMessage_Handler,
		},
		{
			MethodName: "AckMessage",
			Handler:    _Service_AckMessage_Handler,
		},
		{
			MethodName: "PeekMessage",
			Handler:    _Service_PeekMessage_Handler,
		},
		{
			MethodName: "DescribeConsumer",
			Handler:    _Service_DescribeConsumer_Handler,
		},
		{
			MethodName: "GetConsumer",
			Handler:    _Service_GetConsumer_Handler,
		},
		{
			MethodName: "GetConsumerOffset",
			Handler:    _Service_GetConsumerOffset_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeConsumer",
			Handler:       _Service_SubscribeConsumer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}