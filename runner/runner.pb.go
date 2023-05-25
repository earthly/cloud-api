// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: runner.proto

package runner

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

type RunnerStatus int32

const (
	RunnerStatus_RUNNER_STATUS_UNKNOWN RunnerStatus = 0
	RunnerStatus_RUNNER_STATUS_WAITING RunnerStatus = 1
	RunnerStatus_RUNNER_STATUS_READY   RunnerStatus = 2
	RunnerStatus_RUNNER_STATUS_ERROR   RunnerStatus = 3
)

// Enum value maps for RunnerStatus.
var (
	RunnerStatus_name = map[int32]string{
		0: "RUNNER_STATUS_UNKNOWN",
		1: "RUNNER_STATUS_WAITING",
		2: "RUNNER_STATUS_READY",
		3: "RUNNER_STATUS_ERROR",
	}
	RunnerStatus_value = map[string]int32{
		"RUNNER_STATUS_UNKNOWN": 0,
		"RUNNER_STATUS_WAITING": 1,
		"RUNNER_STATUS_READY":   2,
		"RUNNER_STATUS_ERROR":   3,
	}
)

func (x RunnerStatus) Enum() *RunnerStatus {
	p := new(RunnerStatus)
	*p = x
	return p
}

func (x RunnerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunnerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_runner_proto_enumTypes[0].Descriptor()
}

func (RunnerStatus) Type() protoreflect.EnumType {
	return &file_runner_proto_enumTypes[0]
}

func (x RunnerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunnerStatus.Descriptor instead.
func (RunnerStatus) EnumDescriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{0}
}

type GitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://docs.github.com/en/rest/checks/runs?apiVersion=2022-11-28
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo  string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Hash  string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	// head_branch represents the branch the build is running on (where changes are made)
	// Whereas, base_branch would be the branch changes are merged into (i.e., main)
	// Builds run on a SHA, but there we want to add the branch context for the appropriate trigger (push / pr)
	HeadBranch string `protobuf:"bytes,15,opt,name=head_branch,json=headBranch,proto3" json:"head_branch,omitempty"`
}

func (x *GitInfo) Reset() {
	*x = GitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitInfo) ProtoMessage() {}

func (x *GitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitInfo.ProtoReflect.Descriptor instead.
func (*GitInfo) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{0}
}

func (x *GitInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GitInfo) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GitInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *GitInfo) GetHeadBranch() string {
	if x != nil {
		return x.HeadBranch
	}
	return ""
}

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Path         string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Tag          string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Target       string   `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	RepoToken    string   `protobuf:"bytes,5,opt,name=repo_token,json=repoToken,proto3" json:"repo_token,omitempty"`
	BuildId      string   `protobuf:"bytes,6,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	PipelineId   string   `protobuf:"bytes,7,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	UserId       string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrgId        string   `protobuf:"bytes,9,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ProjectId    string   `protobuf:"bytes,10,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	PipelineName string   `protobuf:"bytes,11,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	GitInfo      *GitInfo `protobuf:"bytes,12,opt,name=git_info,json=gitInfo,proto3" json:"git_info,omitempty"`
	NoCache      bool     `protobuf:"varint,13,opt,name=no_cache,json=noCache,proto3" json:"no_cache,omitempty"`
	IsPush       bool     `protobuf:"varint,14,opt,name=is_push,json=isPush,proto3" json:"is_push,omitempty"`
	OrgName      string   `protobuf:"bytes,15,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	ProjectName  string   `protobuf:"bytes,16,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Timeout      string   `protobuf:"bytes,17,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{1}
}

func (x *BuildRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BuildRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BuildRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *BuildRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *BuildRequest) GetRepoToken() string {
	if x != nil {
		return x.RepoToken
	}
	return ""
}

func (x *BuildRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *BuildRequest) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *BuildRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BuildRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *BuildRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BuildRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *BuildRequest) GetGitInfo() *GitInfo {
	if x != nil {
		return x.GitInfo
	}
	return nil
}

func (x *BuildRequest) GetNoCache() bool {
	if x != nil {
		return x.NoCache
	}
	return false
}

func (x *BuildRequest) GetIsPush() bool {
	if x != nil {
		return x.IsPush
	}
	return false
}

