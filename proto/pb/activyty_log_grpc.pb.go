// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ActivityLogDMSServiceClient is the client API for ActivityLogDMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityLogDMSServiceClient interface {
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error)
	GetAllLogAction(ctx context.Context, in *GetLogActionRequest, opts ...grpc.CallOption) (*GetLogActionResponse, error)
}

type activityLogDMSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityLogDMSServiceClient(cc grpc.ClientConnInterface) ActivityLogDMSServiceClient {
	return &activityLogDMSServiceClient{cc}
}

func (c *activityLogDMSServiceClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error) {
	out := new(CreateActionResponse)
	err := c.cc.Invoke(ctx, "/user.ActivityLogDMSService/CreateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityLogDMSServiceClient) GetAllLogAction(ctx context.Context, in *GetLogActionRequest, opts ...grpc.CallOption) (*GetLogActionResponse, error) {
	out := new(GetLogActionResponse)
	err := c.cc.Invoke(ctx, "/user.ActivityLogDMSService/GetAllLogAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityLogDMSServiceServer is the server API for ActivityLogDMSService service.
// All implementations must embed UnimplementedActivityLogDMSServiceServer
// for forward compatibility
type ActivityLogDMSServiceServer interface {
	CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error)
	GetAllLogAction(context.Context, *GetLogActionRequest) (*GetLogActionResponse, error)
	mustEmbedUnimplementedActivityLogDMSServiceServer()
}

// UnimplementedActivityLogDMSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActivityLogDMSServiceServer struct {
}

func (UnimplementedActivityLogDMSServiceServer) CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAction not implemented")
}
func (UnimplementedActivityLogDMSServiceServer) GetAllLogAction(context.Context, *GetLogActionRequest) (*GetLogActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLogAction not implemented")
}
func (UnimplementedActivityLogDMSServiceServer) mustEmbedUnimplementedActivityLogDMSServiceServer() {}

// UnsafeActivityLogDMSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityLogDMSServiceServer will
// result in compilation errors.
type UnsafeActivityLogDMSServiceServer interface {
	mustEmbedUnimplementedActivityLogDMSServiceServer()
}

func RegisterActivityLogDMSServiceServer(s grpc.ServiceRegistrar, srv ActivityLogDMSServiceServer) {
	s.RegisterService(&ActivityLogDMSService_ServiceDesc, srv)
}

func _ActivityLogDMSService_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityLogDMSServiceServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ActivityLogDMSService/CreateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityLogDMSServiceServer).CreateAction(ctx, req.(*CreateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityLogDMSService_GetAllLogAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityLogDMSServiceServer).GetAllLogAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ActivityLogDMSService/GetAllLogAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityLogDMSServiceServer).GetAllLogAction(ctx, req.(*GetLogActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityLogDMSService_ServiceDesc is the grpc.ServiceDesc for ActivityLogDMSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityLogDMSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.ActivityLogDMSService",
	HandlerType: (*ActivityLogDMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAction",
			Handler:    _ActivityLogDMSService_CreateAction_Handler,
		},
		{
			MethodName: "GetAllLogAction",
			Handler:    _ActivityLogDMSService_GetAllLogAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activyty_log.proto",
}
