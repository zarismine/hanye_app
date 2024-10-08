// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: cart.proto

package pb

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
	Cart_AddSubCart_FullMethodName  = "/pb.cart/addSubCart"
	Cart_GetCartList_FullMethodName = "/pb.cart/getCartList"
	Cart_CleanCart_FullMethodName   = "/pb.cart/cleanCart"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	AddSubCart(ctx context.Context, in *AddSubCartReq, opts ...grpc.CallOption) (*AddSubCartResp, error)
	GetCartList(ctx context.Context, in *GetCartListReq, opts ...grpc.CallOption) (*GetCartListResp, error)
	CleanCart(ctx context.Context, in *CleanCartReq, opts ...grpc.CallOption) (*CleanCartResp, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) AddSubCart(ctx context.Context, in *AddSubCartReq, opts ...grpc.CallOption) (*AddSubCartResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSubCartResp)
	err := c.cc.Invoke(ctx, Cart_AddSubCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) GetCartList(ctx context.Context, in *GetCartListReq, opts ...grpc.CallOption) (*GetCartListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCartListResp)
	err := c.cc.Invoke(ctx, Cart_GetCartList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) CleanCart(ctx context.Context, in *CleanCartReq, opts ...grpc.CallOption) (*CleanCartResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CleanCartResp)
	err := c.cc.Invoke(ctx, Cart_CleanCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility.
type CartServer interface {
	AddSubCart(context.Context, *AddSubCartReq) (*AddSubCartResp, error)
	GetCartList(context.Context, *GetCartListReq) (*GetCartListResp, error)
	CleanCart(context.Context, *CleanCartReq) (*CleanCartResp, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartServer struct{}

func (UnimplementedCartServer) AddSubCart(context.Context, *AddSubCartReq) (*AddSubCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubCart not implemented")
}
func (UnimplementedCartServer) GetCartList(context.Context, *GetCartListReq) (*GetCartListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartList not implemented")
}
func (UnimplementedCartServer) CleanCart(context.Context, *CleanCartReq) (*CleanCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}
func (UnimplementedCartServer) testEmbeddedByValue()              {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	// If the following call pancis, it indicates UnimplementedCartServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_AddSubCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).AddSubCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_AddSubCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddSubCart(ctx, req.(*AddSubCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_GetCartList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetCartList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_GetCartList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetCartList(ctx, req.(*GetCartListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_CleanCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).CleanCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_CleanCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).CleanCart(ctx, req.(*CleanCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addSubCart",
			Handler:    _Cart_AddSubCart_Handler,
		},
		{
			MethodName: "getCartList",
			Handler:    _Cart_GetCartList_Handler,
		},
		{
			MethodName: "cleanCart",
			Handler:    _Cart_CleanCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
