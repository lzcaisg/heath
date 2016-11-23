// Code generated by protoc-gen-go.
// source: block.proto
// DO NOT EDIT!

/*
Package block is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Signature
	Block
*/
package block

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

// metadat about a block signature
type Signature struct {
	Timestamp        *int64 `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	ContentHash      []byte `protobuf:"bytes,2,req,name=contentHash" json:"contentHash,omitempty"`
	SignatureA       []byte `protobuf:"bytes,3,req,name=signatureA" json:"signatureA,omitempty"`
	SignatureB       []byte `protobuf:"bytes,4,req,name=signatureB" json:"signatureB,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Signature) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Signature) GetContentHash() []byte {
	if m != nil {
		return m.ContentHash
	}
	return nil
}

func (m *Signature) GetSignatureA() []byte {
	if m != nil {
		return m.SignatureA
	}
	return nil
}

func (m *Signature) GetSignatureB() []byte {
	if m != nil {
		return m.SignatureB
	}
	return nil
}

// Metadata about a block
type Block struct {
	Parent           *Signature `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	Signature        *Signature `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	Payload          []byte     `protobuf:"bytes,3,req,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetParent() *Signature {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *Block) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*Block)(nil), "Block")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8e, 0xbd, 0x0e, 0x82, 0x30,
	0x14, 0x46, 0x53, 0xf0, 0x27, 0x5c, 0x9c, 0x3a, 0x75, 0x30, 0xa6, 0x61, 0xea, 0xc4, 0xe0, 0x1b,
	0xc8, 0xe4, 0x8c, 0x4f, 0x70, 0xc5, 0x46, 0x89, 0xd0, 0xdb, 0xc0, 0x75, 0xf0, 0x19, 0x7c, 0x69,
	0x23, 0xe1, 0xd7, 0xb1, 0xe7, 0x7c, 0xbd, 0x39, 0x10, 0x5f, 0x2b, 0x2a, 0x9e, 0xa9, 0x6f, 0x88,
	0x29, 0xf9, 0x08, 0x88, 0x2e, 0xe5, 0xdd, 0x21, 0xbf, 0x1a, 0x2b, 0xf7, 0x10, 0x71, 0x59, 0xdb,
	0x96, 0xb1, 0xf6, 0x4a, 0xe8, 0xc0, 0x84, 0xf9, 0x04, 0xa4, 0x86, 0xb8, 0x20, 0xc7, 0xd6, 0xf1,
	0x19, 0xdb, 0x87, 0x0a, 0x74, 0x60, 0x76, 0xf9, 0x1c, 0xc9, 0x03, 0x40, 0x3b, 0x1c, 0x3b, 0xa9,
	0xb0, 0x1b, 0xcc, 0xc8, 0xc2, 0x67, 0x6a, 0xf5, 0xe7, 0xb3, 0x84, 0x60, 0x9d, 0xfd, 0xe2, 0x64,
	0x02, 0x1b, 0x8f, 0x8d, 0x75, 0xac, 0x84, 0x16, 0x26, 0x3e, 0x42, 0x3a, 0x46, 0xe6, 0xbd, 0x91,
	0x06, 0xa2, 0xf1, 0x6b, 0x17, 0xb3, 0x9c, 0x4d, 0x52, 0x2a, 0xd8, 0x7a, 0x7c, 0x57, 0x84, 0xb7,
	0xbe, 0x69, 0x78, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x76, 0x20, 0x0b, 0x3c, 0x0c, 0x01, 0x00,
	0x00,
}