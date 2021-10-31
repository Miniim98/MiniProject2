// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChittychatClient is the client API for Chittychat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittychatClient interface {
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (Chittychat_PublishClient, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (Chittychat_BroadcastClient, error)
}

type chittychatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittychatClient(cc grpc.ClientConnInterface) ChittychatClient {
	return &chittychatClient{cc}
}

func (c *chittychatClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/Chat.Chittychat/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Publish(ctx context.Context, opts ...grpc.CallOption) (Chittychat_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chittychat_ServiceDesc.Streams[0], "/Chat.Chittychat/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittychatPublishClient{stream}
	return x, nil
}

type Chittychat_PublishClient interface {
	Send(*PublishRequest) error
	CloseAndRecv() (*PublishResponse, error)
	grpc.ClientStream
}

type chittychatPublishClient struct {
	grpc.ClientStream
}

func (x *chittychatPublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittychatPublishClient) CloseAndRecv() (*PublishResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittychatClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (Chittychat_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chittychat_ServiceDesc.Streams[1], "/Chat.Chittychat/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittychatBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chittychat_BroadcastClient interface {
	Recv() (*BroadcastResponse, error)
	grpc.ClientStream
}

type chittychatBroadcastClient struct {
	grpc.ClientStream
}

func (x *chittychatBroadcastClient) Recv() (*BroadcastResponse, error) {
	m := new(BroadcastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittychatServer is the server API for Chittychat service.
// All implementations must embed UnimplementedChittychatServer
// for forward compatibility
type ChittychatServer interface {
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	Publish(Chittychat_PublishServer) error
	Broadcast(*BroadcastRequest, Chittychat_BroadcastServer) error
	mustEmbedUnimplementedChittychatServer()
}

// UnimplementedChittychatServer must be embedded to have forward compatible implementations.
type UnimplementedChittychatServer struct {
}

func (UnimplementedChittychatServer) Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChittychatServer) Publish(Chittychat_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittychatServer) Broadcast(*BroadcastRequest, Chittychat_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChittychatServer) mustEmbedUnimplementedChittychatServer() {}

// UnsafeChittychatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittychatServer will
// result in compilation errors.
type UnsafeChittychatServer interface {
	mustEmbedUnimplementedChittychatServer()
}

func RegisterChittychatServer(s grpc.ServiceRegistrar, srv ChittychatServer) {
	s.RegisterService(&Chittychat_ServiceDesc, srv)
}

func _Chittychat_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat.Chittychat/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittychatServer).Publish(&chittychatPublishServer{stream})
}

type Chittychat_PublishServer interface {
	SendAndClose(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type chittychatPublishServer struct {
	grpc.ServerStream
}

func (x *chittychatPublishServer) SendAndClose(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittychatPublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chittychat_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittychatServer).Broadcast(m, &chittychatBroadcastServer{stream})
}

type Chittychat_BroadcastServer interface {
	Send(*BroadcastResponse) error
	grpc.ServerStream
}

type chittychatBroadcastServer struct {
	grpc.ServerStream
}

func (x *chittychatBroadcastServer) Send(m *BroadcastResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Chittychat_ServiceDesc is the grpc.ServiceDesc for Chittychat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chittychat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chat.Chittychat",
	HandlerType: (*ChittychatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Chittychat_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _Chittychat_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _Chittychat_Broadcast_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Chat/Chat.proto",
}
