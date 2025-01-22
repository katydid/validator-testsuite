// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: banana.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BananaSource int32

const (
	BananaSource_BananaSource_q      BananaSource = 0
	BananaSource_BananaSource_w      BananaSource = 1
	BananaSource_BananaSource_e      BananaSource = 2
	BananaSource_BananaSource_r      BananaSource = 3
	BananaSource_BananaSource_t      BananaSource = 4
	BananaSource_BananaSource_y      BananaSource = 5
	BananaSource_BananaSource_u      BananaSource = 6
	BananaSource_BananaSource_i      BananaSource = 7
	BananaSource_BananaSource_o      BananaSource = 8
	BananaSource_BananaSource_p      BananaSource = 9
	BananaSource_BananaSource_a      BananaSource = 10
	BananaSource_BananaSource_s      BananaSource = 11
	BananaSource_BananaSource_d      BananaSource = 12
	BananaSource_BananaSource_f      BananaSource = 13
	BananaSource_BananaSource_g      BananaSource = 14
	BananaSource_BananaSource_h      BananaSource = 15
	BananaSource_BananaSource_j      BananaSource = 16
	BananaSource_BananaSource_k      BananaSource = 17
	BananaSource_BananaSource_l      BananaSource = 18
	BananaSource_BananaSource_z      BananaSource = 19
	BananaSource_BananaSource_x      BananaSource = 20
	BananaSource_BananaSource_c      BananaSource = 21
	BananaSource_BananaSource_v      BananaSource = 22
	BananaSource_BananaSource_b      BananaSource = 23
	BananaSource_BananaSource_n      BananaSource = 24
	BananaSource_BananaSourceMaximum BananaSource = 25
)

// Enum value maps for BananaSource.
var (
	BananaSource_name = map[int32]string{
		0:  "BananaSource_q",
		1:  "BananaSource_w",
		2:  "BananaSource_e",
		3:  "BananaSource_r",
		4:  "BananaSource_t",
		5:  "BananaSource_y",
		6:  "BananaSource_u",
		7:  "BananaSource_i",
		8:  "BananaSource_o",
		9:  "BananaSource_p",
		10: "BananaSource_a",
		11: "BananaSource_s",
		12: "BananaSource_d",
		13: "BananaSource_f",
		14: "BananaSource_g",
		15: "BananaSource_h",
		16: "BananaSource_j",
		17: "BananaSource_k",
		18: "BananaSource_l",
		19: "BananaSource_z",
		20: "BananaSource_x",
		21: "BananaSource_c",
		22: "BananaSource_v",
		23: "BananaSource_b",
		24: "BananaSource_n",
		25: "BananaSourceMaximum",
	}
	BananaSource_value = map[string]int32{
		"BananaSource_q":      0,
		"BananaSource_w":      1,
		"BananaSource_e":      2,
		"BananaSource_r":      3,
		"BananaSource_t":      4,
		"BananaSource_y":      5,
		"BananaSource_u":      6,
		"BananaSource_i":      7,
		"BananaSource_o":      8,
		"BananaSource_p":      9,
		"BananaSource_a":      10,
		"BananaSource_s":      11,
		"BananaSource_d":      12,
		"BananaSource_f":      13,
		"BananaSource_g":      14,
		"BananaSource_h":      15,
		"BananaSource_j":      16,
		"BananaSource_k":      17,
		"BananaSource_l":      18,
		"BananaSource_z":      19,
		"BananaSource_x":      20,
		"BananaSource_c":      21,
		"BananaSource_v":      22,
		"BananaSource_b":      23,
		"BananaSource_n":      24,
		"BananaSourceMaximum": 25,
	}
)

func (x BananaSource) Enum() *BananaSource {
	p := new(BananaSource)
	*p = x
	return p
}

func (x BananaSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BananaSource) Descriptor() protoreflect.EnumDescriptor {
	return file_banana_proto_enumTypes[0].Descriptor()
}

