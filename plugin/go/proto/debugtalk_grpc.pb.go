// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: plugin/proto/debugtalk.proto

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

// DebugTalkClient is the client API for DebugTalk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugTalkClient interface {
	GetNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetNamesResponse, error)
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type debugTalkClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugTalkClient(cc grpc.ClientConnInterface) DebugTalkClient {
	return &debugTalkClient{cc}
}

func (c *debugTalkClient) GetNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetNamesResponse, error) {
	out := new(GetNamesResponse)
	err := c.cc.Invoke(ctx, "/proto.DebugTalk/GetNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugTalkClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/proto.DebugTalk/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugTalkServer is the server API for DebugTalk service.
// All implementations must embed UnimplementedDebugTalkServer
// for forward compatibility
type DebugTalkServer interface {
	GetNames(context.Context, *Empty) (*GetNamesResponse, error)
	Call(context.Context, *CallRequest) (*CallResponse, error)
	mustEmbedUnimplementedDebugTalkServer()
}

// UnimplementedDebugTalkServer must be embedded to have forward compatible implementations.
type UnimplementedDebugTalkServer struct {
}

func (UnimplementedDebugTalkServer) GetNames(context.Context, *Empty) (*GetNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNames not implemented")
}
func (UnimplementedDebugTalkServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedDebugTalkServer) mustEmbedUnimplementedDebugTalkServer() {}

// UnsafeDebugTalkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugTalkServer will
// result in compilation errors.
type UnsafeDebugTalkServer interface {
	mustEmbedUnimplementedDebugTalkServer()
}

func RegisterDebugTalkServer(s grpc.ServiceRegistrar, srv DebugTalkServer) {
	s.RegisterService(&DebugTalk_ServiceDesc, srv)
}

func _DebugTalk_GetNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugTalkServer).GetNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DebugTalk/GetNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugTalkServer).GetNames(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebugTalk_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugTalkServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DebugTalk/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugTalkServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DebugTalk_ServiceDesc is the grpc.ServiceDesc for DebugTalk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugTalk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DebugTalk",
	HandlerType: (*DebugTalkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNames",
			Handler:    _DebugTalk_GetNames_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _DebugTalk_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/proto/debugtalk.proto",
}
