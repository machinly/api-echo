// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: apiecho/v1/myna.proto

package v1

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

// MynaClient is the client API for Myna service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MynaClient interface {
	// Get Header from request
	Header(ctx context.Context, in *HeaderRequest, opts ...grpc.CallOption) (*HeaderReply, error)
	// FillData
	FillData(ctx context.Context, in *FillDataRequest, opts ...grpc.CallOption) (*FillDataReply, error)
}

type mynaClient struct {
	cc grpc.ClientConnInterface
}

func NewMynaClient(cc grpc.ClientConnInterface) MynaClient {
	return &mynaClient{cc}
}

func (c *mynaClient) Header(ctx context.Context, in *HeaderRequest, opts ...grpc.CallOption) (*HeaderReply, error) {
	out := new(HeaderReply)
	err := c.cc.Invoke(ctx, "/api.apiecho.v1.Myna/Header", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mynaClient) FillData(ctx context.Context, in *FillDataRequest, opts ...grpc.CallOption) (*FillDataReply, error) {
	out := new(FillDataReply)
	err := c.cc.Invoke(ctx, "/api.apiecho.v1.Myna/FillData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MynaServer is the server API for Myna service.
// All implementations must embed UnimplementedMynaServer
// for forward compatibility
type MynaServer interface {
	// Get Header from request
	Header(context.Context, *HeaderRequest) (*HeaderReply, error)
	// FillData
	FillData(context.Context, *FillDataRequest) (*FillDataReply, error)
	mustEmbedUnimplementedMynaServer()
}

// UnimplementedMynaServer must be embedded to have forward compatible implementations.
type UnimplementedMynaServer struct {
}

func (UnimplementedMynaServer) Header(context.Context, *HeaderRequest) (*HeaderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Header not implemented")
}
func (UnimplementedMynaServer) FillData(context.Context, *FillDataRequest) (*FillDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FillData not implemented")
}
func (UnimplementedMynaServer) mustEmbedUnimplementedMynaServer() {}

// UnsafeMynaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MynaServer will
// result in compilation errors.
type UnsafeMynaServer interface {
	mustEmbedUnimplementedMynaServer()
}

func RegisterMynaServer(s grpc.ServiceRegistrar, srv MynaServer) {
	s.RegisterService(&Myna_ServiceDesc, srv)
}

func _Myna_Header_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MynaServer).Header(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.apiecho.v1.Myna/Header",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MynaServer).Header(ctx, req.(*HeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Myna_FillData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FillDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MynaServer).FillData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.apiecho.v1.Myna/FillData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MynaServer).FillData(ctx, req.(*FillDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Myna_ServiceDesc is the grpc.ServiceDesc for Myna service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Myna_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.apiecho.v1.Myna",
	HandlerType: (*MynaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Header",
			Handler:    _Myna_Header_Handler,
		},
		{
			MethodName: "FillData",
			Handler:    _Myna_FillData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiecho/v1/myna.proto",
}
