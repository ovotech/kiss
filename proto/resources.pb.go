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

type ListSecretsRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListSecretsRequest) Reset()         { *m = ListSecretsRequest{} }
func (m *ListSecretsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSecretsRequest) ProtoMessage()    {}
func (*ListSecretsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{5}
}

func (m *ListSecretsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSecretsRequest.Unmarshal(m, b)
}
func (m *ListSecretsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSecretsRequest.Marshal(b, m, deterministic)
}
func (m *ListSecretsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSecretsRequest.Merge(m, src)
}
func (m *ListSecretsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSecretsRequest.Size(m)
}
func (m *ListSecretsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSecretsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSecretsRequest proto.InternalMessageInfo

func (m *ListSecretsRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ListSecretsResponse struct {
	Secrets              []string `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSecretsResponse) Reset()         { *m = ListSecretsResponse{} }
func (m *ListSecretsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSecretsResponse) ProtoMessage()    {}
func (*ListSecretsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{6}
}

func (m *ListSecretsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSecretsResponse.Unmarshal(m, b)
}
func (m *ListSecretsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSecretsResponse.Marshal(b, m, deterministic)
}
func (m *ListSecretsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSecretsResponse.Merge(m, src)
}
func (m *ListSecretsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSecretsResponse.Size(m)
}
func (m *ListSecretsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSecretsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSecretsResponse proto.InternalMessageInfo

func (m *ListSecretsResponse) GetSecrets() []string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type BindSecretRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ServiceAccountName   string      `protobuf:"bytes,3,opt,name=serviceAccountName,proto3" json:"serviceAccountName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BindSecretRequest) Reset()         { *m = BindSecretRequest{} }
func (m *BindSecretRequest) String() string { return proto.CompactTextString(m) }
func (*BindSecretRequest) ProtoMessage()    {}
func (*BindSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{7}
}

func (m *BindSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindSecretRequest.Unmarshal(m, b)
}
func (m *BindSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindSecretRequest.Marshal(b, m, deterministic)
}
func (m *BindSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindSecretRequest.Merge(m, src)
}
func (m *BindSecretRequest) XXX_Size() int {
	return xxx_messageInfo_BindSecretRequest.Size(m)
}
func (m *BindSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindSecretRequest proto.InternalMessageInfo

func (m *BindSecretRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *BindSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BindSecretRequest) GetServiceAccountName() string {
	if m != nil {
		return m.ServiceAccountName
	}
	return ""
}

type BindSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindSecretResponse) Reset()         { *m = BindSecretResponse{} }
func (m *BindSecretResponse) String() string { return proto.CompactTextString(m) }
func (*BindSecretResponse) ProtoMessage()    {}
func (*BindSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{8}
}

func (m *BindSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindSecretResponse.Unmarshal(m, b)
}
func (m *BindSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindSecretResponse.Marshal(b, m, deterministic)
}
func (m *BindSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindSecretResponse.Merge(m, src)
}
func (m *BindSecretResponse) XXX_Size() int {
	return xxx_messageInfo_BindSecretResponse.Size(m)
}
func (m *BindSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindSecretResponse proto.InternalMessageInfo

type CreateSecretIAMPolicyRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateSecretIAMPolicyRequest) Reset()         { *m = CreateSecretIAMPolicyRequest{} }
func (m *CreateSecretIAMPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSecretIAMPolicyRequest) ProtoMessage()    {}
func (*CreateSecretIAMPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{9}
}

func (m *CreateSecretIAMPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretIAMPolicyRequest.Unmarshal(m, b)
}
func (m *CreateSecretIAMPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretIAMPolicyRequest.Marshal(b, m, deterministic)
}
func (m *CreateSecretIAMPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretIAMPolicyRequest.Merge(m, src)
}
func (m *CreateSecretIAMPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSecretIAMPolicyRequest.Size(m)
}
func (m *CreateSecretIAMPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretIAMPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretIAMPolicyRequest proto.InternalMessageInfo

func (m *CreateSecretIAMPolicyRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateSecretIAMPolicyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateSecretIAMPolicyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSecretIAMPolicyResponse) Reset()         { *m = CreateSecretIAMPolicyResponse{} }
func (m *CreateSecretIAMPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSecretIAMPolicyResponse) ProtoMessage()    {}
func (*CreateSecretIAMPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{10}
}

func (m *CreateSecretIAMPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretIAMPolicyResponse.Unmarshal(m, b)
}
func (m *CreateSecretIAMPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretIAMPolicyResponse.Marshal(b, m, deterministic)
}
func (m *CreateSecretIAMPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretIAMPolicyResponse.Merge(m, src)
}
func (m *CreateSecretIAMPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSecretIAMPolicyResponse.Size(m)
}
func (m *CreateSecretIAMPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretIAMPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretIAMPolicyResponse proto.InternalMessageInfo

type DeleteSecretRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteSecretRequest) Reset()         { *m = DeleteSecretRequest{} }
func (m *DeleteSecretRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSecretRequest) ProtoMessage()    {}
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{11}
}

func (m *DeleteSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSecretRequest.Unmarshal(m, b)
}
func (m *DeleteSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSecretRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSecretRequest.Merge(m, src)
}
func (m *DeleteSecretRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSecretRequest.Size(m)
}
func (m *DeleteSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSecretRequest proto.InternalMessageInfo

func (m *DeleteSecretRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DeleteSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSecretResponse) Reset()         { *m = DeleteSecretResponse{} }
func (m *DeleteSecretResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSecretResponse) ProtoMessage()    {}
func (*DeleteSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{12}
}

func (m *DeleteSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSecretResponse.Unmarshal(m, b)
}
func (m *DeleteSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSecretResponse.Marshal(b, m, deterministic)
}
func (m *DeleteSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSecretResponse.Merge(m, src)
}
func (m *DeleteSecretResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSecretResponse.Size(m)
}
func (m *DeleteSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSecretResponse proto.InternalMessageInfo

type DeleteSecretIAMPolicyRequest struct {
	Metadata             *ClientMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteSecretIAMPolicyRequest) Reset()         { *m = DeleteSecretIAMPolicyRequest{} }
func (m *DeleteSecretIAMPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSecretIAMPolicyRequest) ProtoMessage()    {}
func (*DeleteSecretIAMPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{13}
}

func (m *DeleteSecretIAMPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSecretIAMPolicyRequest.Unmarshal(m, b)
}
func (m *DeleteSecretIAMPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSecretIAMPolicyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSecretIAMPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSecretIAMPolicyRequest.Merge(m, src)
}
func (m *DeleteSecretIAMPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSecretIAMPolicyRequest.Size(m)
}
func (m *DeleteSecretIAMPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSecretIAMPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSecretIAMPolicyRequest proto.InternalMessageInfo

func (m *DeleteSecretIAMPolicyRequest) GetMetadata() *ClientMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DeleteSecretIAMPolicyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteSecretIAMPolicyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSecretIAMPolicyResponse) Reset()         { *m = DeleteSecretIAMPolicyResponse{} }
func (m *DeleteSecretIAMPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteSecretIAMPolicyResponse) ProtoMessage()    {}
func (*DeleteSecretIAMPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1b13971fe4c19d, []int{14}
}

func (m *DeleteSecretIAMPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSecretIAMPolicyResponse.Unmarshal(m, b)
}
func (m *DeleteSecretIAMPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSecretIAMPolicyResponse.Marshal(b, m, deterministic)
}
func (m *DeleteSecretIAMPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSecretIAMPolicyResponse.Merge(m, src)
}
func (m *DeleteSecretIAMPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteSecretIAMPolicyResponse.Size(m)
}
func (m *DeleteSecretIAMPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSecretIAMPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSecretIAMPolicyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientMeta)(nil), "kiss.resources.ClientMeta")
	proto.RegisterType((*PingRequest)(nil), "kiss.resources.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "kiss.resources.PingResponse")
	proto.RegisterType((*CreateSecretRequest)(nil), "kiss.resources.CreateSecretRequest")
	proto.RegisterType((*CreateSecretResponse)(nil), "kiss.resources.CreateSecretResponse")
	proto.RegisterType((*ListSecretsRequest)(nil), "kiss.resources.ListSecretsRequest")
	proto.RegisterType((*ListSecretsResponse)(nil), "kiss.resources.ListSecretsResponse")
	proto.RegisterType((*BindSecretRequest)(nil), "kiss.resources.BindSecretRequest")
	proto.RegisterType((*BindSecretResponse)(nil), "kiss.resources.BindSecretResponse")
	proto.RegisterType((*CreateSecretIAMPolicyRequest)(nil), "kiss.resources.CreateSecretIAMPolicyRequest")
	proto.RegisterType((*CreateSecretIAMPolicyResponse)(nil), "kiss.resources.CreateSecretIAMPolicyResponse")
	proto.RegisterType((*DeleteSecretRequest)(nil), "kiss.resources.DeleteSecretRequest")
	proto.RegisterType((*DeleteSecretResponse)(nil), "kiss.resources.DeleteSecretResponse")
	proto.RegisterType((*DeleteSecretIAMPolicyRequest)(nil), "kiss.resources.DeleteSecretIAMPolicyRequest")
	proto.RegisterType((*DeleteSecretIAMPolicyResponse)(nil), "kiss.resources.DeleteSecretIAMPolicyResponse")
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor_cf1b13971fe4c19d) }

var fileDescriptor_cf1b13971fe4c19d = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0x59, 0x07, 0xf4, 0x76, 0x1a, 0xe2, 0xb6, 0xa0, 0x28, 0x74, 0xda, 0xf0, 0x78, 0x98,
	0x10, 0xa4, 0xd2, 0x90, 0x78, 0xdf, 0x0a, 0x0f, 0x13, 0x1b, 0x9a, 0x5a, 0x21, 0x21, 0x78, 0xf2,
	0xbc, 0xab, 0xcd, 0xd0, 0xc6, 0x25, 0x76, 0x8a, 0xf8, 0x0a, 0xfe, 0x82, 0xef, 0x44, 0x49, 0x4a,
	0xe3, 0xc4, 0x59, 0x84, 0x54, 0xe5, 0x29, 0xf1, 0xf5, 0xf5, 0x39, 0xc7, 0x27, 0xe7, 0x2a, 0xf0,
	0x28, 0x22, 0xad, 0xe2, 0x48, 0x90, 0x0e, 0x16, 0x91, 0x32, 0x0a, 0x77, 0xbf, 0x4b, 0xad, 0x83,
	0x75, 0x95, 0xbd, 0x04, 0x18, 0xcf, 0x24, 0x85, 0xe6, 0x82, 0x0c, 0xc7, 0x21, 0x74, 0x43, 0x3e,
	0x27, 0xbd, 0xe0, 0x82, 0xbc, 0xf6, 0x41, 0xfb, 0xa8, 0x3b, 0xc9, 0x0b, 0xec, 0x3d, 0xf4, 0x2e,
	0x65, 0x78, 0x33, 0xa1, 0x1f, 0x31, 0x69, 0x83, 0x6f, 0xe1, 0xe1, 0x9c, 0x0c, 0xbf, 0xe6, 0x86,
	0xa7, 0xbd, 0xbd, 0x63, 0x3f, 0x28, 0xa2, 0x07, 0x39, 0xf4, 0x64, 0xdd, 0xcb, 0x76, 0x61, 0x27,
	0x83, 0xd1, 0x0b, 0x15, 0x6a, 0x62, 0x3f, 0xa1, 0x3f, 0x8e, 0x88, 0x1b, 0x9a, 0x92, 0x88, 0xc8,
	0x6c, 0x08, 0x8f, 0x08, 0x9d, 0x44, 0xb2, 0x77, 0x2f, 0x95, 0x9f, 0xbe, 0xe3, 0x00, 0xb6, 0x97,
	0x7c, 0x16, 0x93, 0xb7, 0x95, 0x16, 0xb3, 0x05, 0x7b, 0x0a, 0x83, 0x22, 0xf1, 0x4a, 0xd0, 0x39,
	0xe0, 0xb9, 0xd4, 0x26, 0xab, 0xea, 0x4d, 0xaf, 0x3b, 0x82, 0x7e, 0x01, 0x2d, 0x23, 0x41, 0x0f,
	0x1e, 0xe8, 0xac, 0xe4, 0xb5, 0x0f, 0xb6, 0x8e, 0xba, 0x93, 0x7f, 0x4b, 0xf6, 0xbb, 0x0d, 0x8f,
	0x4f, 0x65, 0x78, 0xdd, 0x9c, 0x1d, 0x01, 0xa0, 0xa6, 0x68, 0x29, 0x05, 0x9d, 0x08, 0xa1, 0xe2,
	0xd0, 0x7c, 0x4c, 0x3a, 0x32, 0x6f, 0x2a, 0x76, 0xd8, 0x00, 0xd0, 0x16, 0xb4, 0xb2, 0xe9, 0x1b,
	0x0c, 0x6d, 0xfb, 0xce, 0x4e, 0x2e, 0x2e, 0xd5, 0x4c, 0x8a, 0x5f, 0x0d, 0x28, 0x66, 0xfb, 0xb0,
	0x77, 0x07, 0xd7, 0x4a, 0x0c, 0x87, 0xfe, 0x3b, 0x9a, 0x51, 0x83, 0x21, 0x4a, 0xe2, 0x52, 0xa4,
	0xc8, 0x7d, 0xb0, 0xeb, 0x4d, 0xfb, 0x70, 0x07, 0x57, 0x26, 0xe6, 0xf8, 0xcf, 0x36, 0x74, 0x3e,
	0x9c, 0x4d, 0xa7, 0x38, 0x86, 0x4e, 0x32, 0x65, 0xf8, 0xac, 0xcc, 0x65, 0x8d, 0xb0, 0x3f, 0xac,
	0xde, 0x5c, 0x5d, 0xac, 0x85, 0x5f, 0x61, 0xc7, 0xb6, 0x1d, 0x0f, 0x1d, 0xe1, 0xee, 0xe0, 0xfa,
	0x2f, 0xea, 0x9b, 0xd6, 0xe0, 0x9f, 0xa1, 0x67, 0x0d, 0x06, 0xb2, 0xf2, 0x31, 0x77, 0x06, 0xfd,
	0xc3, 0xda, 0x1e, 0x5b, 0xb6, 0xed, 0x92, 0x2b, 0xbb, 0x22, 0x2a, 0xae, 0xec, 0xca, 0x8f, 0xdd,
	0xc2, 0x4f, 0x00, 0xf9, 0x30, 0xe0, 0xf3, 0xf2, 0x29, 0x67, 0x72, 0x7d, 0x56, 0xd7, 0xb2, 0x86,
	0x5d, 0xc2, 0x93, 0xca, 0x84, 0xe3, 0xab, 0x3a, 0x3b, 0xcb, 0x61, 0xf3, 0x5f, 0xff, 0x67, 0xb7,
	0xcd, 0x5b, 0x99, 0x28, 0x97, 0xb7, 0x2e, 0xe4, 0x2e, 0x6f, 0x6d, 0x4c, 0x59, 0xeb, 0x74, 0xff,
	0xcb, 0xde, 0x8d, 0x34, 0xb7, 0xf1, 0x55, 0x20, 0xd4, 0x7c, 0xa4, 0x96, 0xca, 0x90, 0xb8, 0x1d,
	0x25, 0x20, 0xa3, 0xf4, 0x4f, 0x75, 0x75, 0x3f, 0x7d, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x3d, 0x30, 0xf2, 0xc3, 0x06, 0x00, 0x00,
}
