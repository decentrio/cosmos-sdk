// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/interfaces/account_abstraction/v1/interface.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
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

// MsgAuthenticate is a message that an x/account account abstraction implementer
// must handle to authenticate a transaction. Always ensure the caller is the Accounts module.
type MsgAuthenticate struct {
	// bundler defines the address of the bundler that sent the operation.
	// NOTE: in case the operation was sent directly by the user, this field will reflect
	// the user address.
	Bundler string `protobuf:"bytes,1,opt,name=bundler,proto3" json:"bundler,omitempty"`
	// raw_tx defines the raw version of the tx, this is useful to compute the signature quickly.
	RawTx *tx.TxRaw `protobuf:"bytes,2,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// tx defines the decoded version of the tx, coming from raw_tx.
	Tx *tx.Tx `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	// signer_index defines the index of the signer in the tx.
	// Specifically this can be used to extract the signature at the correct
	// index.
	SignerIndex uint32 `protobuf:"varint,4,opt,name=signer_index,json=signerIndex,proto3" json:"signer_index,omitempty"`
}

func (m *MsgAuthenticate) Reset()         { *m = MsgAuthenticate{} }
func (m *MsgAuthenticate) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticate) ProtoMessage()    {}
func (*MsgAuthenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{0}
}
func (m *MsgAuthenticate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticate.Merge(m, src)
}
func (m *MsgAuthenticate) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticate proto.InternalMessageInfo

func (m *MsgAuthenticate) GetBundler() string {
	if m != nil {
		return m.Bundler
	}
	return ""
}

func (m *MsgAuthenticate) GetRawTx() *tx.TxRaw {
	if m != nil {
		return m.RawTx
	}
	return nil
}

func (m *MsgAuthenticate) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *MsgAuthenticate) GetSignerIndex() uint32 {
	if m != nil {
		return m.SignerIndex
	}
	return 0
}

// MsgAuthenticateResponse is the response to MsgAuthenticate.
// The authentication either fails or succeeds, this is why
// there are no auxiliary fields to the response.
type MsgAuthenticateResponse struct {
}

func (m *MsgAuthenticateResponse) Reset()         { *m = MsgAuthenticateResponse{} }
func (m *MsgAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticateResponse) ProtoMessage()    {}
func (*MsgAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{1}
}
func (m *MsgAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticateResponse.Merge(m, src)
}
func (m *MsgAuthenticateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticateResponse proto.InternalMessageInfo

// QueryAuthenticationMethods is a query that an x/account account abstraction implementer
// must handle to return the authentication methods that the account supports.
type QueryAuthenticationMethods struct {
}

func (m *QueryAuthenticationMethods) Reset()         { *m = QueryAuthenticationMethods{} }
func (m *QueryAuthenticationMethods) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethods) ProtoMessage()    {}
func (*QueryAuthenticationMethods) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{2}
}
func (m *QueryAuthenticationMethods) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethods.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethods.Merge(m, src)
}
func (m *QueryAuthenticationMethods) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethods) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethods.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethods proto.InternalMessageInfo

// QueryAuthenticationMethodsResponse is the response to QueryAuthenticationMethods.
type QueryAuthenticationMethodsResponse struct {
	// authentication_methods are the authentication methods that the account supports.
	AuthenticationMethods []string `protobuf:"bytes,1,rep,name=authentication_methods,json=authenticationMethods,proto3" json:"authentication_methods,omitempty"`
}

