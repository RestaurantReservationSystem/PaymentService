// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: menu.proto

package genproto

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

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	CreateReservation(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateReservation(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MenuResponse, error)
	GetAllReservation(ctx context.Context, in *GetAllMenuRequest, opts ...grpc.CallOption) (*MenusResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) CreateReservation(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.MenuService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateReservation(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.MenuService/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.MenuService/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetByIdReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MenuResponse, error) {
	out := new(MenuResponse)
	err := c.cc.Invoke(ctx, "/protos.MenuService/GetByIdReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetAllReservation(ctx context.Context, in *GetAllMenuRequest, opts ...grpc.CallOption) (*MenusResponse, error) {
	out := new(MenusResponse)
	err := c.cc.Invoke(ctx, "/protos.MenuService/GetAllReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility
type MenuServiceServer interface {
	CreateReservation(context.Context, *CreateMenuRequest) (*Void, error)
	UpdateReservation(context.Context, *UpdateMenuRequest) (*Void, error)
	DeleteReservation(context.Context, *IdRequest) (*Void, error)
	GetByIdReservation(context.Context, *IdRequest) (*MenuResponse, error)
	GetAllReservation(context.Context, *GetAllMenuRequest) (*MenusResponse, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (UnimplementedMenuServiceServer) CreateReservation(context.Context, *CreateMenuRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedMenuServiceServer) UpdateReservation(context.Context, *UpdateMenuRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedMenuServiceServer) DeleteReservation(context.Context, *IdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedMenuServiceServer) GetByIdReservation(context.Context, *IdRequest) (*MenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdReservation not implemented")
}
func (UnimplementedMenuServiceServer) GetAllReservation(context.Context, *GetAllMenuRequest) (*MenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservation not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MenuService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateReservation(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MenuService/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateReservation(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MenuService/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteReservation(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetByIdReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetByIdReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MenuService/GetByIdReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetByIdReservation(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetAllReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetAllReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MenuService/GetAllReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetAllReservation(ctx, req.(*GetAllMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReservation",
			Handler:    _MenuService_CreateReservation_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _MenuService_UpdateReservation_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _MenuService_DeleteReservation_Handler,
		},
		{
			MethodName: "GetByIdReservation",
			Handler:    _MenuService_GetByIdReservation_Handler,
		},
		{
			MethodName: "GetAllReservation",
			Handler:    _MenuService_GetAllReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu.proto",
}