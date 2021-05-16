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

// CasbinClient is the client API for Casbin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CasbinClient interface {
	HasPermissionForUserInDomain(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*BoolReply, error)
}

type casbinClient struct {
	cc grpc.ClientConnInterface
}

func NewCasbinClient(cc grpc.ClientConnInterface) CasbinClient {
	return &casbinClient{cc}
}

func (c *casbinClient) HasPermissionForUserInDomain(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/proto.Casbin/HasPermissionForUserInDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CasbinServer is the server API for Casbin service.
// All implementations should embed UnimplementedCasbinServer
// for forward compatibility
type CasbinServer interface {
	HasPermissionForUserInDomain(context.Context, *PermissionRequest) (*BoolReply, error)
}

// UnimplementedCasbinServer should be embedded to have forward compatible implementations.
type UnimplementedCasbinServer struct {
}

func (UnimplementedCasbinServer) HasPermissionForUserInDomain(context.Context, *PermissionRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasPermissionForUserInDomain not implemented")
}

// UnsafeCasbinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CasbinServer will
// result in compilation errors.
type UnsafeCasbinServer interface {
	mustEmbedUnimplementedCasbinServer()
}

func RegisterCasbinServer(s grpc.ServiceRegistrar, srv CasbinServer) {
	s.RegisterService(&Casbin_ServiceDesc, srv)
}

func _Casbin_HasPermissionForUserInDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CasbinServer).HasPermissionForUserInDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Casbin/HasPermissionForUserInDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CasbinServer).HasPermissionForUserInDomain(ctx, req.(*PermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Casbin_ServiceDesc is the grpc.ServiceDesc for Casbin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Casbin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Casbin",
	HandlerType: (*CasbinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasPermissionForUserInDomain",
			Handler:    _Casbin_HasPermissionForUserInDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/casbin.proto",
}
