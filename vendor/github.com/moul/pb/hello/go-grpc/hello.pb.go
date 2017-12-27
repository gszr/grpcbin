// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package hello

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Greeting         *string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetGreeting() string {
	if m != nil && m.Greeting != nil {
		return *m.Greeting
	}
	return ""
}

type HelloResponse struct {
	Reply            *string `protobuf:"bytes,1,req,name=reply" json:"reply,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetReply() string {
	if m != nil && m.Reply != nil {
		return *m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_LotsOfRepliesClient, error)
	LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (HelloService_LotsOfGreetingsClient, error)
	BidiHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_BidiHelloClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/hello.HelloService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_LotsOfRepliesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HelloService_serviceDesc.Streams[0], c.cc, "/hello.HelloService/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_LotsOfRepliesClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *helloServiceLotsOfRepliesClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) LotsOfGreetings(ctx context.Context, opts ...grpc.CallOption) (HelloService_LotsOfGreetingsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HelloService_serviceDesc.Streams[1], c.cc, "/hello.HelloService/LotsOfGreetings", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceLotsOfGreetingsClient{stream}
	return x, nil
}

type HelloService_LotsOfGreetingsClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceLotsOfGreetingsClient struct {
	grpc.ClientStream
}

func (x *helloServiceLotsOfGreetingsClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceLotsOfGreetingsClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) BidiHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_BidiHelloClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HelloService_serviceDesc.Streams[2], c.cc, "/hello.HelloService/BidiHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceBidiHelloClient{stream}
	return x, nil
}

type HelloService_BidiHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceBidiHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceBidiHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceBidiHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	LotsOfReplies(*HelloRequest, HelloService_LotsOfRepliesServer) error
	LotsOfGreetings(HelloService_LotsOfGreetingsServer) error
	BidiHello(HelloService_BidiHelloServer) error
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).LotsOfReplies(m, &helloServiceLotsOfRepliesServer{stream})
}

type HelloService_LotsOfRepliesServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *helloServiceLotsOfRepliesServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_LotsOfGreetings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).LotsOfGreetings(&helloServiceLotsOfGreetingsServer{stream})
}

type HelloService_LotsOfGreetingsServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceLotsOfGreetingsServer struct {
	grpc.ServerStream
}

func (x *helloServiceLotsOfGreetingsServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceLotsOfGreetingsServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_BidiHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).BidiHello(&helloServiceBidiHelloServer{stream})
}

type HelloService_BidiHelloServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceBidiHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceBidiHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceBidiHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _HelloService_LotsOfReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LotsOfGreetings",
			Handler:       _HelloService_LotsOfGreetings_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidiHello",
			Handler:       _HelloService_BidiHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xb4, 0xb8, 0x78, 0x3c,
	0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e, 0xf4, 0xa2, 0xd4,
	0xd4, 0x92, 0xcc, 0xbc, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x49, 0x95,
	0x8b, 0x17, 0xaa, 0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0xb5, 0x28, 0xb5,
	0x20, 0xa7, 0x52, 0x82, 0x51, 0x81, 0x49, 0x83, 0x33, 0x08, 0xc2, 0x31, 0xfa, 0xcf, 0x08, 0x35,
	0x33, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x94, 0x8b, 0x23, 0x38, 0xb1, 0x12, 0x2c,
	0x24, 0x24, 0xac, 0x07, 0x71, 0x04, 0xb2, 0xa5, 0x52, 0x22, 0xa8, 0x82, 0x50, 0xd3, 0x6d, 0xb8,
	0x78, 0x7d, 0xf2, 0x4b, 0x8a, 0xfd, 0xd3, 0x82, 0x52, 0x0b, 0x72, 0x32, 0x53, 0x8b, 0x49, 0xd0,
	0x6b, 0xc0, 0x28, 0x64, 0xc7, 0xc5, 0x0f, 0xd1, 0xed, 0x0e, 0x75, 0x3e, 0x29, 0xfa, 0x35, 0x18,
	0x85, 0xac, 0xb8, 0x38, 0x9d, 0x32, 0x53, 0x32, 0x49, 0x75, 0xb5, 0x06, 0xa3, 0x01, 0x23, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x22, 0xff, 0xdc, 0xc9, 0x69, 0x01, 0x00, 0x00,
}
