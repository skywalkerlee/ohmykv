// Code generated by protoc-gen-go.
// source: msg.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	Readreq
	Readresp
	Iterreq
	Iterresp
	Manreq
	Manresp
	Writereq
	Writeresp
	Dump
*/
package msg

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

type Readreq struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *Readreq) Reset()                    { *m = Readreq{} }
func (m *Readreq) String() string            { return proto.CompactTextString(m) }
func (*Readreq) ProtoMessage()               {}
func (*Readreq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Readreq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Readresp struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Readresp) Reset()                    { *m = Readresp{} }
func (m *Readresp) String() string            { return proto.CompactTextString(m) }
func (*Readresp) ProtoMessage()               {}
func (*Readresp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Readresp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Iterreq struct {
	Op   int32  `protobuf:"varint,1,opt,name=op" json:"op,omitempty"`
	Seek []byte `protobuf:"bytes,2,opt,name=seek,proto3" json:"seek,omitempty"`
}

func (m *Iterreq) Reset()                    { *m = Iterreq{} }
func (m *Iterreq) String() string            { return proto.CompactTextString(m) }
func (*Iterreq) ProtoMessage()               {}
func (*Iterreq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Iterreq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *Iterreq) GetSeek() []byte {
	if m != nil {
		return m.Seek
	}
	return nil
}

type Iterresp struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Iterresp) Reset()                    { *m = Iterresp{} }
func (m *Iterresp) String() string            { return proto.CompactTextString(m) }
func (*Iterresp) ProtoMessage()               {}
func (*Iterresp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Iterresp) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Iterresp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Manreq struct {
	Op    int32  `protobuf:"varint,1,opt,name=op" json:"op,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Manreq) Reset()                    { *m = Manreq{} }
func (m *Manreq) String() string            { return proto.CompactTextString(m) }
func (*Manreq) ProtoMessage()               {}
func (*Manreq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Manreq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *Manreq) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Manresp struct {
	Status int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Manresp) Reset()                    { *m = Manresp{} }
func (m *Manresp) String() string            { return proto.CompactTextString(m) }
func (*Manresp) ProtoMessage()               {}
func (*Manresp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Manresp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Manresp) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Writereq struct {
	Op    int32  `protobuf:"varint,1,opt,name=op" json:"op,omitempty"`
	Key   []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Writereq) Reset()                    { *m = Writereq{} }
func (m *Writereq) String() string            { return proto.CompactTextString(m) }
func (*Writereq) ProtoMessage()               {}
func (*Writereq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Writereq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *Writereq) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Writereq) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Writeresp struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Writeresp) Reset()                    { *m = Writeresp{} }
func (m *Writeresp) String() string            { return proto.CompactTextString(m) }
func (*Writeresp) ProtoMessage()               {}
func (*Writeresp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Writeresp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Dump struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Dump) Reset()                    { *m = Dump{} }
func (m *Dump) String() string            { return proto.CompactTextString(m) }
func (*Dump) ProtoMessage()               {}
func (*Dump) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Dump) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Dump) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Readreq)(nil), "msg.readreq")
	proto.RegisterType((*Readresp)(nil), "msg.readresp")
	proto.RegisterType((*Iterreq)(nil), "msg.iterreq")
	proto.RegisterType((*Iterresp)(nil), "msg.iterresp")
	proto.RegisterType((*Manreq)(nil), "msg.manreq")
	proto.RegisterType((*Manresp)(nil), "msg.manresp")
	proto.RegisterType((*Writereq)(nil), "msg.writereq")
	proto.RegisterType((*Writeresp)(nil), "msg.writeresp")
	proto.RegisterType((*Dump)(nil), "msg.dump")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Kv service

type KvClient interface {
	Write(ctx context.Context, in *Writereq, opts ...grpc.CallOption) (*Writeresp, error)
	Read(ctx context.Context, in *Readreq, opts ...grpc.CallOption) (*Readresp, error)
	Man(ctx context.Context, in *Manreq, opts ...grpc.CallOption) (*Manresp, error)
	Iterator(ctx context.Context, in *Iterreq, opts ...grpc.CallOption) (Kv_IteratorClient, error)
}

type kvClient struct {
	cc *grpc.ClientConn
}

func NewKvClient(cc *grpc.ClientConn) KvClient {
	return &kvClient{cc}
}

func (c *kvClient) Write(ctx context.Context, in *Writereq, opts ...grpc.CallOption) (*Writeresp, error) {
	out := new(Writeresp)
	err := grpc.Invoke(ctx, "/msg.kv/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvClient) Read(ctx context.Context, in *Readreq, opts ...grpc.CallOption) (*Readresp, error) {
	out := new(Readresp)
	err := grpc.Invoke(ctx, "/msg.kv/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvClient) Man(ctx context.Context, in *Manreq, opts ...grpc.CallOption) (*Manresp, error) {
	out := new(Manresp)
	err := grpc.Invoke(ctx, "/msg.kv/Man", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvClient) Iterator(ctx context.Context, in *Iterreq, opts ...grpc.CallOption) (Kv_IteratorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Kv_serviceDesc.Streams[0], c.cc, "/msg.kv/Iterator", opts...)
	if err != nil {
		return nil, err
	}
	x := &kvIteratorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kv_IteratorClient interface {
	Recv() (*Iterresp, error)
	grpc.ClientStream
}

type kvIteratorClient struct {
	grpc.ClientStream
}

func (x *kvIteratorClient) Recv() (*Iterresp, error) {
	m := new(Iterresp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Kv service

type KvServer interface {
	Write(context.Context, *Writereq) (*Writeresp, error)
	Read(context.Context, *Readreq) (*Readresp, error)
	Man(context.Context, *Manreq) (*Manresp, error)
	Iterator(*Iterreq, Kv_IteratorServer) error
}

func RegisterKvServer(s *grpc.Server, srv KvServer) {
	s.RegisterService(&_Kv_serviceDesc, srv)
}

func _Kv_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Writereq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.kv/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Write(ctx, req.(*Writereq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kv_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Readreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.kv/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Read(ctx, req.(*Readreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kv_Man_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Man(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.kv/Man",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Man(ctx, req.(*Manreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kv_Iterator_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Iterreq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KvServer).Iterator(m, &kvIteratorServer{stream})
}

type Kv_IteratorServer interface {
	Send(*Iterresp) error
	grpc.ServerStream
}

type kvIteratorServer struct {
	grpc.ServerStream
}

func (x *kvIteratorServer) Send(m *Iterresp) error {
	return x.ServerStream.SendMsg(m)
}

var _Kv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.kv",
	HandlerType: (*KvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Kv_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Kv_Read_Handler,
		},
		{
			MethodName: "Man",
			Handler:    _Kv_Man_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Iterator",
			Handler:       _Kv_Iterator_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "msg.proto",
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x6c, 0x36, 0x7f, 0xfb, 0x7e, 0x69, 0xf9, 0xf1, 0x10, 0x29, 0xf1, 0x52, 0x56, 0x84, 0x82,
	0x18, 0xa4, 0xe2, 0x17, 0xf0, 0xe6, 0xc1, 0x4b, 0x2e, 0x9e, 0xb7, 0x64, 0x29, 0x12, 0xd3, 0x5d,
	0x77, 0x37, 0x95, 0x7e, 0x21, 0x3f, 0xa7, 0xec, 0x66, 0x43, 0x03, 0x46, 0xf0, 0xf6, 0x26, 0x6f,
	0xde, 0xcc, 0x24, 0x13, 0x98, 0xb7, 0x7a, 0x5f, 0x4a, 0x25, 0x8c, 0xc0, 0xb0, 0xd5, 0x7b, 0x7a,
	0x05, 0xa9, 0xe2, 0xac, 0x56, 0xfc, 0x03, 0xff, 0x43, 0xd8, 0xf0, 0xd3, 0x2a, 0x58, 0x07, 0x9b,
	0xbc, 0xb2, 0x23, 0x5d, 0x43, 0xd6, 0x2f, 0xb5, 0xc4, 0x0b, 0x88, 0x8f, 0xec, 0xbd, 0xe3, 0x7e,
	0xdf, 0x03, 0x7a, 0x07, 0xe9, 0x9b, 0xe1, 0xca, 0x9e, 0x2f, 0x81, 0x08, 0xe9, 0xb6, 0x71, 0x45,
	0x84, 0x44, 0x84, 0x48, 0x73, 0xde, 0xac, 0x88, 0xe3, 0xbb, 0x99, 0x6e, 0x21, 0xeb, 0xe9, 0x5a,
	0xfe, 0xb4, 0x3b, 0x5b, 0x90, 0xb1, 0x45, 0x09, 0x49, 0xcb, 0x0e, 0x53, 0x0e, 0xd3, 0xfc, 0x47,
	0x48, 0x1d, 0x5f, 0x4b, 0xbc, 0x84, 0x44, 0x1b, 0x66, 0x3a, 0xed, 0x8f, 0x3c, 0xb2, 0xd1, 0x76,
	0xa2, 0x3e, 0x0d, 0xd1, 0xec, 0x4c, 0x9f, 0x20, 0xfb, 0x54, 0x36, 0xdc, 0x84, 0x91, 0x8f, 0x4a,
	0x26, 0xa2, 0x86, 0x63, 0xeb, 0x6b, 0x98, 0x7b, 0x8d, 0xdf, 0xcd, 0x69, 0x09, 0x51, 0xdd, 0xb5,
	0x7f, 0x7e, 0xff, 0xed, 0x57, 0x00, 0xa4, 0x39, 0xe2, 0x06, 0xe2, 0x57, 0xab, 0x8d, 0x8b, 0xd2,
	0x56, 0x38, 0x64, 0x2d, 0x96, 0x63, 0xa8, 0x25, 0x9d, 0xe1, 0x0d, 0x44, 0x15, 0x67, 0x35, 0xe6,
	0x6e, 0xe3, 0xdb, 0x2d, 0x16, 0x23, 0xe4, 0x68, 0x14, 0xc2, 0x17, 0x76, 0xc0, 0x7f, 0xee, 0x79,
	0xff, 0x85, 0x8b, 0xfc, 0x0c, 0x1c, 0xe7, 0x16, 0xb2, 0x67, 0xc3, 0x15, 0x33, 0x42, 0x79, 0x39,
	0xdf, 0xb6, 0x97, 0x1b, 0xca, 0xa4, 0xb3, 0xfb, 0x60, 0x97, 0xb8, 0xdf, 0xea, 0xe1, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0xb7, 0x04, 0xf4, 0x63, 0x02, 0x00, 0x00,
}
