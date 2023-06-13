// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gashub/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryParamsRequest defines the request type for querying x/gashub parameters.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af85680fb3beada8, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse defines the response type for querying x/gashub parameters.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af85680fb3beada8, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryMsgGasParamsRequest defines the RPC request for looking up MsgGasParams entries.
type QueryMsgGasParamsRequest struct {
	// msg_type_urls is the specific type urls you want look up. Leave empty to get all entries.
	MsgTypeUrls []string `protobuf:"bytes,1,rep,name=msg_type_urls,json=msgTypeUrls,proto3" json:"msg_type_urls,omitempty"`
	// pagination defines an optional pagination for the request. This field is
	// only read if the msg_type_urls field is empty.
	Pagination *query.PageRequest `protobuf:"bytes,99,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryMsgGasParamsRequest) Reset()         { *m = QueryMsgGasParamsRequest{} }
func (m *QueryMsgGasParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMsgGasParamsRequest) ProtoMessage()    {}
func (*QueryMsgGasParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af85680fb3beada8, []int{2}
}
func (m *QueryMsgGasParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMsgGasParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMsgGasParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMsgGasParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMsgGasParamsRequest.Merge(m, src)
}
func (m *QueryMsgGasParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMsgGasParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMsgGasParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMsgGasParamsRequest proto.InternalMessageInfo

func (m *QueryMsgGasParamsRequest) GetMsgTypeUrls() []string {
	if m != nil {
		return m.MsgTypeUrls
	}
	return nil
}

func (m *QueryMsgGasParamsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryMsgGasParamsResponse defines the RPC response of a MsgGasParams query.
type QueryMsgGasParamsResponse struct {
	MsgGasParams []*MsgGasParams `protobuf:"bytes,1,rep,name=msg_gas_params,json=msgGasParams,proto3" json:"msg_gas_params,omitempty"`
	// pagination defines the pagination in the response. This field is only
	// populated if the msg_type_urls field in the request is empty.
	Pagination *query.PageResponse `protobuf:"bytes,99,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryMsgGasParamsResponse) Reset()         { *m = QueryMsgGasParamsResponse{} }
func (m *QueryMsgGasParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMsgGasParamsResponse) ProtoMessage()    {}
func (*QueryMsgGasParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af85680fb3beada8, []int{3}
}
func (m *QueryMsgGasParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMsgGasParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMsgGasParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMsgGasParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMsgGasParamsResponse.Merge(m, src)
}
func (m *QueryMsgGasParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMsgGasParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMsgGasParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMsgGasParamsResponse proto.InternalMessageInfo

func (m *QueryMsgGasParamsResponse) GetMsgGasParams() []*MsgGasParams {
	if m != nil {
		return m.MsgGasParams
	}
	return nil
}

func (m *QueryMsgGasParamsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.gashub.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.gashub.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryMsgGasParamsRequest)(nil), "cosmos.gashub.v1beta1.QueryMsgGasParamsRequest")
	proto.RegisterType((*QueryMsgGasParamsResponse)(nil), "cosmos.gashub.v1beta1.QueryMsgGasParamsResponse")
}

func init() { proto.RegisterFile("cosmos/gashub/v1beta1/query.proto", fileDescriptor_af85680fb3beada8) }