func (m *QueryAuthenticationMethodsResponse) Reset()         { *m = QueryAuthenticationMethodsResponse{} }
func (m *QueryAuthenticationMethodsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethodsResponse) ProtoMessage()    {}
func (*QueryAuthenticationMethodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{3}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.Merge(m, src)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethodsResponse proto.InternalMessageInfo

func (m *QueryAuthenticationMethodsResponse) GetAuthenticationMethods() []string {
	if m != nil {
		return m.AuthenticationMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAuthenticate)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticate")
	proto.RegisterType((*MsgAuthenticateResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticateResponse")
	proto.RegisterType((*QueryAuthenticationMethods)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethods")
	proto.RegisterType((*QueryAuthenticationMethodsResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethodsResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/interfaces/account_abstraction/v1/interface.proto", fileDescriptor_56b360422260e9d1)
}

var fileDescriptor_56b360422260e9d1 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x6e, 0xe2, 0x30,
	0x1c, 0xc6, 0x31, 0xdc, 0x71, 0xc2, 0xdc, 0xe9, 0xa4, 0x48, 0xdc, 0x85, 0xe8, 0x14, 0xe5, 0x22,
	0x55, 0xca, 0x64, 0x2b, 0xad, 0x3a, 0x74, 0xa4, 0x5b, 0x07, 0x86, 0xa6, 0x4c, 0xed, 0x10, 0x39,
	0x89, 0x09, 0x51, 0xc1, 0x46, 0xf6, 0x3f, 0x60, 0xde, 0xa2, 0x4f, 0xd1, 0x67, 0xe9, 0xc8, 0xd8,
	0xb1, 0x82, 0x17, 0xa9, 0x4a, 0xa0, 0xb4, 0x15, 0x1d, 0x3a, 0xe6, 0xfb, 0x7e, 0xf9, 0xe9, 0xb3,
	0x6c, 0xdc, 0x4b, 0xa5, 0x9e, 0x48, 0x4d, 0x59, 0x9a, 0xca, 0x52, 0x80, 0xa6, 0x85, 0x00, 0xae,
	0x86, 0x2c, 0xe5, 0xaf, 0x59, 0xcc, 0x12, 0x0d, 0x8a, 0xa5, 0x50, 0x48, 0x41, 0x67, 0xe1, 0x9e,
	0x20, 0x53, 0x25, 0x41, 0x5a, 0x61, 0xa5, 0x20, 0x3b, 0x05, 0xd9, 0x2b, 0xc8, 0x01, 0x05, 0x99,
	0x85, 0x4e, 0x37, 0x97, 0x32, 0x1f, 0x73, 0xba, 0x11, 0x24, 0xe5, 0x90, 0x32, 0xb1, 0xa8, 0x6c,
	0x8e, 0xb3, 0x1d, 0x04, 0x86, 0xce, 0xc2, 0x84, 0x03, 0x0b, 0x29, 0x98, 0xaa, 0xf3, 0xef, 0x11,
	0xfe, 0xdd, 0xd7, 0x79, 0xaf, 0x84, 0x11, 0x17, 0x50, 0xa4, 0x0c, 0xb8, 0x65, 0xe3, 0x1f, 0x49,
	0x29, 0xb2, 0x31, 0x57, 0x36, 0xf2, 0x50, 0xd0, 0x8a, 0x76, 0x9f, 0x16, 0xc5, 0x4d, 0xc5, 0xe6,
	0x31, 0x18, 0xbb, 0xee, 0xa1, 0xa0, 0x7d, 0x6c, 0x93, 0xed, 0x50, 0x30, 0x64, 0xab, 0x26, 0x03,
	0x13, 0xb1, 0x79, 0xf4, 0x5d, 0xb1, 0xf9, 0xc0, 0x58, 0x47, 0xb8, 0x0e, 0xc6, 0x6e, 0x6c, 0xe0,
	0xce, 0x61, 0xb8, 0x0e, 0xc6, 0xfa, 0x8f, 0x7f, 0xea, 0x22, 0x17, 0x5c, 0xc5, 0x85, 0xc8, 0xb8,
	0xb1, 0xbf, 0x79, 0x28, 0xf8, 0x15, 0xb5, 0xab, 0xec, 0xe2, 0x25, 0xf2, 0xbb, 0xf8, 0xef, 0x87,
	0x9d, 0x11, 0xd7, 0x53, 0x29, 0x34, 0xf7, 0xff, 0x61, 0xe7, 0xb2, 0xe4, 0x6a, 0xf1, 0xa6, 0x2c,
	0xa4, 0xe8, 0x73, 0x18, 0xc9, 0x4c, 0xfb, 0x37, 0xd8, 0xff, 0xbc, 0xdd, 0x39, 0xac, 0x53, 0xfc,
	0x87, 0xbd, 0x03, 0xe2, 0x49, 0x45, 0xd8, 0xc8, 0x6b, 0x04, 0xad, 0xa8, 0xc3, 0x0e, 0xfd, 0x7e,
	0x7e, 0xf5, 0xb0, 0x72, 0xd1, 0x72, 0xe5, 0xa2, 0xa7, 0x95, 0x8b, 0xee, 0xd6, 0x6e, 0x6d, 0xb9,
	0x76, 0x6b, 0x8f, 0x6b, 0xb7, 0x76, 0x7d, 0x56, 0x1d, 0x56, 0x67, 0xb7, 0xa4, 0x90, 0xd4, 0x7c,
	0xe1, 0x35, 0x24, 0xcd, 0xcd, 0xd5, 0x9c, 0x3c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x1c, 0xbd,
	0x8a, 0x49, 0x02, 0x00, 0x00,
}

func (m *MsgAuthenticate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignerIndex != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.SignerIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RawTx != nil {
		{
			size, err := m.RawTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bundler) > 0 {
		i -= len(m.Bundler)
		copy(dAtA[i:], m.Bundler)
		i = encodeVarintInterface(dAtA, i, uint64(len(m.Bundler)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthenticateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethods) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethods) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethods) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for iNdEx := len(m.AuthenticationMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthenticationMethods[iNdEx])
			copy(dAtA[i:], m.AuthenticationMethods[iNdEx])
			i = encodeVarintInterface(dAtA, i, uint64(len(m.AuthenticationMethods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterface(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthenticate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bundler)
	if l > 0 {
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.RawTx != nil {
		l = m.RawTx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.SignerIndex != 0 {
		n += 1 + sovInterface(uint64(m.SignerIndex))
	}
	return n
}

func (m *MsgAuthenticateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethods) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for _, s := range m.AuthenticationMethods {
			l = len(s)
			n += 1 + l + sovInterface(uint64(l))
		}
	}
	return n
}

func sovInterface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterface(x uint64) (n int) {
	return sovInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthenticate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RawTx == nil {
				m.RawTx = &tx.TxRaw{}
			}
			if err := m.RawTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerIndex", wireType)
			}
			m.SignerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignerIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *MsgAuthenticateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethods) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethods: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethods: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationMethods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticationMethods = append(m.AuthenticationMethods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func skipInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
				return 0, ErrInvalidLengthInterface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterface = fmt.Errorf("proto: unexpected end of group")
)
