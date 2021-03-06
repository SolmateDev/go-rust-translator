// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/solana.proto

package solana

import (
	context "context"
	basic "github.com/SolmateDev/go-rust-translator/basic"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastClient interface {
	//rpc SendTx(SendTxRequest) returns (SendTxResponse) {}
	SendTx(ctx context.Context, in *SendBatchRequest, opts ...grpc.CallOption) (*SendBatchResponse, error)
	RunGenesis(ctx context.Context, in *Genesis, opts ...grpc.CallOption) (*basic.Empty, error)
	RunInitializeTokenAccount(ctx context.Context, in *InitializeTokenAccount, opts ...grpc.CallOption) (*basic.Keypair, error)
	RunCreateAccount(ctx context.Context, in *CreateAccount, opts ...grpc.CallOption) (*basic.Keypair, error)
	RunMint(ctx context.Context, in *Mint, opts ...grpc.CallOption) (*basic.Empty, error)
}

type broadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastClient(cc grpc.ClientConnInterface) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) SendTx(ctx context.Context, in *SendBatchRequest, opts ...grpc.CallOption) (*SendBatchResponse, error) {
	out := new(SendBatchResponse)
	err := c.cc.Invoke(ctx, "/solana.Broadcast/SendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) RunGenesis(ctx context.Context, in *Genesis, opts ...grpc.CallOption) (*basic.Empty, error) {
	out := new(basic.Empty)
	err := c.cc.Invoke(ctx, "/solana.Broadcast/RunGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) RunInitializeTokenAccount(ctx context.Context, in *InitializeTokenAccount, opts ...grpc.CallOption) (*basic.Keypair, error) {
	out := new(basic.Keypair)
	err := c.cc.Invoke(ctx, "/solana.Broadcast/RunInitializeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) RunCreateAccount(ctx context.Context, in *CreateAccount, opts ...grpc.CallOption) (*basic.Keypair, error) {
	out := new(basic.Keypair)
	err := c.cc.Invoke(ctx, "/solana.Broadcast/RunCreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) RunMint(ctx context.Context, in *Mint, opts ...grpc.CallOption) (*basic.Empty, error) {
	out := new(basic.Empty)
	err := c.cc.Invoke(ctx, "/solana.Broadcast/RunMint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
// All implementations must embed UnimplementedBroadcastServer
// for forward compatibility
type BroadcastServer interface {
	//rpc SendTx(SendTxRequest) returns (SendTxResponse) {}
	SendTx(context.Context, *SendBatchRequest) (*SendBatchResponse, error)
	RunGenesis(context.Context, *Genesis) (*basic.Empty, error)
	RunInitializeTokenAccount(context.Context, *InitializeTokenAccount) (*basic.Keypair, error)
	RunCreateAccount(context.Context, *CreateAccount) (*basic.Keypair, error)
	RunMint(context.Context, *Mint) (*basic.Empty, error)
	mustEmbedUnimplementedBroadcastServer()
}

// UnimplementedBroadcastServer must be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (UnimplementedBroadcastServer) SendTx(context.Context, *SendBatchRequest) (*SendBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedBroadcastServer) RunGenesis(context.Context, *Genesis) (*basic.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunGenesis not implemented")
}
func (UnimplementedBroadcastServer) RunInitializeTokenAccount(context.Context, *InitializeTokenAccount) (*basic.Keypair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunInitializeTokenAccount not implemented")
}
func (UnimplementedBroadcastServer) RunCreateAccount(context.Context, *CreateAccount) (*basic.Keypair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCreateAccount not implemented")
}
func (UnimplementedBroadcastServer) RunMint(context.Context, *Mint) (*basic.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunMint not implemented")
}
func (UnimplementedBroadcastServer) mustEmbedUnimplementedBroadcastServer() {}

// UnsafeBroadcastServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServer will
// result in compilation errors.
type UnsafeBroadcastServer interface {
	mustEmbedUnimplementedBroadcastServer()
}

func RegisterBroadcastServer(s grpc.ServiceRegistrar, srv BroadcastServer) {
	s.RegisterService(&Broadcast_ServiceDesc, srv)
}

func _Broadcast_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.Broadcast/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).SendTx(ctx, req.(*SendBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_RunGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Genesis)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).RunGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.Broadcast/RunGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).RunGenesis(ctx, req.(*Genesis))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_RunInitializeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeTokenAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).RunInitializeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.Broadcast/RunInitializeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).RunInitializeTokenAccount(ctx, req.(*InitializeTokenAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_RunCreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).RunCreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.Broadcast/RunCreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).RunCreateAccount(ctx, req.(*CreateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_RunMint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).RunMint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solana.Broadcast/RunMint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).RunMint(ctx, req.(*Mint))
	}
	return interceptor(ctx, in, info, handler)
}

// Broadcast_ServiceDesc is the grpc.ServiceDesc for Broadcast service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broadcast_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solana.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTx",
			Handler:    _Broadcast_SendTx_Handler,
		},
		{
			MethodName: "RunGenesis",
			Handler:    _Broadcast_RunGenesis_Handler,
		},
		{
			MethodName: "RunInitializeTokenAccount",
			Handler:    _Broadcast_RunInitializeTokenAccount_Handler,
		},
		{
			MethodName: "RunCreateAccount",
			Handler:    _Broadcast_RunCreateAccount_Handler,
		},
		{
			MethodName: "RunMint",
			Handler:    _Broadcast_RunMint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/solana.proto",
}
