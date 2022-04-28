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

// SmartCampaignSettingServiceClient is the client API for SmartCampaignSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartCampaignSettingServiceClient interface {
	// Returns the requested Smart campaign setting in full detail.
	GetSmartCampaignSetting(ctx context.Context, in *GetSmartCampaignSettingRequest, opts ...grpc.CallOption) (*resources.SmartCampaignSetting, error)
	// Updates Smart campaign settings for campaigns.
	MutateSmartCampaignSettings(ctx context.Context, in *MutateSmartCampaignSettingsRequest, opts ...grpc.CallOption) (*MutateSmartCampaignSettingsResponse, error)
}

type smartCampaignSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartCampaignSettingServiceClient(cc grpc.ClientConnInterface) SmartCampaignSettingServiceClient {
	return &smartCampaignSettingServiceClient{cc}
}

func (c *smartCampaignSettingServiceClient) GetSmartCampaignSetting(ctx context.Context, in *GetSmartCampaignSettingRequest, opts ...grpc.CallOption) (*resources.SmartCampaignSetting, error) {
	out := new(resources.SmartCampaignSetting)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.SmartCampaignSettingService/GetSmartCampaignSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCampaignSettingServiceClient) MutateSmartCampaignSettings(ctx context.Context, in *MutateSmartCampaignSettingsRequest, opts ...grpc.CallOption) (*MutateSmartCampaignSettingsResponse, error) {
	out := new(MutateSmartCampaignSettingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.SmartCampaignSettingService/MutateSmartCampaignSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartCampaignSettingServiceServer is the server API for SmartCampaignSettingService service.
// All implementations must embed UnimplementedSmartCampaignSettingServiceServer
// for forward compatibility
type SmartCampaignSettingServiceServer interface {
	// Returns the requested Smart campaign setting in full detail.
	GetSmartCampaignSetting(context.Context, *GetSmartCampaignSettingRequest) (*resources.SmartCampaignSetting, error)
	// Updates Smart campaign settings for campaigns.
	MutateSmartCampaignSettings(context.Context, *MutateSmartCampaignSettingsRequest) (*MutateSmartCampaignSettingsResponse, error)
	mustEmbedUnimplementedSmartCampaignSettingServiceServer()
}

// UnimplementedSmartCampaignSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmartCampaignSettingServiceServer struct {
}

func (UnimplementedSmartCampaignSettingServiceServer) GetSmartCampaignSetting(context.Context, *GetSmartCampaignSettingRequest) (*resources.SmartCampaignSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmartCampaignSetting not implemented")
}
func (UnimplementedSmartCampaignSettingServiceServer) MutateSmartCampaignSettings(context.Context, *MutateSmartCampaignSettingsRequest) (*MutateSmartCampaignSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateSmartCampaignSettings not implemented")
}
func (UnimplementedSmartCampaignSettingServiceServer) mustEmbedUnimplementedSmartCampaignSettingServiceServer() {
}

// UnsafeSmartCampaignSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartCampaignSettingServiceServer will
// result in compilation errors.
type UnsafeSmartCampaignSettingServiceServer interface {
	mustEmbedUnimplementedSmartCampaignSettingServiceServer()
}

func RegisterSmartCampaignSettingServiceServer(s grpc.ServiceRegistrar, srv SmartCampaignSettingServiceServer) {
	s.RegisterService(&SmartCampaignSettingService_ServiceDesc, srv)
}

func _SmartCampaignSettingService_GetSmartCampaignSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSmartCampaignSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSettingServiceServer).GetSmartCampaignSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.SmartCampaignSettingService/GetSmartCampaignSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSettingServiceServer).GetSmartCampaignSetting(ctx, req.(*GetSmartCampaignSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCampaignSettingService_MutateSmartCampaignSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSmartCampaignSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSettingServiceServer).MutateSmartCampaignSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.SmartCampaignSettingService/MutateSmartCampaignSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSettingServiceServer).MutateSmartCampaignSettings(ctx, req.(*MutateSmartCampaignSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartCampaignSettingService_ServiceDesc is the grpc.ServiceDesc for SmartCampaignSettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartCampaignSettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.SmartCampaignSettingService",
	HandlerType: (*SmartCampaignSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSmartCampaignSetting",
			Handler:    _SmartCampaignSettingService_GetSmartCampaignSetting_Handler,
		},
		{
			MethodName: "MutateSmartCampaignSettings",
			Handler:    _SmartCampaignSettingService_MutateSmartCampaignSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/smart_campaign_setting_service.proto",
}
