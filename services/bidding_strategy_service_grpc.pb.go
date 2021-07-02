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

// BiddingStrategyServiceClient is the client API for BiddingStrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingStrategyServiceClient interface {
	// Returns the requested bidding strategy in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AdxError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BiddingError]()
	//   [BiddingStrategyError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error)
}

type biddingStrategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingStrategyServiceClient(cc grpc.ClientConnInterface) BiddingStrategyServiceClient {
	return &biddingStrategyServiceClient{cc}
}

func (c *biddingStrategyServiceClient) GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error) {
	out := new(resources.BiddingStrategy)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.BiddingStrategyService/GetBiddingStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biddingStrategyServiceClient) MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error) {
	out := new(MutateBiddingStrategiesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.BiddingStrategyService/MutateBiddingStrategies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingStrategyServiceServer is the server API for BiddingStrategyService service.
// All implementations must embed UnimplementedBiddingStrategyServiceServer
// for forward compatibility
type BiddingStrategyServiceServer interface {
	// Returns the requested bidding strategy in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetBiddingStrategy(context.Context, *GetBiddingStrategyRequest) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AdxError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BiddingError]()
	//   [BiddingStrategyError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperationAccessDeniedError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error)
	mustEmbedUnimplementedBiddingStrategyServiceServer()
}

// UnimplementedBiddingStrategyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingStrategyServiceServer struct {
}

func (UnimplementedBiddingStrategyServiceServer) GetBiddingStrategy(context.Context, *GetBiddingStrategyRequest) (*resources.BiddingStrategy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBiddingStrategy not implemented")
}
func (UnimplementedBiddingStrategyServiceServer) MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBiddingStrategies not implemented")
}
func (UnimplementedBiddingStrategyServiceServer) mustEmbedUnimplementedBiddingStrategyServiceServer() {
}

// UnsafeBiddingStrategyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingStrategyServiceServer will
// result in compilation errors.
type UnsafeBiddingStrategyServiceServer interface {
	mustEmbedUnimplementedBiddingStrategyServiceServer()
}

func RegisterBiddingStrategyServiceServer(s grpc.ServiceRegistrar, srv BiddingStrategyServiceServer) {
	s.RegisterService(&BiddingStrategyService_ServiceDesc, srv)
}

func _BiddingStrategyService_GetBiddingStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBiddingStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.BiddingStrategyService/GetBiddingStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, req.(*GetBiddingStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiddingStrategyService_MutateBiddingStrategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.BiddingStrategyService/MutateBiddingStrategies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, req.(*MutateBiddingStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingStrategyService_ServiceDesc is the grpc.ServiceDesc for BiddingStrategyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingStrategyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.BiddingStrategyService",
	HandlerType: (*BiddingStrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBiddingStrategy",
			Handler:    _BiddingStrategyService_GetBiddingStrategy_Handler,
		},
		{
			MethodName: "MutateBiddingStrategies",
			Handler:    _BiddingStrategyService_MutateBiddingStrategies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/bidding_strategy_service.proto",
}
