// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_category.proto

/*
Package flextime_v1 is a generated protocol buffer package.

It is generated from these files:
	task_category.proto

It has these top-level messages:
	Task
	TaskCategory
*/
package flextime_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Name     string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Estimate string                     `protobuf:"bytes,2,opt,name=estimate" json:"estimate,omitempty"`
	Repeat   string                     `protobuf:"bytes,3,opt,name=repeat" json:"repeat,omitempty"`
	Due      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=due" json:"due,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetEstimate() string {
	if m != nil {
		return m.Estimate
	}
	return ""
}

func (m *Task) GetRepeat() string {
	if m != nil {
		return m.Repeat
	}
	return ""
}

func (m *Task) GetDue() *google_protobuf.Timestamp {
	if m != nil {
		return m.Due
	}
	return nil
}

type TaskCategory struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Contexts []string `protobuf:"bytes,2,rep,name=contexts" json:"contexts,omitempty"`
	Tasks    []*Task  `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TaskCategory) Reset()                    { *m = TaskCategory{} }
func (m *TaskCategory) String() string            { return proto.CompactTextString(m) }
func (*TaskCategory) ProtoMessage()               {}
func (*TaskCategory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskCategory) GetContexts() []string {
	if m != nil {
		return m.Contexts
	}
	return nil
}

func (m *TaskCategory) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "flextime.v1.Task")
	proto.RegisterType((*TaskCategory)(nil), "flextime.v1.TaskCategory")
}

func init() { proto.RegisterFile("task_category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xa6, 0x2e, 0x3a, 0xf5, 0x62, 0x04, 0x09, 0xbd, 0x18, 0xf6, 0x62, 0x0f, 0x92,
	0xc5, 0xf5, 0x11, 0x7c, 0x83, 0xb2, 0x77, 0xc9, 0xd6, 0x69, 0x28, 0x6d, 0x9a, 0xd2, 0x4c, 0xa5,
	0x82, 0x0f, 0x2f, 0x49, 0x5b, 0xf1, 0xe2, 0x2d, 0x7f, 0xbe, 0x21, 0xdf, 0x3f, 0x81, 0x7b, 0xd2,
	0xbe, 0x7d, 0xaf, 0x34, 0xa1, 0x71, 0xe3, 0x97, 0x1a, 0x46, 0x47, 0x8e, 0x67, 0x75, 0x87, 0x33,
	0x35, 0x16, 0xd5, 0xe7, 0x4b, 0xfe, 0x68, 0x9c, 0x33, 0x1d, 0x1e, 0x23, 0xba, 0x4c, 0xf5, 0x31,
	0x00, 0x4f, 0xda, 0x0e, 0xcb, 0xf4, 0xe1, 0x1b, 0xd2, 0xb3, 0xf6, 0x2d, 0xe7, 0x90, 0xf6, 0xda,
	0xa2, 0x48, 0x64, 0x52, 0xdc, 0x94, 0xf1, 0xcc, 0x73, 0xb8, 0x46, 0x4f, 0x8d, 0xd5, 0x84, 0x62,
	0x17, 0xef, 0x7f, 0x33, 0x7f, 0x80, 0xfd, 0x88, 0x03, 0x6a, 0x12, 0x2c, 0x92, 0x35, 0xf1, 0x67,
	0x60, 0x1f, 0x13, 0x8a, 0x54, 0x26, 0x45, 0x76, 0xca, 0xd5, 0xa2, 0x57, 0x9b, 0x5e, 0x9d, 0x37,
	0x7d, 0x19, 0xc6, 0x0e, 0x06, 0x6e, 0x83, 0xfd, 0x6d, 0xdd, 0xe0, 0xbf, 0x16, 0x95, 0xeb, 0x09,
	0x67, 0xf2, 0x62, 0x27, 0x59, 0x68, 0xb1, 0x65, 0xfe, 0x04, 0x57, 0xe1, 0x0b, 0xbc, 0x60, 0x92,
	0x15, 0xd9, 0xe9, 0x4e, 0xfd, 0xd9, 0x5d, 0x85, 0x97, 0xcb, 0x85, 0x5f, 0xf6, 0xb1, 0xc1, 0xeb,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xa8, 0x95, 0x63, 0x32, 0x01, 0x00, 0x00,
}
