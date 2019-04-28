// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integral.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//protoc --go_out=plugins=grpc:. xx.proto
type FindUserByMobileRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByMobileRequest) Reset()         { *m = FindUserByMobileRequest{} }
func (m *FindUserByMobileRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByMobileRequest) ProtoMessage()    {}
func (*FindUserByMobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{0}
}

func (m *FindUserByMobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByMobileRequest.Unmarshal(m, b)
}
func (m *FindUserByMobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByMobileRequest.Marshal(b, m, deterministic)
}
func (m *FindUserByMobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByMobileRequest.Merge(m, src)
}
func (m *FindUserByMobileRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByMobileRequest.Size(m)
}
func (m *FindUserByMobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByMobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByMobileRequest proto.InternalMessageInfo

func (m *FindUserByMobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type IntegralResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Integral             int64    `protobuf:"varint,2,opt,name=integral,proto3" json:"integral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegralResponse) Reset()         { *m = IntegralResponse{} }
func (m *IntegralResponse) String() string { return proto.CompactTextString(m) }
func (*IntegralResponse) ProtoMessage()    {}
func (*IntegralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{1}
}

func (m *IntegralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegralResponse.Unmarshal(m, b)
}
func (m *IntegralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegralResponse.Marshal(b, m, deterministic)
}
func (m *IntegralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegralResponse.Merge(m, src)
}
func (m *IntegralResponse) XXX_Size() int {
	return xxx_messageInfo_IntegralResponse.Size(m)
}
func (m *IntegralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegralResponse proto.InternalMessageInfo

func (m *IntegralResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IntegralResponse) GetIntegral() int64 {
	if m != nil {
		return m.Integral
	}
	return 0
}

type AddIntegralRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Integral             int64    `protobuf:"varint,2,opt,name=integral,proto3" json:"integral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddIntegralRequest) Reset()         { *m = AddIntegralRequest{} }
func (m *AddIntegralRequest) String() string { return proto.CompactTextString(m) }
func (*AddIntegralRequest) ProtoMessage()    {}
func (*AddIntegralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{2}
}

func (m *AddIntegralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddIntegralRequest.Unmarshal(m, b)
}
func (m *AddIntegralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddIntegralRequest.Marshal(b, m, deterministic)
}
func (m *AddIntegralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddIntegralRequest.Merge(m, src)
}
func (m *AddIntegralRequest) XXX_Size() int {
	return xxx_messageInfo_AddIntegralRequest.Size(m)
}
func (m *AddIntegralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddIntegralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddIntegralRequest proto.InternalMessageInfo

func (m *AddIntegralRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AddIntegralRequest) GetIntegral() int64 {
	if m != nil {
		return m.Integral
	}
	return 0
}

type ConsumerIntegralRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ConsumerIntegral     int64    `protobuf:"varint,2,opt,name=consumerIntegral,proto3" json:"consumerIntegral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerIntegralRequest) Reset()         { *m = ConsumerIntegralRequest{} }
func (m *ConsumerIntegralRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumerIntegralRequest) ProtoMessage()    {}
func (*ConsumerIntegralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d836c8d70e1b53, []int{3}
}

func (m *ConsumerIntegralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerIntegralRequest.Unmarshal(m, b)
}
func (m *ConsumerIntegralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerIntegralRequest.Marshal(b, m, deterministic)
}
func (m *ConsumerIntegralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerIntegralRequest.Merge(m, src)
}
func (m *ConsumerIntegralRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumerIntegralRequest.Size(m)
}
func (m *ConsumerIntegralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerIntegralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerIntegralRequest proto.InternalMessageInfo

func (m *ConsumerIntegralRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ConsumerIntegralRequest) GetConsumerIntegral() int64 {
	if m != nil {
		return m.ConsumerIntegral
	}
	return 0
}

func init() {
	proto.RegisterType((*FindUserByMobileRequest)(nil), "protos.FindUserByMobileRequest")
	proto.RegisterType((*IntegralResponse)(nil), "protos.IntegralResponse")
	proto.RegisterType((*AddIntegralRequest)(nil), "protos.AddIntegralRequest")
	proto.RegisterType((*ConsumerIntegralRequest)(nil), "protos.ConsumerIntegralRequest")
}

func init() { proto.RegisterFile("integral.proto", fileDescriptor_d3d836c8d70e1b53) }

var fileDescriptor_d3d836c8d70e1b53 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x4d, 0x2f, 0x4a, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a,
	0x86, 0x5c, 0xe2, 0x6e, 0x99, 0x79, 0x29, 0xa1, 0xc5, 0xa9, 0x45, 0x4e, 0x95, 0xbe, 0xf9, 0x49,
	0x99, 0x39, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xb9, 0x60,
	0x01, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0xc9, 0x8d, 0x4b, 0xc0, 0x13, 0x6a,
	0x58, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x6d, 0x69, 0x71, 0x6a, 0x91, 0x67,
	0x0a, 0x58, 0x2d, 0x73, 0x10, 0x94, 0x27, 0x24, 0xc5, 0xc5, 0x01, 0xb3, 0x58, 0x82, 0x09, 0x2c,
	0x03, 0xe7, 0x2b, 0x79, 0x70, 0x09, 0x39, 0xa6, 0xa4, 0x20, 0x8c, 0x82, 0xdb, 0x4a, 0xb2, 0x49,
	0xb1, 0x5c, 0xe2, 0xce, 0xf9, 0x79, 0xc5, 0xa5, 0xb9, 0xa9, 0x45, 0xc4, 0x1a, 0xa7, 0xc5, 0x25,
	0x90, 0x8c, 0xa6, 0x05, 0x6a, 0x2c, 0x86, 0xb8, 0xd1, 0x5c, 0x46, 0x2e, 0x76, 0x50, 0x00, 0x05,
	0x15, 0x24, 0x0b, 0x39, 0x73, 0x71, 0x23, 0x39, 0x5a, 0x48, 0x0a, 0x12, 0x9c, 0xc5, 0x7a, 0x98,
	0x3e, 0x91, 0x92, 0x80, 0xc9, 0x61, 0x84, 0x96, 0x2f, 0x97, 0x00, 0xba, 0x7b, 0x85, 0xe4, 0x61,
	0xaa, 0x71, 0xf8, 0x04, 0xb7, 0x71, 0x49, 0x90, 0xb8, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x3b, 0x5e, 0x6b, 0xe4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserRpcClient is the client API for UserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRpcClient interface {
	AddIntegral(ctx context.Context, in *AddIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error)
	ConsumerIntegral(ctx context.Context, in *ConsumerIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error)
}

type userRpcClient struct {
	cc *grpc.ClientConn
}

func NewUserRpcClient(cc *grpc.ClientConn) UserRpcClient {
	return &userRpcClient{cc}
}

func (c *userRpcClient) AddIntegral(ctx context.Context, in *AddIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error) {
	out := new(IntegralResponse)
	err := c.cc.Invoke(ctx, "/protos.UserRpc/AddIntegral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) ConsumerIntegral(ctx context.Context, in *ConsumerIntegralRequest, opts ...grpc.CallOption) (*IntegralResponse, error) {
	out := new(IntegralResponse)
	err := c.cc.Invoke(ctx, "/protos.UserRpc/ConsumerIntegral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRpcServer is the server API for UserRpc service.
type UserRpcServer interface {
	AddIntegral(context.Context, *AddIntegralRequest) (*IntegralResponse, error)
	ConsumerIntegral(context.Context, *ConsumerIntegralRequest) (*IntegralResponse, error)
}

func RegisterUserRpcServer(s *grpc.Server, srv UserRpcServer) {
	s.RegisterService(&_UserRpc_serviceDesc, srv)
}

func _UserRpc_AddIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).AddIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRpc/AddIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).AddIntegral(ctx, req.(*AddIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_ConsumerIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumerIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).ConsumerIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserRpc/ConsumerIntegral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).ConsumerIntegral(ctx, req.(*ConsumerIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserRpc",
	HandlerType: (*UserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIntegral",
			Handler:    _UserRpc_AddIntegral_Handler,
		},
		{
			MethodName: "ConsumerIntegral",
			Handler:    _UserRpc_ConsumerIntegral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integral.proto",
}
