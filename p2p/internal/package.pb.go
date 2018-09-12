// Code generated by protoc-gen-go. DO NOT EDIT.
// source: package.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Package struct {
	Anything             *any.Any `protobuf:"bytes,1,opt,name=anything,proto3" json:"anything,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae8103ff0e06fb71, []int{0}
}
func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (dst *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(dst, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetAnything() *any.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Package) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Package)(nil), "internal.Package")
}

func init() { proto.RegisterFile("package.proto", fileDescriptor_ae8103ff0e06fb71) }

var fileDescriptor_ae8103ff0e06fb71 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0x4c, 0xce,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0x91, 0x92, 0x05, 0x0b, 0x24, 0x95, 0xa6, 0xe9, 0x17, 0x94, 0x54, 0x16, 0xa4,
	0x16, 0xeb, 0x27, 0xe6, 0x55, 0x82, 0x30, 0x44, 0xa1, 0x52, 0x24, 0x17, 0x7b, 0x00, 0x44, 0xa7,
	0x90, 0x01, 0x17, 0x47, 0x62, 0x5e, 0x65, 0x49, 0x46, 0x66, 0x5e, 0xba, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0xb7, 0x91, 0x88, 0x5e, 0x7a, 0x7e, 0x7e, 0x7a, 0x0e, 0xd4, 0xd0, 0xa4, 0xd2, 0x34, 0x3d,
	0xc7, 0xbc, 0xca, 0x20, 0xb8, 0x2a, 0x21, 0x19, 0x2e, 0xce, 0xe2, 0xcc, 0xf4, 0xbc, 0xc4, 0x92,
	0xd2, 0xa2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x84, 0x40, 0x12, 0x1b, 0x58, 0x97,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x12, 0x34, 0xaf, 0x9b, 0x00, 0x00, 0x00,
}
