// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/shortened_links/service.proto

package shortened_links

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
	ShortenedLinks_SaveLink_FullMethodName = "/shortenedLinks/SaveLink"
	ShortenedLinks_GetLink_FullMethodName  = "/shortenedLinks/GetLink"
)

// ShortenedLinksClient is the client API for ShortenedLinks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenedLinksClient interface {
	SaveLink(ctx context.Context, in *OriginalLink, opts ...grpc.CallOption) (*ShortenedLink, error)
	GetLink(ctx context.Context, in *ShortenedLink, opts ...grpc.CallOption) (*OriginalLink, error)
}

type shortenedLinksClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenedLinksClient(cc grpc.ClientConnInterface) ShortenedLinksClient {
	return &shortenedLinksClient{cc}
}

func (c *shortenedLinksClient) SaveLink(ctx context.Context, in *OriginalLink, opts ...grpc.CallOption) (*ShortenedLink, error) {
	out := new(ShortenedLink)
	err := c.cc.Invoke(ctx, ShortenedLinks_SaveLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenedLinksClient) GetLink(ctx context.Context, in *ShortenedLink, opts ...grpc.CallOption) (*OriginalLink, error) {
	out := new(OriginalLink)
	err := c.cc.Invoke(ctx, ShortenedLinks_GetLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenedLinksServer is the server API for ShortenedLinks service.
// All implementations must embed UnimplementedShortenedLinksServer
// for forward compatibility
type ShortenedLinksServer interface {
	SaveLink(context.Context, *OriginalLink) (*ShortenedLink, error)
	GetLink(context.Context, *ShortenedLink) (*OriginalLink, error)
	mustEmbedUnimplementedShortenedLinksServer()
}

// UnimplementedShortenedLinksServer must be embedded to have forward compatible implementations.
type UnimplementedShortenedLinksServer struct {
}

func (UnimplementedShortenedLinksServer) SaveLink(context.Context, *OriginalLink) (*ShortenedLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLink not implemented")
}
func (UnimplementedShortenedLinksServer) GetLink(context.Context, *ShortenedLink) (*OriginalLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (UnimplementedShortenedLinksServer) mustEmbedUnimplementedShortenedLinksServer() {}

// UnsafeShortenedLinksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenedLinksServer will
// result in compilation errors.
type UnsafeShortenedLinksServer interface {
	mustEmbedUnimplementedShortenedLinksServer()
}

func RegisterShortenedLinksServer(s grpc.ServiceRegistrar, srv ShortenedLinksServer) {
	s.RegisterService(&ShortenedLinks_ServiceDesc, srv)
}

func _ShortenedLinks_SaveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginalLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenedLinksServer).SaveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenedLinks_SaveLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenedLinksServer).SaveLink(ctx, req.(*OriginalLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenedLinks_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenedLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenedLinksServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenedLinks_GetLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenedLinksServer).GetLink(ctx, req.(*ShortenedLink))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenedLinks_ServiceDesc is the grpc.ServiceDesc for ShortenedLinks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenedLinks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortenedLinks",
	HandlerType: (*ShortenedLinksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveLink",
			Handler:    _ShortenedLinks_SaveLink_Handler,
		},
		{
			MethodName: "GetLink",
			Handler:    _ShortenedLinks_GetLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shortened_links/service.proto",
}