// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.0
// source: proto/test_grpc.proto

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

const (
	ServiceStream_DownLoadFile_FullMethodName = "/grpcFd.ServiceStream/DownLoadFile"
)

// ServiceStreamClient is the client API for ServiceStream service.
//
// For semantics around ctx user and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceStreamClient interface {
	DownLoadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (ServiceStream_DownLoadFileClient, error)
}

type serviceStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceStreamClient(cc grpc.ClientConnInterface) ServiceStreamClient {
	return &serviceStreamClient{cc}
}

func (c *serviceStreamClient) DownLoadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (ServiceStream_DownLoadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceStream_ServiceDesc.Streams[0], ServiceStream_DownLoadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceStreamDownLoadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceStream_DownLoadFileClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type serviceStreamDownLoadFileClient struct {
	grpc.ClientStream
}

func (x *serviceStreamDownLoadFileClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceStreamServer is the server API for ServiceStream service.
// All implementations must embed UnimplementedServiceStreamServer
// for forward compatibility
type ServiceStreamServer interface {
	DownLoadFile(*Request, ServiceStream_DownLoadFileServer) error
	mustEmbedUnimplementedServiceStreamServer()
}

// UnimplementedServiceStreamServer must be embedded to have forward compatible implementations.
type UnimplementedServiceStreamServer struct {
}

func (UnimplementedServiceStreamServer) DownLoadFile(*Request, ServiceStream_DownLoadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownLoadFile not implemented")
}
func (UnimplementedServiceStreamServer) mustEmbedUnimplementedServiceStreamServer() {}

// UnsafeServiceStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceStreamServer will
// result in compilation errors.
type UnsafeServiceStreamServer interface {
	mustEmbedUnimplementedServiceStreamServer()
}

func RegisterServiceStreamServer(s grpc.ServiceRegistrar, srv ServiceStreamServer) {
	s.RegisterService(&ServiceStream_ServiceDesc, srv)
}

func _ServiceStream_DownLoadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceStreamServer).DownLoadFile(m, &serviceStreamDownLoadFileServer{stream})
}

type ServiceStream_DownLoadFileServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type serviceStreamDownLoadFileServer struct {
	grpc.ServerStream
}

func (x *serviceStreamDownLoadFileServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServiceStream_ServiceDesc is the grpc.ServiceDesc for ServiceStream service.
// It's only intended for direct user with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcFd.ServiceStream",
	HandlerType: (*ServiceStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownLoadFile",
			Handler:       _ServiceStream_DownLoadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/test_grpc.proto",
}
