// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/vesting/v1/query.proto

package v1

import (
	_ "cosmossdk.io/x/auth/vesting/types"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// QueryOriginalVestingRequest is used to query the account original vesting.
type QueryOriginalVestingRequest struct {
}

func (m *QueryOriginalVestingRequest) Reset()         { *m = QueryOriginalVestingRequest{} }
func (m *QueryOriginalVestingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOriginalVestingRequest) ProtoMessage()    {}
func (*QueryOriginalVestingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{0}
}
func (m *QueryOriginalVestingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOriginalVestingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOriginalVestingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOriginalVestingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOriginalVestingRequest.Merge(m, src)
}
func (m *QueryOriginalVestingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryOriginalVestingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOriginalVestingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOriginalVestingRequest proto.InternalMessageInfo

// QueryOriginalVestingResponse returns the account original vesting.
type QueryOriginalVestingResponse struct {
	// original_vesting defines the value of the account original vesting.
	OriginalVesting github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=original_vesting,json=originalVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"original_vesting"`
}

func (m *QueryOriginalVestingResponse) Reset()         { *m = QueryOriginalVestingResponse{} }
func (m *QueryOriginalVestingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOriginalVestingResponse) ProtoMessage()    {}
func (*QueryOriginalVestingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{1}
}
func (m *QueryOriginalVestingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOriginalVestingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOriginalVestingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOriginalVestingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOriginalVestingResponse.Merge(m, src)
}
func (m *QueryOriginalVestingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOriginalVestingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOriginalVestingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOriginalVestingResponse proto.InternalMessageInfo

func (m *QueryOriginalVestingResponse) GetOriginalVesting() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.OriginalVesting
	}
	return nil
}

// QueryDelegatedFreeRequest is used to query the account delegated free amount.
type QueryDelegatedFreeRequest struct {
}

func (m *QueryDelegatedFreeRequest) Reset()         { *m = QueryDelegatedFreeRequest{} }
func (m *QueryDelegatedFreeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedFreeRequest) ProtoMessage()    {}
func (*QueryDelegatedFreeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{2}
}
func (m *QueryDelegatedFreeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedFreeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedFreeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedFreeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedFreeRequest.Merge(m, src)
}
func (m *QueryDelegatedFreeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedFreeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedFreeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedFreeRequest proto.InternalMessageInfo

// QueryDelegatedFreeResponse returns the account delegated free amount.
type QueryDelegatedFreeResponse struct {
	// delegated_free defines the value of the account delegated free.
	DelegatedFree github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=delegated_free,json=delegatedFree,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_free"`
}

func (m *QueryDelegatedFreeResponse) Reset()         { *m = QueryDelegatedFreeResponse{} }
func (m *QueryDelegatedFreeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedFreeResponse) ProtoMessage()    {}
func (*QueryDelegatedFreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{3}
}
func (m *QueryDelegatedFreeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedFreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedFreeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedFreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedFreeResponse.Merge(m, src)
}
func (m *QueryDelegatedFreeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedFreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedFreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedFreeResponse proto.InternalMessageInfo

func (m *QueryDelegatedFreeResponse) GetDelegatedFree() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DelegatedFree
	}
	return nil
}

// QueryDelegatedVestingRequest is used to query the account delegated vesting amount.
type QueryDelegatedVestingRequest struct {
}

func (m *QueryDelegatedVestingRequest) Reset()         { *m = QueryDelegatedVestingRequest{} }
func (m *QueryDelegatedVestingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedVestingRequest) ProtoMessage()    {}
func (*QueryDelegatedVestingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{4}
}
func (m *QueryDelegatedVestingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedVestingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedVestingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedVestingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedVestingRequest.Merge(m, src)
}
func (m *QueryDelegatedVestingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedVestingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedVestingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedVestingRequest proto.InternalMessageInfo

// QueryDelegatedVestingResponse returns the account delegated vesting amount.
type QueryDelegatedVestingResponse struct {
	// delegated_vesting defines the the account delegated vesting.
	DelegatedVesting github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=delegated_vesting,json=delegatedVesting,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"delegated_vesting"`
}

func (m *QueryDelegatedVestingResponse) Reset()         { *m = QueryDelegatedVestingResponse{} }
func (m *QueryDelegatedVestingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedVestingResponse) ProtoMessage()    {}
func (*QueryDelegatedVestingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{5}
}
func (m *QueryDelegatedVestingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedVestingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedVestingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedVestingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedVestingResponse.Merge(m, src)
}
func (m *QueryDelegatedVestingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedVestingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedVestingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedVestingResponse proto.InternalMessageInfo

func (m *QueryDelegatedVestingResponse) GetDelegatedVesting() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DelegatedVesting
	}
	return nil
}

// QueryCounterRequest is used to query the account vesting end time.
type QueryEndTimeRequest struct {
}

func (m *QueryEndTimeRequest) Reset()         { *m = QueryEndTimeRequest{} }
func (m *QueryEndTimeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEndTimeRequest) ProtoMessage()    {}
func (*QueryEndTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{6}
}
func (m *QueryEndTimeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndTimeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndTimeRequest.Merge(m, src)
}
func (m *QueryEndTimeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndTimeRequest proto.InternalMessageInfo

// QueryEndTimeResponse returns the account vesting end time.
type QueryEndTimeResponse struct {
	// end_time defines the value of the account vesting end time.
	EndTime int64 `protobuf:"varint,1,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (m *QueryEndTimeResponse) Reset()         { *m = QueryEndTimeResponse{} }
func (m *QueryEndTimeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEndTimeResponse) ProtoMessage()    {}
func (*QueryEndTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f809c257a71960c, []int{7}
}
func (m *QueryEndTimeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEndTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEndTimeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEndTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEndTimeResponse.Merge(m, src)
}
func (m *QueryEndTimeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEndTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEndTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEndTimeResponse proto.InternalMessageInfo

func (m *QueryEndTimeResponse) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryOriginalVestingRequest)(nil), "cosmos.accounts.vesting.v1.QueryOriginalVestingRequest")
	proto.RegisterType((*QueryOriginalVestingResponse)(nil), "cosmos.accounts.vesting.v1.QueryOriginalVestingResponse")
	proto.RegisterType((*QueryDelegatedFreeRequest)(nil), "cosmos.accounts.vesting.v1.QueryDelegatedFreeRequest")
	proto.RegisterType((*QueryDelegatedFreeResponse)(nil), "cosmos.accounts.vesting.v1.QueryDelegatedFreeResponse")
	proto.RegisterType((*QueryDelegatedVestingRequest)(nil), "cosmos.accounts.vesting.v1.QueryDelegatedVestingRequest")
	proto.RegisterType((*QueryDelegatedVestingResponse)(nil), "cosmos.accounts.vesting.v1.QueryDelegatedVestingResponse")
	proto.RegisterType((*QueryEndTimeRequest)(nil), "cosmos.accounts.vesting.v1.QueryEndTimeRequest")
	proto.RegisterType((*QueryEndTimeResponse)(nil), "cosmos.accounts.vesting.v1.QueryEndTimeResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/vesting/v1/query.proto", fileDescriptor_5f809c257a71960c)
}

var fileDescriptor_5f809c257a71960c = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6e, 0xd4, 0x30,
	0x14, 0xc6, 0xc7, 0xaa, 0x04, 0xc8, 0xfc, 0x6b, 0x4b, 0x11, 0xcc, 0x94, 0xba, 0x28, 0x42, 0x28,
	0xaa, 0x84, 0xad, 0xc0, 0x0d, 0x06, 0xe8, 0x16, 0x31, 0x42, 0x2c, 0xd8, 0x44, 0x99, 0xf8, 0x11,
	0xac, 0x36, 0xf6, 0x34, 0x76, 0x22, 0xe6, 0x06, 0x2c, 0x10, 0x62, 0xcd, 0x05, 0x40, 0x48, 0x48,
	0x73, 0x8c, 0x2e, 0xbb, 0x64, 0x05, 0x68, 0x66, 0xd1, 0x6b, 0xa0, 0xd8, 0x2f, 0x2d, 0x2d, 0xed,
	0xb6, 0x9b, 0x99, 0xe4, 0x7d, 0x9f, 0xfd, 0x7e, 0x9f, 0x5f, 0x4c, 0x1f, 0xe6, 0xc6, 0x96, 0xc6,
	0x8a, 0x2c, 0xcf, 0x4d, 0xad, 0x9d, 0x15, 0x0d, 0x58, 0xa7, 0x74, 0x21, 0x9a, 0x44, 0xec, 0xd5,
	0x50, 0x4d, 0xf9, 0xa4, 0x32, 0xce, 0xac, 0x0e, 0x82, 0x8f, 0x77, 0x3e, 0x8e, 0x3e, 0xde, 0x24,
	0x83, 0x95, 0xac, 0x54, 0xda, 0x08, 0xff, 0x1b, 0xec, 0x03, 0x86, 0xdb, 0x8e, 0x33, 0x0b, 0xa2,
	0x49, 0xc6, 0xe0, 0xb2, 0x44, 0xe4, 0x46, 0x69, 0xd4, 0xef, 0xa0, 0x5e, 0x5a, 0xdf, 0xa9, 0xb4,
	0x05, 0x0a, 0x0f, 0x50, 0x38, 0xc6, 0x08, 0x6b, 0xbb, 0x76, 0xc1, 0xd5, 0x0f, 0xae, 0xd4, 0xbf,
	0x09, 0x44, 0x0b, 0xd2, 0x5a, 0x61, 0x0a, 0x13, 0xea, 0xed, 0x53, 0xa8, 0x46, 0x1b, 0x74, 0xfd,
	0x65, 0x9b, 0xe6, 0x45, 0xa5, 0x0a, 0xa5, 0xb3, 0xdd, 0xd7, 0x61, 0xbb, 0x11, 0xec, 0xd5, 0x60,
	0x5d, 0xf4, 0x83, 0xd0, 0x7b, 0x67, 0xeb, 0x76, 0x62, 0xb4, 0x85, 0xd5, 0x8f, 0x84, 0x2e, 0x1b,
	0xd4, 0x52, 0x64, 0xb9, 0x4b, 0xee, 0x2f, 0xc5, 0x57, 0x1f, 0xf7, 0x39, 0xf6, 0x6f, 0xb3, 0x72,
	0xe4, 0xe5, 0x4f, 0x8d, 0xd2, 0xc3, 0xed, 0xfd, 0x5f, 0x9b, 0xbd, 0xef, 0xbf, 0x37, 0xe3, 0x42,
	0xb9, 0x77, 0xf5, 0x98, 0xe7, 0xa6, 0x44, 0x58, 0xfc, 0x7b, 0x64, 0xe5, 0x8e, 0x70, 0xd3, 0x09,
	0x58, 0xbf, 0xc0, 0x7e, 0x39, 0x9c, 0x6d, 0x5d, 0xdb, 0x85, 0x22, 0xcb, 0xa7, 0x69, 0x7b, 0x5a,
	0xf6, 0xdb, 0xe1, 0x6c, 0x8b, 0x8c, 0x6e, 0x9a, 0x93, 0x58, 0xd1, 0x3a, 0xed, 0x7b, 0xdc, 0x67,
	0xd0, 0x9a, 0x1d, 0xc8, 0xed, 0x0a, 0xa0, 0x0b, 0xf3, 0x95, 0xd0, 0xc1, 0x59, 0x2a, 0x46, 0xf9,
	0x40, 0xe8, 0x0d, 0xd9, 0x29, 0xe9, 0xdb, 0x0a, 0xe0, 0xe2, 0x82, 0x5c, 0x97, 0xff, 0x22, 0x45,
	0x0c, 0x4f, 0xfd, 0x08, 0xf4, 0xd4, 0x58, 0x66, 0x84, 0x6e, 0x9c, 0x63, 0xc0, 0x30, 0x9f, 0x08,
	0x5d, 0x39, 0x0e, 0x73, 0xe1, 0x83, 0x59, 0x96, 0xa7, 0xc0, 0xa2, 0xdb, 0xf4, 0x96, 0x27, 0x7e,
	0xae, 0xe5, 0x2b, 0x55, 0x1e, 0xcd, 0x24, 0xa1, 0x6b, 0x27, 0xcb, 0xc8, 0xdf, 0xa7, 0x57, 0x40,
	0xcb, 0xd4, 0xa9, 0xb2, 0x9d, 0x02, 0x89, 0x97, 0x46, 0x97, 0x21, 0x58, 0x86, 0xc3, 0xfd, 0x39,
	0x23, 0x07, 0x73, 0x46, 0xfe, 0xcc, 0x19, 0xf9, 0xbc, 0x60, 0xbd, 0x83, 0x05, 0xeb, 0xfd, 0x5c,
	0xb0, 0xde, 0x9b, 0x38, 0x30, 0x5a, 0xb9, 0xc3, 0x95, 0x11, 0xef, 0xff, 0xbf, 0xbb, 0x1e, 0xbc,
	0xbd, 0x3a, 0x97, 0xfc, 0xd7, 0xff, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xb1, 0xf6,
	0xa3, 0xe6, 0x03, 0x00, 0x00,
}

