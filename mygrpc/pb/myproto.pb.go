// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddReq struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReq) Reset()         { *m = AddReq{} }
func (m *AddReq) String() string { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()    {}
func (*AddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

func (m *AddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReq.Unmarshal(m, b)
}
func (m *AddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReq.Marshal(b, m, deterministic)
}
func (m *AddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReq.Merge(m, src)
}
func (m *AddReq) XXX_Size() int {
	return xxx_messageInfo_AddReq.Size(m)
}
func (m *AddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddReq proto.InternalMessageInfo

func (m *AddReq) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddReq) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type UserReq struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{1}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserReq) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type ReplyResp struct {
	RespStr              string   `protobuf:"bytes,1,opt,name=respStr,proto3" json:"respStr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyResp) Reset()         { *m = ReplyResp{} }
func (m *ReplyResp) String() string { return proto.CompactTextString(m) }
func (*ReplyResp) ProtoMessage()    {}
func (*ReplyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{2}
}

func (m *ReplyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyResp.Unmarshal(m, b)
}
func (m *ReplyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyResp.Marshal(b, m, deterministic)
}
func (m *ReplyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyResp.Merge(m, src)
}
func (m *ReplyResp) XXX_Size() int {
	return xxx_messageInfo_ReplyResp.Size(m)
}
func (m *ReplyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyResp proto.InternalMessageInfo

func (m *ReplyResp) GetRespStr() string {
	if m != nil {
		return m.RespStr
	}
	return ""
}

func init() {
	proto.RegisterType((*AddReq)(nil), "pb.AddReq")
	proto.RegisterType((*UserReq)(nil), "pb.UserReq")
	proto.RegisterType((*ReplyResp)(nil), "pb.ReplyResp")
}

func init() {
	proto.RegisterFile("myproto.proto", fileDescriptor_04877df457807402)
}

var fileDescriptor_04877df457807402 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0xaa, 0xc2, 0x30,
	0x10, 0x45, 0x5f, 0xfa, 0xb4, 0xa5, 0xa3, 0x05, 0x99, 0x55, 0x71, 0x25, 0x41, 0xa1, 0xb8, 0xa8,
	0xa0, 0x5f, 0xe0, 0xaa, 0xfb, 0x54, 0x3f, 0xa0, 0xa1, 0x83, 0x1b, 0xa5, 0x31, 0x13, 0x85, 0xfe,
	0xbd, 0x24, 0x6d, 0xdd, 0xb8, 0x09, 0xf7, 0x92, 0x33, 0xc3, 0x19, 0xc8, 0x1e, 0xbd, 0xb1, 0x9d,
	0xeb, 0xca, 0xf0, 0x62, 0x64, 0xb4, 0xdc, 0x42, 0x7c, 0x6e, 0x5b, 0x45, 0x4f, 0x5c, 0x82, 0x68,
	0x72, 0xb1, 0x11, 0xc5, 0x5c, 0x89, 0xc6, 0x37, 0x9d, 0x47, 0x43, 0xd3, 0xf2, 0x00, 0xc9, 0x95,
	0xc9, 0x7a, 0x0c, 0x61, 0xf6, 0x62, 0xb2, 0x81, 0x4c, 0x55, 0xc8, 0xb8, 0x82, 0xff, 0xe6, 0x46,
	0x23, 0xee, 0xa3, 0xdc, 0x41, 0xaa, 0xc8, 0xdc, 0x7b, 0x45, 0x6c, 0x30, 0x87, 0xc4, 0x12, 0x9b,
	0xda, 0x4d, 0x53, 0x53, 0x3d, 0x6a, 0x80, 0x0b, 0xb1, 0xab, 0xc9, 0xbe, 0xc9, 0x62, 0x01, 0x69,
	0x45, 0x2e, 0xe8, 0x30, 0x42, 0x69, 0x74, 0x39, 0xa8, 0xad, 0x33, 0x9f, 0xbf, 0xfb, 0xe4, 0x1f,
	0xee, 0x01, 0x2a, 0x72, 0x83, 0x12, 0xe3, 0xc2, 0x7f, 0x8f, 0x7e, 0x3f, 0xac, 0x8e, 0xc3, 0xb1,
	0xa7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xad, 0xcb, 0xb0, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServerClient is the client API for TestServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServerClient interface {
	GetAddRes(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*ReplyResp, error)
	GetUserRes(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ReplyResp, error)
}

type testServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServerClient(cc grpc.ClientConnInterface) TestServerClient {
	return &testServerClient{cc}
}

func (c *testServerClient) GetAddRes(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*ReplyResp, error) {
	out := new(ReplyResp)
	err := c.cc.Invoke(ctx, "/pb.TestServer/GetAddRes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServerClient) GetUserRes(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*ReplyResp, error) {
	out := new(ReplyResp)
	err := c.cc.Invoke(ctx, "/pb.TestServer/GetUserRes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServerServer is the server API for TestServer service.
type TestServerServer interface {
	GetAddRes(context.Context, *AddReq) (*ReplyResp, error)
	GetUserRes(context.Context, *UserReq) (*ReplyResp, error)
}

// UnimplementedTestServerServer can be embedded to have forward compatible implementations.
type UnimplementedTestServerServer struct {
}

func (*UnimplementedTestServerServer) GetAddRes(ctx context.Context, req *AddReq) (*ReplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddRes not implemented")
}
func (*UnimplementedTestServerServer) GetUserRes(ctx context.Context, req *UserReq) (*ReplyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRes not implemented")
}

func RegisterTestServerServer(s *grpc.Server, srv TestServerServer) {
	s.RegisterService(&_TestServer_serviceDesc, srv)
}

func _TestServer_GetAddRes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServerServer).GetAddRes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestServer/GetAddRes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServerServer).GetAddRes(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestServer_GetUserRes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServerServer).GetUserRes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestServer/GetUserRes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServerServer).GetUserRes(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TestServer",
	HandlerType: (*TestServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddRes",
			Handler:    _TestServer_GetAddRes_Handler,
		},
		{
			MethodName: "GetUserRes",
			Handler:    _TestServer_GetUserRes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}