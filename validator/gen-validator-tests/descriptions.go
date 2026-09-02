package main

import (
	descriptor "google.golang.org/protobuf/types/descriptorpb"
	"katydid.org.za/go/parser-go-proto/proto/desc"
)

func (this *BananaTuple) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_banana_proto)
}

func (this *SrcTree) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_srctree_proto)
}

func (this *Person) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_person_proto)
}

func (this *Knot) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_knot_proto)
}

func (this *PuddingMilkshake) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_puddingmilkshake_proto)
}

func (this *FinanceJudo) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_puddingmilkshake_proto)
}

func (this *TopsyTurvy) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_topsyturvy_proto)
}

func (this *TypewriterPrison) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_typewriterprison_proto)
}

func (this *PocketRoses) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_typewriterprison_proto)
}

func (this *Conf) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_conf_proto)
}

func (this *Date) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_conf_proto)
}

func (this *Location) Description() *descriptor.FileDescriptorSet {
	return desc.NewFileDescriptorSet(File_conf_proto)
}
