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

// AdGroupCriterionLabelServiceClient is the client API for AdGroupCriterionLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupCriterionLabelServiceClient interface {
	// Returns the requested ad group criterion label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error)
}

type adGroupCriterionLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionLabelServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionLabelServiceClient {
	return &adGroupCriterionLabelServiceClient{cc}
}

func (c *adGroupCriterionLabelServiceClient) GetAdGroupCriterionLabel(ctx context.Context, in *GetAdGroupCriterionLabelRequest, opts ...grpc.CallOption) (*resources.AdGroupCriterionLabel, error) {
	out := new(resources.AdGroupCriterionLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupCriterionLabelServiceClient) MutateAdGroupCriterionLabels(ctx context.Context, in *MutateAdGroupCriterionLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupCriterionLabelsResponse, error) {
	out := new(MutateAdGroupCriterionLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionLabelServiceServer is the server API for AdGroupCriterionLabelService service.
// All implementations must embed UnimplementedAdGroupCriterionLabelServiceServer
// for forward compatibility
type AdGroupCriterionLabelServiceServer interface {
	// Returns the requested ad group criterion label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdGroupCriterionLabel(context.Context, *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error)
	// Creates and removes ad group criterion labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error)
	mustEmbedUnimplementedAdGroupCriterionLabelServiceServer()
}

// UnimplementedAdGroupCriterionLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionLabelServiceServer struct {
}

func (UnimplementedAdGroupCriterionLabelServiceServer) GetAdGroupCriterionLabel(context.Context, *GetAdGroupCriterionLabelRequest) (*resources.AdGroupCriterionLabel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupCriterionLabel not implemented")
}
func (UnimplementedAdGroupCriterionLabelServiceServer) MutateAdGroupCriterionLabels(context.Context, *MutateAdGroupCriterionLabelsRequest) (*MutateAdGroupCriterionLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupCriterionLabels not implemented")
}
func (UnimplementedAdGroupCriterionLabelServiceServer) mustEmbedUnimplementedAdGroupCriterionLabelServiceServer() {
}

// UnsafeAdGroupCriterionLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupCriterionLabelServiceServer will
// result in compilation errors.
type UnsafeAdGroupCriterionLabelServiceServer interface {
	mustEmbedUnimplementedAdGroupCriterionLabelServiceServer()
}

func RegisterAdGroupCriterionLabelServiceServer(s grpc.ServiceRegistrar, srv AdGroupCriterionLabelServiceServer) {
	s.RegisterService(&AdGroupCriterionLabelService_ServiceDesc, srv)
}

func _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupCriterionLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupCriterionLabelService/GetAdGroupCriterionLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).GetAdGroupCriterionLabel(ctx, req.(*GetAdGroupCriterionLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriterionLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.AdGroupCriterionLabelService/MutateAdGroupCriterionLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionLabelServiceServer).MutateAdGroupCriterionLabels(ctx, req.(*MutateAdGroupCriterionLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupCriterionLabelService_ServiceDesc is the grpc.ServiceDesc for AdGroupCriterionLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupCriterionLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.AdGroupCriterionLabelService",
	HandlerType: (*AdGroupCriterionLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupCriterionLabel",
			Handler:    _AdGroupCriterionLabelService_GetAdGroupCriterionLabel_Handler,
		},
		{
			MethodName: "MutateAdGroupCriterionLabels",
			Handler:    _AdGroupCriterionLabelService_MutateAdGroupCriterionLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/ad_group_criterion_label_service.proto",
}
