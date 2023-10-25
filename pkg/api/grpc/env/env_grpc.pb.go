// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: env/env.proto

package env

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
	EnvService_Env_FullMethodName = "/env.EnvService/Env"
)

// EnvServiceClient is the client API for EnvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvServiceClient interface {
	Env(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error)
}

type envServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvServiceClient(cc grpc.ClientConnInterface) EnvServiceClient {
	return &envServiceClient{cc}
}

func (c *envServiceClient) Env(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error) {
	out := new(EnvResponse)
	err := c.cc.Invoke(ctx, EnvService_Env_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvServiceServer is the server API for EnvService service.
// All implementations must embed UnimplementedEnvServiceServer
// for forward compatibility
type EnvServiceServer interface {
	Env(context.Context, *EnvRequest) (*EnvResponse, error)
	mustEmbedUnimplementedEnvServiceServer()
}

// UnimplementedEnvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnvServiceServer struct {
}

func (UnimplementedEnvServiceServer) Env(context.Context, *EnvRequest) (*EnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Env not implemented")
}
func (UnimplementedEnvServiceServer) mustEmbedUnimplementedEnvServiceServer() {}

// UnsafeEnvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvServiceServer will
// result in compilation errors.
type UnsafeEnvServiceServer interface {
	mustEmbedUnimplementedEnvServiceServer()
}

func RegisterEnvServiceServer(s grpc.ServiceRegistrar, srv EnvServiceServer) {
	s.RegisterService(&EnvService_ServiceDesc, srv)
}

func _EnvService_Env_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServiceServer).Env(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvService_Env_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServiceServer).Env(ctx, req.(*EnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvService_ServiceDesc is the grpc.ServiceDesc for EnvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "env.EnvService",
	HandlerType: (*EnvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Env",
			Handler:    _EnvService_Env_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "env/env.proto",
}