func (BananaSource) Type() protoreflect.EnumType {
	return &file_banana_proto_enumTypes[0]
}

func (x BananaSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BananaSource) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BananaSource(num)
	return nil
}

// Deprecated: Use BananaSource.Descriptor instead.
func (BananaSource) EnumDescriptor() ([]byte, []int) {
	return file_banana_proto_rawDescGZIP(), []int{0}
}

type PeelType int32

const (
	PeelType_PeelFirm        PeelType = 0
	PeelType_PeelGreen       PeelType = 1
	PeelType_PeelYellow      PeelType = 2
	PeelType_PeelDarkened    PeelType = 3
	PeelType_PeelDecomposing PeelType = 4
)

// Enum value maps for PeelType.
var (
	PeelType_name = map[int32]string{
		0: "PeelFirm",
		1: "PeelGreen",
		2: "PeelYellow",
		3: "PeelDarkened",
		4: "PeelDecomposing",
	}
	PeelType_value = map[string]int32{
		"PeelFirm":        0,
		"PeelGreen":       1,
		"PeelYellow":      2,
		"PeelDarkened":    3,
		"PeelDecomposing": 4,
	}
)

func (x PeelType) Enum() *PeelType {
	p := new(PeelType)
	*p = x
	return p
}

func (x PeelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeelType) Descriptor() protoreflect.EnumDescriptor {
	return file_banana_proto_enumTypes[1].Descriptor()
}

func (PeelType) Type() protoreflect.EnumType {
	return &file_banana_proto_enumTypes[1]
}

func (x PeelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PeelType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PeelType(num)
	return nil
}

// Deprecated: Use PeelType.Descriptor instead.
func (PeelType) EnumDescriptor() ([]byte, []int) {
	return file_banana_proto_rawDescGZIP(), []int{1}
}

