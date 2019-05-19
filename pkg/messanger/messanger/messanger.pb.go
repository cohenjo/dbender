// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messanger.proto

package messanger

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

// The request message containing the user's name.
type MessageRequest struct {
	Msg                  string            `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Body                 string            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Channel              string            `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Type                 string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Bag                  map[string]string `protobuf:"bytes,5,rep,name=bag,proto3" json:"bag,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messanger_a9b32a3beb198ac9, []int{0}
}
func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (dst *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(dst, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MessageRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MessageRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *MessageRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MessageRequest) GetBag() map[string]string {
	if m != nil {
		return m.Bag
	}
	return nil
}

// The response message containing the greetings
type MessageReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReply) Reset()         { *m = MessageReply{} }
func (m *MessageReply) String() string { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()    {}
func (*MessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_messanger_a9b32a3beb198ac9, []int{1}
}
func (m *MessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReply.Unmarshal(m, b)
}
func (m *MessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReply.Marshal(b, m, deterministic)
}
func (dst *MessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReply.Merge(dst, src)
}
func (m *MessageReply) XXX_Size() int {
	return xxx_messageInfo_MessageReply.Size(m)
}
func (m *MessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReply proto.InternalMessageInfo

func (m *MessageReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "MessageRequest")
	proto.RegisterMapType((map[string]string)(nil), "MessageRequest.BagEntry")
	proto.RegisterType((*MessageReply)(nil), "MessageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessangerClient is the client API for Messanger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessangerClient interface {
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type messangerClient struct {
	cc *grpc.ClientConn
}

func NewMessangerClient(cc *grpc.ClientConn) MessangerClient {
	return &messangerClient{cc}
}

func (c *messangerClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/messanger/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessangerServer is the server API for Messanger service.
type MessangerServer interface {
	SendMessage(context.Context, *MessageRequest) (*MessageReply, error)
}

func RegisterMessangerServer(s *grpc.Server, srv MessangerServer) {
	s.RegisterService(&_Messanger_serviceDesc, srv)
}

func _Messanger_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessangerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messanger/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessangerServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messanger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messanger",
	HandlerType: (*MessangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messanger_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messanger.proto",
}

func init() { proto.RegisterFile("messanger.proto", fileDescriptor_messanger_a9b32a3beb198ac9) }

var fileDescriptor_messanger_a9b32a3beb198ac9 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0xcc, 0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x3a, 0xc1, 0xc8, 0xc5,
	0xe7, 0x0b, 0x12, 0x4b, 0x4f, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62,
	0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x84, 0xb8, 0x58,
	0x92, 0xf2, 0x53, 0x2a, 0x25, 0x98, 0xc0, 0x42, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x46,
	0x62, 0x5e, 0x5e, 0x6a, 0x8e, 0x04, 0x33, 0x58, 0x18, 0xc6, 0x05, 0xa9, 0x2e, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x81, 0xa8, 0x06, 0xb1, 0x85, 0xb4, 0xb8, 0x98, 0x93, 0x12, 0xd3, 0x25, 0x58, 0x15,
	0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf4, 0x50, 0x6d, 0xd4, 0x73, 0x4a, 0x4c, 0x77, 0xcd, 0x2b, 0x29,
	0xaa, 0x0c, 0x02, 0x29, 0x92, 0x32, 0xe3, 0xe2, 0x80, 0x09, 0x80, 0xdc, 0x92, 0x9d, 0x5a, 0x09,
	0x73, 0x4b, 0x76, 0x6a, 0xa5, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0xd4, 0x31,
	0x10, 0x8e, 0x15, 0x93, 0x05, 0xa3, 0x92, 0x06, 0x17, 0x0f, 0xdc, 0xdc, 0x82, 0x1c, 0xb0, 0x0b,
	0x73, 0x21, 0x7c, 0xa8, 0x7e, 0x18, 0xd7, 0xc8, 0x86, 0x8b, 0x13, 0x1e, 0x0e, 0x42, 0xfa, 0x5c,
	0xdc, 0xc1, 0xa9, 0x79, 0x29, 0x50, 0xad, 0x42, 0xfc, 0x68, 0x8e, 0x93, 0xe2, 0xd5, 0x43, 0x36,
	0x55, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xa9,
	0x9d, 0x38, 0x4c, 0x01, 0x00, 0x00,
}
