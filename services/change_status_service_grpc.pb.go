// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/felicson/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChangeStatusServiceClient is the client API for ChangeStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChangeStatusServiceClient interface {
	// Returns the requested change status in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetChangeStatus(ctx context.Context, in *GetChangeStatusRequest, opts ...grpc.CallOption) (*resources.ChangeStatus, error)
}

type changeStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChangeStatusServiceClient(cc grpc.ClientConnInterface) ChangeStatusServiceClient {
	return &changeStatusServiceClient{cc}
}

func (c *changeStatusServiceClient) GetChangeStatus(ctx context.Context, in *GetChangeStatusRequest, opts ...grpc.CallOption) (*resources.ChangeStatus, error) {
	out := new(resources.ChangeStatus)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ChangeStatusService/GetChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeStatusServiceServer is the server API for ChangeStatusService service.
// All implementations must embed UnimplementedChangeStatusServiceServer
// for forward compatibility
type ChangeStatusServiceServer interface {
	// Returns the requested change status in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetChangeStatus(context.Context, *GetChangeStatusRequest) (*resources.ChangeStatus, error)
	mustEmbedUnimplementedChangeStatusServiceServer()
}

// UnimplementedChangeStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChangeStatusServiceServer struct {
}

func (UnimplementedChangeStatusServiceServer) GetChangeStatus(context.Context, *GetChangeStatusRequest) (*resources.ChangeStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangeStatus not implemented")
}
func (UnimplementedChangeStatusServiceServer) mustEmbedUnimplementedChangeStatusServiceServer() {}

// UnsafeChangeStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChangeStatusServiceServer will
// result in compilation errors.
type UnsafeChangeStatusServiceServer interface {
	mustEmbedUnimplementedChangeStatusServiceServer()
}

func RegisterChangeStatusServiceServer(s grpc.ServiceRegistrar, srv ChangeStatusServiceServer) {
	s.RegisterService(&ChangeStatusService_ServiceDesc, srv)
}

func _ChangeStatusService_GetChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeStatusServiceServer).GetChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ChangeStatusService/GetChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeStatusServiceServer).GetChangeStatus(ctx, req.(*GetChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChangeStatusService_ServiceDesc is the grpc.ServiceDesc for ChangeStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChangeStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ChangeStatusService",
	HandlerType: (*ChangeStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChangeStatus",
			Handler:    _ChangeStatusService_GetChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/change_status_service.proto",
}
