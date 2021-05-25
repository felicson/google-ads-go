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

// GeographicViewServiceClient is the client API for GeographicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeographicViewServiceClient interface {
	// Returns the requested geographic view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error)
}

type geographicViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographicViewServiceClient(cc grpc.ClientConnInterface) GeographicViewServiceClient {
	return &geographicViewServiceClient{cc}
}

func (c *geographicViewServiceClient) GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error) {
	out := new(resources.GeographicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.GeographicViewService/GetGeographicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographicViewServiceServer is the server API for GeographicViewService service.
// All implementations must embed UnimplementedGeographicViewServiceServer
// for forward compatibility
type GeographicViewServiceServer interface {
	// Returns the requested geographic view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGeographicView(context.Context, *GetGeographicViewRequest) (*resources.GeographicView, error)
	mustEmbedUnimplementedGeographicViewServiceServer()
}

// UnimplementedGeographicViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeographicViewServiceServer struct {
}

func (UnimplementedGeographicViewServiceServer) GetGeographicView(context.Context, *GetGeographicViewRequest) (*resources.GeographicView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeographicView not implemented")
}
func (UnimplementedGeographicViewServiceServer) mustEmbedUnimplementedGeographicViewServiceServer() {}

// UnsafeGeographicViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeographicViewServiceServer will
// result in compilation errors.
type UnsafeGeographicViewServiceServer interface {
	mustEmbedUnimplementedGeographicViewServiceServer()
}

func RegisterGeographicViewServiceServer(s grpc.ServiceRegistrar, srv GeographicViewServiceServer) {
	s.RegisterService(&GeographicViewService_ServiceDesc, srv)
}

func _GeographicViewService_GetGeographicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeographicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.GeographicViewService/GetGeographicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, req.(*GetGeographicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeographicViewService_ServiceDesc is the grpc.ServiceDesc for GeographicViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeographicViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.GeographicViewService",
	HandlerType: (*GeographicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeographicView",
			Handler:    _GeographicViewService_GetGeographicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/geographic_view_service.proto",
}
