// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peerTalk.proto

package peerTalk

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Register struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ac8e6e2c44fa6b, []int{0}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Register) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Bootstrap struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bootstrap) Reset()         { *m = Bootstrap{} }
func (m *Bootstrap) String() string { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()    {}
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ac8e6e2c44fa6b, []int{1}
}

func (m *Bootstrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap.Unmarshal(m, b)
}
func (m *Bootstrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap.Marshal(b, m, deterministic)
}
func (m *Bootstrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap.Merge(m, src)
}
func (m *Bootstrap) XXX_Size() int {
	return xxx_messageInfo_Bootstrap.Size(m)
}
func (m *Bootstrap) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap proto.InternalMessageInfo

func (m *Bootstrap) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Grouping struct {
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Ids                  [][]byte `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Grouping) Reset()         { *m = Grouping{} }
func (m *Grouping) String() string { return proto.CompactTextString(m) }
func (*Grouping) ProtoMessage()    {}
func (*Grouping) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ac8e6e2c44fa6b, []int{2}
}

func (m *Grouping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grouping.Unmarshal(m, b)
}
func (m *Grouping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grouping.Marshal(b, m, deterministic)
}
func (m *Grouping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grouping.Merge(m, src)
}
func (m *Grouping) XXX_Size() int {
	return xxx_messageInfo_Grouping.Size(m)
}
func (m *Grouping) XXX_DiscardUnknown() {
	xxx_messageInfo_Grouping.DiscardUnknown(m)
}

var xxx_messageInfo_Grouping proto.InternalMessageInfo

func (m *Grouping) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Grouping) GetIds() [][]byte {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Done struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Done) Reset()         { *m = Done{} }
func (m *Done) String() string { return proto.CompactTextString(m) }
func (*Done) ProtoMessage()    {}
func (*Done) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ac8e6e2c44fa6b, []int{3}
}

func (m *Done) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Done.Unmarshal(m, b)
}
func (m *Done) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Done.Marshal(b, m, deterministic)
}
func (m *Done) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Done.Merge(m, src)
}
func (m *Done) XXX_Size() int {
	return xxx_messageInfo_Done.Size(m)
}
func (m *Done) XXX_DiscardUnknown() {
	xxx_messageInfo_Done.DiscardUnknown(m)
}

var xxx_messageInfo_Done proto.InternalMessageInfo

func (m *Done) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*Register)(nil), "peerTalk.Register")
	proto.RegisterType((*Bootstrap)(nil), "peerTalk.Bootstrap")
	proto.RegisterType((*Grouping)(nil), "peerTalk.Grouping")
	proto.RegisterType((*Done)(nil), "peerTalk.Done")
}

func init() { proto.RegisterFile("peerTalk.proto", fileDescriptor_87ac8e6e2c44fa6b) }

var fileDescriptor_87ac8e6e2c44fa6b = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0x4d, 0x2d,
	0x0a, 0x49, 0xcc, 0xc9, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x4c,
	0xb8, 0x38, 0x82, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x98, 0x32, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x13, 0x53,
	0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x55,
	0x2e, 0x4e, 0xa7, 0xfc, 0xfc, 0x92, 0xe2, 0x92, 0xa2, 0xc4, 0x02, 0x64, 0x65, 0x8c, 0xa8, 0xca,
	0x0c, 0xb8, 0x38, 0xdc, 0x8b, 0xf2, 0x4b, 0x0b, 0x32, 0xf3, 0xd2, 0x85, 0x84, 0xb8, 0x58, 0x8a,
	0x33, 0xab, 0x52, 0xc1, 0x4a, 0x78, 0x83, 0xc0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xcc, 0x14, 0x90,
	0xe1, 0xcc, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x92, 0x18, 0x17, 0x8b, 0x4b, 0x7e, 0x5e, 0x2a, 0xba,
	0x53, 0x92, 0xd8, 0xc0, 0xee, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x23, 0xaa, 0xf9,
	0xc9, 0x00, 0x00, 0x00,
}