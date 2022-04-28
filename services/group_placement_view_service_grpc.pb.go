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

// GroupPlacementViewServiceClient is the client API for GroupPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupPlacementViewServiceClient interface {
	// Returns the requested Group Placement view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error)
}

type groupPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupPlacementViewServiceClient(cc grpc.ClientConnInterface) GroupPlacementViewServiceClient {
	return &groupPlacementViewServiceClient{cc}
}

func (c *groupPlacementViewServiceClient) GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error) {
	out := new(resources.GroupPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.GroupPlacementViewService/GetGroupPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupPlacementViewServiceServer is the server API for GroupPlacementViewService service.
// All implementations must embed UnimplementedGroupPlacementViewServiceServer
// for forward compatibility
type GroupPlacementViewServiceServer interface {
	// Returns the requested Group Placement view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error)
	mustEmbedUnimplementedGroupPlacementViewServiceServer()
}

// UnimplementedGroupPlacementViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupPlacementViewServiceServer struct {
}

func (UnimplementedGroupPlacementViewServiceServer) GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPlacementView not implemented")
}
func (UnimplementedGroupPlacementViewServiceServer) mustEmbedUnimplementedGroupPlacementViewServiceServer() {
}

// UnsafeGroupPlacementViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupPlacementViewServiceServer will
// result in compilation errors.
type UnsafeGroupPlacementViewServiceServer interface {
	mustEmbedUnimplementedGroupPlacementViewServiceServer()
}

func RegisterGroupPlacementViewServiceServer(s grpc.ServiceRegistrar, srv GroupPlacementViewServiceServer) {
	s.RegisterService(&GroupPlacementViewService_ServiceDesc, srv)
}

func _GroupPlacementViewService_GetGroupPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.GroupPlacementViewService/GetGroupPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, req.(*GetGroupPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupPlacementViewService_ServiceDesc is the grpc.ServiceDesc for GroupPlacementViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupPlacementViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.GroupPlacementViewService",
	HandlerType: (*GroupPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupPlacementView",
			Handler:    _GroupPlacementViewService_GetGroupPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/group_placement_view_service.proto",
}
