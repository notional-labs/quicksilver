// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Minter represents the minting state.
type Minter struct {
	// current epoch provisions
	EpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=epoch_provisions,json=epochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"epoch_provisions" yaml:"epoch_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3179bac36b5b0964, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

type DistributionProportions struct {
	// staking defines the proportion of the minted minted_denom that is to be
	// allocated as staking rewards.
	Staking github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=staking,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking" yaml:"staking"`
	// pool_incentives defines the proportion of the minted minted_denom that is
	// to be allocated as pool incentives.
	PoolIncentives github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=pool_incentives,json=poolIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"pool_incentives" yaml:"pool_incentives"`
	// participation_rewards defines the proportion of the minted minted_denom that is
	// to be allocated to participation rewards address.
	ParticipationRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=participation_rewards,json=participationRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"participation_rewards" yaml:"participation_rewards"`
	// community_pool defines the proportion of the minted minted_denom that is
	// to be allocated to the community pool.
	CommunityPool github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool" yaml:"community_pool"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_3179bac36b5b0964, []int{1}
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

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// epoch provisions from the first epoch
	GenesisEpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_epoch_provisions" yaml:"genesis_epoch_provisions"`
	// mint epoch identifier
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty" yaml:"epoch_identifier"`
	// number of epochs take to reduce rewards
	ReductionPeriodInEpochs int64 `protobuf:"varint,4,opt,name=reduction_period_in_epochs,json=reductionPeriodInEpochs,proto3" json:"reduction_period_in_epochs,omitempty" yaml:"reduction_period_in_epochs"`
	// reduction multiplier to execute on each period
	ReductionFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor" yaml:"reduction_factor"`
	// distribution_proportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,6,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	// start epoch to distribute minting rewards
	MintingRewardsDistributionStartEpoch int64 `protobuf:"varint,7,opt,name=minting_rewards_distribution_start_epoch,json=mintingRewardsDistributionStartEpoch,proto3" json:"minting_rewards_distribution_start_epoch,omitempty" yaml:"minting_rewards_distribution_start_epoch"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3179bac36b5b0964, []int{2}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *Params) GetReductionPeriodInEpochs() int64 {
	if m != nil {
		return m.ReductionPeriodInEpochs
	}
	return 0
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetMintingRewardsDistributionStartEpoch() int64 {
	if m != nil {
		return m.MintingRewardsDistributionStartEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "quicksilver.mint.v1beta1.Minter")
	proto.RegisterType((*DistributionProportions)(nil), "quicksilver.mint.v1beta1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "quicksilver.mint.v1beta1.Params")
}

func init() {
	proto.RegisterFile("quicksilver/mint/v1beta1/mint.proto", fileDescriptor_3179bac36b5b0964)
}

var fileDescriptor_3179bac36b5b0964 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0x8e, 0x7f, 0x6d, 0x53, 0xf5, 0x7e, 0xa2, 0x45, 0x56, 0xdb, 0x98, 0x02, 0x76, 0x31, 0x7f,
	0xd4, 0xa5, 0xb1, 0x4a, 0xb7, 0x4e, 0x28, 0x2a, 0x85, 0x0c, 0x45, 0xc1, 0x6c, 0x5d, 0x2c, 0xff,
	0xb9, 0xba, 0xaf, 0x1a, 0xdf, 0xb9, 0x77, 0xe7, 0x40, 0x16, 0x16, 0xc6, 0x2e, 0x8c, 0x8c, 0x7c,
	0x9c, 0x8e, 0x1d, 0x11, 0x43, 0x40, 0xed, 0x37, 0xe8, 0x27, 0x40, 0x77, 0xe7, 0xa4, 0x89, 0x21,
	0x12, 0x11, 0x53, 0x7c, 0xcf, 0xfb, 0xf8, 0x79, 0x9f, 0xbc, 0x7f, 0xce, 0xe8, 0xf1, 0x59, 0x01,
	0xf1, 0x29, 0x87, 0x6e, 0x0f, 0x33, 0x2f, 0x03, 0x22, 0xbc, 0xde, 0x4e, 0x84, 0x45, 0xb8, 0xa3,
	0x0e, 0xcd, 0x9c, 0x51, 0x41, 0x4d, 0x6b, 0x8c, 0xd4, 0x54, 0x78, 0x49, 0xda, 0x58, 0x4d, 0x69,
	0x4a, 0x15, 0xc9, 0x93, 0x4f, 0x9a, 0xbf, 0xe1, 0xa4, 0x94, 0xa6, 0x5d, 0xec, 0xa9, 0x53, 0x54,
	0x1c, 0x7b, 0x02, 0x32, 0xcc, 0x45, 0x98, 0xe5, 0x25, 0xe1, 0x5e, 0x95, 0x10, 0x92, 0x7e, 0x19,
	0xb2, 0xab, 0xa1, 0xa4, 0x60, 0xa1, 0x00, 0x4a, 0x74, 0xdc, 0xfd, 0x88, 0xea, 0x87, 0x40, 0x04,
	0x66, 0xa6, 0x40, 0x77, 0x71, 0x4e, 0xe3, 0x93, 0x20, 0x67, 0xb4, 0x07, 0x1c, 0x28, 0xe1, 0x96,
	0xb1, 0x69, 0x6c, 0x2d, 0xb5, 0xda, 0x17, 0x03, 0xa7, 0xf6, 0x7d, 0xe0, 0x3c, 0x4b, 0x41, 0x9c,
	0x14, 0x51, 0x33, 0xa6, 0x99, 0x17, 0x53, 0x9e, 0x51, 0x5e, 0xfe, 0x6c, 0xf3, 0xe4, 0xd4, 0x13,
	0xfd, 0x1c, 0xf3, 0xe6, 0x3e, 0x8e, 0x6f, 0x06, 0x4e, 0xa3, 0x1f, 0x66, 0xdd, 0x3d, 0xb7, 0xaa,
	0xe7, 0xfa, 0x2b, 0x0a, 0xea, 0xdc, 0x22, 0x83, 0x39, 0xd4, 0xd8, 0x07, 0x2e, 0x18, 0x44, 0x85,
	0xb4, 0xd5, 0x61, 0x34, 0xa7, 0x4c, 0x3e, 0x71, 0xf3, 0x08, 0x2d, 0x72, 0x11, 0x9e, 0x02, 0x49,
	0x4b, 0x23, 0x2f, 0x66, 0x36, 0xb2, 0xac, 0x8d, 0x94, 0x32, 0xae, 0x3f, 0x14, 0x34, 0xcf, 0xd0,
	0x4a, 0x4e, 0x69, 0x37, 0x00, 0x12, 0x63, 0x22, 0xa0, 0x87, 0xb9, 0xf5, 0x9f, 0xca, 0xf1, 0x7a,
	0xe6, 0x1c, 0xeb, 0x3a, 0x47, 0x45, 0xce, 0xf5, 0x97, 0x25, 0xd2, 0x1e, 0x01, 0xe6, 0x27, 0x03,
	0xad, 0xe5, 0x21, 0x13, 0x10, 0x43, 0xae, 0x5a, 0x10, 0x30, 0xfc, 0x3e, 0x64, 0x09, 0xb7, 0xe6,
	0x54, 0xe6, 0x37, 0x33, 0x67, 0x7e, 0x50, 0x66, 0xfe, 0x93, 0xa8, 0xeb, 0xaf, 0x4e, 0xe0, 0xbe,
	0x86, 0x4d, 0x82, 0x96, 0x63, 0x9a, 0x65, 0x05, 0x01, 0xd1, 0x0f, 0xa4, 0x43, 0x6b, 0x5e, 0x65,
	0x7f, 0x35, 0x73, 0xf6, 0x35, 0x9d, 0x7d, 0x52, 0xcd, 0xf5, 0xef, 0x8c, 0x80, 0x8e, 0x3c, 0xff,
	0x58, 0x40, 0xf5, 0x4e, 0xc8, 0xc2, 0x8c, 0x9b, 0x0f, 0x11, 0x92, 0xd3, 0x1e, 0x24, 0x98, 0xd0,
	0x4c, 0xb7, 0xd4, 0x5f, 0x92, 0xc8, 0xbe, 0x04, 0xcc, 0x73, 0x03, 0x59, 0x29, 0x26, 0x98, 0x03,
	0x0f, 0x7e, 0x9b, 0x44, 0xdd, 0x9c, 0xb7, 0x33, 0x9b, 0x74, 0xb4, 0xc9, 0x69, 0xba, 0xae, 0xbf,
	0x5e, 0x86, 0x5e, 0x4e, 0x0e, 0xa6, 0x79, 0x30, 0x5c, 0x07, 0x48, 0x64, 0x03, 0x8f, 0x01, 0xb3,
	0xb2, 0x4f, 0xf7, 0xab, 0x03, 0x7e, 0xcb, 0x18, 0x0e, 0x78, 0x7b, 0x84, 0x98, 0x11, 0xda, 0x60,
	0x38, 0x29, 0x62, 0xd5, 0x9b, 0x1c, 0x33, 0xa0, 0x49, 0x00, 0x44, 0x1b, 0xe1, 0xaa, 0xf6, 0x73,
	0xad, 0xa7, 0x37, 0x03, 0xe7, 0x91, 0x56, 0x9c, 0xce, 0x75, 0xfd, 0xc6, 0x28, 0xd8, 0x51, 0xb1,
	0x36, 0x51, 0xa6, 0xb9, 0x5c, 0xdd, 0xdb, 0xf7, 0x8e, 0xc3, 0x58, 0x50, 0x66, 0x2d, 0xfc, 0xdb,
	0xea, 0x56, 0xf5, 0x5c, 0x7f, 0x65, 0x04, 0x1d, 0x28, 0xc4, 0x64, 0xc8, 0x4a, 0xc6, 0x36, 0x57,
	0x56, 0x75, 0xb8, 0xba, 0x56, 0x7d, 0xd3, 0xd8, 0xfa, 0xff, 0xf9, 0x4e, 0x73, 0xda, 0x4d, 0xd7,
	0x9c, 0xb2, 0xf3, 0xad, 0x79, 0x69, 0xd8, 0x6f, 0x24, 0x53, 0xae, 0x84, 0x73, 0x03, 0x6d, 0x49,
	0x1d, 0x20, 0xe9, 0x70, 0xd0, 0x83, 0x09, 0x13, 0x5c, 0x84, 0x4c, 0xe8, 0x8a, 0x59, 0x8b, 0xaa,
	0xb8, 0xbb, 0x37, 0x03, 0xc7, 0xd3, 0x7f, 0xea, 0x6f, 0xdf, 0x74, 0xfd, 0x27, 0x25, 0xb5, 0xdc,
	0x9a, 0x71, 0xb7, 0xef, 0x24, 0x4f, 0x15, 0x7e, 0x6f, 0xfe, 0xcb, 0x57, 0xa7, 0xd6, 0x3a, 0xbc,
	0xb8, 0xb2, 0x8d, 0xcb, 0x2b, 0xdb, 0xf8, 0x79, 0x65, 0x1b, 0x9f, 0xaf, 0xed, 0xda, 0xe5, 0xb5,
	0x5d, 0xfb, 0x76, 0x6d, 0xd7, 0x8e, 0x76, 0xc7, 0xaa, 0x0e, 0x24, 0xc5, 0xa4, 0x00, 0xd1, 0xdf,
	0x8e, 0x0a, 0xe8, 0x26, 0xde, 0xf8, 0x87, 0xe2, 0x83, 0xfe, 0x54, 0xa8, 0x36, 0x44, 0x75, 0x75,
	0x31, 0xef, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xa7, 0x6b, 0xfb, 0x4b, 0x06, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EpochProvisions.Size()
		i -= size
		if _, err := m.EpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ParticipationRewards.Size()
		i -= size
		if _, err := m.ParticipationRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PoolIncentives.Size()
		i -= size
		if _, err := m.PoolIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
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
	if m.MintingRewardsDistributionStartEpoch != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.MintingRewardsDistributionStartEpoch))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ReductionPeriodInEpochs != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ReductionPeriodInEpochs))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintMint(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.GenesisEpochProvisions.Size()
		i -= size
		if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Staking.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.PoolIncentives.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.ParticipationRewards.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.GenesisEpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.ReductionPeriodInEpochs != 0 {
		n += 1 + sovMint(uint64(m.ReductionPeriodInEpochs))
	}
	l = m.ReductionFactor.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.DistributionProportions.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.MintingRewardsDistributionStartEpoch != 0 {
		n += 1 + sovMint(uint64(m.MintingRewardsDistributionStartEpoch))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ParticipationRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInEpochs", wireType)
			}
			m.ReductionPeriodInEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingRewardsDistributionStartEpoch", wireType)
			}
			m.MintingRewardsDistributionStartEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintingRewardsDistributionStartEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)