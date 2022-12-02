// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: contract.proto

package tks_pb

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

// ContractServiceClient is the client API for ContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractServiceClient interface {
	// CreateContract creates new contract.
	CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error)
	// UpdateQuota updates quota of the contract.
	UpdateQuota(ctx context.Context, in *UpdateQuotaRequest, opts ...grpc.CallOption) (*UpdateQuotaResponse, error)
	// UpdateServices updates list of available services of the contract.
	UpdateServices(ctx context.Context, in *UpdateServicesRequest, opts ...grpc.CallOption) (*UpdateServicesResponse, error)
	// GetContract returns a contract if exists.
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
	// GetDefaultContract returns a default contract.
	GetDefaultContract(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetContractResponse, error)
	// Getcontracts return a list of contract.
	GetContracts(ctx context.Context, in *GetContractsRequest, opts ...grpc.CallOption) (*GetContractsResponse, error)
	// GetQuota returns a quota for the contract.
	GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error)
	// GetAvailableServices returns list of available services for the contract.
	GetAvailableServices(ctx context.Context, in *GetAvailableServicesRequest, opts ...grpc.CallOption) (*GetAvailableServicesResponse, error)
}

type contractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContractServiceClient(cc grpc.ClientConnInterface) ContractServiceClient {
	return &contractServiceClient{cc}
}

func (c *contractServiceClient) CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error) {
	out := new(CreateContractResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/CreateContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) UpdateQuota(ctx context.Context, in *UpdateQuotaRequest, opts ...grpc.CallOption) (*UpdateQuotaResponse, error) {
	out := new(UpdateQuotaResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/UpdateQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) UpdateServices(ctx context.Context, in *UpdateServicesRequest, opts ...grpc.CallOption) (*UpdateServicesResponse, error) {
	out := new(UpdateServicesResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/UpdateServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/GetContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetDefaultContract(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/GetDefaultContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetContracts(ctx context.Context, in *GetContractsRequest, opts ...grpc.CallOption) (*GetContractsResponse, error) {
	out := new(GetContractsResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/GetContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error) {
	out := new(GetQuotaResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/GetQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractServiceClient) GetAvailableServices(ctx context.Context, in *GetAvailableServicesRequest, opts ...grpc.CallOption) (*GetAvailableServicesResponse, error) {
	out := new(GetAvailableServicesResponse)
	err := c.cc.Invoke(ctx, "/tks_pb.ContractService/GetAvailableServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractServiceServer is the server API for ContractService service.
// All implementations must embed UnimplementedContractServiceServer
// for forward compatibility
type ContractServiceServer interface {
	// CreateContract creates new contract.
	CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error)
	// UpdateQuota updates quota of the contract.
	UpdateQuota(context.Context, *UpdateQuotaRequest) (*UpdateQuotaResponse, error)
	// UpdateServices updates list of available services of the contract.
	UpdateServices(context.Context, *UpdateServicesRequest) (*UpdateServicesResponse, error)
	// GetContract returns a contract if exists.
	GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error)
	// GetDefaultContract returns a default contract.
	GetDefaultContract(context.Context, *emptypb.Empty) (*GetContractResponse, error)
	// Getcontracts return a list of contract.
	GetContracts(context.Context, *GetContractsRequest) (*GetContractsResponse, error)
	// GetQuota returns a quota for the contract.
	GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error)
	// GetAvailableServices returns list of available services for the contract.
	GetAvailableServices(context.Context, *GetAvailableServicesRequest) (*GetAvailableServicesResponse, error)
	mustEmbedUnimplementedContractServiceServer()
}

// UnimplementedContractServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContractServiceServer struct {
}

func (UnimplementedContractServiceServer) CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedContractServiceServer) UpdateQuota(context.Context, *UpdateQuotaRequest) (*UpdateQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuota not implemented")
}
func (UnimplementedContractServiceServer) UpdateServices(context.Context, *UpdateServicesRequest) (*UpdateServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServices not implemented")
}
func (UnimplementedContractServiceServer) GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedContractServiceServer) GetDefaultContract(context.Context, *emptypb.Empty) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultContract not implemented")
}
func (UnimplementedContractServiceServer) GetContracts(context.Context, *GetContractsRequest) (*GetContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContracts not implemented")
}
func (UnimplementedContractServiceServer) GetQuota(context.Context, *GetQuotaRequest) (*GetQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuota not implemented")
}
func (UnimplementedContractServiceServer) GetAvailableServices(context.Context, *GetAvailableServicesRequest) (*GetAvailableServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableServices not implemented")
}
func (UnimplementedContractServiceServer) mustEmbedUnimplementedContractServiceServer() {}

// UnsafeContractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractServiceServer will
// result in compilation errors.
type UnsafeContractServiceServer interface {
	mustEmbedUnimplementedContractServiceServer()
}

func RegisterContractServiceServer(s grpc.ServiceRegistrar, srv ContractServiceServer) {
	s.RegisterService(&ContractService_ServiceDesc, srv)
}

func _ContractService_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/CreateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).CreateContract(ctx, req.(*CreateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_UpdateQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).UpdateQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/UpdateQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).UpdateQuota(ctx, req.(*UpdateQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_UpdateServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).UpdateServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/UpdateServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).UpdateServices(ctx, req.(*UpdateServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/GetContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetDefaultContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetDefaultContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/GetDefaultContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetDefaultContract(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/GetContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetContracts(ctx, req.(*GetContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/GetQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetQuota(ctx, req.(*GetQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractService_GetAvailableServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).GetAvailableServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tks_pb.ContractService/GetAvailableServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).GetAvailableServices(ctx, req.(*GetAvailableServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractService_ServiceDesc is the grpc.ServiceDesc for ContractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tks_pb.ContractService",
	HandlerType: (*ContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContract",
			Handler:    _ContractService_CreateContract_Handler,
		},
		{
			MethodName: "UpdateQuota",
			Handler:    _ContractService_UpdateQuota_Handler,
		},
		{
			MethodName: "UpdateServices",
			Handler:    _ContractService_UpdateServices_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _ContractService_GetContract_Handler,
		},
		{
			MethodName: "GetDefaultContract",
			Handler:    _ContractService_GetDefaultContract_Handler,
		},
		{
			MethodName: "GetContracts",
			Handler:    _ContractService_GetContracts_Handler,
		},
		{
			MethodName: "GetQuota",
			Handler:    _ContractService_GetQuota_Handler,
		},
		{
			MethodName: "GetAvailableServices",
			Handler:    _ContractService_GetAvailableServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract.proto",
}
