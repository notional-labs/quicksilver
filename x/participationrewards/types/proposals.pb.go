// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/proposals.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
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

type AddProtocolDataProposal struct {
	Title       string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        string                   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Data        encoding_json.RawMessage `protobuf:"bytes,5,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty" yaml:"data"`
	Key         string                   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
}

func (m *AddProtocolDataProposal) Reset()      { *m = AddProtocolDataProposal{} }
func (*AddProtocolDataProposal) ProtoMessage() {}
func (*AddProtocolDataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94433b2236a43ef, []int{0}
}
func (m *AddProtocolDataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddProtocolDataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddProtocolDataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddProtocolDataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProtocolDataProposal.Merge(m, src)
}
func (m *AddProtocolDataProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddProtocolDataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProtocolDataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddProtocolDataProposal proto.InternalMessageInfo

type AddProtocolDataProposalWithDeposit struct {
	Title       string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Protocol    string                   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty" yaml:"protocol"`
	Type        string                   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Key         string                   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Data        encoding_json.RawMessage `protobuf:"bytes,6,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty" yaml:"data"`
	Deposit     string                   `protobuf:"bytes,7,opt,name=deposit,proto3" json:"deposit,omitempty" yaml:"deposit"`
}

func (m *AddProtocolDataProposalWithDeposit) Reset()         { *m = AddProtocolDataProposalWithDeposit{} }
func (m *AddProtocolDataProposalWithDeposit) String() string { return proto.CompactTextString(m) }
func (*AddProtocolDataProposalWithDeposit) ProtoMessage()    {}
func (*AddProtocolDataProposalWithDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94433b2236a43ef, []int{1}
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddProtocolDataProposalWithDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProtocolDataProposalWithDeposit.Merge(m, src)
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Size() int {
	return m.Size()
}
func (m *AddProtocolDataProposalWithDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProtocolDataProposalWithDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_AddProtocolDataProposalWithDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddProtocolDataProposal)(nil), "quicksilver.participationrewards.v1.AddProtocolDataProposal")
	proto.RegisterType((*AddProtocolDataProposalWithDeposit)(nil), "quicksilver.participationrewards.v1.AddProtocolDataProposalWithDeposit")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/proposals.proto", fileDescriptor_d94433b2236a43ef)
}

