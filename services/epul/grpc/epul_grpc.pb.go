// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: epul.proto

package epul

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PeopleServiceClient is the client API for PeopleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeopleServiceClient interface {
	ListPeople(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Peoples, error)
	DetailPeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error)
	AddPeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error)
}

type peopleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeopleServiceClient(cc grpc.ClientConnInterface) PeopleServiceClient {
	return &peopleServiceClient{cc}
}

func (c *peopleServiceClient) ListPeople(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Peoples, error) {
	out := new(Peoples)
	err := c.cc.Invoke(ctx, "/PeopleService/ListPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peopleServiceClient) DetailPeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error) {
	out := new(People)
	err := c.cc.Invoke(ctx, "/PeopleService/DetailPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peopleServiceClient) AddPeople(ctx context.Context, in *People, opts ...grpc.CallOption) (*People, error) {
	out := new(People)
	err := c.cc.Invoke(ctx, "/PeopleService/AddPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeopleServiceServer is the server API for PeopleService service.
// All implementations must embed UnimplementedPeopleServiceServer
// for forward compatibility
type PeopleServiceServer interface {
	ListPeople(context.Context, *emptypb.Empty) (*Peoples, error)
	DetailPeople(context.Context, *People) (*People, error)
	AddPeople(context.Context, *People) (*People, error)
	mustEmbedUnimplementedPeopleServiceServer()
}

// UnimplementedPeopleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeopleServiceServer struct {
}

func (UnimplementedPeopleServiceServer) ListPeople(context.Context, *emptypb.Empty) (*Peoples, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeople not implemented")
}
func (UnimplementedPeopleServiceServer) DetailPeople(context.Context, *People) (*People, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailPeople not implemented")
}
func (UnimplementedPeopleServiceServer) AddPeople(context.Context, *People) (*People, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeople not implemented")
}
func (UnimplementedPeopleServiceServer) mustEmbedUnimplementedPeopleServiceServer() {}

// UnsafePeopleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeopleServiceServer will
// result in compilation errors.
type UnsafePeopleServiceServer interface {
	mustEmbedUnimplementedPeopleServiceServer()
}

func RegisterPeopleServiceServer(s grpc.ServiceRegistrar, srv PeopleServiceServer) {
	s.RegisterService(&PeopleService_ServiceDesc, srv)
}

func _PeopleService_ListPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).ListPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeopleService/ListPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).ListPeople(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeopleService_DetailPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).DetailPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeopleService/DetailPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).DetailPeople(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeopleService_AddPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).AddPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeopleService/AddPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).AddPeople(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

// PeopleService_ServiceDesc is the grpc.ServiceDesc for PeopleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeopleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PeopleService",
	HandlerType: (*PeopleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPeople",
			Handler:    _PeopleService_ListPeople_Handler,
		},
		{
			MethodName: "DetailPeople",
			Handler:    _PeopleService_DetailPeople_Handler,
		},
		{
			MethodName: "AddPeople",
			Handler:    _PeopleService_AddPeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "epul.proto",
}
