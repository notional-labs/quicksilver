// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/participationrewards.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type ProtocolDataType int32

const (
	// Undefined action (per protobuf spec)
	ProtocolDataTypeUndefined                     ProtocolDataType = 0
	ProtocolDataTypeConnection                    ProtocolDataType = 1
	ProtocolDataTypeOsmosisParams                 ProtocolDataType = 2
	ProtocolDataTypeLiquidToken                   ProtocolDataType = 3
	ProtocolDataTypeOsmosisPool                   ProtocolDataType = 4
	ProtocolDataTypeCrescentPool                  ProtocolDataType = 5
	ProtocolDataTypeSifchainPool                  ProtocolDataType = 6
	ProtocolDataTypeUmeeParams                    ProtocolDataType = 7
	ProtocolDataTypeUmeeReserves                  ProtocolDataType = 8
	ProtocolDataTypeUmeeInterestScalar            ProtocolDataType = 9
	ProtocolDataTypeUmeeTotalBorrows              ProtocolDataType = 10
	ProtocolDataTypeUmeeUTokenSupply              ProtocolDataType = 11
	ProtocolDataTypeUmeeLeverageModuleBalance     ProtocolDataType = 12
	ProtocolDataTypeCrescentParams                ProtocolDataType = 13
	ProtocolDataTypeCrescentReserveAddressBalance ProtocolDataType = 14
	ProtocolDataTypeCrescentPoolCoinSupply        ProtocolDataType = 15
)

var ProtocolDataType_name = map[int32]string{
	0:  "ProtocolDataTypeUndefined",
	1:  "ProtocolDataTypeConnection",
	2:  "ProtocolDataTypeOsmosisParams",
	3:  "ProtocolDataTypeLiquidToken",
	4:  "ProtocolDataTypeOsmosisPool",
	5:  "ProtocolDataTypeCrescentPool",
	6:  "ProtocolDataTypeSifchainPool",
	7:  "ProtocolDataTypeUmeeParams",
	8:  "ProtocolDataTypeUmeeReserves",
	9:  "ProtocolDataTypeUmeeInterestScalar",
	10: "ProtocolDataTypeUmeeTotalBorrows",
	11: "ProtocolDataTypeUmeeUTokenSupply",
	12: "ProtocolDataTypeUmeeLeverageModuleBalance",
	13: "ProtocolDataTypeCrescentParams",
	14: "ProtocolDataTypeCrescentReserveAddressBalance",
	15: "ProtocolDataTypeCrescentPoolCoinSupply",
}