func (m *QueryOriginalVestingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOriginalVestingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOriginalVestingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryOriginalVestingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOriginalVestingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOriginalVestingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OriginalVesting) > 0 {
		for iNdEx := len(m.OriginalVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OriginalVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedFreeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedFreeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedFreeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedFreeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedFreeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedFreeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatedFree) > 0 {
		for iNdEx := len(m.DelegatedFree) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedFree[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedVestingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedVestingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedVestingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedVestingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedVestingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedVestingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DelegatedVesting) > 0 {
		for iNdEx := len(m.DelegatedVesting) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatedVesting[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryEndTimeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndTimeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndTimeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEndTimeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEndTimeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEndTimeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryOriginalVestingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryOriginalVestingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.OriginalVesting) > 0 {
		for _, e := range m.OriginalVesting {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDelegatedFreeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryDelegatedFreeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DelegatedFree) > 0 {
		for _, e := range m.DelegatedFree {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDelegatedVestingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryDelegatedVestingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DelegatedVesting) > 0 {
		for _, e := range m.DelegatedVesting {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryEndTimeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEndTimeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndTime != 0 {
		n += 1 + sovQuery(uint64(m.EndTime))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryOriginalVestingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOriginalVestingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOriginalVestingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryOriginalVestingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOriginalVestingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOriginalVestingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalVesting = append(m.OriginalVesting, types.Coin{})
			if err := m.OriginalVesting[len(m.OriginalVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatedFreeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatedFreeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedFreeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatedFreeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatedFreeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedFreeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedFree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedFree = append(m.DelegatedFree, types.Coin{})
			if err := m.DelegatedFree[len(m.DelegatedFree)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatedVestingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatedVestingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedVestingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDelegatedVestingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDelegatedVestingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedVestingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatedVesting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatedVesting = append(m.DelegatedVesting, types.Coin{})
			if err := m.DelegatedVesting[len(m.DelegatedVesting)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEndTimeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEndTimeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndTimeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEndTimeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEndTimeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEndTimeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
