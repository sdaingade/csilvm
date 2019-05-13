// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/not_empty_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible not empty errors.
type NotEmptyErrorEnum_NotEmptyError int32

const (
	// Enum unspecified.
	NotEmptyErrorEnum_UNSPECIFIED NotEmptyErrorEnum_NotEmptyError = 0
	// The received error code is not known in this version.
	NotEmptyErrorEnum_UNKNOWN NotEmptyErrorEnum_NotEmptyError = 1
	// Empty list.
	NotEmptyErrorEnum_EMPTY_LIST NotEmptyErrorEnum_NotEmptyError = 2
)

var NotEmptyErrorEnum_NotEmptyError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EMPTY_LIST",
}
var NotEmptyErrorEnum_NotEmptyError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EMPTY_LIST":  2,
}

func (x NotEmptyErrorEnum_NotEmptyError) String() string {
	return proto.EnumName(NotEmptyErrorEnum_NotEmptyError_name, int32(x))
}
func (NotEmptyErrorEnum_NotEmptyError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_not_empty_error_60da7a8d2d549cbf, []int{0, 0}
}

// Container for enum describing possible not empty errors.
type NotEmptyErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotEmptyErrorEnum) Reset()         { *m = NotEmptyErrorEnum{} }
func (m *NotEmptyErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NotEmptyErrorEnum) ProtoMessage()    {}
func (*NotEmptyErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_not_empty_error_60da7a8d2d549cbf, []int{0}
}
func (m *NotEmptyErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotEmptyErrorEnum.Unmarshal(m, b)
}
func (m *NotEmptyErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotEmptyErrorEnum.Marshal(b, m, deterministic)
}
func (dst *NotEmptyErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotEmptyErrorEnum.Merge(dst, src)
}
func (m *NotEmptyErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NotEmptyErrorEnum.Size(m)
}
func (m *NotEmptyErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NotEmptyErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NotEmptyErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NotEmptyErrorEnum)(nil), "google.ads.googleads.v1.errors.NotEmptyErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.NotEmptyErrorEnum_NotEmptyError", NotEmptyErrorEnum_NotEmptyError_name, NotEmptyErrorEnum_NotEmptyError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/not_empty_error.proto", fileDescriptor_not_empty_error_60da7a8d2d549cbf)
}

var fileDescriptor_not_empty_error_60da7a8d2d549cbf = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0x6d, 0x05, 0x85, 0x0c, 0xb5, 0xf6, 0x52, 0x64, 0x17, 0x7d, 0x80, 0x84, 0xa2, 0x57,
	0x11, 0x2f, 0x3a, 0x17, 0x47, 0x51, 0x6b, 0x71, 0x5b, 0x45, 0x29, 0x94, 0x6a, 0x4b, 0x28, 0xac,
	0xf9, 0x4b, 0x13, 0x07, 0xbe, 0x8e, 0x97, 0x3e, 0x8a, 0x8f, 0x22, 0xf8, 0x0e, 0x92, 0xc6, 0x16,
	0x76, 0xe1, 0xae, 0x72, 0x38, 0x7c, 0xe7, 0xe4, 0xf0, 0xa3, 0x73, 0x0e, 0xc0, 0x57, 0x25, 0xc9,
	0x0b, 0x49, 0x8c, 0xd4, 0x6a, 0xed, 0x93, 0xb2, 0x6d, 0xa1, 0x95, 0x44, 0x80, 0xca, 0xca, 0xba,
	0x51, 0xef, 0x59, 0x67, 0xe0, 0xa6, 0x05, 0x05, 0xee, 0xd8, 0xa0, 0x38, 0x2f, 0x24, 0x1e, 0x52,
	0x78, 0xed, 0x63, 0x93, 0x3a, 0x39, 0xed, 0x5b, 0x9b, 0x8a, 0xe4, 0x42, 0x80, 0xca, 0x55, 0x05,
	0x42, 0x9a, 0xb4, 0xf7, 0x80, 0x8e, 0x23, 0x50, 0x4c, 0xb7, 0x32, 0xcd, 0x33, 0xf1, 0x56, 0x7b,
	0x97, 0xe8, 0x60, 0xc3, 0x74, 0x8f, 0xd0, 0x68, 0x19, 0xcd, 0x63, 0x76, 0x15, 0x5e, 0x87, 0x6c,
	0xea, 0xec, 0xb8, 0x23, 0xb4, 0xbf, 0x8c, 0x6e, 0xa2, 0xfb, 0xc7, 0xc8, 0xb1, 0xdc, 0x43, 0x84,
	0xd8, 0x5d, 0xbc, 0x78, 0xca, 0x6e, 0xc3, 0xf9, 0xc2, 0xb1, 0x27, 0x3f, 0x16, 0xf2, 0x5e, 0xa1,
	0xc6, 0xdb, 0x87, 0x4d, 0xdc, 0x8d, 0x3f, 0x62, 0x3d, 0x27, 0xb6, 0x9e, 0xa7, 0x7f, 0x29, 0x0e,
	0xab, 0x5c, 0x70, 0x0c, 0x2d, 0x27, 0xbc, 0x14, 0xdd, 0xd8, 0xfe, 0x28, 0x4d, 0x25, 0xff, 0xbb,
	0xd1, 0x85, 0x79, 0x3e, 0xec, 0xdd, 0x59, 0x10, 0x7c, 0xda, 0xe3, 0x99, 0x29, 0x0b, 0x0a, 0x89,
	0x8d, 0xd4, 0x2a, 0xf1, 0x71, 0xf7, 0xa5, 0xfc, 0xea, 0x81, 0x34, 0x28, 0x64, 0x3a, 0x00, 0x69,
	0xe2, 0xa7, 0x06, 0xf8, 0xb6, 0x3d, 0xe3, 0x52, 0x1a, 0x14, 0x92, 0xd2, 0x01, 0xa1, 0x34, 0xf1,
	0x29, 0x35, 0xd0, 0xcb, 0x5e, 0xb7, 0xee, 0xec, 0x37, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x55, 0x9c,
	0x20, 0xc0, 0x01, 0x00, 0x00,
}
