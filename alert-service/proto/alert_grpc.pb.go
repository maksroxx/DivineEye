// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: alert.proto

package alert

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AlertService_CreateAlert_FullMethodName = "/alert.AlertService/CreateAlert"
	AlertService_DeleteAlert_FullMethodName = "/alert.AlertService/DeleteAlert"
	AlertService_GetAlerts_FullMethodName   = "/alert.AlertService/GetAlerts"
)

// AlertServiceClient is the client API for AlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertServiceClient interface {
	CreateAlert(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*CreateAlertResponse, error)
	DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error)
	GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error)
}

type alertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertServiceClient(cc grpc.ClientConnInterface) AlertServiceClient {
	return &alertServiceClient{cc}
}

func (c *alertServiceClient) CreateAlert(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*CreateAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAlertResponse)
	err := c.cc.Invoke(ctx, AlertService_CreateAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAlertResponse)
	err := c.cc.Invoke(ctx, AlertService_DeleteAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_GetAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServiceServer is the server API for AlertService service.
// All implementations must embed UnimplementedAlertServiceServer
// for forward compatibility.
type AlertServiceServer interface {
	CreateAlert(context.Context, *CreateAlertRequest) (*CreateAlertResponse, error)
	DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error)
	GetAlerts(context.Context, *GetAlertsRequest) (*GetAlertsResponse, error)
	mustEmbedUnimplementedAlertServiceServer()
}

// UnimplementedAlertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertServiceServer struct{}

func (UnimplementedAlertServiceServer) CreateAlert(context.Context, *CreateAlertRequest) (*CreateAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlert not implemented")
}
func (UnimplementedAlertServiceServer) DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlert not implemented")
}
func (UnimplementedAlertServiceServer) GetAlerts(context.Context, *GetAlertsRequest) (*GetAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlerts not implemented")
}
func (UnimplementedAlertServiceServer) mustEmbedUnimplementedAlertServiceServer() {}
func (UnimplementedAlertServiceServer) testEmbeddedByValue()                      {}

// UnsafeAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServiceServer will
// result in compilation errors.
type UnsafeAlertServiceServer interface {
	mustEmbedUnimplementedAlertServiceServer()
}

func RegisterAlertServiceServer(s grpc.ServiceRegistrar, srv AlertServiceServer) {
	// If the following call pancis, it indicates UnimplementedAlertServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AlertService_ServiceDesc, srv)
}

func _AlertService_CreateAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).CreateAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_CreateAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).CreateAlert(ctx, req.(*CreateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_DeleteAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).DeleteAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_DeleteAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).DeleteAlert(ctx, req.(*DeleteAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAlerts(ctx, req.(*GetAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertService_ServiceDesc is the grpc.ServiceDesc for AlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alert.AlertService",
	HandlerType: (*AlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlert",
			Handler:    _AlertService_CreateAlert_Handler,
		},
		{
			MethodName: "DeleteAlert",
			Handler:    _AlertService_DeleteAlert_Handler,
		},
		{
			MethodName: "GetAlerts",
			Handler:    _AlertService_GetAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert.proto",
}
