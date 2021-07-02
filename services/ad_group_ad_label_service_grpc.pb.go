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

// AdGroupAdLabelServiceClient is the client API for AdGroupAdLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupAdLabelServiceClient interface {
	// Returns the requested ad group ad label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error)
}

type adGroupAdLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdLabelServiceClient(cc grpc.ClientConnInterface) AdGroupAdLabelServiceClient {
	return &adGroupAdLabelServiceClient{cc}
}

func (c *adGroupAdLabelServiceClient) GetAdGroupAdLabel(ctx context.Context, in *GetAdGroupAdLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupAdLabel, error) {
	out := new(resources.AdGroupAdLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupAdLabelService/GetAdGroupAdLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdLabelServiceClient) MutateAdGroupAdLabels(ctx context.Context, in *MutateAdGroupAdLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdLabelsResponse, error) {
	out := new(MutateAdGroupAdLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupAdLabelService/MutateAdGroupAdLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdLabelServiceServer is the server API for AdGroupAdLabelService service.
// All implementations must embed UnimplementedAdGroupAdLabelServiceServer
// for forward compatibility
type AdGroupAdLabelServiceServer interface {
	// Returns the requested ad group ad label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupAdLabel(context.Context, *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error)
	// Creates and removes ad group ad labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupAdLabels(context.Context, *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error)
	mustEmbedUnimplementedAdGroupAdLabelServiceServer()
}

// UnimplementedAdGroupAdLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdLabelServiceServer struct {
}

func (UnimplementedAdGroupAdLabelServiceServer) GetAdGroupAdLabel(context.Context, *GetAdGroupAdLabelRequest) (*resources.AdGroupAdLabel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupAdLabel not implemented")
}
func (UnimplementedAdGroupAdLabelServiceServer) MutateAdGroupAdLabels(context.Context, *MutateAdGroupAdLabelsRequest) (*MutateAdGroupAdLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupAdLabels not implemented")
}
func (UnimplementedAdGroupAdLabelServiceServer) mustEmbedUnimplementedAdGroupAdLabelServiceServer() {}

// UnsafeAdGroupAdLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupAdLabelServiceServer will
// result in compilation errors.
type UnsafeAdGroupAdLabelServiceServer interface {
	mustEmbedUnimplementedAdGroupAdLabelServiceServer()
}

func RegisterAdGroupAdLabelServiceServer(s grpc.ServiceRegistrar, srv AdGroupAdLabelServiceServer) {
	s.RegisterService(&AdGroupAdLabelService_ServiceDesc, srv)
}

func _AdGroupAdLabelService_GetAdGroupAdLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupAdLabelService/GetAdGroupAdLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).GetAdGroupAdLabel(ctx, req.(*GetAdGroupAdLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupAdLabelService/MutateAdGroupAdLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdLabelServiceServer).MutateAdGroupAdLabels(ctx, req.(*MutateAdGroupAdLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupAdLabelService_ServiceDesc is the grpc.ServiceDesc for AdGroupAdLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupAdLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.AdGroupAdLabelService",
	HandlerType: (*AdGroupAdLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAdLabel",
			Handler:    _AdGroupAdLabelService_GetAdGroupAdLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupAdLabels",
			Handler:    _AdGroupAdLabelService_MutateAdGroupAdLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/ad_group_ad_label_service.proto",
}
