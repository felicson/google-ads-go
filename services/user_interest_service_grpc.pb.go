// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/ercling/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInterestServiceClient(cc grpc.ClientConnInterface) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
// All implementations must embed UnimplementedUserInterestServiceServer
// for forward compatibility
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
	mustEmbedUnimplementedUserInterestServiceServer()
}

// UnimplementedUserInterestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserInterestServiceServer struct {
}

func (UnimplementedUserInterestServiceServer) GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInterest not implemented")
}
func (UnimplementedUserInterestServiceServer) mustEmbedUnimplementedUserInterestServiceServer() {}

// UnsafeUserInterestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInterestServiceServer will
// result in compilation errors.
type UnsafeUserInterestServiceServer interface {
	mustEmbedUnimplementedUserInterestServiceServer()
}

func RegisterUserInterestServiceServer(s grpc.ServiceRegistrar, srv UserInterestServiceServer) {
	s.RegisterService(&UserInterestService_ServiceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInterestService_ServiceDesc is the grpc.ServiceDesc for UserInterestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInterestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/user_interest_service.proto",
}
