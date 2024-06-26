// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: engine/engine.proto

package engine

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       map[string]*any1.Any `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkingDir string               `protobuf:"bytes,2,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	EnvVars    map[string]string    `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetMeta() map[string]*any1.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *InitRequest) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *InitRequest) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout     string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ResultCode int32  `protobuf:"varint,3,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *InitResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *InitResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       map[string]*any1.Any `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkingDir string               `protobuf:"bytes,2,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	EnvVars    map[string]string    `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{2}
}

func (x *ShutdownRequest) GetMeta() map[string]*any1.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ShutdownRequest) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *ShutdownRequest) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout     string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ResultCode int32  `protobuf:"varint,3,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{3}
}

func (x *ShutdownResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ShutdownResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *ShutdownResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta              map[string]*any1.Any `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkingDir        string               `protobuf:"bytes,2,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Command           string               `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Args              []string             `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	AllocatePseudoTty bool                 `protobuf:"varint,5,opt,name=allocate_pseudo_tty,json=allocatePseudoTty,proto3" json:"allocate_pseudo_tty,omitempty"`
	EnvVars           map[string]string    `protobuf:"bytes,6,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{4}
}

func (x *RunRequest) GetMeta() map[string]*any1.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RunRequest) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *RunRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *RunRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *RunRequest) GetAllocatePseudoTty() bool {
	if x != nil {
		return x.AllocatePseudoTty
	}
	return false
}

func (x *RunRequest) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout     string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	ResultCode int32  `protobuf:"varint,3,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_engine_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_engine_engine_proto_rawDescGZIP(), []int{5}
}

func (x *RunResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *RunResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *RunResponse) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

var File_engine_engine_proto protoreflect.FileDescriptor

var file_engine_engine_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x0b, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x12, 0x3b, 0x0a, 0x08,
	0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x4d, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56,
	0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xb5, 0x02, 0x0a, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69,
	0x72, 0x12, 0x3f, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76,
	0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x73, 0x1a, 0x4d, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a,
	0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x84, 0x03, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x73, 0x65, 0x75, 0x64, 0x6f, 0x5f, 0x74, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x50, 0x73, 0x65, 0x75, 0x64, 0x6f, 0x54,
	0x74, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x4d,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a,
	0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x0b, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x33, 0x0a,
	0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x30, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x12, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x12, 0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_engine_proto_rawDescOnce sync.Once
	file_engine_engine_proto_rawDescData = file_engine_engine_proto_rawDesc
)

func file_engine_engine_proto_rawDescGZIP() []byte {
	file_engine_engine_proto_rawDescOnce.Do(func() {
		file_engine_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_engine_proto_rawDescData)
	})
	return file_engine_engine_proto_rawDescData
}

var file_engine_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_engine_engine_proto_goTypes = []any{
	(*InitRequest)(nil),      // 0: engine.InitRequest
	(*InitResponse)(nil),     // 1: engine.InitResponse
	(*ShutdownRequest)(nil),  // 2: engine.ShutdownRequest
	(*ShutdownResponse)(nil), // 3: engine.ShutdownResponse
	(*RunRequest)(nil),       // 4: engine.RunRequest
	(*RunResponse)(nil),      // 5: engine.RunResponse
	nil,                      // 6: engine.InitRequest.MetaEntry
	nil,                      // 7: engine.InitRequest.EnvVarsEntry
	nil,                      // 8: engine.ShutdownRequest.MetaEntry
	nil,                      // 9: engine.ShutdownRequest.EnvVarsEntry
	nil,                      // 10: engine.RunRequest.MetaEntry
	nil,                      // 11: engine.RunRequest.EnvVarsEntry
	(*any1.Any)(nil),         // 12: google.protobuf.Any
}
var file_engine_engine_proto_depIdxs = []int32{
	6,  // 0: engine.InitRequest.meta:type_name -> engine.InitRequest.MetaEntry
	7,  // 1: engine.InitRequest.env_vars:type_name -> engine.InitRequest.EnvVarsEntry
	8,  // 2: engine.ShutdownRequest.meta:type_name -> engine.ShutdownRequest.MetaEntry
	9,  // 3: engine.ShutdownRequest.env_vars:type_name -> engine.ShutdownRequest.EnvVarsEntry
	10, // 4: engine.RunRequest.meta:type_name -> engine.RunRequest.MetaEntry
	11, // 5: engine.RunRequest.env_vars:type_name -> engine.RunRequest.EnvVarsEntry
	12, // 6: engine.InitRequest.MetaEntry.value:type_name -> google.protobuf.Any
	12, // 7: engine.ShutdownRequest.MetaEntry.value:type_name -> google.protobuf.Any
	12, // 8: engine.RunRequest.MetaEntry.value:type_name -> google.protobuf.Any
	0,  // 9: engine.CommandExecutor.Init:input_type -> engine.InitRequest
	4,  // 10: engine.CommandExecutor.Run:input_type -> engine.RunRequest
	2,  // 11: engine.CommandExecutor.Shutdown:input_type -> engine.ShutdownRequest
	1,  // 12: engine.CommandExecutor.Init:output_type -> engine.InitResponse
	5,  // 13: engine.CommandExecutor.Run:output_type -> engine.RunResponse
	3,  // 14: engine.CommandExecutor.Shutdown:output_type -> engine.ShutdownResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_engine_engine_proto_init() }
func file_engine_engine_proto_init() {
	if File_engine_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_engine_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitRequest); i {
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
		file_engine_engine_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InitResponse); i {
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
		file_engine_engine_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ShutdownRequest); i {
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
		file_engine_engine_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ShutdownResponse); i {
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
		file_engine_engine_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RunRequest); i {
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
		file_engine_engine_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_engine_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_engine_engine_proto_goTypes,
		DependencyIndexes: file_engine_engine_proto_depIdxs,
		MessageInfos:      file_engine_engine_proto_msgTypes,
	}.Build()
	File_engine_engine_proto = out.File
	file_engine_engine_proto_rawDesc = nil
	file_engine_engine_proto_goTypes = nil
	file_engine_engine_proto_depIdxs = nil
}
