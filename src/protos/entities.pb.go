// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entities.proto

package entities

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Country struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}
func (*Country) Descriptor() ([]byte, []int) {
	return fileDescriptor_entities_3e6a3a4703cf427e, []int{0}
}
func (m *Country) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Country.Unmarshal(m, b)
}
func (m *Country) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Country.Marshal(b, m, deterministic)
}
func (dst *Country) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Country.Merge(dst, src)
}
func (m *Country) XXX_Size() int {
	return xxx_messageInfo_Country.Size(m)
}
func (m *Country) XXX_DiscardUnknown() {
	xxx_messageInfo_Country.DiscardUnknown(m)
}

var xxx_messageInfo_Country proto.InternalMessageInfo

func (m *Country) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Country) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GetCountriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCountriesRequest) Reset()         { *m = GetCountriesRequest{} }
func (m *GetCountriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCountriesRequest) ProtoMessage()    {}
func (*GetCountriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_entities_3e6a3a4703cf427e, []int{1}
}
func (m *GetCountriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCountriesRequest.Unmarshal(m, b)
}
func (m *GetCountriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCountriesRequest.Marshal(b, m, deterministic)
}
func (dst *GetCountriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCountriesRequest.Merge(dst, src)
}
func (m *GetCountriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCountriesRequest.Size(m)
}
func (m *GetCountriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCountriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCountriesRequest proto.InternalMessageInfo

type GetCountriesReply struct {
	Countries            []*Country `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetCountriesReply) Reset()         { *m = GetCountriesReply{} }
func (m *GetCountriesReply) String() string { return proto.CompactTextString(m) }
func (*GetCountriesReply) ProtoMessage()    {}
func (*GetCountriesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_entities_3e6a3a4703cf427e, []int{2}
}
func (m *GetCountriesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCountriesReply.Unmarshal(m, b)
}
func (m *GetCountriesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCountriesReply.Marshal(b, m, deterministic)
}
func (dst *GetCountriesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCountriesReply.Merge(dst, src)
}
func (m *GetCountriesReply) XXX_Size() int {
	return xxx_messageInfo_GetCountriesReply.Size(m)
}
func (m *GetCountriesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCountriesReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetCountriesReply proto.InternalMessageInfo

func (m *GetCountriesReply) GetCountries() []*Country {
	if m != nil {
		return m.Countries
	}
	return nil
}

func init() {
	proto.RegisterType((*Country)(nil), "Country")
	proto.RegisterType((*GetCountriesRequest)(nil), "GetCountriesRequest")
	proto.RegisterType((*GetCountriesReply)(nil), "GetCountriesReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesReply, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesReply, error) {
	out := new(GetCountriesReply)
	err := c.cc.Invoke(ctx, "/Database/GetCountries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesReply, error)
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_GetCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).GetCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/GetCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).GetCountries(ctx, req.(*GetCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCountries",
			Handler:    _Database_GetCountries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entities.proto",
}

func init() { proto.RegisterFile("entities.proto", fileDescriptor_entities_3e6a3a4703cf427e) }

var fileDescriptor_entities_3e6a3a4703cf427e = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x2b, 0xc9,
	0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe4, 0x62, 0x77, 0xce,
	0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x10, 0x31,
	0x10, 0x5b, 0x49, 0x94, 0x4b, 0xd8, 0x3d, 0xb5, 0x04, 0xa2, 0x2b, 0x33, 0xb5, 0x38, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc9, 0x9a, 0x4b, 0x10, 0x55, 0xb8, 0x20, 0xa7, 0x52, 0x48, 0x8d,
	0x8b, 0x33, 0x19, 0x26, 0x22, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa1, 0x07, 0xb5, 0x30,
	0x08, 0x21, 0x65, 0xe4, 0xc6, 0xc5, 0xe1, 0x92, 0x58, 0x92, 0x98, 0x94, 0x58, 0x9c, 0x2a, 0x64,
	0xc5, 0xc5, 0x83, 0x6c, 0x90, 0x90, 0x88, 0x1e, 0x16, 0xeb, 0xa4, 0x84, 0xf4, 0x30, 0x6c, 0x53,
	0x62, 0x48, 0x62, 0x03, 0xfb, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0x44, 0x30, 0xbe,
	0xe7, 0x00, 0x00, 0x00,
}
