// Code generated by protoc-gen-go. DO NOT EDIT.
// source: location.proto

package locationpb // import "github.com/aefalcon/github-keystore-protobuf/go/locationpb"

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

type Location struct {
	// Types that are valid to be assigned to Location:
	//	*Location_S3
	//	*Location_Url
	Location             isLocation_Location `protobuf_oneof:"location"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_location_950fe5b99e666d12, []int{0}
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

type isLocation_Location interface {
	isLocation_Location()
}

type Location_S3 struct {
	S3 *S3Ref `protobuf:"bytes,1,opt,name=s3,proto3,oneof"`
}
type Location_Url struct {
	Url string `protobuf:"bytes,2,opt,name=url,proto3,oneof"`
}

func (*Location_S3) isLocation_Location()  {}
func (*Location_Url) isLocation_Location() {}

func (m *Location) GetLocation() isLocation_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Location) GetS3() *S3Ref {
	if x, ok := m.GetLocation().(*Location_S3); ok {
		return x.S3
	}
	return nil
}

func (m *Location) GetUrl() string {
	if x, ok := m.GetLocation().(*Location_Url); ok {
		return x.Url
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Location) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Location_OneofMarshaler, _Location_OneofUnmarshaler, _Location_OneofSizer, []interface{}{
		(*Location_S3)(nil),
		(*Location_Url)(nil),
	}
}

func _Location_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Location)
	// location
	switch x := m.Location.(type) {
	case *Location_S3:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.S3); err != nil {
			return err
		}
	case *Location_Url:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Url)
	case nil:
	default:
		return fmt.Errorf("Location.Location has unexpected type %T", x)
	}
	return nil
}

func _Location_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Location)
	switch tag {
	case 1: // location.s3
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(S3Ref)
		err := b.DecodeMessage(msg)
		m.Location = &Location_S3{msg}
		return true, err
	case 2: // location.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Location = &Location_Url{x}
		return true, err
	default:
		return false, nil
	}
}

func _Location_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Location)
	// location
	switch x := m.Location.(type) {
	case *Location_S3:
		s := proto.Size(x.S3)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Location_Url:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type S3Ref struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S3Ref) Reset()         { *m = S3Ref{} }
func (m *S3Ref) String() string { return proto.CompactTextString(m) }
func (*S3Ref) ProtoMessage()    {}
func (*S3Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_location_950fe5b99e666d12, []int{1}
}
func (m *S3Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3Ref.Unmarshal(m, b)
}
func (m *S3Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3Ref.Marshal(b, m, deterministic)
}
func (dst *S3Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3Ref.Merge(dst, src)
}
func (m *S3Ref) XXX_Size() int {
	return xxx_messageInfo_S3Ref.Size(m)
}
func (m *S3Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_S3Ref.DiscardUnknown(m)
}

var xxx_messageInfo_S3Ref proto.InternalMessageInfo

func (m *S3Ref) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *S3Ref) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *S3Ref) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*Location)(nil), "mobettersoftware.protobuf.location.Location")
	proto.RegisterType((*S3Ref)(nil), "mobettersoftware.protobuf.location.S3Ref")
}

func init() { proto.RegisterFile("location.proto", fileDescriptor_location_950fe5b99e666d12) }

var fileDescriptor_location_950fe5b99e666d12 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8e, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x69, 0x23, 0xaa, 0xf6, 0x90, 0x10, 0xf2, 0x80, 0x32, 0xa2, 0x4e, 0x30, 0xd4, 0x96,
	0xc8, 0x06, 0x4c, 0x9d, 0x40, 0x62, 0x32, 0x1b, 0x9b, 0x6d, 0x9d, 0x4d, 0xd5, 0xb4, 0x17, 0x39,
	0x67, 0x21, 0xfe, 0x3d, 0x4e, 0x9c, 0xb0, 0xb2, 0xdd, 0x7b, 0xfa, 0xbe, 0xbb, 0x83, 0xeb, 0x96,
	0x9c, 0xe1, 0x03, 0x9d, 0x65, 0x17, 0x89, 0x49, 0x6c, 0x4f, 0x64, 0x91, 0x19, 0x63, 0x4f, 0x9e,
	0xbf, 0x4d, 0xc4, 0xd2, 0xdb, 0xe4, 0xe5, 0x4c, 0x6e, 0x03, 0xac, 0xdf, 0xa7, 0x59, 0x3c, 0xc3,
	0xb2, 0x6f, 0xea, 0xc5, 0xdd, 0xe2, 0xfe, 0xea, 0xf1, 0x41, 0xfe, 0x2f, 0xcb, 0x8f, 0x46, 0xa3,
	0x7f, 0xbd, 0xd0, 0x59, 0x13, 0x02, 0xaa, 0x14, 0xdb, 0x7a, 0x99, 0xed, 0x4d, 0xae, 0x86, 0xb0,
	0x07, 0x58, 0xff, 0x1d, 0x7a, 0x83, 0xcb, 0x11, 0x17, 0xb7, 0xb0, 0xb2, 0xc9, 0x1d, 0x91, 0xc7,
	0x4b, 0x1b, 0x3d, 0x25, 0x71, 0x03, 0xd5, 0x11, 0x7f, 0xca, 0x02, 0x3d, 0x8c, 0x03, 0x19, 0x31,
	0x64, 0xb9, 0xae, 0x0a, 0x59, 0xd2, 0xfe, 0xe5, 0xf3, 0x29, 0x1c, 0xf8, 0x2b, 0x59, 0xe9, 0xe8,
	0xa4, 0x0c, 0x7a, 0xd3, 0x3a, 0x3a, 0xab, 0xd2, 0xed, 0xb2, 0xd9, 0x33, 0x45, 0xdc, 0xcd, 0xff,
	0xaa, 0x40, 0x6a, 0x7e, 0xa3, 0xb3, 0x76, 0x35, 0xf6, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe3, 0xd1, 0x41, 0x50, 0x2e, 0x01, 0x00, 0x00,
}
