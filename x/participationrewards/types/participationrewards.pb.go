// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/participationrewards.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ProtocolDataType int32

const (
	// Undefined action (per protobuf spec)
	ProtocolDataTypeUndefined     ProtocolDataType = 0
	ProtocolDataTypeConnection    ProtocolDataType = 1
	ProtocolDataTypeOsmosisParams ProtocolDataType = 2
	ProtocolDataTypeLiquidToken   ProtocolDataType = 3
	ProtocolDataTypeOsmosisPool   ProtocolDataType = 4
	ProtocolDataTypeCrescentPool  ProtocolDataType = 5
	ProtocolDataTypeSifchainPool  ProtocolDataType = 6
)

var ProtocolDataType_name = map[int32]string{
	0: "ProtocolDataTypeUndefined",
	1: "ProtocolDataTypeConnection",
	2: "ProtocolDataTypeOsmosisParams",
	3: "ProtocolDataTypeLiquidToken",
	4: "ProtocolDataTypeOsmosisPool",
	5: "ProtocolDataTypeCrescentPool",
	6: "ProtocolDataTypeSifchainPool",
}

var ProtocolDataType_value = map[string]int32{
	"ProtocolDataTypeUndefined":     0,
	"ProtocolDataTypeConnection":    1,
	"ProtocolDataTypeOsmosisParams": 2,
	"ProtocolDataTypeLiquidToken":   3,
	"ProtocolDataTypeOsmosisPool":   4,
	"ProtocolDataTypeCrescentPool":  5,
	"ProtocolDataTypeSifchainPool":  6,
}

func (x ProtocolDataType) String() string {
	return proto.EnumName(ProtocolDataType_name, int32(x))
}

func (ProtocolDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4fb4e5bb851c124, []int{0}
}

