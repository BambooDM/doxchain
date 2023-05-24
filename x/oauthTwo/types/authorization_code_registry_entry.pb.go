// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauthtwo/v1beta1/authorization_code_registry_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/did/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AuthorizationCodeRegistryEntry struct {
	Owner             types.Did `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner"`
	AuthorizationCode string    `protobuf:"bytes,2,opt,name=authorizationCode,proto3" json:"authorizationCode,omitempty"`
	ExpiresAt         int64     `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (m *AuthorizationCodeRegistryEntry) Reset()         { *m = AuthorizationCodeRegistryEntry{} }
func (m *AuthorizationCodeRegistryEntry) String() string { return proto.CompactTextString(m) }
func (*AuthorizationCodeRegistryEntry) ProtoMessage()    {}
func (*AuthorizationCodeRegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1763526c68c589aa, []int{0}
}
func (m *AuthorizationCodeRegistryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizationCodeRegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizationCodeRegistryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizationCodeRegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationCodeRegistryEntry.Merge(m, src)
}
func (m *AuthorizationCodeRegistryEntry) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizationCodeRegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationCodeRegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationCodeRegistryEntry proto.InternalMessageInfo

func (m *AuthorizationCodeRegistryEntry) GetOwner() types.Did {
	if m != nil {
		return m.Owner
	}
	return types.Did{}
}

func (m *AuthorizationCodeRegistryEntry) GetAuthorizationCode() string {
	if m != nil {
		return m.AuthorizationCode
	}
	return ""
}

func (m *AuthorizationCodeRegistryEntry) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthorizationCodeRegistryEntry)(nil), "beheroes.doxchain.oauthtwo.v1beta1.AuthorizationCodeRegistryEntry")
}

func init() {
	proto.RegisterFile("doxchain/oauthtwo/v1beta1/authorization_code_registry_entry.proto", fileDescriptor_1763526c68c589aa)
}

var fileDescriptor_1763526c68c589aa = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4f, 0x83, 0x40,
	0x14, 0xc7, 0x39, 0xab, 0x26, 0xc5, 0x49, 0xe2, 0x40, 0x1a, 0x73, 0x12, 0x26, 0x06, 0xbd, 0x4b,
	0x75, 0x37, 0xa1, 0x6a, 0xe2, 0xcc, 0xe8, 0x42, 0x80, 0x7b, 0x81, 0x1b, 0xe4, 0x91, 0xe3, 0x6a,
	0xc1, 0x4f, 0xe1, 0x07, 0xf1, 0x83, 0x74, 0xec, 0xe8, 0x64, 0x0c, 0x7c, 0x11, 0x43, 0xa1, 0xd4,
	0xd8, 0xed, 0xee, 0xde, 0xbd, 0xdf, 0xff, 0xf7, 0x9e, 0xe9, 0x0b, 0xac, 0x92, 0x2c, 0x92, 0x39,
	0xc7, 0x68, 0xa9, 0x33, 0xbd, 0x42, 0xfe, 0x36, 0x8f, 0x41, 0x47, 0x73, 0xde, 0xdd, 0x51, 0xc9,
	0xf7, 0x48, 0x4b, 0xcc, 0xc3, 0x04, 0x05, 0x84, 0x0a, 0x52, 0x59, 0x6a, 0x55, 0x87, 0x90, 0x6b,
	0x55, 0xb3, 0x42, 0xa1, 0x46, 0xcb, 0x8d, 0x21, 0x03, 0x85, 0x50, 0xb2, 0x1d, 0x8b, 0xed, 0x58,
	0x6c, 0x60, 0xcd, 0x2e, 0x52, 0x4c, 0x71, 0xfb, 0x9d, 0x77, 0xa7, 0xbe, 0x73, 0x46, 0xc7, 0x70,
	0x21, 0xc5, 0x98, 0x2b, 0xa4, 0xe8, 0xeb, 0xee, 0x27, 0x31, 0xa9, 0xff, 0xd7, 0xe2, 0x01, 0x05,
	0x04, 0x83, 0xc3, 0x53, 0xa7, 0x60, 0xdd, 0x9b, 0x27, 0xb8, 0xca, 0x41, 0xd9, 0xc4, 0x21, 0xde,
	0xd9, 0xad, 0xcb, 0x0e, 0x65, 0x3a, 0xde, 0xc0, 0x66, 0x8f, 0x52, 0x2c, 0x8e, 0xd7, 0xdf, 0x57,
	0x46, 0xd0, 0xb7, 0x59, 0xd7, 0xe6, 0x79, 0xf4, 0x3f, 0xc1, 0x3e, 0x72, 0x88, 0x37, 0x0d, 0x0e,
	0x0b, 0xd6, 0xa5, 0x39, 0x85, 0xaa, 0x90, 0x0a, 0x4a, 0x5f, 0xdb, 0x13, 0x87, 0x78, 0x93, 0x60,
	0xff, 0xb0, 0x78, 0x5e, 0x37, 0x94, 0x6c, 0x1a, 0x4a, 0x7e, 0x1a, 0x4a, 0x3e, 0x5a, 0x6a, 0x6c,
	0x5a, 0x6a, 0x7c, 0xb5, 0xd4, 0x78, 0x61, 0xa9, 0xd4, 0xd9, 0x32, 0x66, 0x09, 0xbe, 0xf2, 0x18,
	0x6e, 0x7a, 0x43, 0x3e, 0x4e, 0x5f, 0xed, 0x97, 0xaf, 0xeb, 0x02, 0xca, 0xf8, 0x74, 0x3b, 0xff,
	0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x38, 0xd5, 0x55, 0x9e, 0x01, 0x00, 0x00,
}

func (m *AuthorizationCodeRegistryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationCodeRegistryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizationCodeRegistryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AuthorizationCode) > 0 {
		i -= len(m.AuthorizationCode)
		copy(dAtA[i:], m.AuthorizationCode)
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(len(m.AuthorizationCode)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuthorizationCodeRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAuthorizationCodeRegistryEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthorizationCodeRegistryEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthorizationCodeRegistryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovAuthorizationCodeRegistryEntry(uint64(l))
	l = len(m.AuthorizationCode)
	if l > 0 {
		n += 1 + l + sovAuthorizationCodeRegistryEntry(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovAuthorizationCodeRegistryEntry(uint64(m.ExpiresAt))
	}
	return n
}

func sovAuthorizationCodeRegistryEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthorizationCodeRegistryEntry(x uint64) (n int) {
	return sovAuthorizationCodeRegistryEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthorizationCodeRegistryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthorizationCodeRegistryEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthorizationCodeRegistryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationCodeRegistryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizationCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthorizationCodeRegistryEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuthorizationCodeRegistryEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuthorizationCodeRegistryEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuthorizationCodeRegistryEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuthorizationCodeRegistryEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthorizationCodeRegistryEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthorizationCodeRegistryEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthorizationCodeRegistryEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthorizationCodeRegistryEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthorizationCodeRegistryEntry = fmt.Errorf("proto: unexpected end of group")
)
