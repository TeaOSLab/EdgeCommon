// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_sms_sender.proto

package pb

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
	SMSSenderService_SendSMS_FullMethodName = "/pb.SMSSenderService/sendSMS"
)

// SMSSenderServiceClient is the client API for SMSSenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSSenderServiceClient interface {
	// 发送短信
	SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error)
}

type sMSSenderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSSenderServiceClient(cc grpc.ClientConnInterface) SMSSenderServiceClient {
	return &sMSSenderServiceClient{cc}
}

func (c *sMSSenderServiceClient) SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error) {
	out := new(SendSMSResponse)
	err := c.cc.Invoke(ctx, SMSSenderService_SendSMS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSSenderServiceServer is the server API for SMSSenderService service.
// All implementations should embed UnimplementedSMSSenderServiceServer
// for forward compatibility
type SMSSenderServiceServer interface {
	// 发送短信
	SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
}

// UnimplementedSMSSenderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSMSSenderServiceServer struct {
}

func (UnimplementedSMSSenderServiceServer) SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMS not implemented")
}

// UnsafeSMSSenderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSSenderServiceServer will
// result in compilation errors.
type UnsafeSMSSenderServiceServer interface {
	mustEmbedUnimplementedSMSSenderServiceServer()
}

func RegisterSMSSenderServiceServer(s grpc.ServiceRegistrar, srv SMSSenderServiceServer) {
	s.RegisterService(&SMSSenderService_ServiceDesc, srv)
}

func _SMSSenderService_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSSenderServiceServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SMSSenderService_SendSMS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSSenderServiceServer).SendSMS(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SMSSenderService_ServiceDesc is the grpc.ServiceDesc for SMSSenderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSSenderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SMSSenderService",
	HandlerType: (*SMSSenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendSMS",
			Handler:    _SMSSenderService_SendSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_sms_sender.proto",
}
