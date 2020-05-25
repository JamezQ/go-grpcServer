// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chatRpc.proto

package proto

import (
	context "context"
	fmt "fmt"
	protoImport "github.com/JamezQ/go-grpcServer/protos/protoImport"
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

type ChannelChat struct {
	Msg                  *protoImport.ChatMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Channel              *Channel                 `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ChannelChat) Reset()         { *m = ChannelChat{} }
func (m *ChannelChat) String() string { return proto.CompactTextString(m) }
func (*ChannelChat) ProtoMessage()    {}
func (*ChannelChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_37a4d8a79ac6f38c, []int{0}
}

func (m *ChannelChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelChat.Unmarshal(m, b)
}
func (m *ChannelChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelChat.Marshal(b, m, deterministic)
}
func (m *ChannelChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelChat.Merge(m, src)
}
func (m *ChannelChat) XXX_Size() int {
	return xxx_messageInfo_ChannelChat.Size(m)
}
func (m *ChannelChat) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelChat.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelChat proto.InternalMessageInfo

func (m *ChannelChat) GetMsg() *protoImport.ChatMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ChannelChat) GetChannel() *Channel {
	if m != nil {
		return m.Channel
	}
	return nil
}

type ChannelHistory struct {
	Msgs                 []*ChannelChat `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChannelHistory) Reset()         { *m = ChannelHistory{} }
func (m *ChannelHistory) String() string { return proto.CompactTextString(m) }
func (*ChannelHistory) ProtoMessage()    {}
func (*ChannelHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_37a4d8a79ac6f38c, []int{1}
}

func (m *ChannelHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelHistory.Unmarshal(m, b)
}
func (m *ChannelHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelHistory.Marshal(b, m, deterministic)
}
func (m *ChannelHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelHistory.Merge(m, src)
}
func (m *ChannelHistory) XXX_Size() int {
	return xxx_messageInfo_ChannelHistory.Size(m)
}
func (m *ChannelHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelHistory proto.InternalMessageInfo

func (m *ChannelHistory) GetMsgs() []*ChannelChat {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type Channel struct {
	ChName               string   `protobuf:"bytes,1,opt,name=chName,proto3" json:"chName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_37a4d8a79ac6f38c, []int{2}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetChName() string {
	if m != nil {
		return m.ChName
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_37a4d8a79ac6f38c, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChannelChat)(nil), "ChannelChat")
	proto.RegisterType((*ChannelHistory)(nil), "ChannelHistory")
	proto.RegisterType((*Channel)(nil), "Channel")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() {
	proto.RegisterFile("proto/chatRpc.proto", fileDescriptor_37a4d8a79ac6f38c)
}

var fileDescriptor_37a4d8a79ac6f38c = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0x4d, 0x4b, 0xc4, 0x30,
	0x14, 0xdc, 0xb8, 0xda, 0xea, 0xab, 0x1f, 0x10, 0x61, 0x29, 0x3d, 0x48, 0xcd, 0xc5, 0x05, 0x35,
	0x5d, 0xea, 0x3f, 0xb0, 0x88, 0x1f, 0xa8, 0xb0, 0xf5, 0xe6, 0x2d, 0x1b, 0x42, 0x5a, 0x30, 0x4d,
	0x49, 0xa2, 0xb0, 0xfe, 0x0d, 0xff, 0xb0, 0x34, 0x6d, 0x85, 0x7a, 0x1a, 0xde, 0xcc, 0xbc, 0x37,
	0xc3, 0x83, 0xd3, 0xd6, 0x68, 0xa7, 0x33, 0x5e, 0x31, 0x57, 0xb6, 0x9c, 0xfa, 0x29, 0x59, 0x78,
	0x78, 0x54, 0xad, 0x36, 0xce, 0x4b, 0x3d, 0x4f, 0xd6, 0x10, 0x15, 0x15, 0x6b, 0x1a, 0xf1, 0x51,
	0x54, 0xcc, 0xe1, 0x33, 0x98, 0x2b, 0x2b, 0x63, 0x94, 0xa2, 0x65, 0x94, 0x1f, 0xd2, 0x8e, 0x7b,
	0x11, 0xd6, 0x32, 0x29, 0xca, 0x4e, 0xc0, 0x04, 0x42, 0xde, 0xdb, 0xe3, 0x1d, 0xef, 0xd9, 0xa7,
	0xc3, 0x7a, 0x39, 0x0a, 0x24, 0x87, 0xe3, 0x81, 0x7b, 0xa8, 0xad, 0xd3, 0x66, 0x8b, 0x53, 0xd8,
	0x55, 0x56, 0xda, 0x18, 0xa5, 0xf3, 0xf1, 0xec, 0x98, 0x58, 0x7a, 0x85, 0x9c, 0x43, 0x38, 0x90,
	0x78, 0x01, 0x01, 0xaf, 0x5e, 0x99, 0x12, 0xbe, 0xc5, 0x41, 0x39, 0x4c, 0x24, 0x84, 0xbd, 0x3b,
	0xd5, 0xba, 0x6d, 0xfe, 0x83, 0x00, 0xba, 0xd5, 0x37, 0x61, 0xbe, 0x84, 0xc1, 0x17, 0x00, 0xf7,
	0xc2, 0x8d, 0x51, 0x01, 0xf5, 0xa6, 0xe4, 0x84, 0x4e, 0x3b, 0x90, 0x19, 0xbe, 0x84, 0xa3, 0xe7,
	0xda, 0x3a, 0xd1, 0x8c, 0x49, 0x7f, 0xdd, 0x93, 0x49, 0x25, 0x32, 0x5b, 0x21, 0x9c, 0x41, 0x54,
	0xe8, 0xa6, 0x11, 0xdc, 0xf9, 0xbf, 0x4c, 0x0c, 0xff, 0xed, 0x4b, 0xb4, 0x42, 0xb7, 0xf4, 0xfd,
	0x4a, 0xd6, 0xae, 0xfa, 0xdc, 0x50, 0xae, 0x55, 0xf6, 0xc4, 0x94, 0xf8, 0x5e, 0x67, 0x52, 0x5f,
	0x4b, 0xd3, 0xf2, 0xbe, 0x69, 0xe6, 0x3f, 0x6e, 0x7b, 0xd8, 0x04, 0x1e, 0x6e, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x59, 0xa5, 0xff, 0xae, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServerClient is the client API for ChatServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServerClient interface {
	GetHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChannelHistory, error)
	ListenChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatServer_ListenChannelClient, error)
	ConnectChat(ctx context.Context, opts ...grpc.CallOption) (ChatServer_ConnectChatClient, error)
}

type chatServerClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServerClient(cc grpc.ClientConnInterface) ChatServerClient {
	return &chatServerClient{cc}
}

func (c *chatServerClient) GetHistory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChannelHistory, error) {
	out := new(ChannelHistory)
	err := c.cc.Invoke(ctx, "/ChatServer/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServerClient) ListenChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatServer_ListenChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatServer_serviceDesc.Streams[0], "/ChatServer/ListenChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServerListenChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatServer_ListenChannelClient interface {
	Recv() (*ChannelChat, error)
	grpc.ClientStream
}

type chatServerListenChannelClient struct {
	grpc.ClientStream
}

func (x *chatServerListenChannelClient) Recv() (*ChannelChat, error) {
	m := new(ChannelChat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServerClient) ConnectChat(ctx context.Context, opts ...grpc.CallOption) (ChatServer_ConnectChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatServer_serviceDesc.Streams[1], "/ChatServer/ConnectChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServerConnectChatClient{stream}
	return x, nil
}

type ChatServer_ConnectChatClient interface {
	Send(*ChannelChat) error
	Recv() (*ChannelChat, error)
	grpc.ClientStream
}

type chatServerConnectChatClient struct {
	grpc.ClientStream
}

func (x *chatServerConnectChatClient) Send(m *ChannelChat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServerConnectChatClient) Recv() (*ChannelChat, error) {
	m := new(ChannelChat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServerServer is the server API for ChatServer service.
type ChatServerServer interface {
	GetHistory(context.Context, *Empty) (*ChannelHistory, error)
	ListenChannel(*Channel, ChatServer_ListenChannelServer) error
	ConnectChat(ChatServer_ConnectChatServer) error
}

// UnimplementedChatServerServer can be embedded to have forward compatible implementations.
type UnimplementedChatServerServer struct {
}

func (*UnimplementedChatServerServer) GetHistory(ctx context.Context, req *Empty) (*ChannelHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (*UnimplementedChatServerServer) ListenChannel(req *Channel, srv ChatServer_ListenChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenChannel not implemented")
}
func (*UnimplementedChatServerServer) ConnectChat(srv ChatServer_ConnectChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectChat not implemented")
}

func RegisterChatServerServer(s *grpc.Server, srv ChatServerServer) {
	s.RegisterService(&_ChatServer_serviceDesc, srv)
}

func _ChatServer_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServerServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatServer/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServerServer).GetHistory(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServer_ListenChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServerServer).ListenChannel(m, &chatServerListenChannelServer{stream})
}

type ChatServer_ListenChannelServer interface {
	Send(*ChannelChat) error
	grpc.ServerStream
}

type chatServerListenChannelServer struct {
	grpc.ServerStream
}

func (x *chatServerListenChannelServer) Send(m *ChannelChat) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatServer_ConnectChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServerServer).ConnectChat(&chatServerConnectChatServer{stream})
}

type ChatServer_ConnectChatServer interface {
	Send(*ChannelChat) error
	Recv() (*ChannelChat, error)
	grpc.ServerStream
}

type chatServerConnectChatServer struct {
	grpc.ServerStream
}

func (x *chatServerConnectChatServer) Send(m *ChannelChat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServerConnectChatServer) Recv() (*ChannelChat, error) {
	m := new(ChannelChat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatServer",
	HandlerType: (*ChatServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHistory",
			Handler:    _ChatServer_GetHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenChannel",
			Handler:       _ChatServer_ListenChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConnectChat",
			Handler:       _ChatServer_ConnectChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chatRpc.proto",
}