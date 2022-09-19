// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pipelines.proto

package pipelines

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

// PipelinesClient is the client API for Pipelines service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelinesClient interface {
	// ManualBuild manually triggers a Pipeline build with the
	// provided Targets, Args and Secrets. The new build_id is returned.
	ManualBuild(ctx context.Context, in *ManualBuildRequest, opts ...grpc.CallOption) (*ManualBuildResponse, error)
	// LaunchSatellite starts a new Earthly Satellite instance on the latest version.
	// The instance can be used by users to build their local, Cloud, or 3rd-party CI builds.
	LaunchSatellite(ctx context.Context, in *LaunchSatelliteRequest, opts ...grpc.CallOption) (*LaunchSatelliteResponse, error)
	// ListSatellites returns a list of Earthly Satellite instances available in the organization.
	ListSatellites(ctx context.Context, in *ListSatellitesRequest, opts ...grpc.CallOption) (*ListSatellitesResponse, error)
	// UpdateSatellite updates a Satellite instance to the latest version.
	// (I.e. the latest AMI we have in AWS).
	// Calling this may result in some down-time on the instance while it updates.
	UpdateSatellite(ctx context.Context, in *UpdateSatelliteRequest, opts ...grpc.CallOption) (*UpdateSatelliteResponse, error)
	// DeleteSatellite decommissions a Satellite instance.
	DeleteSatellite(ctx context.Context, in *DeleteSatelliteRequest, opts ...grpc.CallOption) (*DeleteSatelliteResponse, error)
	// GetSatellite retrieves the details of a particular Satellite instance.
	// Mainly intended for use by Buildkit Proxy when establishing a new connection to an instance.
	GetSatellite(ctx context.Context, in *GetSatelliteRequest, opts ...grpc.CallOption) (*GetSatelliteResponse, error)
	WakeSatellite(ctx context.Context, in *WakeSatelliteRequest, opts ...grpc.CallOption) (*WakeSatelliteResponse, error)
	// ListRemoteRepos uses the GitHub API to list remote repositories.
	ListRemoteRepos(ctx context.Context, in *ListRemoteReposRequest, opts ...grpc.CallOption) (*ListRemoteReposResponse, error)
	// AddProjectRepos adds one or more repositories to a project.
	AddProjectRepos(ctx context.Context, in *AddProjectReposRequest, opts ...grpc.CallOption) (*AddProjectReposResponse, error)
	// RemoveProjectRepo removes a repository from a project.
	RemoveProjectRepo(ctx context.Context, in *RemoveProjectRepoRequest, opts ...grpc.CallOption) (*RemoveProjectRepoResponse, error)
	// ListProjectRespos lists all project repositories.
	ListProjectRepos(ctx context.Context, in *ListProjectReposRequest, opts ...grpc.CallOption) (*ListProjectReposResponse, error)
	// ListRemotePipelines uses the GitHub API to list pipeline definitions present in a remote repository.
	ListRemotePipelines(ctx context.Context, in *ListRemotePipelinesRequest, opts ...grpc.CallOption) (*ListRemotePipelinesResponse, error)
}

type pipelinesClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelinesClient(cc grpc.ClientConnInterface) PipelinesClient {
	return &pipelinesClient{cc}
}

