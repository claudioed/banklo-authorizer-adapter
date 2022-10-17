// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// LedgerServiceClient is the client API for LedgerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerServiceClient interface {
	Confirmation(ctx context.Context, in *RequestConfirmation, opts ...grpc.CallOption) (*TransactionResponse, error)
	Cancellation(ctx context.Context, in *RequestCancellation, opts ...grpc.CallOption) (*TransactionResponse, error)
	Reversal(ctx context.Context, in *RequestReversal, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type ledgerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerServiceClient(cc grpc.ClientConnInterface) LedgerServiceClient {
	return &ledgerServiceClient{cc}
}

func (c *ledgerServiceClient) Confirmation(ctx context.Context, in *RequestConfirmation, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/LedgerService/Confirmation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) Cancellation(ctx context.Context, in *RequestCancellation, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/LedgerService/Cancellation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledgerServiceClient) Reversal(ctx context.Context, in *RequestReversal, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/LedgerService/Reversal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerServiceServer is the server API for LedgerService service.
// All implementations must embed UnimplementedLedgerServiceServer
// for forward compatibility
type LedgerServiceServer interface {
	Confirmation(context.Context, *RequestConfirmation) (*TransactionResponse, error)
	Cancellation(context.Context, *RequestCancellation) (*TransactionResponse, error)
	Reversal(context.Context, *RequestReversal) (*TransactionResponse, error)
	mustEmbedUnimplementedLedgerServiceServer()
}

// UnimplementedLedgerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLedgerServiceServer struct {
}

func (UnimplementedLedgerServiceServer) Confirmation(context.Context, *RequestConfirmation) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirmation not implemented")
}
func (UnimplementedLedgerServiceServer) Cancellation(context.Context, *RequestCancellation) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancellation not implemented")
}
func (UnimplementedLedgerServiceServer) Reversal(context.Context, *RequestReversal) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reversal not implemented")
}
func (UnimplementedLedgerServiceServer) mustEmbedUnimplementedLedgerServiceServer() {}

// UnsafeLedgerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerServiceServer will
// result in compilation errors.
type UnsafeLedgerServiceServer interface {
	mustEmbedUnimplementedLedgerServiceServer()
}

func RegisterLedgerServiceServer(s grpc.ServiceRegistrar, srv LedgerServiceServer) {
	s.RegisterService(&LedgerService_ServiceDesc, srv)
}

func _LedgerService_Confirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConfirmation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).Confirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LedgerService/Confirmation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).Confirmation(ctx, req.(*RequestConfirmation))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_Cancellation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCancellation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).Cancellation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LedgerService/Cancellation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).Cancellation(ctx, req.(*RequestCancellation))
	}
	return interceptor(ctx, in, info, handler)
}

func _LedgerService_Reversal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReversal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerServiceServer).Reversal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LedgerService/Reversal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerServiceServer).Reversal(ctx, req.(*RequestReversal))
	}
	return interceptor(ctx, in, info, handler)
}

// LedgerService_ServiceDesc is the grpc.ServiceDesc for LedgerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LedgerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LedgerService",
	HandlerType: (*LedgerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Confirmation",
			Handler:    _LedgerService_Confirmation_Handler,
		},
		{
			MethodName: "Cancellation",
			Handler:    _LedgerService_Cancellation_Handler,
		},
		{
			MethodName: "Reversal",
			Handler:    _LedgerService_Reversal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ledger.proto",
}
