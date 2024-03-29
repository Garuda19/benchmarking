// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delta.proto

package benchmarking

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

type Delta struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Chips                int64    `protobuf:"varint,3,opt,name=chips,proto3" json:"chips,omitempty"`
	Tracking             []string `protobuf:"bytes,4,rep,name=tracking,proto3" json:"tracking,omitempty"`
	InstallOS            string   `protobuf:"bytes,5,opt,name=installOS,proto3" json:"installOS,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ClearTimeStamp       bool     `protobuf:"varint,7,opt,name=clearTimeStamp,proto3" json:"clearTimeStamp,omitempty"`
	IsTournament         bool     `protobuf:"varint,8,opt,name=isTournament,proto3" json:"isTournament,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Delta) Reset()         { *m = Delta{} }
func (m *Delta) String() string { return proto.CompactTextString(m) }
func (*Delta) ProtoMessage()    {}
func (*Delta) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9f1fd9d66f078a0, []int{0}
}

func (m *Delta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delta.Unmarshal(m, b)
}
func (m *Delta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delta.Marshal(b, m, deterministic)
}
func (m *Delta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delta.Merge(m, src)
}
func (m *Delta) XXX_Size() int {
	return xxx_messageInfo_Delta.Size(m)
}
func (m *Delta) XXX_DiscardUnknown() {
	xxx_messageInfo_Delta.DiscardUnknown(m)
}

var xxx_messageInfo_Delta proto.InternalMessageInfo

func (m *Delta) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Delta) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Delta) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *Delta) GetTracking() []string {
	if m != nil {
		return m.Tracking
	}
	return nil
}

func (m *Delta) GetInstallOS() string {
	if m != nil {
		return m.InstallOS
	}
	return ""
}

func (m *Delta) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Delta) GetClearTimeStamp() bool {
	if m != nil {
		return m.ClearTimeStamp
	}
	return false
}

func (m *Delta) GetIsTournament() bool {
	if m != nil {
		return m.IsTournament
	}
	return false
}

func init() {
	proto.RegisterType((*Delta)(nil), "benchmarking.Delta")
}

func init() { proto.RegisterFile("delta.proto", fileDescriptor_c9f1fd9d66f078a0) }

var fileDescriptor_c9f1fd9d66f078a0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0x65, 0x2e, 0x17, 0xee, 0x96, 0x88, 0x62, 0x45, 0x61, 0x21, 0x0a, 0x2b, 0x05, 0x72,
	0x45, 0xc3, 0x2b, 0xd0, 0x50, 0x21, 0x39, 0x79, 0x01, 0xe7, 0xb2, 0x22, 0x16, 0xe7, 0x1f, 0xd9,
	0x4b, 0xc1, 0x73, 0xf3, 0x02, 0xc8, 0x8e, 0x94, 0x13, 0xe9, 0x76, 0xbe, 0x19, 0xcd, 0x48, 0x0b,
	0x77, 0x47, 0x9a, 0xd9, 0xbe, 0xa4, 0x1c, 0x39, 0xe2, 0xe6, 0x40, 0x61, 0x3a, 0x79, 0x9b, 0xbf,
	0x5c, 0xf8, 0xdc, 0xfe, 0x0a, 0xe8, 0xdf, 0xaa, 0x8b, 0x4f, 0x30, 0x16, 0x2a, 0xc5, 0xc5, 0xf0,
	0x7e, 0x94, 0x42, 0x09, 0x3d, 0x9a, 0x05, 0x20, 0xc2, 0x8a, 0x7f, 0x12, 0xc9, 0x1b, 0x25, 0x74,
	0x6f, 0xda, 0x8d, 0x0f, 0xd0, 0x4f, 0x27, 0x97, 0x8a, 0xec, 0x94, 0xd0, 0x9d, 0x39, 0x0b, 0x7c,
	0x84, 0x81, 0xb3, 0x9d, 0x6a, 0xbb, 0x5c, 0xa9, 0x4e, 0x8f, 0xe6, 0xa2, 0xeb, 0x86, 0x0b, 0x85,
	0xed, 0x3c, 0x7f, 0xec, 0x64, 0x7f, 0xde, 0xb8, 0x80, 0xea, 0xb2, 0xf3, 0x54, 0xd8, 0xfa, 0x24,
	0xd7, 0xad, 0x73, 0x01, 0xf8, 0x0c, 0xf7, 0xd3, 0x4c, 0x36, 0xef, 0x9d, 0xa7, 0x5d, 0x8b, 0xdc,
	0x2a, 0xa1, 0x07, 0x73, 0x45, 0x71, 0x0b, 0x1b, 0x57, 0xf6, 0xf1, 0x3b, 0x07, 0xeb, 0x29, 0xb0,
	0x1c, 0x5a, 0xea, 0x1f, 0x3b, 0xac, 0xdb, 0x2b, 0x5e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x58, 0x74, 0x0d, 0x19, 0x01, 0x00, 0x00,
}
