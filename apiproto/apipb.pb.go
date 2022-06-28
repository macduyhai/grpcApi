// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apiproto/apipb.proto

package apiproto

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

type Addrequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Fullname             string   `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Old                  int32    `protobuf:"varint,4,opt,name=old,proto3" json:"old,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addrequest) Reset()         { *m = Addrequest{} }
func (m *Addrequest) String() string { return proto.CompactTextString(m) }
func (*Addrequest) ProtoMessage()    {}
func (*Addrequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1769b40471972737, []int{0}
}

func (m *Addrequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addrequest.Unmarshal(m, b)
}
func (m *Addrequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addrequest.Marshal(b, m, deterministic)
}
func (m *Addrequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addrequest.Merge(m, src)
}
func (m *Addrequest) XXX_Size() int {
	return xxx_messageInfo_Addrequest.Size(m)
}
func (m *Addrequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Addrequest.DiscardUnknown(m)
}

var xxx_messageInfo_Addrequest proto.InternalMessageInfo

func (m *Addrequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Addrequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Addrequest) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *Addrequest) GetOld() int32 {
	if m != nil {
		return m.Old
	}
	return 0
}

type AddResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1769b40471972737, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AddResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Addrequest)(nil), "api.Addrequest")
	proto.RegisterType((*AddResponse)(nil), "api.AddResponse")
}

func init() { proto.RegisterFile("apiproto/apipb.proto", fileDescriptor_1769b40471972737) }

var fileDescriptor_1769b40471972737 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc4, 0x30,
	0x0c, 0x86, 0x55, 0x72, 0x07, 0x87, 0x19, 0x38, 0x59, 0x0c, 0xd1, 0x4d, 0xa7, 0x4e, 0x9d, 0x8a,
	0x04, 0x12, 0x0b, 0x53, 0x79, 0x84, 0xb0, 0xb1, 0xa5, 0xb5, 0x41, 0x91, 0xda, 0x26, 0xc4, 0x2d,
	0xbc, 0x3e, 0x4a, 0x4a, 0x60, 0xca, 0xff, 0xe5, 0x53, 0xf4, 0xc7, 0x86, 0x3b, 0x1b, 0x5c, 0x88,
	0x7e, 0xf1, 0xf7, 0x29, 0xf4, 0x6d, 0xce, 0xa8, 0x6c, 0x70, 0x75, 0x04, 0xe8, 0x88, 0x22, 0x7f,
	0xae, 0x2c, 0x0b, 0x9e, 0xe0, 0xb0, 0x0a, 0xc7, 0xd9, 0x4e, 0xac, 0xab, 0x73, 0xd5, 0x5c, 0x9b,
	0x3f, 0x4e, 0x2e, 0x58, 0x91, 0x6f, 0x1f, 0x49, 0x5f, 0x6c, 0xae, 0x70, 0x72, 0xef, 0xeb, 0x38,
	0xe6, 0x77, 0x6a, 0x73, 0x85, 0xf1, 0x08, 0xca, 0x8f, 0xa4, 0x77, 0xe7, 0xaa, 0xd9, 0x9b, 0x14,
	0xeb, 0x67, 0xb8, 0xe9, 0x88, 0x0c, 0x4b, 0xf0, 0xb3, 0x30, 0x22, 0xec, 0x06, 0x4f, 0x5b, 0xe1,
	0xde, 0xe4, 0x8c, 0x1a, 0xae, 0x26, 0x16, 0xb1, 0x1f, 0xfc, 0xdb, 0x55, 0xf0, 0xe1, 0x29, 0x7f,
	0xf8, 0x95, 0xe3, 0x97, 0x1b, 0x18, 0x1b, 0x50, 0x1d, 0x11, 0xde, 0xb6, 0x36, 0xb8, 0xf6, 0x7f,
	0x90, 0xd3, 0xb1, 0x5c, 0x94, 0x96, 0x17, 0x78, 0x3b, 0x94, 0x2d, 0xf4, 0x97, 0xf9, 0x78, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x15, 0x2e, 0x0e, 0xc2, 0x18, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Add(ctx context.Context, in *Addrequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *Addrequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/api.AddService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Add(context.Context, *Addrequest) (*AddResponse, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) Add(ctx context.Context, req *Addrequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Addrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*Addrequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiproto/apipb.proto",
}