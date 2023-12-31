// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/guessing-game.proto

package proto

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

type GuessStatus int32

const (
	GuessStatus_GUESS_NONE     GuessStatus = 0
	GuessStatus_GUESS_TOO_LOW  GuessStatus = 1
	GuessStatus_GUESS_TOO_HIGH GuessStatus = 2
	GuessStatus_GUESS_CORRECT  GuessStatus = 3
)

var GuessStatus_name = map[int32]string{
	0: "GUESS_NONE",
	1: "GUESS_TOO_LOW",
	2: "GUESS_TOO_HIGH",
	3: "GUESS_CORRECT",
}

var GuessStatus_value = map[string]int32{
	"GUESS_NONE":     0,
	"GUESS_TOO_LOW":  1,
	"GUESS_TOO_HIGH": 2,
	"GUESS_CORRECT":  3,
}

func (x GuessStatus) String() string {
	return proto.EnumName(GuessStatus_name, int32(x))
}

func (GuessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4c0d186668d3726, []int{0}
}

type Guess struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Guess) Reset()         { *m = Guess{} }
func (m *Guess) String() string { return proto.CompactTextString(m) }
func (*Guess) ProtoMessage()    {}
func (*Guess) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c0d186668d3726, []int{0}
}

func (m *Guess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guess.Unmarshal(m, b)
}
func (m *Guess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guess.Marshal(b, m, deterministic)
}
func (m *Guess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guess.Merge(m, src)
}
func (m *Guess) XXX_Size() int {
	return xxx_messageInfo_Guess.Size(m)
}
func (m *Guess) XXX_DiscardUnknown() {
	xxx_messageInfo_Guess.DiscardUnknown(m)
}

var xxx_messageInfo_Guess proto.InternalMessageInfo

func (m *Guess) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GuessResult struct {
	Result               GuessStatus `protobuf:"varint,1,opt,name=result,proto3,enum=GuessStatus" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GuessResult) Reset()         { *m = GuessResult{} }
func (m *GuessResult) String() string { return proto.CompactTextString(m) }
func (*GuessResult) ProtoMessage()    {}
func (*GuessResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c0d186668d3726, []int{1}
}

func (m *GuessResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuessResult.Unmarshal(m, b)
}
func (m *GuessResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuessResult.Marshal(b, m, deterministic)
}
func (m *GuessResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuessResult.Merge(m, src)
}
func (m *GuessResult) XXX_Size() int {
	return xxx_messageInfo_GuessResult.Size(m)
}
func (m *GuessResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GuessResult.DiscardUnknown(m)
}

var xxx_messageInfo_GuessResult proto.InternalMessageInfo

func (m *GuessResult) GetResult() GuessStatus {
	if m != nil {
		return m.Result
	}
	return GuessStatus_GUESS_NONE
}

func init() {
	proto.RegisterEnum("GuessStatus", GuessStatus_name, GuessStatus_value)
	proto.RegisterType((*Guess)(nil), "Guess")
	proto.RegisterType((*GuessResult)(nil), "GuessResult")
}

func init() {
	proto.RegisterFile("proto/guessing-game.proto", fileDescriptor_f4c0d186668d3726)
}

var fileDescriptor_f4c0d186668d3726 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x4d, 0x2d, 0x2e, 0xce, 0xcc, 0x4b, 0xd7, 0x4d, 0x4f, 0xcc, 0x4d, 0xd5,
	0x03, 0x8b, 0x29, 0xc9, 0x73, 0xb1, 0xba, 0x83, 0x84, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73,
	0x93, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x63, 0x2e, 0x6e,
	0xb0, 0x82, 0xa0, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x21, 0x15, 0x2e, 0xb6, 0x22, 0x30, 0x0b, 0xac,
	0x8c, 0xcf, 0x88, 0x47, 0x0f, 0x2c, 0x1b, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x1c, 0x04, 0x95, 0xd3,
	0x0a, 0x87, 0x6a, 0x82, 0x08, 0x0b, 0xf1, 0x71, 0x71, 0xb9, 0x87, 0xba, 0x06, 0x07, 0xc7, 0xfb,
	0xf9, 0xfb, 0xb9, 0x0a, 0x30, 0x08, 0x09, 0x72, 0xf1, 0x42, 0xf8, 0x21, 0xfe, 0xfe, 0xf1, 0x3e,
	0xfe, 0xe1, 0x02, 0x8c, 0x42, 0x42, 0x5c, 0x7c, 0x08, 0x21, 0x0f, 0x4f, 0x77, 0x0f, 0x01, 0x26,
	0x84, 0x32, 0x67, 0xff, 0xa0, 0x20, 0x57, 0xe7, 0x10, 0x01, 0x66, 0x23, 0x63, 0x2e, 0x1e, 0x77,
	0xa8, 0x2f, 0xdc, 0x13, 0x73, 0x53, 0x85, 0x94, 0xb9, 0x38, 0x7d, 0x13, 0xb3, 0x53, 0x21, 0x5e,
	0x60, 0x83, 0xb8, 0x45, 0x0a, 0xea, 0x26, 0x88, 0x8b, 0x95, 0x18, 0x9c, 0xb8, 0xa3, 0x38, 0x0b,
	0xb2, 0xd3, 0xf5, 0xc1, 0x1e, 0x4e, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x1f, 0x15, 0xd8, 0x14, 0x01, 0x00, 0x00,
}
