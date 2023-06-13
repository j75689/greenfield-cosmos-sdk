// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gashub/v1beta1/gashub.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the gashub module.
type Params struct {
	// max_tx_size is the maximum size of a transaction's bytes.
	MaxTxSize uint64 `protobuf:"varint,1,opt,name=max_tx_size,json=maxTxSize,proto3" json:"max_tx_size,omitempty"`
	// min_gas_per_byte is the minimum gas to be paid per byte of a transaction's
	MinGasPerByte uint64 `protobuf:"varint,2,opt,name=min_gas_per_byte,json=minGasPerByte,proto3" json:"min_gas_per_byte,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2f12e3606fbd41, []int{0}
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

func (m *Params) GetMaxTxSize() uint64 {
	if m != nil {
		return m.MaxTxSize
	}
	return 0
}

func (m *Params) GetMinGasPerByte() uint64 {
	if m != nil {
		return m.MinGasPerByte
	}
	return 0
}

// MsgGasParams defines gas consumption for a msg type
type MsgGasParams struct {
	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// gas_params is the oneof that represents either fixed_gas_params or dynamic_gas_params
	//
	// Types that are valid to be assigned to GasParams:
	//
	//	*MsgGasParams_FixedType
	//	*MsgGasParams_GrantType
	//	*MsgGasParams_MultiSendType
	//	*MsgGasParams_GrantAllowanceType
	GasParams isMsgGasParams_GasParams `protobuf_oneof:"gas_params"`
}

func (m *MsgGasParams) Reset()         { *m = MsgGasParams{} }
func (m *MsgGasParams) String() string { return proto.CompactTextString(m) }
func (*MsgGasParams) ProtoMessage()    {}
func (*MsgGasParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2f12e3606fbd41, []int{1}
}
func (m *MsgGasParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGasParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGasParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGasParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGasParams.Merge(m, src)
}
func (m *MsgGasParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgGasParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGasParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGasParams proto.InternalMessageInfo

type isMsgGasParams_GasParams interface {
	isMsgGasParams_GasParams()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type MsgGasParams_FixedType struct {
	FixedType *MsgGasParams_FixedGasParams `protobuf:"bytes,2,opt,name=fixed_type,json=fixedType,proto3,oneof" json:"fixed_type,omitempty"`
}
type MsgGasParams_GrantType struct {
	GrantType *MsgGasParams_DynamicGasParams `protobuf:"bytes,3,opt,name=grant_type,json=grantType,proto3,oneof" json:"grant_type,omitempty"`
}
type MsgGasParams_MultiSendType struct {
	MultiSendType *MsgGasParams_DynamicGasParams `protobuf:"bytes,4,opt,name=multi_send_type,json=multiSendType,proto3,oneof" json:"multi_send_type,omitempty"`
}
type MsgGasParams_GrantAllowanceType struct {
	GrantAllowanceType *MsgGasParams_DynamicGasParams `protobuf:"bytes,5,opt,name=grant_allowance_type,json=grantAllowanceType,proto3,oneof" json:"grant_allowance_type,omitempty"`
}

func (*MsgGasParams_FixedType) isMsgGasParams_GasParams()          {}
func (*MsgGasParams_GrantType) isMsgGasParams_GasParams()          {}
func (*MsgGasParams_MultiSendType) isMsgGasParams_GasParams()      {}
func (*MsgGasParams_GrantAllowanceType) isMsgGasParams_GasParams() {}

func (m *MsgGasParams) GetGasParams() isMsgGasParams_GasParams {
	if m != nil {
		return m.GasParams
	}
	return nil
}

func (m *MsgGasParams) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *MsgGasParams) GetFixedType() *MsgGasParams_FixedGasParams {
	if x, ok := m.GetGasParams().(*MsgGasParams_FixedType); ok {
		return x.FixedType
	}
	return nil
}

func (m *MsgGasParams) GetGrantType() *MsgGasParams_DynamicGasParams {
	if x, ok := m.GetGasParams().(*MsgGasParams_GrantType); ok {
		return x.GrantType
	}
	return nil
}

func (m *MsgGasParams) GetMultiSendType() *MsgGasParams_DynamicGasParams {
	if x, ok := m.GetGasParams().(*MsgGasParams_MultiSendType); ok {
		return x.MultiSendType
	}
	return nil
}

func (m *MsgGasParams) GetGrantAllowanceType() *MsgGasParams_DynamicGasParams {
	if x, ok := m.GetGasParams().(*MsgGasParams_GrantAllowanceType); ok {
		return x.GrantAllowanceType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MsgGasParams) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MsgGasParams_FixedType)(nil),
		(*MsgGasParams_GrantType)(nil),
		(*MsgGasParams_MultiSendType)(nil),
		(*MsgGasParams_GrantAllowanceType)(nil),
	}
}

// FixedGasParams defines the parameters for fixed gas type.
type MsgGasParams_FixedGasParams struct {
	// fixed_gas is the gas cost for a fixed type msg
	FixedGas uint64 `protobuf:"varint,1,opt,name=fixed_gas,json=fixedGas,proto3" json:"fixed_gas,omitempty"`
}

func (m *MsgGasParams_FixedGasParams) Reset()         { *m = MsgGasParams_FixedGasParams{} }
func (m *MsgGasParams_FixedGasParams) String() string { return proto.CompactTextString(m) }
func (*MsgGasParams_FixedGasParams) ProtoMessage()    {}
func (*MsgGasParams_FixedGasParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2f12e3606fbd41, []int{1, 0}
}
func (m *MsgGasParams_FixedGasParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGasParams_FixedGasParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGasParams_FixedGasParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGasParams_FixedGasParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGasParams_FixedGasParams.Merge(m, src)
}
func (m *MsgGasParams_FixedGasParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgGasParams_FixedGasParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGasParams_FixedGasParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGasParams_FixedGasParams proto.InternalMessageInfo

func (m *MsgGasParams_FixedGasParams) GetFixedGas() uint64 {
	if m != nil {
		return m.FixedGas
	}
	return 0
}

// DynamicGasParams defines the parameters for dynamic gas type.
type MsgGasParams_DynamicGasParams struct {
	// fixed_gas is the base gas cost for a dynamic type msg
	FixedGas uint64 `protobuf:"varint,1,opt,name=fixed_gas,json=fixedGas,proto3" json:"fixed_gas,omitempty"`
	// gas_per_item is the gas cost for a dynamic type msg per item
	GasPerItem uint64 `protobuf:"varint,2,opt,name=gas_per_item,json=gasPerItem,proto3" json:"gas_per_item,omitempty"`
}

func (m *MsgGasParams_DynamicGasParams) Reset()         { *m = MsgGasParams_DynamicGasParams{} }
func (m *MsgGasParams_DynamicGasParams) String() string { return proto.CompactTextString(m) }
func (*MsgGasParams_DynamicGasParams) ProtoMessage()    {}
func (*MsgGasParams_DynamicGasParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa2f12e3606fbd41, []int{1, 1}
}
func (m *MsgGasParams_DynamicGasParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGasParams_DynamicGasParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGasParams_DynamicGasParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGasParams_DynamicGasParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGasParams_DynamicGasParams.Merge(m, src)
}
func (m *MsgGasParams_DynamicGasParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgGasParams_DynamicGasParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGasParams_DynamicGasParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGasParams_DynamicGasParams proto.InternalMessageInfo

func (m *MsgGasParams_DynamicGasParams) GetFixedGas() uint64 {
	if m != nil {
		return m.FixedGas
	}
	return 0
}

func (m *MsgGasParams_DynamicGasParams) GetGasPerItem() uint64 {
	if m != nil {
		return m.GasPerItem
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.gashub.v1beta1.Params")
	proto.RegisterType((*MsgGasParams)(nil), "cosmos.gashub.v1beta1.MsgGasParams")
	proto.RegisterType((*MsgGasParams_FixedGasParams)(nil), "cosmos.gashub.v1beta1.MsgGasParams.FixedGasParams")
	proto.RegisterType((*MsgGasParams_DynamicGasParams)(nil), "cosmos.gashub.v1beta1.MsgGasParams.DynamicGasParams")
}

func init() {
	proto.RegisterFile("cosmos/gashub/v1beta1/gashub.proto", fileDescriptor_aa2f12e3606fbd41)
}

var fileDescriptor_aa2f12e3606fbd41 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x63, 0x48, 0xab, 0xe6, 0x6d, 0x52, 0x5a, 0xab, 0x48, 0x95, 0x07, 0x1b, 0x85, 0x85,
	0x3f, 0xaa, 0x4d, 0x0b, 0x53, 0xb6, 0x5a, 0xfc, 0x1d, 0x22, 0x55, 0x4e, 0xbb, 0x30, 0x60, 0x5d,
	0xd2, 0xeb, 0xf5, 0x84, 0xcf, 0x17, 0xf9, 0x2e, 0x60, 0x77, 0x66, 0x62, 0x62, 0x64, 0xe4, 0x23,
	0xf0, 0x31, 0x18, 0x3b, 0x32, 0x59, 0xc8, 0x19, 0xe0, 0x63, 0xa0, 0xbb, 0xb3, 0xa3, 0x16, 0x75,
	0xa0, 0xca, 0x62, 0xbd, 0x77, 0x7a, 0xdf, 0xdf, 0xf3, 0xf8, 0xd5, 0x73, 0xd0, 0x9f, 0x70, 0xc1,
	0xb8, 0x08, 0x08, 0x12, 0x67, 0xb3, 0x71, 0xf0, 0x61, 0x6f, 0x8c, 0x25, 0xda, 0xab, 0x8f, 0xfe,
	0x34, 0xe3, 0x92, 0xdb, 0x77, 0x4d, 0x8f, 0x5f, 0x5f, 0xd6, 0x3d, 0xce, 0x36, 0xe1, 0x84, 0xeb,
	0x8e, 0x40, 0x55, 0xa6, 0xd9, 0xd9, 0x42, 0x8c, 0xa6, 0x3c, 0xd0, 0x5f, 0x73, 0xd5, 0xff, 0x6a,
	0xc1, 0xea, 0x21, 0xca, 0x10, 0x13, 0xf6, 0x2e, 0xac, 0x33, 0x94, 0xc7, 0x32, 0x8f, 0x05, 0x3d,
	0xc7, 0x3b, 0xd6, 0x3d, 0xeb, 0x41, 0x3b, 0xec, 0x55, 0xa5, 0xd7, 0x19, 0xa2, 0xfc, 0x28, 0x1f,
	0xd1, 0x73, 0x1c, 0x75, 0x58, 0x53, 0xda, 0x03, 0xd8, 0x64, 0x34, 0x8d, 0x09, 0x12, 0xf1, 0x14,
	0x67, 0xf1, 0xb8, 0x90, 0x78, 0xe7, 0x96, 0x9e, 0xd9, 0xaa, 0x4a, 0xaf, 0x37, 0xa4, 0xe9, 0x2b,
	0x24, 0x0e, 0x71, 0x16, 0x16, 0x12, 0x47, 0x3d, 0x76, 0xf9, 0x38, 0xb8, 0xff, 0xe7, 0x9b, 0x67,
	0x7d, 0xfe, 0xfd, 0xfd, 0x91, 0x63, 0xec, 0xef, 0x8a, 0x93, 0xf7, 0x41, 0xde, 0xfc, 0xa8, 0xf1,
	0xd3, 0xff, 0xb4, 0x02, 0xdd, 0xa1, 0x20, 0x6a, 0xcc, 0x18, 0x7c, 0x02, 0x5d, 0x26, 0x48, 0x2c,
	0x8b, 0x29, 0x8e, 0x67, 0x59, 0xa2, 0x1d, 0x76, 0xc2, 0x8d, 0xaa, 0xf4, 0x60, 0x28, 0xc8, 0x51,
	0x31, 0xc5, 0xc7, 0x59, 0x12, 0x01, 0x5b, 0xd4, 0xf6, 0x08, 0xe0, 0x94, 0xe6, 0xf8, 0x44, 0xcf,
	0x68, 0x77, 0xeb, 0xfb, 0xfb, 0xfe, 0xb5, 0x2b, 0xf3, 0x2f, 0x4b, 0xf9, 0x2f, 0xd5, 0xd4, 0xe2,
	0xf8, 0xba, 0x15, 0x75, 0x34, 0x47, 0x71, 0xed, 0x63, 0x00, 0x92, 0xa1, 0x54, 0x1a, 0xe8, 0x6d,
	0x0d, 0x7d, 0xf6, 0x3f, 0xd0, 0xe7, 0x45, 0x8a, 0x18, 0x9d, 0x5c, 0xc1, 0x6a, 0x92, 0xc6, 0xbe,
	0x83, 0x3b, 0x6c, 0x96, 0x48, 0x1a, 0x0b, 0x9c, 0xd6, 0x86, 0xdb, 0x4b, 0xb1, 0x7b, 0x1a, 0x37,
	0xc2, 0xa9, 0xb1, 0x7d, 0x06, 0xdb, 0xc6, 0x36, 0x4a, 0x12, 0xfe, 0x11, 0xa5, 0x13, 0x6c, 0x44,
	0x56, 0x96, 0x12, 0xb1, 0x35, 0xf3, 0xa0, 0x41, 0x2a, 0x25, 0xe7, 0x00, 0x36, 0xae, 0xee, 0xcf,
	0x7e, 0x08, 0x66, 0x7f, 0x2a, 0x2d, 0x75, 0xb0, 0xba, 0x55, 0xe9, 0xad, 0x35, 0x6d, 0xd1, 0xda,
	0x69, 0x5d, 0x0d, 0xda, 0x2a, 0x1a, 0xce, 0x0c, 0x36, 0xff, 0x15, 0xbb, 0x01, 0x44, 0x25, 0xa5,
	0xc9, 0x25, 0x95, 0x98, 0xd5, 0xb9, 0xd4, 0x49, 0x31, 0x29, 0x7c, 0x23, 0x31, 0x8b, 0x80, 0x2c,
	0x6a, 0x23, 0x6b, 0xbe, 0x61, 0x17, 0x40, 0x4f, 0x6b, 0xd9, 0xf0, 0xc5, 0x8f, 0xca, 0xb5, 0x2e,
	0x2a, 0xd7, 0xfa, 0x55, 0xb9, 0xd6, 0x97, 0xb9, 0xdb, 0xba, 0x98, 0xbb, 0xad, 0x9f, 0x73, 0xb7,
	0xf5, 0xf6, 0x31, 0xa1, 0x52, 0xed, 0x6a, 0xc2, 0x59, 0x50, 0x3f, 0xd5, 0xeb, 0xe2, 0xac, 0xb6,
	0x2c, 0xc6, 0xab, 0xfa, 0xbd, 0x3d, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x92, 0xbd, 0xe8,
	0xd5, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxTxSize != that1.MaxTxSize {
		return false
	}
	if this.MinGasPerByte != that1.MinGasPerByte {
		return false
	}
	return true
}
func (this *MsgGasParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams)
	if !ok {
		that2, ok := that.(MsgGasParams)
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
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	if that1.GasParams == nil {
		if this.GasParams != nil {
			return false
		}
	} else if this.GasParams == nil {
		return false
	} else if !this.GasParams.Equal(that1.GasParams) {
		return false
	}
	return true
}
func (this *MsgGasParams_FixedType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_FixedType)
	if !ok {
		that2, ok := that.(MsgGasParams_FixedType)
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
	if !this.FixedType.Equal(that1.FixedType) {
		return false
	}
	return true
}
func (this *MsgGasParams_GrantType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_GrantType)
	if !ok {
		that2, ok := that.(MsgGasParams_GrantType)
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
	if !this.GrantType.Equal(that1.GrantType) {
		return false
	}
	return true
}
func (this *MsgGasParams_MultiSendType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_MultiSendType)
	if !ok {
		that2, ok := that.(MsgGasParams_MultiSendType)
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
	if !this.MultiSendType.Equal(that1.MultiSendType) {
		return false
	}
	return true
}
func (this *MsgGasParams_GrantAllowanceType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_GrantAllowanceType)
	if !ok {
		that2, ok := that.(MsgGasParams_GrantAllowanceType)
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
	if !this.GrantAllowanceType.Equal(that1.GrantAllowanceType) {
		return false
	}
	return true
}
func (this *MsgGasParams_FixedGasParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_FixedGasParams)
	if !ok {
		that2, ok := that.(MsgGasParams_FixedGasParams)
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
	if this.FixedGas != that1.FixedGas {
		return false
	}
	return true
}
func (this *MsgGasParams_DynamicGasParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGasParams_DynamicGasParams)
	if !ok {
		that2, ok := that.(MsgGasParams_DynamicGasParams)
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
	if this.FixedGas != that1.FixedGas {
		return false
	}
	if this.GasPerItem != that1.GasPerItem {
		return false
	}
	return true
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
	if m.MinGasPerByte != 0 {
		i = encodeVarintGashub(dAtA, i, uint64(m.MinGasPerByte))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxTxSize != 0 {
		i = encodeVarintGashub(dAtA, i, uint64(m.MaxTxSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgGasParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGasParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasParams != nil {
		{
			size := m.GasParams.Size()
			i -= size
			if _, err := m.GasParams.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintGashub(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGasParams_FixedType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_FixedType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FixedType != nil {
		{
			size, err := m.FixedType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGashub(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *MsgGasParams_GrantType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_GrantType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GrantType != nil {
		{
			size, err := m.GrantType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGashub(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *MsgGasParams_MultiSendType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_MultiSendType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MultiSendType != nil {
		{
			size, err := m.MultiSendType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGashub(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *MsgGasParams_GrantAllowanceType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_GrantAllowanceType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GrantAllowanceType != nil {
		{
			size, err := m.GrantAllowanceType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGashub(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *MsgGasParams_FixedGasParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGasParams_FixedGasParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_FixedGasParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FixedGas != 0 {
		i = encodeVarintGashub(dAtA, i, uint64(m.FixedGas))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgGasParams_DynamicGasParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGasParams_DynamicGasParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGasParams_DynamicGasParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasPerItem != 0 {
		i = encodeVarintGashub(dAtA, i, uint64(m.GasPerItem))
		i--
		dAtA[i] = 0x10
	}
	if m.FixedGas != 0 {
		i = encodeVarintGashub(dAtA, i, uint64(m.FixedGas))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGashub(dAtA []byte, offset int, v uint64) int {
	offset -= sovGashub(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTxSize != 0 {
		n += 1 + sovGashub(uint64(m.MaxTxSize))
	}
	if m.MinGasPerByte != 0 {
		n += 1 + sovGashub(uint64(m.MinGasPerByte))
	}
	return n
}

func (m *MsgGasParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovGashub(uint64(l))
	}
	if m.GasParams != nil {
		n += m.GasParams.Size()
	}
	return n
}

func (m *MsgGasParams_FixedType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedType != nil {
		l = m.FixedType.Size()
		n += 1 + l + sovGashub(uint64(l))
	}
	return n
}
func (m *MsgGasParams_GrantType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GrantType != nil {
		l = m.GrantType.Size()
		n += 1 + l + sovGashub(uint64(l))
	}
	return n
}
func (m *MsgGasParams_MultiSendType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MultiSendType != nil {
		l = m.MultiSendType.Size()
		n += 1 + l + sovGashub(uint64(l))
	}
	return n
}
func (m *MsgGasParams_GrantAllowanceType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GrantAllowanceType != nil {
		l = m.GrantAllowanceType.Size()
		n += 1 + l + sovGashub(uint64(l))
	}
	return n
}
func (m *MsgGasParams_FixedGasParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedGas != 0 {
		n += 1 + sovGashub(uint64(m.FixedGas))
	}
	return n
}

func (m *MsgGasParams_DynamicGasParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedGas != 0 {
		n += 1 + sovGashub(uint64(m.FixedGas))
	}
	if m.GasPerItem != 0 {
		n += 1 + sovGashub(uint64(m.GasPerItem))
	}
	return n
}

func sovGashub(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGashub(x uint64) (n int) {
	return sovGashub(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGashub
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxSize", wireType)
			}
			m.MaxTxSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPerByte", wireType)
			}
			m.MinGasPerByte = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinGasPerByte |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGashub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGashub
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
func (m *MsgGasParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGashub
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
			return fmt.Errorf("proto: MsgGasParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGasParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
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
				return ErrInvalidLengthGashub
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGashub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
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
				return ErrInvalidLengthGashub
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGashub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgGasParams_FixedGasParams{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.GasParams = &MsgGasParams_FixedType{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
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
				return ErrInvalidLengthGashub
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGashub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgGasParams_DynamicGasParams{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.GasParams = &MsgGasParams_GrantType{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiSendType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
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
				return ErrInvalidLengthGashub
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGashub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgGasParams_DynamicGasParams{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.GasParams = &MsgGasParams_MultiSendType{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantAllowanceType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
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
				return ErrInvalidLengthGashub
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGashub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgGasParams_DynamicGasParams{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.GasParams = &MsgGasParams_GrantAllowanceType{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGashub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGashub
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
func (m *MsgGasParams_FixedGasParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGashub
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
			return fmt.Errorf("proto: FixedGasParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedGasParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedGas", wireType)
			}
			m.FixedGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FixedGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGashub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGashub
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
func (m *MsgGasParams_DynamicGasParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGashub
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
			return fmt.Errorf("proto: DynamicGasParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DynamicGasParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedGas", wireType)
			}
			m.FixedGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FixedGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPerItem", wireType)
			}
			m.GasPerItem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGashub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPerItem |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGashub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGashub
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
func skipGashub(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGashub
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
					return 0, ErrIntOverflowGashub
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
					return 0, ErrIntOverflowGashub
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
				return 0, ErrInvalidLengthGashub
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGashub
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGashub
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGashub        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGashub          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGashub = fmt.Errorf("proto: unexpected end of group")
)