type BananaTuple struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BananaTime    *int64                 `protobuf:"varint,1,opt,name=BananaTime" json:"BananaTime,omitempty"`
	Pulpy         *string                `protobuf:"bytes,2,opt,name=Pulpy" json:"Pulpy,omitempty"`
	IsPulpySet    *bool                  `protobuf:"varint,3,opt,name=IsPulpySet" json:"IsPulpySet,omitempty"`
	Yummy         *string                `protobuf:"bytes,4,opt,name=Yummy" json:"Yummy,omitempty"`
	IsYummySet    *bool                  `protobuf:"varint,5,opt,name=IsYummySet" json:"IsYummySet,omitempty"`
	Chewy         *string                `protobuf:"bytes,6,opt,name=Chewy" json:"Chewy,omitempty"`
	IsChewySet    *bool                  `protobuf:"varint,7,opt,name=IsChewySet" json:"IsChewySet,omitempty"`
	Squishy       *string                `protobuf:"bytes,8,opt,name=Squishy" json:"Squishy,omitempty"`
	IsSquishySet  *bool                  `protobuf:"varint,9,opt,name=IsSquishySet" json:"IsSquishySet,omitempty"`
	Brown         *string                `protobuf:"bytes,10,opt,name=Brown" json:"Brown,omitempty"`
	IsBrownSet    *bool                  `protobuf:"varint,11,opt,name=IsBrownSet" json:"IsBrownSet,omitempty"`
	Rotten        *uint32                `protobuf:"varint,12,opt,name=Rotten" json:"Rotten,omitempty"`
	IsRottenSet   *bool                  `protobuf:"varint,13,opt,name=IsRottenSet" json:"IsRottenSet,omitempty"`
	Peel          *uint32                `protobuf:"varint,14,opt,name=Peel" json:"Peel,omitempty"`
	IsPeelSet     *bool                  `protobuf:"varint,15,opt,name=IsPeelSet" json:"IsPeelSet,omitempty"`
	PeelType      *PeelType              `protobuf:"varint,16,opt,name=PeelType,enum=main.PeelType" json:"PeelType,omitempty"`
	IsPeelTypeSet *bool                  `protobuf:"varint,17,opt,name=IsPeelTypeSet" json:"IsPeelTypeSet,omitempty"`
	BananaSource  *BananaSource          `protobuf:"varint,18,opt,name=BananaSource,enum=main.BananaSource" json:"BananaSource,omitempty"`
	Slippery      *bool                  `protobuf:"varint,19,opt,name=Slippery" json:"Slippery,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BananaTuple) Reset() {
	*x = BananaTuple{}
	mi := &file_banana_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BananaTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BananaTuple) ProtoMessage() {}

func (x *BananaTuple) ProtoReflect() protoreflect.Message {
	mi := &file_banana_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BananaTuple.ProtoReflect.Descriptor instead.
func (*BananaTuple) Descriptor() ([]byte, []int) {
	return file_banana_proto_rawDescGZIP(), []int{0}
}

func (x *BananaTuple) GetBananaTime() int64 {
	if x != nil && x.BananaTime != nil {
		return *x.BananaTime
	}
	return 0
}

func (x *BananaTuple) GetPulpy() string {
	if x != nil && x.Pulpy != nil {
		return *x.Pulpy
	}
	return ""
}

func (x *BananaTuple) GetIsPulpySet() bool {
	if x != nil && x.IsPulpySet != nil {
		return *x.IsPulpySet
	}
	return false
}

func (x *BananaTuple) GetYummy() string {
	if x != nil && x.Yummy != nil {
		return *x.Yummy
	}
	return ""
}

func (x *BananaTuple) GetIsYummySet() bool {
	if x != nil && x.IsYummySet != nil {
		return *x.IsYummySet
	}
	return false
}

func (x *BananaTuple) GetChewy() string {
	if x != nil && x.Chewy != nil {
		return *x.Chewy
	}
	return ""
}

func (x *BananaTuple) GetIsChewySet() bool {
	if x != nil && x.IsChewySet != nil {
		return *x.IsChewySet
	}
	return false
}

func (x *BananaTuple) GetSquishy() string {
	if x != nil && x.Squishy != nil {
		return *x.Squishy
	}
	return ""
}

func (x *BananaTuple) GetIsSquishySet() bool {
	if x != nil && x.IsSquishySet != nil {
		return *x.IsSquishySet
	}
	return false
}

func (x *BananaTuple) GetBrown() string {
	if x != nil && x.Brown != nil {
		return *x.Brown
	}
	return ""
}

func (x *BananaTuple) GetIsBrownSet() bool {
	if x != nil && x.IsBrownSet != nil {
		return *x.IsBrownSet
	}
	return false
}

func (x *BananaTuple) GetRotten() uint32 {
	if x != nil && x.Rotten != nil {
		return *x.Rotten
	}
	return 0
}

func (x *BananaTuple) GetIsRottenSet() bool {
	if x != nil && x.IsRottenSet != nil {
		return *x.IsRottenSet
	}
	return false
}

func (x *BananaTuple) GetPeel() uint32 {
	if x != nil && x.Peel != nil {
		return *x.Peel
	}
	return 0
}

func (x *BananaTuple) GetIsPeelSet() bool {
	if x != nil && x.IsPeelSet != nil {
		return *x.IsPeelSet
	}
	return false
}

func (x *BananaTuple) GetPeelType() PeelType {
	if x != nil && x.PeelType != nil {
		return *x.PeelType
	}
	return PeelType_PeelFirm
}

func (x *BananaTuple) GetIsPeelTypeSet() bool {
	if x != nil && x.IsPeelTypeSet != nil {
		return *x.IsPeelTypeSet
	}
	return false
}

func (x *BananaTuple) GetBananaSource() BananaSource {
	if x != nil && x.BananaSource != nil {
		return *x.BananaSource
	}
	return BananaSource_BananaSource_q
}

func (x *BananaTuple) GetSlippery() bool {
	if x != nil && x.Slippery != nil {
		return *x.Slippery
	}
	return false
}

var File_banana_proto protoreflect.FileDescriptor

var file_banana_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0xd5, 0x04, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x75, 0x6c, 0x70, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x75, 0x6c, 0x70, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73,
	0x50, 0x75, 0x6c, 0x70, 0x79, 0x53, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x49, 0x73, 0x50, 0x75, 0x6c, 0x70, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x59, 0x75,
	0x6d, 0x6d, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x59, 0x75, 0x6d, 0x6d, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x59, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x59, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x77, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x43, 0x68, 0x65, 0x77, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x43, 0x68, 0x65, 0x77,
	0x79, 0x53, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x43, 0x68,
	0x65, 0x77, 0x79, 0x53, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68, 0x79, 0x53, 0x65, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49, 0x73, 0x53, 0x71, 0x75, 0x69, 0x73, 0x68,
	0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73,
	0x42, 0x72, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x49, 0x73, 0x42, 0x72, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x74, 0x74, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x74, 0x74,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x52, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x53, 0x65,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73, 0x52, 0x6f, 0x74, 0x74, 0x65,
	0x6e, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x50, 0x65, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x50, 0x65,
	0x65, 0x6c, 0x53, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x50,
	0x65, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x50, 0x65, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x50, 0x65, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x73, 0x50, 0x65, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x73, 0x50, 0x65, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x42, 0x61, 0x6e, 0x61,
	0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x0c, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x79, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x53, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x72, 0x79, 0x2a, 0x9b, 0x04, 0x0a,
	0x0c, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x77, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x10,
	0x04, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x79, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x10, 0x07, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x10,
	0x08, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x70, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x10, 0x0b, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x10,
	0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x66, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x10, 0x0f, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6a, 0x10,
	0x10, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6b, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x7a, 0x10, 0x13, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x78, 0x10,
	0x14, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x10, 0x16, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x10, 0x17, 0x12, 0x12, 0x0a,
	0x0e, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x10,
	0x18, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x10, 0x19, 0x2a, 0x5e, 0x0a, 0x08, 0x50, 0x65,
	0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x6c, 0x46, 0x69,
	0x72, 0x6d, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x65, 0x65, 0x6c, 0x47, 0x72, 0x65, 0x65,
	0x6e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x6c, 0x59, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x6c, 0x44, 0x61, 0x72, 0x6b, 0x65,
	0x6e, 0x65, 0x64, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x65, 0x65, 0x6c, 0x44, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x74, 0x79, 0x64, 0x69, 0x64,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6e,
}

var (
	file_banana_proto_rawDescOnce sync.Once
	file_banana_proto_rawDescData = file_banana_proto_rawDesc
)

func file_banana_proto_rawDescGZIP() []byte {
	file_banana_proto_rawDescOnce.Do(func() {
		file_banana_proto_rawDescData = protoimpl.X.CompressGZIP(file_banana_proto_rawDescData)
	})
	return file_banana_proto_rawDescData
}

var file_banana_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_banana_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_banana_proto_goTypes = []any{
	(BananaSource)(0),   // 0: main.BananaSource
	(PeelType)(0),       // 1: main.PeelType
	(*BananaTuple)(nil), // 2: main.BananaTuple
}
var file_banana_proto_depIdxs = []int32{
	1, // 0: main.BananaTuple.PeelType:type_name -> main.PeelType
	0, // 1: main.BananaTuple.BananaSource:type_name -> main.BananaSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_banana_proto_init() }
func file_banana_proto_init() {
	if File_banana_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_banana_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_banana_proto_goTypes,
		DependencyIndexes: file_banana_proto_depIdxs,
		EnumInfos:         file_banana_proto_enumTypes,
		MessageInfos:      file_banana_proto_msgTypes,
	}.Build()
	File_banana_proto = out.File
	file_banana_proto_rawDesc = nil
	file_banana_proto_goTypes = nil
	file_banana_proto_depIdxs = nil
}
