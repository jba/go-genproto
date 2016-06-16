// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/range.proto
// DO NOT EDIT!

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A 0-based half-open genomic coordinate range for search requests.
type Range struct {
	// The reference sequence name, for example `chr1`,
	// `1`, or `chrX`.
	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName" json:"reference_name,omitempty"`
	// The start position of the range on the reference, 0-based inclusive.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// The end position of the range on the reference, 0-based exclusive.
	End int64 `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func init() {
	proto.RegisterType((*Range)(nil), "google.genomics.v1.Range")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/range.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8e, 0xb1, 0x4b, 0xc4, 0x30,
	0x18, 0xc5, 0x89, 0xa5, 0x82, 0x01, 0x45, 0x82, 0x48, 0x71, 0x12, 0x41, 0xd0, 0x25, 0xa1, 0x38,
	0xbb, 0x74, 0x73, 0x91, 0xd2, 0xc1, 0x55, 0x62, 0xfc, 0x1a, 0x02, 0xed, 0xf7, 0x95, 0x24, 0xf4,
	0x6f, 0xbf, 0xf1, 0x92, 0xf4, 0xee, 0x96, 0x5b, 0x6e, 0x09, 0xc9, 0xef, 0xbd, 0xbc, 0xf7, 0xf8,
	0xa7, 0x25, 0xb2, 0x13, 0x48, 0x4b, 0x93, 0x46, 0x2b, 0xc9, 0x5b, 0x65, 0x01, 0x17, 0x4f, 0x91,
	0xd4, 0x26, 0xe9, 0xc5, 0x85, 0xcc, 0x68, 0x76, 0x26, 0xa8, 0xb5, 0x55, 0x3e, 0x19, 0x41, 0x16,
	0x8b, 0x10, 0xc7, 0xef, 0x07, 0x5d, 0xae, 0xed, 0xd3, 0xd7, 0x65, 0x91, 0xe9, 0x50, 0x01, 0xfc,
	0xea, 0x0c, 0x18, 0xc2, 0xd1, 0x59, 0xa5, 0x11, 0x29, 0xea, 0xe8, 0x08, 0xc3, 0x16, 0xff, 0xf2,
	0xc3, 0xeb, 0x21, 0xb7, 0x89, 0x57, 0x7e, 0xe7, 0x61, 0x04, 0x0f, 0x68, 0xe0, 0x17, 0xf5, 0x0c,
	0x0d, 0x7b, 0x66, 0x6f, 0x37, 0xc3, 0xed, 0x89, 0x7e, 0x27, 0x28, 0x1e, 0x78, 0x1d, 0xa2, 0xf6,
	0xb1, 0xb9, 0x4a, 0x6a, 0x35, 0x6c, 0x0f, 0x71, 0xcf, 0x2b, 0xc0, 0xff, 0xa6, 0x2a, 0x2c, 0x5f,
	0xbb, 0x77, 0xfe, 0x68, 0x68, 0x96, 0xe7, 0xe3, 0x3b, 0x5e, 0xfa, 0xfa, 0xdc, 0xde, 0xb3, 0x1d,
	0x63, 0x7f, 0xd7, 0x65, 0xc9, 0xc7, 0x3e, 0x00, 0x00, 0xff, 0xff, 0x33, 0x75, 0x3d, 0xcd, 0x29,
	0x01, 0x00, 0x00,
}
