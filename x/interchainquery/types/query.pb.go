// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainquery/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	types2 "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/cometbft/cometbft/proto/tendermint/types"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryRequestsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	ChainId    string             `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryRequestsRequest) Reset()         { *m = QueryRequestsRequest{} }
func (m *QueryRequestsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequestsRequest) ProtoMessage()    {}
func (*QueryRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4aadfdae61bcbb1, []int{0}
}
func (m *QueryRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequestsRequest.Merge(m, src)
}
func (m *QueryRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequestsRequest proto.InternalMessageInfo

func (m *QueryRequestsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryRequestsRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryRequestsResponse struct {
	// params defines the parameters of the module.
	Queries    []Query             `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRequestsResponse) Reset()         { *m = QueryRequestsResponse{} }
func (m *QueryRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRequestsResponse) ProtoMessage()    {}
func (*QueryRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4aadfdae61bcbb1, []int{1}
}
func (m *QueryRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequestsResponse.Merge(m, src)
}
func (m *QueryRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequestsResponse proto.InternalMessageInfo

func (m *QueryRequestsResponse) GetQueries() []Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *QueryRequestsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// GetTxResponse is the response type for the Service.GetTx method.
type GetTxWithProofResponse struct {
	// tx is the queried transaction.
	Tx *tx.Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	// tx_response is the queried TxResponses.
	TxResponse *types.TxResponse `protobuf:"bytes,2,opt,name=tx_response,json=txResponse,proto3" json:"tx_response,omitempty"`
	// proof is the tmproto.TxProof for the queried tx
	Proof *types1.TxProof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	// ibc-go header to validate txs
	Header *types2.Header `protobuf:"bytes,4,opt,name=header,proto3" json:"header,omitempty"`
}

func (m *GetTxWithProofResponse) Reset()         { *m = GetTxWithProofResponse{} }
func (m *GetTxWithProofResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxWithProofResponse) ProtoMessage()    {}
func (*GetTxWithProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4aadfdae61bcbb1, []int{2}
}
func (m *GetTxWithProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxWithProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxWithProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxWithProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxWithProofResponse.Merge(m, src)
}
func (m *GetTxWithProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTxWithProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxWithProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxWithProofResponse proto.InternalMessageInfo

func (m *GetTxWithProofResponse) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *GetTxWithProofResponse) GetTxResponse() *types.TxResponse {
	if m != nil {
		return m.TxResponse
	}
	return nil
}

func (m *GetTxWithProofResponse) GetProof() *types1.TxProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *GetTxWithProofResponse) GetHeader() *types2.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequestsRequest)(nil), "quicksilver.interchainquery.v1.QueryRequestsRequest")
	proto.RegisterType((*QueryRequestsResponse)(nil), "quicksilver.interchainquery.v1.QueryRequestsResponse")
	proto.RegisterType((*GetTxWithProofResponse)(nil), "quicksilver.interchainquery.v1.GetTxWithProofResponse")
}

func init() {
	proto.RegisterFile("quicksilver/interchainquery/v1/query.proto", fileDescriptor_e4aadfdae61bcbb1)
}

var fileDescriptor_e4aadfdae61bcbb1 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xcd, 0xa4, 0xfd, 0xda, 0xaf, 0xd3, 0x9d, 0xd5, 0xa2, 0x34, 0x42, 0x26, 0x0a, 0x14, 0xa2,
	0x4a, 0xcc, 0x28, 0x21, 0x6c, 0x58, 0x74, 0x51, 0xa9, 0x14, 0x76, 0xad, 0x89, 0x04, 0x62, 0x53,
	0xd9, 0xce, 0xe0, 0x8c, 0x48, 0x66, 0x1c, 0xcf, 0xb5, 0xe5, 0x08, 0xb1, 0xe1, 0x09, 0x90, 0x78,
	0x08, 0xf6, 0x3c, 0x01, 0xcb, 0x2e, 0x2b, 0xb1, 0x61, 0x85, 0x50, 0xc2, 0x3b, 0xb0, 0x45, 0x1e,
	0x8f, 0x8d, 0x13, 0x24, 0x5a, 0x36, 0xf6, 0x8c, 0xee, 0x39, 0xf7, 0x9c, 0xfb, 0x63, 0xe3, 0x83,
	0x69, 0xcc, 0xfd, 0xd7, 0x8a, 0x8f, 0x13, 0x16, 0x51, 0x2e, 0x80, 0x45, 0xfe, 0xc8, 0xe5, 0x62,
	0x1a, 0xb3, 0x68, 0x46, 0x93, 0x2e, 0xd5, 0x07, 0x12, 0x46, 0x12, 0xa4, 0x65, 0x57, 0xb0, 0x64,
	0x05, 0x4b, 0x92, 0x6e, 0x73, 0x27, 0x90, 0x81, 0xd4, 0x50, 0x9a, 0x9d, 0x72, 0x56, 0xf3, 0x66,
	0x20, 0x65, 0x30, 0x66, 0xd4, 0x0d, 0x39, 0x75, 0x85, 0x90, 0xe0, 0x02, 0x97, 0x42, 0x99, 0x68,
	0xff, 0x0a, 0xfd, 0x55, 0x99, 0x9c, 0x75, 0xe0, 0x4b, 0x35, 0x91, 0x8a, 0x7a, 0xae, 0x62, 0xb4,
	0xc0, 0x7a, 0x0c, 0xdc, 0x2e, 0x0d, 0xdd, 0x80, 0x0b, 0x2d, 0x61, 0xb0, 0xb7, 0xab, 0x58, 0xd7,
	0xf3, 0x79, 0x09, 0xcd, 0x2e, 0x06, 0xd4, 0x34, 0x20, 0x48, 0xcb, 0x28, 0xa4, 0x45, 0x01, 0xc0,
	0xc4, 0x90, 0x45, 0x13, 0x2e, 0x80, 0xc2, 0x2c, 0x64, 0x2a, 0x7f, 0x9a, 0x28, 0xe5, 0x9e, 0x4f,
	0xc7, 0x3c, 0x18, 0x81, 0x3f, 0xe6, 0x4c, 0x80, 0xa2, 0x15, 0x78, 0xd2, 0xad, 0xdc, 0x72, 0x42,
	0x7b, 0x86, 0x77, 0xce, 0x32, 0xc7, 0x0e, 0x9b, 0xc6, 0x4c, 0x81, 0x32, 0x6f, 0xeb, 0x31, 0xc6,
	0xbf, 0xbd, 0x37, 0x50, 0x0b, 0x75, 0xb6, 0x7b, 0x77, 0x49, 0xee, 0x8b, 0x64, 0xe6, 0x49, 0xd1,
	0x68, 0xed, 0x8f, 0x9c, 0xba, 0x01, 0x33, 0x5c, 0xa7, 0xc2, 0xb4, 0xf6, 0xf0, 0xff, 0xba, 0x5f,
	0xe7, 0x7c, 0xd8, 0xa8, 0xb7, 0x50, 0x67, 0xcb, 0xd9, 0xd4, 0xf7, 0xa7, 0xc3, 0xf6, 0x47, 0x84,
	0x77, 0x57, 0xb4, 0x55, 0x28, 0x85, 0x62, 0xd6, 0x31, 0xde, 0xcc, 0xb2, 0x73, 0xa6, 0x1a, 0xa8,
	0xb5, 0xd6, 0xd9, 0xee, 0xed, 0x93, 0xbf, 0x0f, 0x9b, 0xe8, 0x3c, 0x47, 0xeb, 0x17, 0xdf, 0x6e,
	0xd5, 0x9c, 0x82, 0x6b, 0x9d, 0x2c, 0xd5, 0x50, 0xd7, 0x35, 0xdc, 0xbb, 0xb2, 0x86, 0xdc, 0x43,
	0xb5, 0x88, 0xf6, 0x4f, 0x84, 0x6f, 0x9c, 0x30, 0x18, 0xa4, 0xcf, 0x39, 0x8c, 0x4e, 0x23, 0x29,
	0x5f, 0x95, 0x56, 0xf7, 0x71, 0x1d, 0x52, 0xd3, 0x9f, 0xdd, 0x22, 0x37, 0xa4, 0x65, 0xce, 0x41,
	0xea, 0xd4, 0x21, 0xb5, 0x8e, 0xf1, 0x36, 0xa4, 0xe7, 0x91, 0x61, 0x19, 0x2f, 0x77, 0x96, 0xbc,
	0xe8, 0xf9, 0x57, 0x68, 0xa5, 0x11, 0x28, 0xcf, 0x16, 0xc5, 0xff, 0x85, 0x99, 0x7c, 0x63, 0x4d,
	0x27, 0xd8, 0x23, 0x95, 0x79, 0xe6, 0x6b, 0x30, 0x48, 0x73, 0x7f, 0x39, 0xce, 0x3a, 0xc4, 0x1b,
	0x23, 0xe6, 0x0e, 0x59, 0xd4, 0x58, 0x37, 0x23, 0xe4, 0x9e, 0x4f, 0xaa, 0x0b, 0x52, 0x4d, 0x91,
	0x74, 0xc9, 0x13, 0x8d, 0x76, 0x0c, 0xab, 0xf7, 0x19, 0xe1, 0x2d, 0xdd, 0xdb, 0x67, 0x51, 0x12,
	0x59, 0x9f, 0x10, 0xde, 0x3c, 0x33, 0xcd, 0xed, 0x5f, 0x6b, 0x24, 0x2b, 0x6b, 0xd5, 0x7c, 0xf8,
	0x8f, 0xac, 0xbc, 0xee, 0xf6, 0xa3, 0x77, 0x5f, 0x7e, 0x7c, 0xa8, 0xf7, 0xad, 0x1e, 0xbd, 0xc6,
	0x0f, 0x82, 0x33, 0x45, 0xdf, 0x14, 0x4b, 0xf7, 0xf6, 0xe8, 0xc5, 0xc5, 0xdc, 0x46, 0x97, 0x73,
	0x1b, 0x7d, 0x9f, 0xdb, 0xe8, 0xfd, 0xc2, 0xae, 0x5d, 0x2e, 0xec, 0xda, 0xd7, 0x85, 0x5d, 0x7b,
	0x79, 0x18, 0x70, 0x18, 0xc5, 0x1e, 0xf1, 0xe5, 0x84, 0x72, 0x11, 0x30, 0x11, 0x73, 0x98, 0xdd,
	0xf7, 0x62, 0x3e, 0x1e, 0x2e, 0xe9, 0xa4, 0x7f, 0x28, 0xe9, 0x5e, 0x7b, 0x1b, 0xfa, 0x13, 0x7a,
	0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x89, 0x06, 0x2a, 0x3e, 0xb6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuerySrvrClient is the client API for QuerySrvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuerySrvrClient interface {
	// Params returns the total set of minting parameters.
	Queries(ctx context.Context, in *QueryRequestsRequest, opts ...grpc.CallOption) (*QueryRequestsResponse, error)
}

type querySrvrClient struct {
	cc grpc1.ClientConn
}

func NewQuerySrvrClient(cc grpc1.ClientConn) QuerySrvrClient {
	return &querySrvrClient{cc}
}

func (c *querySrvrClient) Queries(ctx context.Context, in *QueryRequestsRequest, opts ...grpc.CallOption) (*QueryRequestsResponse, error) {
	out := new(QueryRequestsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainquery.v1.QuerySrvr/Queries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerySrvrServer is the server API for QuerySrvr service.
type QuerySrvrServer interface {
	// Params returns the total set of minting parameters.
	Queries(context.Context, *QueryRequestsRequest) (*QueryRequestsResponse, error)
}

// UnimplementedQuerySrvrServer can be embedded to have forward compatible implementations.
type UnimplementedQuerySrvrServer struct {
}

func (*UnimplementedQuerySrvrServer) Queries(ctx context.Context, req *QueryRequestsRequest) (*QueryRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Queries not implemented")
}

func RegisterQuerySrvrServer(s grpc1.Server, srv QuerySrvrServer) {
	s.RegisterService(&_QuerySrvr_serviceDesc, srv)
}

func _QuerySrvr_Queries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerySrvrServer).Queries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainquery.v1.QuerySrvr/Queries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerySrvrServer).Queries(ctx, req.(*QueryRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuerySrvr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.interchainquery.v1.QuerySrvr",
	HandlerType: (*QuerySrvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Queries",
			Handler:    _QuerySrvr_Queries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/interchainquery/v1/query.proto",
}

func (m *QueryRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Queries) > 0 {
		for iNdEx := len(m.Queries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Queries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetTxWithProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxWithProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxWithProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TxResponse != nil {
		{
			size, err := m.TxResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
func (m *QueryRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetTxWithProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.TxResponse != nil {
		l = m.TxResponse.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequestsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
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
func (m *QueryRequestsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
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
			m.Queries = append(m.Queries, Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetTxWithProofResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetTxWithProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxWithProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxResponse", wireType)
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
			if m.TxResponse == nil {
				m.TxResponse = &types.TxResponse{}
			}
			if err := m.TxResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
			if m.Proof == nil {
				m.Proof = &types1.TxProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
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
			if m.Header == nil {
				m.Header = &types2.Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
