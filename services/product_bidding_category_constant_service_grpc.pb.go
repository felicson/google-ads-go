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

// ProductBiddingCategoryConstantServiceClient is the client API for ProductBiddingCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductBiddingCategoryConstantServiceClient interface {
	// Returns the requested Product Bidding Category in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error)
}

type productBiddingCategoryConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductBiddingCategoryConstantServiceClient(cc grpc.ClientConnInterface) ProductBiddingCategoryConstantServiceClient {
	return &productBiddingCategoryConstantServiceClient{cc}
}

func (c *productBiddingCategoryConstantServiceClient) GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error) {
	out := new(resources.ProductBiddingCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductBiddingCategoryConstantServiceServer is the server API for ProductBiddingCategoryConstantService service.
// All implementations must embed UnimplementedProductBiddingCategoryConstantServiceServer
// for forward compatibility
type ProductBiddingCategoryConstantServiceServer interface {
	// Returns the requested Product Bidding Category in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetProductBiddingCategoryConstant(context.Context, *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error)
	mustEmbedUnimplementedProductBiddingCategoryConstantServiceServer()
}

// UnimplementedProductBiddingCategoryConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductBiddingCategoryConstantServiceServer struct {
}

func (UnimplementedProductBiddingCategoryConstantServiceServer) GetProductBiddingCategoryConstant(context.Context, *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBiddingCategoryConstant not implemented")
}
func (UnimplementedProductBiddingCategoryConstantServiceServer) mustEmbedUnimplementedProductBiddingCategoryConstantServiceServer() {
}

// UnsafeProductBiddingCategoryConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductBiddingCategoryConstantServiceServer will
// result in compilation errors.
type UnsafeProductBiddingCategoryConstantServiceServer interface {
	mustEmbedUnimplementedProductBiddingCategoryConstantServiceServer()
}

func RegisterProductBiddingCategoryConstantServiceServer(s grpc.ServiceRegistrar, srv ProductBiddingCategoryConstantServiceServer) {
	s.RegisterService(&ProductBiddingCategoryConstantService_ServiceDesc, srv)
}

func _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBiddingCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, req.(*GetProductBiddingCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductBiddingCategoryConstantService_ServiceDesc is the grpc.ServiceDesc for ProductBiddingCategoryConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductBiddingCategoryConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ProductBiddingCategoryConstantService",
	HandlerType: (*ProductBiddingCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductBiddingCategoryConstant",
			Handler:    _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/product_bidding_category_constant_service.proto",
}
