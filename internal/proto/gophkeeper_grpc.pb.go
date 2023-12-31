// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/gophkeeper.proto

package gophkeeper

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
	Storage_SaveCredentials_FullMethodName = "/gophkeeper.Storage/SaveCredentials"
	Storage_LoadCredentials_FullMethodName = "/gophkeeper.Storage/LoadCredentials"
	Storage_SaveText_FullMethodName        = "/gophkeeper.Storage/SaveText"
	Storage_LoadText_FullMethodName        = "/gophkeeper.Storage/LoadText"
	Storage_SaveBinary_FullMethodName      = "/gophkeeper.Storage/SaveBinary"
	Storage_LoadBinary_FullMethodName      = "/gophkeeper.Storage/LoadBinary"
	Storage_SaveBankCard_FullMethodName    = "/gophkeeper.Storage/SaveBankCard"
	Storage_LoadBankCard_FullMethodName    = "/gophkeeper.Storage/LoadBankCard"
	Storage_RegisterUser_FullMethodName    = "/gophkeeper.Storage/RegisterUser"
	Storage_LoginUser_FullMethodName       = "/gophkeeper.Storage/LoginUser"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	SaveCredentials(ctx context.Context, in *SaveCredentialsDataRequest, opts ...grpc.CallOption) (*SaveCredentialsDataResponse, error)
	LoadCredentials(ctx context.Context, in *LoadCredentialsDataRequest, opts ...grpc.CallOption) (*LoadCredentialsDataResponse, error)
	SaveText(ctx context.Context, in *SaveTextDataRequest, opts ...grpc.CallOption) (*SaveTextDataResponse, error)
	LoadText(ctx context.Context, in *LoadTextDataRequest, opts ...grpc.CallOption) (*LoadTextDataResponse, error)
	SaveBinary(ctx context.Context, in *SaveBinaryDataRequest, opts ...grpc.CallOption) (*SaveBinaryDataResponse, error)
	LoadBinary(ctx context.Context, in *LoadBinaryDataRequest, opts ...grpc.CallOption) (*LoadBinaryDataResponse, error)
	SaveBankCard(ctx context.Context, in *SaveBankCardDataRequest, opts ...grpc.CallOption) (*SaveBankCardDataResponse, error)
	LoadBankCard(ctx context.Context, in *LoadBankCardDataRequest, opts ...grpc.CallOption) (*LoadBankCardDataResponse, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) SaveCredentials(ctx context.Context, in *SaveCredentialsDataRequest, opts ...grpc.CallOption) (*SaveCredentialsDataResponse, error) {
	out := new(SaveCredentialsDataResponse)
	err := c.cc.Invoke(ctx, Storage_SaveCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) LoadCredentials(ctx context.Context, in *LoadCredentialsDataRequest, opts ...grpc.CallOption) (*LoadCredentialsDataResponse, error) {
	out := new(LoadCredentialsDataResponse)
	err := c.cc.Invoke(ctx, Storage_LoadCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveText(ctx context.Context, in *SaveTextDataRequest, opts ...grpc.CallOption) (*SaveTextDataResponse, error) {
	out := new(SaveTextDataResponse)
	err := c.cc.Invoke(ctx, Storage_SaveText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) LoadText(ctx context.Context, in *LoadTextDataRequest, opts ...grpc.CallOption) (*LoadTextDataResponse, error) {
	out := new(LoadTextDataResponse)
	err := c.cc.Invoke(ctx, Storage_LoadText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveBinary(ctx context.Context, in *SaveBinaryDataRequest, opts ...grpc.CallOption) (*SaveBinaryDataResponse, error) {
	out := new(SaveBinaryDataResponse)
	err := c.cc.Invoke(ctx, Storage_SaveBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) LoadBinary(ctx context.Context, in *LoadBinaryDataRequest, opts ...grpc.CallOption) (*LoadBinaryDataResponse, error) {
	out := new(LoadBinaryDataResponse)
	err := c.cc.Invoke(ctx, Storage_LoadBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SaveBankCard(ctx context.Context, in *SaveBankCardDataRequest, opts ...grpc.CallOption) (*SaveBankCardDataResponse, error) {
	out := new(SaveBankCardDataResponse)
	err := c.cc.Invoke(ctx, Storage_SaveBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) LoadBankCard(ctx context.Context, in *LoadBankCardDataRequest, opts ...grpc.CallOption) (*LoadBankCardDataResponse, error) {
	out := new(LoadBankCardDataResponse)
	err := c.cc.Invoke(ctx, Storage_LoadBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, Storage_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Storage_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	SaveCredentials(context.Context, *SaveCredentialsDataRequest) (*SaveCredentialsDataResponse, error)
	LoadCredentials(context.Context, *LoadCredentialsDataRequest) (*LoadCredentialsDataResponse, error)
	SaveText(context.Context, *SaveTextDataRequest) (*SaveTextDataResponse, error)
	LoadText(context.Context, *LoadTextDataRequest) (*LoadTextDataResponse, error)
	SaveBinary(context.Context, *SaveBinaryDataRequest) (*SaveBinaryDataResponse, error)
	LoadBinary(context.Context, *LoadBinaryDataRequest) (*LoadBinaryDataResponse, error)
	SaveBankCard(context.Context, *SaveBankCardDataRequest) (*SaveBankCardDataResponse, error)
	LoadBankCard(context.Context, *LoadBankCardDataRequest) (*LoadBankCardDataResponse, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) SaveCredentials(context.Context, *SaveCredentialsDataRequest) (*SaveCredentialsDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCredentials not implemented")
}
func (UnimplementedStorageServer) LoadCredentials(context.Context, *LoadCredentialsDataRequest) (*LoadCredentialsDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCredentials not implemented")
}
func (UnimplementedStorageServer) SaveText(context.Context, *SaveTextDataRequest) (*SaveTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedStorageServer) LoadText(context.Context, *LoadTextDataRequest) (*LoadTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadText not implemented")
}
func (UnimplementedStorageServer) SaveBinary(context.Context, *SaveBinaryDataRequest) (*SaveBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinary not implemented")
}
func (UnimplementedStorageServer) LoadBinary(context.Context, *LoadBinaryDataRequest) (*LoadBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadBinary not implemented")
}
func (UnimplementedStorageServer) SaveBankCard(context.Context, *SaveBankCardDataRequest) (*SaveBankCardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBankCard not implemented")
}
func (UnimplementedStorageServer) LoadBankCard(context.Context, *LoadBankCardDataRequest) (*LoadBankCardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadBankCard not implemented")
}
func (UnimplementedStorageServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedStorageServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_SaveCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCredentialsDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveCredentials(ctx, req.(*SaveCredentialsDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_LoadCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCredentialsDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).LoadCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_LoadCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).LoadCredentials(ctx, req.(*LoadCredentialsDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveText(ctx, req.(*SaveTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_LoadText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).LoadText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_LoadText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).LoadText(ctx, req.(*LoadTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveBinary(ctx, req.(*SaveBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_LoadBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).LoadBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_LoadBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).LoadBinary(ctx, req.(*LoadBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SaveBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBankCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SaveBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_SaveBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SaveBankCard(ctx, req.(*SaveBankCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_LoadBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBankCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).LoadBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_LoadBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).LoadBankCard(ctx, req.(*LoadBankCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveCredentials",
			Handler:    _Storage_SaveCredentials_Handler,
		},
		{
			MethodName: "LoadCredentials",
			Handler:    _Storage_LoadCredentials_Handler,
		},
		{
			MethodName: "SaveText",
			Handler:    _Storage_SaveText_Handler,
		},
		{
			MethodName: "LoadText",
			Handler:    _Storage_LoadText_Handler,
		},
		{
			MethodName: "SaveBinary",
			Handler:    _Storage_SaveBinary_Handler,
		},
		{
			MethodName: "LoadBinary",
			Handler:    _Storage_LoadBinary_Handler,
		},
		{
			MethodName: "SaveBankCard",
			Handler:    _Storage_SaveBankCard_Handler,
		},
		{
			MethodName: "LoadBankCard",
			Handler:    _Storage_LoadBankCard_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _Storage_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Storage_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gophkeeper.proto",
}
