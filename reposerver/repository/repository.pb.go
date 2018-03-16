// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reposerver/repository/repository.proto

/*
Package repository is a generated protocol buffer package.

It is generated from these files:
	reposerver/repository/repository.proto

It has these top-level messages:
	ManifestRequest
	ManifestResponse
	EnvParamsRequest
	EnvParamsResponse
*/
package repository

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "k8s.io/api/core/v1"
import github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ManifestRequest is a query for manifest generation.
type ManifestRequest struct {
	Repo        *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Revision    string                                                                `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Path        string                                                                `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Environment string                                                                `protobuf:"bytes,4,opt,name=environment" json:"environment,omitempty"`
}

func (m *ManifestRequest) Reset()                    { *m = ManifestRequest{} }
func (m *ManifestRequest) String() string            { return proto.CompactTextString(m) }
func (*ManifestRequest) ProtoMessage()               {}
func (*ManifestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ManifestRequest) GetRepo() *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *ManifestRequest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ManifestRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ManifestRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

type ManifestResponse struct {
	Manifests []string `protobuf:"bytes,1,rep,name=manifests" json:"manifests,omitempty"`
	Namespace string   `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Server    string   `protobuf:"bytes,3,opt,name=server" json:"server,omitempty"`
}

func (m *ManifestResponse) Reset()                    { *m = ManifestResponse{} }
func (m *ManifestResponse) String() string            { return proto.CompactTextString(m) }
func (*ManifestResponse) ProtoMessage()               {}
func (*ManifestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ManifestResponse) GetManifests() []string {
	if m != nil {
		return m.Manifests
	}
	return nil
}

func (m *ManifestResponse) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ManifestResponse) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type EnvParamsRequest struct {
	Repo        *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Revision    string                                                                `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Path        string                                                                `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Environment string                                                                `protobuf:"bytes,4,opt,name=environment" json:"environment,omitempty"`
}

func (m *EnvParamsRequest) Reset()                    { *m = EnvParamsRequest{} }
func (m *EnvParamsRequest) String() string            { return proto.CompactTextString(m) }
func (*EnvParamsRequest) ProtoMessage()               {}
func (*EnvParamsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EnvParamsRequest) GetRepo() *github_com_argoproj_argo_cd_pkg_apis_application_v1alpha1.Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *EnvParamsRequest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *EnvParamsRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *EnvParamsRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

type EnvParamsResponse struct {
	Params map[string]string `protobuf:"bytes,1,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnvParamsResponse) Reset()                    { *m = EnvParamsResponse{} }
func (m *EnvParamsResponse) String() string            { return proto.CompactTextString(m) }
func (*EnvParamsResponse) ProtoMessage()               {}
func (*EnvParamsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnvParamsResponse) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*ManifestRequest)(nil), "repository.ManifestRequest")
	proto.RegisterType((*ManifestResponse)(nil), "repository.ManifestResponse")
	proto.RegisterType((*EnvParamsRequest)(nil), "repository.EnvParamsRequest")
	proto.RegisterType((*EnvParamsResponse)(nil), "repository.EnvParamsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepositoryService service

type RepositoryServiceClient interface {
	// Generate manifest for application in specified repo name and revision
	GenerateManifest(ctx context.Context, in *ManifestRequest, opts ...grpc.CallOption) (*ManifestResponse, error)
	// Retrieve Ksonnet environment params in specified repo name and revision
	GetEnvParams(ctx context.Context, in *EnvParamsRequest, opts ...grpc.CallOption) (*EnvParamsResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) GenerateManifest(ctx context.Context, in *ManifestRequest, opts ...grpc.CallOption) (*ManifestResponse, error) {
	out := new(ManifestResponse)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/GenerateManifest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryServiceClient) GetEnvParams(ctx context.Context, in *EnvParamsRequest, opts ...grpc.CallOption) (*EnvParamsResponse, error) {
	out := new(EnvParamsResponse)
	err := grpc.Invoke(ctx, "/repository.RepositoryService/GetEnvParams", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepositoryService service

type RepositoryServiceServer interface {
	// Generate manifest for application in specified repo name and revision
	GenerateManifest(context.Context, *ManifestRequest) (*ManifestResponse, error)
	// Retrieve Ksonnet environment params in specified repo name and revision
	GetEnvParams(context.Context, *EnvParamsRequest) (*EnvParamsResponse, error)
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_GenerateManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GenerateManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/GenerateManifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GenerateManifest(ctx, req.(*ManifestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryService_GetEnvParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).GetEnvParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/repository.RepositoryService/GetEnvParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).GetEnvParams(ctx, req.(*EnvParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "repository.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateManifest",
			Handler:    _RepositoryService_GenerateManifest_Handler,
		},
		{
			MethodName: "GetEnvParams",
			Handler:    _RepositoryService_GetEnvParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reposerver/repository/repository.proto",
}

func init() { proto.RegisterFile("reposerver/repository/repository.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x41, 0x8f, 0xd3, 0x3c,
	0x10, 0xdd, 0x6c, 0xfb, 0x55, 0x5f, 0xa7, 0x48, 0x74, 0xad, 0x15, 0x8a, 0x42, 0x91, 0xaa, 0x1c,
	0x50, 0x39, 0x60, 0xab, 0xdd, 0xcb, 0xc2, 0x0d, 0xa4, 0x6a, 0xc5, 0x61, 0x05, 0x0a, 0x27, 0xb8,
	0x20, 0x6f, 0x3a, 0x9b, 0x9a, 0x36, 0xb6, 0xb1, 0xdd, 0x48, 0xfd, 0x19, 0xfc, 0x16, 0xee, 0xf0,
	0xd7, 0x50, 0x1c, 0xb7, 0x89, 0xd8, 0x65, 0xef, 0xdc, 0xde, 0xbc, 0x71, 0xe6, 0xcd, 0x3c, 0x8f,
	0x03, 0xcf, 0x0d, 0x6a, 0x65, 0xd1, 0x54, 0x68, 0x98, 0x87, 0xc2, 0x29, 0xb3, 0xef, 0x40, 0xaa,
	0x8d, 0x72, 0x8a, 0x40, 0xcb, 0x24, 0xe7, 0x85, 0x2a, 0x94, 0xa7, 0x59, 0x8d, 0x9a, 0x13, 0xc9,
	0xa4, 0x50, 0xaa, 0xd8, 0x22, 0xe3, 0x5a, 0x30, 0x2e, 0xa5, 0x72, 0xdc, 0x09, 0x25, 0x6d, 0xc8,
	0xa6, 0x9b, 0x4b, 0x4b, 0x85, 0xf2, 0xd9, 0x5c, 0x19, 0x64, 0xd5, 0x9c, 0x15, 0x28, 0xd1, 0x70,
	0x87, 0xab, 0x70, 0xe6, 0x5d, 0x21, 0xdc, 0x7a, 0x77, 0x43, 0x73, 0x55, 0x32, 0x6e, 0xbc, 0xc4,
	0x57, 0x0f, 0x5e, 0xe6, 0x2b, 0xa6, 0x37, 0x45, 0xfd, 0xb1, 0x65, 0x5c, 0xeb, 0xad, 0xc8, 0x7d,
	0x71, 0x56, 0xcd, 0xf9, 0x56, 0xaf, 0xf9, 0x9d, 0x52, 0xe9, 0xcf, 0x08, 0x1e, 0x5f, 0x73, 0x29,
	0x6e, 0xd1, 0xba, 0x0c, 0xbf, 0xed, 0xd0, 0x3a, 0xf2, 0x09, 0xfa, 0xf5, 0x10, 0x71, 0x34, 0x8d,
	0x66, 0xa3, 0xc5, 0x92, 0xb6, 0x6a, 0xf4, 0xa0, 0xe6, 0xc1, 0x97, 0x7c, 0x45, 0xf5, 0xa6, 0xa0,
	0xb5, 0x1a, 0xed, 0xa8, 0xd1, 0x83, 0x1a, 0xcd, 0x8e, 0x5e, 0x64, 0xbe, 0x24, 0x49, 0xe0, 0x7f,
	0x83, 0x95, 0xb0, 0x42, 0xc9, 0xf8, 0x74, 0x1a, 0xcd, 0x86, 0xd9, 0x31, 0x26, 0x04, 0xfa, 0x9a,
	0xbb, 0x75, 0xdc, 0xf3, 0xbc, 0xc7, 0x64, 0x0a, 0x23, 0x94, 0x95, 0x30, 0x4a, 0x96, 0x28, 0x5d,
	0xdc, 0xf7, 0xa9, 0x2e, 0x95, 0xde, 0xc2, 0xb8, 0xed, 0xdf, 0x6a, 0x25, 0x2d, 0x92, 0x09, 0x0c,
	0xcb, 0xc0, 0xd9, 0x38, 0x9a, 0xf6, 0x66, 0xc3, 0xac, 0x25, 0xea, 0xac, 0xe4, 0x25, 0x5a, 0xcd,
	0x73, 0x0c, 0x4d, 0xb4, 0x04, 0x79, 0x02, 0x83, 0xe6, 0x96, 0x43, 0x1f, 0x21, 0x4a, 0x7f, 0x45,
	0x30, 0x5e, 0xca, 0xea, 0x03, 0x37, 0xbc, 0xb4, 0xff, 0xa4, 0x53, 0xdf, 0x23, 0x38, 0xeb, 0x4c,
	0x10, 0xbc, 0x7a, 0x03, 0x03, 0xed, 0x19, 0x6f, 0xd4, 0x68, 0xf1, 0x82, 0x76, 0x56, 0xfa, 0xce,
	0x71, 0xda, 0x84, 0x4b, 0xe9, 0xcc, 0x3e, 0x0b, 0x1f, 0x26, 0xaf, 0x60, 0xd4, 0xa1, 0xc9, 0x18,
	0x7a, 0x1b, 0xdc, 0x7b, 0x4f, 0x86, 0x59, 0x0d, 0xc9, 0x39, 0xfc, 0x57, 0xf1, 0xed, 0xee, 0xe0,
	0x76, 0x13, 0xbc, 0x3e, 0xbd, 0x8c, 0x16, 0x3f, 0x22, 0x38, 0x6b, 0x47, 0xff, 0x88, 0xa6, 0x12,
	0x39, 0x92, 0xf7, 0x30, 0xbe, 0x0a, 0x7b, 0x7a, 0xb8, 0x5b, 0xf2, 0xb4, 0xdb, 0xd7, 0x1f, 0x1b,
	0x9b, 0x4c, 0xee, 0x4f, 0x36, 0x3d, 0xa7, 0x27, 0xe4, 0x1a, 0x1e, 0x5d, 0xa1, 0x3b, 0x4e, 0x43,
	0x26, 0x7f, 0x19, 0xb2, 0xa9, 0xf6, 0xec, 0x41, 0x0b, 0xd2, 0x93, 0xb7, 0x17, 0x9f, 0xe7, 0x0f,
	0xbd, 0xc0, 0x7b, 0xff, 0x14, 0x37, 0x03, 0xff, 0xe0, 0x2e, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x31, 0x78, 0x89, 0x02, 0x49, 0x04, 0x00, 0x00,
}