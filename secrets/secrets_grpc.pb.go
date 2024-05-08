// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: secrets.proto

package secrets

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
	SecretsService_GetAWSCredentials_FullMethodName = "/api.public.secrets.SecretsService/GetAWSCredentials"
)

// SecretsServiceClient is the client API for SecretsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretsServiceClient interface {
	// generates and returns a temporary aws session token for the given role_arn
	GetAWSCredentials(ctx context.Context, in *GetAWSCredentialsRequest, opts ...grpc.CallOption) (*GetAWSCredentialsResponse, error)
}

type secretsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsServiceClient(cc grpc.ClientConnInterface) SecretsServiceClient {
	return &secretsServiceClient{cc}
}

func (c *secretsServiceClient) GetAWSCredentials(ctx context.Context, in *GetAWSCredentialsRequest, opts ...grpc.CallOption) (*GetAWSCredentialsResponse, error) {
	out := new(GetAWSCredentialsResponse)
	err := c.cc.Invoke(ctx, SecretsService_GetAWSCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsServiceServer is the server API for SecretsService service.
// All implementations must embed UnimplementedSecretsServiceServer
// for forward compatibility
type SecretsServiceServer interface {
	// generates and returns a temporary aws session token for the given role_arn
	GetAWSCredentials(context.Context, *GetAWSCredentialsRequest) (*GetAWSCredentialsResponse, error)
	mustEmbedUnimplementedSecretsServiceServer()
}

// UnimplementedSecretsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretsServiceServer struct {
}

func (UnimplementedSecretsServiceServer) GetAWSCredentials(context.Context, *GetAWSCredentialsRequest) (*GetAWSCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAWSCredentials not implemented")
}
func (UnimplementedSecretsServiceServer) mustEmbedUnimplementedSecretsServiceServer() {}

// UnsafeSecretsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretsServiceServer will
// result in compilation errors.
type UnsafeSecretsServiceServer interface {
	mustEmbedUnimplementedSecretsServiceServer()
}

func RegisterSecretsServiceServer(s grpc.ServiceRegistrar, srv SecretsServiceServer) {
	s.RegisterService(&SecretsService_ServiceDesc, srv)
}

func _SecretsService_GetAWSCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAWSCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).GetAWSCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecretsService_GetAWSCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).GetAWSCredentials(ctx, req.(*GetAWSCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretsService_ServiceDesc is the grpc.ServiceDesc for SecretsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.secrets.SecretsService",
	HandlerType: (*SecretsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAWSCredentials",
			Handler:    _SecretsService_GetAWSCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secrets.proto",
}