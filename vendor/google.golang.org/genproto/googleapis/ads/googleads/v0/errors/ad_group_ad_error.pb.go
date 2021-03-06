// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_ad_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible ad group ad errors.
type AdGroupAdErrorEnum_AdGroupAdError int32

const (
	// Enum unspecified.
	AdGroupAdErrorEnum_UNSPECIFIED AdGroupAdErrorEnum_AdGroupAdError = 0
	// The received error code is not known in this version.
	AdGroupAdErrorEnum_UNKNOWN AdGroupAdErrorEnum_AdGroupAdError = 1
	// No link found between the adgroup ad and the label.
	AdGroupAdErrorEnum_AD_GROUP_AD_LABEL_DOES_NOT_EXIST AdGroupAdErrorEnum_AdGroupAdError = 2
	// The label has already been attached to the adgroup ad.
	AdGroupAdErrorEnum_AD_GROUP_AD_LABEL_ALREADY_EXISTS AdGroupAdErrorEnum_AdGroupAdError = 3
	// The specified ad was not found in the adgroup
	AdGroupAdErrorEnum_AD_NOT_UNDER_ADGROUP AdGroupAdErrorEnum_AdGroupAdError = 4
	// Removed ads may not be modified
	AdGroupAdErrorEnum_CANNOT_OPERATE_ON_REMOVED_ADGROUPAD AdGroupAdErrorEnum_AdGroupAdError = 5
	// An ad of this type is deprecated and cannot be created. Only deletions
	// are permitted.
	AdGroupAdErrorEnum_CANNOT_CREATE_DEPRECATED_ADS AdGroupAdErrorEnum_AdGroupAdError = 6
	// Text ads are deprecated and cannot be created. Use expanded text ads
	// instead.
	AdGroupAdErrorEnum_CANNOT_CREATE_TEXT_ADS AdGroupAdErrorEnum_AdGroupAdError = 7
	// A required field was not specified or is an empty string.
	AdGroupAdErrorEnum_EMPTY_FIELD AdGroupAdErrorEnum_AdGroupAdError = 8
	// An ad may only be modified once per call
	AdGroupAdErrorEnum_RESOURCE_REFERENCED_IN_MULTIPLE_OPS AdGroupAdErrorEnum_AdGroupAdError = 9
)

var AdGroupAdErrorEnum_AdGroupAdError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_AD_LABEL_DOES_NOT_EXIST",
	3: "AD_GROUP_AD_LABEL_ALREADY_EXISTS",
	4: "AD_NOT_UNDER_ADGROUP",
	5: "CANNOT_OPERATE_ON_REMOVED_ADGROUPAD",
	6: "CANNOT_CREATE_DEPRECATED_ADS",
	7: "CANNOT_CREATE_TEXT_ADS",
	8: "EMPTY_FIELD",
	9: "RESOURCE_REFERENCED_IN_MULTIPLE_OPS",
}
var AdGroupAdErrorEnum_AdGroupAdError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"AD_GROUP_AD_LABEL_DOES_NOT_EXIST":    2,
	"AD_GROUP_AD_LABEL_ALREADY_EXISTS":    3,
	"AD_NOT_UNDER_ADGROUP":                4,
	"CANNOT_OPERATE_ON_REMOVED_ADGROUPAD": 5,
	"CANNOT_CREATE_DEPRECATED_ADS":        6,
	"CANNOT_CREATE_TEXT_ADS":              7,
	"EMPTY_FIELD":                         8,
	"RESOURCE_REFERENCED_IN_MULTIPLE_OPS": 9,
}

func (x AdGroupAdErrorEnum_AdGroupAdError) String() string {
	return proto.EnumName(AdGroupAdErrorEnum_AdGroupAdError_name, int32(x))
}
func (AdGroupAdErrorEnum_AdGroupAdError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_ad_error_f2e4af8e18d61396, []int{0, 0}
}

// Container for enum describing possible ad group ad errors.
type AdGroupAdErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdErrorEnum) Reset()         { *m = AdGroupAdErrorEnum{} }
func (m *AdGroupAdErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdErrorEnum) ProtoMessage()    {}
func (*AdGroupAdErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_ad_error_f2e4af8e18d61396, []int{0}
}
func (m *AdGroupAdErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupAdErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdErrorEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupAdErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdErrorEnum.Merge(dst, src)
}
func (m *AdGroupAdErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdErrorEnum.Size(m)
}
func (m *AdGroupAdErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupAdErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupAdErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupAdErrorEnum_AdGroupAdError", AdGroupAdErrorEnum_AdGroupAdError_name, AdGroupAdErrorEnum_AdGroupAdError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_ad_error.proto", fileDescriptor_ad_group_ad_error_f2e4af8e18d61396)
}

