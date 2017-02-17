// Code generated by protoc-gen-go.
// source: signature.proto
// DO NOT EDIT!

/*
Package signature is a generated protocol buffer package.

It is generated from these files:
	signature.proto

It has these top-level messages:
	DigitallySigned
*/
package signature

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

// HashAlgorithm defines the approved methods for object hashing.
type DigitallySigned_HashAlgorithm int32

const (
	// No hash algorithm is used.
	DigitallySigned_NONE DigitallySigned_HashAlgorithm = 0
	// SHA256 is used.
	DigitallySigned_SHA256 DigitallySigned_HashAlgorithm = 4
	// SHA512 is used.
	DigitallySigned_SHA512 DigitallySigned_HashAlgorithm = 6
)

var DigitallySigned_HashAlgorithm_name = map[int32]string{
	0: "NONE",
	4: "SHA256",
	6: "SHA512",
}
var DigitallySigned_HashAlgorithm_value = map[string]int32{
	"NONE":   0,
	"SHA256": 4,
	"SHA512": 6,
}

func (x DigitallySigned_HashAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_HashAlgorithm_name, int32(x))
}
func (DigitallySigned_HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// SignatureAlgorithm defines the algorithm used to sign the object.
type DigitallySigned_SignatureAlgorithm int32

const (
	// Anonymous signature scheme.
	DigitallySigned_ANONYMOUS DigitallySigned_SignatureAlgorithm = 0
	// ECDSA signature scheme.
	DigitallySigned_ECDSA DigitallySigned_SignatureAlgorithm = 1
	// RSA 3072-bit signature scheme.
	DigitallySigned_RSA_3072 DigitallySigned_SignatureAlgorithm = 2
)

var DigitallySigned_SignatureAlgorithm_name = map[int32]string{
	0: "ANONYMOUS",
	1: "ECDSA",
	2: "RSA_3072",
}
var DigitallySigned_SignatureAlgorithm_value = map[string]int32{
	"ANONYMOUS": 0,
	"ECDSA":     1,
	"RSA_3072":  2,
}

func (x DigitallySigned_SignatureAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_SignatureAlgorithm_name, int32(x))
}
func (DigitallySigned_SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

// DigitallySigned defines a way to digitally sign objects.
type DigitallySigned struct {
	// hash_algorithm contains the hash algorithm used.
	HashAlgorithm DigitallySigned_HashAlgorithm `protobuf:"varint,1,opt,name=hash_algorithm,json=hashAlgorithm,enum=signature.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// sig_algorithm contains the signing algorithm used.
	SigAlgorithm DigitallySigned_SignatureAlgorithm `protobuf:"varint,2,opt,name=sig_algorithm,json=sigAlgorithm,enum=signature.DigitallySigned_SignatureAlgorithm" json:"sig_algorithm,omitempty"`
	// signature contains the object signature.
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *DigitallySigned) Reset()                    { *m = DigitallySigned{} }
func (m *DigitallySigned) String() string            { return proto.CompactTextString(m) }
func (*DigitallySigned) ProtoMessage()               {}
func (*DigitallySigned) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*DigitallySigned)(nil), "signature.DigitallySigned")
	proto.RegisterEnum("signature.DigitallySigned_HashAlgorithm", DigitallySigned_HashAlgorithm_name, DigitallySigned_HashAlgorithm_value)
	proto.RegisterEnum("signature.DigitallySigned_SignatureAlgorithm", DigitallySigned_SignatureAlgorithm_name, DigitallySigned_SignatureAlgorithm_value)
}

func init() { proto.RegisterFile("signature.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0x4c, 0xcf,
	0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x1d, 0x65, 0xe2, 0xe2, 0x77, 0xc9, 0x4c, 0xcf, 0x2c, 0x49, 0xcc, 0xc9, 0xa9, 0x0c, 0xce, 0x4c,
	0xcf, 0x4b, 0x4d, 0x11, 0xf2, 0xe7, 0xe2, 0xcb, 0x48, 0x2c, 0xce, 0x88, 0x4f, 0xcc, 0x49, 0xcf,
	0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd0, 0x43, 0x18,
	0x84, 0xa6, 0x47, 0xcf, 0x23, 0xb1, 0x38, 0xc3, 0x11, 0xa6, 0x3e, 0x88, 0x37, 0x03, 0x99, 0x2b,
	0x14, 0xc4, 0xc5, 0x5b, 0x9c, 0x99, 0x8e, 0x64, 0x1e, 0x13, 0xd8, 0x3c, 0x5d, 0x3c, 0xe6, 0x05,
	0xc3, 0x64, 0x10, 0x86, 0xf2, 0x14, 0x67, 0xa6, 0x23, 0xcc, 0x94, 0xe1, 0x42, 0xf8, 0x42, 0x82,
	0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc9, 0x5b, 0x86, 0x5c, 0xbc, 0x28, 0x2e, 0x12, 0xe2, 0xe0,
	0x62, 0xf1, 0xf3, 0xf7, 0x73, 0x15, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0x0b, 0xf6, 0x70, 0x34, 0x32,
	0x35, 0x13, 0x60, 0x81, 0xb2, 0x4d, 0x0d, 0x8d, 0x04, 0xd8, 0x94, 0x6c, 0xb8, 0x84, 0x30, 0x2d,
	0x15, 0xe2, 0xe5, 0xe2, 0x74, 0xf4, 0xf3, 0xf7, 0x8b, 0xf4, 0xf5, 0x0f, 0x0d, 0x16, 0x60, 0x10,
	0xe2, 0xe4, 0x62, 0x75, 0x75, 0x76, 0x09, 0x76, 0x14, 0x60, 0x14, 0xe2, 0xe1, 0xe2, 0x08, 0x0a,
	0x76, 0x8c, 0x37, 0x36, 0x30, 0x37, 0x12, 0x60, 0x4a, 0x62, 0x03, 0x87, 0xac, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x67, 0x5c, 0x8c, 0x5d, 0x6c, 0x01, 0x00, 0x00,
}