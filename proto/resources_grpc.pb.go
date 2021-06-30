// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// KISSClient is the client API for KISS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KISSClient interface {
	// Temporary RPC to test authorization; will be removed.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error)
	BindSecret(ctx context.Context, in *BindSecretRequest, opts ...grpc.CallOption) (*BindSecretResponse, error)
	CreateSecretIAMPolicy(ctx context.Context, in *CreateSecretIAMPolicyRequest, opts ...grpc.CallOption) (*CreateSecretIAMPolicyResponse, error)
	DeleteSecretIAMPolicy(ctx context.Context, in *DeleteSecretIAMPolicyRequest, opts ...grpc.CallOption) (*DeleteSecretIAMPolicyResponse, error)
}

type kISSClient struct {
	cc grpc.ClientConnInterface
}

func NewKISSClient(cc grpc.ClientConnInterface) KISSClient {
	return &kISSClient{cc}
}

func (c *kISSClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kISSClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error) {
	out := new(CreateSecretResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kISSClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error) {
	out := new(DeleteSecretResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kISSClient) BindSecret(ctx context.Context, in *BindSecretRequest, opts ...grpc.CallOption) (*BindSecretResponse, error) {
	out := new(BindSecretResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/BindSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kISSClient) CreateSecretIAMPolicy(ctx context.Context, in *CreateSecretIAMPolicyRequest, opts ...grpc.CallOption) (*CreateSecretIAMPolicyResponse, error) {
	out := new(CreateSecretIAMPolicyResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/CreateSecretIAMPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kISSClient) DeleteSecretIAMPolicy(ctx context.Context, in *DeleteSecretIAMPolicyRequest, opts ...grpc.CallOption) (*DeleteSecretIAMPolicyResponse, error) {
	out := new(DeleteSecretIAMPolicyResponse)
	err := c.cc.Invoke(ctx, "/kiss.resources.KISS/DeleteSecretIAMPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KISSServer is the server API for KISS service.
// All implementations must embed UnimplementedKISSServer
// for forward compatibility
type KISSServer interface {
	// Temporary RPC to test authorization; will be removed.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error)
	BindSecret(context.Context, *BindSecretRequest) (*BindSecretResponse, error)
	CreateSecretIAMPolicy(context.Context, *CreateSecretIAMPolicyRequest) (*CreateSecretIAMPolicyResponse, error)
	DeleteSecretIAMPolicy(context.Context, *DeleteSecretIAMPolicyRequest) (*DeleteSecretIAMPolicyResponse, error)
	mustEmbedUnimplementedKISSServer()
}

// UnimplementedKISSServer must be embedded to have forward compatible implementations.
type UnimplementedKISSServer struct {
}

func (UnimplementedKISSServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKISSServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedKISSServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedKISSServer) BindSecret(context.Context, *BindSecretRequest) (*BindSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindSecret not implemented")
}
func (UnimplementedKISSServer) CreateSecretIAMPolicy(context.Context, *CreateSecretIAMPolicyRequest) (*CreateSecretIAMPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecretIAMPolicy not implemented")
}
func (UnimplementedKISSServer) DeleteSecretIAMPolicy(context.Context, *DeleteSecretIAMPolicyRequest) (*DeleteSecretIAMPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecretIAMPolicy not implemented")
}
func (UnimplementedKISSServer) mustEmbedUnimplementedKISSServer() {}

// UnsafeKISSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KISSServer will
// result in compilation errors.
type UnsafeKISSServer interface {
	mustEmbedUnimplementedKISSServer()
}

func RegisterKISSServer(s grpc.ServiceRegistrar, srv KISSServer) {
	s.RegisterService(&KISS_ServiceDesc, srv)
}

func _KISS_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KISS_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KISS_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KISS_BindSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).BindSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/BindSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).BindSecret(ctx, req.(*BindSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KISS_CreateSecretIAMPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretIAMPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).CreateSecretIAMPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/CreateSecretIAMPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).CreateSecretIAMPolicy(ctx, req.(*CreateSecretIAMPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KISS_DeleteSecretIAMPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretIAMPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KISSServer).DeleteSecretIAMPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kiss.resources.KISS/DeleteSecretIAMPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KISSServer).DeleteSecretIAMPolicy(ctx, req.(*DeleteSecretIAMPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KISS_ServiceDesc is the grpc.ServiceDesc for KISS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KISS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kiss.resources.KISS",
	HandlerType: (*KISSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _KISS_Ping_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _KISS_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _KISS_DeleteSecret_Handler,
		},
		{
			MethodName: "BindSecret",
			Handler:    _KISS_BindSecret_Handler,
		},
		{
			MethodName: "CreateSecretIAMPolicy",
			Handler:    _KISS_CreateSecretIAMPolicy_Handler,
		},
		{
			MethodName: "DeleteSecretIAMPolicy",
			Handler:    _KISS_DeleteSecretIAMPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources.proto",
}
