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

// AssetFieldTypeViewServiceClient is the client API for AssetFieldTypeViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetFieldTypeViewServiceClient interface {
	// Returns the requested asset field type view in full detail.
	GetAssetFieldTypeView(ctx context.Context, in *GetAssetFieldTypeViewRequest, opts ...grpc.CallOption) (*resources.AssetFieldTypeView, error)
}

type assetFieldTypeViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetFieldTypeViewServiceClient(cc grpc.ClientConnInterface) AssetFieldTypeViewServiceClient {
	return &assetFieldTypeViewServiceClient{cc}
}

func (c *assetFieldTypeViewServiceClient) GetAssetFieldTypeView(ctx context.Context, in *GetAssetFieldTypeViewRequest, opts ...grpc.CallOption) (*resources.AssetFieldTypeView, error) {
	out := new(resources.AssetFieldTypeView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AssetFieldTypeViewService/GetAssetFieldTypeView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetFieldTypeViewServiceServer is the server API for AssetFieldTypeViewService service.
// All implementations must embed UnimplementedAssetFieldTypeViewServiceServer
// for forward compatibility
type AssetFieldTypeViewServiceServer interface {
	// Returns the requested asset field type view in full detail.
	GetAssetFieldTypeView(context.Context, *GetAssetFieldTypeViewRequest) (*resources.AssetFieldTypeView, error)
	mustEmbedUnimplementedAssetFieldTypeViewServiceServer()
}

// UnimplementedAssetFieldTypeViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetFieldTypeViewServiceServer struct {
}

func (UnimplementedAssetFieldTypeViewServiceServer) GetAssetFieldTypeView(context.Context, *GetAssetFieldTypeViewRequest) (*resources.AssetFieldTypeView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetFieldTypeView not implemented")
}
func (UnimplementedAssetFieldTypeViewServiceServer) mustEmbedUnimplementedAssetFieldTypeViewServiceServer() {
}

// UnsafeAssetFieldTypeViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetFieldTypeViewServiceServer will
// result in compilation errors.
type UnsafeAssetFieldTypeViewServiceServer interface {
	mustEmbedUnimplementedAssetFieldTypeViewServiceServer()
}

func RegisterAssetFieldTypeViewServiceServer(s grpc.ServiceRegistrar, srv AssetFieldTypeViewServiceServer) {
	s.RegisterService(&AssetFieldTypeViewService_ServiceDesc, srv)
}

func _AssetFieldTypeViewService_GetAssetFieldTypeView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetFieldTypeViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetFieldTypeViewServiceServer).GetAssetFieldTypeView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AssetFieldTypeViewService/GetAssetFieldTypeView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetFieldTypeViewServiceServer).GetAssetFieldTypeView(ctx, req.(*GetAssetFieldTypeViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetFieldTypeViewService_ServiceDesc is the grpc.ServiceDesc for AssetFieldTypeViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetFieldTypeViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AssetFieldTypeViewService",
	HandlerType: (*AssetFieldTypeViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAssetFieldTypeView",
			Handler:    _AssetFieldTypeViewService_GetAssetFieldTypeView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/asset_field_type_view_service.proto",
}
