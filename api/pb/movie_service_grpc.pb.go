// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: movie_service.proto

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

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*CreateMovieResponse, error)
	ReadMovie(ctx context.Context, in *ReadMovieRequest, opts ...grpc.CallOption) (*ReadMovieResponse, error)
	UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*UpdateMovieResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error)
	ListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (MovieService_ListMovieClient, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*CreateMovieResponse, error) {
	out := new(CreateMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieService/CreateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) ReadMovie(ctx context.Context, in *ReadMovieRequest, opts ...grpc.CallOption) (*ReadMovieResponse, error) {
	out := new(ReadMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieService/ReadMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*UpdateMovieResponse, error) {
	out := new(UpdateMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieService/UpdateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error) {
	out := new(DeleteMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieService/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) ListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (MovieService_ListMovieClient, error) {
	stream, err := c.cc.NewStream(ctx, &MovieService_ServiceDesc.Streams[0], "/pb.MovieService/ListMovie", opts...)
	if err != nil {
		return nil, err
	}
	x := &movieServiceListMovieClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MovieService_ListMovieClient interface {
	Recv() (*ListMovieResponse, error)
	grpc.ClientStream
}

type movieServiceListMovieClient struct {
	grpc.ClientStream
}

func (x *movieServiceListMovieClient) Recv() (*ListMovieResponse, error) {
	m := new(ListMovieResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	CreateMovie(context.Context, *CreateMovieRequest) (*CreateMovieResponse, error)
	ReadMovie(context.Context, *ReadMovieRequest) (*ReadMovieResponse, error)
	UpdateMovie(context.Context, *UpdateMovieRequest) (*UpdateMovieResponse, error)
	DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error)
	ListMovie(*ListMovieRequest, MovieService_ListMovieServer) error
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) CreateMovie(context.Context, *CreateMovieRequest) (*CreateMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedMovieServiceServer) ReadMovie(context.Context, *ReadMovieRequest) (*ReadMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMovie not implemented")
}
func (UnimplementedMovieServiceServer) UpdateMovie(context.Context, *UpdateMovieRequest) (*UpdateMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieServiceServer) DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieServiceServer) ListMovie(*ListMovieRequest, MovieService_ListMovieServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMovie not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieService/CreateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).CreateMovie(ctx, req.(*CreateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_ReadMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).ReadMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieService/ReadMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).ReadMovie(ctx, req.(*ReadMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieService/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).UpdateMovie(ctx, req.(*UpdateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieService/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_ListMovie_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMovieRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieServiceServer).ListMovie(m, &movieServiceListMovieServer{stream})
}

type MovieService_ListMovieServer interface {
	Send(*ListMovieResponse) error
	grpc.ServerStream
}

type movieServiceListMovieServer struct {
	grpc.ServerStream
}

func (x *movieServiceListMovieServer) Send(m *ListMovieResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMovie",
			Handler:    _MovieService_CreateMovie_Handler,
		},
		{
			MethodName: "ReadMovie",
			Handler:    _MovieService_ReadMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieService_DeleteMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMovie",
			Handler:       _MovieService_ListMovie_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "movie_service.proto",
}
