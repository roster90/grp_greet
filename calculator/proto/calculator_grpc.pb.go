// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: calculator.proto

package proto

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
	CalculatorService_Sum_FullMethodName            = "/calculator.CalculatorService/Sum"
	CalculatorService_Primes_FullMethodName         = "/calculator.CalculatorService/Primes"
	CalculatorService_ComputeAverage_FullMethodName = "/calculator.CalculatorService/ComputeAverage"
	CalculatorService_FindMaximum_FullMethodName    = "/calculator.CalculatorService/FindMaximum"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimesResponse], error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse], error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse], error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Sum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_Primes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PrimesRequest, PrimesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_PrimesClient = grpc.ServerStreamingClient[PrimesResponse]

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], CalculatorService_ComputeAverage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CalculatorRequest, CalculatorResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_ComputeAverageClient = grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse]

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], CalculatorService_FindMaximum_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CalculatorRequest, CalculatorResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_FindMaximumClient = grpc.ClientStreamingClient[CalculatorRequest, CalculatorResponse]

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility.
type CalculatorServiceServer interface {
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	Primes(*PrimesRequest, grpc.ServerStreamingServer[PrimesResponse]) error
	ComputeAverage(grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]) error
	FindMaximum(grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServiceServer struct{}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Primes(*PrimesRequest, grpc.ServerStreamingServer[PrimesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCalculatorServiceServer) ComputeAverage(grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) FindMaximum(grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}
func (UnimplementedCalculatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Primes(m, &grpc.GenericServerStream[PrimesRequest, PrimesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_PrimesServer = grpc.ServerStreamingServer[PrimesResponse]

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&grpc.GenericServerStream[CalculatorRequest, CalculatorResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_ComputeAverageServer = grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&grpc.GenericServerStream[CalculatorRequest, CalculatorResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_FindMaximumServer = grpc.ClientStreamingServer[CalculatorRequest, CalculatorResponse]

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _CalculatorService_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