var fileDescriptor_ad_group_ad_error_f2e4af8e18d61396 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0x86, 0x9d, 0xae, 0xee, 0x6a, 0x16, 0xb4, 0x54, 0x11, 0x11, 0x59, 0x96, 0x51, 0xf0, 0x2e,
	0x2d, 0x08, 0x5e, 0xd4, 0xab, 0x33, 0xcd, 0x99, 0xa1, 0xd8, 0x49, 0x42, 0xda, 0x8e, 0xbb, 0x32,
	0x10, 0xaa, 0x19, 0x8a, 0xb0, 0x3b, 0x19, 0x5a, 0x77, 0x1f, 0xc8, 0x4b, 0xef, 0x7d, 0x09, 0xf1,
	0x49, 0xbc, 0xf0, 0x19, 0x24, 0xcd, 0xce, 0xc0, 0x80, 0xee, 0x55, 0x7f, 0xfe, 0xf3, 0xfd, 0x3d,
	0xc9, 0x39, 0x21, 0x6f, 0x5b, 0x6b, 0xdb, 0x8b, 0x55, 0xdc, 0x98, 0x3e, 0xf6, 0xd2, 0xa9, 0xeb,
	0x24, 0x5e, 0x75, 0x9d, 0xed, 0xfa, 0xb8, 0x31, 0xba, 0xed, 0xec, 0xd5, 0x46, 0x37, 0x46, 0x0f,
	0x16, 0xdd, 0x74, 0xf6, 0xab, 0x8d, 0x4e, 0x3c, 0x4c, 0x1b, 0xd3, 0xd3, 0x5d, 0x8e, 0x5e, 0x27,
	0xd4, 0xe7, 0xc6, 0xbf, 0x02, 0x12, 0x81, 0x99, 0xb9, 0x28, 0x18, 0x74, 0x1e, 0xae, 0xaf, 0x2e,
	0xc7, 0x3f, 0x02, 0xf2, 0x70, 0xdf, 0x8e, 0x1e, 0x91, 0xe3, 0x9a, 0x97, 0x12, 0xb3, 0x7c, 0x9a,
	0x23, 0x0b, 0xef, 0x44, 0xc7, 0xe4, 0xa8, 0xe6, 0xef, 0xb9, 0xf8, 0xc0, 0xc3, 0x51, 0xf4, 0x8a,
	0x9c, 0x02, 0xd3, 0x33, 0x25, 0x6a, 0xa9, 0x81, 0xe9, 0x02, 0x26, 0x58, 0x68, 0x26, 0xb0, 0xd4,
	0x5c, 0x54, 0x1a, 0xcf, 0xf2, 0xb2, 0x0a, 0x83, 0x7f, 0x53, 0x50, 0x28, 0x04, 0x76, 0xee, 0xa1,
	0x32, 0x3c, 0x88, 0x9e, 0x91, 0x27, 0xc0, 0x86, 0x5c, 0xcd, 0x19, 0x2a, 0x0d, 0x6c, 0x48, 0x84,
	0x77, 0xa3, 0xd7, 0xe4, 0x65, 0x06, 0xdc, 0x55, 0x84, 0x44, 0x05, 0x15, 0x6a, 0xc1, 0xb5, 0xc2,
	0xb9, 0x58, 0x20, 0xdb, 0x62, 0xc0, 0xc2, 0x7b, 0xd1, 0x29, 0x79, 0x71, 0x03, 0x66, 0x0a, 0x1d,
	0xc7, 0x50, 0x2a, 0xcc, 0xa0, 0x1a, 0xb8, 0x32, 0x3c, 0x8c, 0x9e, 0x93, 0xa7, 0xfb, 0x44, 0x85,
	0x67, 0xd5, 0x50, 0x3b, 0x72, 0x57, 0xc5, 0xb9, 0xac, 0xce, 0xf5, 0x34, 0xc7, 0x82, 0x85, 0xf7,
	0x5d, 0x5f, 0x85, 0xa5, 0xa8, 0x55, 0x86, 0x5a, 0xe1, 0x14, 0x15, 0xf2, 0x0c, 0x99, 0xce, 0xb9,
	0x9e, 0xd7, 0x45, 0x95, 0xcb, 0x02, 0xb5, 0x90, 0x65, 0xf8, 0x60, 0xf2, 0x67, 0x44, 0xc6, 0x9f,
	0xed, 0x25, 0xbd, 0x7d, 0xea, 0x93, 0xc7, 0xfb, 0xb3, 0x95, 0x6e, 0x55, 0x72, 0xf4, 0x91, 0xdd,
	0xc4, 0x5a, 0x7b, 0xd1, 0xac, 0x5b, 0x6a, 0xbb, 0x36, 0x6e, 0x57, 0xeb, 0x61, 0x91, 0xdb, 0xa5,
	0x6f, 0xbe, 0xf4, 0xff, 0x7b, 0x03, 0xef, 0xfc, 0xe7, 0x5b, 0x70, 0x30, 0x03, 0xf8, 0x1e, 0x9c,
	0xcc, 0xfc, 0xcf, 0xc0, 0xf4, 0xd4, 0x4b, 0xa7, 0x16, 0x09, 0x1d, 0x5a, 0xf6, 0x3f, 0xb7, 0xc0,
	0x12, 0x4c, 0xbf, 0xdc, 0x01, 0xcb, 0x45, 0xb2, 0xf4, 0xc0, 0xef, 0x60, 0xec, 0xdd, 0x34, 0x05,
	0xd3, 0xa7, 0xe9, 0x0e, 0x49, 0xd3, 0x45, 0x92, 0xa6, 0x1e, 0xfa, 0x74, 0x38, 0x9c, 0xee, 0xcd,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x3e, 0x79, 0x8a, 0xa0, 0x02, 0x00, 0x00,
}
