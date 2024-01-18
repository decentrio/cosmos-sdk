// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/vesting/v1/tx.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// MsgInitVestingAccount defines a message that enables creating a vesting
// account.
type MsgInitVestingAccount struct {
	Owner  string                                   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	// end of vesting as unix time (in seconds).
	EndTime int64 `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Delayed bool  `protobuf:"varint,4,opt,name=delayed,proto3" json:"delayed,omitempty"`
	// start of vesting as unix time (in seconds).
	//
	// Since 0.51.x
	StartTime int64 `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (m *MsgInitVestingAccount) Reset()         { *m = MsgInitVestingAccount{} }
func (m *MsgInitVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgInitVestingAccount) ProtoMessage()    {}
func (*MsgInitVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b8f3ca3aa753c3, []int{0}
}
func (m *MsgInitVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitVestingAccount.Merge(m, src)
}
func (m *MsgInitVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitVestingAccount proto.InternalMessageInfo

func (m *MsgInitVestingAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgInitVestingAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgInitVestingAccount) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *MsgInitVestingAccount) GetDelayed() bool {
	if m != nil {
		return m.Delayed
	}
	return false
}

func (m *MsgInitVestingAccount) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

// MsgInitVestingAccountResponse defines the Msg/InitVestingAccount response type.
type MsgInitVestingAccountResponse struct {
}

func (m *MsgInitVestingAccountResponse) Reset()         { *m = MsgInitVestingAccountResponse{} }
func (m *MsgInitVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitVestingAccountResponse) ProtoMessage()    {}
func (*MsgInitVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b8f3ca3aa753c3, []int{1}
}
func (m *MsgInitVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitVestingAccountResponse.Merge(m, src)
}
func (m *MsgInitVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitVestingAccountResponse proto.InternalMessageInfo

// MsgInitVestingAccount defines a message that enables creating a vesting
// account.
type MsgInitPeriodicVestingAccount struct {
	// start of vesting as unix time (in seconds).
	StartTime      int64    `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods []Period `protobuf:"bytes,2,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods"`
}

func (m *MsgInitPeriodicVestingAccount) Reset()         { *m = MsgInitPeriodicVestingAccount{} }
func (m *MsgInitPeriodicVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgInitPeriodicVestingAccount) ProtoMessage()    {}
func (*MsgInitPeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b8f3ca3aa753c3, []int{2}
}
func (m *MsgInitPeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitPeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitPeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitPeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitPeriodicVestingAccount.Merge(m, src)
}
func (m *MsgInitPeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitPeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitPeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitPeriodicVestingAccount proto.InternalMessageInfo

func (m *MsgInitPeriodicVestingAccount) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *MsgInitPeriodicVestingAccount) GetVestingPeriods() []Period {
	if m != nil {
		return m.VestingPeriods
	}
	return nil
}

// MsgInitVestingAccountResponse defines the Msg/InitPeriodicVestingAccount
// response type.
type MsgInitPeriodicVestingAccountResponse struct {
}

func (m *MsgInitPeriodicVestingAccountResponse) Reset()         { *m = MsgInitPeriodicVestingAccountResponse{} }
func (m *MsgInitPeriodicVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitPeriodicVestingAccountResponse) ProtoMessage()    {}
func (*MsgInitPeriodicVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3b8f3ca3aa753c3, []int{3}
}
func (m *MsgInitPeriodicVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitPeriodicVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitPeriodicVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitPeriodicVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitPeriodicVestingAccountResponse.Merge(m, src)
}
func (m *MsgInitPeriodicVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitPeriodicVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitPeriodicVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitPeriodicVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInitVestingAccount)(nil), "cosmos.accounts.vesting.v1.MsgInitVestingAccount")
	proto.RegisterType((*MsgInitVestingAccountResponse)(nil), "cosmos.accounts.vesting.v1.MsgInitVestingAccountResponse")
	proto.RegisterType((*MsgInitPeriodicVestingAccount)(nil), "cosmos.accounts.vesting.v1.MsgInitPeriodicVestingAccount")
	proto.RegisterType((*MsgInitPeriodicVestingAccountResponse)(nil), "cosmos.accounts.vesting.v1.MsgInitPeriodicVestingAccountResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/vesting/v1/tx.proto", fileDescriptor_e3b8f3ca3aa753c3)
}

var fileDescriptor_e3b8f3ca3aa753c3 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x35, 0xf4, 0x47, 0x0e, 0x04, 0xc2, 0x2a, 0x92, 0x13, 0xa9, 0xb6, 0x95, 0x0a, 0x61,
	0x22, 0x7a, 0xa7, 0xc0, 0x44, 0xb7, 0x06, 0x09, 0x89, 0x01, 0x09, 0x19, 0xd4, 0x81, 0x25, 0x72,
	0xec, 0x93, 0x39, 0xb5, 0xbe, 0x8b, 0xfc, 0xae, 0xa1, 0xf9, 0x17, 0x98, 0x98, 0x99, 0x3a, 0x22,
	0xa6, 0x0c, 0xfc, 0x03, 0x48, 0x0c, 0x1d, 0x2b, 0x26, 0x26, 0x40, 0xc9, 0x90, 0xfe, 0x19, 0xe8,
	0x7c, 0xe7, 0x0a, 0x4a, 0xc9, 0xe2, 0x1f, 0xf7, 0x7d, 0xdf, 0x7b, 0xef, 0xfb, 0xee, 0xe1, 0xed,
	0x54, 0x42, 0x21, 0x81, 0x26, 0x69, 0x2a, 0x8f, 0x84, 0x02, 0x3a, 0x66, 0xa0, 0xb8, 0xc8, 0xe9,
	0xb8, 0x47, 0xd5, 0x31, 0x19, 0x95, 0x52, 0x49, 0xb7, 0x6d, 0x48, 0xa4, 0x26, 0x11, 0x4b, 0x22,
	0xe3, 0x5e, 0x7b, 0x33, 0x97, 0xb9, 0xac, 0x68, 0x54, 0x7f, 0x19, 0x45, 0xdb, 0xb7, 0x65, 0x87,
	0x09, 0x30, 0x3a, 0xee, 0x0d, 0x99, 0x4a, 0x7a, 0x34, 0x95, 0x5c, 0x58, 0xbc, 0x65, 0xf0, 0x81,
	0x11, 0xda, 0xf2, 0x06, 0x8a, 0x96, 0x4c, 0x54, 0xf7, 0x35, 0xcc, 0xdb, 0x49, 0xc1, 0x85, 0xa4,
	0xd5, 0xd3, 0x1c, 0x75, 0xbe, 0xac, 0xe0, 0x3b, 0xcf, 0x21, 0x7f, 0x26, 0xb8, 0xda, 0x37, 0xdc,
	0x3d, 0x53, 0xc6, 0x25, 0x78, 0x55, 0xbe, 0x15, 0xac, 0xf4, 0x50, 0x88, 0xa2, 0x66, 0xdf, 0xfb,
	0xf6, 0x79, 0x67, 0xd3, 0xf6, 0xdd, 0xcb, 0xb2, 0x92, 0x01, 0xbc, 0x54, 0x25, 0x17, 0x79, 0x6c,
	0x68, 0xee, 0x04, 0xaf, 0x25, 0x85, 0x56, 0x7a, 0x2b, 0x61, 0x23, 0xba, 0xfe, 0xb0, 0x45, 0x2c,
	0x5b, 0x5b, 0x22, 0xd6, 0x12, 0x79, 0x22, 0xb9, 0xe8, 0x3f, 0x3d, 0xfd, 0x11, 0x38, 0x9f, 0x7e,
	0x06, 0x51, 0xce, 0xd5, 0x9b, 0xa3, 0x21, 0x49, 0x65, 0x61, 0x2d, 0xd9, 0xd7, 0x0e, 0x64, 0x07,
	0x54, 0x4d, 0x46, 0x0c, 0x2a, 0x01, 0x7c, 0x58, 0x4c, 0xbb, 0x37, 0x0e, 0x59, 0x9e, 0xa4, 0x93,
	0x81, 0x0e, 0x05, 0x3e, 0x2e, 0xa6, 0x5d, 0x14, 0xdb, 0x86, 0x6e, 0x0b, 0x6f, 0x30, 0x91, 0x0d,
	0x14, 0x2f, 0x98, 0xd7, 0x08, 0x51, 0xd4, 0x88, 0xd7, 0x99, 0xc8, 0x5e, 0xf1, 0x82, 0xb9, 0x1e,
	0x5e, 0xcf, 0xd8, 0x61, 0x32, 0x61, 0x99, 0x77, 0x2d, 0x44, 0xd1, 0x46, 0x5c, 0xff, 0xba, 0x5b,
	0x18, 0x83, 0x4a, 0x4a, 0x65, 0x64, 0xab, 0x95, 0xac, 0x59, 0x9d, 0x68, 0xe1, 0xee, 0xfd, 0xf3,
	0x93, 0x00, 0xbd, 0x5b, 0x4c, 0xbb, 0xe1, 0x1f, 0x23, 0x5d, 0x99, 0x54, 0x27, 0xc0, 0x5b, 0x57,
	0x02, 0x31, 0x83, 0x91, 0x14, 0xc0, 0x3a, 0x5f, 0xd1, 0x05, 0xe3, 0x05, 0x2b, 0xb9, 0xcc, 0x78,
	0x7a, 0x29, 0xec, 0xbf, 0x87, 0x41, 0x97, 0x86, 0x71, 0xf7, 0xf1, 0x2d, 0x7b, 0x93, 0x83, 0x51,
	0x55, 0x00, 0x6c, 0xc8, 0x1d, 0xf2, 0xff, 0x4d, 0x23, 0xa6, 0x57, 0xbf, 0xa9, 0xd3, 0x36, 0x81,
	0xdd, 0xb4, 0xa8, 0x41, 0x60, 0xf7, 0xc1, 0xf9, 0x49, 0xe0, 0x68, 0x93, 0xdb, 0xff, 0x9a, 0x34,
	0x1c, 0x3d, 0x67, 0xed, 0xf3, 0x1e, 0xbe, 0xbb, 0xd4, 0x45, 0xed, 0xb7, 0xff, 0xf8, 0x74, 0xe6,
	0xa3, 0xb3, 0x99, 0x8f, 0x7e, 0xcd, 0x7c, 0xf4, 0x7e, 0xee, 0x3b, 0x67, 0x73, 0xdf, 0xf9, 0x3e,
	0xf7, 0x9d, 0xd7, 0x81, 0xe9, 0x03, 0xd9, 0x01, 0xe1, 0x92, 0x1e, 0x5f, 0xac, 0x6a, 0x75, 0xd7,
	0x7a, 0xef, 0xd7, 0xaa, 0xb5, 0x7c, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x3d, 0xcc, 0x17,
	0x67, 0x03, 0x00, 0x00,
}

func (this *MsgInitVestingAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgInitVestingAccount)
	if !ok {
		that2, ok := that.(MsgInitVestingAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Delayed != that1.Delayed {
		return false
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	return true
}
func (m *MsgInitVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x28
	}
	if m.Delayed {
		i--
		if m.Delayed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.EndTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgInitPeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitPeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitPeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitPeriodicVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitPeriodicVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitPeriodicVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInitVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovTx(uint64(m.EndTime))
	}
	if m.Delayed {
		n += 2
	}
	if m.StartTime != 0 {
		n += 1 + sovTx(uint64(m.StartTime))
	}
	return n
}

func (m *MsgInitVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgInitPeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != 0 {
		n += 1 + sovTx(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgInitPeriodicVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delayed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Delayed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitPeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitPeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitPeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitPeriodicVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitPeriodicVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitPeriodicVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
