// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: user_service.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetMoviesByCar(ctx context.Context, in *GetMoviesByCarRequest, opts ...grpc.CallOption) (*GetMoviesByCarResponse, error)
	GetCarsByMovie(ctx context.Context, in *GetCarsByMovieRequest, opts ...grpc.CallOption) (*GetCarsByMovieResponse, error)
	GetCarInfo(ctx context.Context, in *GetCarInfoRequest, opts ...grpc.CallOption) (*GetCarInfoResponse, error)
	GetMovieInfo(ctx context.Context, in *GetMovieInfoRequest, opts ...grpc.CallOption) (*GetMovieInfoResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetMoviesByCar(ctx context.Context, in *GetMoviesByCarRequest, opts ...grpc.CallOption) (*GetMoviesByCarResponse, error) {
	out := new(GetMoviesByCarResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetMoviesByCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCarsByMovie(ctx context.Context, in *GetCarsByMovieRequest, opts ...grpc.CallOption) (*GetCarsByMovieResponse, error) {
	out := new(GetCarsByMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetCarsByMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCarInfo(ctx context.Context, in *GetCarInfoRequest, opts ...grpc.CallOption) (*GetCarInfoResponse, error) {
	out := new(GetCarInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetCarInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMovieInfo(ctx context.Context, in *GetMovieInfoRequest, opts ...grpc.CallOption) (*GetMovieInfoResponse, error) {
	out := new(GetMovieInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetMovieInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetMoviesByCar(context.Context, *GetMoviesByCarRequest) (*GetMoviesByCarResponse, error)
	GetCarsByMovie(context.Context, *GetCarsByMovieRequest) (*GetCarsByMovieResponse, error)
	GetCarInfo(context.Context, *GetCarInfoRequest) (*GetCarInfoResponse, error)
	GetMovieInfo(context.Context, *GetMovieInfoRequest) (*GetMovieInfoResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetMoviesByCar(context.Context, *GetMoviesByCarRequest) (*GetMoviesByCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoviesByCar not implemented")
}
func (UnimplementedUserServiceServer) GetCarsByMovie(context.Context, *GetCarsByMovieRequest) (*GetCarsByMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarsByMovie not implemented")
}
func (UnimplementedUserServiceServer) GetCarInfo(context.Context, *GetCarInfoRequest) (*GetCarInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarInfo not implemented")
}
func (UnimplementedUserServiceServer) GetMovieInfo(context.Context, *GetMovieInfoRequest) (*GetMovieInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieInfo not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetMoviesByCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoviesByCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMoviesByCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetMoviesByCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMoviesByCar(ctx, req.(*GetMoviesByCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCarsByMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarsByMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCarsByMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetCarsByMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCarsByMovie(ctx, req.(*GetCarsByMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCarInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCarInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetCarInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCarInfo(ctx, req.(*GetCarInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMovieInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMovieInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetMovieInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMovieInfo(ctx, req.(*GetMovieInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMoviesByCar",
			Handler:    _UserService_GetMoviesByCar_Handler,
		},
		{
			MethodName: "GetCarsByMovie",
			Handler:    _UserService_GetCarsByMovie_Handler,
		},
		{
			MethodName: "GetCarInfo",
			Handler:    _UserService_GetCarInfo_Handler,
		},
		{
			MethodName: "GetMovieInfo",
			Handler:    _UserService_GetMovieInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}