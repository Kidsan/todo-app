// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/todo-app.proto

package proto

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

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tasks       []*TaskRequest `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{0}
}

func (x *TodoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoRequest) GetTasks() []*TaskRequest {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tasks       []*TaskReply `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TodoReply) Reset() {
	*x = TodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoReply) ProtoMessage() {}

func (x *TodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoReply.ProtoReflect.Descriptor instead.
func (*TodoReply) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{1}
}

func (x *TodoReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoReply) GetTasks() []*TaskReply {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TodoUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tasks       []*TaskRequest `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TodoUpdate) Reset() {
	*x = TodoUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUpdate) ProtoMessage() {}

func (x *TodoUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUpdate.ProtoReflect.Descriptor instead.
func (*TodoUpdate) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{2}
}

func (x *TodoUpdate) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoUpdate) GetTasks() []*TaskRequest {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TodoIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TodoIdentifier) Reset() {
	*x = TodoIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoIdentifier) ProtoMessage() {}

func (x *TodoIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoIdentifier.ProtoReflect.Descriptor instead.
func (*TodoIdentifier) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{3}
}

func (x *TodoIdentifier) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TodoListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*TodoReply `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *TodoListReply) Reset() {
	*x = TodoListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListReply) ProtoMessage() {}

func (x *TodoListReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListReply.ProtoReflect.Descriptor instead.
func (*TodoListReply) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{4}
}

func (x *TodoListReply) GetTodos() []*TodoReply {
	if x != nil {
		return x.Todos
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{5}
}

type TaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TaskReply) Reset() {
	*x = TaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReply) ProtoMessage() {}

func (x *TaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReply.ProtoReflect.Descriptor instead.
func (*TaskReply) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{6}
}

func (x *TaskReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todo_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_app_proto_rawDescGZIP(), []int{7}
}

func (x *TaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_todo_app_proto protoreflect.FileDescriptor

var file_proto_todo_app_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a,
	0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x79, 0x0a, 0x09,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x7c, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f,
	0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x21, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x30, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x69, 0x64, 0x73, 0x61, 0x6e, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x61, 0x70, 0x70,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todo_app_proto_rawDescOnce sync.Once
	file_proto_todo_app_proto_rawDescData = file_proto_todo_app_proto_rawDesc
)

func file_proto_todo_app_proto_rawDescGZIP() []byte {
	file_proto_todo_app_proto_rawDescOnce.Do(func() {
		file_proto_todo_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todo_app_proto_rawDescData)
	})
	return file_proto_todo_app_proto_rawDescData
}

var file_proto_todo_app_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_todo_app_proto_goTypes = []interface{}{
	(*TodoRequest)(nil),    // 0: proto.TodoRequest
	(*TodoReply)(nil),      // 1: proto.TodoReply
	(*TodoUpdate)(nil),     // 2: proto.TodoUpdate
	(*TodoIdentifier)(nil), // 3: proto.TodoIdentifier
	(*TodoListReply)(nil),  // 4: proto.TodoListReply
	(*GetRequest)(nil),     // 5: proto.GetRequest
	(*TaskReply)(nil),      // 6: proto.TaskReply
	(*TaskRequest)(nil),    // 7: proto.TaskRequest
}
var file_proto_todo_app_proto_depIdxs = []int32{
	7, // 0: proto.TodoRequest.tasks:type_name -> proto.TaskRequest
	6, // 1: proto.TodoReply.tasks:type_name -> proto.TaskReply
	7, // 2: proto.TodoUpdate.tasks:type_name -> proto.TaskRequest
	1, // 3: proto.TodoListReply.todos:type_name -> proto.TodoReply
	5, // 4: proto.Todos.Get:input_type -> proto.GetRequest
	0, // 5: proto.Todos.Save:input_type -> proto.TodoRequest
	3, // 6: proto.Todos.Find:input_type -> proto.TodoIdentifier
	2, // 7: proto.Todos.Update:input_type -> proto.TodoUpdate
	4, // 8: proto.Todos.Get:output_type -> proto.TodoListReply
	1, // 9: proto.Todos.Save:output_type -> proto.TodoReply
	1, // 10: proto.Todos.Find:output_type -> proto.TodoReply
	1, // 11: proto.Todos.Update:output_type -> proto.TodoReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_todo_app_proto_init() }
func file_proto_todo_app_proto_init() {
	if File_proto_todo_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todo_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoRequest); i {
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
		file_proto_todo_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoReply); i {
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
		file_proto_todo_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoUpdate); i {
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
		file_proto_todo_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoIdentifier); i {
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
		file_proto_todo_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListReply); i {
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
		file_proto_todo_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_todo_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReply); i {
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
		file_proto_todo_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
			RawDescriptor: file_proto_todo_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_app_proto_goTypes,
		DependencyIndexes: file_proto_todo_app_proto_depIdxs,
		MessageInfos:      file_proto_todo_app_proto_msgTypes,
	}.Build()
	File_proto_todo_app_proto = out.File
	file_proto_todo_app_proto_rawDesc = nil
	file_proto_todo_app_proto_goTypes = nil
	file_proto_todo_app_proto_depIdxs = nil
}
