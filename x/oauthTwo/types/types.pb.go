// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/v1beta1/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GrantType int32

const (
	GrantType_GRANT_TYPE_UNSPECIFIED              GrantType = 0
	GrantType_GRANT_TYPE_CLIENT_CREDENTIALS_GRANT GrantType = 1
	GrantType_GRANT_TYPE_DEVICE_CODE_GRANT        GrantType = 2
	GrantType_GRANT_TYPE_AUTHORIZATION_CODE_GRANT GrantType = 3
)

var GrantType_name = map[int32]string{
	0: "GRANT_TYPE_UNSPECIFIED",
	1: "GRANT_TYPE_CLIENT_CREDENTIALS_GRANT",
	2: "GRANT_TYPE_DEVICE_CODE_GRANT",
	3: "GRANT_TYPE_AUTHORIZATION_CODE_GRANT",
}

var GrantType_value = map[string]int32{
	"GRANT_TYPE_UNSPECIFIED":              0,
	"GRANT_TYPE_CLIENT_CREDENTIALS_GRANT": 1,
	"GRANT_TYPE_DEVICE_CODE_GRANT":        2,
	"GRANT_TYPE_AUTHORIZATION_CODE_GRANT": 3,
}

func (x GrantType) String() string {
	return proto.EnumName(GrantType_name, int32(x))
}

func (GrantType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43019df9ca413057, []int{0}
}

func init() {
	proto.RegisterEnum("beheroes.doxchain.oauthtwo.v1beta1.GrantType", GrantType_name, GrantType_value)
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/v1beta1/types.proto", fileDescriptor_43019df9ca413057)
}

var fileDescriptor_43019df9ca413057 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xc9, 0xaf, 0x48,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4f, 0x2c, 0x2d, 0xc9, 0x28, 0x29, 0xcf, 0xd7, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x4a, 0x4a, 0xcd, 0x48, 0x2d, 0xca, 0x4f, 0x2d, 0xd6, 0x83, 0xa9, 0xd7, 0x83,
	0xa9, 0xd7, 0x83, 0xaa, 0xd7, 0x9a, 0xcd, 0xc8, 0xc5, 0xe9, 0x5e, 0x94, 0x98, 0x57, 0x12, 0x52,
	0x59, 0x90, 0x2a, 0x24, 0xc5, 0x25, 0xe6, 0x1e, 0xe4, 0xe8, 0x17, 0x12, 0x1f, 0x12, 0x19, 0xe0,
	0x1a, 0x1f, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0xa4,
	0xce, 0xa5, 0x8c, 0x24, 0xe7, 0xec, 0xe3, 0xe9, 0xea, 0x17, 0x12, 0xef, 0x1c, 0xe4, 0xea, 0xe2,
	0xea, 0x17, 0xe2, 0xe9, 0xe8, 0x13, 0x1c, 0x0f, 0x96, 0x15, 0x60, 0x14, 0x52, 0xe0, 0x92, 0x41,
	0x52, 0xe8, 0xe2, 0x1a, 0xe6, 0xe9, 0xec, 0x1a, 0xef, 0xec, 0xef, 0xe2, 0x0a, 0x55, 0xc1, 0x84,
	0x66, 0x94, 0x63, 0x68, 0x88, 0x87, 0x7f, 0x90, 0x67, 0x94, 0x63, 0x88, 0xa7, 0xbf, 0x1f, 0xb2,
	0x42, 0x66, 0x27, 0x8f, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4a, 0xd5, 0x85, 0xf8, 0x53,
	0x1f, 0x1e, 0x2e, 0x15, 0x88, 0x90, 0x01, 0x87, 0x48, 0x12, 0x1b, 0x38, 0x48, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x52, 0x6d, 0x47, 0x3b, 0x01, 0x00, 0x00,
}
