// Code generated by protoc-gen-gogo.
// source: srctree.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	srctree.proto

It has these top-level messages:
	SrcTree
*/
package main

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SrcTree struct {
	PackageName      *string    `protobuf:"bytes,1,opt,name=PackageName" json:"PackageName,omitempty"`
	Imports          []*SrcTree `protobuf:"bytes,2,rep,name=Imports" json:"Imports,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SrcTree) Reset()                    { *m = SrcTree{} }
func (m *SrcTree) String() string            { return proto.CompactTextString(m) }
func (*SrcTree) ProtoMessage()               {}
func (*SrcTree) Descriptor() ([]byte, []int) { return fileDescriptorSrctree, []int{0} }

func (m *SrcTree) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *SrcTree) GetImports() []*SrcTree {
	if m != nil {
		return m.Imports
	}
	return nil
}

func init() {
	proto.RegisterType((*SrcTree)(nil), "main.SrcTree")
}
func (this *SrcTree) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return SrctreeDescription()
}
func SrctreeDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3411 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x4b, 0x6c, 0x1b, 0xd7,
		0xd5, 0x0e, 0x5f, 0x12, 0x79, 0x48, 0x51, 0xa3, 0x2b, 0x45, 0xa6, 0x95, 0xc4, 0x96, 0x15, 0xe7,
		0xb7, 0x92, 0xfc, 0x91, 0x03, 0xc7, 0x76, 0x6c, 0xfa, 0x4f, 0x0c, 0x4a, 0xa2, 0x15, 0x19, 0x7a,
		0xf0, 0x1f, 0x4a, 0x89, 0x93, 0x2c, 0x06, 0x57, 0xc3, 0x4b, 0x6a, 0xec, 0xe1, 0x0c, 0xff, 0x99,
		0xa1, 0x6d, 0x79, 0xe5, 0x20, 0xff, 0x03, 0x41, 0xf0, 0xf7, 0x0d, 0x34, 0xef, 0xa4, 0x01, 0xda,
		0xb4, 0xe9, 0x33, 0x7d, 0x2d, 0xba, 0xea, 0x26, 0xed, 0xae, 0x40, 0xf6, 0xdd, 0x04, 0x08, 0xd0,
		0x57, 0xda, 0xa6, 0xad, 0x81, 0x16, 0xf0, 0xa6, 0xb8, 0xaf, 0xe1, 0x0c, 0x49, 0x79, 0xa8, 0x00,
		0x49, 0xba, 0x12, 0xef, 0xb9, 0xe7, 0xfb, 0xe6, 0xde, 0x73, 0xce, 0x3d, 0xe7, 0xcc, 0x1d, 0xc1,
		0x33, 0xc7, 0x61, 0xba, 0x61, 0xdb, 0x0d, 0x93, 0x1c, 0x6d, 0x39, 0xb6, 0x67, 0x6f, 0xb5, 0xeb,
		0x47, 0x6b, 0xc4, 0xd5, 0x1d, 0xa3, 0xe5, 0xd9, 0xce, 0x1c, 0x93, 0xa1, 0x51, 0xae, 0x31, 0x27,
		0x35, 0x66, 0x56, 0x61, 0xec, 0x9c, 0x61, 0x92, 0x45, 0x5f, 0xb1, 0x4a, 0x3c, 0x74, 0x0a, 0x92,
		0x75, 0xc3, 0x24, 0x85, 0xd8, 0x74, 0x62, 0x36, 0x7b, 0xec, 0xf0, 0x5c, 0x17, 0x68, 0x2e, 0x8c,
		0xa8, 0x50, 0xb1, 0xca, 0x10, 0x33, 0x1f, 0x24, 0x61, 0xbc, 0xcf, 0x2c, 0x42, 0x90, 0xb4, 0x70,
		0x93, 0x32, 0xc6, 0x66, 0x33, 0x2a, 0xfb, 0x8d, 0x0a, 0x30, 0xdc, 0xc2, 0xfa, 0x25, 0xdc, 0x20,
		0x85, 0x38, 0x13, 0xcb, 0x21, 0x3a, 0x00, 0x50, 0x23, 0x2d, 0x62, 0xd5, 0x88, 0xa5, 0xef, 0x14,
		0x12, 0xd3, 0x89, 0xd9, 0x8c, 0x1a, 0x90, 0xa0, 0xfb, 0x61, 0xac, 0xd5, 0xde, 0x32, 0x0d, 0x5d,
		0x0b, 0xa8, 0xc1, 0x74, 0x62, 0x36, 0xa5, 0x2a, 0x7c, 0x62, 0xb1, 0xa3, 0x7c, 0x04, 0x46, 0xaf,
		0x10, 0x7c, 0x29, 0xa8, 0x9a, 0x65, 0xaa, 0x79, 0x2a, 0x0e, 0x28, 0x2e, 0x40, 0xae, 0x49, 0x5c,
		0x17, 0x37, 0x88, 0xe6, 0xed, 0xb4, 0x48, 0x21, 0xc9, 0x76, 0x3f, 0xdd, 0xb3, 0xfb, 0xee, 0x9d,
		0x67, 0x05, 0x6a, 0x63, 0xa7, 0x45, 0x50, 0x09, 0x32, 0xc4, 0x6a, 0x37, 0x39, 0x43, 0x6a, 0x17,
		0xfb, 0x95, 0xad, 0x76, 0xb3, 0x9b, 0x25, 0x4d, 0x61, 0x82, 0x62, 0xd8, 0x25, 0xce, 0x65, 0x43,
		0x27, 0x85, 0x21, 0x46, 0x70, 0xa4, 0x87, 0xa0, 0xca, 0xe7, 0xbb, 0x39, 0x24, 0x0e, 0x2d, 0x40,
		0x86, 0x5c, 0xf5, 0x88, 0xe5, 0x1a, 0xb6, 0x55, 0x18, 0x66, 0x24, 0xf7, 0xf4, 0xf1, 0x22, 0x31,
		0x6b, 0xdd, 0x14, 0x1d, 0x1c, 0x3a, 0x09, 0xc3, 0x76, 0xcb, 0x33, 0x6c, 0xcb, 0x2d, 0xa4, 0xa7,
		0x63, 0xb3, 0xd9, 0x63, 0x77, 0xf6, 0x0d, 0x84, 0x75, 0xae, 0xa3, 0x4a, 0x65, 0xb4, 0x0c, 0x8a,
		0x6b, 0xb7, 0x1d, 0x9d, 0x68, 0xba, 0x5d, 0x23, 0x9a, 0x61, 0xd5, 0xed, 0x42, 0x86, 0x11, 0x1c,
		0xec, 0xdd, 0x08, 0x53, 0x5c, 0xb0, 0x6b, 0x64, 0xd9, 0xaa, 0xdb, 0x6a, 0xde, 0x0d, 0x8d, 0xd1,
		0x24, 0x0c, 0xb9, 0x3b, 0x96, 0x87, 0xaf, 0x16, 0x72, 0x2c, 0x42, 0xc4, 0x68, 0xe6, 0xef, 0x29,
		0x18, 0x1d, 0x24, 0xc4, 0xce, 0x40, 0xaa, 0x4e, 0x77, 0x59, 0x88, 0xef, 0xc5, 0x06, 0x1c, 0x13,
		0x36, 0xe2, 0xd0, 0xc7, 0x34, 0x62, 0x09, 0xb2, 0x16, 0x71, 0x3d, 0x52, 0xe3, 0x11, 0x91, 0x18,
		0x30, 0xa6, 0x80, 0x83, 0x7a, 0x43, 0x2a, 0xf9, 0xb1, 0x42, 0xea, 0x02, 0x8c, 0xfa, 0x4b, 0xd2,
		0x1c, 0x6c, 0x35, 0x64, 0x6c, 0x1e, 0x8d, 0x5a, 0xc9, 0x5c, 0x59, 0xe2, 0x54, 0x0a, 0x53, 0xf3,
		0x24, 0x34, 0x46, 0x8b, 0x00, 0xb6, 0x45, 0xec, 0xba, 0x56, 0x23, 0xba, 0x59, 0x48, 0xef, 0x62,
		0xa5, 0x75, 0xaa, 0xd2, 0x63, 0x25, 0x9b, 0x4b, 0x75, 0x13, 0x9d, 0xee, 0x84, 0xda, 0xf0, 0x2e,
		0x91, 0xb2, 0xca, 0x0f, 0x59, 0x4f, 0xb4, 0x6d, 0x42, 0xde, 0x21, 0x34, 0xee, 0x49, 0x4d, 0xec,
		0x2c, 0xc3, 0x16, 0x31, 0x17, 0xb9, 0x33, 0x55, 0xc0, 0xf8, 0xc6, 0x46, 0x9c, 0xe0, 0x10, 0xdd,
		0x0d, 0xbe, 0x40, 0x63, 0x61, 0x05, 0x2c, 0x0b, 0xe5, 0xa4, 0x70, 0x0d, 0x37, 0xc9, 0xd4, 0x29,
		0xc8, 0x87, 0xcd, 0x83, 0x26, 0x20, 0xe5, 0x7a, 0xd8, 0xf1, 0x58, 0x14, 0xa6, 0x54, 0x3e, 0x40,
		0x0a, 0x24, 0x88, 0x55, 0x63, 0x59, 0x2e, 0xa5, 0xd2, 0x9f, 0x53, 0x0f, 0xc3, 0x48, 0xe8, 0xf1,
		0x83, 0x02, 0x67, 0x5e, 0x18, 0x82, 0x89, 0x7e, 0x31, 0xd7, 0x37, 0xfc, 0x27, 0x61, 0xc8, 0x6a,
		0x37, 0xb7, 0x88, 0x53, 0x48, 0x30, 0x06, 0x31, 0x42, 0x25, 0x48, 0x99, 0x78, 0x8b, 0x98, 0x85,
		0xe4, 0x74, 0x6c, 0x36, 0x7f, 0xec, 0xfe, 0x81, 0xa2, 0x7a, 0x6e, 0x85, 0x42, 0x54, 0x8e, 0x44,
		0x8f, 0x42, 0x52, 0xa4, 0x38, 0xca, 0x70, 0xdf, 0x60, 0x0c, 0x34, 0x16, 0x55, 0x86, 0x43, 0x77,
		0x40, 0x86, 0xfe, 0xe5, 0xb6, 0x1d, 0x62, 0x6b, 0x4e, 0x53, 0x01, 0xb5, 0x2b, 0x9a, 0x82, 0x34,
		0x0b, 0xb3, 0x1a, 0x91, 0xa5, 0xc1, 0x1f, 0x53, 0xc7, 0xd4, 0x48, 0x1d, 0xb7, 0x4d, 0x4f, 0xbb,
		0x8c, 0xcd, 0x36, 0x61, 0x01, 0x93, 0x51, 0x73, 0x42, 0xf8, 0x38, 0x95, 0xa1, 0x83, 0x90, 0xe5,
		0x51, 0x69, 0x58, 0x35, 0x72, 0x95, 0x65, 0x9f, 0x94, 0xca, 0x03, 0x75, 0x99, 0x4a, 0xe8, 0xe3,
		0x2f, 0xba, 0xb6, 0x25, 0x5d, 0xcb, 0x1e, 0x41, 0x05, 0xec, 0xf1, 0x0f, 0x77, 0x27, 0xbe, 0xbb,
		0xfa, 0x6f, 0xaf, 0x3b, 0x16, 0x67, 0x7e, 0x1a, 0x87, 0x24, 0x3b, 0x6f, 0xa3, 0x90, 0xdd, 0x78,
		0xb2, 0x52, 0xd6, 0x16, 0xd7, 0x37, 0xe7, 0x57, 0xca, 0x4a, 0x0c, 0xe5, 0x01, 0x98, 0xe0, 0xdc,
		0xca, 0x7a, 0x69, 0x43, 0x89, 0xfb, 0xe3, 0xe5, 0xb5, 0x8d, 0x93, 0xc7, 0x95, 0x84, 0x0f, 0xd8,
		0xe4, 0x82, 0x64, 0x50, 0xe1, 0xa1, 0x63, 0x4a, 0x0a, 0x29, 0x90, 0xe3, 0x04, 0xcb, 0x17, 0xca,
		0x8b, 0x27, 0x8f, 0x2b, 0x43, 0x61, 0xc9, 0x43, 0xc7, 0x94, 0x61, 0x34, 0x02, 0x19, 0x26, 0x99,
		0x5f, 0x5f, 0x5f, 0x51, 0xd2, 0x3e, 0x67, 0x75, 0x43, 0x5d, 0x5e, 0x5b, 0x52, 0x32, 0x3e, 0xe7,
		0x92, 0xba, 0xbe, 0x59, 0x51, 0xc0, 0x67, 0x58, 0x2d, 0x57, 0xab, 0xa5, 0xa5, 0xb2, 0x92, 0xf5,
		0x35, 0xe6, 0x9f, 0xdc, 0x28, 0x57, 0x95, 0x5c, 0x68, 0x59, 0x0f, 0x1d, 0x53, 0x46, 0xfc, 0x47,
		0x94, 0xd7, 0x36, 0x57, 0x95, 0x3c, 0x1a, 0x83, 0x11, 0xfe, 0x08, 0xb9, 0x88, 0xd1, 0x2e, 0xd1,
		0xc9, 0xe3, 0x8a, 0xd2, 0x59, 0x08, 0x67, 0x19, 0x0b, 0x09, 0x4e, 0x1e, 0x57, 0xd0, 0xcc, 0x02,
		0xa4, 0x58, 0x74, 0x21, 0x04, 0xf9, 0x95, 0xd2, 0x7c, 0x79, 0x45, 0x5b, 0xaf, 0x6c, 0x2c, 0xaf,
		0xaf, 0x95, 0x56, 0x94, 0x58, 0x47, 0xa6, 0x96, 0xff, 0x73, 0x73, 0x59, 0x2d, 0x2f, 0x2a, 0xf1,
		0xa0, 0xac, 0x52, 0x2e, 0x6d, 0x94, 0x17, 0x95, 0xc4, 0x8c, 0x0e, 0x13, 0xfd, 0xf2, 0x4c, 0xdf,
		0x93, 0x11, 0x70, 0x71, 0x7c, 0x17, 0x17, 0x33, 0xae, 0x1e, 0x17, 0xbf, 0x19, 0x83, 0xf1, 0x3e,
		0xb9, 0xb6, 0xef, 0x43, 0xce, 0x42, 0x8a, 0x87, 0x28, 0xaf, 0x3e, 0xf7, 0xf6, 0x4d, 0xda, 0x2c,
		0x60, 0x7b, 0x2a, 0x10, 0xc3, 0x05, 0x2b, 0x70, 0x62, 0x97, 0x0a, 0x4c, 0x29, 0x7a, 0x16, 0xf9,
		0x6c, 0x0c, 0x0a, 0xbb, 0x71, 0x47, 0x24, 0x8a, 0x78, 0x28, 0x51, 0x9c, 0xe9, 0x5e, 0xc0, 0xa1,
		0xdd, 0xf7, 0xd0, 0xb3, 0x8a, 0xb7, 0x62, 0x30, 0xd9, 0xbf, 0x51, 0xe9, 0xbb, 0x86, 0x47, 0x61,
		0xa8, 0x49, 0xbc, 0x6d, 0x5b, 0x16, 0xeb, 0x7f, 0xeb, 0x53, 0x02, 0xe8, 0x74, 0xb7, 0xad, 0x04,
		0x2a, 0x58, 0x43, 0x12, 0xbb, 0x75, 0x1b, 0x7c, 0x35, 0x3d, 0x2b, 0x7d, 0x2e, 0x0e, 0xb7, 0xf7,
		0x25, 0xef, 0xbb, 0xd0, 0xbb, 0x00, 0x0c, 0xab, 0xd5, 0xf6, 0x78, 0x41, 0xe6, 0xf9, 0x29, 0xc3,
		0x24, 0xec, 0xec, 0xd3, 0xdc, 0xd3, 0xf6, 0xfc, 0xf9, 0x04, 0x9b, 0x07, 0x2e, 0x62, 0x0a, 0xa7,
		0x3a, 0x0b, 0x4d, 0xb2, 0x85, 0x1e, 0xd8, 0x65, 0xa7, 0x3d, 0xb5, 0xee, 0x41, 0x50, 0x74, 0xd3,
		0x20, 0x96, 0xa7, 0xb9, 0x9e, 0x43, 0x70, 0xd3, 0xb0, 0x1a, 0x2c, 0x01, 0xa7, 0x8b, 0xa9, 0x3a,
		0x36, 0x5d, 0xa2, 0x8e, 0xf2, 0xe9, 0xaa, 0x9c, 0xa5, 0x08, 0x56, 0x65, 0x9c, 0x00, 0x62, 0x28,
		0x84, 0xe0, 0xd3, 0x3e, 0x62, 0xe6, 0xf9, 0x61, 0xc8, 0x06, 0xda, 0x3a, 0x74, 0x08, 0x72, 0x17,
		0xf1, 0x65, 0xac, 0xc9, 0x56, 0x9d, 0x5b, 0x22, 0x4b, 0x65, 0x15, 0xd1, 0xae, 0x3f, 0x08, 0x13,
		0x4c, 0xc5, 0x6e, 0x7b, 0xc4, 0xd1, 0x74, 0x13, 0xbb, 0x2e, 0x33, 0x5a, 0x9a, 0xa9, 0x22, 0x3a,
		0xb7, 0x4e, 0xa7, 0x16, 0xe4, 0x0c, 0x3a, 0x01, 0xe3, 0x0c, 0xd1, 0x6c, 0x9b, 0x9e, 0xd1, 0x32,
		0x89, 0x46, 0x5f, 0x1e, 0x5c, 0x96, 0x88, 0xfd, 0x95, 0x8d, 0x51, 0x8d, 0x55, 0xa1, 0x40, 0x57,
		0xe4, 0xa2, 0x25, 0xb8, 0x8b, 0xc1, 0x1a, 0xc4, 0x22, 0x0e, 0xf6, 0x88, 0x46, 0xfe, 0xab, 0x8d,
		0x4d, 0x57, 0xc3, 0x56, 0x4d, 0xdb, 0xc6, 0xee, 0x76, 0x61, 0x22, 0x48, 0xb0, 0x9f, 0xea, 0x2e,
		0x09, 0xd5, 0x32, 0xd3, 0x2c, 0x59, 0xb5, 0xc7, 0xb0, 0xbb, 0x8d, 0x8a, 0x30, 0xc9, 0x88, 0x5c,
		0xcf, 0x31, 0xac, 0x86, 0xa6, 0x6f, 0x13, 0xfd, 0x92, 0xd6, 0xf6, 0xea, 0xa7, 0x0a, 0x77, 0x04,
		0x19, 0xd8, 0x22, 0xab, 0x4c, 0x67, 0x81, 0xaa, 0x6c, 0x7a, 0xf5, 0x53, 0xa8, 0x0a, 0x39, 0xea,
		0x8f, 0xa6, 0x71, 0x8d, 0x68, 0x75, 0xdb, 0x61, 0xc5, 0x25, 0xdf, 0xe7, 0x70, 0x07, 0x8c, 0x38,
		0xb7, 0x2e, 0x00, 0xab, 0x76, 0x8d, 0x14, 0x53, 0xd5, 0x4a, 0xb9, 0xbc, 0xa8, 0x66, 0x25, 0xcb,
		0x39, 0xdb, 0xa1, 0x31, 0xd5, 0xb0, 0x7d, 0x1b, 0x67, 0x79, 0x4c, 0x35, 0x6c, 0x69, 0xe1, 0x13,
		0x30, 0xae, 0xeb, 0x7c, 0xdb, 0x86, 0xae, 0x89, 0x2e, 0xdf, 0x2d, 0x28, 0x21, 0x7b, 0xe9, 0xfa,
		0x12, 0x57, 0x10, 0x61, 0xee, 0xa2, 0xd3, 0x70, 0x7b, 0xc7, 0x5e, 0x41, 0xe0, 0x58, 0xcf, 0x2e,
		0xbb, 0xa1, 0x27, 0x60, 0xbc, 0xb5, 0xd3, 0x0b, 0x44, 0xa1, 0x27, 0xb6, 0x76, 0xba, 0x61, 0xf7,
		0xb0, 0x37, 0x37, 0x87, 0xe8, 0xd8, 0x23, 0xb5, 0xc2, 0xbe, 0xa0, 0x76, 0x60, 0x02, 0x1d, 0x05,
		0x45, 0xd7, 0x35, 0x62, 0xe1, 0x2d, 0x93, 0x68, 0xd8, 0x21, 0x16, 0x76, 0x0b, 0x07, 0x83, 0xca,
		0x79, 0x5d, 0x2f, 0xb3, 0xd9, 0x12, 0x9b, 0x44, 0xf7, 0xc1, 0x98, 0xbd, 0x75, 0x51, 0xe7, 0xc1,
		0xa5, 0xb5, 0x1c, 0x52, 0x37, 0xae, 0x16, 0x0e, 0x33, 0x33, 0x8d, 0xd2, 0x09, 0x16, 0x5a, 0x15,
		0x26, 0x46, 0xf7, 0x82, 0xa2, 0xbb, 0xdb, 0xd8, 0x69, 0xb1, 0xea, 0xee, 0xb6, 0xb0, 0x4e, 0x0a,
		0xf7, 0x70, 0x55, 0x2e, 0x5f, 0x93, 0x62, 0x74, 0x01, 0x26, 0xda, 0x96, 0x61, 0x79, 0xc4, 0x69,
		0x39, 0x84, 0x36, 0xe9, 0xfc, 0xa4, 0x15, 0x7e, 0x33, 0xbc, 0x4b, 0x9b, 0xbd, 0x19, 0xd4, 0xe6,
		0xde, 0x55, 0xc7, 0xdb, 0xbd, 0xc2, 0x99, 0x22, 0xe4, 0x82, 0x4e, 0x47, 0x19, 0xe0, 0x6e, 0x57,
		0x62, 0xb4, 0x86, 0x2e, 0xac, 0x2f, 0xd2, 0xea, 0xf7, 0x54, 0x59, 0x89, 0xd3, 0x2a, 0xbc, 0xb2,
		0xbc, 0x51, 0xd6, 0xd4, 0xcd, 0xb5, 0x8d, 0xe5, 0xd5, 0xb2, 0x92, 0xb8, 0x2f, 0x93, 0xfe, 0xed,
		0xb0, 0x72, 0xfd, 0xfa, 0xf5, 0xeb, 0xf1, 0x99, 0x77, 0xe3, 0x90, 0x0f, 0x77, 0xbe, 0xe8, 0x3f,
		0x60, 0x9f, 0x7c, 0x4d, 0x75, 0x89, 0xa7, 0x5d, 0x31, 0x1c, 0x16, 0x87, 0x4d, 0xcc, 0x7b, 0x47,
		0xdf, 0x84, 0x13, 0x42, 0xab, 0x4a, 0xbc, 0x27, 0x0c, 0x87, 0x46, 0x59, 0x13, 0x7b, 0x68, 0x05,
		0x0e, 0x5a, 0xb6, 0xe6, 0x7a, 0xd8, 0xaa, 0x61, 0xa7, 0xa6, 0x75, 0x2e, 0x08, 0x34, 0xac, 0xeb,
		0xc4, 0x75, 0x6d, 0x5e, 0x02, 0x7c, 0x96, 0x3b, 0x2d, 0xbb, 0x2a, 0x94, 0x3b, 0xb9, 0xb1, 0x24,
		0x54, 0xbb, 0xdc, 0x9d, 0xd8, 0xcd, 0xdd, 0x77, 0x40, 0xa6, 0x89, 0x5b, 0x1a, 0xb1, 0x3c, 0x67,
		0x87, 0xf5, 0x6b, 0x69, 0x35, 0xdd, 0xc4, 0xad, 0x32, 0x1d, 0x7f, 0x72, 0x3e, 0x08, 0xda, 0xf1,
		0xd7, 0x09, 0xc8, 0x05, 0x7b, 0x36, 0xda, 0x02, 0xeb, 0x2c, 0x3f, 0xc7, 0xd8, 0xf1, 0xbd, 0xfb,
		0x96, 0x1d, 0xde, 0xdc, 0x02, 0x4d, 0xdc, 0xc5, 0x21, 0xde, 0x49, 0xa9, 0x1c, 0x49, 0x8b, 0x26,
		0x3d, 0xb0, 0x84, 0xf7, 0xe7, 0x69, 0x55, 0x8c, 0xd0, 0x12, 0x0c, 0x5d, 0x74, 0x19, 0xf7, 0x10,
		0xe3, 0x3e, 0x7c, 0x6b, 0xee, 0xf3, 0x55, 0x46, 0x9e, 0x39, 0x5f, 0xd5, 0xd6, 0xd6, 0xd5, 0xd5,
		0xd2, 0x8a, 0x2a, 0xe0, 0x68, 0x3f, 0x24, 0x4d, 0x7c, 0x6d, 0x27, 0x9c, 0xe2, 0x99, 0x68, 0x50,
		0xc3, 0xef, 0x87, 0xe4, 0x15, 0x82, 0x2f, 0x85, 0x13, 0x2b, 0x13, 0x7d, 0x82, 0xa1, 0x7f, 0x14,
		0x52, 0xcc, 0x5e, 0x08, 0x40, 0x58, 0x4c, 0xb9, 0x0d, 0xa5, 0x21, 0xb9, 0xb0, 0xae, 0xd2, 0xf0,
		0x57, 0x20, 0xc7, 0xa5, 0x5a, 0x65, 0xb9, 0xbc, 0x50, 0x56, 0xe2, 0x33, 0x27, 0x60, 0x88, 0x1b,
		0x81, 0x1e, 0x0d, 0xdf, 0x0c, 0xca, 0x6d, 0x62, 0x28, 0x38, 0x62, 0x72, 0x76, 0x73, 0x75, 0xbe,
		0xac, 0x2a, 0xf1, 0xa0, 0x7b, 0x5d, 0xc8, 0x05, 0xdb, 0xb5, 0x4f, 0x27, 0xa6, 0x7e, 0x16, 0x83,
		0x6c, 0xa0, 0xfd, 0xa2, 0x85, 0x1f, 0x9b, 0xa6, 0x7d, 0x45, 0xc3, 0xa6, 0x81, 0x5d, 0x11, 0x14,
		0xc0, 0x44, 0x25, 0x2a, 0x19, 0xd4, 0x69, 0x9f, 0xca, 0xe2, 0x5f, 0x8b, 0x81, 0xd2, 0xdd, 0xba,
		0x75, 0x2d, 0x30, 0xf6, 0x99, 0x2e, 0xf0, 0x95, 0x18, 0xe4, 0xc3, 0xfd, 0x5a, 0xd7, 0xf2, 0x0e,
		0x7d, 0xa6, 0xcb, 0x7b, 0x39, 0x06, 0x23, 0xa1, 0x2e, 0xed, 0x5f, 0x6a, 0x75, 0x2f, 0x25, 0x60,
		0xbc, 0x0f, 0x0e, 0x95, 0x44, 0x3b, 0xcb, 0x3b, 0xec, 0x07, 0x06, 0x79, 0xd6, 0x1c, 0xad, 0x96,
		0x15, 0xec, 0x78, 0xa2, 0xfb, 0xbd, 0x17, 0x14, 0xa3, 0x46, 0x2c, 0xcf, 0xa8, 0x1b, 0xc4, 0x11,
		0xaf, 0xe0, 0xbc, 0xc7, 0x1d, 0xed, 0xc8, 0xf9, 0x5b, 0xf8, 0xbf, 0x03, 0x6a, 0xd9, 0xae, 0xe1,
		0x19, 0x97, 0x89, 0x66, 0x58, 0xf2, 0x7d, 0x9d, 0xf6, 0xbc, 0x49, 0x55, 0x91, 0x33, 0xcb, 0x96,
		0xe7, 0x6b, 0x5b, 0xa4, 0x81, 0xbb, 0xb4, 0x69, 0xee, 0x4b, 0xa8, 0x8a, 0x9c, 0xf1, 0xb5, 0x0f,
		0x41, 0xae, 0x66, 0xb7, 0x69, 0xfb, 0xc0, 0xf5, 0x68, 0xaa, 0x8d, 0xa9, 0x59, 0x2e, 0xf3, 0x55,
		0x44, 0x7f, 0xd7, 0xb9, 0x28, 0xc8, 0xa9, 0x59, 0x2e, 0xe3, 0x2a, 0x47, 0x60, 0x14, 0x37, 0x1a,
		0x0e, 0x25, 0x97, 0x44, 0xbc, 0x69, 0xcd, 0xfb, 0x62, 0xa6, 0x38, 0x75, 0x1e, 0xd2, 0xd2, 0x0e,
		0xb4, 0x9a, 0x51, 0x4b, 0x68, 0x2d, 0x7e, 0x5d, 0x13, 0x9f, 0xcd, 0xa8, 0x69, 0x4b, 0x4e, 0x1e,
		0x82, 0x9c, 0xe1, 0x6a, 0x9d, 0x7b, 0xc3, 0xf8, 0x74, 0x7c, 0x36, 0xad, 0x66, 0x0d, 0xd7, 0xbf,
		0x28, 0x9a, 0x79, 0x2b, 0x0e, 0xf9, 0xf0, 0xbd, 0x27, 0x5a, 0x84, 0xb4, 0x69, 0xeb, 0x98, 0x05,
		0x02, 0xbf, 0x74, 0x9f, 0x8d, 0xb8, 0x2a, 0x9d, 0x5b, 0x11, 0xfa, 0xaa, 0x8f, 0x9c, 0xfa, 0x55,
		0x0c, 0xd2, 0x52, 0x8c, 0x26, 0x21, 0xd9, 0xc2, 0xde, 0x36, 0xa3, 0x4b, 0xcd, 0xc7, 0x95, 0x98,
		0xca, 0xc6, 0x54, 0xee, 0xb6, 0xb0, 0xc5, 0x42, 0x40, 0xc8, 0xe9, 0x98, 0xfa, 0xd5, 0x24, 0xb8,
		0xc6, 0xda, 0x61, 0xbb, 0xd9, 0x24, 0x96, 0xe7, 0x4a, 0xbf, 0x0a, 0xf9, 0x82, 0x10, 0xa3, 0xfb,
		0x61, 0xcc, 0x73, 0xb0, 0x61, 0x86, 0x74, 0x93, 0x4c, 0x57, 0x91, 0x13, 0xbe, 0x72, 0x11, 0xf6,
		0x4b, 0xde, 0x1a, 0xf1, 0xb0, 0xbe, 0x4d, 0x6a, 0x1d, 0xd0, 0x10, 0xbb, 0x54, 0xdb, 0x27, 0x14,
		0x16, 0xc5, 0xbc, 0xc4, 0xce, 0xbc, 0x17, 0x83, 0x31, 0xd9, 0xc0, 0xd7, 0x7c, 0x63, 0xad, 0x02,
		0x60, 0xcb, 0xb2, 0xbd, 0xa0, 0xb9, 0x7a, 0x43, 0xb9, 0x07, 0x37, 0x57, 0xf2, 0x41, 0x6a, 0x80,
		0x60, 0xaa, 0x09, 0xd0, 0x99, 0xd9, 0xd5, 0x6c, 0x07, 0x21, 0x2b, 0x2e, 0xb5, 0xd9, 0x97, 0x11,
		0xfe, 0xd6, 0x07, 0x5c, 0x44, 0x3b, 0x7d, 0x34, 0x01, 0xa9, 0x2d, 0xd2, 0x30, 0x2c, 0x71, 0xd5,
		0xc6, 0x07, 0xf2, 0x02, 0x2f, 0xe9, 0x5f, 0xe0, 0xcd, 0x3f, 0x0d, 0xe3, 0xba, 0xdd, 0xec, 0x5e,
		0xee, 0xbc, 0xd2, 0xf5, 0xe6, 0xe9, 0x3e, 0x16, 0x7b, 0x0a, 0x3a, 0xdd, 0xd9, 0x1b, 0xb1, 0xd8,
		0x9b, 0xf1, 0xc4, 0x52, 0x65, 0xfe, 0xed, 0xf8, 0xd4, 0x12, 0x87, 0x56, 0xe4, 0x4e, 0x55, 0x52,
		0x37, 0x89, 0x4e, 0x57, 0x0f, 0xaf, 0x1f, 0x86, 0x07, 0x1a, 0x86, 0xb7, 0xdd, 0xde, 0x9a, 0xd3,
		0xed, 0xe6, 0xd1, 0x86, 0xdd, 0xb0, 0x3b, 0x1f, 0x83, 0xe8, 0x88, 0x0d, 0xd8, 0x2f, 0xf1, 0x41,
		0x28, 0xe3, 0x4b, 0xa7, 0x22, 0xbf, 0x1e, 0x15, 0xd7, 0x60, 0x5c, 0x28, 0x6b, 0xec, 0x46, 0x9a,
		0xf7, 0xe1, 0xe8, 0x96, 0xb7, 0x12, 0x85, 0x77, 0x3e, 0x60, 0x95, 0x4e, 0x1d, 0x13, 0x50, 0x3a,
		0xc7, 0x3b, 0xf5, 0xa2, 0x0a, 0xb7, 0x87, 0xf8, 0xf8, 0xd1, 0x24, 0x4e, 0x04, 0xe3, 0xbb, 0x82,
		0x71, 0x3c, 0xc0, 0x58, 0x15, 0xd0, 0xe2, 0x02, 0x8c, 0xec, 0x85, 0xeb, 0x17, 0x82, 0x2b, 0x47,
		0x82, 0x24, 0x4b, 0x30, 0xca, 0x48, 0xf4, 0xb6, 0xeb, 0xd9, 0x4d, 0x96, 0xf7, 0x6e, 0x4d, 0xf3,
		0xcb, 0x0f, 0xf8, 0x59, 0xc9, 0x53, 0xd8, 0x82, 0x8f, 0x2a, 0x3e, 0x0e, 0x13, 0x54, 0xc2, 0x52,
		0x4b, 0x90, 0x2d, 0xfa, 0x1e, 0xa5, 0xf0, 0xde, 0xb3, 0xfc, 0x48, 0x8d, 0xfb, 0x04, 0x01, 0xde,
		0x80, 0x27, 0x1a, 0xc4, 0xf3, 0x88, 0xe3, 0x6a, 0xd8, 0x34, 0xd1, 0x2d, 0xbf, 0xd0, 0x14, 0x5e,
		0xfc, 0x30, 0xec, 0x89, 0x25, 0x8e, 0x2c, 0x99, 0x66, 0x71, 0x13, 0xf6, 0xf5, 0xf1, 0xec, 0x00,
		0x9c, 0x2f, 0x09, 0xce, 0x89, 0x1e, 0xef, 0x52, 0xda, 0x0a, 0x48, 0xb9, 0xef, 0x8f, 0x01, 0x38,
		0x5f, 0x16, 0x9c, 0x48, 0x60, 0xa5, 0x5b, 0x28, 0xe3, 0x79, 0x18, 0xbb, 0x4c, 0x9c, 0x2d, 0xdb,
		0x15, 0x2f, 0xff, 0x03, 0xd0, 0xbd, 0x22, 0xe8, 0x46, 0x05, 0x90, 0x5d, 0x05, 0x50, 0xae, 0xd3,
		0x90, 0xae, 0x63, 0x9d, 0x0c, 0x40, 0xf1, 0xaa, 0xa0, 0x18, 0xa6, 0xfa, 0x14, 0x5a, 0x82, 0x5c,
		0xc3, 0x16, 0xd5, 0x25, 0x1a, 0xfe, 0x9a, 0x80, 0x67, 0x25, 0x46, 0x50, 0xb4, 0xec, 0x56, 0xdb,
		0xa4, 0xa5, 0x27, 0x9a, 0xe2, 0x75, 0x49, 0x21, 0x31, 0x82, 0x62, 0x0f, 0x66, 0x7d, 0x43, 0x52,
		0xb8, 0x01, 0x7b, 0x9e, 0x85, 0xac, 0x6d, 0x99, 0x3b, 0xb6, 0x35, 0xc8, 0x22, 0xbe, 0x26, 0x18,
		0x40, 0x40, 0x28, 0xc1, 0x19, 0xc8, 0x0c, 0xea, 0x88, 0xaf, 0x0b, 0x78, 0x9a, 0x48, 0x0f, 0x2c,
		0xc1, 0xa8, 0x4c, 0x32, 0x86, 0x6d, 0x0d, 0x40, 0xf1, 0x0d, 0x41, 0x91, 0x0f, 0xc0, 0xc4, 0x36,
		0x3c, 0xe2, 0x7a, 0x0d, 0x32, 0x08, 0xc9, 0x5b, 0x72, 0x1b, 0x02, 0x22, 0x4c, 0xb9, 0x45, 0x2c,
		0x7d, 0x7b, 0x30, 0x86, 0x6f, 0x4a, 0x53, 0x4a, 0x0c, 0xa5, 0x58, 0x80, 0x91, 0x26, 0x76, 0xdc,
		0x6d, 0x6c, 0x0e, 0xe4, 0x8e, 0x6f, 0x09, 0x8e, 0x9c, 0x0f, 0x12, 0x16, 0x69, 0x5b, 0x7b, 0xa1,
		0x79, 0x5b, 0x5a, 0x24, 0x00, 0x13, 0x47, 0xcf, 0xf5, 0xd8, 0xfd, 0xca, 0x5e, 0xd8, 0xbe, 0x2d,
		0x8f, 0x1e, 0xc7, 0xae, 0x06, 0x19, 0xcf, 0x40, 0xc6, 0x35, 0xae, 0x0d, 0x44, 0xf3, 0x1d, 0xe9,
		0x69, 0x06, 0xa0, 0xe0, 0x27, 0x61, 0x7f, 0xdf, 0x54, 0x3f, 0x00, 0xd9, 0x77, 0x05, 0xd9, 0x64,
		0x9f, 0x74, 0x2f, 0x52, 0xc2, 0x5e, 0x29, 0xbf, 0x27, 0x53, 0x02, 0xe9, 0xe2, 0xaa, 0xd0, 0xee,
		0xdc, 0xc5, 0xf5, 0xbd, 0x59, 0xed, 0xfb, 0xd2, 0x6a, 0x1c, 0x1b, 0xb2, 0xda, 0x06, 0x4c, 0x0a,
		0xc6, 0xbd, 0xf9, 0xf5, 0x07, 0x32, 0xb1, 0x72, 0xf4, 0x66, 0xd8, 0xbb, 0x4f, 0xc3, 0x94, 0x6f,
		0x4e, 0xd9, 0x58, 0xba, 0x5a, 0x13, 0xb7, 0x06, 0x60, 0x7e, 0x47, 0x30, 0xcb, 0x8c, 0xef, 0x77,
		0xa6, 0xee, 0x2a, 0x6e, 0x51, 0xf2, 0x0b, 0x50, 0x90, 0xe4, 0x6d, 0xcb, 0x21, 0xba, 0xdd, 0xb0,
		0x8c, 0x6b, 0xa4, 0x36, 0x00, 0xf5, 0x0f, 0xbb, 0x5c, 0xb5, 0x19, 0x80, 0x53, 0xe6, 0x65, 0x50,
		0xfc, 0x7e, 0x43, 0x33, 0x9a, 0x2d, 0xdb, 0xf1, 0x22, 0x18, 0x7f, 0x24, 0x3d, 0xe5, 0xe3, 0x96,
		0x19, 0xac, 0x58, 0x86, 0x3c, 0x1b, 0x0e, 0x1a, 0x92, 0x3f, 0x16, 0x44, 0x23, 0x1d, 0x94, 0x48,
		0x1c, 0xba, 0xdd, 0x6c, 0x61, 0x67, 0x90, 0xfc, 0xf7, 0x13, 0x99, 0x38, 0x04, 0x84, 0x47, 0xdf,
		0x68, 0x57, 0x25, 0x46, 0x51, 0x1f, 0xaf, 0x0b, 0xcf, 0xdc, 0x10, 0x67, 0x36, 0x5c, 0x88, 0x8b,
		0x2b, 0xd4, 0x3c, 0xe1, 0x72, 0x19, 0x4d, 0xf6, 0xec, 0x0d, 0xdf, 0x42, 0xa1, 0x6a, 0x59, 0x3c,
		0x07, 0x23, 0xa1, 0x52, 0x19, 0x4d, 0xf5, 0xdf, 0x82, 0x2a, 0x17, 0xac, 0x94, 0xc5, 0x13, 0x90,
		0xa4, 0x65, 0x2f, 0x1a, 0xfe, 0x3f, 0x02, 0xce, 0xd4, 0x8b, 0x8f, 0x40, 0x5a, 0x96, 0xbb, 0x68,
		0xe8, 0xff, 0x0a, 0xa8, 0x0f, 0xa1, 0x70, 0x59, 0xea, 0xa2, 0xe1, 0xff, 0x27, 0xe1, 0x12, 0x42,
		0xe1, 0x83, 0x9b, 0xf0, 0xe7, 0xcf, 0x27, 0x45, 0xba, 0x92, 0xb6, 0x3b, 0x03, 0xc3, 0xa2, 0xc6,
		0x45, 0xa3, 0x9f, 0x13, 0x0f, 0x97, 0x88, 0xe2, 0xc3, 0x90, 0x1a, 0xd0, 0xe0, 0xff, 0x2f, 0xa0,
		0x5c, 0xbf, 0xb8, 0x00, 0xd9, 0x40, 0x5d, 0x8b, 0x86, 0x7f, 0x4e, 0xc0, 0x83, 0x28, 0xba, 0x74,
		0x51, 0xd7, 0xa2, 0x09, 0x3e, 0x2f, 0x97, 0x2e, 0x10, 0xd4, 0x6c, 0xb2, 0xa4, 0x45, 0xa3, 0xbf,
		0x20, 0xad, 0x2e, 0x21, 0xc5, 0xb3, 0x90, 0xf1, 0xd3, 0x54, 0x34, 0xfe, 0x8b, 0x02, 0xdf, 0xc1,
		0x50, 0x0b, 0x04, 0xd2, 0x64, 0x34, 0xc5, 0x97, 0xa4, 0x05, 0x02, 0x28, 0x7a, 0x8c, 0xba, 0x4b,
		0x5f, 0x34, 0xd3, 0x97, 0xe5, 0x31, 0xea, 0xaa, 0x7c, 0xd4, 0x9b, 0x2c, 0x5b, 0x44, 0x53, 0x7c,
		0x45, 0x7a, 0x93, 0xe9, 0xd3, 0x65, 0x74, 0xd7, 0x92, 0x68, 0x8e, 0xaf, 0xca, 0x65, 0x74, 0x95,
		0x92, 0x62, 0x05, 0x50, 0x6f, 0x1d, 0x89, 0xe6, 0x7b, 0x41, 0xf0, 0x8d, 0xf5, 0x94, 0x91, 0xe2,
		0x13, 0x30, 0xd9, 0xbf, 0x86, 0x44, 0xb3, 0xbe, 0x78, 0xa3, 0xab, 0xeb, 0x0f, 0x96, 0x90, 0xe2,
		0x46, 0xa7, 0xeb, 0x0f, 0xd6, 0x8f, 0x68, 0xda, 0x97, 0x6e, 0x84, 0x5f, 0xec, 0x82, 0xe5, 0xa3,
		0x58, 0x02, 0xe8, 0xa4, 0xee, 0x68, 0xae, 0x57, 0x04, 0x57, 0x00, 0x44, 0x8f, 0x86, 0xc8, 0xdc,
		0xd1, 0xf8, 0x57, 0xe5, 0xd1, 0x10, 0x88, 0xe2, 0x19, 0x48, 0x5b, 0x6d, 0xd3, 0xa4, 0xc1, 0x81,
		0x6e, 0xfd, 0x0f, 0x21, 0x85, 0xdf, 0xdd, 0x14, 0x07, 0x43, 0x02, 0x8a, 0x27, 0x20, 0x45, 0x9a,
		0x5b, 0xa4, 0x16, 0x85, 0xfc, 0xfd, 0x4d, 0x99, 0x10, 0xa8, 0x76, 0xf1, 0x2c, 0x00, 0x7f, 0x69,
		0x64, 0xdf, 0x03, 0x22, 0xb0, 0x7f, 0xb8, 0x29, 0xbe, 0x35, 0x77, 0x20, 0x1d, 0x02, 0xfe, 0xe5,
		0xfa, 0xd6, 0x04, 0x1f, 0x86, 0x09, 0xd8, 0x8b, 0xe6, 0x69, 0x18, 0xbe, 0xe8, 0xda, 0x96, 0x87,
		0x1b, 0x51, 0xe8, 0x3f, 0x0a, 0xb4, 0xd4, 0xa7, 0x06, 0x6b, 0xda, 0x0e, 0xf1, 0x70, 0xc3, 0x8d,
		0xc2, 0xfe, 0x49, 0x60, 0x7d, 0x00, 0x05, 0xeb, 0xd8, 0xf5, 0x06, 0xd9, 0xf7, 0x9f, 0x25, 0x58,
		0x02, 0xe8, 0xa2, 0xe9, 0xef, 0x4b, 0x64, 0x27, 0x0a, 0xfb, 0x91, 0x5c, 0xb4, 0xd0, 0x2f, 0x3e,
		0x02, 0x19, 0xfa, 0x93, 0xff, 0xff, 0x45, 0x04, 0xf8, 0x2f, 0x02, 0xdc, 0x41, 0xd0, 0x27, 0xbb,
		0x5e, 0xcd, 0x33, 0xa2, 0x8d, 0xfd, 0x57, 0xe1, 0x69, 0xa9, 0x5f, 0x2c, 0x41, 0xd6, 0xf5, 0x6a,
		0xb5, 0xb6, 0xc3, 0x2f, 0xa2, 0x22, 0xe0, 0x7f, 0xbb, 0xe9, 0xbf, 0xcc, 0xf9, 0x98, 0xf9, 0x43,
		0xfd, 0xef, 0x96, 0x60, 0xc9, 0x5e, 0xb2, 0xf9, 0xad, 0x12, 0xbc, 0x1d, 0x83, 0x11, 0xd7, 0xd1,
		0x3d, 0x87, 0x08, 0x05, 0x94, 0x6c, 0x62, 0xc3, 0x9a, 0xda, 0xdb, 0xb5, 0xd1, 0xcc, 0x53, 0x30,
		0x5c, 0x75, 0xf4, 0x0d, 0x87, 0x10, 0x34, 0x0d, 0x59, 0xf1, 0x99, 0x7a, 0xad, 0xf3, 0x9f, 0x12,
		0x41, 0x11, 0x3a, 0x02, 0xc3, 0xbc, 0x69, 0x73, 0xc5, 0xc5, 0xf3, 0xc8, 0x1c, 0x7d, 0xe6, 0x9c,
		0x60, 0x50, 0xe5, 0x6c, 0x31, 0xf9, 0xd1, 0x1b, 0x07, 0x63, 0xf3, 0xe9, 0x8f, 0xde, 0x3f, 0x10,
		0xfb, 0xc7, 0xfb, 0x07, 0x62, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xc3, 0x22, 0xc9, 0xd3,
		0x2c, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *SrcTree) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&main.SrcTree{")
	if this.PackageName != nil {
		s = append(s, "PackageName: "+valueToGoStringSrctree(this.PackageName, "string")+",\n")
	}
	if this.Imports != nil {
		s = append(s, "Imports: "+fmt.Sprintf("%#v", this.Imports)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSrctree(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedSrcTree(r randySrctree, easy bool) *SrcTree {
	this := &SrcTree{}
	if r.Intn(10) != 0 {
		v1 := string(randStringSrctree(r))
		this.PackageName = &v1
	}
	if r.Intn(10) == 0 {
		v2 := r.Intn(5)
		this.Imports = make([]*SrcTree, v2)
		for i := 0; i < v2; i++ {
			this.Imports[i] = NewPopulatedSrcTree(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSrctree(r, 3)
	}
	return this
}

type randySrctree interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSrctree(r randySrctree) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSrctree(r randySrctree) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneSrctree(r)
	}
	return string(tmps)
}
func randUnrecognizedSrctree(r randySrctree, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSrctree(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSrctree(dAtA []byte, r randySrctree, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSrctree(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSrctree(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("srctree.proto", fileDescriptorSrctree) }

var fileDescriptorSrctree = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x4a, 0x2e,
	0x29, 0x4a, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93,
	0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf,
	0xd7, 0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0x14, 0xc5,
	0xc5, 0x1e, 0x5c, 0x94, 0x1c, 0x52, 0x94, 0x9a, 0x2a, 0xa4, 0xc0, 0xc5, 0x1d, 0x90, 0x98, 0x9c,
	0x9d, 0x98, 0x9e, 0xea, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2c,
	0x24, 0xa4, 0xce, 0xc5, 0xee, 0x99, 0x5b, 0x90, 0x5f, 0x54, 0x52, 0x2c, 0xc1, 0xa4, 0xc0, 0xac,
	0xc1, 0x6d, 0xc4, 0xab, 0x07, 0xb2, 0x53, 0x0f, 0x6a, 0x42, 0x10, 0x4c, 0xd6, 0x8a, 0xe5, 0xc3,
	0x02, 0x79, 0x46, 0x27, 0x8e, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x04, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x84, 0x99, 0x13, 0x2f, 0xaa, 0x00, 0x00, 0x00,
}
