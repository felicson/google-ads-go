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

// CampaignAudienceViewServiceClient is the client API for CampaignAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignAudienceViewServiceClient interface {
	// Returns the requested campaign audience view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error)
}

type campaignAudienceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAudienceViewServiceClient(cc grpc.ClientConnInterface) CampaignAudienceViewServiceClient {
	return &campaignAudienceViewServiceClient{cc}
}

func (c *campaignAudienceViewServiceClient) GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error) {
	out := new(resources.CampaignAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAudienceViewServiceServer is the server API for CampaignAudienceViewService service.
// All implementations must embed UnimplementedCampaignAudienceViewServiceServer
// for forward compatibility
type CampaignAudienceViewServiceServer interface {
	// Returns the requested campaign audience view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error)
	mustEmbedUnimplementedCampaignAudienceViewServiceServer()
}

// UnimplementedCampaignAudienceViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignAudienceViewServiceServer struct {
}

func (UnimplementedCampaignAudienceViewServiceServer) GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignAudienceView not implemented")
}
func (UnimplementedCampaignAudienceViewServiceServer) mustEmbedUnimplementedCampaignAudienceViewServiceServer() {
}

// UnsafeCampaignAudienceViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignAudienceViewServiceServer will
// result in compilation errors.
type UnsafeCampaignAudienceViewServiceServer interface {
	mustEmbedUnimplementedCampaignAudienceViewServiceServer()
}

func RegisterCampaignAudienceViewServiceServer(s grpc.ServiceRegistrar, srv CampaignAudienceViewServiceServer) {
	s.RegisterService(&CampaignAudienceViewService_ServiceDesc, srv)
}

func _CampaignAudienceViewService_GetCampaignAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignAudienceViewService_ServiceDesc is the grpc.ServiceDesc for CampaignAudienceViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignAudienceViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/campaign_audience_view_service.proto",
}
