// Code generated by protoc-gen-go. DO NOT EDIT.
// source: requestf/requestf.proto

package requestf

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

type ReqPacket struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ReqId                int32    `protobuf:"varint,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
	FuncName             string   `protobuf:"bytes,3,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Servant              string   `protobuf:"bytes,4,opt,name=servant,proto3" json:"servant,omitempty"`
	Content              []byte   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Timeout              int32    `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqPacket) Reset()         { *m = ReqPacket{} }
func (m *ReqPacket) String() string { return proto.CompactTextString(m) }
func (*ReqPacket) ProtoMessage()    {}
func (*ReqPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_009a0fa200021b02, []int{0}
}

func (m *ReqPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqPacket.Unmarshal(m, b)
}
func (m *ReqPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqPacket.Marshal(b, m, deterministic)
}
func (m *ReqPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqPacket.Merge(m, src)
}
func (m *ReqPacket) XXX_Size() int {
	return xxx_messageInfo_ReqPacket.Size(m)
}
func (m *ReqPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqPacket.DiscardUnknown(m)
}

var xxx_messageInfo_ReqPacket proto.InternalMessageInfo

func (m *ReqPacket) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReqPacket) GetReqId() int32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *ReqPacket) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *ReqPacket) GetServant() string {
	if m != nil {
		return m.Servant
	}
	return ""
}

func (m *ReqPacket) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ReqPacket) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type RspPacket struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ReqId                int32    `protobuf:"varint,2,opt,name=reqId,proto3" json:"reqId,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	IsErr                bool     `protobuf:"varint,4,opt,name=isErr,proto3" json:"isErr,omitempty"`
	ResDesc              string   `protobuf:"bytes,5,opt,name=resDesc,proto3" json:"resDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspPacket) Reset()         { *m = RspPacket{} }
func (m *RspPacket) String() string { return proto.CompactTextString(m) }
func (*RspPacket) ProtoMessage()    {}
func (*RspPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_009a0fa200021b02, []int{1}
}

func (m *RspPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspPacket.Unmarshal(m, b)
}
func (m *RspPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspPacket.Marshal(b, m, deterministic)
}
func (m *RspPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspPacket.Merge(m, src)
}
func (m *RspPacket) XXX_Size() int {
	return xxx_messageInfo_RspPacket.Size(m)
}
func (m *RspPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_RspPacket.DiscardUnknown(m)
}

var xxx_messageInfo_RspPacket proto.InternalMessageInfo

func (m *RspPacket) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RspPacket) GetReqId() int32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *RspPacket) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RspPacket) GetIsErr() bool {
	if m != nil {
		return m.IsErr
	}
	return false
}

func (m *RspPacket) GetResDesc() string {
	if m != nil {
		return m.ResDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqPacket)(nil), "requestf.ReqPacket")
	proto.RegisterType((*RspPacket)(nil), "requestf.RspPacket")
}

func init() { proto.RegisterFile("requestf/requestf.proto", fileDescriptor_009a0fa200021b02) }

var fileDescriptor_009a0fa200021b02 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd0, 0x31, 0x6e, 0x83, 0x30,
	0x14, 0x06, 0x60, 0xb9, 0xd4, 0x14, 0xac, 0x4e, 0x16, 0x52, 0xad, 0x4e, 0x88, 0x89, 0xa9, 0x1d,
	0x7a, 0x85, 0x76, 0xe8, 0x52, 0x55, 0xbe, 0x01, 0x71, 0x1e, 0x12, 0x8a, 0xb0, 0xe1, 0xf9, 0xc1,
	0x0d, 0x72, 0x94, 0xdc, 0x33, 0xb2, 0x89, 0xa3, 0x64, 0xcd, 0xc6, 0xc7, 0x6f, 0xf9, 0xff, 0x65,
	0xf1, 0x86, 0x30, 0x2f, 0xe0, 0xa9, 0xff, 0x4c, 0x1f, 0x1f, 0x13, 0x3a, 0x72, 0xb2, 0x48, 0x6e,
	0x4e, 0x4c, 0x94, 0x1a, 0xe6, 0xff, 0xce, 0x1c, 0x80, 0xa4, 0x12, 0x2f, 0x2b, 0xa0, 0x1f, 0x9c,
	0x55, 0xac, 0x66, 0x2d, 0xd7, 0x89, 0xb2, 0x12, 0x1c, 0x61, 0xfe, 0xdd, 0xab, 0xa7, 0xf8, 0x7f,
	0x83, 0x7c, 0x17, 0x45, 0xbf, 0x58, 0xf3, 0xd7, 0x8d, 0xa0, 0xb2, 0x9a, 0xb5, 0xa5, 0xbe, 0x3a,
	0xdc, 0xe5, 0x01, 0xd7, 0xce, 0x92, 0x7a, 0x8e, 0x51, 0x62, 0x48, 0x8c, 0xb3, 0x04, 0x96, 0x14,
	0xaf, 0x59, 0xfb, 0xaa, 0x13, 0x43, 0x42, 0xc3, 0x08, 0x6e, 0x21, 0x95, 0x6f, 0xfd, 0x17, 0x36,
	0xc7, 0xb0, 0xd3, 0x4f, 0x0f, 0xee, 0xbc, 0x69, 0xcc, 0xee, 0x1b, 0x2b, 0xc1, 0x07, 0xff, 0x83,
	0x18, 0x37, 0x16, 0x7a, 0x43, 0x38, 0x8f, 0xe0, 0xbf, 0xc1, 0x9b, 0xb8, 0xb0, 0xd4, 0x89, 0xbb,
	0x3c, 0x3e, 0xe0, 0xd7, 0x39, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x54, 0x81, 0x13, 0x5b, 0x01, 0x00,
	0x00,
}

// fffffffffffffffff
