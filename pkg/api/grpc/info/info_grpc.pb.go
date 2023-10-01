// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: info/info.proto

package info

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
	InfoService_Info_FullMethodName = "/info.InfoService/Info"
)

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoServiceClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, InfoService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
// All implementations must embed UnimplementedInfoServiceServer
// for forward compatibility
type InfoServiceServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedInfoServiceServer()
}

// UnimplementedInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (UnimplementedInfoServiceServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedInfoServiceServer) mustEmbedUnimplementedInfoServiceServer() {}

// UnsafeInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServiceServer will
// result in compilation errors.
type UnsafeInfoServiceServer interface {
	mustEmbedUnimplementedInfoServiceServer()
}

func RegisterInfoServiceServer(s grpc.ServiceRegistrar, srv InfoServiceServer) {
	s.RegisterService(&InfoService_ServiceDesc, srv)
}

func _InfoService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoService_ServiceDesc is the grpc.ServiceDesc for InfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _InfoService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info/info.proto",
}
