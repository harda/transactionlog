// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transactionlog

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	SaveTransactionLog(ctx context.Context, in *TransactionLog, opts ...grpc.CallOption) (*ID, error)
	SelectTransactionLog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TransactionLog, error)
	UpdateTransactionLog(ctx context.Context, in *TransactionLog, opts ...grpc.CallOption) (*ID, error)
	DeleteTransactionLog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) SaveTransactionLog(ctx context.Context, in *TransactionLog, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/TransactionService/SaveTransactionLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SelectTransactionLog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TransactionLog, error) {
	out := new(TransactionLog)
	err := c.cc.Invoke(ctx, "/TransactionService/SelectTransactionLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UpdateTransactionLog(ctx context.Context, in *TransactionLog, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/TransactionService/UpdateTransactionLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransactionLog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/TransactionService/DeleteTransactionLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	SaveTransactionLog(context.Context, *TransactionLog) (*ID, error)
	SelectTransactionLog(context.Context, *ID) (*TransactionLog, error)
	UpdateTransactionLog(context.Context, *TransactionLog) (*ID, error)
	DeleteTransactionLog(context.Context, *ID) (*ID, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) SaveTransactionLog(context.Context, *TransactionLog) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTransactionLog not implemented")
}
func (UnimplementedTransactionServiceServer) SelectTransactionLog(context.Context, *ID) (*TransactionLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectTransactionLog not implemented")
}
func (UnimplementedTransactionServiceServer) UpdateTransactionLog(context.Context, *TransactionLog) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactionLog not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteTransactionLog(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransactionLog not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_SaveTransactionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SaveTransactionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/SaveTransactionLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SaveTransactionLog(ctx, req.(*TransactionLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_SelectTransactionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SelectTransactionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/SelectTransactionLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SelectTransactionLog(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UpdateTransactionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UpdateTransactionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/UpdateTransactionLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UpdateTransactionLog(ctx, req.(*TransactionLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransactionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransactionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionService/DeleteTransactionLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransactionLog(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveTransactionLog",
			Handler:    _TransactionService_SaveTransactionLog_Handler,
		},
		{
			MethodName: "SelectTransactionLog",
			Handler:    _TransactionService_SelectTransactionLog_Handler,
		},
		{
			MethodName: "UpdateTransactionLog",
			Handler:    _TransactionService_UpdateTransactionLog_Handler,
		},
		{
			MethodName: "DeleteTransactionLog",
			Handler:    _TransactionService_DeleteTransactionLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
