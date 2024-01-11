// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: dagger.proto

package daggerProto

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

type DbContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General search DB service
	Res []string `protobuf:"bytes,1,rep,name=res,proto3" json:"res,omitempty"`
}

func (x *DbContents) Reset() {
	*x = DbContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbContents) ProtoMessage() {}

func (x *DbContents) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbContents.ProtoReflect.Descriptor instead.
func (*DbContents) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{0}
}

func (x *DbContents) GetRes() []string {
	if x != nil {
		return x.Res
	}
	return nil
}

type UpdateObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID        string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Whoami      string `protobuf:"bytes,2,opt,name=Whoami,proto3" json:"Whoami,omitempty"`
	Signature   string `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Retrieved   int32  `protobuf:"varint,4,opt,name=Retrieved,proto3" json:"Retrieved,omitempty"`
	Command     string `protobuf:"bytes,5,opt,name=Command,proto3" json:"Command,omitempty"`
	LastCheckIn string `protobuf:"bytes,6,opt,name=LastCheckIn,proto3" json:"LastCheckIn,omitempty"`
	Result      string `protobuf:"bytes,7,opt,name=Result,proto3" json:"Result,omitempty"`
	GotIt       int32  `protobuf:"varint,8,opt,name=GotIt,proto3" json:"GotIt,omitempty"`
}

func (x *UpdateObject) Reset() {
	*x = UpdateObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObject) ProtoMessage() {}

func (x *UpdateObject) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObject.ProtoReflect.Descriptor instead.
func (*UpdateObject) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateObject) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *UpdateObject) GetWhoami() string {
	if x != nil {
		return x.Whoami
	}
	return ""
}

func (x *UpdateObject) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UpdateObject) GetRetrieved() int32 {
	if x != nil {
		return x.Retrieved
	}
	return 0
}

func (x *UpdateObject) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *UpdateObject) GetLastCheckIn() string {
	if x != nil {
		return x.LastCheckIn
	}
	return ""
}

func (x *UpdateObject) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *UpdateObject) GetGotIt() int32 {
	if x != nil {
		return x.GotIt
	}
	return 0
}

type BuildRoutine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform        string `protobuf:"bytes,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Architecture    string `protobuf:"bytes,2,opt,name=Architecture,proto3" json:"Architecture,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ListenerAddress string `protobuf:"bytes,4,opt,name=Listener_address,json=ListenerAddress,proto3" json:"Listener_address,omitempty"`
}

func (x *BuildRoutine) Reset() {
	*x = BuildRoutine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRoutine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRoutine) ProtoMessage() {}

func (x *BuildRoutine) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRoutine.ProtoReflect.Descriptor instead.
func (*BuildRoutine) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{2}
}

func (x *BuildRoutine) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *BuildRoutine) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *BuildRoutine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildRoutine) GetListenerAddress() string {
	if x != nil {
		return x.ListenerAddress
	}
	return ""
}

type ReponseCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *ReponseCode) Reset() {
	*x = ReponseCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReponseCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReponseCode) ProtoMessage() {}

func (x *ReponseCode) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReponseCode.ProtoReflect.Descriptor instead.
func (*ReponseCode) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{3}
}

func (x *ReponseCode) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *GetUUID) Reset() {
	*x = GetUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUUID) ProtoMessage() {}

func (x *GetUUID) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUUID.ProtoReflect.Descriptor instead.
func (*GetUUID) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{4}
}

func (x *GetUUID) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type GetKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *GetKey) Reset() {
	*x = GetKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dagger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKey) ProtoMessage() {}

func (x *GetKey) ProtoReflect() protoreflect.Message {
	mi := &file_dagger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKey.ProtoReflect.Descriptor instead.
func (*GetKey) Descriptor() ([]byte, []int) {
	return file_dagger_proto_rawDescGZIP(), []int{5}
}

func (x *GetKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_dagger_proto protoreflect.FileDescriptor

var file_dagger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0a, 0x44, 0x62, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x57,
	0x68, 0x6f, 0x61, 0x6d, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x68, 0x6f,
	0x61, 0x6d, 0x69, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x6f, 0x74, 0x49, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x47, 0x6f, 0x74, 0x49, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x07,
	0x67, 0x65, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x1a, 0x0a, 0x06, 0x67,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x32, 0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x2e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x64, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x64, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x00, 0x32, 0x3d, 0x0a, 0x0a, 0x68, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x2f, 0x0a, 0x04, 0x48, 0x67, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x67, 0x65, 0x74, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00,
	0x32, 0x3b, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a,
	0x03, 0x67, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x65,
	0x74, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x32, 0x47, 0x0a,
	0x07, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x1a,
	0x13, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x32, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x00, 0x42, 0x1a, 0x5a, 0x18, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x61, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dagger_proto_rawDescOnce sync.Once
	file_dagger_proto_rawDescData = file_dagger_proto_rawDesc
)

func file_dagger_proto_rawDescGZIP() []byte {
	file_dagger_proto_rawDescOnce.Do(func() {
		file_dagger_proto_rawDescData = protoimpl.X.CompressGZIP(file_dagger_proto_rawDescData)
	})
	return file_dagger_proto_rawDescData
}

var file_dagger_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dagger_proto_goTypes = []interface{}{
	(*DbContents)(nil),   // 0: dagger.DbContents
	(*UpdateObject)(nil), // 1: dagger.UpdateObject
	(*BuildRoutine)(nil), // 2: dagger.BuildRoutine
	(*ReponseCode)(nil),  // 3: dagger.ReponseCode
	(*GetUUID)(nil),      // 4: dagger.getUUID
	(*GetKey)(nil),       // 5: dagger.getKey
}
var file_dagger_proto_depIdxs = []int32{
	5, // 0: dagger.GetAll.GetAll:input_type -> dagger.getKey
	4, // 1: dagger.hgetRecord.Hget:input_type -> dagger.getUUID
	4, // 2: dagger.getRecord.get:input_type -> dagger.getUUID
	2, // 3: dagger.Builder.StartBuilding:input_type -> dagger.BuildRoutine
	1, // 4: dagger.UpdateRecord.SendUpdate:input_type -> dagger.UpdateObject
	0, // 5: dagger.GetAll.GetAll:output_type -> dagger.DbContents
	1, // 6: dagger.hgetRecord.Hget:output_type -> dagger.UpdateObject
	1, // 7: dagger.getRecord.get:output_type -> dagger.UpdateObject
	3, // 8: dagger.Builder.StartBuilding:output_type -> dagger.ReponseCode
	3, // 9: dagger.UpdateRecord.SendUpdate:output_type -> dagger.ReponseCode
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dagger_proto_init() }
func file_dagger_proto_init() {
	if File_dagger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dagger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbContents); i {
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
		file_dagger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObject); i {
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
		file_dagger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRoutine); i {
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
		file_dagger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReponseCode); i {
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
		file_dagger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUUID); i {
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
		file_dagger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKey); i {
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
			RawDescriptor: file_dagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_dagger_proto_goTypes,
		DependencyIndexes: file_dagger_proto_depIdxs,
		MessageInfos:      file_dagger_proto_msgTypes,
	}.Build()
	File_dagger_proto = out.File
	file_dagger_proto_rawDesc = nil
	file_dagger_proto_goTypes = nil
	file_dagger_proto_depIdxs = nil
}