// DistributionProportions defines the proportions of minted QCK that is to be
// allocated as participation rewards.
type DistributionProportions struct {
	ValidatorSelectionAllocation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=validator_selection_allocation,json=validatorSelectionAllocation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"validator_selection_allocation"`
	HoldingsAllocation           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=holdings_allocation,json=holdingsAllocation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"holdings_allocation"`
	LockupAllocation             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=lockup_allocation,json=lockupAllocation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"lockup_allocation"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fb4e5bb851c124, []int{0}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

// Params holds parameters for the participationrewards module.
type Params struct {
	// distribution_proportions defines the proportions of the minted
	// participation rewards;
	DistributionProportions DistributionProportions `protobuf:"bytes,1,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	ClaimsEnabled           bool                    `protobuf:"varint,2,opt,name=claims_enabled,json=claimsEnabled,proto3" json:"claims_enabled,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fb4e5bb851c124, []int{1}
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

type KeyedProtocolData struct {
	Key          string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ProtocolData *ProtocolData `protobuf:"bytes,2,opt,name=protocol_data,json=protocolData,proto3" json:"protocol_data,omitempty"`
}

func (m *KeyedProtocolData) Reset()         { *m = KeyedProtocolData{} }
func (m *KeyedProtocolData) String() string { return proto.CompactTextString(m) }
func (*KeyedProtocolData) ProtoMessage()    {}
func (*KeyedProtocolData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fb4e5bb851c124, []int{2}
}
func (m *KeyedProtocolData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyedProtocolData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyedProtocolData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyedProtocolData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyedProtocolData.Merge(m, src)
}
func (m *KeyedProtocolData) XXX_Size() int {
	return m.Size()
}
func (m *KeyedProtocolData) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyedProtocolData.DiscardUnknown(m)
}

var xxx_messageInfo_KeyedProtocolData proto.InternalMessageInfo

func (m *KeyedProtocolData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyedProtocolData) GetProtocolData() *ProtocolData {
	if m != nil {
		return m.ProtocolData
	}
	return nil
}

// Protocol Data is an arbitrary data type held against a given zone for the
// determination of rewards.
type ProtocolData struct {
	Type string                   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data encoding_json.RawMessage `protobuf:"bytes,2,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty"`
}

func (m *ProtocolData) Reset()         { *m = ProtocolData{} }
func (m *ProtocolData) String() string { return proto.CompactTextString(m) }
func (*ProtocolData) ProtoMessage()    {}
func (*ProtocolData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fb4e5bb851c124, []int{3}
}
func (m *ProtocolData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolData.Merge(m, src)
}
func (m *ProtocolData) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolData.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolData proto.InternalMessageInfo

func (m *ProtocolData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ProtocolData) GetData() encoding_json.RawMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("quicksilver.participationrewards.v1.ProtocolDataType", ProtocolDataType_name, ProtocolDataType_value)
	proto.RegisterType((*DistributionProportions)(nil), "quicksilver.participationrewards.v1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "quicksilver.participationrewards.v1.Params")
	proto.RegisterType((*KeyedProtocolData)(nil), "quicksilver.participationrewards.v1.KeyedProtocolData")
	proto.RegisterType((*ProtocolData)(nil), "quicksilver.participationrewards.v1.ProtocolData")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/participationrewards.proto", fileDescriptor_d4fb4e5bb851c124)
}

var fileDescriptor_d4fb4e5bb851c124 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0x4e, 0xba, 0xfc, 0xa6, 0xfd, 0xbc, 0x0d, 0x65, 0x06, 0x69, 0x5b, 0xd9, 0xd2, 0x31, 0x04,
	0x42, 0x48, 0x4b, 0xd8, 0xb8, 0x4d, 0x13, 0x12, 0x5d, 0x39, 0x01, 0x62, 0xca, 0x0a, 0x07, 0x24,
	0x54, 0xb9, 0xb6, 0x97, 0x9a, 0xba, 0x76, 0x16, 0x27, 0x1d, 0x3d, 0x4c, 0x42, 0x9c, 0x76, 0xe4,
	0xc8, 0x11, 0x89, 0x7f, 0x81, 0x03, 0x7f, 0xc2, 0x8e, 0x13, 0x27, 0xc4, 0xa1, 0x42, 0xed, 0x7f,
	0xc1, 0x01, 0x21, 0x27, 0x6d, 0x09, 0xa5, 0x93, 0x76, 0xd8, 0x29, 0xcf, 0xcf, 0x9f, 0xbf, 0xef,
	0xf9, 0x7d, 0x2f, 0x06, 0x0f, 0x0e, 0x13, 0x86, 0x9b, 0x8a, 0xf1, 0x36, 0x8d, 0xbc, 0x10, 0x45,
	0x31, 0xc3, 0x2c, 0x44, 0x31, 0x93, 0x22, 0xa2, 0x47, 0x28, 0x22, 0xca, 0x6b, 0x6f, 0x4e, 0xcc,
	0xbb, 0x61, 0x24, 0x63, 0x09, 0x6f, 0xe6, 0xce, 0xbb, 0x13, 0x71, 0xed, 0xcd, 0xe2, 0x32, 0x96,
	0xaa, 0x25, 0x55, 0x2d, 0x3d, 0xe2, 0x65, 0x8b, 0xec, 0x7c, 0xf1, 0x5a, 0x20, 0x03, 0x99, 0xe5,
	0x75, 0x94, 0x65, 0xd7, 0x7f, 0x15, 0xc0, 0x62, 0x85, 0xa9, 0x38, 0x62, 0xf5, 0x44, 0x73, 0xed,
	0x45, 0x32, 0x94, 0x91, 0x8e, 0x14, 0x7c, 0x67, 0x02, 0xa7, 0x8d, 0x38, 0x23, 0x28, 0x96, 0x51,
	0x4d, 0x51, 0x4e, 0xb1, 0xde, 0xa8, 0x21, 0xce, 0x25, 0x4e, 0x95, 0x97, 0xcc, 0x35, 0xf3, 0xce,
	0xff, 0xe5, 0x9d, 0xd3, 0x6e, 0xc9, 0xf8, 0xde, 0x2d, 0xdd, 0x0e, 0x58, 0xdc, 0x48, 0xea, 0x2e,
	0x96, 0xad, 0x81, 0xf6, 0xe0, 0xb3, 0xa1, 0x48, 0xd3, 0x8b, 0x3b, 0x21, 0x55, 0x6e, 0x85, 0xe2,
	0xaf, 0x9f, 0x37, 0xc0, 0xa0, 0xb4, 0x0a, 0xc5, 0xfe, 0xca, 0x48, 0x63, 0x7f, 0x28, 0xf1, 0x70,
	0xa4, 0x00, 0x5b, 0xe0, 0x6a, 0x43, 0x72, 0xc2, 0x44, 0xa0, 0xf2, 0xc2, 0x85, 0x4b, 0x10, 0x86,
	0x43, 0xe2, 0x9c, 0x1c, 0x03, 0x0b, 0x5c, 0xe2, 0x66, 0x12, 0xe6, 0xc5, 0xa6, 0x2e, 0x41, 0xcc,
	0xce, 0x68, 0xff, 0x48, 0x6d, 0x5b, 0x27, 0x1f, 0x4b, 0xc6, 0xfa, 0x17, 0x13, 0x4c, 0xef, 0xa1,
	0x08, 0xb5, 0x14, 0x3c, 0x06, 0x4b, 0x24, 0x67, 0x85, 0x36, 0x71, 0xe8, 0x45, 0xda, 0xe8, 0xd9,
	0xad, 0x1d, 0xf7, 0x02, 0x43, 0xe0, 0x9e, 0xe3, 0x67, 0xd9, 0xd2, 0x17, 0xf0, 0x17, 0xc9, 0x39,
	0x76, 0xdf, 0x02, 0x57, 0x30, 0x47, 0xac, 0xa5, 0x6a, 0x54, 0xa0, 0x3a, 0xa7, 0x24, 0x6d, 0xf2,
	0x8c, 0x3f, 0x9f, 0x65, 0x1f, 0x65, 0xc9, 0xed, 0x19, 0x5d, 0xf6, 0x07, 0x5d, 0xfa, 0x31, 0x58,
	0x78, 0x4c, 0x3b, 0x94, 0xec, 0xe9, 0x49, 0xc2, 0x92, 0x57, 0x50, 0x8c, 0xa0, 0x0d, 0xa6, 0x9a,
	0xb4, 0x93, 0x0d, 0x86, 0xaf, 0x43, 0xf8, 0x02, 0xcc, 0x87, 0x03, 0x44, 0x8d, 0xa0, 0x18, 0xa5,
	0xb4, 0xb3, 0x5b, 0x9b, 0x17, 0xba, 0x4b, 0x9e, 0xdb, 0x9f, 0x0b, 0x73, 0xab, 0xf5, 0x2a, 0x98,
	0xfb, 0x4b, 0x19, 0x02, 0x4b, 0x37, 0x7f, 0x20, 0x9d, 0xc6, 0xf0, 0x1e, 0xb0, 0x46, 0x92, 0x73,
	0xe5, 0x95, 0x9f, 0xdd, 0xd2, 0x12, 0x15, 0x58, 0x6a, 0xd7, 0xbd, 0xd7, 0x4a, 0x0a, 0xd7, 0x47,
	0x47, 0x4f, 0xa9, 0x52, 0x28, 0xa0, 0x7e, 0x8a, 0xbc, 0xfb, 0xb6, 0x00, 0xec, 0x3c, 0x6d, 0x55,
	0xd3, 0xac, 0x82, 0xe5, 0xf1, 0xdc, 0x73, 0x41, 0xe8, 0x01, 0x13, 0x94, 0xd8, 0x06, 0x74, 0x40,
	0x71, 0x7c, 0x7b, 0x57, 0x0a, 0x91, 0xcd, 0xb2, 0x6d, 0xc2, 0x1b, 0x60, 0x75, 0x7c, 0xff, 0x99,
	0x1e, 0x0c, 0xa6, 0x32, 0xe7, 0xed, 0x02, 0x2c, 0x81, 0xeb, 0xe3, 0x90, 0x27, 0xec, 0x30, 0x61,
	0xa4, 0x2a, 0x9b, 0x54, 0xd8, 0x53, 0x93, 0x00, 0x43, 0x0e, 0x29, 0xb9, 0x6d, 0xc1, 0x35, 0xb0,
	0xf2, 0x4f, 0x11, 0x11, 0x55, 0x98, 0x8a, 0x38, 0x45, 0xfc, 0x37, 0x09, 0xb1, 0xcf, 0x0e, 0x70,
	0x03, 0x31, 0x91, 0x22, 0xa6, 0x8b, 0xd6, 0xc9, 0x27, 0xc7, 0x28, 0xbf, 0x3a, 0xed, 0x39, 0xe6,
	0x59, 0xcf, 0x31, 0x7f, 0xf4, 0x1c, 0xf3, 0x7d, 0xdf, 0x31, 0xce, 0xfa, 0x8e, 0xf1, 0xad, 0xef,
	0x18, 0x2f, 0x77, 0x73, 0xa3, 0xcf, 0x44, 0x40, 0x45, 0xc2, 0xe2, 0xce, 0x46, 0x3d, 0x61, 0x9c,
	0x78, 0xf9, 0xe7, 0xed, 0xcd, 0xe4, 0x07, 0x2e, 0xfd, 0x37, 0xea, 0xd3, 0xa9, 0x8b, 0xf7, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x87, 0x46, 0x1b, 0x11, 0x05, 0x00, 0x00,
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LockupAllocation.Size()
		i -= size
		if _, err := m.LockupAllocation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipationrewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.HoldingsAllocation.Size()
		i -= size
		if _, err := m.HoldingsAllocation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipationrewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ValidatorSelectionAllocation.Size()
		i -= size
		if _, err := m.ValidatorSelectionAllocation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipationrewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if m.ClaimsEnabled {
		i--
		if m.ClaimsEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParticipationrewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *KeyedProtocolData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyedProtocolData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyedProtocolData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProtocolData != nil {
		{
			size, err := m.ProtocolData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParticipationrewards(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintParticipationrewards(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintParticipationrewards(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintParticipationrewards(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParticipationrewards(dAtA []byte, offset int, v uint64) int {
	offset -= sovParticipationrewards(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ValidatorSelectionAllocation.Size()
	n += 1 + l + sovParticipationrewards(uint64(l))
	l = m.HoldingsAllocation.Size()
	n += 1 + l + sovParticipationrewards(uint64(l))
	l = m.LockupAllocation.Size()
	n += 1 + l + sovParticipationrewards(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParticipationrewards(uint64(l))
	if m.ClaimsEnabled {
		n += 2
	}
	return n
}

func (m *KeyedProtocolData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovParticipationrewards(uint64(l))
	}
	if m.ProtocolData != nil {
		l = m.ProtocolData.Size()
		n += 1 + l + sovParticipationrewards(uint64(l))
	}
	return n
}

func (m *ProtocolData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovParticipationrewards(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovParticipationrewards(uint64(l))
	}
	return n
}

func sovParticipationrewards(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParticipationrewards(x uint64) (n int) {
	return sovParticipationrewards(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipationrewards
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSelectionAllocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorSelectionAllocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HoldingsAllocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HoldingsAllocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupAllocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockupAllocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParticipationrewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipationrewards
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipationrewards
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
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimsEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
			m.ClaimsEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParticipationrewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipationrewards
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
func (m *KeyedProtocolData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipationrewards
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
			return fmt.Errorf("proto: KeyedProtocolData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyedProtocolData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProtocolData == nil {
				m.ProtocolData = &ProtocolData{}
			}
			if err := m.ProtocolData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParticipationrewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipationrewards
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
func (m *ProtocolData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipationrewards
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
			return fmt.Errorf("proto: ProtocolData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
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
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipationrewards
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipationrewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParticipationrewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipationrewards
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
func skipParticipationrewards(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParticipationrewards
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
					return 0, ErrIntOverflowParticipationrewards
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
					return 0, ErrIntOverflowParticipationrewards
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
				return 0, ErrInvalidLengthParticipationrewards
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParticipationrewards
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParticipationrewards
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParticipationrewards        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParticipationrewards          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParticipationrewards = fmt.Errorf("proto: unexpected end of group")
)
