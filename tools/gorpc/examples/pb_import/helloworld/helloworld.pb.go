// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld2019

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

type Love struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Love) Reset()         { *m = Love{} }
func (m *Love) String() string { return proto.CompactTextString(m) }
func (*Love) ProtoMessage()    {}
func (*Love) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *Love) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Love.Unmarshal(m, b)
}
func (m *Love) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Love.Marshal(b, m, deterministic)
}
func (m *Love) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Love.Merge(m, src)
}
func (m *Love) XXX_Size() int {
	return xxx_messageInfo_Love.Size(m)
}
func (m *Love) XXX_DiscardUnknown() {
	xxx_messageInfo_Love.DiscardUnknown(m)
}

var xxx_messageInfo_Love proto.InternalMessageInfo

type Hate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hate) Reset()         { *m = Hate{} }
func (m *Hate) String() string { return proto.CompactTextString(m) }
func (*Hate) ProtoMessage()    {}
func (*Hate) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *Hate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hate.Unmarshal(m, b)
}
func (m *Hate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hate.Marshal(b, m, deterministic)
}
func (m *Hate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hate.Merge(m, src)
}
func (m *Hate) XXX_Size() int {
	return xxx_messageInfo_Hate.Size(m)
}
func (m *Hate) XXX_DiscardUnknown() {
	xxx_messageInfo_Hate.DiscardUnknown(m)
}

var xxx_messageInfo_Hate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Love)(nil), "helloworld.Love")
	proto.RegisterType((*Hate)(nil), "helloworld.Hate")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 74 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xb1, 0x71, 0xb1, 0xf8, 0xe4, 0x97, 0xa5, 0x82, 0x68, 0x8f, 0xc4, 0x92, 0x54, 0x27, 0x81,
	0x28, 0x3e, 0x84, 0xac, 0x91, 0x81, 0xa1, 0x65, 0x12, 0x1b, 0x58, 0x93, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x19, 0xeb, 0x27, 0x8e, 0x48, 0x00, 0x00, 0x00,
}