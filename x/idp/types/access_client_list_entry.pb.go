// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/v1beta1/access_client_list_entry.proto

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

type AccessClientListEntry struct {
	User       types.Did `protobuf:"bytes,1,opt,name=user,proto3" json:"user"`
	CanRead    bool      `protobuf:"varint,2,opt,name=canRead,proto3" json:"canRead,omitempty"`
	CanWrite   bool      `protobuf:"varint,3,opt,name=canWrite,proto3" json:"canWrite,omitempty"`
	CanExecute bool      `protobuf:"varint,4,opt,name=canExecute,proto3" json:"canExecute,omitempty"`
}

func (m *AccessClientListEntry) Reset()         { *m = AccessClientListEntry{} }
func (m *AccessClientListEntry) String() string { return proto.CompactTextString(m) }
func (*AccessClientListEntry) ProtoMessage()    {}
func (*AccessClientListEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_80127c65a2f83c7f, []int{0}
}
func (m *AccessClientListEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessClientListEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessClientListEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessClientListEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessClientListEntry.Merge(m, src)
}
func (m *AccessClientListEntry) XXX_Size() int {
	return m.Size()
}
func (m *AccessClientListEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessClientListEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AccessClientListEntry proto.InternalMessageInfo

func (m *AccessClientListEntry) GetUser() types.Did {
	if m != nil {
		return m.User
	}
	return types.Did{}
}

func (m *AccessClientListEntry) GetCanRead() bool {
	if m != nil {
		return m.CanRead
	}
	return false
}

func (m *AccessClientListEntry) GetCanWrite() bool {
	if m != nil {
		return m.CanWrite
	}
	return false
}

func (m *AccessClientListEntry) GetCanExecute() bool {
	if m != nil {
		return m.CanExecute
	}
	return false
}

func init() {
	proto.RegisterType((*AccessClientListEntry)(nil), "doxchain.idp.v1beta1.AccessClientListEntry")
}

func init() {
	proto.RegisterFile("doxchain/idp/v1beta1/access_client_list_entry.proto", fileDescriptor_80127c65a2f83c7f)
}

var fileDescriptor_80127c65a2f83c7f = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x63, 0xa8, 0xa0, 0x32, 0x5b, 0x54, 0xa4, 0xd0, 0xc1, 0x54, 0x4c, 0x65, 0xc0, 0x56,
	0xc9, 0x09, 0x68, 0xe9, 0xc6, 0x94, 0x05, 0x89, 0x25, 0x72, 0xec, 0x5f, 0x89, 0xa5, 0x62, 0x47,
	0xb1, 0x83, 0xd2, 0x5b, 0x70, 0x06, 0x4e, 0xd3, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x45, 0x50, 0xdc,
	0x36, 0x62, 0xf3, 0xf3, 0x7b, 0xcf, 0x9f, 0x9e, 0x71, 0x2c, 0x4d, 0x23, 0x0a, 0xae, 0x34, 0x53,
	0xb2, 0x64, 0x1f, 0x8b, 0x0c, 0x1c, 0x5f, 0x30, 0x2e, 0x04, 0x58, 0x9b, 0x8a, 0x8d, 0x02, 0xed,
	0xd2, 0x8d, 0xb2, 0x2e, 0x05, 0xed, 0xaa, 0x2d, 0x2d, 0x2b, 0xe3, 0x4c, 0x38, 0x39, 0x95, 0xa8,
	0x92, 0x25, 0x3d, 0x96, 0xa6, 0x93, 0xdc, 0xe4, 0xc6, 0x07, 0x58, 0x7f, 0x3a, 0x64, 0xa7, 0x64,
	0x00, 0x48, 0x25, 0x07, 0x80, 0x54, 0xf2, 0xe0, 0xdf, 0x7d, 0x21, 0x7c, 0xfd, 0xe4, 0x71, 0x2b,
	0x4f, 0x7b, 0x51, 0xd6, 0xad, 0x7b, 0x56, 0x18, 0xe3, 0x51, 0x6d, 0xa1, 0x8a, 0xd0, 0x0c, 0xcd,
	0xaf, 0x1e, 0x6f, 0xe8, 0x00, 0xed, 0xcb, 0xc7, 0x87, 0xe8, 0xb3, 0x92, 0xcb, 0xd1, 0xee, 0xe7,
	0x36, 0x48, 0x7c, 0x38, 0x8c, 0xf0, 0xa5, 0xe0, 0x3a, 0x01, 0x2e, 0xa3, 0xb3, 0x19, 0x9a, 0x8f,
	0x93, 0x93, 0x0c, 0xa7, 0x78, 0x2c, 0xb8, 0x7e, 0xad, 0x94, 0x83, 0xe8, 0xdc, 0x5b, 0x83, 0x0e,
	0x09, 0xc6, 0x82, 0xeb, 0x75, 0x03, 0xa2, 0x76, 0x10, 0x8d, 0xbc, 0xfb, 0xef, 0x66, 0xb9, 0xda,
	0xb5, 0x04, 0xed, 0x5b, 0x82, 0x7e, 0x5b, 0x82, 0x3e, 0x3b, 0x12, 0xec, 0x3b, 0x12, 0x7c, 0x77,
	0x24, 0x78, 0xbb, 0xcf, 0x95, 0x2b, 0xea, 0x8c, 0x0a, 0xf3, 0xce, 0x32, 0x78, 0x28, 0xa0, 0x32,
	0x60, 0xd9, 0xb0, 0xb9, 0xf1, 0xdf, 0xea, 0xb6, 0x25, 0xd8, 0xec, 0xc2, 0x0f, 0x8e, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x44, 0xc1, 0x7a, 0x52, 0x73, 0x01, 0x00, 0x00,
}

func (m *AccessClientListEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessClientListEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessClientListEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanExecute {
		i--
		if m.CanExecute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.CanWrite {
		i--
		if m.CanWrite {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.CanRead {
		i--
		if m.CanRead {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAccessClientListEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAccessClientListEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessClientListEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessClientListEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.User.Size()
	n += 1 + l + sovAccessClientListEntry(uint64(l))
	if m.CanRead {
		n += 2
	}
	if m.CanWrite {
		n += 2
	}
	if m.CanExecute {
		n += 2
	}
	return n
}

func sovAccessClientListEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessClientListEntry(x uint64) (n int) {
	return sovAccessClientListEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessClientListEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessClientListEntry
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
			return fmt.Errorf("proto: AccessClientListEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessClientListEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientListEntry
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
				return ErrInvalidLengthAccessClientListEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccessClientListEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanRead", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientListEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanRead = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanWrite", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientListEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanWrite = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanExecute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessClientListEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CanExecute = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccessClientListEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessClientListEntry
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
func skipAccessClientListEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessClientListEntry
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
					return 0, ErrIntOverflowAccessClientListEntry
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
					return 0, ErrIntOverflowAccessClientListEntry
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
				return 0, ErrInvalidLengthAccessClientListEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessClientListEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessClientListEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessClientListEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessClientListEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessClientListEntry = fmt.Errorf("proto: unexpected end of group")
)
