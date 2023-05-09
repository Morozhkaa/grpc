// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/proto/mafia.proto

package proto

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

const (
	Mafia_JoinGame_FullMethodName      = "/mafia.Mafia/JoinGame"
	Mafia_GetGameStatus_FullMethodName = "/mafia.Mafia/GetGameStatus"
	Mafia_SendMessage_FullMethodName   = "/mafia.Mafia/SendMessage"
	Mafia_Kick_FullMethodName          = "/mafia.Mafia/Kick"
	Mafia_Kill_FullMethodName          = "/mafia.Mafia/Kill"
	Mafia_CheckTeam_FullMethodName     = "/mafia.Mafia/CheckTeam"
)

// MafiaClient is the client API for Mafia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MafiaClient interface {
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (Mafia_JoinGameClient, error)
	GetGameStatus(ctx context.Context, in *GetGameStatusRequest, opts ...grpc.CallOption) (*GetGameStatusResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	Kick(ctx context.Context, in *KickRequest, opts ...grpc.CallOption) (*KickResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*KillResponse, error)
	CheckTeam(ctx context.Context, in *CheckTeamRequest, opts ...grpc.CallOption) (*CheckTeamResponse, error)
}

type mafiaClient struct {
	cc grpc.ClientConnInterface
}

func NewMafiaClient(cc grpc.ClientConnInterface) MafiaClient {
	return &mafiaClient{cc}
}

func (c *mafiaClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (Mafia_JoinGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &Mafia_ServiceDesc.Streams[0], Mafia_JoinGame_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &mafiaJoinGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mafia_JoinGameClient interface {
	Recv() (*GameEvent, error)
	grpc.ClientStream
}

type mafiaJoinGameClient struct {
	grpc.ClientStream
}

func (x *mafiaJoinGameClient) Recv() (*GameEvent, error) {
	m := new(GameEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mafiaClient) GetGameStatus(ctx context.Context, in *GetGameStatusRequest, opts ...grpc.CallOption) (*GetGameStatusResponse, error) {
	out := new(GetGameStatusResponse)
	err := c.cc.Invoke(ctx, Mafia_GetGameStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Mafia_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) Kick(ctx context.Context, in *KickRequest, opts ...grpc.CallOption) (*KickResponse, error) {
	out := new(KickResponse)
	err := c.cc.Invoke(ctx, Mafia_Kick_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) Kill(ctx context.Context, in *KillRequest, opts ...grpc.CallOption) (*KillResponse, error) {
	out := new(KillResponse)
	err := c.cc.Invoke(ctx, Mafia_Kill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) CheckTeam(ctx context.Context, in *CheckTeamRequest, opts ...grpc.CallOption) (*CheckTeamResponse, error) {
	out := new(CheckTeamResponse)
	err := c.cc.Invoke(ctx, Mafia_CheckTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MafiaServer is the server API for Mafia service.
// All implementations must embed UnimplementedMafiaServer
// for forward compatibility
type MafiaServer interface {
	JoinGame(*JoinGameRequest, Mafia_JoinGameServer) error
	GetGameStatus(context.Context, *GetGameStatusRequest) (*GetGameStatusResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	Kick(context.Context, *KickRequest) (*KickResponse, error)
	Kill(context.Context, *KillRequest) (*KillResponse, error)
	CheckTeam(context.Context, *CheckTeamRequest) (*CheckTeamResponse, error)
	mustEmbedUnimplementedMafiaServer()
}

// UnimplementedMafiaServer must be embedded to have forward compatible implementations.
type UnimplementedMafiaServer struct {
}

func (UnimplementedMafiaServer) JoinGame(*JoinGameRequest, Mafia_JoinGameServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedMafiaServer) GetGameStatus(context.Context, *GetGameStatusRequest) (*GetGameStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameStatus not implemented")
}
func (UnimplementedMafiaServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMafiaServer) Kick(context.Context, *KickRequest) (*KickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kick not implemented")
}
func (UnimplementedMafiaServer) Kill(context.Context, *KillRequest) (*KillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kill not implemented")
}
func (UnimplementedMafiaServer) CheckTeam(context.Context, *CheckTeamRequest) (*CheckTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTeam not implemented")
}
func (UnimplementedMafiaServer) mustEmbedUnimplementedMafiaServer() {}

// UnsafeMafiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MafiaServer will
// result in compilation errors.
type UnsafeMafiaServer interface {
	mustEmbedUnimplementedMafiaServer()
}

func RegisterMafiaServer(s grpc.ServiceRegistrar, srv MafiaServer) {
	s.RegisterService(&Mafia_ServiceDesc, srv)
}

func _Mafia_JoinGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinGameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MafiaServer).JoinGame(m, &mafiaJoinGameServer{stream})
}

type Mafia_JoinGameServer interface {
	Send(*GameEvent) error
	grpc.ServerStream
}

type mafiaJoinGameServer struct {
	grpc.ServerStream
}

func (x *mafiaJoinGameServer) Send(m *GameEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Mafia_GetGameStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).GetGameStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mafia_GetGameStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).GetGameStatus(ctx, req.(*GetGameStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mafia_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_Kick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).Kick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mafia_Kick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).Kick(ctx, req.(*KickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_Kill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).Kill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mafia_Kill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).Kill(ctx, req.(*KillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_CheckTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).CheckTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mafia_CheckTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).CheckTeam(ctx, req.(*CheckTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mafia_ServiceDesc is the grpc.ServiceDesc for Mafia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mafia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mafia.Mafia",
	HandlerType: (*MafiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGameStatus",
			Handler:    _Mafia_GetGameStatus_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Mafia_SendMessage_Handler,
		},
		{
			MethodName: "Kick",
			Handler:    _Mafia_Kick_Handler,
		},
		{
			MethodName: "Kill",
			Handler:    _Mafia_Kill_Handler,
		},
		{
			MethodName: "CheckTeam",
			Handler:    _Mafia_CheckTeam_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinGame",
			Handler:       _Mafia_JoinGame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/proto/mafia.proto",
}