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

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelServiceClient interface {
	// Returns the requested label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error) {
	out := new(resources.Label)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error) {
	out := new(MutateLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.LabelService/MutateLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
// All implementations must embed UnimplementedLabelServiceServer
// for forward compatibility
type LabelServiceServer interface {
	// Returns the requested label in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error)
	mustEmbedUnimplementedLabelServiceServer()
}

// UnimplementedLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (UnimplementedLabelServiceServer) GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}
func (UnimplementedLabelServiceServer) MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateLabels not implemented")
}
func (UnimplementedLabelServiceServer) mustEmbedUnimplementedLabelServiceServer() {}

// UnsafeLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelServiceServer will
// result in compilation errors.
type UnsafeLabelServiceServer interface {
	mustEmbedUnimplementedLabelServiceServer()
}

func RegisterLabelServiceServer(s grpc.ServiceRegistrar, srv LabelServiceServer) {
	s.RegisterService(&LabelService_ServiceDesc, srv)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_MutateLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).MutateLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.LabelService/MutateLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).MutateLabels(ctx, req.(*MutateLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelService_ServiceDesc is the grpc.ServiceDesc for LabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
		{
			MethodName: "MutateLabels",
			Handler:    _LabelService_MutateLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/label_service.proto",
}
