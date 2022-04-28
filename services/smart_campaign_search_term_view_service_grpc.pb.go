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

// SmartCampaignSearchTermViewServiceClient is the client API for SmartCampaignSearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartCampaignSearchTermViewServiceClient interface {
	// Returns the attributes of the requested Smart campaign search term view.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetSmartCampaignSearchTermView(ctx context.Context, in *GetSmartCampaignSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SmartCampaignSearchTermView, error)
}

type smartCampaignSearchTermViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartCampaignSearchTermViewServiceClient(cc grpc.ClientConnInterface) SmartCampaignSearchTermViewServiceClient {
	return &smartCampaignSearchTermViewServiceClient{cc}
}

func (c *smartCampaignSearchTermViewServiceClient) GetSmartCampaignSearchTermView(ctx context.Context, in *GetSmartCampaignSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SmartCampaignSearchTermView, error) {
	out := new(resources.SmartCampaignSearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.SmartCampaignSearchTermViewService/GetSmartCampaignSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartCampaignSearchTermViewServiceServer is the server API for SmartCampaignSearchTermViewService service.
// All implementations must embed UnimplementedSmartCampaignSearchTermViewServiceServer
// for forward compatibility
type SmartCampaignSearchTermViewServiceServer interface {
	// Returns the attributes of the requested Smart campaign search term view.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetSmartCampaignSearchTermView(context.Context, *GetSmartCampaignSearchTermViewRequest) (*resources.SmartCampaignSearchTermView, error)
	mustEmbedUnimplementedSmartCampaignSearchTermViewServiceServer()
}

// UnimplementedSmartCampaignSearchTermViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmartCampaignSearchTermViewServiceServer struct {
}

func (UnimplementedSmartCampaignSearchTermViewServiceServer) GetSmartCampaignSearchTermView(context.Context, *GetSmartCampaignSearchTermViewRequest) (*resources.SmartCampaignSearchTermView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmartCampaignSearchTermView not implemented")
}
func (UnimplementedSmartCampaignSearchTermViewServiceServer) mustEmbedUnimplementedSmartCampaignSearchTermViewServiceServer() {
}

// UnsafeSmartCampaignSearchTermViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartCampaignSearchTermViewServiceServer will
// result in compilation errors.
type UnsafeSmartCampaignSearchTermViewServiceServer interface {
	mustEmbedUnimplementedSmartCampaignSearchTermViewServiceServer()
}

func RegisterSmartCampaignSearchTermViewServiceServer(s grpc.ServiceRegistrar, srv SmartCampaignSearchTermViewServiceServer) {
	s.RegisterService(&SmartCampaignSearchTermViewService_ServiceDesc, srv)
}

func _SmartCampaignSearchTermViewService_GetSmartCampaignSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSmartCampaignSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSearchTermViewServiceServer).GetSmartCampaignSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.SmartCampaignSearchTermViewService/GetSmartCampaignSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSearchTermViewServiceServer).GetSmartCampaignSearchTermView(ctx, req.(*GetSmartCampaignSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartCampaignSearchTermViewService_ServiceDesc is the grpc.ServiceDesc for SmartCampaignSearchTermViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartCampaignSearchTermViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.SmartCampaignSearchTermViewService",
	HandlerType: (*SmartCampaignSearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSmartCampaignSearchTermView",
			Handler:    _SmartCampaignSearchTermViewService_GetSmartCampaignSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/smart_campaign_search_term_view_service.proto",
}
