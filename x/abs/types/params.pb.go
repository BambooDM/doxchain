// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/abs/v1beta1/params.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/did/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Params struct {
	Operators         []types.Did `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators"`
	BlockExpireOffset uint64      `protobuf:"varint,2,opt,name=blockExpireOffset,proto3" json:"blockExpireOffset,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5ba7663fa1ce94, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetOperators() []types.Did {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *Params) GetBlockExpireOffset() uint64 {
	if m != nil {
		return m.BlockExpireOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "beheroes.doxchain.abs.v1beta1.Params")
}

func init() { proto.RegisterFile("doxchain/abs/v1beta1/params.proto", fileDescriptor_2f5ba7663fa1ce94) }

var fileDescriptor_2f5ba7663fa1ce94 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xc9, 0xaf, 0x48,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4c, 0x2a, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4d,
	0x4a, 0xcd, 0x48, 0x2d, 0xca, 0x4f, 0x2d, 0xd6, 0x83, 0xa9, 0xd5, 0x4b, 0x4c, 0x2a, 0xd6, 0x83,
	0xaa, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd4, 0x07, 0xb1, 0x20, 0x9a, 0xa4, 0xe4,
	0xe0, 0xe6, 0xa6, 0x64, 0xa6, 0xc0, 0xcd, 0x4d, 0xc9, 0x4c, 0x81, 0xc8, 0x2b, 0xb5, 0x30, 0x72,
	0xb1, 0x05, 0x80, 0x6d, 0x11, 0x72, 0xe3, 0xe2, 0xcc, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f,
	0x2a, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd2, 0xc3, 0xb4, 0x13, 0xa4, 0x17, 0x6a,
	0x8e, 0x9e, 0x4b, 0x66, 0x8a, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x08, 0xad, 0x42, 0x3a,
	0x5c, 0x82, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xae, 0x15, 0x05, 0x99, 0x45, 0xa9, 0xfe, 0x69, 0x69,
	0xc5, 0xa9, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x98, 0x12, 0x56, 0x2c, 0x33, 0x16,
	0xc8, 0x33, 0x38, 0x39, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x66,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x52, 0xaa, 0x2e, 0xc4, 0x35,
	0xfa, 0x70, 0x5f, 0x55, 0x80, 0xc3, 0xab, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x25,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x7d, 0xab, 0xf5, 0x4c, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockExpireOffset != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlockExpireOffset))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Operators) > 0 {
		for iNdEx := len(m.Operators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Operators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Operators) > 0 {
		for _, e := range m.Operators {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.BlockExpireOffset != 0 {
		n += 1 + sovParams(uint64(m.BlockExpireOffset))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operators = append(m.Operators, types.Did{})
			if err := m.Operators[len(m.Operators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockExpireOffset", wireType)
			}
			m.BlockExpireOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockExpireOffset |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
