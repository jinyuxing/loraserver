// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common // import "github.com/brocaar/loraserver/api/common"

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}
var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}
func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_f2f882223dc106a9, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
}
var Region_value = map[string]int32{
	"EU868": 0,
	"US915": 2,
	"CN779": 3,
	"EU433": 4,
	"AU915": 5,
	"CN470": 6,
	"AS923": 7,
	"KR920": 8,
	"IN865": 9,
	"RU864": 10,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}
func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_f2f882223dc106a9, []int{1}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver.
	LocationSource_GEO_RESOLVER LocationSource = 3
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER",
}
var LocationSource_value = map[string]int32{
	"UNKNOWN":      0,
	"GPS":          1,
	"CONFIG":       2,
	"GEO_RESOLVER": 3,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}
func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_f2f882223dc106a9, []int{2}
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_f2f882223dc106a9, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterType((*Location)(nil), "common.Location")
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_f2f882223dc106a9) }

var fileDescriptor_common_f2f882223dc106a9 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0xab, 0xda, 0x40,
	0x14, 0xc5, 0x1d, 0xa3, 0x31, 0xde, 0x5a, 0x19, 0x66, 0x51, 0xa4, 0x14, 0x2a, 0x5d, 0x85, 0x2c,
	0x8c, 0x35, 0xfe, 0x49, 0x96, 0xd6, 0x46, 0x11, 0x6d, 0x52, 0x26, 0xa4, 0x85, 0x6e, 0xca, 0x18,
	0x83, 0x2f, 0x10, 0x1d, 0x19, 0x13, 0xe1, 0x7d, 0xa7, 0xf7, 0x21, 0x1f, 0x93, 0x51, 0x1f, 0x6f,
	0xf7, 0x3b, 0xf7, 0x9c, 0x3b, 0x73, 0xe7, 0x0e, 0x74, 0x12, 0x7e, 0x3c, 0xf2, 0xd3, 0xe0, 0x2c,
	0x78, 0xc1, 0x89, 0xae, 0xd4, 0xb7, 0x17, 0x04, 0xc6, 0x96, 0x27, 0xac, 0xc8, 0xf8, 0x89, 0x7c,
	0x06, 0x23, 0x67, 0x45, 0x56, 0x94, 0xfb, 0xb4, 0x87, 0xfa, 0xc8, 0x44, 0xf4, 0xa1, 0xc9, 0x17,
	0x68, 0xe7, 0xfc, 0x74, 0x50, 0x66, 0xbd, 0x32, 0xdf, 0x0a, 0xb2, 0x93, 0xe5, 0xb7, 0x4e, 0x4d,
	0x75, 0xde, 0x35, 0x19, 0x80, 0x7e, 0xe1, 0xa5, 0x48, 0xd2, 0x5e, 0xa3, 0x8f, 0xcc, 0xee, 0xe8,
	0xd3, 0xe0, 0x36, 0xc9, 0xfd, 0xde, 0xa8, 0x72, 0xe9, 0x2d, 0x55, 0x9d, 0x95, 0x24, 0xa5, 0x60,
	0xc9, 0x73, 0xaf, 0xd9, 0x47, 0xe6, 0x47, 0xfa, 0xd0, 0xd6, 0x57, 0x80, 0x5f, 0x7c, 0x5f, 0xe6,
	0x6a, 0x5e, 0x03, 0x1a, 0xdb, 0x90, 0xce, 0x71, 0x8d, 0xb4, 0x40, 0x5b, 0x46, 0x1b, 0x8c, 0xac,
	0x2b, 0xe8, 0x34, 0x3d, 0x48, 0xb3, 0x0d, 0x4d, 0x3f, 0x76, 0xa7, 0x2e, 0xae, 0x49, 0x8c, 0x23,
	0xef, 0xfb, 0x04, 0xd7, 0x25, 0x2e, 0x82, 0xd9, 0xcc, 0xc3, 0x9a, 0x0a, 0x8c, 0x1d, 0x07, 0x37,
	0x24, 0xce, 0x63, 0x19, 0x68, 0xaa, 0xc0, 0x78, 0x36, 0xc4, 0x7a, 0x55, 0x8d, 0xbc, 0x91, 0x83,
	0x5b, 0x12, 0x37, 0xd4, 0x1b, 0x0d, 0xb1, 0x21, 0x71, 0x1d, 0xb8, 0xd3, 0x09, 0x6e, 0x4b, 0xa4,
	0xb1, 0x3b, 0x1d, 0x63, 0xb0, 0x7e, 0x42, 0xf7, 0xfd, 0x73, 0xc8, 0x07, 0x68, 0xc5, 0xc1, 0x26,
	0x08, 0xff, 0x06, 0x6a, 0xbe, 0xd5, 0xef, 0x08, 0x23, 0x02, 0xa0, 0x2f, 0xc2, 0x60, 0xb9, 0x5e,
	0xe1, 0x3a, 0xc1, 0xd0, 0x59, 0xf9, 0xe1, 0x7f, 0xea, 0x47, 0xe1, 0xf6, 0x8f, 0x4f, 0xb1, 0xf6,
	0xc3, 0xfa, 0x67, 0x1e, 0xb2, 0xe2, 0xa9, 0xdc, 0xc9, 0x15, 0xd9, 0x3b, 0xc1, 0x13, 0xc6, 0x84,
	0x9d, 0x73, 0xc1, 0x2e, 0xa9, 0xb8, 0xa6, 0xc2, 0x66, 0xe7, 0xcc, 0x56, 0xdb, 0xdb, 0xe9, 0xd5,
	0x47, 0x3a, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x51, 0x0f, 0x52, 0xd8, 0x01, 0x00, 0x00,
}
