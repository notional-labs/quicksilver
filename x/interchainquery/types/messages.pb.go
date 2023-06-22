// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainquery/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSubmitQueryResponse represents a message type to fulfil a query request.
type MsgSubmitQueryResponse struct {
	ChainId     string           `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	QueryId     string           `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty" yaml:"query_id"`
	Result      []byte           `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty" yaml:"result"`
	ProofOps    *crypto.ProofOps `protobuf:"bytes,4,opt,name=proof_ops,json=proofOps,proto3" json:"proof_ops,omitempty" yaml:"proof_ops"`
	Height      int64            `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty" yaml:"height"`
	FromAddress string           `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSubmitQueryResponse) Reset()         { *m = MsgSubmitQueryResponse{} }
func (m *MsgSubmitQueryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0640fcbc3e895a79, []int{0}
}
func (m *MsgSubmitQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponse proto.InternalMessageInfo

// MsgSubmitQueryResponseResponse defines the MsgSubmitQueryResponse response
// type.
type MsgSubmitQueryResponseResponse struct {
}

func (m *MsgSubmitQueryResponseResponse) Reset()         { *m = MsgSubmitQueryResponseResponse{} }
func (m *MsgSubmitQueryResponseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponseResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0640fcbc3e895a79, []int{1}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitQueryResponse)(nil), "quicksilver.interchainquery.v1.MsgSubmitQueryResponse")
	proto.RegisterType((*MsgSubmitQueryResponseResponse)(nil), "quicksilver.interchainquery.v1.MsgSubmitQueryResponseResponse")
}

func init() {
	proto.RegisterFile("quicksilver/interchainquery/v1/messages.proto", fileDescriptor_0640fcbc3e895a79)
}

var fileDescriptor_0640fcbc3e895a79 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3f, 0x8f, 0xd3, 0x3e,
	0x18, 0xc7, 0xeb, 0xf6, 0xf7, 0xeb, 0xdd, 0xe5, 0x8a, 0x80, 0x5c, 0x85, 0x42, 0x81, 0xa4, 0xca,
	0x42, 0x41, 0xd4, 0x56, 0x8b, 0xc4, 0x50, 0xa4, 0x93, 0xe8, 0x76, 0xc3, 0xf1, 0x27, 0xb7, 0x20,
	0x96, 0x2a, 0x6d, 0x7c, 0xae, 0x45, 0x63, 0xe7, 0x6c, 0xa7, 0xba, 0xae, 0x4c, 0x8c, 0x48, 0x2c,
	0x8c, 0x7d, 0x11, 0xbc, 0x01, 0x36, 0xd8, 0x4e, 0xb0, 0x30, 0x55, 0xa8, 0x65, 0x80, 0xb5, 0xaf,
	0x00, 0xc5, 0x4e, 0xa1, 0xba, 0xab, 0x18, 0x98, 0xfc, 0xe4, 0xfb, 0x7c, 0x9e, 0x7f, 0xf1, 0x63,
	0xab, 0x79, 0x92, 0xd2, 0xc1, 0x4b, 0x49, 0x47, 0x63, 0x2c, 0x10, 0x65, 0x0a, 0x8b, 0xc1, 0x30,
	0xa4, 0xec, 0x24, 0xc5, 0x62, 0x82, 0xc6, 0x2d, 0x14, 0x63, 0x29, 0x43, 0x82, 0x25, 0x4c, 0x04,
	0x57, 0xdc, 0x76, 0xd7, 0x70, 0x78, 0x0e, 0x87, 0xe3, 0x56, 0xed, 0xfa, 0x80, 0xcb, 0x98, 0xcb,
	0x9e, 0xa6, 0x91, 0xf9, 0x30, 0xa1, 0xb5, 0x2a, 0xe1, 0x84, 0x1b, 0x3d, 0xb3, 0x72, 0xf5, 0x26,
	0xe1, 0x9c, 0x8c, 0x30, 0x0a, 0x13, 0x8a, 0x42, 0xc6, 0xb8, 0x0a, 0x15, 0xe5, 0x6c, 0x15, 0x73,
	0x4b, 0x61, 0x16, 0x61, 0x11, 0x53, 0xa6, 0xd0, 0x40, 0x4c, 0x12, 0xc5, 0x51, 0x22, 0x38, 0x3f,
	0x36, 0x6e, 0xff, 0x67, 0xd1, 0xba, 0x76, 0x28, 0xc9, 0x51, 0xda, 0x8f, 0xa9, 0x7a, 0x96, 0xf5,
	0x10, 0x60, 0x99, 0x70, 0x26, 0xb1, 0x0d, 0xad, 0x6d, 0xdd, 0x59, 0x8f, 0x46, 0x0e, 0xa8, 0x83,
	0xc6, 0x4e, 0x77, 0x6f, 0x39, 0xf3, 0x2e, 0x4f, 0xc2, 0x78, 0xd4, 0xf1, 0x57, 0x1e, 0x3f, 0xd8,
	0xd2, 0xe6, 0x41, 0x94, 0xf1, 0x7a, 0x88, 0x8c, 0x2f, 0x9e, 0xe7, 0x57, 0x1e, 0x3f, 0xd8, 0xd2,
	0xe6, 0x41, 0x64, 0xdf, 0xb1, 0xca, 0x02, 0xcb, 0x74, 0xa4, 0x9c, 0x52, 0x1d, 0x34, 0x2a, 0xdd,
	0xab, 0xcb, 0x99, 0x77, 0xc9, 0xd0, 0x46, 0xf7, 0x83, 0x1c, 0xb0, 0x1f, 0x5b, 0x3b, 0xba, 0xe9,
	0x1e, 0x4f, 0xa4, 0xf3, 0x5f, 0x1d, 0x34, 0x76, 0xdb, 0x37, 0xe0, 0x9f, 0xc1, 0xa0, 0x19, 0x0c,
	0x3e, 0xcd, 0x98, 0x27, 0x89, 0xec, 0x56, 0x97, 0x33, 0xef, 0x8a, 0x49, 0xf5, 0x3b, 0xce, 0x0f,
	0xb6, 0x93, 0xdc, 0x9f, 0x95, 0x1e, 0x62, 0x4a, 0x86, 0xca, 0xf9, 0xbf, 0x0e, 0x1a, 0xa5, 0xf5,
	0xd2, 0x46, 0xf7, 0x83, 0x1c, 0xb0, 0x1f, 0x5a, 0x95, 0x63, 0xc1, 0xe3, 0x5e, 0x18, 0x45, 0x02,
	0x4b, 0xe9, 0x94, 0xf5, 0x64, 0xce, 0xe7, 0xf7, 0xcd, 0x6a, 0x7e, 0x37, 0x8f, 0x8c, 0xe7, 0x48,
	0x09, 0xca, 0x48, 0xb0, 0x9b, 0xd1, 0xb9, 0xd4, 0xa9, 0xbc, 0x9e, 0x7a, 0x85, 0x77, 0x53, 0x0f,
	0xfc, 0x98, 0x7a, 0x05, 0xbf, 0x6e, 0xb9, 0x9b, 0x7f, 0xf5, 0xea, 0x6c, 0x7f, 0x02, 0x56, 0xe9,
	0x50, 0x12, 0xfb, 0x03, 0xb0, 0xf6, 0x36, 0x5d, 0xc9, 0x03, 0xf8, 0xf7, 0xe5, 0x81, 0x9b, 0xf3,
	0xd7, 0xf6, 0xff, 0x2d, 0x6e, 0x75, 0xfa, 0xed, 0x57, 0x5f, 0xbe, 0xbf, 0x2d, 0xde, 0xf3, 0x6f,
	0x5f, 0xd8, 0x6f, 0x75, 0x8a, 0xc6, 0xad, 0x3e, 0x56, 0x61, 0x0b, 0x49, 0x9d, 0x40, 0xcb, 0x1d,
	0x70, 0xb7, 0xfb, 0xfc, 0xe3, 0xdc, 0x05, 0x67, 0x73, 0x17, 0x7c, 0x9b, 0xbb, 0xe0, 0xcd, 0xc2,
	0x2d, 0x9c, 0x2d, 0xdc, 0xc2, 0xd7, 0x85, 0x5b, 0x78, 0xb1, 0x4f, 0xa8, 0x1a, 0xa6, 0x7d, 0x38,
	0xe0, 0x31, 0xa2, 0x8c, 0x60, 0x96, 0x52, 0x35, 0x69, 0xf6, 0x53, 0x3a, 0x8a, 0xd0, 0xfa, 0x5b,
	0x3a, 0xbd, 0x58, 0x6d, 0x92, 0x60, 0xd9, 0x2f, 0xeb, 0xd5, 0xbd, 0xff, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x80, 0xb2, 0x83, 0xbd, 0x79, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error) {
	out := new(MsgSubmitQueryResponseResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainquery.v1.Msg/SubmitQueryResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(context.Context, *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitQueryResponse(ctx context.Context, req *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitQueryResponse not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitQueryResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitQueryResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitQueryResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainquery.v1.Msg/SubmitQueryResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitQueryResponse(ctx, req.(*MsgSubmitQueryResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.interchainquery.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitQueryResponse",
			Handler:    _Msg_SubmitQueryResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/interchainquery/v1/messages.proto",
}

func (m *MsgSubmitQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.Height != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x28
	}
	if m.ProofOps != nil {
		{
			size, err := m.ProofOps.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitQueryResponseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ProofOps != nil {
		l = m.ProofOps.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovMessages(uint64(m.Height))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSubmitQueryResponseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofOps == nil {
				m.ProofOps = &crypto.ProofOps{}
			}
			if err := m.ProofOps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSubmitQueryResponseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