func (c *pipelinesClient) ManualBuild(ctx context.Context, in *ManualBuildRequest, opts ...grpc.CallOption) (*ManualBuildResponse, error) {
	out := new(ManualBuildResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/ManualBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) LaunchSatellite(ctx context.Context, in *LaunchSatelliteRequest, opts ...grpc.CallOption) (*LaunchSatelliteResponse, error) {
	out := new(LaunchSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/LaunchSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) ListSatellites(ctx context.Context, in *ListSatellitesRequest, opts ...grpc.CallOption) (*ListSatellitesResponse, error) {
	out := new(ListSatellitesResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/ListSatellites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) UpdateSatellite(ctx context.Context, in *UpdateSatelliteRequest, opts ...grpc.CallOption) (*UpdateSatelliteResponse, error) {
	out := new(UpdateSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/UpdateSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) DeleteSatellite(ctx context.Context, in *DeleteSatelliteRequest, opts ...grpc.CallOption) (*DeleteSatelliteResponse, error) {
	out := new(DeleteSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/DeleteSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) GetSatellite(ctx context.Context, in *GetSatelliteRequest, opts ...grpc.CallOption) (*GetSatelliteResponse, error) {
	out := new(GetSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/GetSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) WakeSatellite(ctx context.Context, in *WakeSatelliteRequest, opts ...grpc.CallOption) (*WakeSatelliteResponse, error) {
	out := new(WakeSatelliteResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/WakeSatellite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) ListRemoteRepos(ctx context.Context, in *ListRemoteReposRequest, opts ...grpc.CallOption) (*ListRemoteReposResponse, error) {
	out := new(ListRemoteReposResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/ListRemoteRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) AddProjectRepos(ctx context.Context, in *AddProjectReposRequest, opts ...grpc.CallOption) (*AddProjectReposResponse, error) {
	out := new(AddProjectReposResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/AddProjectRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) RemoveProjectRepo(ctx context.Context, in *RemoveProjectRepoRequest, opts ...grpc.CallOption) (*RemoveProjectRepoResponse, error) {
	out := new(RemoveProjectRepoResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/RemoveProjectRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) ListProjectRepos(ctx context.Context, in *ListProjectReposRequest, opts ...grpc.CallOption) (*ListProjectReposResponse, error) {
	out := new(ListProjectReposResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/ListProjectRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelinesClient) ListRemotePipelines(ctx context.Context, in *ListRemotePipelinesRequest, opts ...grpc.CallOption) (*ListRemotePipelinesResponse, error) {
	out := new(ListRemotePipelinesResponse)
	err := c.cc.Invoke(ctx, "/api.public.pipelines.Pipelines/ListRemotePipelines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelinesServer is the server API for Pipelines service.
// All implementations must embed UnimplementedPipelinesServer
// for forward compatibility
type PipelinesServer interface {
	// ManualBuild manually triggers a Pipeline build with the
	// provided Targets, Args and Secrets. The new build_id is returned.
	ManualBuild(context.Context, *ManualBuildRequest) (*ManualBuildResponse, error)
	// LaunchSatellite starts a new Earthly Satellite instance on the latest version.
	// The instance can be used by users to build their local, Cloud, or 3rd-party CI builds.
	LaunchSatellite(context.Context, *LaunchSatelliteRequest) (*LaunchSatelliteResponse, error)
	// ListSatellites returns a list of Earthly Satellite instances available in the organization.
	ListSatellites(context.Context, *ListSatellitesRequest) (*ListSatellitesResponse, error)
	// UpdateSatellite updates a Satellite instance to the latest version.
	// (I.e. the latest AMI we have in AWS).
	// Calling this may result in some down-time on the instance while it updates.
	UpdateSatellite(context.Context, *UpdateSatelliteRequest) (*UpdateSatelliteResponse, error)
	// DeleteSatellite decommissions a Satellite instance.
	DeleteSatellite(context.Context, *DeleteSatelliteRequest) (*DeleteSatelliteResponse, error)
	// GetSatellite retrieves the details of a particular Satellite instance.
	// Mainly intended for use by Buildkit Proxy when establishing a new connection to an instance.
	GetSatellite(context.Context, *GetSatelliteRequest) (*GetSatelliteResponse, error)
	WakeSatellite(context.Context, *WakeSatelliteRequest) (*WakeSatelliteResponse, error)
	// ListRemoteRepos uses the GitHub API to list remote repositories.
	ListRemoteRepos(context.Context, *ListRemoteReposRequest) (*ListRemoteReposResponse, error)
	// AddProjectRepos adds one or more repositories to a project.
	AddProjectRepos(context.Context, *AddProjectReposRequest) (*AddProjectReposResponse, error)
	// RemoveProjectRepo removes a repository from a project.
	RemoveProjectRepo(context.Context, *RemoveProjectRepoRequest) (*RemoveProjectRepoResponse, error)
	// ListProjectRespos lists all project repositories.
	ListProjectRepos(context.Context, *ListProjectReposRequest) (*ListProjectReposResponse, error)
	// ListRemotePipelines uses the GitHub API to list pipeline definitions present in a remote repository.
	ListRemotePipelines(context.Context, *ListRemotePipelinesRequest) (*ListRemotePipelinesResponse, error)
	mustEmbedUnimplementedPipelinesServer()
}

// UnimplementedPipelinesServer must be embedded to have forward compatible implementations.
type UnimplementedPipelinesServer struct {
}

func (UnimplementedPipelinesServer) ManualBuild(context.Context, *ManualBuildRequest) (*ManualBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManualBuild not implemented")
}
func (UnimplementedPipelinesServer) LaunchSatellite(context.Context, *LaunchSatelliteRequest) (*LaunchSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchSatellite not implemented")
}
func (UnimplementedPipelinesServer) ListSatellites(context.Context, *ListSatellitesRequest) (*ListSatellitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSatellites not implemented")
}
func (UnimplementedPipelinesServer) UpdateSatellite(context.Context, *UpdateSatelliteRequest) (*UpdateSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSatellite not implemented")
}
func (UnimplementedPipelinesServer) DeleteSatellite(context.Context, *DeleteSatelliteRequest) (*DeleteSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSatellite not implemented")
}
func (UnimplementedPipelinesServer) GetSatellite(context.Context, *GetSatelliteRequest) (*GetSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatellite not implemented")
}
func (UnimplementedPipelinesServer) WakeSatellite(context.Context, *WakeSatelliteRequest) (*WakeSatelliteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WakeSatellite not implemented")
}
func (UnimplementedPipelinesServer) ListRemoteRepos(context.Context, *ListRemoteReposRequest) (*ListRemoteReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemoteRepos not implemented")
}
func (UnimplementedPipelinesServer) AddProjectRepos(context.Context, *AddProjectReposRequest) (*AddProjectReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProjectRepos not implemented")
}
func (UnimplementedPipelinesServer) RemoveProjectRepo(context.Context, *RemoveProjectRepoRequest) (*RemoveProjectRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProjectRepo not implemented")
}
func (UnimplementedPipelinesServer) ListProjectRepos(context.Context, *ListProjectReposRequest) (*ListProjectReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectRepos not implemented")
}
func (UnimplementedPipelinesServer) ListRemotePipelines(context.Context, *ListRemotePipelinesRequest) (*ListRemotePipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemotePipelines not implemented")
}
func (UnimplementedPipelinesServer) mustEmbedUnimplementedPipelinesServer() {}

// UnsafePipelinesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelinesServer will
// result in compilation errors.
type UnsafePipelinesServer interface {
	mustEmbedUnimplementedPipelinesServer()
}

func RegisterPipelinesServer(s grpc.ServiceRegistrar, srv PipelinesServer) {
	s.RegisterService(&Pipelines_ServiceDesc, srv)
}

func _Pipelines_ManualBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManualBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).ManualBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/ManualBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).ManualBuild(ctx, req.(*ManualBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_LaunchSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).LaunchSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/LaunchSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).LaunchSatellite(ctx, req.(*LaunchSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_ListSatellites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSatellitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).ListSatellites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/ListSatellites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).ListSatellites(ctx, req.(*ListSatellitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_UpdateSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).UpdateSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/UpdateSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).UpdateSatellite(ctx, req.(*UpdateSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_DeleteSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).DeleteSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/DeleteSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).DeleteSatellite(ctx, req.(*DeleteSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_GetSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).GetSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/GetSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).GetSatellite(ctx, req.(*GetSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_WakeSatellite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WakeSatelliteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).WakeSatellite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/WakeSatellite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).WakeSatellite(ctx, req.(*WakeSatelliteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_ListRemoteRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRemoteReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).ListRemoteRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/ListRemoteRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).ListRemoteRepos(ctx, req.(*ListRemoteReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_AddProjectRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProjectReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).AddProjectRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/AddProjectRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).AddProjectRepos(ctx, req.(*AddProjectReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_RemoveProjectRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProjectRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).RemoveProjectRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/RemoveProjectRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).RemoveProjectRepo(ctx, req.(*RemoveProjectRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_ListProjectRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).ListProjectRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/ListProjectRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).ListProjectRepos(ctx, req.(*ListProjectReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipelines_ListRemotePipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRemotePipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelinesServer).ListRemotePipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.public.pipelines.Pipelines/ListRemotePipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelinesServer).ListRemotePipelines(ctx, req.(*ListRemotePipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pipelines_ServiceDesc is the grpc.ServiceDesc for Pipelines service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pipelines_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.public.pipelines.Pipelines",
	HandlerType: (*PipelinesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManualBuild",
			Handler:    _Pipelines_ManualBuild_Handler,
		},
		{
			MethodName: "LaunchSatellite",
			Handler:    _Pipelines_LaunchSatellite_Handler,
		},
		{
			MethodName: "ListSatellites",
			Handler:    _Pipelines_ListSatellites_Handler,
		},
		{
			MethodName: "UpdateSatellite",
			Handler:    _Pipelines_UpdateSatellite_Handler,
		},
		{
			MethodName: "DeleteSatellite",
			Handler:    _Pipelines_DeleteSatellite_Handler,
		},
		{
			MethodName: "GetSatellite",
			Handler:    _Pipelines_GetSatellite_Handler,
		},
		{
			MethodName: "WakeSatellite",
			Handler:    _Pipelines_WakeSatellite_Handler,
		},
		{
			MethodName: "ListRemoteRepos",
			Handler:    _Pipelines_ListRemoteRepos_Handler,
		},
		{
			MethodName: "AddProjectRepos",
			Handler:    _Pipelines_AddProjectRepos_Handler,
		},
		{
			MethodName: "RemoveProjectRepo",
			Handler:    _Pipelines_RemoveProjectRepo_Handler,
		},
		{
			MethodName: "ListProjectRepos",
			Handler:    _Pipelines_ListProjectRepos_Handler,
		},
		{
			MethodName: "ListRemotePipelines",
			Handler:    _Pipelines_ListRemotePipelines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipelines.proto",
}
