package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckUsernameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUsernameRequest) Reset()         { *m = CheckUsernameRequest{} }
func (m *CheckUsernameRequest) String() string { return proto.CompactTextString(m) }
func (*CheckUsernameRequest) ProtoMessage()    {}
func (*CheckUsernameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e04ae7447678ce, []int{0}
}

func (m *CheckUsernameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUsernameRequest.Unmarshal(m, b)
}
func (m *CheckUsernameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUsernameRequest.Marshal(b, m, deterministic)
}
func (m *CheckUsernameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUsernameRequest.Merge(m, src)
}
func (m *CheckUsernameRequest) XXX_Size() int {
	return xxx_messageInfo_CheckUsernameRequest.Size(m)
}
func (m *CheckUsernameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUsernameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUsernameRequest proto.InternalMessageInfo

func (m *CheckUsernameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CheckUsernameResponse struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUsernameResponse) Reset()         { *m = CheckUsernameResponse{} }
func (m *CheckUsernameResponse) String() string { return proto.CompactTextString(m) }
func (*CheckUsernameResponse) ProtoMessage()    {}
func (*CheckUsernameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e04ae7447678ce, []int{1}
}

func (m *CheckUsernameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUsernameResponse.Unmarshal(m, b)
}
func (m *CheckUsernameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUsernameResponse.Marshal(b, m, deterministic)
}
func (m *CheckUsernameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUsernameResponse.Merge(m, src)
}
func (m *CheckUsernameResponse) XXX_Size() int {
	return xxx_messageInfo_CheckUsernameResponse.Size(m)
}
func (m *CheckUsernameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUsernameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUsernameResponse proto.InternalMessageInfo

func (m *CheckUsernameResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type GetUserDetailsRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserDetailsRequest) Reset()         { *m = GetUserDetailsRequest{} }
func (m *GetUserDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserDetailsRequest) ProtoMessage()    {}
func (*GetUserDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e04ae7447678ce, []int{2}
}

func (m *GetUserDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDetailsRequest.Unmarshal(m, b)
}
func (m *GetUserDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDetailsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDetailsRequest.Merge(m, src)
}
func (m *GetUserDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserDetailsRequest.Size(m)
}
func (m *GetUserDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDetailsRequest proto.InternalMessageInfo

func (m *GetUserDetailsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserDetailsResponse struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ImageUrl             string   `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	About                string   `protobuf:"bytes,4,opt,name=about,proto3" json:"about,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserDetailsResponse) Reset()         { *m = GetUserDetailsResponse{} }
func (m *GetUserDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserDetailsResponse) ProtoMessage()    {}
func (*GetUserDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e04ae7447678ce, []int{3}
}

func (m *GetUserDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDetailsResponse.Unmarshal(m, b)
}
func (m *GetUserDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDetailsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDetailsResponse.Merge(m, src)
}
func (m *GetUserDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserDetailsResponse.Size(m)
}
func (m *GetUserDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDetailsResponse proto.InternalMessageInfo

func (m *GetUserDetailsResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserDetailsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *GetUserDetailsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserDetailsResponse) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckUsernameRequest)(nil), "heurist_proto.CheckUsernameRequest")
	proto.RegisterType((*CheckUsernameResponse)(nil), "heurist_proto.CheckUsernameResponse")
	proto.RegisterType((*GetUserDetailsRequest)(nil), "heurist_proto.GetUserDetailsRequest")
	proto.RegisterType((*GetUserDetailsResponse)(nil), "heurist_proto.GetUserDetailsResponse")
}

func init() { proto.RegisterFile("heurist.proto", fileDescriptor_b5e04ae7447678ce) }

var fileDescriptor_b5e04ae7447678ce = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x48, 0x2d, 0x2d,
	0xca, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0x71, 0xe3, 0xc1, 0x5c, 0x25,
	0x23, 0x2e, 0x11, 0xe7, 0x8c, 0xd4, 0xe4, 0xec, 0xd0, 0xe2, 0xd4, 0xa2, 0xbc, 0xc4, 0xdc, 0xd4,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e, 0x52, 0xa8, 0x90, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf, 0x64, 0xc8, 0x25, 0x8a, 0xa6, 0xa7, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0xb3, 0x38, 0x2c, 0x31, 0x27, 0x33, 0x05, 0xac, 0x87,
	0x23, 0x08, 0xc6, 0x55, 0x32, 0xe6, 0x12, 0x75, 0x4f, 0x2d, 0x01, 0x69, 0x70, 0x49, 0x2d, 0x49,
	0xcc, 0xcc, 0x29, 0x26, 0xc6, 0x9e, 0x2a, 0x2e, 0x31, 0x74, 0x4d, 0x50, 0x8b, 0xf0, 0xe8, 0x02,
	0xc9, 0x65, 0xe6, 0x26, 0xa6, 0xa7, 0x86, 0x16, 0xe5, 0x48, 0x30, 0x41, 0xe4, 0x60, 0x7c, 0x21,
	0x21, 0x2e, 0x16, 0xb0, 0x1e, 0x66, 0xb0, 0x38, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x98, 0x94,
	0x5f, 0x5a, 0x22, 0xc1, 0x02, 0x16, 0x84, 0x70, 0x8c, 0xf6, 0x32, 0x72, 0x71, 0x7b, 0x40, 0x42,
	0xca, 0xbd, 0xa8, 0x20, 0x59, 0x28, 0x84, 0x8b, 0x15, 0xec, 0x67, 0x21, 0x65, 0x3d, 0x94, 0x00,
	0xd4, 0xc3, 0x16, 0x7a, 0x52, 0x2a, 0xf8, 0x15, 0x41, 0x7c, 0xa1, 0xc4, 0x20, 0x14, 0xc1, 0xc5,
	0x0e, 0xf5, 0xa1, 0x10, 0xba, 0x16, 0xac, 0xc1, 0x25, 0xa5, 0x4a, 0x40, 0x15, 0xcc, 0xe4, 0x24,
	0x36, 0xb0, 0xbc, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xe3, 0x60, 0x14, 0xfe, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeuristGrpcClient is the client API for HeuristGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeuristGrpcClient interface {
	Check(ctx context.Context, in *CheckUsernameRequest, opts ...grpc.CallOption) (*CheckUsernameResponse, error)
	GetUser(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
}

type heuristGrpcClient struct {
	cc *grpc.ClientConn
}

func NewHeuristGrpcClient(cc *grpc.ClientConn) HeuristGrpcClient {
	return &heuristGrpcClient{cc}
}

func (c *heuristGrpcClient) Check(ctx context.Context, in *CheckUsernameRequest, opts ...grpc.CallOption) (*CheckUsernameResponse, error) {
	out := new(CheckUsernameResponse)
	err := c.cc.Invoke(ctx, "/heurist_proto.HeuristGrpc/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heuristGrpcClient) GetUser(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/heurist_proto.HeuristGrpc/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeuristGrpcServer is the server API for HeuristGrpc service.
type HeuristGrpcServer interface {
	Check(context.Context, *CheckUsernameRequest) (*CheckUsernameResponse, error)
	GetUser(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
}

func RegisterHeuristGrpcServer(s *grpc.Server, srv HeuristGrpcServer) {
	s.RegisterService(&_HeuristGrpc_serviceDesc, srv)
}

func _HeuristGrpc_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeuristGrpcServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heurist_proto.HeuristGrpc/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeuristGrpcServer).Check(ctx, req.(*CheckUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeuristGrpc_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeuristGrpcServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heurist_proto.HeuristGrpc/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeuristGrpcServer).GetUser(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeuristGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heurist_proto.HeuristGrpc",
	HandlerType: (*HeuristGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HeuristGrpc_Check_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _HeuristGrpc_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heurist.proto",
}
