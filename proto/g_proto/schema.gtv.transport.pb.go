// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.gtv.transport.proto

package g_proto

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

type GTVRpcError struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrorCode            int32    `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	WaitFor              int32    `protobuf:"varint,4,opt,name=wait_for,json=waitFor,proto3" json:"wait_for,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GTVRpcError) Reset()         { *m = GTVRpcError{} }
func (m *GTVRpcError) String() string { return proto.CompactTextString(m) }
func (*GTVRpcError) ProtoMessage()    {}
func (*GTVRpcError) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f4fa8781c73c6d, []int{0}
}

func (m *GTVRpcError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GTVRpcError.Unmarshal(m, b)
}
func (m *GTVRpcError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GTVRpcError.Marshal(b, m, deterministic)
}
func (m *GTVRpcError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GTVRpcError.Merge(m, src)
}
func (m *GTVRpcError) XXX_Size() int {
	return xxx_messageInfo_GTVRpcError.Size(m)
}
func (m *GTVRpcError) XXX_DiscardUnknown() {
	xxx_messageInfo_GTVRpcError.DiscardUnknown(m)
}

var xxx_messageInfo_GTVRpcError proto.InternalMessageInfo

func (m *GTVRpcError) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GTVRpcError) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *GTVRpcError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GTVRpcError) GetWaitFor() int32 {
	if m != nil {
		return m.WaitFor
	}
	return 0
}

func init() {
	proto.RegisterType((*GTVRpcError)(nil), "g_proto.GTVRpcError")
}

func init() { proto.RegisterFile("schema.gtv.transport.proto", fileDescriptor_d1f4fa8781c73c6d) }

var fileDescriptor_d1f4fa8781c73c6d = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x4b, 0x2f, 0x29, 0xd3, 0x2b, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f, 0x2a,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x8f, 0x07, 0x33, 0x94, 0x4a, 0xb9, 0xb8,
	0xdd, 0x43, 0xc2, 0x82, 0x0a, 0x92, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0x92,
	0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x59, 0x2e, 0xae,
	0x54, 0x90, 0x64, 0x3c, 0x58, 0x86, 0x09, 0x2c, 0xc3, 0x09, 0x16, 0x71, 0x06, 0x49, 0x4b, 0x70,
	0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0xc1,
	0xb8, 0x42, 0x92, 0x5c, 0x1c, 0xe5, 0x89, 0x99, 0x25, 0xf1, 0x69, 0xf9, 0x45, 0x12, 0x2c, 0x60,
	0x6d, 0xec, 0x20, 0xbe, 0x5b, 0x7e, 0x91, 0x93, 0x34, 0x17, 0x5f, 0x59, 0x1e, 0xd8, 0x65, 0x50,
	0x87, 0x38, 0xb1, 0xbb, 0xc7, 0x07, 0x80, 0x18, 0x1e, 0x4c, 0x49, 0x6c, 0x60, 0x11, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xdb, 0x15, 0x7b, 0xc1, 0x00, 0x00, 0x00,
}
