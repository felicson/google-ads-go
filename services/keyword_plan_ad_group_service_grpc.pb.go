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

// KeywordPlanAdGroupServiceClient is the client API for KeywordPlanAdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordPlanAdGroupServiceClient interface {
	// Returns the requested Keyword Plan ad group in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetKeywordPlanAdGroup(ctx context.Context, in *GetKeywordPlanAdGroupRequest, opts ...grpc.CallOption) (*resources.KeywordPlanAdGroup, error)
	// Creates, updates, or removes Keyword Plan ad groups. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanAdGroupError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error)
}

type keywordPlanAdGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanAdGroupServiceClient(cc grpc.ClientConnInterface) KeywordPlanAdGroupServiceClient {
	return &keywordPlanAdGroupServiceClient{cc}
}

func (c *keywordPlanAdGroupServiceClient) GetKeywordPlanAdGroup(ctx context.Context, in *GetKeywordPlanAdGroupRequest, opts ...grpc.CallOption) (*resources.KeywordPlanAdGroup, error) {
	out := new(resources.KeywordPlanAdGroup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanAdGroupService/GetKeywordPlanAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanAdGroupServiceClient) MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error) {
	out := new(MutateKeywordPlanAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanAdGroupServiceServer is the server API for KeywordPlanAdGroupService service.
// All implementations must embed UnimplementedKeywordPlanAdGroupServiceServer
// for forward compatibility
type KeywordPlanAdGroupServiceServer interface {
	// Returns the requested Keyword Plan ad group in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetKeywordPlanAdGroup(context.Context, *GetKeywordPlanAdGroupRequest) (*resources.KeywordPlanAdGroup, error)
	// Creates, updates, or removes Keyword Plan ad groups. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanAdGroupError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroups(context.Context, *MutateKeywordPlanAdGroupsRequest) (*MutateKeywordPlanAdGroupsResponse, error)
	mustEmbedUnimplementedKeywordPlanAdGroupServiceServer()
}

// UnimplementedKeywordPlanAdGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanAdGroupServiceServer struct {
}

func (UnimplementedKeywordPlanAdGroupServiceServer) GetKeywordPlanAdGroup(context.Context, *GetKeywordPlanAdGroupRequest) (*resources.KeywordPlanAdGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordPlanAdGroup not implemented")
}
func (UnimplementedKeywordPlanAdGroupServiceServer) MutateKeywordPlanAdGroups(context.Context, *MutateKeywordPlanAdGroupsRequest) (*MutateKeywordPlanAdGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlanAdGroups not implemented")
}
func (UnimplementedKeywordPlanAdGroupServiceServer) mustEmbedUnimplementedKeywordPlanAdGroupServiceServer() {
}

// UnsafeKeywordPlanAdGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanAdGroupServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanAdGroupServiceServer interface {
	mustEmbedUnimplementedKeywordPlanAdGroupServiceServer()
}

func RegisterKeywordPlanAdGroupServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanAdGroupServiceServer) {
	s.RegisterService(&KeywordPlanAdGroupService_ServiceDesc, srv)
}

func _KeywordPlanAdGroupService_GetKeywordPlanAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanAdGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupServiceServer).GetKeywordPlanAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanAdGroupService/GetKeywordPlanAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupServiceServer).GetKeywordPlanAdGroup(ctx, req.(*GetKeywordPlanAdGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, req.(*MutateKeywordPlanAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanAdGroupService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanAdGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanAdGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.KeywordPlanAdGroupService",
	HandlerType: (*KeywordPlanAdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanAdGroup",
			Handler:    _KeywordPlanAdGroupService_GetKeywordPlanAdGroup_Handler,
		},
		{
			MethodName: "MutateKeywordPlanAdGroups",
			Handler:    _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/keyword_plan_ad_group_service.proto",
}
