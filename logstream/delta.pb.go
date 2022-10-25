// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: delta.proto

package logstream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Types that are assignable to DeltaTypeOneof:
	//	*Delta_DeltaManifests
	//	*Delta_DeltaLogs
	DeltaTypeOneof isDelta_DeltaTypeOneof `protobuf_oneof:"delta_type_oneof"`
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

func (m *Delta) GetDeltaTypeOneof() isDelta_DeltaTypeOneof {
	if m != nil {
		return m.DeltaTypeOneof
	}
	return nil
}

func (x *Delta) GetDeltaManifests() *DeltaManifest {
	if x, ok := x.GetDeltaTypeOneof().(*Delta_DeltaManifests); ok {
		return x.DeltaManifests
	}
	return nil
}

func (x *Delta) GetDeltaLogs() *DeltaLog {
	if x, ok := x.GetDeltaTypeOneof().(*Delta_DeltaLogs); ok {
		return x.DeltaLogs
	}
	return nil
}

type isDelta_DeltaTypeOneof interface {
	isDelta_DeltaTypeOneof()
}

type Delta_DeltaManifests struct {
	DeltaManifests *DeltaManifest `protobuf:"bytes,2,opt,name=delta_manifests,json=deltaManifests,proto3,oneof"`
}

type Delta_DeltaLogs struct {
	DeltaLogs *DeltaLog `protobuf:"bytes,3,opt,name=delta_logs,json=deltaLogs,proto3,oneof"`
}

func (*Delta_DeltaManifests) isDelta_DeltaTypeOneof() {}

func (*Delta_DeltaLogs) isDelta_DeltaTypeOneof() {}

type DeltaLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId     string                 `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	CommandIndex int32                  `protobuf:"varint,2,opt,name=command_index,json=commandIndex,proto3" json:"command_index,omitempty"`
	Stream       int32                  `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"` // stdout or stderr
	Timestamp    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Log          []byte                 `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *DeltaLog) Reset() {
	*x = DeltaLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaLog) ProtoMessage() {}

func (x *DeltaLog) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeltaLog.ProtoReflect.Descriptor instead.
func (*DeltaLog) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{1}
}

func (x *DeltaLog) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *DeltaLog) GetCommandIndex() int32 {
	if x != nil {
		return x.CommandIndex
	}
	return 0
}

func (x *DeltaLog) GetStream() int32 {
	if x != nil {
		return x.Stream
	}
	return 0
}

func (x *DeltaLog) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DeltaLog) GetLog() []byte {
	if x != nil {
		return x.Log
	}
	return nil
}

type DeltaManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to DeltaManifestOneof:
	//	*DeltaManifest_ResetAll
	//	*DeltaManifest_Fields
	DeltaManifestOneof isDeltaManifest_DeltaManifestOneof `protobuf_oneof:"delta_manifest_oneof"`
}

func (x *DeltaManifest) Reset() {
	*x = DeltaManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaManifest) ProtoMessage() {}

func (x *DeltaManifest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeltaManifest.ProtoReflect.Descriptor instead.
func (*DeltaManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{2}
}

func (m *DeltaManifest) GetDeltaManifestOneof() isDeltaManifest_DeltaManifestOneof {
	if m != nil {
		return m.DeltaManifestOneof
	}
	return nil
}

func (x *DeltaManifest) GetResetAll() *RunManifest {
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

type DeltaManifest_ResetAll struct {
	ResetAll *RunManifest `protobuf:"bytes,1,opt,name=reset_all,json=resetAll,proto3,oneof"`
}

type DeltaManifest_Fields struct {
	Fields *DeltaManifest_FieldsDelta `protobuf:"bytes,2,opt,name=fields,proto3,oneof"`
}

func (*DeltaManifest_ResetAll) isDeltaManifest_DeltaManifestOneof() {}

func (*DeltaManifest_Fields) isDeltaManifest_DeltaManifestOneof() {}

type DeltaTargetManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CanonicalName string                          `protobuf:"bytes,2,opt,name=canonical_name,json=canonicalName,proto3" json:"canonical_name,omitempty"`
	OverrideArgs  []string                        `protobuf:"bytes,3,rep,name=override_args,json=overrideArgs,proto3" json:"override_args,omitempty"`
	Platform      string                          `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Status        RunStatus                       `protobuf:"varint,5,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	StartedAt     *timestamppb.Timestamp          `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt       *timestamppb.Timestamp          `protobuf:"bytes,7,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	Commands      map[int32]*DeltaCommandManifest `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeltaTargetManifest) Reset() {
	*x = DeltaTargetManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaTargetManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaTargetManifest) ProtoMessage() {}

