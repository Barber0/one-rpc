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
	Packettype           int32    `protobuf:"varint,2,opt,name=packettype,proto3" json:"packettype,omitempty"`
	ReqId                int32    `protobuf:"varint,3,opt,name=reqId,proto3" json:"reqId,omitempty"`
	FuncName             string   `protobuf:"bytes,4,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Servant              string   `protobuf:"bytes,5,opt,name=servant,proto3" json:"servant,omitempty"`
	Content              []byte   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Timeout              int32    `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
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

func (m *ReqPacket) GetPackettype() int32 {
	if m != nil {
		return m.Packettype
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
	Packettype           int32    `protobuf:"varint,2,opt,name=packettype,proto3" json:"packettype,omitempty"`
	ReqId                int32    `protobuf:"varint,3,opt,name=reqId,proto3" json:"reqId,omitempty"`
	IRet                 string   `protobuf:"bytes,4,opt,name=iRet,proto3" json:"iRet,omitempty"`
	Content              []byte   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	ResDesc              int32    `protobuf:"varint,6,opt,name=resDesc,proto3" json:"resDesc,omitempty"`
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

func (m *RspPacket) GetPackettype() int32 {
	if m != nil {
		return m.Packettype
	}
	return 0
}

func (m *RspPacket) GetReqId() int32 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *RspPacket) GetIRet() string {
	if m != nil {
		return m.IRet
	}
	return ""
}

func (m *RspPacket) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RspPacket) GetResDesc() int32 {
	if m != nil {
		return m.ResDesc
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqPacket)(nil), "requestf.ReqPacket")
	proto.RegisterType((*RspPacket)(nil), "requestf.RspPacket")
}

func init() { proto.RegisterFile("requestf/requestf.proto", fileDescriptor_009a0fa200021b02) }

var fileDescriptor_009a0fa200021b02 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0x41, 0x4a, 0xc5, 0x30,
	0x14, 0x45, 0x89, 0xfe, 0xfc, 0xff, 0xfb, 0x70, 0x14, 0x04, 0x1f, 0x0e, 0xa4, 0x74, 0xd4, 0x91,
	0x0e, 0xdc, 0x82, 0x13, 0x27, 0x22, 0xd9, 0x41, 0x8d, 0xaf, 0x50, 0xa4, 0x49, 0x9a, 0xbc, 0x16,
	0x5c, 0x8f, 0x6b, 0x71, 0x5f, 0x92, 0xb4, 0x91, 0x82, 0x63, 0x67, 0xf7, 0xe4, 0x04, 0x72, 0x6f,
	0xe0, 0x26, 0xd0, 0x34, 0x53, 0xe4, 0xfe, 0xa1, 0x84, 0x7b, 0x1f, 0x1c, 0x3b, 0x75, 0x2e, 0xdc,
	0x7c, 0x0b, 0xa8, 0x34, 0x4d, 0xaf, 0x9d, 0xf9, 0x20, 0x56, 0x08, 0xa7, 0x85, 0x42, 0x1c, 0x9c,
	0x45, 0x51, 0x8b, 0x56, 0xea, 0x82, 0xea, 0x0e, 0xc0, 0xe7, 0x3b, 0xfc, 0xe9, 0x09, 0x2f, 0xb2,
	0xdc, 0x9d, 0xa8, 0x6b, 0x90, 0x81, 0xa6, 0xe7, 0x77, 0xbc, 0xcc, 0x6a, 0x05, 0x75, 0x0b, 0xe7,
	0x7e, 0xb6, 0xe6, 0xa5, 0x1b, 0x09, 0x0f, 0xb5, 0x68, 0x2b, 0xfd, 0xcb, 0xe9, 0xad, 0x48, 0x61,
	0xe9, 0x2c, 0xa3, 0xcc, 0xaa, 0x60, 0x32, 0xc6, 0x59, 0x26, 0xcb, 0x78, 0xac, 0x45, 0x7b, 0xa5,
	0x0b, 0x26, 0xc3, 0xc3, 0x48, 0x6e, 0x66, 0x3c, 0xad, 0xfd, 0x36, 0x6c, 0xbe, 0xd2, 0x8e, 0xe8,
	0xff, 0x69, 0x87, 0x82, 0xc3, 0xa0, 0x89, 0xb7, 0x0d, 0x39, 0xef, 0x5b, 0xca, 0x3f, 0x2d, 0x03,
	0xc5, 0x27, 0x8a, 0x26, 0xf7, 0x97, 0xba, 0xe0, 0xdb, 0x31, 0x7f, 0xff, 0xe3, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0xe5, 0x9f, 0x7a, 0x99, 0x01, 0x00, 0x00,
}

// fffffffffffffffff
