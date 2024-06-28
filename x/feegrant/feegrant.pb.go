// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/feegrant.proto

package feegrant

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BasicAllowance implements Allowance with a one-time grant of coins
// that optionally expires. The grantee can use up to SpendLimit to cover fees.
type BasicAllowance struct {
	// spend_limit specifies the maximum amount of coins that can be spent
	// by this allowance and will be updated as coins are spent. If it is
	// empty, there is no spend limit and any amount of coins can be spent.
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit"`
	// expiration specifies an optional time when this allowance expires
	Expiration *time.Time `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
}

func (m *BasicAllowance) Reset()         { *m = BasicAllowance{} }
func (m *BasicAllowance) String() string { return proto.CompactTextString(m) }
func (*BasicAllowance) ProtoMessage()    {}
func (*BasicAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{0}
}
func (m *BasicAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicAllowance.Merge(m, src)
}
func (m *BasicAllowance) XXX_Size() int {
	return m.Size()
}
func (m *BasicAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_BasicAllowance proto.InternalMessageInfo

func (m *BasicAllowance) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *BasicAllowance) GetExpiration() *time.Time {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// PeriodicAllowance extends Allowance to allow for both a maximum cap,
// as well as a limit per time period.
type PeriodicAllowance struct {
	// basic specifies a struct of `BasicAllowance`
	Basic BasicAllowance `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic"`
	// period specifies the time duration in which period_spend_limit coins can
	// be spent before that allowance is reset
	Period time.Duration `protobuf:"bytes,2,opt,name=period,proto3,stdduration" json:"period"`
	// period_spend_limit specifies the maximum number of coins that can be spent
	// in the period
	PeriodSpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=period_spend_limit,json=periodSpendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"period_spend_limit"`
	// period_can_spend is the number of coins left to be spent before the period_reset time
	PeriodCanSpend github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=period_can_spend,json=periodCanSpend,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"period_can_spend"`
	// period_reset is the time at which this period resets and a new one begins,
	// it is calculated from the start time of the first transaction after the
	// last period ended
	PeriodReset time.Time `protobuf:"bytes,5,opt,name=period_reset,json=periodReset,proto3,stdtime" json:"period_reset"`
}

func (m *PeriodicAllowance) Reset()         { *m = PeriodicAllowance{} }
func (m *PeriodicAllowance) String() string { return proto.CompactTextString(m) }
func (*PeriodicAllowance) ProtoMessage()    {}
func (*PeriodicAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{1}
}
func (m *PeriodicAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicAllowance.Merge(m, src)
}
func (m *PeriodicAllowance) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicAllowance proto.InternalMessageInfo

func (m *PeriodicAllowance) GetBasic() BasicAllowance {
	if m != nil {
		return m.Basic
	}
	return BasicAllowance{}
}

func (m *PeriodicAllowance) GetPeriod() time.Duration {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *PeriodicAllowance) GetPeriodSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodSpendLimit
	}
	return nil
}

func (m *PeriodicAllowance) GetPeriodCanSpend() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodCanSpend
	}
	return nil
}

func (m *PeriodicAllowance) GetPeriodReset() time.Time {
	if m != nil {
		return m.PeriodReset
	}
	return time.Time{}
}

