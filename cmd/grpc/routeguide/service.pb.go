// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package routeguide is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Post
	Result
*/
package routeguide

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

type Post struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Body string `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Result struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Post)(nil), "routeguide.Post")
	proto.RegisterType((*Result)(nil), "routeguide.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteGuide service

type RouteGuideClient interface {
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Result, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/routeguide.RouteGuide/CreatePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	CreatePost(context.Context, *Post) (*Result, error)
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _RouteGuide_CreatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x3f, 0x0b, 0xc2, 0x40,
	0x0c, 0x47, 0x6d, 0xad, 0x05, 0x03, 0x8a, 0x64, 0x2a, 0xe2, 0x20, 0x9d, 0x9c, 0x6e, 0x50, 0x67,
	0x07, 0x1d, 0x5c, 0xe5, 0xbe, 0x81, 0xed, 0x45, 0x39, 0x50, 0x22, 0xf7, 0x47, 0xf0, 0xdb, 0x4b,
	0xd2, 0x41, 0xb7, 0xc7, 0x0b, 0xe4, 0xfd, 0x60, 0x16, 0x29, 0xbc, 0x7d, 0x4f, 0xe6, 0x15, 0x38,
	0x31, 0x42, 0xe0, 0x9c, 0xe8, 0x9e, 0xbd, 0xa3, 0xf6, 0x00, 0xd5, 0x85, 0x63, 0xc2, 0x39, 0x94,
	0xde, 0x35, 0xc5, 0xba, 0xd8, 0x4c, 0x6d, 0xe9, 0x1d, 0x22, 0x54, 0xb7, 0xc0, 0xcf, 0xa6, 0x54,
	0xa3, 0x2c, 0xae, 0x63, 0xf7, 0x69, 0xc6, 0x83, 0x13, 0x6e, 0x57, 0x50, 0x5b, 0x8a, 0xf9, 0x91,
	0xe4, 0xda, 0xb3, 0x23, 0xfd, 0x31, 0xb1, 0xca, 0xdb, 0x23, 0x80, 0x95, 0xd6, 0x59, 0x5a, 0xb8,
	0x07, 0x38, 0x05, 0xba, 0x26, 0xd2, 0xe2, 0xc2, 0xfc, 0x66, 0x18, 0x31, 0x4b, 0xfc, 0x37, 0xc3,
	0xd7, 0x76, 0xd4, 0xd5, 0x3a, 0x7a, 0xf7, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x49, 0x3f, 0x13,
	0xc5, 0x00, 0x00, 0x00,
}
