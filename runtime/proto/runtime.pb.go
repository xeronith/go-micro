// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime.proto

package go_micro_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Service struct {
	// name of the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// git url of the source
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// local path of the source
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// command to execute
	Exec                 []string `protobuf:"bytes,4,rep,name=exec,proto3" json:"exec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Service) GetExec() []string {
	if m != nil {
		return m.Exec
	}
	return nil
}

type CreateRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{4}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.runtime.Service")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.runtime.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.runtime.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.runtime.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.runtime.DeleteResponse")
}

func init() { proto.RegisterFile("runtime.proto", fileDescriptor_86e2dd377c869464) }

var fileDescriptor_86e2dd377c869464 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xbf, 0x4e, 0x86, 0x30,
	0x14, 0xc5, 0x41, 0x08, 0xc4, 0x6b, 0x30, 0xa4, 0x83, 0xa9, 0x2e, 0x92, 0x4e, 0x4e, 0x1d, 0xe0,
	0x11, 0x64, 0x75, 0xa9, 0xb3, 0x03, 0x36, 0x37, 0x4a, 0x22, 0x14, 0xdb, 0x62, 0x7c, 0x23, 0x5f,
	0xd3, 0xf4, 0x0f, 0x03, 0xca, 0x37, 0x7d, 0xdb, 0xe9, 0xc9, 0xc9, 0xef, 0x9c, 0x9b, 0x42, 0xa5,
	0xd7, 0xd9, 0x8e, 0x13, 0xf2, 0x45, 0x2b, 0xab, 0x48, 0xfd, 0xa6, 0xf8, 0x34, 0x4a, 0xad, 0x78,
	0xf4, 0xd9, 0x0b, 0x94, 0xcf, 0xa8, 0xbf, 0x46, 0x89, 0x84, 0x40, 0x3e, 0x0f, 0x13, 0xd2, 0xb4,
	0x49, 0x1f, 0x2e, 0x85, 0xd7, 0xe4, 0x06, 0x0a, 0xa3, 0x56, 0x2d, 0x91, 0x5e, 0x78, 0x37, 0xbe,
	0x5c, 0x76, 0x19, 0xec, 0x3b, 0xcd, 0x42, 0xd6, 0x69, 0xe7, 0xe1, 0x37, 0x4a, 0x9a, 0x37, 0x99,
	0xf3, 0x9c, 0x66, 0x3d, 0x54, 0x8f, 0x1a, 0x07, 0x8b, 0x02, 0x3f, 0x57, 0x34, 0x96, 0x74, 0x50,
	0x9a, 0xd0, 0xe7, 0x7b, 0xae, 0xda, 0x5b, 0xfe, 0x77, 0x13, 0x8f, 0x83, 0xc4, 0x96, 0x64, 0x35,
	0x5c, 0x6f, 0x14, 0xb3, 0xa8, 0xd9, 0xa0, 0xe3, 0xf6, 0xf8, 0x81, 0xe7, 0x73, 0x37, 0x4a, 0xe0,
	0xb6, 0x3f, 0x29, 0x94, 0x22, 0xc4, 0xc9, 0x13, 0x14, 0xa1, 0x95, 0xdc, 0xff, 0x67, 0xed, 0xae,
	0xba, 0x6b, 0x4e, 0x07, 0xe2, 0xe0, 0xc4, 0xe1, 0x42, 0xd9, 0x11, 0x6e, 0x77, 0xcc, 0x11, 0x6e,
	0xbf, 0x93, 0x25, 0xaf, 0x85, 0xff, 0xd1, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x99, 0xc5, 0x97,
	0x12, 0xe2, 0x01, 0x00, 0x00,
}