var fileDescriptor_af85680fb3beada8 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x3b, 0xbb, 0x58, 0xd8, 0xe9, 0x2a, 0x38, 0xae, 0x50, 0xa3, 0x9b, 0xd5, 0x88, 0x6e,
	0xad, 0x98, 0x71, 0xe3, 0x17, 0x90, 0x05, 0x5d, 0x3c, 0x08, 0x1a, 0x14, 0xc1, 0x4b, 0x99, 0xd4,
	0x61, 0x0c, 0x36, 0x99, 0x6c, 0xde, 0x64, 0xb1, 0x57, 0x0f, 0x22, 0x78, 0x11, 0xfc, 0x0c, 0x82,
	0x37, 0xfd, 0x02, 0xde, 0xf7, 0xb8, 0xe0, 0xc5, 0x93, 0x48, 0x2b, 0xf8, 0x35, 0x64, 0xfe, 0x74,
	0x9b, 0x60, 0xea, 0xf6, 0xd2, 0x26, 0x33, 0xcf, 0xfb, 0xbc, 0xbf, 0x67, 0xde, 0x09, 0xbe, 0x32,
	0x94, 0x90, 0x48, 0xa0, 0x82, 0xc1, 0xcb, 0x32, 0xa2, 0x07, 0x3b, 0x11, 0x2f, 0xd8, 0x0e, 0xdd,
	0x2f, 0x79, 0x3e, 0xf6, 0xb3, 0x5c, 0x16, 0x92, 0x9c, 0x37, 0x12, 0xdf, 0x48, 0x7c, 0x2b, 0x71,
	0x36, 0x84, 0x14, 0x52, 0x2b, 0xa8, 0x7a, 0x32, 0x62, 0xe7, 0x92, 0x90, 0x52, 0x8c, 0x38, 0x65,
	0x59, 0x4c, 0x59, 0x9a, 0xca, 0x82, 0x15, 0xb1, 0x4c, 0xc1, 0xee, 0x7a, 0xcd, 0xdd, 0xac, 0xb3,
	0xd1, 0x9c, 0x65, 0x49, 0x9c, 0x4a, 0xaa, 0x7f, 0xed, 0xd2, 0x45, 0x5b, 0xa6, 0xa9, 0xe8, 0x41,
	0x0d, 0xcf, 0xe9, 0xdb, 0xcd, 0x88, 0x01, 0x3f, 0x56, 0x18, 0xdf, 0x8c, 0x89, 0x38, 0xd5, 0x00,
	0x46, 0xeb, 0x6d, 0x60, 0xf2, 0x58, 0x29, 0x1e, 0xb1, 0x9c, 0x25, 0x10, 0xf2, 0xfd, 0x92, 0x43,
	0xe1, 0x3d, 0xc3, 0xe7, 0x6a, 0xab, 0x90, 0xc9, 0x14, 0x38, 0xb9, 0x8b, 0xdb, 0x99, 0x5e, 0xe9,
	0xa2, 0xcb, 0xa8, 0xd7, 0x09, 0x36, 0xfd, 0xc6, 0x83, 0xf0, 0x4d, 0xd9, 0xee, 0xda, 0xe1, 0xcf,
	0xad, 0xd6, 0xe7, 0x3f, 0x5f, 0xfb, 0x28, 0xb4, 0x75, 0xde, 0x5b, 0x84, 0xbb, 0xda, 0xf9, 0x21,
	0x88, 0x3d, 0x06, 0xb5, 0xae, 0xc4, 0xc3, 0xa7, 0x13, 0x10, 0x83, 0x62, 0x9c, 0xf1, 0x41, 0x99,
	0x8f, 0x54, 0x97, 0xd5, 0xde, 0x5a, 0xd8, 0x49, 0x40, 0x3c, 0x19, 0x67, 0xfc, 0x69, 0x3e, 0x02,
	0x72, 0x1f, 0xe3, 0x79, 0x86, 0xee, 0x50, 0x63, 0x5c, 0x9f, 0x61, 0xa8, 0xc0, 0xbe, 0x39, 0x89,
	0x39, 0x8a, 0xe0, 0xd6, 0x3f, 0xac, 0x54, 0x7a, 0x5f, 0x10, 0xbe, 0xd0, 0x00, 0x62, 0x83, 0x3e,
	0xc0, 0x67, 0x14, 0x89, 0x60, 0x30, 0x38, 0x0e, 0xbc, 0xda, 0xeb, 0x04, 0x57, 0x17, 0x04, 0xae,
	0x99, 0xac, 0x27, 0x95, 0x37, 0xb2, 0xd7, 0x00, 0xbc, 0x7d, 0x22, 0xb0, 0xe1, 0xa8, 0x12, 0x07,
	0xdf, 0x56, 0xf0, 0x29, 0x4d, 0x4c, 0xde, 0x23, 0xdc, 0xb6, 0xee, 0x37, 0x16, 0x00, 0xfd, 0x3b,
	0x53, 0xa7, 0xbf, 0x8c, 0xd4, 0xf4, 0xf5, 0xfa, 0xef, 0xd4, 0xd4, 0xde, 0x7c, 0xff, 0xfd, 0x71,
	0x65, 0x8b, 0x6c, 0xd2, 0xe6, 0x3b, 0x6a, 0x4e, 0x86, 0x7c, 0x42, 0x78, 0xbd, 0x9a, 0x9f, 0xd0,
	0xff, 0x35, 0x6a, 0x98, 0xbb, 0x73, 0x7b, 0xf9, 0x02, 0xcb, 0x17, 0xcc, 0xf9, 0xb6, 0xc9, 0xb5,
	0x05, 0x7c, 0xf5, 0x09, 0xee, 0xde, 0x3b, 0x9c, 0xb8, 0xe8, 0x68, 0xe2, 0xa2, 0x5f, 0x13, 0x17,
	0x7d, 0x98, 0xba, 0xad, 0xa3, 0xa9, 0xdb, 0xfa, 0x31, 0x75, 0x5b, 0xcf, 0x6f, 0x8a, 0xb8, 0x50,
	0x7d, 0x87, 0x32, 0x99, 0x59, 0x99, 0xbf, 0x5b, 0xf0, 0xe2, 0x15, 0x7d, 0x3d, 0xf3, 0x55, 0xf7,
	0x12, 0xa2, 0xb6, 0xfe, 0x6e, 0xee, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x33, 0x63, 0x00, 0x9e,
	0x27, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of x/gashub module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// MsgGasParams queries for MsgGasParams entries.
	//
	// This query only returns params that have specific MsgGasParams settings.
	// Any msg type that does not have a specific setting will not be returned by this query.
	MsgGasParams(ctx context.Context, in *QueryMsgGasParamsRequest, opts ...grpc.CallOption) (*QueryMsgGasParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gashub.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MsgGasParams(ctx context.Context, in *QueryMsgGasParamsRequest, opts ...grpc.CallOption) (*QueryMsgGasParamsResponse, error) {
	out := new(QueryMsgGasParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gashub.v1beta1.Query/MsgGasParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of x/gashub module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// MsgGasParams queries for MsgGasParams entries.
	//
	// This query only returns params that have specific MsgGasParams settings.
	// Any msg type that does not have a specific setting will not be returned by this query.
	MsgGasParams(context.Context, *QueryMsgGasParamsRequest) (*QueryMsgGasParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) MsgGasParams(ctx context.Context, req *QueryMsgGasParamsRequest) (*QueryMsgGasParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgGasParams not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gashub.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MsgGasParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMsgGasParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MsgGasParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gashub.v1beta1.Query/MsgGasParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MsgGasParams(ctx, req.(*QueryMsgGasParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gashub.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "MsgGasParams",
			Handler:    _Query_MsgGasParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/gashub/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryMsgGasParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMsgGasParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMsgGasParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0x9a
	}
	if len(m.MsgTypeUrls) > 0 {
		for iNdEx := len(m.MsgTypeUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MsgTypeUrls[iNdEx])
			copy(dAtA[i:], m.MsgTypeUrls[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.MsgTypeUrls[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryMsgGasParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMsgGasParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMsgGasParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0x9a
	}
	if len(m.MsgGasParams) > 0 {
		for iNdEx := len(m.MsgGasParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MsgGasParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryMsgGasParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgTypeUrls) > 0 {
		for _, s := range m.MsgTypeUrls {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 2 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryMsgGasParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgGasParams) > 0 {
		for _, e := range m.MsgGasParams {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 2 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryMsgGasParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMsgGasParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMsgGasParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrls", wireType)
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
			m.MsgTypeUrls = append(m.MsgTypeUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 99:
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
func (m *QueryMsgGasParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMsgGasParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMsgGasParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgGasParams", wireType)
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
			m.MsgGasParams = append(m.MsgGasParams, &MsgGasParams{})
			if err := m.MsgGasParams[len(m.MsgGasParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 99:
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