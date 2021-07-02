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

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClientLinkServiceClient interface {
	// Returns the requested CustomerClientLink in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [ManagerLinkError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error)
}

type customerClientLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientLinkServiceClient(cc grpc.ClientConnInterface) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error) {
	out := new(resources.CustomerClientLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.CustomerClientLinkService/GetCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClientLinkServiceClient) MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error) {
	out := new(MutateCustomerClientLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.CustomerClientLinkService/MutateCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
// All implementations must embed UnimplementedCustomerClientLinkServiceServer
// for forward compatibility
type CustomerClientLinkServiceServer interface {
	// Returns the requested CustomerClientLink in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [ManagerLinkError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error)
	mustEmbedUnimplementedCustomerClientLinkServiceServer()
}

// UnimplementedCustomerClientLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerClientLinkServiceServer struct {
}

func (UnimplementedCustomerClientLinkServiceServer) GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClientLink not implemented")
}
func (UnimplementedCustomerClientLinkServiceServer) MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerClientLink not implemented")
}
func (UnimplementedCustomerClientLinkServiceServer) mustEmbedUnimplementedCustomerClientLinkServiceServer() {
}

// UnsafeCustomerClientLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerClientLinkServiceServer will
// result in compilation errors.
type UnsafeCustomerClientLinkServiceServer interface {
	mustEmbedUnimplementedCustomerClientLinkServiceServer()
}

func RegisterCustomerClientLinkServiceServer(s grpc.ServiceRegistrar, srv CustomerClientLinkServiceServer) {
	s.RegisterService(&CustomerClientLinkService_ServiceDesc, srv)
}

func _CustomerClientLinkService_GetCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.CustomerClientLinkService/GetCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, req.(*GetCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerClientLinkService_MutateCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.CustomerClientLinkService/MutateCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, req.(*MutateCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerClientLinkService_ServiceDesc is the grpc.ServiceDesc for CustomerClientLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerClientLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClientLink",
			Handler:    _CustomerClientLinkService_GetCustomerClientLink_Handler,
		},
		{
			MethodName: "MutateCustomerClientLink",
			Handler:    _CustomerClientLinkService_MutateCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/customer_client_link_service.proto",
}
