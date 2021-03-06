// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package Test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 请求用户信息
type UserInfoRequest struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *UserInfoRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 请求用户信息的结果
type UserInfoResponse struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age   uint32 `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Sex   uint32 `protobuf:"varint,3,opt,name=sex" json:"sex,omitempty"`
	Count uint32 `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *UserInfoResponse) Reset()                    { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()               {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoResponse) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfoResponse) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserInfoResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "Test.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "Test.UserInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Data service

type DataClient interface {
	// 简单Rpc
	// 获取用户数据
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	//  修改用户 双向流模式
	ChangeUserInfo(ctx context.Context, opts ...grpc.CallOption) (Data_ChangeUserInfoClient, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := grpc.Invoke(ctx, "/Test.Data/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) ChangeUserInfo(ctx context.Context, opts ...grpc.CallOption) (Data_ChangeUserInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Data_serviceDesc.Streams[0], c.cc, "/Test.Data/ChangeUserInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataChangeUserInfoClient{stream}
	return x, nil
}

type Data_ChangeUserInfoClient interface {
	Send(*UserInfoResponse) error
	Recv() (*UserInfoResponse, error)
	grpc.ClientStream
}

type dataChangeUserInfoClient struct {
	grpc.ClientStream
}

func (x *dataChangeUserInfoClient) Send(m *UserInfoResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataChangeUserInfoClient) Recv() (*UserInfoResponse, error) {
	m := new(UserInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Data service

type DataServer interface {
	// 简单Rpc
	// 获取用户数据
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	//  修改用户 双向流模式
	ChangeUserInfo(Data_ChangeUserInfoServer) error
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test.Data/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_ChangeUserInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).ChangeUserInfo(&dataChangeUserInfoServer{stream})
}

type Data_ChangeUserInfoServer interface {
	Send(*UserInfoResponse) error
	Recv() (*UserInfoResponse, error)
	grpc.ServerStream
}

type dataChangeUserInfoServer struct {
	grpc.ServerStream
}

func (x *dataChangeUserInfoServer) Send(m *UserInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataChangeUserInfoServer) Recv() (*UserInfoResponse, error) {
	m := new(UserInfoResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Test.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Data_GetUserInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChangeUserInfo",
			Handler:       _Data_ChangeUserInfo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x09, 0x49, 0x2d, 0x2e, 0x51, 0x52, 0xe6, 0xe2,
	0x0f, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x95,
	0x12, 0xb8, 0x04, 0x10, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0xf2,
	0x12, 0x73, 0x53, 0xc1, 0xca, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0xce, 0xc4, 0xf4, 0x54, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0x10, 0x13, 0x24, 0x52, 0x9c, 0x5a, 0x21, 0xc1, 0x0c, 0x11, 0x29,
	0x4e, 0xad, 0x10, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x01, 0x8b, 0x41,
	0x38, 0x46, 0x7d, 0x8c, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0x89, 0x42, 0x76, 0x5c, 0xdc, 0xee, 0xa9,
	0x25, 0x30, 0xdb, 0x84, 0x44, 0xf5, 0x40, 0xae, 0xd4, 0x43, 0x73, 0xa2, 0x94, 0x18, 0xba, 0x30,
	0xc4, 0x51, 0x4a, 0x0c, 0x42, 0x6e, 0x5c, 0x7c, 0xce, 0x19, 0x89, 0x79, 0xe9, 0xa9, 0x70, 0x23,
	0x70, 0xa8, 0xc5, 0x6d, 0x86, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0x38, 0x90, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x6a, 0xa9, 0xeb, 0x32, 0x01, 0x00, 0x00,
}
