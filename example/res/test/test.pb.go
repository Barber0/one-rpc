// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/test.proto

package test

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

type Obj struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Obj) Reset()         { *m = Obj{} }
func (m *Obj) String() string { return proto.CompactTextString(m) }
func (*Obj) ProtoMessage()    {}
func (*Obj) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eb23d74a64bdab, []int{0}
}

func (m *Obj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Obj.Unmarshal(m, b)
}
func (m *Obj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Obj.Marshal(b, m, deterministic)
}
func (m *Obj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Obj.Merge(m, src)
}
func (m *Obj) XXX_Size() int {
	return xxx_messageInfo_Obj.Size(m)
}
func (m *Obj) XXX_DiscardUnknown() {
	xxx_messageInfo_Obj.DiscardUnknown(m)
}

var xxx_messageInfo_Obj proto.InternalMessageInfo

func (m *Obj) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Obj)(nil), "test.Obj")
}

func init() { proto.RegisterFile("test/test.proto", fileDescriptor_84eb23d74a64bdab) }

var fileDescriptor_84eb23d74a64bdab = []byte{
	// 71 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x49, 0x2d, 0x2e,
	0xd1, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x24, 0x17,
	0xb3, 0x7f, 0x52, 0x96, 0x90, 0x10, 0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x6b, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0x56, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1b,
	0xb6, 0x2d, 0x26, 0x3a, 0x00, 0x00, 0x00,
}

// fffffffffffffffff