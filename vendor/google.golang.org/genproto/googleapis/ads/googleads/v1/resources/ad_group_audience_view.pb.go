// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_audience_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

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

// An ad group audience view.
// Includes performance data from interests and remarketing lists for Display
// Network and YouTube Network ads, and remarketing lists for search ads (RLSA),
// aggregated at the audience level.
type AdGroupAudienceView struct {
	// The resource name of the ad group audience view.
	// Ad group audience view resource names have the form:
	//
	// `customers/{customer_id}/adGroupAudienceViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAudienceView) Reset()         { *m = AdGroupAudienceView{} }
func (m *AdGroupAudienceView) String() string { return proto.CompactTextString(m) }
func (*AdGroupAudienceView) ProtoMessage()    {}
func (*AdGroupAudienceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_audience_view_52e87a3873ae29dc, []int{0}
}
func (m *AdGroupAudienceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAudienceView.Unmarshal(m, b)
}
func (m *AdGroupAudienceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAudienceView.Marshal(b, m, deterministic)
}
func (dst *AdGroupAudienceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAudienceView.Merge(dst, src)
}
func (m *AdGroupAudienceView) XXX_Size() int {
	return xxx_messageInfo_AdGroupAudienceView.Size(m)
}
func (m *AdGroupAudienceView) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAudienceView.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAudienceView proto.InternalMessageInfo

func (m *AdGroupAudienceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AdGroupAudienceView)(nil), "google.ads.googleads.v1.resources.AdGroupAudienceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_audience_view.proto", fileDescriptor_ad_group_audience_view_52e87a3873ae29dc)
}

var fileDescriptor_ad_group_audience_view_52e87a3873ae29dc = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0xc6, 0x99, 0x0a, 0x82, 0x45, 0x37, 0xe3, 0x66, 0x10, 0x17, 0x8e, 0x32, 0xe0, 0x2a, 0x21,
	0xb8, 0x8b, 0x20, 0x64, 0x36, 0x05, 0x17, 0x32, 0xcc, 0xa2, 0x0b, 0x29, 0x94, 0xd8, 0x3c, 0x42,
	0x60, 0x9a, 0x57, 0x92, 0xb6, 0x73, 0x02, 0x2f, 0xe2, 0xd2, 0xa3, 0x78, 0x14, 0x4f, 0x21, 0x9d,
	0x4c, 0xb2, 0x12, 0xdd, 0x7d, 0x24, 0xbf, 0xef, 0x0f, 0x2f, 0x7f, 0xd2, 0x88, 0x7a, 0x07, 0x54,
	0x2a, 0x4f, 0x83, 0x9c, 0xd4, 0xc8, 0xa8, 0x03, 0x8f, 0x83, 0x6b, 0xc0, 0x53, 0xa9, 0x6a, 0xed,
	0x70, 0xe8, 0x6a, 0x39, 0x28, 0x03, 0xb6, 0x81, 0x7a, 0x34, 0xb0, 0x27, 0x9d, 0xc3, 0x1e, 0xe7,
	0xcb, 0x60, 0x22, 0x52, 0x79, 0x92, 0xfc, 0x64, 0x64, 0x24, 0xf9, 0xaf, 0xae, 0x63, 0x45, 0x67,
	0xa8, 0xb4, 0x16, 0x7b, 0xd9, 0x1b, 0xb4, 0x3e, 0x04, 0xdc, 0xf2, 0xfc, 0x52, 0xa8, 0x62, 0xca,
	0x17, 0xc7, 0xf8, 0xd2, 0xc0, 0x7e, 0x7e, 0x97, 0x5f, 0xc4, 0x84, 0xda, 0xca, 0x16, 0x16, 0xb3,
	0x9b, 0xd9, 0xfd, 0xd9, 0xf6, 0x3c, 0x3e, 0xbe, 0xc8, 0x16, 0xd6, 0xef, 0x59, 0xbe, 0x6a, 0xb0,
	0x25, 0xff, 0x6e, 0x58, 0x2f, 0x7e, 0xe9, 0xd8, 0x4c, 0xfd, 0x9b, 0xd9, 0xeb, 0xf3, 0xd1, 0xae,
	0x71, 0x27, 0xad, 0x26, 0xe8, 0x34, 0xd5, 0x60, 0x0f, 0xeb, 0xe2, 0x49, 0x3a, 0xe3, 0xff, 0xb8,
	0xd0, 0x63, 0x52, 0x1f, 0xd9, 0x49, 0x21, 0xc4, 0x67, 0xb6, 0x2c, 0x42, 0xa4, 0x50, 0x9e, 0x04,
	0x39, 0xa9, 0x92, 0x91, 0x6d, 0x24, 0xbf, 0x22, 0x53, 0x09, 0xe5, 0xab, 0xc4, 0x54, 0x25, 0xab,
	0x12, 0xf3, 0x9d, 0xad, 0xc2, 0x07, 0xe7, 0x42, 0x79, 0xce, 0x13, 0xc5, 0x79, 0xc9, 0x38, 0x4f,
	0xdc, 0xdb, 0xe9, 0x61, 0xec, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xe3, 0x2c, 0x1c,
	0xcd, 0x01, 0x00, 0x00,
}
