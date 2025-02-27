// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: main.proto

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

// UrlShortenerClient is the client API for UrlShortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerClient interface {
	CreateShortenedUrl(ctx context.Context, in *CreateShortenedUrlRequest, opts ...grpc.CallOption) (*CreateShortenedUrlResponse, error)
	GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error)
}

type urlShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerClient(cc grpc.ClientConnInterface) UrlShortenerClient {
	return &urlShortenerClient{cc}
}

func (c *urlShortenerClient) CreateShortenedUrl(ctx context.Context, in *CreateShortenedUrlRequest, opts ...grpc.CallOption) (*CreateShortenedUrlResponse, error) {
	out := new(CreateShortenedUrlResponse)
	err := c.cc.Invoke(ctx, "/api.UrlShortener/CreateShortenedUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerClient) GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error) {
	out := new(GetOriginalURLResponse)
	err := c.cc.Invoke(ctx, "/api.UrlShortener/GetOriginalURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServer is the server API for UrlShortener service.
// All implementations must embed UnimplementedUrlShortenerServer
// for forward compatibility
type UrlShortenerServer interface {
	CreateShortenedUrl(context.Context, *CreateShortenedUrlRequest) (*CreateShortenedUrlResponse, error)
	GetOriginalURL(context.Context, *GetOriginalURLRequest) (*GetOriginalURLResponse, error)
	mustEmbedUnimplementedUrlShortenerServer()
}

// UnimplementedUrlShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortenerServer struct {
}

func (UnimplementedUrlShortenerServer) CreateShortenedUrl(context.Context, *CreateShortenedUrlRequest) (*CreateShortenedUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortenedUrl not implemented")
}
func (UnimplementedUrlShortenerServer) GetOriginalURL(context.Context, *GetOriginalURLRequest) (*GetOriginalURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalURL not implemented")
}
func (UnimplementedUrlShortenerServer) mustEmbedUnimplementedUrlShortenerServer() {}

// UnsafeUrlShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServer will
// result in compilation errors.
type UnsafeUrlShortenerServer interface {
	mustEmbedUnimplementedUrlShortenerServer()
}

func RegisterUrlShortenerServer(s grpc.ServiceRegistrar, srv UrlShortenerServer) {
	s.RegisterService(&UrlShortener_ServiceDesc, srv)
}

func _UrlShortener_CreateShortenedUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShortenedUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).CreateShortenedUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UrlShortener/CreateShortenedUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).CreateShortenedUrl(ctx, req.(*CreateShortenedUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortener_GetOriginalURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginalURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).GetOriginalURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UrlShortener/GetOriginalURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).GetOriginalURL(ctx, req.(*GetOriginalURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortener_ServiceDesc is the grpc.ServiceDesc for UrlShortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UrlShortener",
	HandlerType: (*UrlShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortenedUrl",
			Handler:    _UrlShortener_CreateShortenedUrl_Handler,
		},
		{
			MethodName: "GetOriginalURL",
			Handler:    _UrlShortener_GetOriginalURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
