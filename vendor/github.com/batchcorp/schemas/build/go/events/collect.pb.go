// Code generated by protoc-gen-go. DO NOT EDIT.
// source: collect.proto

package events

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

// This type is used to convey info about a collection request; it includes
// auth details that enables the collector to associate intake events with a
// particular client's collection request.
type Collect struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collect) Reset()         { *m = Collect{} }
func (m *Collect) String() string { return proto.CompactTextString(m) }
func (*Collect) ProtoMessage()    {}
func (*Collect) Descriptor() ([]byte, []int) {
	return fileDescriptor_37ca5d531a1e5e54, []int{0}
}

func (m *Collect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collect.Unmarshal(m, b)
}
func (m *Collect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collect.Marshal(b, m, deterministic)
}
func (m *Collect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collect.Merge(m, src)
}
func (m *Collect) XXX_Size() int {
	return xxx_messageInfo_Collect.Size(m)
}
func (m *Collect) XXX_DiscardUnknown() {
	xxx_messageInfo_Collect.DiscardUnknown(m)
}

var xxx_messageInfo_Collect proto.InternalMessageInfo

func (m *Collect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Collect) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Collect)(nil), "events.Collect")
}

func init() { proto.RegisterFile("collect.proto", fileDescriptor_37ca5d531a1e5e54) }

var fileDescriptor_37ca5d531a1e5e54 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xc9,
	0x49, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b,
	0x29, 0x56, 0xd2, 0xe7, 0x62, 0x77, 0x86, 0x48, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7,
	0xe6, 0x49, 0x30, 0x81, 0x85, 0x20, 0x1c, 0x27, 0xbd, 0x28, 0x9d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xa4, 0xc4, 0x92, 0xe4, 0x8c, 0xe4, 0xfc, 0xa2, 0x02, 0xfd,
	0xe2, 0xe4, 0x8c, 0xd4, 0xdc, 0xc4, 0x62, 0xfd, 0xa4, 0xd2, 0xcc, 0x9c, 0x14, 0xfd, 0xf4, 0x7c,
	0x7d, 0x88, 0x05, 0x49, 0x6c, 0x60, 0xfb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x1f,
	0x3b, 0x90, 0x80, 0x00, 0x00, 0x00,
}