// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: structs.proto

package arithmetic

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

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	SendArithmetic(ctx context.Context, in *ArithmeticRequest, opts ...grpc.CallOption) (*ArithmeticResponse, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) SendArithmetic(ctx context.Context, in *ArithmeticRequest, opts ...grpc.CallOption) (*ArithmeticResponse, error) {
	out := new(ArithmeticResponse)
	err := c.cc.Invoke(ctx, "/structs.ArithmeticService/SendArithmetic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations must embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	SendArithmetic(context.Context, *ArithmeticRequest) (*ArithmeticResponse, error)
	mustEmbedUnimplementedArithmeticServiceServer()
}

// UnimplementedArithmeticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) SendArithmetic(context.Context, *ArithmeticRequest) (*ArithmeticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendArithmetic not implemented")
}
func (UnimplementedArithmeticServiceServer) mustEmbedUnimplementedArithmeticServiceServer() {}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_SendArithmetic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithmeticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).SendArithmetic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/structs.ArithmeticService/SendArithmetic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).SendArithmetic(ctx, req.(*ArithmeticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "structs.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendArithmetic",
			Handler:    _ArithmeticService_SendArithmetic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "structs.proto",
}
