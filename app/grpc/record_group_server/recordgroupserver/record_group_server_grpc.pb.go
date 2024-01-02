// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: recordgroupserver/record_group_server.proto

package __

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

// RecordGroupServerClient is the client API for RecordGroupServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordGroupServerClient interface {
	RecordGroupList(ctx context.Context, in *RecordGroupListRequest, opts ...grpc.CallOption) (*RecordGroupListResponse, error)
	RecordGroupUpdate(ctx context.Context, in *RecordGroupUpRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error)
	RecordGroupAdd(ctx context.Context, in *RecordGroupAddRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error)
	RecordGroupDel(ctx context.Context, in *RecordGroupDelRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error)
}

type recordGroupServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordGroupServerClient(cc grpc.ClientConnInterface) RecordGroupServerClient {
	return &recordGroupServerClient{cc}
}

func (c *recordGroupServerClient) RecordGroupList(ctx context.Context, in *RecordGroupListRequest, opts ...grpc.CallOption) (*RecordGroupListResponse, error) {
	out := new(RecordGroupListResponse)
	err := c.cc.Invoke(ctx, "/record_group_server.RecordGroupServer/RecordGroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordGroupServerClient) RecordGroupUpdate(ctx context.Context, in *RecordGroupUpRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error) {
	out := new(RecordGroupResponse)
	err := c.cc.Invoke(ctx, "/record_group_server.RecordGroupServer/RecordGroupUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordGroupServerClient) RecordGroupAdd(ctx context.Context, in *RecordGroupAddRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error) {
	out := new(RecordGroupResponse)
	err := c.cc.Invoke(ctx, "/record_group_server.RecordGroupServer/RecordGroupAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordGroupServerClient) RecordGroupDel(ctx context.Context, in *RecordGroupDelRequest, opts ...grpc.CallOption) (*RecordGroupResponse, error) {
	out := new(RecordGroupResponse)
	err := c.cc.Invoke(ctx, "/record_group_server.RecordGroupServer/RecordGroupDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordGroupServerServer is the server API for RecordGroupServer service.
// All implementations must embed UnimplementedRecordGroupServerServer
// for forward compatibility
type RecordGroupServerServer interface {
	RecordGroupList(context.Context, *RecordGroupListRequest) (*RecordGroupListResponse, error)
	RecordGroupUpdate(context.Context, *RecordGroupUpRequest) (*RecordGroupResponse, error)
	RecordGroupAdd(context.Context, *RecordGroupAddRequest) (*RecordGroupResponse, error)
	RecordGroupDel(context.Context, *RecordGroupDelRequest) (*RecordGroupResponse, error)
	mustEmbedUnimplementedRecordGroupServerServer()
}

// UnimplementedRecordGroupServerServer must be embedded to have forward compatible implementations.
type UnimplementedRecordGroupServerServer struct {
}

func (UnimplementedRecordGroupServerServer) RecordGroupList(context.Context, *RecordGroupListRequest) (*RecordGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordGroupList not implemented")
}
func (UnimplementedRecordGroupServerServer) RecordGroupUpdate(context.Context, *RecordGroupUpRequest) (*RecordGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordGroupUpdate not implemented")
}
func (UnimplementedRecordGroupServerServer) RecordGroupAdd(context.Context, *RecordGroupAddRequest) (*RecordGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordGroupAdd not implemented")
}
func (UnimplementedRecordGroupServerServer) RecordGroupDel(context.Context, *RecordGroupDelRequest) (*RecordGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordGroupDel not implemented")
}
func (UnimplementedRecordGroupServerServer) mustEmbedUnimplementedRecordGroupServerServer() {}

// UnsafeRecordGroupServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordGroupServerServer will
// result in compilation errors.
type UnsafeRecordGroupServerServer interface {
	mustEmbedUnimplementedRecordGroupServerServer()
}

func RegisterRecordGroupServerServer(s grpc.ServiceRegistrar, srv RecordGroupServerServer) {
	s.RegisterService(&RecordGroupServer_ServiceDesc, srv)
}

func _RecordGroupServer_RecordGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordGroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordGroupServerServer).RecordGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record_group_server.RecordGroupServer/RecordGroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordGroupServerServer).RecordGroupList(ctx, req.(*RecordGroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordGroupServer_RecordGroupUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordGroupUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordGroupServerServer).RecordGroupUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record_group_server.RecordGroupServer/RecordGroupUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordGroupServerServer).RecordGroupUpdate(ctx, req.(*RecordGroupUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordGroupServer_RecordGroupAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordGroupAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordGroupServerServer).RecordGroupAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record_group_server.RecordGroupServer/RecordGroupAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordGroupServerServer).RecordGroupAdd(ctx, req.(*RecordGroupAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordGroupServer_RecordGroupDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordGroupDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordGroupServerServer).RecordGroupDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record_group_server.RecordGroupServer/RecordGroupDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordGroupServerServer).RecordGroupDel(ctx, req.(*RecordGroupDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordGroupServer_ServiceDesc is the grpc.ServiceDesc for RecordGroupServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordGroupServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "record_group_server.RecordGroupServer",
	HandlerType: (*RecordGroupServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordGroupList",
			Handler:    _RecordGroupServer_RecordGroupList_Handler,
		},
		{
			MethodName: "RecordGroupUpdate",
			Handler:    _RecordGroupServer_RecordGroupUpdate_Handler,
		},
		{
			MethodName: "RecordGroupAdd",
			Handler:    _RecordGroupServer_RecordGroupAdd_Handler,
		},
		{
			MethodName: "RecordGroupDel",
			Handler:    _RecordGroupServer_RecordGroupDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recordgroupserver/record_group_server.proto",
}