// AllowedMsgAllowance creates allowance only for specified message types.
type AllowedMsgAllowance struct {
	// allowance can be any of basic and periodic fee allowance.
	Allowance *types1.Any `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
	// allowed_messages are the messages for which the grantee has the access.
	AllowedMessages []string `protobuf:"bytes,2,rep,name=allowed_messages,json=allowedMessages,proto3" json:"allowed_messages,omitempty"`
}

func (m *AllowedMsgAllowance) Reset()         { *m = AllowedMsgAllowance{} }
func (m *AllowedMsgAllowance) String() string { return proto.CompactTextString(m) }
func (*AllowedMsgAllowance) ProtoMessage()    {}
func (*AllowedMsgAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{2}
}
func (m *AllowedMsgAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedMsgAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedMsgAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedMsgAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedMsgAllowance.Merge(m, src)
}
func (m *AllowedMsgAllowance) XXX_Size() int {
	return m.Size()
}
func (m *AllowedMsgAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedMsgAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedMsgAllowance proto.InternalMessageInfo

// Grant is stored in the KVStore to record a grant with full context
type Grant struct {
	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// allowance can be any of basic, periodic, allowed fee allowance.
	Allowance *types1.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{3}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *Grant) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *Grant) GetAllowance() *types1.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}

func init() {
	proto.RegisterType((*BasicAllowance)(nil), "cosmos.feegrant.v1beta1.BasicAllowance")
	proto.RegisterType((*PeriodicAllowance)(nil), "cosmos.feegrant.v1beta1.PeriodicAllowance")
	proto.RegisterType((*AllowedMsgAllowance)(nil), "cosmos.feegrant.v1beta1.AllowedMsgAllowance")
	proto.RegisterType((*Grant)(nil), "cosmos.feegrant.v1beta1.Grant")
}

func init() {
	proto.RegisterFile("cosmos/feegrant/v1beta1/feegrant.proto", fileDescriptor_7279582900c30aea)
}

var fileDescriptor_7279582900c30aea = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x9b, 0xa6, 0x90, 0x0b, 0x94, 0xd6, 0x04, 0xe1, 0x64, 0x70, 0xa2, 0x0e, 0x34, 0x0c,
	0xb1, 0x69, 0xd8, 0xca, 0x42, 0x1c, 0xa0, 0x42, 0xa2, 0x12, 0x72, 0x99, 0x58, 0xa2, 0xb3, 0xfd,
	0x6a, 0x4e, 0xc4, 0x3e, 0xcb, 0x77, 0x81, 0xe6, 0x1b, 0x30, 0x76, 0x64, 0x42, 0xcc, 0xcc, 0x15,
	0x9f, 0xa1, 0x62, 0xaa, 0x60, 0x61, 0xa2, 0x28, 0xf9, 0x22, 0xc8, 0x77, 0xe7, 0x24, 0x24, 0xfc,
	0x13, 0xea, 0x14, 0xdf, 0xbb, 0xf7, 0xfb, 0xf7, 0xde, 0x29, 0xe8, 0x96, 0x4f, 0x59, 0x44, 0x99,
	0x7d, 0x08, 0x10, 0xa6, 0x38, 0xe6, 0xf6, 0xab, 0x1d, 0x0f, 0x38, 0xde, 0x99, 0x16, 0xac, 0x24,
	0xa5, 0x9c, 0xea, 0x37, 0x65, 0x9f, 0x35, 0x2d, 0xab, 0xbe, 0x7a, 0x35, 0xa4, 0x21, 0x15, 0x3d,
	0x76, 0xf6, 0x25, 0xdb, 0xeb, 0xb5, 0x90, 0xd2, 0x70, 0x00, 0xb6, 0x38, 0x79, 0xc3, 0x43, 0x1b,
	0xc7, 0xa3, 0xfc, 0x4a, 0x32, 0xf5, 0x25, 0x46, 0xd1, 0xca, 0x2b, 0x53, 0x99, 0xf1, 0x30, 0x83,
	0xa9, 0x11, 0x9f, 0x92, 0x58, 0xdd, 0x37, 0x16, 0x59, 0x39, 0x89, 0x80, 0x71, 0x1c, 0x25, 0x39,
	0xc1, 0x62, 0x43, 0x30, 0x4c, 0x31, 0x27, 0x54, 0x11, 0x6c, 0x7d, 0xd1, 0xd0, 0xba, 0x83, 0x19,
	0xf1, 0xbb, 0x83, 0x01, 0x7d, 0x8d, 0x63, 0x1f, 0xf4, 0x01, 0xaa, 0xb0, 0x04, 0xe2, 0xa0, 0x3f,
	0x20, 0x11, 0xe1, 0x86, 0xd6, 0x2c, 0xb6, 0x2a, 0x9d, 0x9a, 0xa5, 0x7c, 0x65, 0x4e, 0xf2, 0xa8,
	0x56, 0x8f, 0x92, 0xd8, 0xb9, 0x73, 0xfa, 0xad, 0x51, 0xf8, 0x70, 0xde, 0x68, 0x85, 0x84, 0xbf,
	0x18, 0x7a, 0x96, 0x4f, 0x23, 0x15, 0x42, 0xfd, 0xb4, 0x59, 0xf0, 0xd2, 0xe6, 0xa3, 0x04, 0x98,
	0x00, 0x30, 0x17, 0x09, 0xfe, 0x27, 0x19, 0xbd, 0x7e, 0x1f, 0x21, 0x38, 0x4a, 0x88, 0x34, 0x65,
	0xac, 0x34, 0xb5, 0x56, 0xa5, 0x53, 0xb7, 0xa4, 0x6b, 0x2b, 0x77, 0x6d, 0x3d, 0xcb, 0x63, 0x39,
	0xab, 0xc7, 0xe7, 0x0d, 0xcd, 0x9d, 0xc3, 0xec, 0x6e, 0x7e, 0x3a, 0x69, 0x5f, 0x7d, 0x04, 0x30,
	0x4d, 0xf0, 0x78, 0x6b, 0x52, 0x44, 0x9b, 0x4f, 0x21, 0x25, 0x34, 0x98, 0x0f, 0xd6, 0x43, 0x25,
	0x2f, 0x8b, 0x6a, 0x68, 0x42, 0x65, 0xdb, 0xfa, 0xcd, 0x06, 0xad, 0x9f, 0x07, 0xe2, 0xac, 0x66,
	0x01, 0x5d, 0x89, 0xd5, 0xef, 0xa1, 0xb5, 0x44, 0x30, 0x2b, 0xaf, 0xb5, 0x25, 0xaf, 0x0f, 0xd4,
	0x84, 0x9d, 0xcb, 0x19, 0xee, 0x6d, 0x66, 0x57, 0x41, 0xf4, 0x11, 0xd2, 0xe5, 0x57, 0x7f, 0x7e,
	0xc2, 0xc5, 0x8b, 0x9f, 0xf0, 0x86, 0x94, 0x39, 0x98, 0xcd, 0x79, 0x88, 0x54, 0xad, 0xef, 0xe3,
	0x58, 0xca, 0x1b, 0xab, 0x17, 0x2f, 0xbc, 0x2e, 0x45, 0x7a, 0x38, 0x16, 0xda, 0xfa, 0x1e, 0xba,
	0xa2, 0x64, 0x53, 0x60, 0xc0, 0x8d, 0xd2, 0x5f, 0x17, 0x2c, 0xa6, 0x26, 0x96, 0x5c, 0x91, 0x48,
	0x37, 0x03, 0xfe, 0x6a, 0xcb, 0xef, 0x34, 0x74, 0x5d, 0x1c, 0x21, 0xd8, 0x67, 0xe1, 0x6c, 0xcf,
	0x0f, 0x51, 0x19, 0xe7, 0x07, 0xb5, 0xeb, 0xea, 0x92, 0x60, 0x37, 0x1e, 0x39, 0xcb, 0x9c, 0xee,
	0x0c, 0xa9, 0xdf, 0x46, 0x1b, 0x58, 0xb2, 0xf7, 0x23, 0x60, 0x0c, 0x87, 0xc0, 0x8c, 0x95, 0x66,
	0xb1, 0x55, 0x76, 0xaf, 0xa9, 0xfa, 0xbe, 0x2a, 0xef, 0xde, 0x78, 0xf3, 0xbe, 0x51, 0x58, 0x36,
	0xf8, 0x51, 0x43, 0xa5, 0xbd, 0xec, 0x65, 0xe9, 0x1d, 0x74, 0x49, 0x3c, 0x31, 0x48, 0x85, 0xa1,
	0xb2, 0x63, 0x7c, 0x3e, 0x69, 0x57, 0xd5, 0xdc, 0xbb, 0x41, 0x90, 0x02, 0x63, 0x07, 0x3c, 0x25,
	0x71, 0xe8, 0xe6, 0x8d, 0x33, 0x0c, 0x88, 0xa7, 0xf6, 0x0f, 0x98, 0x85, 0xe8, 0xc5, 0xff, 0x8d,
	0xee, 0x74, 0x4f, 0xc7, 0xa6, 0x76, 0x36, 0x36, 0xb5, 0xef, 0x63, 0x53, 0x3b, 0x9e, 0x98, 0x85,
	0xb3, 0x89, 0x59, 0xf8, 0x3a, 0x31, 0x0b, 0xcf, 0xb7, 0xff, 0xf8, 0x12, 0x8e, 0xa6, 0x7f, 0x92,
	0xde, 0x9a, 0x90, 0xbb, 0xfb, 0x23, 0x00, 0x00, 0xff, 0xff, 0x41, 0xd8, 0x15, 0xc6, 0x4f, 0x05,
	0x00, 0x00,
}

func (m *BasicAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Expiration):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintFeegrant(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PeriodReset, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PeriodReset):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintFeegrant(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.PeriodCanSpend) > 0 {
		for iNdEx := len(m.PeriodCanSpend) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodCanSpend[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PeriodSpendLimit) > 0 {
		for iNdEx := len(m.PeriodSpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodSpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Period, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Period):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintFeegrant(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Basic.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllowedMsgAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedMsgAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedMsgAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedMessages) > 0 {
		for iNdEx := len(m.AllowedMessages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedMessages[iNdEx])
			copy(dAtA[i:], m.AllowedMessages[iNdEx])
			i = encodeVarintFeegrant(dAtA, i, uint64(len(m.AllowedMessages[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeegrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeegrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasicAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	if m.Expiration != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovFeegrant(uint64(l))
	}
	return n
}

func (m *PeriodicAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Basic.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Period)
	n += 1 + l + sovFeegrant(uint64(l))
	if len(m.PeriodSpendLimit) > 0 {
		for _, e := range m.PeriodSpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	if len(m.PeriodCanSpend) > 0 {
		for _, e := range m.PeriodCanSpend {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PeriodReset)
	n += 1 + l + sovFeegrant(uint64(l))
	return n
}

func (m *AllowedMsgAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovFeegrant(uint64(l))
	}
	if len(m.AllowedMessages) > 0 {
		for _, s := range m.AllowedMessages {
			l = len(s)
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	return n
}

func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovFeegrant(uint64(l))
	}
	return n
}

func sovFeegrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeegrant(x uint64) (n int) {
	return sovFeegrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasicAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
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
			return fmt.Errorf("proto: BasicAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
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
func (m *PeriodicAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
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
			return fmt.Errorf("proto: PeriodicAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Basic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Basic.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Period, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodSpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodSpendLimit = append(m.PeriodSpendLimit, types.Coin{})
			if err := m.PeriodSpendLimit[len(m.PeriodSpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodCanSpend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodCanSpend = append(m.PeriodCanSpend, types.Coin{})
			if err := m.PeriodCanSpend[len(m.PeriodCanSpend)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodReset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PeriodReset, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
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
func (m *AllowedMsgAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
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
			return fmt.Errorf("proto: AllowedMsgAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedMsgAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types1.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedMessages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedMessages = append(m.AllowedMessages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
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
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
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
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types1.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
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
func skipFeegrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeegrant
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
					return 0, ErrIntOverflowFeegrant
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
					return 0, ErrIntOverflowFeegrant
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
				return 0, ErrInvalidLengthFeegrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeegrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeegrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeegrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeegrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeegrant = fmt.Errorf("proto: unexpected end of group")
)
