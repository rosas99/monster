// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/usercenter/v1/usercenter.proto

package v1

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
	UserCenter_Authorize_FullMethodName = "/usercenter.v1.UserCenter/Authorize"
)

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	Authorize(ctx context.Context, in *AuthzRequest, opts ...grpc.CallOption) (*AuthzResponse, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) Authorize(ctx context.Context, in *AuthzRequest, opts ...grpc.CallOption) (*AuthzResponse, error) {
	out := new(AuthzResponse)
	err := c.cc.Invoke(ctx, UserCenter_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	Authorize(context.Context, *AuthzRequest) (*AuthzResponse, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) Authorize(context.Context, *AuthzRequest) (*AuthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Authorize(ctx, req.(*AuthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.v1.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _UserCenter_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/usercenter/v1/usercenter.proto",
}