func (x *BuildRequest) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *BuildRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *BuildRequest) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{2}
}

type CancelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{3}
}

func (x *CancelRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type CancelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cancelled bool `protobuf:"varint,1,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
}

func (x *CancelResponse) Reset() {
	*x = CancelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelResponse) ProtoMessage() {}

func (x *CancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelResponse.ProtoReflect.Descriptor instead.
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{4}
}

func (x *CancelResponse) GetCancelled() bool {
	if x != nil {
		return x.Cancelled
	}
	return false
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{5}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       RunnerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.public.runner.RunnerStatus" json:"status,omitempty"`
	StatusReason string       `protobuf:"bytes,2,opt,name=status_reason,json=statusReason,proto3" json:"status_reason,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{6}
}

func (x *StatusResponse) GetStatus() RunnerStatus {
	if x != nil {
		return x.Status
	}
	return RunnerStatus_RUNNER_STATUS_UNKNOWN
}

func (x *StatusResponse) GetStatusReason() string {
	if x != nil {
		return x.StatusReason
	}
	return ""
}

var File_runner_proto protoreflect.FileDescriptor

var file_runner_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x22, 0x68, 0x0a, 0x07, 0x47, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65,
	0x61, 0x64, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x65, 0x61, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0xf0, 0x03, 0x0a, 0x0c,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x67,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x67, 0x69, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x0f,
	0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2a, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x76, 0x0a, 0x0c,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15,
	0x52, 0x55, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x55, 0x4e, 0x4e, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x55, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52,
	0x55, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x03, 0x32, 0xf8, 0x01, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x4c, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_runner_proto_rawDescOnce sync.Once
	file_runner_proto_rawDescData = file_runner_proto_rawDesc
)

func file_runner_proto_rawDescGZIP() []byte {
	file_runner_proto_rawDescOnce.Do(func() {
		file_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_runner_proto_rawDescData)
	})
	return file_runner_proto_rawDescData
}

var file_runner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_runner_proto_goTypes = []interface{}{
	(RunnerStatus)(0),      // 0: api.public.runner.RunnerStatus
	(*GitInfo)(nil),        // 1: api.public.runner.GitInfo
	(*BuildRequest)(nil),   // 2: api.public.runner.BuildRequest
	(*BuildResponse)(nil),  // 3: api.public.runner.BuildResponse
	(*CancelRequest)(nil),  // 4: api.public.runner.CancelRequest
	(*CancelResponse)(nil), // 5: api.public.runner.CancelResponse
	(*StatusRequest)(nil),  // 6: api.public.runner.StatusRequest
	(*StatusResponse)(nil), // 7: api.public.runner.StatusResponse
}
var file_runner_proto_depIdxs = []int32{
	1, // 0: api.public.runner.BuildRequest.git_info:type_name -> api.public.runner.GitInfo
	0, // 1: api.public.runner.StatusResponse.status:type_name -> api.public.runner.RunnerStatus
	2, // 2: api.public.runner.Runner.Build:input_type -> api.public.runner.BuildRequest
	6, // 3: api.public.runner.Runner.Status:input_type -> api.public.runner.StatusRequest
	4, // 4: api.public.runner.Runner.Cancel:input_type -> api.public.runner.CancelRequest
	3, // 5: api.public.runner.Runner.Build:output_type -> api.public.runner.BuildResponse
	7, // 6: api.public.runner.Runner.Status:output_type -> api.public.runner.StatusResponse
	5, // 7: api.public.runner.Runner.Cancel:output_type -> api.public.runner.CancelResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_runner_proto_init() }
func file_runner_proto_init() {
	if File_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitInfo); i {
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
		file_runner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
		file_runner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
		file_runner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelRequest); i {
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
		file_runner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelResponse); i {
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
		file_runner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_runner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
			RawDescriptor: file_runner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runner_proto_goTypes,
		DependencyIndexes: file_runner_proto_depIdxs,
		EnumInfos:         file_runner_proto_enumTypes,
		MessageInfos:      file_runner_proto_msgTypes,
	}.Build()
	File_runner_proto = out.File
	file_runner_proto_rawDesc = nil
	file_runner_proto_goTypes = nil
	file_runner_proto_depIdxs = nil
}
