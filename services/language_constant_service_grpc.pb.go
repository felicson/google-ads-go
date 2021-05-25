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

// LanguageConstantServiceClient is the client API for LanguageConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LanguageConstantServiceClient interface {
	// Returns the requested language constant.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error)
}

type languageConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLanguageConstantServiceClient(cc grpc.ClientConnInterface) LanguageConstantServiceClient {
	return &languageConstantServiceClient{cc}
}

func (c *languageConstantServiceClient) GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error) {
	out := new(resources.LanguageConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v7.services.LanguageConstantService/GetLanguageConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageConstantServiceServer is the server API for LanguageConstantService service.
// All implementations must embed UnimplementedLanguageConstantServiceServer
// for forward compatibility
type LanguageConstantServiceServer interface {
	// Returns the requested language constant.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetLanguageConstant(context.Context, *GetLanguageConstantRequest) (*resources.LanguageConstant, error)
	mustEmbedUnimplementedLanguageConstantServiceServer()
}

// UnimplementedLanguageConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLanguageConstantServiceServer struct {
}

func (UnimplementedLanguageConstantServiceServer) GetLanguageConstant(context.Context, *GetLanguageConstantRequest) (*resources.LanguageConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguageConstant not implemented")
}
func (UnimplementedLanguageConstantServiceServer) mustEmbedUnimplementedLanguageConstantServiceServer() {
}

// UnsafeLanguageConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LanguageConstantServiceServer will
// result in compilation errors.
type UnsafeLanguageConstantServiceServer interface {
	mustEmbedUnimplementedLanguageConstantServiceServer()
}

func RegisterLanguageConstantServiceServer(s grpc.ServiceRegistrar, srv LanguageConstantServiceServer) {
	s.RegisterService(&LanguageConstantService_ServiceDesc, srv)
}

func _LanguageConstantService_GetLanguageConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLanguageConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v7.services.LanguageConstantService/GetLanguageConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, req.(*GetLanguageConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LanguageConstantService_ServiceDesc is the grpc.ServiceDesc for LanguageConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LanguageConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v7.services.LanguageConstantService",
	HandlerType: (*LanguageConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguageConstant",
			Handler:    _LanguageConstantService_GetLanguageConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v7/services/language_constant_service.proto",
}
