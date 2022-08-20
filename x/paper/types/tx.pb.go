// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paper/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreatePaper struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Host      string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	PaperName string `protobuf:"bytes,3,opt,name=paperName,proto3" json:"paperName,omitempty"`
	Owner     string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Price     string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MsgCreatePaper) Reset()         { *m = MsgCreatePaper{} }
func (m *MsgCreatePaper) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePaper) ProtoMessage()    {}
func (*MsgCreatePaper) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e0eb91fb56ab61, []int{0}
}
func (m *MsgCreatePaper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePaper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePaper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePaper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePaper.Merge(m, src)
}
func (m *MsgCreatePaper) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePaper) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePaper.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePaper proto.InternalMessageInfo

func (m *MsgCreatePaper) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePaper) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MsgCreatePaper) GetPaperName() string {
	if m != nil {
		return m.PaperName
	}
	return ""
}

func (m *MsgCreatePaper) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgCreatePaper) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type MsgCreatePaperResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreatePaperResponse) Reset()         { *m = MsgCreatePaperResponse{} }
func (m *MsgCreatePaperResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePaperResponse) ProtoMessage()    {}
func (*MsgCreatePaperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e0eb91fb56ab61, []int{1}
}
func (m *MsgCreatePaperResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePaperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePaperResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePaperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePaperResponse.Merge(m, src)
}
func (m *MsgCreatePaperResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePaperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePaperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePaperResponse proto.InternalMessageInfo

func (m *MsgCreatePaperResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdatePaper struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NewOwner string `protobuf:"bytes,3,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
	NewPrice string `protobuf:"bytes,4,opt,name=newPrice,proto3" json:"newPrice,omitempty"`
}

func (m *MsgUpdatePaper) Reset()         { *m = MsgUpdatePaper{} }
func (m *MsgUpdatePaper) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePaper) ProtoMessage()    {}
func (*MsgUpdatePaper) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e0eb91fb56ab61, []int{2}
}
func (m *MsgUpdatePaper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePaper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePaper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePaper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePaper.Merge(m, src)
}
func (m *MsgUpdatePaper) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePaper) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePaper.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePaper proto.InternalMessageInfo

func (m *MsgUpdatePaper) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdatePaper) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgUpdatePaper) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

func (m *MsgUpdatePaper) GetNewPrice() string {
	if m != nil {
		return m.NewPrice
	}
	return ""
}

type MsgUpdatePaperResponse struct {
}

func (m *MsgUpdatePaperResponse) Reset()         { *m = MsgUpdatePaperResponse{} }
func (m *MsgUpdatePaperResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePaperResponse) ProtoMessage()    {}
func (*MsgUpdatePaperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e0eb91fb56ab61, []int{3}
}
func (m *MsgUpdatePaperResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePaperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePaperResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePaperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePaperResponse.Merge(m, src)
}
func (m *MsgUpdatePaperResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePaperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePaperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePaperResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePaper)(nil), "paper.paper.MsgCreatePaper")
	proto.RegisterType((*MsgCreatePaperResponse)(nil), "paper.paper.MsgCreatePaperResponse")
	proto.RegisterType((*MsgUpdatePaper)(nil), "paper.paper.MsgUpdatePaper")
	proto.RegisterType((*MsgUpdatePaperResponse)(nil), "paper.paper.MsgUpdatePaperResponse")
}

func init() { proto.RegisterFile("paper/tx.proto", fileDescriptor_87e0eb91fb56ab61) }

var fileDescriptor_87e0eb91fb56ab61 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0x34, 0x05, 0x7a, 0x91, 0x32, 0x18, 0x84, 0xac, 0x82, 0x2c, 0x14, 0x96, 0x2e,
	0xa4, 0x12, 0xbc, 0x01, 0xcc, 0x25, 0x55, 0x24, 0x16, 0xb6, 0xd0, 0x58, 0x25, 0x03, 0xb1, 0x65,
	0x47, 0x6a, 0x79, 0x03, 0x46, 0xde, 0x82, 0x57, 0x61, 0xec, 0xc8, 0x88, 0x92, 0x17, 0x41, 0x39,
	0xb7, 0x21, 0x55, 0x11, 0xea, 0x62, 0xf9, 0xff, 0x7f, 0xf9, 0xfc, 0xdd, 0xe9, 0x20, 0x50, 0xa9,
	0x12, 0x7a, 0x5c, 0x2e, 0x23, 0xa5, 0x65, 0x29, 0xa9, 0x8f, 0x3a, 0xc2, 0x33, 0x7c, 0x23, 0x10,
	0x4c, 0xcc, 0xfc, 0x4e, 0x8b, 0xb4, 0x14, 0xd3, 0xc6, 0xa2, 0x0c, 0x0e, 0x67, 0x8d, 0x94, 0x9a,
	0x91, 0x0b, 0x32, 0x1a, 0x24, 0x1b, 0x49, 0x29, 0x78, 0xcf, 0xd2, 0x94, 0xcc, 0x45, 0x1b, 0xef,
	0xf4, 0x1c, 0x06, 0x58, 0xe9, 0x3e, 0x7d, 0x11, 0xac, 0x87, 0xc1, 0xaf, 0x41, 0x4f, 0xa0, 0x2f,
	0x17, 0x85, 0xd0, 0xcc, 0xc3, 0xc4, 0x8a, 0xc6, 0x55, 0x3a, 0x9f, 0x09, 0xd6, 0xb7, 0x2e, 0x8a,
	0x70, 0x04, 0xa7, 0xdb, 0x24, 0x89, 0x30, 0x4a, 0x16, 0x46, 0xd0, 0x00, 0xdc, 0x3c, 0x43, 0x18,
	0x2f, 0x71, 0xf3, 0x2c, 0xd4, 0xc8, 0xfc, 0xa0, 0xb2, 0x3d, 0x98, 0xed, 0x5b, 0x4b, 0xec, 0xe6,
	0x19, 0x1d, 0xc2, 0x51, 0x21, 0x16, 0x31, 0x42, 0x59, 0xdc, 0x56, 0xaf, 0xb3, 0x29, 0xa2, 0x79,
	0x6d, 0x86, 0x3a, 0x64, 0x48, 0xd7, 0xf9, 0x73, 0x43, 0x77, 0xfd, 0x41, 0xa0, 0x37, 0x31, 0x73,
	0x1a, 0x83, 0xdf, 0x1d, 0xe3, 0x59, 0xd4, 0x99, 0x73, 0xb4, 0xdd, 0xd9, 0xf0, 0xf2, 0x9f, 0xb0,
	0x6d, 0x3b, 0x06, 0xbf, 0xdb, 0xe3, 0x4e, 0xc1, 0x4e, 0xb8, 0x5b, 0xf0, 0x0f, 0xd2, 0xdb, 0xab,
	0xcf, 0x8a, 0x93, 0x55, 0xc5, 0xc9, 0x77, 0xc5, 0xc9, 0x7b, 0xcd, 0x9d, 0x55, 0xcd, 0x9d, 0xaf,
	0x9a, 0x3b, 0x8f, 0xc7, 0x76, 0x47, 0x96, 0xe3, 0xf5, 0xae, 0xbc, 0x2a, 0x61, 0x9e, 0x0e, 0x70,
	0x5f, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x9f, 0x0f, 0x52, 0x41, 0x02, 0x00, 0x00,
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
	CreatePaper(ctx context.Context, in *MsgCreatePaper, opts ...grpc.CallOption) (*MsgCreatePaperResponse, error)
	UpdatePaper(ctx context.Context, in *MsgUpdatePaper, opts ...grpc.CallOption) (*MsgUpdatePaperResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePaper(ctx context.Context, in *MsgCreatePaper, opts ...grpc.CallOption) (*MsgCreatePaperResponse, error) {
	out := new(MsgCreatePaperResponse)
	err := c.cc.Invoke(ctx, "/paper.paper.Msg/CreatePaper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePaper(ctx context.Context, in *MsgUpdatePaper, opts ...grpc.CallOption) (*MsgUpdatePaperResponse, error) {
	out := new(MsgUpdatePaperResponse)
	err := c.cc.Invoke(ctx, "/paper.paper.Msg/UpdatePaper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePaper(context.Context, *MsgCreatePaper) (*MsgCreatePaperResponse, error)
	UpdatePaper(context.Context, *MsgUpdatePaper) (*MsgUpdatePaperResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePaper(ctx context.Context, req *MsgCreatePaper) (*MsgCreatePaperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaper not implemented")
}
func (*UnimplementedMsgServer) UpdatePaper(ctx context.Context, req *MsgUpdatePaper) (*MsgUpdatePaperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaper not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePaper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePaper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paper.paper.Msg/CreatePaper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePaper(ctx, req.(*MsgCreatePaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePaper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePaper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paper.paper.Msg/UpdatePaper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePaper(ctx, req.(*MsgUpdatePaper))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paper.paper.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePaper",
			Handler:    _Msg_CreatePaper_Handler,
		},
		{
			MethodName: "UpdatePaper",
			Handler:    _Msg_UpdatePaper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paper/tx.proto",
}

func (m *MsgCreatePaper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePaper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePaper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PaperName) > 0 {
		i -= len(m.PaperName)
		copy(dAtA[i:], m.PaperName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PaperName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePaperResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePaperResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePaperResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePaper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePaper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePaper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewPrice) > 0 {
		i -= len(m.NewPrice)
		copy(dAtA[i:], m.NewPrice)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewPrice)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePaperResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePaperResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePaperResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePaper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PaperName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreatePaperResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdatePaper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewPrice)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdatePaperResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePaper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePaper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePaper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaperName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaperName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePaperResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePaperResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePaperResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePaper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePaper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePaper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePaperResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePaperResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePaperResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
