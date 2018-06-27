// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcapi.proto

package iotgrpcapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskStatus int32

const (
	TaskStatus_NOT_SENT  TaskStatus = 0
	TaskStatus_SENT      TaskStatus = 1
	TaskStatus_RECEIVED  TaskStatus = 2
	TaskStatus_COMPLETED TaskStatus = 3
)

var TaskStatus_name = map[int32]string{
	0: "NOT_SENT",
	1: "SENT",
	2: "RECEIVED",
	3: "COMPLETED",
}
var TaskStatus_value = map[string]int32{
	"NOT_SENT":  0,
	"SENT":      1,
	"RECEIVED":  2,
	"COMPLETED": 3,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{0}
}

type TaskDescription struct {
	UserId                string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId                string                 `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskName              string                 `protobuf:"bytes,3,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	HierarchyId           string                 `protobuf:"bytes,4,opt,name=hierarchy_id,json=hierarchyId,proto3" json:"hierarchy_id,omitempty"`
	DueDateTimestamp      int64                  `protobuf:"varint,5,opt,name=due_date_timestamp,json=dueDateTimestamp,proto3" json:"due_date_timestamp,omitempty"`
	IsCompleted           bool                   `protobuf:"varint,6,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"`
	FunctionalLocationIds *FunctionalLocationIds `protobuf:"bytes,7,opt,name=functional_location_ids,json=functionalLocationIds,proto3" json:"functional_location_ids,omitempty"`
	Status                TaskStatus             `protobuf:"varint,8,opt,name=status,proto3,enum=iotgrpcapi.TaskStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *TaskDescription) Reset()         { *m = TaskDescription{} }
func (m *TaskDescription) String() string { return proto.CompactTextString(m) }
func (*TaskDescription) ProtoMessage()    {}
func (*TaskDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{0}
}
func (m *TaskDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDescription.Unmarshal(m, b)
}
func (m *TaskDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDescription.Marshal(b, m, deterministic)
}
func (dst *TaskDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDescription.Merge(dst, src)
}
func (m *TaskDescription) XXX_Size() int {
	return xxx_messageInfo_TaskDescription.Size(m)
}
func (m *TaskDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDescription.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDescription proto.InternalMessageInfo

func (m *TaskDescription) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TaskDescription) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskDescription) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *TaskDescription) GetHierarchyId() string {
	if m != nil {
		return m.HierarchyId
	}
	return ""
}

func (m *TaskDescription) GetDueDateTimestamp() int64 {
	if m != nil {
		return m.DueDateTimestamp
	}
	return 0
}

func (m *TaskDescription) GetIsCompleted() bool {
	if m != nil {
		return m.IsCompleted
	}
	return false
}

func (m *TaskDescription) GetFunctionalLocationIds() *FunctionalLocationIds {
	if m != nil {
		return m.FunctionalLocationIds
	}
	return nil
}

func (m *TaskDescription) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_NOT_SENT
}

type InitialTaskDescription struct {
	UserId                string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskName              string                 `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	HierarchyId           string                 `protobuf:"bytes,3,opt,name=hierarchy_id,json=hierarchyId,proto3" json:"hierarchy_id,omitempty"`
	DueDateTimestamp      int64                  `protobuf:"varint,4,opt,name=due_date_timestamp,json=dueDateTimestamp,proto3" json:"due_date_timestamp,omitempty"`
	FunctionalLocationIds *FunctionalLocationIds `protobuf:"bytes,5,opt,name=functional_location_ids,json=functionalLocationIds,proto3" json:"functional_location_ids,omitempty"`
	ExternalTaskId        string                 `protobuf:"bytes,6,opt,name=external_task_id,json=externalTaskId,proto3" json:"external_task_id,omitempty"`
	Status                TaskStatus             `protobuf:"varint,7,opt,name=status,proto3,enum=iotgrpcapi.TaskStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *InitialTaskDescription) Reset()         { *m = InitialTaskDescription{} }
func (m *InitialTaskDescription) String() string { return proto.CompactTextString(m) }
func (*InitialTaskDescription) ProtoMessage()    {}
func (*InitialTaskDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{1}
}
func (m *InitialTaskDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitialTaskDescription.Unmarshal(m, b)
}
func (m *InitialTaskDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitialTaskDescription.Marshal(b, m, deterministic)
}
func (dst *InitialTaskDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialTaskDescription.Merge(dst, src)
}
func (m *InitialTaskDescription) XXX_Size() int {
	return xxx_messageInfo_InitialTaskDescription.Size(m)
}
func (m *InitialTaskDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialTaskDescription.DiscardUnknown(m)
}

var xxx_messageInfo_InitialTaskDescription proto.InternalMessageInfo

func (m *InitialTaskDescription) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *InitialTaskDescription) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *InitialTaskDescription) GetHierarchyId() string {
	if m != nil {
		return m.HierarchyId
	}
	return ""
}

func (m *InitialTaskDescription) GetDueDateTimestamp() int64 {
	if m != nil {
		return m.DueDateTimestamp
	}
	return 0
}

func (m *InitialTaskDescription) GetFunctionalLocationIds() *FunctionalLocationIds {
	if m != nil {
		return m.FunctionalLocationIds
	}
	return nil
}

func (m *InitialTaskDescription) GetExternalTaskId() string {
	if m != nil {
		return m.ExternalTaskId
	}
	return ""
}

func (m *InitialTaskDescription) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_NOT_SENT
}

type TaskUser struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskUser) Reset()         { *m = TaskUser{} }
func (m *TaskUser) String() string { return proto.CompactTextString(m) }
func (*TaskUser) ProtoMessage()    {}
func (*TaskUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{2}
}
func (m *TaskUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskUser.Unmarshal(m, b)
}
func (m *TaskUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskUser.Marshal(b, m, deterministic)
}
func (dst *TaskUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskUser.Merge(dst, src)
}
func (m *TaskUser) XXX_Size() int {
	return xxx_messageInfo_TaskUser.Size(m)
}
func (m *TaskUser) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskUser.DiscardUnknown(m)
}

var xxx_messageInfo_TaskUser proto.InternalMessageInfo

func (m *TaskUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TaskUser) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type SetTaskStatusInput struct {
	TaskId               string     `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status               TaskStatus `protobuf:"varint,2,opt,name=status,proto3,enum=iotgrpcapi.TaskStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetTaskStatusInput) Reset()         { *m = SetTaskStatusInput{} }
func (m *SetTaskStatusInput) String() string { return proto.CompactTextString(m) }
func (*SetTaskStatusInput) ProtoMessage()    {}
func (*SetTaskStatusInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{3}
}
func (m *SetTaskStatusInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTaskStatusInput.Unmarshal(m, b)
}
func (m *SetTaskStatusInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTaskStatusInput.Marshal(b, m, deterministic)
}
func (dst *SetTaskStatusInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTaskStatusInput.Merge(dst, src)
}
func (m *SetTaskStatusInput) XXX_Size() int {
	return xxx_messageInfo_SetTaskStatusInput.Size(m)
}
func (m *SetTaskStatusInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTaskStatusInput.DiscardUnknown(m)
}

var xxx_messageInfo_SetTaskStatusInput proto.InternalMessageInfo

func (m *SetTaskStatusInput) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SetTaskStatusInput) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_NOT_SENT
}

type TaskDescriptions struct {
	TaskDescriptionArr   []*TaskDescription `protobuf:"bytes,1,rep,name=task_description_arr,json=taskDescriptionArr,proto3" json:"task_description_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TaskDescriptions) Reset()         { *m = TaskDescriptions{} }
func (m *TaskDescriptions) String() string { return proto.CompactTextString(m) }
func (*TaskDescriptions) ProtoMessage()    {}
func (*TaskDescriptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{4}
}
func (m *TaskDescriptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDescriptions.Unmarshal(m, b)
}
func (m *TaskDescriptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDescriptions.Marshal(b, m, deterministic)
}
func (dst *TaskDescriptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDescriptions.Merge(dst, src)
}
func (m *TaskDescriptions) XXX_Size() int {
	return xxx_messageInfo_TaskDescriptions.Size(m)
}
func (m *TaskDescriptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDescriptions.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDescriptions proto.InternalMessageInfo

func (m *TaskDescriptions) GetTaskDescriptionArr() []*TaskDescription {
	if m != nil {
		return m.TaskDescriptionArr
	}
	return nil
}

type FunctionalLocationIds struct {
	IdArr                []string `protobuf:"bytes,1,rep,name=id_arr,json=idArr,proto3" json:"id_arr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FunctionalLocationIds) Reset()         { *m = FunctionalLocationIds{} }
func (m *FunctionalLocationIds) String() string { return proto.CompactTextString(m) }
func (*FunctionalLocationIds) ProtoMessage()    {}
func (*FunctionalLocationIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{5}
}
func (m *FunctionalLocationIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionalLocationIds.Unmarshal(m, b)
}
func (m *FunctionalLocationIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionalLocationIds.Marshal(b, m, deterministic)
}
func (dst *FunctionalLocationIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionalLocationIds.Merge(dst, src)
}
func (m *FunctionalLocationIds) XXX_Size() int {
	return xxx_messageInfo_FunctionalLocationIds.Size(m)
}
func (m *FunctionalLocationIds) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionalLocationIds.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionalLocationIds proto.InternalMessageInfo

func (m *FunctionalLocationIds) GetIdArr() []string {
	if m != nil {
		return m.IdArr
	}
	return nil
}

type PrimitiveString struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveString) Reset()         { *m = PrimitiveString{} }
func (m *PrimitiveString) String() string { return proto.CompactTextString(m) }
func (*PrimitiveString) ProtoMessage()    {}
func (*PrimitiveString) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{6}
}
func (m *PrimitiveString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveString.Unmarshal(m, b)
}
func (m *PrimitiveString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveString.Marshal(b, m, deterministic)
}
func (dst *PrimitiveString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveString.Merge(dst, src)
}
func (m *PrimitiveString) XXX_Size() int {
	return xxx_messageInfo_PrimitiveString.Size(m)
}
func (m *PrimitiveString) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveString.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveString proto.InternalMessageInfo

func (m *PrimitiveString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PrimitiveBool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveBool) Reset()         { *m = PrimitiveBool{} }
func (m *PrimitiveBool) String() string { return proto.CompactTextString(m) }
func (*PrimitiveBool) ProtoMessage()    {}
func (*PrimitiveBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{7}
}
func (m *PrimitiveBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveBool.Unmarshal(m, b)
}
func (m *PrimitiveBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveBool.Marshal(b, m, deterministic)
}
func (dst *PrimitiveBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveBool.Merge(dst, src)
}
func (m *PrimitiveBool) XXX_Size() int {
	return xxx_messageInfo_PrimitiveBool.Size(m)
}
func (m *PrimitiveBool) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveBool.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveBool proto.InternalMessageInfo

func (m *PrimitiveBool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type PrimitiveVoid struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveVoid) Reset()         { *m = PrimitiveVoid{} }
func (m *PrimitiveVoid) String() string { return proto.CompactTextString(m) }
func (*PrimitiveVoid) ProtoMessage()    {}
func (*PrimitiveVoid) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpcapi_1232c0a65f151064, []int{8}
}
func (m *PrimitiveVoid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveVoid.Unmarshal(m, b)
}
func (m *PrimitiveVoid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveVoid.Marshal(b, m, deterministic)
}
func (dst *PrimitiveVoid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveVoid.Merge(dst, src)
}
func (m *PrimitiveVoid) XXX_Size() int {
	return xxx_messageInfo_PrimitiveVoid.Size(m)
}
func (m *PrimitiveVoid) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveVoid.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveVoid proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TaskDescription)(nil), "iotgrpcapi.TaskDescription")
	proto.RegisterType((*InitialTaskDescription)(nil), "iotgrpcapi.InitialTaskDescription")
	proto.RegisterType((*TaskUser)(nil), "iotgrpcapi.TaskUser")
	proto.RegisterType((*SetTaskStatusInput)(nil), "iotgrpcapi.SetTaskStatusInput")
	proto.RegisterType((*TaskDescriptions)(nil), "iotgrpcapi.TaskDescriptions")
	proto.RegisterType((*FunctionalLocationIds)(nil), "iotgrpcapi.FunctionalLocationIds")
	proto.RegisterType((*PrimitiveString)(nil), "iotgrpcapi.PrimitiveString")
	proto.RegisterType((*PrimitiveBool)(nil), "iotgrpcapi.PrimitiveBool")
	proto.RegisterType((*PrimitiveVoid)(nil), "iotgrpcapi.PrimitiveVoid")
	proto.RegisterEnum("iotgrpcapi.TaskStatus", TaskStatus_name, TaskStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IoTClient is the client API for IoT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IoTClient interface {
	DeepPing(ctx context.Context, in *PrimitiveVoid, opts ...grpc.CallOption) (*PrimitiveString, error)
	CreateTask(ctx context.Context, in *InitialTaskDescription, opts ...grpc.CallOption) (*PrimitiveString, error)
	GetAllTasks(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error)
	GetUncompletedTasks(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error)
	SetTaskCompleted(ctx context.Context, in *TaskUser, opts ...grpc.CallOption) (*PrimitiveVoid, error)
	DeleteTask(ctx context.Context, in *TaskUser, opts ...grpc.CallOption) (*PrimitiveVoid, error)
	GetUncompletedTasksByHierarchy(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error)
	SetTaskStatus(ctx context.Context, in *SetTaskStatusInput, opts ...grpc.CallOption) (*PrimitiveVoid, error)
}

type ioTClient struct {
	cc *grpc.ClientConn
}

func NewIoTClient(cc *grpc.ClientConn) IoTClient {
	return &ioTClient{cc}
}

func (c *ioTClient) DeepPing(ctx context.Context, in *PrimitiveVoid, opts ...grpc.CallOption) (*PrimitiveString, error) {
	out := new(PrimitiveString)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/DeepPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) CreateTask(ctx context.Context, in *InitialTaskDescription, opts ...grpc.CallOption) (*PrimitiveString, error) {
	out := new(PrimitiveString)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) GetAllTasks(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error) {
	out := new(TaskDescriptions)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/GetAllTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) GetUncompletedTasks(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error) {
	out := new(TaskDescriptions)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/GetUncompletedTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) SetTaskCompleted(ctx context.Context, in *TaskUser, opts ...grpc.CallOption) (*PrimitiveVoid, error) {
	out := new(PrimitiveVoid)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/SetTaskCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) DeleteTask(ctx context.Context, in *TaskUser, opts ...grpc.CallOption) (*PrimitiveVoid, error) {
	out := new(PrimitiveVoid)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) GetUncompletedTasksByHierarchy(ctx context.Context, in *PrimitiveString, opts ...grpc.CallOption) (*TaskDescriptions, error) {
	out := new(TaskDescriptions)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/GetUncompletedTasksByHierarchy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTClient) SetTaskStatus(ctx context.Context, in *SetTaskStatusInput, opts ...grpc.CallOption) (*PrimitiveVoid, error) {
	out := new(PrimitiveVoid)
	err := c.cc.Invoke(ctx, "/iotgrpcapi.IoT/SetTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IoTServer is the server API for IoT service.
type IoTServer interface {
	DeepPing(context.Context, *PrimitiveVoid) (*PrimitiveString, error)
	CreateTask(context.Context, *InitialTaskDescription) (*PrimitiveString, error)
	GetAllTasks(context.Context, *PrimitiveString) (*TaskDescriptions, error)
	GetUncompletedTasks(context.Context, *PrimitiveString) (*TaskDescriptions, error)
	SetTaskCompleted(context.Context, *TaskUser) (*PrimitiveVoid, error)
	DeleteTask(context.Context, *TaskUser) (*PrimitiveVoid, error)
	GetUncompletedTasksByHierarchy(context.Context, *PrimitiveString) (*TaskDescriptions, error)
	SetTaskStatus(context.Context, *SetTaskStatusInput) (*PrimitiveVoid, error)
}

func RegisterIoTServer(s *grpc.Server, srv IoTServer) {
	s.RegisterService(&_IoT_serviceDesc, srv)
}

func _IoT_DeepPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimitiveVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).DeepPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/DeepPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).DeepPing(ctx, req.(*PrimitiveVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitialTaskDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).CreateTask(ctx, req.(*InitialTaskDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_GetAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimitiveString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).GetAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/GetAllTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).GetAllTasks(ctx, req.(*PrimitiveString))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_GetUncompletedTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimitiveString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).GetUncompletedTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/GetUncompletedTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).GetUncompletedTasks(ctx, req.(*PrimitiveString))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_SetTaskCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).SetTaskCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/SetTaskCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).SetTaskCompleted(ctx, req.(*TaskUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).DeleteTask(ctx, req.(*TaskUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_GetUncompletedTasksByHierarchy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimitiveString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).GetUncompletedTasksByHierarchy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/GetUncompletedTasksByHierarchy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).GetUncompletedTasksByHierarchy(ctx, req.(*PrimitiveString))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoT_SetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaskStatusInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServer).SetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotgrpcapi.IoT/SetTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServer).SetTaskStatus(ctx, req.(*SetTaskStatusInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _IoT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iotgrpcapi.IoT",
	HandlerType: (*IoTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeepPing",
			Handler:    _IoT_DeepPing_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _IoT_CreateTask_Handler,
		},
		{
			MethodName: "GetAllTasks",
			Handler:    _IoT_GetAllTasks_Handler,
		},
		{
			MethodName: "GetUncompletedTasks",
			Handler:    _IoT_GetUncompletedTasks_Handler,
		},
		{
			MethodName: "SetTaskCompleted",
			Handler:    _IoT_SetTaskCompleted_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _IoT_DeleteTask_Handler,
		},
		{
			MethodName: "GetUncompletedTasksByHierarchy",
			Handler:    _IoT_GetUncompletedTasksByHierarchy_Handler,
		},
		{
			MethodName: "SetTaskStatus",
			Handler:    _IoT_SetTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi.proto",
}

func init() { proto.RegisterFile("grpcapi.proto", fileDescriptor_grpcapi_1232c0a65f151064) }

var fileDescriptor_grpcapi_1232c0a65f151064 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5b, 0x4f, 0xdb, 0x4a,
	0x10, 0x8e, 0x73, 0xc3, 0x99, 0x10, 0x88, 0xf6, 0x70, 0xc9, 0x81, 0x23, 0x14, 0x2c, 0x1d, 0xd5,
	0xaa, 0xaa, 0x3c, 0xa4, 0xaf, 0x95, 0x2a, 0x88, 0x5d, 0xea, 0x8a, 0x9b, 0x1c, 0x83, 0xd4, 0x07,
	0x64, 0x6d, 0xb3, 0x0b, 0xac, 0x70, 0x6c, 0x6b, 0x77, 0x8d, 0xca, 0xff, 0xec, 0x7b, 0x7f, 0x43,
	0xff, 0x41, 0xb5, 0xce, 0xc5, 0x0e, 0x0d, 0x21, 0x2d, 0x7d, 0xf3, 0x7e, 0xf3, 0xcd, 0x97, 0x99,
	0x6f, 0x67, 0x36, 0xd0, 0xb8, 0xe1, 0xf1, 0x00, 0xc7, 0xac, 0x13, 0xf3, 0x48, 0x46, 0x08, 0x58,
	0x24, 0xc7, 0x88, 0xf1, 0xbd, 0x08, 0xeb, 0x1e, 0x16, 0x77, 0x16, 0x15, 0x03, 0xce, 0x62, 0xc9,
	0xa2, 0x10, 0x6d, 0xc3, 0x4a, 0x22, 0x28, 0xf7, 0x19, 0x69, 0x69, 0x6d, 0xcd, 0xac, 0xb9, 0x55,
	0x75, 0x74, 0x88, 0x0a, 0x48, 0x2c, 0xee, 0x54, 0xa0, 0x38, 0x0a, 0xa8, 0xa3, 0x43, 0xd0, 0x2e,
	0xd4, 0xd2, 0x40, 0x88, 0x87, 0xb4, 0x55, 0x4a, 0x43, 0xba, 0x02, 0x4e, 0xf1, 0x90, 0xa2, 0x7d,
	0x58, 0xbd, 0x65, 0x94, 0x63, 0x3e, 0xb8, 0x7d, 0x50, 0xa9, 0xe5, 0x34, 0x5e, 0x9f, 0x62, 0x0e,
	0x41, 0x6f, 0x00, 0x91, 0x84, 0xfa, 0x04, 0x4b, 0xea, 0x4b, 0x36, 0xa4, 0x42, 0xe2, 0x61, 0xdc,
	0xaa, 0xb4, 0x35, 0xb3, 0xe4, 0x36, 0x49, 0x42, 0x2d, 0x2c, 0xa9, 0x37, 0xc1, 0x95, 0x20, 0x13,
	0xfe, 0x20, 0x1a, 0xc6, 0x01, 0x95, 0x94, 0xb4, 0xaa, 0x6d, 0xcd, 0xd4, 0xdd, 0x3a, 0x13, 0xbd,
	0x09, 0x84, 0x3e, 0xc3, 0xf6, 0x75, 0x12, 0x0e, 0x54, 0x3b, 0x38, 0xf0, 0x83, 0x68, 0x80, 0xd5,
	0xa7, 0xcf, 0x88, 0x68, 0xad, 0xb4, 0x35, 0xb3, 0xde, 0xdd, 0xef, 0x64, 0x26, 0x74, 0x3e, 0x4c,
	0xa9, 0xc7, 0x63, 0xa6, 0x43, 0x84, 0xbb, 0x79, 0x3d, 0x0f, 0x46, 0x1d, 0xa8, 0x0a, 0x89, 0x65,
	0x22, 0x5a, 0x7a, 0x5b, 0x33, 0xd7, 0xba, 0x5b, 0x79, 0x25, 0x65, 0x65, 0x3f, 0x8d, 0xba, 0x63,
	0x96, 0xf1, 0xad, 0x08, 0x5b, 0x4e, 0xc8, 0x24, 0xc3, 0xc1, 0xd2, 0x46, 0xcf, 0xf8, 0x59, 0x7c,
	0xc6, 0xcf, 0xd2, 0xb2, 0x7e, 0x96, 0x9f, 0xf0, 0x73, 0x81, 0x59, 0x95, 0x17, 0x9a, 0x65, 0x42,
	0x93, 0x7e, 0x95, 0x94, 0x2b, 0xe1, 0xc9, 0xe8, 0x54, 0xd3, 0x7a, 0xd7, 0x26, 0xb8, 0x37, 0x1a,
	0xa1, 0xcc, 0xd6, 0x95, 0xa5, 0x6c, 0x7d, 0x07, 0xba, 0x42, 0x2f, 0x04, 0xe5, 0xbf, 0x3f, 0xb0,
	0xc6, 0x15, 0xa0, 0x3e, 0x95, 0x99, 0xac, 0x13, 0xc6, 0x89, 0xcc, 0xd3, 0xb5, 0x99, 0xf9, 0xce,
	0x8a, 0x2b, 0x2e, 0x55, 0x1c, 0x86, 0xe6, 0xa3, 0xbb, 0x16, 0xe8, 0x04, 0x36, 0x52, 0x71, 0x92,
	0x81, 0x3e, 0xe6, 0xbc, 0xa5, 0xb5, 0x4b, 0x66, 0xbd, 0xbb, 0xfb, 0x58, 0x31, 0x97, 0xeb, 0x22,
	0x39, 0x0b, 0x1c, 0x70, 0x6e, 0x74, 0x60, 0x73, 0xee, 0x4d, 0xa0, 0x4d, 0xa8, 0x32, 0x32, 0x55,
	0xae, 0xb9, 0x15, 0x46, 0x14, 0xff, 0x15, 0xac, 0x9f, 0x73, 0x36, 0x64, 0x92, 0xdd, 0xd3, 0xbe,
	0xe4, 0x2c, 0xbc, 0x41, 0x1b, 0x50, 0xb9, 0xc7, 0x41, 0x42, 0xc7, 0xcd, 0x8e, 0x0e, 0xc6, 0xff,
	0xd0, 0x98, 0x12, 0x0f, 0xa3, 0x28, 0x98, 0xa5, 0xe9, 0x13, 0xda, 0x7a, 0x8e, 0x76, 0x19, 0x31,
	0xf2, 0xfa, 0x00, 0x20, 0x73, 0x02, 0xad, 0x82, 0x7e, 0x7a, 0xe6, 0xf9, 0x7d, 0xfb, 0xd4, 0x6b,
	0x16, 0x90, 0x0e, 0xe5, 0xf4, 0x4b, 0x53, 0xb8, 0x6b, 0xf7, 0x6c, 0xe7, 0xd2, 0xb6, 0x9a, 0x45,
	0xd4, 0x80, 0x5a, 0xef, 0xec, 0xe4, 0xfc, 0xd8, 0xf6, 0x6c, 0xab, 0x59, 0xea, 0xfe, 0x28, 0x43,
	0xc9, 0x89, 0x3c, 0x64, 0x81, 0x6e, 0x51, 0x1a, 0x9f, 0xab, 0x22, 0xff, 0xcd, 0x1b, 0x33, 0xf3,
	0x8b, 0x3b, 0xbb, 0x73, 0x43, 0xa3, 0xe6, 0x8c, 0x02, 0x3a, 0x03, 0xe8, 0x71, 0xaa, 0x26, 0x1d,
	0x8b, 0x3b, 0x64, 0xe4, 0xc9, 0xf3, 0xf7, 0xf1, 0x39, 0xc1, 0x4f, 0x50, 0x3f, 0xa2, 0xf2, 0x20,
	0x48, 0xf3, 0x04, 0x5a, 0xc4, 0xde, 0xf9, 0x6f, 0xc1, 0x7d, 0x0a, 0xa3, 0x80, 0x5c, 0xf8, 0xe7,
	0x88, 0xca, 0x8b, 0x70, 0xfa, 0x8a, 0xfd, 0x05, 0x4d, 0x1b, 0x9a, 0xe3, 0xa1, 0xce, 0x1e, 0xc2,
	0x8d, 0xc7, 0x39, 0x6a, 0x61, 0x76, 0x9e, 0x36, 0xd5, 0x28, 0xa0, 0xf7, 0x00, 0x16, 0x55, 0xd9,
	0xa9, 0x6f, 0x7f, 0x20, 0x70, 0x05, 0x7b, 0x73, 0x7a, 0x3b, 0x7c, 0xf8, 0x38, 0x79, 0xa1, 0x5e,
	0xd6, 0xe6, 0x31, 0x34, 0x66, 0x76, 0x17, 0xed, 0xe5, 0x13, 0x7e, 0x5d, 0xeb, 0x85, 0xc5, 0x7e,
	0xa9, 0xa6, 0xff, 0x89, 0x6f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x25, 0x5e, 0xb1, 0xfd, 0x24,
	0x07, 0x00, 0x00,
}
