// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: receive_service.proto

package pb

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

// ReceiveServiceClient is the client API for ReceiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiveServiceClient interface {
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
}

type receiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiveServiceClient(cc grpc.ClientConnInterface) ReceiveServiceClient {
	return &receiveServiceClient{cc}
}

func (c *receiveServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/pb.ReceiveService/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiveServiceServer is the server API for ReceiveService service.
// All implementations must embed UnimplementedReceiveServiceServer
// for forward compatibility
type ReceiveServiceServer interface {
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	mustEmbedUnimplementedReceiveServiceServer()
}

// UnimplementedReceiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReceiveServiceServer struct {
}

func (UnimplementedReceiveServiceServer) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedReceiveServiceServer) mustEmbedUnimplementedReceiveServiceServer() {}

// UnsafeReceiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiveServiceServer will
// result in compilation errors.
type UnsafeReceiveServiceServer interface {
	mustEmbedUnimplementedReceiveServiceServer()
}

func RegisterReceiveServiceServer(s grpc.ServiceRegistrar, srv ReceiveServiceServer) {
	s.RegisterService(&ReceiveService_ServiceDesc, srv)
}

func _ReceiveService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiveServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ReceiveService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiveServiceServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReceiveService_ServiceDesc is the grpc.ServiceDesc for ReceiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReceiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ReceiveService",
	HandlerType: (*ReceiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _ReceiveService_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receive_service.proto",
}