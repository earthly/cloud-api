// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: manifest.proto

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

type RunStatus int32

const (
	RunStatus_RUN_STATUS_UNKNOWN     RunStatus = 0
	RunStatus_RUN_STATUS_NOT_STARTED RunStatus = 1
	RunStatus_RUN_STATUS_IN_PROGRESS RunStatus = 2
	RunStatus_RUN_STATUS_SUCCESS     RunStatus = 3
	RunStatus_RUN_STATUS_FAILURE     RunStatus = 4
	RunStatus_RUN_STATUS_CANCELED    RunStatus = 5
)

// Enum value maps for RunStatus.
var (
	RunStatus_name = map[int32]string{
		0: "RUN_STATUS_UNKNOWN",
		1: "RUN_STATUS_NOT_STARTED",
		2: "RUN_STATUS_IN_PROGRESS",
		3: "RUN_STATUS_SUCCESS",
		4: "RUN_STATUS_FAILURE",
		5: "RUN_STATUS_CANCELED",
	}
	RunStatus_value = map[string]int32{
		"RUN_STATUS_UNKNOWN":     0,
		"RUN_STATUS_NOT_STARTED": 1,
		"RUN_STATUS_IN_PROGRESS": 2,
		"RUN_STATUS_SUCCESS":     3,
		"RUN_STATUS_FAILURE":     4,
		"RUN_STATUS_CANCELED":    5,
	}
)

func (x RunStatus) Enum() *RunStatus {
	p := new(RunStatus)
	*p = x
	return p
}

func (x RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_manifest_proto_enumTypes[0].Descriptor()
}

func (RunStatus) Type() protoreflect.EnumType {
	return &file_manifest_proto_enumTypes[0]
}

func (x RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStatus.Descriptor instead.
func (RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{0}
}

type FailureType int32

const (
	FailureType_FAILURE_TYPE_UNKNOWN        FailureType = 0
	FailureType_FAILURE_TYPE_NONZERO_EXIT   FailureType = 1
	FailureType_FAILURE_TYPE_FILE_NOT_FOUND FailureType = 2
	FailureType_FAILURE_TYPE_SYNTAX         FailureType = 3
)

// Enum value maps for FailureType.
var (
	FailureType_name = map[int32]string{
		0: "FAILURE_TYPE_UNKNOWN",
		1: "FAILURE_TYPE_NONZERO_EXIT",
		2: "FAILURE_TYPE_FILE_NOT_FOUND",
		3: "FAILURE_TYPE_SYNTAX",
	}
	FailureType_value = map[string]int32{
		"FAILURE_TYPE_UNKNOWN":        0,
		"FAILURE_TYPE_NONZERO_EXIT":   1,
		"FAILURE_TYPE_FILE_NOT_FOUND": 2,
		"FAILURE_TYPE_SYNTAX":         3,
	}
)

func (x FailureType) Enum() *FailureType {
	p := new(FailureType)
	*p = x
	return p
}

func (x FailureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailureType) Descriptor() protoreflect.EnumDescriptor {
	return file_manifest_proto_enumTypes[1].Descriptor()
}

func (FailureType) Type() protoreflect.EnumType {
	return &file_manifest_proto_enumTypes[1]
}

func (x FailureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailureType.Descriptor instead.
func (FailureType) EnumDescriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{1}
}

type RunManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId    string                     `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Version    int32                      `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt  *timestamppb.Timestamp     `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt  *timestamppb.Timestamp     `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt    *timestamppb.Timestamp     `protobuf:"bytes,5,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	Status     RunStatus                  `protobuf:"varint,6,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	Targets    map[string]*TargetManifest `protobuf:"bytes,7,rep,name=targets,proto3" json:"targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MainTarget string                     `protobuf:"bytes,8,opt,name=main_target,json=mainTarget,proto3" json:"main_target,omitempty"`
	Failure    *Failure                   `protobuf:"bytes,9,opt,name=failure,proto3" json:"failure,omitempty"`
	// user_id is the UUID of the user that ran the build.
	UserId string `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// org_id is the Organization the build belongs to.
	OrgId string `protobuf:"bytes,11,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// project_id is the ID of the CI Project containing the build.
	// It my be empty for builds that are not run on the CI.
	ProjectId string `protobuf:"bytes,12,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *RunManifest) Reset() {
	*x = RunManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunManifest) ProtoMessage() {}

func (x *RunManifest) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunManifest.ProtoReflect.Descriptor instead.
func (*RunManifest) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *RunManifest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *RunManifest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RunManifest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RunManifest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *RunManifest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *RunManifest) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
}

func (x *RunManifest) GetTargets() map[string]*TargetManifest {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *RunManifest) GetMainTarget() string {
	if x != nil {
		return x.MainTarget
	}
	return ""
}

func (x *RunManifest) GetFailure() *Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

func (x *RunManifest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RunManifest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *RunManifest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            FailureType `protobuf:"varint,1,opt,name=type,proto3,enum=api.public.logstream.FailureType" json:"type,omitempty"`
	TargetId        string      `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	HasCommandIndex bool        `protobuf:"varint,3,opt,name=has_command_index,json=hasCommandIndex,proto3" json:"has_command_index,omitempty"`
	CommandIndex    int32       `protobuf:"varint,4,opt,name=command_index,json=commandIndex,proto3" json:"command_index,omitempty"`
	Summary         string      `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Output          []byte      `protobuf:"bytes,6,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *Failure) GetType() FailureType {
	if x != nil {
		return x.Type
	}
	return FailureType_FAILURE_TYPE_UNKNOWN
}