var ProtocolDataType_value = map[string]int32{
	"ProtocolDataTypeUndefined":                     0,
	"ProtocolDataTypeConnection":                    1,
	"ProtocolDataTypeOsmosisParams":                 2,
	"ProtocolDataTypeLiquidToken":                   3,
	"ProtocolDataTypeOsmosisPool":                   4,
	"ProtocolDataTypeCrescentPool":                  5,
	"ProtocolDataTypeSifchainPool":                  6,
	"ProtocolDataTypeUmeeParams":                    7,
	"ProtocolDataTypeUmeeReserves":                  8,
	"ProtocolDataTypeUmeeInterestScalar":            9,
	"ProtocolDataTypeUmeeTotalBorrows":              10,
	"ProtocolDataTypeUmeeUTokenSupply":              11,
	"ProtocolDataTypeUmeeLeverageModuleBalance":     12,
	"ProtocolDataTypeCrescentParams":                13,
	"ProtocolDataTypeCrescentReserveAddressBalance": 14,
	"ProtocolDataTypeCrescentPoolCoinSupply":        15,
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
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x3f, 0x4f, 0x1b, 0x49,
	0x14, 0xf7, 0x82, 0x8f, 0x83, 0xc1, 0x70, 0xcb, 0xdc, 0x49, 0x80, 0x0f, 0xd6, 0x9c, 0x2f, 0x41,
	0x09, 0x92, 0xed, 0x98, 0x74, 0x08, 0x45, 0xc2, 0x38, 0x45, 0x14, 0x50, 0xd0, 0xda, 0xa4, 0x48,
	0x11, 0x6b, 0x3c, 0xfb, 0x30, 0x13, 0x8f, 0x77, 0x96, 0x99, 0xb5, 0x89, 0x23, 0xd1, 0xa4, 0xa2,
	0x4c, 0x99, 0x32, 0x52, 0xbe, 0x42, 0x8a, 0x7c, 0x04, 0x4a, 0x94, 0x2a, 0x4a, 0x81, 0x22, 0x68,
	0xf2, 0x19, 0x52, 0x44, 0xd1, 0xec, 0xda, 0x64, 0xe3, 0xac, 0x23, 0x0a, 0x2a, 0xbf, 0x7d, 0xef,
	0x37, 0xbf, 0xdf, 0xfb, 0x37, 0x63, 0x74, 0xef, 0xa0, 0xcd, 0x68, 0x53, 0x31, 0xde, 0x01, 0x59,
	0xf0, 0x88, 0xf4, 0x19, 0x65, 0x1e, 0xf1, 0x99, 0x70, 0x25, 0x1c, 0x12, 0xe9, 0xa8, 0x42, 0xa7,
	0x18, 0xeb, 0xcf, 0x7b, 0x52, 0xf8, 0x02, 0xff, 0x1f, 0x39, 0x9f, 0x8f, 0xc5, 0x75, 0x8a, 0xe9,
	0x79, 0x2a, 0x54, 0x4b, 0xa8, 0x5a, 0x70, 0xa4, 0x10, 0x7e, 0x84, 0xe7, 0xd3, 0xff, 0x34, 0x44,
	0x43, 0x84, 0x7e, 0x6d, 0x85, 0xde, 0xec, 0xb7, 0x11, 0x34, 0x5b, 0x66, 0xca, 0x97, 0xac, 0xde,
	0xd6, 0x5c, 0x3b, 0x52, 0x78, 0x42, 0x6a, 0x4b, 0xe1, 0x97, 0x06, 0xb2, 0x3a, 0x84, 0x33, 0x87,
	0xf8, 0x42, 0xd6, 0x14, 0x70, 0xa0, 0x3a, 0x50, 0x23, 0x9c, 0x0b, 0x1a, 0x28, 0xcf, 0x19, 0x4b,
	0xc6, 0xad, 0x89, 0xd2, 0xfa, 0xc9, 0x59, 0x26, 0xf1, 0xe9, 0x2c, 0xb3, 0xdc, 0x60, 0xfe, 0x7e,
	0xbb, 0x9e, 0xa7, 0xa2, 0xd5, 0xd3, 0xee, 0xfd, 0xe4, 0x94, 0xd3, 0x2c, 0xf8, 0x5d, 0x0f, 0x54,
	0xbe, 0x0c, 0xf4, 0xc3, 0xbb, 0x1c, 0xea, 0xa5, 0x56, 0x06, 0x6a, 0x2f, 0x5c, 0x6a, 0x54, 0xfa,
	0x12, 0x1b, 0x97, 0x0a, 0xb8, 0x85, 0xfe, 0xde, 0x17, 0xdc, 0x61, 0x6e, 0x43, 0x45, 0x85, 0x47,
	0xae, 0x41, 0x18, 0xf7, 0x89, 0x23, 0x72, 0x0c, 0xcd, 0x70, 0x41, 0x9b, 0x6d, 0x2f, 0x2a, 0x36,
	0x7a, 0x0d, 0x62, 0x66, 0x48, 0xfb, 0x43, 0x6a, 0x2d, 0x79, 0xfc, 0x26, 0x93, 0xc8, 0xbe, 0x37,
	0xd0, 0xd8, 0x0e, 0x91, 0xa4, 0xa5, 0xf0, 0x11, 0x9a, 0x73, 0x22, 0xa3, 0xd0, 0x43, 0xec, 0xcf,
	0x22, 0x68, 0xf4, 0xe4, 0xea, 0x7a, 0xfe, 0x0a, 0x4b, 0x90, 0x1f, 0x32, 0xcf, 0x52, 0x52, 0x17,
	0x60, 0xcf, 0x3a, 0x43, 0xc6, 0x7d, 0x13, 0x4d, 0x53, 0x4e, 0x58, 0x4b, 0xd5, 0xc0, 0x25, 0x75,
	0x0e, 0x4e, 0xd0, 0xe4, 0x71, 0x7b, 0x2a, 0xf4, 0xde, 0x0f, 0x9d, 0x6b, 0xe3, 0x3a, 0xed, 0xd7,
	0x3a, 0xf5, 0x23, 0x34, 0xf3, 0x10, 0xba, 0xe0, 0xec, 0xe8, 0x4d, 0xa2, 0x82, 0x97, 0x89, 0x4f,
	0xb0, 0x89, 0x46, 0x9b, 0xd0, 0x0d, 0x17, 0xc3, 0xd6, 0x26, 0x7e, 0x8c, 0xa6, 0xbc, 0x1e, 0xa2,
	0xe6, 0x10, 0x9f, 0x04, 0xb4, 0x93, 0xab, 0xc5, 0x2b, 0xd5, 0x12, 0xe5, 0xb6, 0x53, 0x5e, 0xe4,
	0x2b, 0x5b, 0x45, 0xa9, 0x9f, 0x94, 0x31, 0x4a, 0xea, 0xe6, 0xf7, 0xa4, 0x03, 0x1b, 0xdf, 0x41,
	0xc9, 0x4b, 0xc9, 0x54, 0x69, 0xe1, 0xeb, 0x59, 0x66, 0x0e, 0x5c, 0x2a, 0xf4, 0xd4, 0x0b, 0xcf,
	0x94, 0x70, 0xf3, 0x36, 0x39, 0xdc, 0x06, 0xa5, 0x48, 0x03, 0xec, 0x00, 0xb9, 0xf2, 0x25, 0x89,
	0xcc, 0x28, 0x6d, 0x55, 0xd3, 0x2c, 0xa2, 0xf9, 0x41, 0xdf, 0xae, 0xeb, 0xc0, 0x1e, 0x73, 0xc1,
	0x31, 0x13, 0xd8, 0x42, 0xe9, 0xc1, 0xf0, 0xa6, 0x70, 0xdd, 0x70, 0x97, 0x4d, 0x03, 0xff, 0x87,
	0x16, 0x07, 0xe3, 0x8f, 0xf4, 0x62, 0x30, 0x15, 0x4e, 0xde, 0x1c, 0xc1, 0x19, 0xf4, 0xef, 0x20,
	0x64, 0x8b, 0x1d, 0xb4, 0x99, 0x53, 0x15, 0x4d, 0x70, 0xcd, 0xd1, 0x38, 0x40, 0x9f, 0x43, 0x08,
	0x6e, 0x26, 0xf1, 0x12, 0x5a, 0xf8, 0x25, 0x09, 0x09, 0x8a, 0x82, 0xeb, 0x07, 0x88, 0x3f, 0xe2,
	0x10, 0x15, 0xb6, 0x47, 0xf7, 0x09, 0x73, 0x03, 0xc4, 0x58, 0x5c, 0x21, 0xbb, 0x2d, 0x80, 0x5e,
	0x96, 0x7f, 0xc6, 0x31, 0xe8, 0xb8, 0x0d, 0x0a, 0x64, 0x07, 0x94, 0x39, 0x8e, 0x97, 0x51, 0x36,
	0x0e, 0xf1, 0xc0, 0xf5, 0x41, 0x82, 0xf2, 0x2b, 0x94, 0x70, 0x22, 0xcd, 0x09, 0x7c, 0x03, 0x2d,
	0xc5, 0xe1, 0xaa, 0xc2, 0x27, 0xbc, 0x24, 0xa4, 0x14, 0x87, 0xca, 0x44, 0xc3, 0x50, 0xbb, 0x41,
	0x53, 0x2a, 0x6d, 0xcf, 0xe3, 0x5d, 0x73, 0x12, 0xe7, 0xd0, 0xed, 0x38, 0xd4, 0x16, 0x74, 0x40,
	0x92, 0x06, 0x6c, 0x0b, 0xa7, 0xcd, 0xa1, 0x44, 0x38, 0x71, 0x29, 0x98, 0x29, 0x9c, 0x45, 0xd6,
	0xd0, 0x46, 0x85, 0x85, 0x4e, 0xe1, 0x22, 0xca, 0x0d, 0xc3, 0xf4, 0x8a, 0xdd, 0x70, 0x1c, 0x09,
	0x4a, 0xf5, 0x69, 0xa7, 0xf1, 0x0a, 0x5a, 0xfe, 0x5d, 0xff, 0x37, 0x05, 0xeb, 0x67, 0xfc, 0x57,
	0x3a, 0x79, 0xfc, 0xd6, 0x4a, 0x94, 0x9e, 0x9e, 0x9c, 0x5b, 0xc6, 0xe9, 0xb9, 0x65, 0x7c, 0x3e,
	0xb7, 0x8c, 0x57, 0x17, 0x56, 0xe2, 0xf4, 0xc2, 0x4a, 0x7c, 0xbc, 0xb0, 0x12, 0x4f, 0xca, 0x91,
	0x27, 0x26, 0x72, 0x4b, 0x72, 0x2f, 0x84, 0x0b, 0x51, 0x47, 0xe1, 0x79, 0xfc, 0x3f, 0x49, 0xf0,
	0x08, 0xd5, 0xc7, 0x82, 0xeb, 0x72, 0xf7, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc9, 0x6b,
	0xe2, 0x7a, 0x06, 0x00, 0x00,
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
