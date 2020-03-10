// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/setup.proto

package protocol

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

type CheckSetupRes struct {
	Setup                bool     `protobuf:"varint,1,opt,name=setup,proto3" json:"setup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckSetupRes) Reset()         { *m = CheckSetupRes{} }
func (m *CheckSetupRes) String() string { return proto.CompactTextString(m) }
func (*CheckSetupRes) ProtoMessage()    {}
func (*CheckSetupRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4214ec315ad003e, []int{0}
}

func (m *CheckSetupRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSetupRes.Unmarshal(m, b)
}
func (m *CheckSetupRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSetupRes.Marshal(b, m, deterministic)
}
func (m *CheckSetupRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSetupRes.Merge(m, src)
}
func (m *CheckSetupRes) XXX_Size() int {
	return xxx_messageInfo_CheckSetupRes.Size(m)
}
func (m *CheckSetupRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSetupRes.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSetupRes proto.InternalMessageInfo

func (m *CheckSetupRes) GetSetup() bool {
	if m != nil {
		return m.Setup
	}
	return false
}

type SetupParams struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Org                  string   `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	Bucket               string   `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	RetentionPeriodHrs   int64    `protobuf:"varint,5,opt,name=retentionPeriodHrs,proto3" json:"retentionPeriodHrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupParams) Reset()         { *m = SetupParams{} }
func (m *SetupParams) String() string { return proto.CompactTextString(m) }
func (*SetupParams) ProtoMessage()    {}
func (*SetupParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4214ec315ad003e, []int{1}
}

func (m *SetupParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupParams.Unmarshal(m, b)
}
func (m *SetupParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupParams.Marshal(b, m, deterministic)
}
func (m *SetupParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupParams.Merge(m, src)
}
func (m *SetupParams) XXX_Size() int {
	return xxx_messageInfo_SetupParams.Size(m)
}
func (m *SetupParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupParams.DiscardUnknown(m)
}

var xxx_messageInfo_SetupParams proto.InternalMessageInfo

func (m *SetupParams) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SetupParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SetupParams) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

func (m *SetupParams) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SetupParams) GetRetentionPeriodHrs() int64 {
	if m != nil {
		return m.RetentionPeriodHrs
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckSetupRes)(nil), "protocol.CheckSetupRes")
	proto.RegisterType((*SetupParams)(nil), "protocol.SetupParams")
}

func init() { proto.RegisterFile("protocol/setup.proto", fileDescriptor_f4214ec315ad003e) }

var fileDescriptor_f4214ec315ad003e = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x8a, 0x83, 0x40,
	0x10, 0x86, 0xd9, 0xf3, 0x14, 0x6f, 0x8e, 0x83, 0x63, 0x91, 0xb0, 0xa4, 0x12, 0x21, 0x60, 0x65,
	0x8a, 0x3c, 0x42, 0x9a, 0x94, 0xb2, 0x79, 0x82, 0x55, 0x87, 0x44, 0x8c, 0xae, 0xcc, 0xae, 0xe4,
	0x71, 0xf2, 0xaa, 0xc1, 0x31, 0xa6, 0x4a, 0x37, 0xdf, 0xff, 0x7d, 0xcd, 0x40, 0x32, 0x92, 0xf5,
	0xb6, 0xb6, 0xb7, 0xbd, 0x43, 0x3f, 0x8d, 0x05, 0xa3, 0x8c, 0xd7, 0x35, 0xdb, 0xc1, 0xdf, 0xf1,
	0x8a, 0x75, 0x77, 0x9e, 0xad, 0x46, 0x27, 0x13, 0x08, 0xb9, 0x54, 0x22, 0x15, 0x79, 0xac, 0x17,
	0xc8, 0x1e, 0x02, 0x7e, 0x39, 0x29, 0x0d, 0x99, 0xde, 0xc9, 0x2d, 0xc4, 0x93, 0x43, 0x1a, 0x4c,
	0x8f, 0x1c, 0xfe, 0xe8, 0x37, 0xcf, 0x6e, 0x34, 0xce, 0xdd, 0x2d, 0x35, 0xea, 0x6b, 0x71, 0x2b,
	0xcb, 0x7f, 0x08, 0x2c, 0x5d, 0x54, 0xc0, 0xf3, 0x7c, 0xca, 0x0d, 0x44, 0xd5, 0x54, 0x77, 0xe8,
	0xd5, 0x37, 0x8f, 0x2f, 0x92, 0x05, 0x48, 0x42, 0x8f, 0x83, 0x6f, 0xed, 0x50, 0x22, 0xb5, 0xb6,
	0x39, 0x91, 0x53, 0x61, 0x2a, 0xf2, 0x40, 0x7f, 0x30, 0x55, 0xc4, 0x2f, 0x1d, 0x9e, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0xa1, 0xd7, 0x6e, 0xf1, 0x00, 0x00, 0x00,
}