func (x *Failure) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Failure) GetHasCommandIndex() bool {
	if x != nil {
		return x.HasCommandIndex
	}
	return false
}

func (x *Failure) GetCommandIndex() int32 {
	if x != nil {
		return x.CommandIndex
	}
	return 0
}

func (x *Failure) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Failure) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type TargetManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                        // e.g. +build
	CanonicalName string                 `protobuf:"bytes,2,opt,name=canonical_name,json=canonicalName,proto3" json:"canonical_name,omitempty"` // e.g. github.com/foo/bar/buz:main+build
	OverrideArgs  []string               `protobuf:"bytes,3,rep,name=override_args,json=overrideArgs,proto3" json:"override_args,omitempty"`
	Platform      string                 `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Status        RunStatus              `protobuf:"varint,5,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	StartedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	Commands      []*CommandManifest     `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *TargetManifest) Reset() {
	*x = TargetManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetManifest) ProtoMessage() {}

func (x *TargetManifest) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetManifest.ProtoReflect.Descriptor instead.
func (*TargetManifest) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *TargetManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TargetManifest) GetCanonicalName() string {
	if x != nil {
		return x.CanonicalName
	}
	return ""
}

func (x *TargetManifest) GetOverrideArgs() []string {
	if x != nil {
		return x.OverrideArgs
	}
	return nil
}

func (x *TargetManifest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *TargetManifest) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
}

func (x *TargetManifest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *TargetManifest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *TargetManifest) GetCommands() []*CommandManifest {
	if x != nil {
		return x.Commands
	}
	return nil
}

type CommandManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status      RunStatus              `protobuf:"varint,2,opt,name=status,proto3,enum=api.public.logstream.RunStatus" json:"status,omitempty"`
	IsCached    bool                   `protobuf:"varint,3,opt,name=is_cached,json=isCached,proto3" json:"is_cached,omitempty"`
	IsPush      bool                   `protobuf:"varint,4,opt,name=is_push,json=isPush,proto3" json:"is_push,omitempty"`
	IsLocal     bool                   `protobuf:"varint,5,opt,name=is_local,json=isLocal,proto3" json:"is_local,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	HasProgress bool                   `protobuf:"varint,8,opt,name=has_progress,json=hasProgress,proto3" json:"has_progress,omitempty"`
	Progress    int32                  `protobuf:"varint,9,opt,name=progress,proto3" json:"progress,omitempty"`
	// error_message is a string representation of an error related to this command.
	// The presence of an error here might not mean that the command has failed,
	// if the error is transient.
	ErrorMessage   string          `protobuf:"bytes,10,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	SourceLocation *SourceLocation `protobuf:"bytes,11,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
}

func (x *CommandManifest) Reset() {
	*x = CommandManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandManifest) ProtoMessage() {}

func (x *CommandManifest) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandManifest.ProtoReflect.Descriptor instead.
func (*CommandManifest) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{3}
}

func (x *CommandManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandManifest) GetStatus() RunStatus {
	if x != nil {
		return x.Status
	}
	return RunStatus_RUN_STATUS_UNKNOWN
}

func (x *CommandManifest) GetIsCached() bool {
	if x != nil {
		return x.IsCached
	}
	return false
}

func (x *CommandManifest) GetIsPush() bool {
	if x != nil {
		return x.IsPush
	}
	return false
}

func (x *CommandManifest) GetIsLocal() bool {
	if x != nil {
		return x.IsLocal
	}
	return false
}

func (x *CommandManifest) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *CommandManifest) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *CommandManifest) GetHasProgress() bool {
	if x != nil {
		return x.HasProgress
	}
	return false
}

func (x *CommandManifest) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *CommandManifest) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *CommandManifest) GetSourceLocation() *SourceLocation {
	if x != nil {
		return x.SourceLocation
	}
	return nil
}

type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repository_url is the URL of the repository.
	RepositoryUrl string `protobuf:"bytes,1,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
	// file is the name of the Earthfile relative to the repository.
	File        string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	StartLine   int32  `protobuf:"varint,3,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	StartColumn int32  `protobuf:"varint,4,opt,name=start_column,json=startColumn,proto3" json:"start_column,omitempty"`
	EndLine     int32  `protobuf:"varint,5,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
	EndColumn   int32  `protobuf:"varint,6,opt,name=end_column,json=endColumn,proto3" json:"end_column,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manifest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_manifest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_manifest_proto_rawDescGZIP(), []int{4}
}