func (x *DeltaTargetManifest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeltaTargetManifest.ProtoReflect.Descriptor instead.
func (*DeltaTargetManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{3}
}

func (x *DeltaTargetManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeltaTargetManifest) GetCanonicalName() string {
	if x != nil {
		return x.CanonicalName
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

func (x *DeltaTargetManifest) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
}

func (x *DeltaTargetManifest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *DeltaTargetManifest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
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

	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status         RunStatus              `protobuf:"varint,2,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	HasCached      bool                   `protobuf:"varint,3,opt,name=has_cached,json=hasCached,proto3" json:"has_cached,omitempty"`
	IsCached       bool                   `protobuf:"varint,4,opt,name=is_cached,json=isCached,proto3" json:"is_cached,omitempty"`
	HasPush        bool                   `protobuf:"varint,5,opt,name=has_push,json=hasPush,proto3" json:"has_push,omitempty"`
	IsPush         bool                   `protobuf:"varint,6,opt,name=is_push,json=isPush,proto3" json:"is_push,omitempty"`
	HasLocal       bool                   `protobuf:"varint,7,opt,name=has_local,json=hasLocal,proto3" json:"has_local,omitempty"`
	IsLocal        bool                   `protobuf:"varint,8,opt,name=is_local,json=isLocal,proto3" json:"is_local,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt        *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	HasHasProgress bool                   `protobuf:"varint,11,opt,name=has_has_progress,json=hasHasProgress,proto3" json:"has_has_progress,omitempty"`
	HasProgress    bool                   `protobuf:"varint,12,opt,name=has_progress,json=hasProgress,proto3" json:"has_progress,omitempty"`
	Progress       int32                  `protobuf:"varint,13,opt,name=progress,proto3" json:"progress,omitempty"`
	ErrorMessage   string                 `protobuf:"bytes,14,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	SourceLocation *SourceLocation        `protobuf:"bytes,15,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
}

func (x *DeltaCommandManifest) Reset() {
	*x = DeltaCommandManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaCommandManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaCommandManifest) ProtoMessage() {}

func (x *DeltaCommandManifest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeltaCommandManifest.ProtoReflect.Descriptor instead.
func (*DeltaCommandManifest) Descriptor() ([]byte, []int) {
	return file_delta_proto_rawDescGZIP(), []int{4}
}

func (x *DeltaCommandManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeltaCommandManifest) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
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

func (x *DeltaCommandManifest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *DeltaCommandManifest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
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

func (x *DeltaCommandManifest) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *DeltaCommandManifest) GetSourceLocation() *SourceLocation {
	if x != nil {
		return x.SourceLocation
	}
	return nil
}

type DeltaManifest_FieldsDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartedAt *timestamppb.Timestamp          `protobuf:"bytes,1,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt   *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	Status    RunStatus                       `protobuf:"varint,3,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	Failure   *Failure                        `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
	Targets   map[string]*DeltaTargetManifest `protobuf:"bytes,5,rep,name=targets,proto3" json:"targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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
	return file_delta_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DeltaManifest_FieldsDelta) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *DeltaManifest_FieldsDelta) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *DeltaManifest_FieldsDelta) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
}

func (x *DeltaManifest_FieldsDelta) GetFailure() *Failure {
	if x != nil {
		return x.Failure
	}
	return nil
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
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xb0, 0x01,
	0x0a, 0x08, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x22, 0xe7, 0x04, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x49, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a,
	0xb0, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x65, 0x0a, 0x0c, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xfa, 0x03, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x53, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x67,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xda, 0x04, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73,
	0x50, 0x75, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x68, 0x61, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x5f, 0x68,
	0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x48, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*DeltaLog)(nil),                  // 1: api.public.logstream.DeltaLog
	(*DeltaManifest)(nil),             // 2: api.public.logstream.DeltaManifest
	(*DeltaTargetManifest)(nil),       // 3: api.public.logstream.DeltaTargetManifest
	(*DeltaCommandManifest)(nil),      // 4: api.public.logstream.DeltaCommandManifest
	(*DeltaManifest_FieldsDelta)(nil), // 5: api.public.logstream.DeltaManifest.FieldsDelta
	nil,                               // 6: api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry
	nil,                               // 7: api.public.logstream.DeltaTargetManifest.CommandsEntry
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
	(*RunManifest)(nil),               // 9: api.public.logstream.RunManifest
	(RunStatus)(0),                    // 10: api.public.logstream.RunStatus
	(*SourceLocation)(nil),            // 11: api.public.logstream.SourceLocation
	(*Failure)(nil),                   // 12: api.public.logstream.Failure
}
var file_delta_proto_depIdxs = []int32{
	2,  // 0: api.public.logstream.Delta.delta_manifests:type_name -> api.public.logstream.DeltaManifest
	1,  // 1: api.public.logstream.Delta.delta_logs:type_name -> api.public.logstream.DeltaLog
	8,  // 2: api.public.logstream.DeltaLog.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 3: api.public.logstream.DeltaManifest.reset_all:type_name -> api.public.logstream.RunManifest
	5,  // 4: api.public.logstream.DeltaManifest.fields:type_name -> api.public.logstream.DeltaManifest.FieldsDelta
	10, // 5: api.public.logstream.DeltaTargetManifest.status:type_name -> api.public.logstream.RunStatus
	8,  // 6: api.public.logstream.DeltaTargetManifest.started_at:type_name -> google.protobuf.Timestamp
	8,  // 7: api.public.logstream.DeltaTargetManifest.ended_at:type_name -> google.protobuf.Timestamp
	7,  // 8: api.public.logstream.DeltaTargetManifest.commands:type_name -> api.public.logstream.DeltaTargetManifest.CommandsEntry
	10, // 9: api.public.logstream.DeltaCommandManifest.status:type_name -> api.public.logstream.RunStatus
	8,  // 10: api.public.logstream.DeltaCommandManifest.started_at:type_name -> google.protobuf.Timestamp
	8,  // 11: api.public.logstream.DeltaCommandManifest.ended_at:type_name -> google.protobuf.Timestamp
	11, // 12: api.public.logstream.DeltaCommandManifest.source_location:type_name -> api.public.logstream.SourceLocation
	8,  // 13: api.public.logstream.DeltaManifest.FieldsDelta.started_at:type_name -> google.protobuf.Timestamp
	8,  // 14: api.public.logstream.DeltaManifest.FieldsDelta.ended_at:type_name -> google.protobuf.Timestamp
	10, // 15: api.public.logstream.DeltaManifest.FieldsDelta.status:type_name -> api.public.logstream.RunStatus
	12, // 16: api.public.logstream.DeltaManifest.FieldsDelta.failure:type_name -> api.public.logstream.Failure
	6,  // 17: api.public.logstream.DeltaManifest.FieldsDelta.targets:type_name -> api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry
	3,  // 18: api.public.logstream.DeltaManifest.FieldsDelta.TargetsEntry.value:type_name -> api.public.logstream.DeltaTargetManifest
	4,  // 19: api.public.logstream.DeltaTargetManifest.CommandsEntry.value:type_name -> api.public.logstream.DeltaCommandManifest
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
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
		file_delta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_delta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_delta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	file_delta_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Delta_DeltaManifests)(nil),
		(*Delta_DeltaLogs)(nil),
	}
	file_delta_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DeltaManifest_ResetAll)(nil),
		(*DeltaManifest_Fields)(nil),
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
