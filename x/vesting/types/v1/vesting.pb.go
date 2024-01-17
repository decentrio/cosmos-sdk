// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/vesting/v1/vesting.proto

package v1

import (
	_ "cosmossdk.io/x/auth/types"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Period defines a length of time and amount of coins that will vest.
type Period struct {
	// Period duration in seconds.
	Length int64                                    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_56cb610908811f92, []int{0}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Period) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*Period)(nil), "cosmos.accounts.vesting.v1.Period")
}

func init() {
	proto.RegisterFile("cosmos/accounts/vesting/v1/vesting.proto", fileDescriptor_56cb610908811f92)
}

var fileDescriptor_56cb610908811f92 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x8d, 0xa9, 0x94, 0x21, 0xb0, 0x50, 0x21, 0x54, 0x3a, 0xb8, 0x15, 0x53, 0x54, 0x09, 0x5b,
	0x81, 0x1b, 0x14, 0x89, 0x19, 0x75, 0x64, 0x41, 0xae, 0x6b, 0x39, 0x56, 0x1b, 0xff, 0xaa, 0x76,
	0x22, 0x72, 0x0b, 0x66, 0xb8, 0x00, 0x62, 0xea, 0x31, 0x3a, 0x76, 0x64, 0x02, 0x94, 0x0c, 0xbd,
	0x06, 0x72, 0xe2, 0x30, 0xc0, 0x62, 0xbf, 0xef, 0xf7, 0xfc, 0xdf, 0xfb, 0x3f, 0x8a, 0x39, 0x98,
	0x0c, 0x0c, 0x65, 0x9c, 0x43, 0xae, 0xad, 0xa1, 0x85, 0x30, 0x56, 0x69, 0x49, 0x8b, 0xa4, 0x83,
	0x64, 0xbd, 0x01, 0x0b, 0xfd, 0x61, 0xab, 0x24, 0x9d, 0x92, 0x74, 0x74, 0x91, 0x0c, 0x4f, 0x59,
	0xa6, 0x34, 0xd0, 0xe6, 0x6c, 0xe5, 0xc3, 0x33, 0x09, 0x12, 0x1a, 0x48, 0x1d, 0xf2, 0xaf, 0xd8,
	0xdb, 0xcd, 0x99, 0x11, 0xb4, 0x48, 0xe6, 0xc2, 0xb2, 0x84, 0x72, 0x50, 0xfa, 0x0f, 0xcf, 0x72,
	0x9b, 0xfe, 0xf2, 0xae, 0x68, 0xf9, 0xcb, 0x57, 0x14, 0x85, 0xf7, 0x62, 0xa3, 0x60, 0xd1, 0x3f,
	0x8f, 0xc2, 0x95, 0xd0, 0xd2, 0xa6, 0x03, 0x34, 0x46, 0x71, 0x6f, 0xe6, 0xab, 0x7e, 0x19, 0x85,
	0x2c, 0x73, 0x09, 0x07, 0x47, 0xe3, 0x5e, 0x7c, 0x7c, 0x7d, 0x41, 0x7c, 0x70, 0xe7, 0x49, 0x7c,
	0x4f, 0x72, 0x0b, 0x4a, 0x4f, 0xef, 0x76, 0x9f, 0xa3, 0xe0, 0xfd, 0x6b, 0x14, 0x4b, 0x65, 0xd3,
	0x7c, 0x4e, 0x38, 0x64, 0xd4, 0x07, 0x68, 0xaf, 0x2b, 0xb3, 0x58, 0x52, 0x5b, 0xae, 0x85, 0x69,
	0x3e, 0x98, 0x97, 0xc3, 0x76, 0x72, 0xb2, 0x12, 0x92, 0xf1, 0xf2, 0xd1, 0xa5, 0x36, 0x6f, 0x87,
	0xed, 0x04, 0xcd, 0xbc, 0xe1, 0x74, 0xba, 0xab, 0x30, 0xda, 0x57, 0x18, 0x7d, 0x57, 0x18, 0x3d,
	0xd7, 0x38, 0xd8, 0xd7, 0x38, 0xf8, 0xa8, 0x71, 0xf0, 0xe0, 0xd7, 0x6c, 0x16, 0x4b, 0xa2, 0x80,
	0x3e, 0xfd, 0x5f, 0x77, 0x63, 0xe2, 0x26, 0x0e, 0x9b, 0x41, 0x6f, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0xc0, 0x3b, 0xdf, 0x99, 0x01, 0x00, 0x00,
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Length != 0 {
		i = encodeVarintVesting(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Length != 0 {
		n += 1 + sovVesting(uint64(m.Length))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
