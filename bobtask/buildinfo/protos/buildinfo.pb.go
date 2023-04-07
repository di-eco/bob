// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: buildinfo.proto

package protos

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

type BuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Targets `protobuf:"bytes,1,opt,name=Target,proto3" json:"Target,omitempty"`
	Meta   *Meta    `protobuf:"bytes,2,opt,name=Meta,proto3" json:"Meta,omitempty"`
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{0}
}

func (x *BuildInfo) GetTarget() *Targets {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *BuildInfo) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task      string `protobuf:"bytes,1,opt,name=Task,proto3" json:"Task,omitempty"`
	InputHash string `protobuf:"bytes,2,opt,name=InputHash,proto3" json:"InputHash,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Meta) GetInputHash() string {
	if x != nil {
		return x.InputHash
	}
	return ""
}

type Targets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filesystem *BuildInfoFiles             `protobuf:"bytes,1,opt,name=Filesystem,proto3" json:"Filesystem,omitempty"`
	Docker     map[string]*BuildInfoDocker `protobuf:"bytes,2,rep,name=Docker,proto3" json:"Docker,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Targets) Reset() {
	*x = Targets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Targets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Targets) ProtoMessage() {}

func (x *Targets) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Targets.ProtoReflect.Descriptor instead.
func (*Targets) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{2}
}

func (x *Targets) GetFilesystem() *BuildInfoFiles {
	if x != nil {
		return x.Filesystem
	}
	return nil
}

func (x *Targets) GetDocker() map[string]*BuildInfoDocker {
	if x != nil {
		return x.Docker
	}
	return nil
}

type BuildInfoFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    string                    `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Targets map[string]*BuildInfoFile `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BuildInfoFiles) Reset() {
	*x = BuildInfoFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfoFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfoFiles) ProtoMessage() {}

func (x *BuildInfoFiles) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfoFiles.ProtoReflect.Descriptor instead.
func (*BuildInfoFiles) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{3}
}

func (x *BuildInfoFiles) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BuildInfoFiles) GetTargets() map[string]*BuildInfoFile {
	if x != nil {
		return x.Targets
	}
	return nil
}

type BuildInfoFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64  `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *BuildInfoFile) Reset() {
	*x = BuildInfoFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfoFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfoFile) ProtoMessage() {}

func (x *BuildInfoFile) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfoFile.ProtoReflect.Descriptor instead.
func (*BuildInfoFile) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{4}
}

func (x *BuildInfoFile) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BuildInfoFile) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type BuildInfoDocker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *BuildInfoDocker) Reset() {
	*x = BuildInfoDocker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildinfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfoDocker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfoDocker) ProtoMessage() {}

func (x *BuildInfoDocker) ProtoReflect() protoreflect.Message {
	mi := &file_buildinfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfoDocker.ProtoReflect.Descriptor instead.
func (*BuildInfoDocker) Descriptor() ([]byte, []int) {
	return file_buildinfo_proto_rawDescGZIP(), []int{5}
}

func (x *BuildInfoDocker) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_buildinfo_proto protoreflect.FileDescriptor

var file_buildinfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x62, 0x6f, 0x62, 0x22, 0x50, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x33,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x44,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x4f, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3a, 0x0a,
	0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x62, 0x6f, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x4e, 0x0a, 0x0c, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x62,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x0d, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x25, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x44,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x6f, 0x62,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buildinfo_proto_rawDescOnce sync.Once
	file_buildinfo_proto_rawDescData = file_buildinfo_proto_rawDesc
)

func file_buildinfo_proto_rawDescGZIP() []byte {
	file_buildinfo_proto_rawDescOnce.Do(func() {
		file_buildinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildinfo_proto_rawDescData)
	})
	return file_buildinfo_proto_rawDescData
}

var file_buildinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_buildinfo_proto_goTypes = []interface{}{
	(*BuildInfo)(nil),       // 0: bob.BuildInfo
	(*Meta)(nil),            // 1: bob.Meta
	(*Targets)(nil),         // 2: bob.Targets
	(*BuildInfoFiles)(nil),  // 3: bob.BuildInfoFiles
	(*BuildInfoFile)(nil),   // 4: bob.BuildInfoFile
	(*BuildInfoDocker)(nil), // 5: bob.BuildInfoDocker
	nil,                     // 6: bob.Targets.DockerEntry
	nil,                     // 7: bob.BuildInfoFiles.TargetsEntry
}
var file_buildinfo_proto_depIdxs = []int32{
	2, // 0: bob.BuildInfo.Target:type_name -> bob.Targets
	1, // 1: bob.BuildInfo.Meta:type_name -> bob.Meta
	3, // 2: bob.Targets.Filesystem:type_name -> bob.BuildInfoFiles
	6, // 3: bob.Targets.Docker:type_name -> bob.Targets.DockerEntry
	7, // 4: bob.BuildInfoFiles.targets:type_name -> bob.BuildInfoFiles.TargetsEntry
	5, // 5: bob.Targets.DockerEntry.value:type_name -> bob.BuildInfoDocker
	4, // 6: bob.BuildInfoFiles.TargetsEntry.value:type_name -> bob.BuildInfoFile
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_buildinfo_proto_init() }
func file_buildinfo_proto_init() {
	if File_buildinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buildinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Targets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfoFiles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfoFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildinfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfoDocker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buildinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buildinfo_proto_goTypes,
		DependencyIndexes: file_buildinfo_proto_depIdxs,
		MessageInfos:      file_buildinfo_proto_msgTypes,
	}.Build()
	File_buildinfo_proto = out.File
	file_buildinfo_proto_rawDesc = nil
	file_buildinfo_proto_goTypes = nil
	file_buildinfo_proto_depIdxs = nil
}
