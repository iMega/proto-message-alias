// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alias.proto

/*
Package alias is a generated protocol buffer package.

It is generated from these files:
	alias.proto

It has these top-level messages:
*/
package alias

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_MessageAlias = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         57382,
	Name:          "alias.message_alias",
	Tag:           "bytes,57382,opt,name=message_alias,json=messageAlias",
	Filename:      "alias.proto",
}

func init() {
	proto.RegisterExtension(E_MessageAlias)
}

func init() { proto.RegisterFile("alias.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcc, 0xc9, 0x4c,
	0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x14, 0xd2, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x82, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99,
	0x05, 0x25, 0xf9, 0x45, 0x10, 0x85, 0x56, 0x6e, 0x5c, 0xbc, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0xa9, 0xf1, 0x60, 0x2d, 0x42, 0xf2, 0x7a, 0x10, 0x3d, 0x7a, 0x30, 0x3d, 0x7a, 0xbe, 0x10, 0x79,
	0xff, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x89, 0x65, 0x07, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x78, 0xa0, 0xfa, 0x1c, 0x41, 0xda, 0x92, 0xd8, 0xc0, 0xca, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0xe6, 0xeb, 0xc6, 0x86, 0x00, 0x00, 0x00,
}
