// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc-gen-swagger/options/annotations.proto

package options // import "github.com/contiamo/grpc-gateway/protoc-gen-swagger/options"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Openapiv2Swagger = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*Swagger)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_swagger",
	Tag:           "bytes,1042,opt,name=openapiv2_swagger,json=openapiv2Swagger",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Operation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Operation)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_operation",
	Tag:           "bytes,1042,opt,name=openapiv2_operation,json=openapiv2Operation",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Schema = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Schema)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_schema",
	Tag:           "bytes,1042,opt,name=openapiv2_schema,json=openapiv2Schema",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Tag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Tag)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_tag",
	Tag:           "bytes,1042,opt,name=openapiv2_tag,json=openapiv2Tag",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

var E_Openapiv2Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*JSONSchema)(nil),
	Field:         1042,
	Name:          "grpc.gateway.protoc_gen_swagger.options.openapiv2_field",
	Tag:           "bytes,1042,opt,name=openapiv2_field,json=openapiv2Field",
	Filename:      "protoc-gen-swagger/options/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Openapiv2Swagger)
	proto.RegisterExtension(E_Openapiv2Operation)
	proto.RegisterExtension(E_Openapiv2Schema)
	proto.RegisterExtension(E_Openapiv2Tag)
	proto.RegisterExtension(E_Openapiv2Field)
}

func init() {
	proto.RegisterFile("protoc-gen-swagger/options/annotations.proto", fileDescriptor_annotations_109a17117359d0a6)
}

var fileDescriptor_annotations_109a17117359d0a6 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xea, 0x40,
	0x14, 0xc5, 0xc3, 0xe6, 0xe5, 0xa5, 0xef, 0xa9, 0x58, 0x37, 0x86, 0xf8, 0x87, 0x9d, 0xc6, 0xc0,
	0x8c, 0x81, 0x5d, 0x77, 0x6a, 0xe2, 0xc2, 0x44, 0x49, 0x0a, 0x2b, 0x37, 0x64, 0x18, 0x2e, 0x97,
	0x49, 0x4a, 0xef, 0x64, 0x66, 0x80, 0x90, 0xb0, 0xf4, 0x13, 0xf8, 0x89, 0x8d, 0xd3, 0xd2, 0x9a,
	0x8a, 0xa6, 0xbb, 0xde, 0xdb, 0x39, 0xe7, 0x77, 0x7a, 0x3a, 0x41, 0x47, 0x1b, 0x72, 0x24, 0xbb,
	0x08, 0x69, 0xd7, 0xae, 0x05, 0x22, 0x18, 0x4e, 0xda, 0x29, 0x4a, 0x2d, 0x17, 0x69, 0x4a, 0x4e,
	0xf8, 0x67, 0xe6, 0x8f, 0x85, 0x57, 0x68, 0xb4, 0x64, 0x28, 0x1c, 0xac, 0xc5, 0x26, 0xdb, 0xc9,
	0x31, 0x42, 0x3a, 0xce, 0xa5, 0x2c, 0x97, 0xb6, 0x6e, 0x7e, 0xb1, 0x25, 0x0d, 0xa9, 0xd0, 0x6a,
	0xd5, 0xcb, 0x0c, 0x5a, 0x6d, 0x24, 0xc2, 0x04, 0xb8, 0x9f, 0x26, 0xcb, 0x19, 0x9f, 0x82, 0x95,
	0x46, 0x69, 0x47, 0x26, 0x3b, 0x11, 0x6d, 0x83, 0xe3, 0x42, 0xb4, 0x43, 0x85, 0x67, 0x2c, 0xd3,
	0xb1, 0x9d, 0x8e, 0x3d, 0xaa, 0x04, 0x06, 0x19, 0xe4, 0xf4, 0xfd, 0x6f, 0xbb, 0x71, 0xfd, 0xaf,
	0x77, 0xcb, 0x6a, 0x26, 0x66, 0xc3, 0x6c, 0x8e, 0x9b, 0x05, 0x29, 0xdf, 0x44, 0x6f, 0x8d, 0xe0,
	0xa4, 0xc4, 0x93, 0x06, 0xe3, 0x3b, 0x09, 0x2f, 0xbe, 0x05, 0x78, 0x06, 0x37, 0xa7, 0x69, 0x25,
	0x42, 0xaf, 0x76, 0x84, 0xc1, 0xce, 0x3a, 0x0e, 0x0b, 0x5e, 0xb1, 0x8b, 0xb6, 0x41, 0xf3, 0x4b,
	0x09, 0x72, 0x0e, 0x0b, 0x11, 0x5e, 0xee, 0x89, 0x60, 0xad, 0xc0, 0x6a, 0x0d, 0xbc, 0x7e, 0x0d,
	0xde, 0x38, 0x3e, 0x2a, 0x5b, 0xf0, 0x8b, 0xc8, 0x06, 0x07, 0x25, 0xdd, 0x09, 0xdc, 0x83, 0x1e,
	0x82, 0x59, 0x29, 0x59, 0x45, 0x77, 0x6a, 0xa3, 0x47, 0x02, 0xe3, 0xff, 0x05, 0x64, 0x24, 0x30,
	0xda, 0x06, 0x65, 0x8e, 0xf1, 0x4c, 0x41, 0x32, 0x0d, 0xcf, 0xf7, 0xfc, 0x75, 0x48, 0xaa, 0x9d,
	0xf7, 0x6b, 0x43, 0x9f, 0x86, 0x83, 0x97, 0xfc, 0x9b, 0x0f, 0x0b, 0x96, 0xb7, 0xbc, 0x7f, 0x78,
	0xbd, 0x43, 0xe5, 0xe6, 0xcb, 0x09, 0x93, 0xb4, 0xe0, 0x9f, 0x86, 0x5d, 0x90, 0x64, 0x37, 0xd6,
	0x41, 0x3e, 0xe6, 0xfe, 0xfc, 0xe7, 0xcb, 0x3e, 0xf9, 0xe3, 0xdf, 0xf5, 0x3f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0xc4, 0x41, 0xfb, 0x68, 0x03, 0x00, 0x00,
}
