// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pkg/proto/app.proto

package proto

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

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SetEntry(ctx context.Context, in *SetEntryRequest, opts ...grpc.CallOption) (*SetEntryResponse, error)
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
	GetAllEntries(ctx context.Context, in *GetAllEntriesRequest, opts ...grpc.CallOption) (*GetAllEntriesResponse, error)
	GetAllTypes(ctx context.Context, in *GetAllTypesRequest, opts ...grpc.CallOption) (*GetAllTypesResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/proto.App/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.App/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) SetEntry(ctx context.Context, in *SetEntryRequest, opts ...grpc.CallOption) (*SetEntryResponse, error) {
	out := new(SetEntryResponse)
	err := c.cc.Invoke(ctx, "/proto.App/SetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, "/proto.App/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetAllEntries(ctx context.Context, in *GetAllEntriesRequest, opts ...grpc.CallOption) (*GetAllEntriesResponse, error) {
	out := new(GetAllEntriesResponse)
	err := c.cc.Invoke(ctx, "/proto.App/GetAllEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetAllTypes(ctx context.Context, in *GetAllTypesRequest, opts ...grpc.CallOption) (*GetAllTypesResponse, error) {
	out := new(GetAllTypesResponse)
	err := c.cc.Invoke(ctx, "/proto.App/GetAllTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := c.cc.Invoke(ctx, "/proto.App/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SetEntry(context.Context, *SetEntryRequest) (*SetEntryResponse, error)
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	GetAllEntries(context.Context, *GetAllEntriesRequest) (*GetAllEntriesResponse, error)
	GetAllTypes(context.Context, *GetAllTypesRequest) (*GetAllTypesResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedAppServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAppServer) SetEntry(context.Context, *SetEntryRequest) (*SetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntry not implemented")
}
func (UnimplementedAppServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedAppServer) GetAllEntries(context.Context, *GetAllEntriesRequest) (*GetAllEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEntries not implemented")
}
func (UnimplementedAppServer) GetAllTypes(context.Context, *GetAllTypesRequest) (*GetAllTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTypes not implemented")
}
func (UnimplementedAppServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_SetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).SetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/SetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).SetEntry(ctx, req.(*SetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetAllEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetAllEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/GetAllEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetAllEntries(ctx, req.(*GetAllEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetAllTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetAllTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/GetAllTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetAllTypes(ctx, req.(*GetAllTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.App/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _App_Registration_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _App_Login_Handler,
		},
		{
			MethodName: "SetEntry",
			Handler:    _App_SetEntry_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _App_GetEntry_Handler,
		},
		{
			MethodName: "GetAllEntries",
			Handler:    _App_GetAllEntries_Handler,
		},
		{
			MethodName: "GetAllTypes",
			Handler:    _App_GetAllTypes_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _App_DeleteEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/app.proto",
}