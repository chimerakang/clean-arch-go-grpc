// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: internal/delivery/grpc/proto/product.proto

package product_grpc

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

// ProductHandlerClient is the client API for ProductHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductHandlerClient interface {
	Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error)
	GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error)
}

type productHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewProductHandlerClient(cc grpc.ClientConnInterface) ProductHandlerClient {
	return &productHandlerClient{cc}
}

func (c *productHandlerClient) Create(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product_grpc.ProductHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productHandlerClient) GetList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/product_grpc.ProductHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductHandlerServer is the server API for ProductHandler service.
// All implementations must embed UnimplementedProductHandlerServer
// for forward compatibility
type ProductHandlerServer interface {
	Create(context.Context, *Product) (*Product, error)
	GetList(context.Context, *Empty) (*Products, error)
	mustEmbedUnimplementedProductHandlerServer()
}

// UnimplementedProductHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedProductHandlerServer struct {
}

func (UnimplementedProductHandlerServer) Create(context.Context, *Product) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductHandlerServer) GetList(context.Context, *Empty) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedProductHandlerServer) mustEmbedUnimplementedProductHandlerServer() {}

// UnsafeProductHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductHandlerServer will
// result in compilation errors.
type UnsafeProductHandlerServer interface {
	mustEmbedUnimplementedProductHandlerServer()
}

func RegisterProductHandlerServer(s grpc.ServiceRegistrar, srv ProductHandlerServer) {
	s.RegisterService(&ProductHandler_ServiceDesc, srv)
}

func _ProductHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_grpc.ProductHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).Create(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_grpc.ProductHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductHandlerServer).GetList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductHandler_ServiceDesc is the grpc.ServiceDesc for ProductHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_grpc.ProductHandler",
	HandlerType: (*ProductHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductHandler_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ProductHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/delivery/grpc/proto/product.proto",
}