func (x *SourceLocation) GetRepositoryUrl() string {
	if x != nil {
		return x.RepositoryUrl
	}
	return ""
}

func (x *SourceLocation) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *SourceLocation) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *SourceLocation) GetStartColumn() int32 {
	if x != nil {
		return x.StartColumn
	}
	return 0
}

func (x *SourceLocation) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *SourceLocation) GetEndColumn() int32 {
	if x != nil {
		return x.EndColumn
	}
	return 0
}

var File_manifest_proto protoreflect.FileDescriptor

var file_manifest_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x04, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x48, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x37, 0x0a,
	0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x60, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0, 0x01, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xfa, 0x02, 0x0a, 0x0e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x6f, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xd4, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x73, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68,
	0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4d, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc7,
	0x01, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2a, 0xa4, 0x01, 0x0a, 0x09, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a,
	0x0a, 0x16, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x55,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x2a,
	0x80, 0x01, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x14, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x5a, 0x45, 0x52,
	0x4f, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x54, 0x41, 0x58,
	0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manifest_proto_rawDescOnce sync.Once
	file_manifest_proto_rawDescData = file_manifest_proto_rawDesc
)

func file_manifest_proto_rawDescGZIP() []byte {
	file_manifest_proto_rawDescOnce.Do(func() {
		file_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_manifest_proto_rawDescData)
	})
	return file_manifest_proto_rawDescData
}

var file_manifest_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manifest_proto_goTypes = []interface{}{
	(RunStatus)(0),                // 0: api.public.logstream.RunStatus
	(FailureType)(0),              // 1: api.public.logstream.FailureType
	(*RunManifest)(nil),           // 2: api.public.logstream.RunManifest
	(*Failure)(nil),               // 3: api.public.logstream.Failure
	(*TargetManifest)(nil),        // 4: api.public.logstream.TargetManifest
	(*CommandManifest)(nil),       // 5: api.public.logstream.CommandManifest
	(*SourceLocation)(nil),        // 6: api.public.logstream.SourceLocation
	nil,                           // 7: api.public.logstream.RunManifest.TargetsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_manifest_proto_depIdxs = []int32{
	8,  // 0: api.public.logstream.RunManifest.created_at:type_name -> google.protobuf.Timestamp
	8,  // 1: api.public.logstream.RunManifest.started_at:type_name -> google.protobuf.Timestamp
	8,  // 2: api.public.logstream.RunManifest.ended_at:type_name -> google.protobuf.Timestamp
	0,  // 3: api.public.logstream.RunManifest.status:type_name -> api.public.logstream.RunStatus
	7,  // 4: api.public.logstream.RunManifest.targets:type_name -> api.public.logstream.RunManifest.TargetsEntry
	3,  // 5: api.public.logstream.RunManifest.failure:type_name -> api.public.logstream.Failure
	1,  // 6: api.public.logstream.Failure.type:type_name -> api.public.logstream.FailureType
	0,  // 7: api.public.logstream.TargetManifest.status:type_name -> api.public.logstream.RunStatus
	8,  // 8: api.public.logstream.TargetManifest.started_at:type_name -> google.protobuf.Timestamp
	8,  // 9: api.public.logstream.TargetManifest.ended_at:type_name -> google.protobuf.Timestamp
	5,  // 10: api.public.logstream.TargetManifest.commands:type_name -> api.public.logstream.CommandManifest
	0,  // 11: api.public.logstream.CommandManifest.status:type_name -> api.public.logstream.RunStatus
	8,  // 12: api.public.logstream.CommandManifest.started_at:type_name -> google.protobuf.Timestamp
	8,  // 13: api.public.logstream.CommandManifest.ended_at:type_name -> google.protobuf.Timestamp
	6,  // 14: api.public.logstream.CommandManifest.source_location:type_name -> api.public.logstream.SourceLocation
	4,  // 15: api.public.logstream.RunManifest.TargetsEntry.value:type_name -> api.public.logstream.TargetManifest
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_manifest_proto_init() }
func file_manifest_proto_init() {
	if File_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunManifest); i {
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
		file_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
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
		file_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetManifest); i {
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
		file_manifest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandManifest); i {
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
		file_manifest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
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
			RawDescriptor: file_manifest_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manifest_proto_goTypes,
		DependencyIndexes: file_manifest_proto_depIdxs,
		EnumInfos:         file_manifest_proto_enumTypes,
		MessageInfos:      file_manifest_proto_msgTypes,
	}.Build()
	File_manifest_proto = out.File
	file_manifest_proto_rawDesc = nil
	file_manifest_proto_goTypes = nil
	file_manifest_proto_depIdxs = nil
}
