// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: delta.proto

package logstream

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

type Delta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        int32            `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	DeltaManifests []*DeltaManifest `protobuf:"bytes,2,rep,name=delta_manifests,json=deltaManifests,proto3" json:"delta_manifests,omitempty"`
	DeltaLogs      []*DeltaLog      `protobuf:"bytes,3,rep,name=delta_logs,json=deltaLogs,proto3" json:"delta_logs,omitempty"`
}

func (x *Delta) Reset() {
	*x = Delta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delta) ProtoMessage() {}

func (x *Delta) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delta.ProtoReflect.Descriptor instead.
func (*Delta) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{0}
}

func (x *Delta) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Delta) GetDeltaManifests() []*DeltaManifest {
	if x != nil {
		return x.DeltaManifests
	}
	return nil
}

func (x *Delta) GetDeltaLogs() []*DeltaLog {
	if x != nil {
		return x.DeltaLogs
	}
	return nil
}

type DeltaManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Types that are assignable to DeltaManifestOneof:
	//	*DeltaManifest_Eof
	//	*DeltaManifest_ResetAll
	//	*DeltaManifest_Fields
	DeltaManifestOneof isDeltaManifest_DeltaManifestOneof `protobuf_oneof:"delta_manifest_oneof"`
}

func (x *DeltaManifest) Reset() {
	*x = DeltaManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaManifest) ProtoMessage() {}

func (x *DeltaManifest) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaManifest.ProtoReflect.Descriptor instead.
func (*DeltaManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{1}
}

func (x *DeltaManifest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (m *DeltaManifest) GetDeltaManifestOneof() isDeltaManifest_DeltaManifestOneof {
	if m != nil {
		return m.DeltaManifestOneof
	}
	return nil
}

func (x *DeltaManifest) GetEof() bool {
	if x, ok := x.GetDeltaManifestOneof().(*DeltaManifest_Eof); ok {
		return x.Eof
	}
	return false
}

func (x *DeltaManifest) GetResetAll() *BuildManifest {
	if x, ok := x.GetDeltaManifestOneof().(*DeltaManifest_ResetAll); ok {
		return x.ResetAll
	}
	return nil
}

func (x *DeltaManifest) GetFields() *DeltaManifest_FieldsDelta {
	if x, ok := x.GetDeltaManifestOneof().(*DeltaManifest_Fields); ok {
		return x.Fields
	}
	return nil
}

type isDeltaManifest_DeltaManifestOneof interface {
	isDeltaManifest_DeltaManifestOneof()
}

type DeltaManifest_Eof struct {
	Eof bool `protobuf:"varint,2,opt,name=eof,proto3,oneof"`
}

type DeltaManifest_ResetAll struct {
	ResetAll *BuildManifest `protobuf:"bytes,3,opt,name=reset_all,json=resetAll,proto3,oneof"`
}

type DeltaManifest_Fields struct {
	Fields *DeltaManifest_FieldsDelta `protobuf:"bytes,4,opt,name=fields,proto3,oneof"`
}

func (*DeltaManifest_Eof) isDeltaManifest_DeltaManifestOneof() {}

func (*DeltaManifest_ResetAll) isDeltaManifest_DeltaManifestOneof() {}

func (*DeltaManifest_Fields) isDeltaManifest_DeltaManifestOneof() {}

type DeltaTargetManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OverrideArgs []string                        `protobuf:"bytes,2,rep,name=override_args,json=overrideArgs,proto3" json:"override_args,omitempty"`
	Platform     string                          `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Status       BuildStatus                     `protobuf:"varint,4,opt,name=status,proto3,enum=api.public.logstream.BuildStatus" json:"status,omitempty"`
	StartedAt    int64                           `protobuf:"varint,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt   int64                           `protobuf:"varint,6,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Commands     map[int32]*DeltaCommandManifest `protobuf:"bytes,7,rep,name=commands,proto3" json:"commands,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeltaTargetManifest) Reset() {
	*x = DeltaTargetManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaTargetManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaTargetManifest) ProtoMessage() {}

