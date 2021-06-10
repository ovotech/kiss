// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

package proto

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

// Required information from client for every request
type ClientMeta struct {
	// k8s "namespace" the client wants to operate _for_
	// nb: we're not actually interacting with k8s here
	// we use tags on AWS resources to control which k8s namespace secrets are _for_
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientMeta) Reset()         { *m = ClientMeta{} }
func (m *ClientMeta) String() string { return proto.CompactTextString(m) }
func (*ClientMeta) ProtoMessage()    {}
func (*ClientMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{0}
}

func (m *ClientMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMeta.Unmarshal(m, b)
}
func (m *ClientMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMeta.Marshal(b, m, deterministic)
}
func (m *ClientMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMeta.Merge(m, src)
}
func (m *ClientMeta) XXX_Size() int {
	return xxx_messageInfo_ClientMeta.Size(m)
}
func (m *ClientMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMeta proto.InternalMessageInfo

func (m *ClientMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type PingRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{1}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{2}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

type CreateSecretRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string      `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateSecretRequest) Reset()         { *m = CreateSecretRequest{} }
func (m *CreateSecretRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSecretRequest) ProtoMessage()    {}
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{3}
}

func (m *CreateSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretRequest.Unmarshal(m, b)
}
func (m *CreateSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretRequest.Marshal(b, m, deterministic)
}
func (m *CreateSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretRequest.Merge(m, src)
}
func (m *CreateSecretRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSecretRequest.Size(m)
}
func (m *CreateSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretRequest proto.InternalMessageInfo

func (m *CreateSecretRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateSecretRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSecretResponse) Reset()         { *m = CreateSecretResponse{} }
func (m *CreateSecretResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSecretResponse) ProtoMessage()    {}
func (*CreateSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{4}
}

func (m *CreateSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretResponse.Unmarshal(m, b)
}
func (m *CreateSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretResponse.Marshal(b, m, deterministic)
}
func (m *CreateSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretResponse.Merge(m, src)
}
func (m *CreateSecretResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSecretResponse.Size(m)
}
func (m *CreateSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientMeta)(nil), "kiss.resources.ClientMeta")
	proto.RegisterType((*PingRequest)(nil), "kiss.resources.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "kiss.resources.PingResponse")
	proto.RegisterType((*CreateSecretRequest)(nil), "kiss.resources.CreateSecretRequest")
	proto.RegisterType((*CreateSecretResponse)(nil), "kiss.resources.CreateSecretResponse")
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor_cf1b13971fe4c19d) }

var fileDescriptor_cf1b13971fe4c19d = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x5d, 0x4b, 0xf3, 0x30,
	0x14, 0x7e, 0xfb, 0x5a, 0xc5, 0x9d, 0x8d, 0x09, 0xc7, 0x21, 0xa5, 0x4e, 0x94, 0xea, 0x85, 0x78,
	0x91, 0xc2, 0x04, 0x7f, 0x80, 0xc5, 0x0b, 0x11, 0x41, 0xda, 0x3b, 0xbd, 0xca, 0xe2, 0x61, 0x2b,
	0xae, 0x4d, 0x4d, 0x4e, 0xe7, 0x5f, 0xf2, 0x67, 0x4a, 0x53, 0xdd, 0x87, 0x13, 0x6f, 0xbc, 0x4a,
	0xf2, 0xe4, 0xc9, 0xf3, 0x91, 0x03, 0x7b, 0x86, 0xac, 0xae, 0x8d, 0x22, 0x2b, 0x2a, 0xa3, 0x59,
	0x63, 0xff, 0x25, 0xb7, 0x56, 0x2c, 0xd0, 0xe8, 0x02, 0x20, 0x99, 0xe5, 0x54, 0xf2, 0x3d, 0xb1,
	0xc4, 0x21, 0x74, 0x4a, 0x59, 0x90, 0xad, 0xa4, 0xa2, 0xc0, 0x3b, 0xf1, 0xce, 0x3b, 0xe9, 0x12,
	0x88, 0x6e, 0xa0, 0xfb, 0x90, 0x97, 0x93, 0x94, 0x5e, 0x6b, 0xb2, 0x8c, 0x57, 0xb0, 0x5b, 0x10,
	0xcb, 0x67, 0xc9, 0xd2, 0x71, 0xbb, 0xa3, 0x50, 0xac, 0xab, 0x8b, 0xa5, 0x74, 0xba, 0xe0, 0x46,
	0x7d, 0xe8, 0xb5, 0x32, 0xb6, 0xd2, 0xa5, 0xa5, 0xe8, 0x0d, 0xf6, 0x13, 0x43, 0x92, 0x29, 0x23,
	0x65, 0x88, 0xff, 0x28, 0x8f, 0x08, 0x7e, 0x13, 0x39, 0xf8, 0xef, 0xe2, 0xbb, 0x3d, 0x0e, 0x60,
	0x7b, 0x2e, 0x67, 0x35, 0x05, 0x5b, 0x0e, 0x6c, 0x0f, 0xd1, 0x01, 0x0c, 0xd6, 0x8d, 0xdb, 0x40,
	0xa3, 0x77, 0x0f, 0xfc, 0xbb, 0xdb, 0x2c, 0xc3, 0x04, 0xfc, 0x26, 0x29, 0x1e, 0x7e, 0x37, 0x5e,
	0xf9, 0x86, 0x70, 0xf8, 0xf3, 0xe5, 0x67, 0xb9, 0x7f, 0xf8, 0x04, 0xbd, 0x55, 0x17, 0x3c, 0xdd,
	0x68, 0xb1, 0x59, 0x3e, 0x3c, 0xfb, 0x9d, 0xf4, 0x25, 0x7e, 0x7d, 0xfc, 0x78, 0x34, 0xc9, 0x79,
	0x5a, 0x8f, 0x85, 0xd2, 0x45, 0xac, 0xe7, 0x9a, 0x49, 0x4d, 0xe3, 0xe6, 0x6d, 0xec, 0xe6, 0x3d,
	0xde, 0x71, 0xcb, 0xe5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xa9, 0x8f, 0x0f, 0x09, 0x02,
	0x00, 0x00,
}