var fileDescriptor_d94433b2236a43ef = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0x4e, 0xda, 0xbb, 0xb6, 0xb8, 0x55, 0x2b, 0x85, 0x4a, 0x58, 0x0c, 0xf1, 0xc9, 0x5d, 0x90,
	0x80, 0x58, 0x55, 0xb7, 0x5b, 0x10, 0xa5, 0x2b, 0x52, 0x95, 0x05, 0x09, 0x89, 0xc1, 0x17, 0x5b,
	0xa9, 0xb9, 0x34, 0x36, 0xb1, 0x73, 0x25, 0xff, 0xa0, 0x23, 0x23, 0xe3, 0xfd, 0x1c, 0x16, 0xa4,
	0x8e, 0x4c, 0x11, 0xba, 0x1b, 0x60, 0xce, 0xc8, 0x84, 0xec, 0xe4, 0x8e, 0x1b, 0xee, 0x24, 0x74,
	0x9b, 0xfd, 0xbe, 0xef, 0x7b, 0xef, 0x7d, 0x9f, 0x1e, 0xb8, 0xf8, 0x54, 0x8a, 0x64, 0xac, 0x45,
	0x36, 0xe1, 0x05, 0x51, 0xb4, 0x30, 0x22, 0x11, 0x8a, 0x1a, 0x21, 0xf3, 0x82, 0xdf, 0xd1, 0x82,
	0x69, 0x32, 0x39, 0x27, 0xaa, 0x90, 0x4a, 0x6a, 0x9a, 0xe9, 0x48, 0x15, 0xd2, 0xc8, 0xe0, 0x6c,
	0x45, 0x14, 0xad, 0x13, 0x45, 0x93, 0xf3, 0xa7, 0xa7, 0xa9, 0x4c, 0xa5, 0xe3, 0x13, 0xfb, 0x6a,
	0xa5, 0xf8, 0x97, 0x0f, 0x9e, 0xbc, 0x66, 0xec, 0xda, 0x7e, 0x12, 0x99, 0x5d, 0x51, 0x43, 0xaf,
	0xbb, 0xee, 0xc1, 0x29, 0xe8, 0x1b, 0x61, 0x32, 0x0e, 0xfd, 0x81, 0xff, 0xec, 0x51, 0xdc, 0x7e,
	0x82, 0x01, 0x38, 0x64, 0x5c, 0x27, 0x85, 0x50, 0x76, 0x00, 0xdc, 0x71, 0xd8, 0x6a, 0x29, 0x38,
	0x03, 0x3d, 0x53, 0x29, 0x0e, 0x7b, 0x16, 0xba, 0x3c, 0x69, 0x6a, 0x74, 0x58, 0xd1, 0xdb, 0x6c,
	0x88, 0x6d, 0x15, 0xc7, 0x0e, 0x0c, 0x5e, 0x81, 0x1e, 0xa3, 0x86, 0xc2, 0xbe, 0x23, 0x3d, 0xff,
	0x47, 0xb2, 0x55, 0xfc, 0xa7, 0x46, 0x90, 0xe7, 0x89, 0x64, 0x22, 0x4f, 0xc9, 0x47, 0x2d, 0xf3,
	0x28, 0xa6, 0x77, 0x6f, 0xb9, 0xd6, 0x34, 0xe5, 0xb1, 0x13, 0x06, 0x03, 0xb0, 0x3b, 0xe6, 0x15,
	0xdc, 0x73, 0xfa, 0xe3, 0xa6, 0x46, 0xa0, 0xd5, 0x8f, 0x79, 0x85, 0x63, 0x0b, 0x0d, 0x8f, 0xee,
	0xa7, 0xc8, 0xfb, 0x3a, 0x45, 0xde, 0xef, 0x29, 0xf2, 0xf0, 0xf7, 0x1d, 0x80, 0x37, 0x38, 0x7d,
	0x27, 0xcc, 0xcd, 0x15, 0x57, 0x52, 0x0b, 0xb3, 0xb5, 0x69, 0x02, 0x0e, 0x54, 0xd7, 0x1a, 0xee,
	0xba, 0x9d, 0x1e, 0x37, 0x35, 0x3a, 0x69, 0x77, 0x5a, 0x20, 0x38, 0x5e, 0x92, 0xfe, 0x2f, 0xa5,
	0xce, 0x64, 0x7f, 0xa3, 0xc9, 0x65, 0x8e, 0x36, 0x87, 0xa3, 0x6d, 0x72, 0x7c, 0x01, 0xf6, 0x59,
	0xeb, 0x1d, 0xee, 0xbb, 0x31, 0x41, 0x53, 0xa3, 0xe3, 0xae, 0x47, 0x0b, 0xe0, 0x78, 0x41, 0x19,
	0x1e, 0xdc, 0x77, 0x79, 0x5e, 0x7e, 0xf8, 0x36, 0x0b, 0xfd, 0x87, 0x59, 0xe8, 0xff, 0x9c, 0x85,
	0xfe, 0x97, 0x79, 0xe8, 0x3d, 0xcc, 0x43, 0xef, 0xc7, 0x3c, 0xf4, 0xde, 0xbf, 0x49, 0x85, 0xb9,
	0x29, 0x47, 0x51, 0x22, 0x6f, 0x89, 0xc8, 0x53, 0x9e, 0x97, 0xc2, 0x54, 0x2f, 0x47, 0xa5, 0xc8,
	0x18, 0x59, 0x3d, 0xef, 0xcf, 0xeb, 0x0f, 0xdc, 0x1a, 0xd7, 0xa3, 0x3d, 0x17, 0xd4, 0xc5, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x03, 0x63, 0x3d, 0x11, 0x03, 0x00, 0x00,
}

func (m *AddProtocolDataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddProtocolDataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddProtocolDataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddProtocolDataProposalWithDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddProtocolDataProposalWithDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddProtocolDataProposalWithDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddProtocolDataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func (m *AddProtocolDataProposalWithDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddProtocolDataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddProtocolDataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddProtocolDataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = encoding_json.RawMessage(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *AddProtocolDataProposalWithDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddProtocolDataProposalWithDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddProtocolDataProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)