func (x *DeltaTargetManifest) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaTargetManifest.ProtoReflect.Descriptor instead.
func (*DeltaTargetManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaTargetManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeltaTargetManifest) GetOverrideArgs() []string {
	if x != nil {
		return x.OverrideArgs
	}
	return nil
}

func (x *DeltaTargetManifest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *DeltaTargetManifest) GetStatus() BuildStatus {
	if x != nil {
		return x.Status
	}
	return BuildStatus_BUILD_STATUS_UNKNOWN
}

func (x *DeltaTargetManifest) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *DeltaTargetManifest) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *DeltaTargetManifest) GetCommands() map[int32]*DeltaCommandManifest {
	if x != nil {
		return x.Commands
	}
	return nil
}

type DeltaCommandManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status         BuildStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.public.logstream.BuildStatus" json:"status,omitempty"`
	HasCached      bool        `protobuf:"varint,3,opt,name=has_cached,json=hasCached,proto3" json:"has_cached,omitempty"`
	IsCached       bool        `protobuf:"varint,4,opt,name=is_cached,json=isCached,proto3" json:"is_cached,omitempty"`
	HasPush        bool        `protobuf:"varint,5,opt,name=has_push,json=hasPush,proto3" json:"has_push,omitempty"`
	IsPush         bool        `protobuf:"varint,6,opt,name=is_push,json=isPush,proto3" json:"is_push,omitempty"`
	HasLocal       bool        `protobuf:"varint,7,opt,name=has_local,json=hasLocal,proto3" json:"has_local,omitempty"`
	IsLocal        bool        `protobuf:"varint,8,opt,name=is_local,json=isLocal,proto3" json:"is_local,omitempty"`
	StartedAt      int64       `protobuf:"varint,9,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt     int64       `protobuf:"varint,10,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	HasHasProgress bool        `protobuf:"varint,11,opt,name=has_has_progress,json=hasHasProgress,proto3" json:"has_has_progress,omitempty"`
	HasProgress    bool        `protobuf:"varint,12,opt,name=has_progress,json=hasProgress,proto3" json:"has_progress,omitempty"`
	Progress       int32       `protobuf:"varint,13,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *DeltaCommandManifest) Reset() {
	*x = DeltaCommandManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaCommandManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaCommandManifest) ProtoMessage() {}

func (x *DeltaCommandManifest) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaCommandManifest.ProtoReflect.Descriptor instead.
func (*DeltaCommandManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{3}
}

func (x *DeltaCommandManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeltaCommandManifest) GetStatus() BuildStatus {
	if x != nil {
		return x.Status
	}
	return BuildStatus_BUILD_STATUS_UNKNOWN
}

func (x *DeltaCommandManifest) GetHasCached() bool {
	if x != nil {
		return x.HasCached
	}
	return false
}

func (x *DeltaCommandManifest) GetIsCached() bool {
	if x != nil {
		return x.IsCached
	}
	return false
}

func (x *DeltaCommandManifest) GetHasPush() bool {
	if x != nil {
		return x.HasPush
	}
	return false
}

func (x *DeltaCommandManifest) GetIsPush() bool {
	if x != nil {
		return x.IsPush
	}
	return false
}

func (x *DeltaCommandManifest) GetHasLocal() bool {
	if x != nil {
		return x.HasLocal
	}
	return false
}

func (x *DeltaCommandManifest) GetIsLocal() bool {
	if x != nil {
		return x.IsLocal
	}
	return false
}

func (x *DeltaCommandManifest) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *DeltaCommandManifest) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *DeltaCommandManifest) GetHasHasProgress() bool {
	if x != nil {
		return x.HasHasProgress
	}
	return false
}

func (x *DeltaCommandManifest) GetHasProgress() bool {
	if x != nil {
		return x.HasProgress
	}
	return false
}

func (x *DeltaCommandManifest) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type DeltaLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeekIndex int64  `protobuf:"varint,1,opt,name=seek_index,json=seekIndex,proto3" json:"seek_index,omitempty"`
	TargetId  string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// Types that are assignable to DeltaLogOneof:
	//	*DeltaLog_Eof
	//	*DeltaLog_Data
	DeltaLogOneof isDeltaLog_DeltaLogOneof `protobuf_oneof:"delta_log_oneof"`
}

func (x *DeltaLog) Reset() {
	*x = DeltaLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaLog) ProtoMessage() {}

func (x *DeltaLog) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaLog.ProtoReflect.Descriptor instead.
func (*DeltaLog) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{4}
}

func (x *DeltaLog) GetSeekIndex() int64 {
	if x != nil {
		return x.SeekIndex
	}
	return 0
}

func (x *DeltaLog) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (m *DeltaLog) GetDeltaLogOneof() isDeltaLog_DeltaLogOneof {
	if m != nil {
		return m.DeltaLogOneof
	}
	return nil
}

func (x *DeltaLog) GetEof() bool {
	if x, ok := x.GetDeltaLogOneof().(*DeltaLog_Eof); ok {
		return x.Eof
	}
	return false
}

func (x *DeltaLog) GetData() []byte {
	if x, ok := x.GetDeltaLogOneof().(*DeltaLog_Data); ok {
		return x.Data
	}
	return nil
}

type isDeltaLog_DeltaLogOneof interface {
	isDeltaLog_DeltaLogOneof()
}

type DeltaLog_Eof struct {
	Eof bool `protobuf:"varint,3,opt,name=eof,proto3,oneof"`
}

type DeltaLog_Data struct {
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

func (*DeltaLog_Eof) isDeltaLog_DeltaLogOneof() {}

func (*DeltaLog_Data) isDeltaLog_DeltaLogOneof() {}

type DeltaManifest_FieldsDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartedAt     int64                           `protobuf:"varint,1,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt    int64                           `protobuf:"varint,2,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Status        BuildStatus                     `protobuf:"varint,3,opt,name=status,proto3,enum=api.public.logstream.BuildStatus" json:"status,omitempty"`
	FailedTarget  string                          `protobuf:"bytes,4,opt,name=failed_target,json=failedTarget,proto3" json:"failed_target,omitempty"`
	FailedSummary string                          `protobuf:"bytes,5,opt,name=failed_summary,json=failedSummary,proto3" json:"failed_summary,omitempty"`
	Targets       map[string]*DeltaTargetManifest `protobuf:"bytes,6,rep,name=targets,proto3" json:"targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeltaManifest_FieldsDelta) Reset() {
	*x = DeltaManifest_FieldsDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaManifest_FieldsDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaManifest_FieldsDelta) ProtoMessage() {}

func (x *DeltaManifest_FieldsDelta) ProtoReflect() protoreflect.Message {
	mi := &file_delta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaManifest_FieldsDelta.ProtoReflect.Descriptor instead.
func (*DeltaManifest_FieldsDelta) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DeltaManifest_FieldsDelta) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *DeltaManifest_FieldsDelta) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *DeltaManifest_FieldsDelta) GetStatus() BuildStatus {
	if x != nil {
		return x.Status
	}
	return BuildStatus_BUILD_STATUS_UNKNOWN
}

func (x *DeltaManifest_FieldsDelta) GetFailedTarget() string {
	if x != nil {
		return x.FailedTarget
	}
	return ""
}

func (x *DeltaManifest_FieldsDelta) GetFailedSummary() string {
	if x != nil {
		return x.FailedSummary
	}
	return ""
}

func (x *DeltaManifest_FieldsDelta) GetTargets() map[string]*DeltaTargetManifest {
	if x != nil {
		return x.Targets
	}
	return nil
}

var File_delta_proto protoreflect.FileDescriptor

var file_delta_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x1a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x4c, 0x6f, 0x67, 0x73, 0x22, 0xfb, 0x04, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x42, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x49, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x1a, 0x93, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x07, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x1a, 0x65, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x22, 0xa3, 0x03, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x53, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a,
	0x67, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x03, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x68, 0x61, 0x73, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x68, 0x61, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x75,
	0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x75, 0x73, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f,
	0x68, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x48, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x65, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x03, 0x65, 0x6f,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x14,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x6c, 0x6f,
	0x67, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delta_proto_rawDescOnce sync.Once
	file_delta_proto_rawDescData = file_delta_proto_rawDesc
)

func file_delta_proto_rawDescGZIP() []byte {
	file_delta_proto_rawDescOnce.Do(func() {
		file_delta_proto_rawDescData = protoimpl.X.CompressGZIP(file_delta_proto_rawDescData)
	})
	return file_delta_proto_rawDescData
}

var file_delta_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_delta_proto_goTypes = []interface{}{
	(*Delta)(nil),                     // 0: api.public.logstream.Delta
	(*DeltaManifest)(nil),             // 1: api.public.logstream.DeltaManifest
	(*DeltaTargetManifest)(nil),       // 2: api.public.logstream.DeltaTargetManifest
	(*DeltaCommandManifest)(nil),      // 3: api.public.logstream.DeltaCommandManifest
	(*DeltaLog)(nil),                  // 4: api.public.logstream.DeltaLog
	(*DeltaManifest_FieldsDelta)(nil), // 5: api.public.logstream.DeltaManifest.FieldsDelta
	nil,                               // 6: api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry
	nil,                               // 7: api.public.logstream.DeltaTargetManifest.CommandsEntry
	(*BuildManifest)(nil),             // 8: api.public.logstream.BuildManifest
	(BuildStatus)(0),                  // 9: api.public.logstream.BuildStatus
}
var file_delta_proto_depIdxs = []int32{
	1,  // 0: api.public.logstream.Delta.delta_manifests:type_name -> api.public.logstream.DeltaManifest
	4,  // 1: api.public.logstream.Delta.delta_logs:type_name -> api.public.logstream.DeltaLog
	8,  // 2: api.public.logstream.DeltaManifest.reset_all:type_name -> api.public.logstream.BuildManifest
	5,  // 3: api.public.logstream.DeltaManifest.fields:type_name -> api.public.logstream.DeltaManifest.FieldsDelta
	9,  // 4: api.public.logstream.DeltaTargetManifest.status:type_name -> api.public.logstream.BuildStatus
	7,  // 5: api.public.logstream.DeltaTargetManifest.commands:type_name -> api.public.logstream.DeltaTargetManifest.CommandsEntry
	9,  // 6: api.public.logstream.DeltaCommandManifest.status:type_name -> api.public.logstream.BuildStatus
	9,  // 7: api.public.logstream.DeltaManifest.FieldsDelta.status:type_name -> api.public.logstream.BuildStatus
	6,  // 8: api.public.logstream.DeltaManifest.FieldsDelta.targets:type_name -> api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry
	2,  // 9: api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry.value:type_name -> api.public.logstream.DeltaTargetManifest
	3,  // 10: api.public.logstream.DeltaTargetManifest.CommandsEntry.value:type_name -> api.public.logstream.DeltaCommandManifest
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_delta_proto_init() }
func file_delta_proto_init() {
	if File_delta_proto != nil {
		return
	}
	file_manifest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_delta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delta); i {
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
		file_delta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaManifest); i {
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
		file_delta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaTargetManifest); i {
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
		file_delta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaCommandManifest); i {
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
		file_delta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaLog); i {
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
		file_delta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaManifest_FieldsDelta); i {
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
	file_delta_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DeltaManifest_Eof)(nil),
		(*DeltaManifest_ResetAll)(nil),
		(*DeltaManifest_Fields)(nil),
	}
	file_delta_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*DeltaLog_Eof)(nil),
		(*DeltaLog_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_delta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_delta_proto_goTypes,
		DependencyIndexes: file_delta_proto_depIdxs,
		MessageInfos:      file_delta_proto_msgTypes,
	}.Build()
	File_delta_proto = out.File
	file_delta_proto_rawDesc = nil
	file_delta_proto_goTypes = nil
	file_delta_proto_depIdxs = nil
}