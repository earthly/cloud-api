// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: api/public/secrets/api.proto

package secrets

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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email                 string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	VerificationToken     string `protobuf:"bytes,2,opt,name=verificationToken,proto3" json:"verificationToken,omitempty"`
	Password              string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PublicKey             string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	AcceptTermsConditions bool   `protobuf:"varint,5,opt,name=acceptTermsConditions,proto3" json:"acceptTermsConditions,omitempty"`
	AcceptPrivacyPolicy   bool   `protobuf:"varint,6,opt,name=acceptPrivacyPolicy,proto3" json:"acceptPrivacyPolicy,omitempty"`
	DisplayName           string `protobuf:"bytes,7,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetVerificationToken() string {
	if x != nil {
		return x.VerificationToken
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *CreateAccountRequest) GetAcceptTermsConditions() bool {
	if x != nil {
		return x.AcceptTermsConditions
	}
	return false
}

func (x *CreateAccountRequest) GetAcceptPrivacyPolicy() bool {
	if x != nil {
		return x.AcceptPrivacyPolicy
	}
	return false
}

func (x *CreateAccountRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// only a response is required, because nothing is sent on request
type AuthChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Challenge string `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *AuthChallengeResponse) Reset() {
	*x = AuthChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthChallengeResponse) ProtoMessage() {}

func (x *AuthChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthChallengeResponse.ProtoReflect.Descriptor instead.
func (*AuthChallengeResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{2}
}

func (x *AuthChallengeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuthChallengeResponse) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

// only a response is required, because nothing is sent on request
type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	WriteAccess bool   `protobuf:"varint,3,opt,name=writeAccess,proto3" json:"writeAccess,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{3}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PingResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PingResponse) GetWriteAccess() bool {
	if x != nil {
		return x.WriteAccess
	}
	return false
}

type OrgPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Write bool   `protobuf:"varint,3,opt,name=write,proto3" json:"write,omitempty"`
}

func (x *OrgPermissions) Reset() {
	*x = OrgPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgPermissions) ProtoMessage() {}

func (x *OrgPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgPermissions.ProtoReflect.Descriptor instead.
func (*OrgPermissions) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{4}
}

func (x *OrgPermissions) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *OrgPermissions) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OrgPermissions) GetWrite() bool {
	if x != nil {
		return x.Write
	}
	return false
}

type ListOrgPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Permissions []*OrgPermissions `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *ListOrgPermissionsResponse) Reset() {
	*x = ListOrgPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrgPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrgPermissionsResponse) ProtoMessage() {}

func (x *ListOrgPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrgPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListOrgPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListOrgPermissionsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListOrgPermissionsResponse) GetPermissions() []*OrgPermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type OrgDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Admin bool   `protobuf:"varint,2,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *OrgDetail) Reset() {
	*x = OrgDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgDetail) ProtoMessage() {}

func (x *OrgDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgDetail.ProtoReflect.Descriptor instead.
func (*OrgDetail) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{6}
}

func (x *OrgDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrgDetail) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

type ListOrgsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details []*OrgDetail `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ListOrgsResponse) Reset() {
	*x = ListOrgsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrgsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrgsResponse) ProtoMessage() {}

func (x *ListOrgsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrgsResponse.ProtoReflect.Descriptor instead.
func (*ListOrgsResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrgsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListOrgsResponse) GetDetails() []*OrgDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Write  bool                   `protobuf:"varint,2,opt,name=write,proto3" json:"write,omitempty"`
	Expiry *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{8}
}

func (x *AuthToken) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthToken) GetWrite() bool {
	if x != nil {
		return x.Write
	}
	return false
}

func (x *AuthToken) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

type ListAuthTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Tokens  []*AuthToken `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *ListAuthTokensResponse) Reset() {
	*x = ListAuthTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokensResponse) ProtoMessage() {}

func (x *ListAuthTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokensResponse.ProtoReflect.Descriptor instead.
func (*ListAuthTokensResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{9}
}

func (x *ListAuthTokensResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListAuthTokensResponse) GetTokens() []*AuthToken {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type OAuthConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *OAuthConnectRequest) Reset() {
	*x = OAuthConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthConnectRequest) ProtoMessage() {}

func (x *OAuthConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthConnectRequest.ProtoReflect.Descriptor instead.
func (*OAuthConnectRequest) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{10}
}

func (x *OAuthConnectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type OAuthConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name   string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Write  bool                   `protobuf:"varint,3,opt,name=write,proto3" json:"write,omitempty"`
	Expiry *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *OAuthConnectResponse) Reset() {
	*x = OAuthConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthConnectResponse) ProtoMessage() {}

func (x *OAuthConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthConnectResponse.ProtoReflect.Descriptor instead.
func (*OAuthConnectResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{11}
}

func (x *OAuthConnectResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *OAuthConnectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OAuthConnectResponse) GetWrite() bool {
	if x != nil {
		return x.Write
	}
	return false
}

func (x *OAuthConnectResponse) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

// only a response is required, because nothing is sent on request
// login data is in headers
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expiry *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_public_secrets_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_public_secrets_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_public_secrets_api_proto_rawDescGZIP(), []int{12}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

var File_api_public_secrets_api_proto protoreflect.FileDescriptor

var file_api_public_secrets_api_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x15, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x30, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x0e, 0x4f, 0x72,
	0x67, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x7c, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x4f,
	0x72, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x4f, 0x72,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x22, 0x65, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x72, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x69, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x2b,
	0x0a, 0x13, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x14,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x59, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_public_secrets_api_proto_rawDescOnce sync.Once
	file_api_public_secrets_api_proto_rawDescData = file_api_public_secrets_api_proto_rawDesc
)

func file_api_public_secrets_api_proto_rawDescGZIP() []byte {
	file_api_public_secrets_api_proto_rawDescOnce.Do(func() {
		file_api_public_secrets_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_public_secrets_api_proto_rawDescData)
	})
	return file_api_public_secrets_api_proto_rawDescData
}

var file_api_public_secrets_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_public_secrets_api_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),       // 0: api.public.secrets.CreateAccountRequest
	(*CreateAccountResponse)(nil),      // 1: api.public.secrets.CreateAccountResponse
	(*AuthChallengeResponse)(nil),      // 2: api.public.secrets.AuthChallengeResponse
	(*PingResponse)(nil),               // 3: api.public.secrets.PingResponse
	(*OrgPermissions)(nil),             // 4: api.public.secrets.OrgPermissions
	(*ListOrgPermissionsResponse)(nil), // 5: api.public.secrets.ListOrgPermissionsResponse
	(*OrgDetail)(nil),                  // 6: api.public.secrets.OrgDetail
	(*ListOrgsResponse)(nil),           // 7: api.public.secrets.ListOrgsResponse
	(*AuthToken)(nil),                  // 8: api.public.secrets.AuthToken
	(*ListAuthTokensResponse)(nil),     // 9: api.public.secrets.ListAuthTokensResponse
	(*OAuthConnectRequest)(nil),        // 10: api.public.secrets.OAuthConnectRequest
	(*OAuthConnectResponse)(nil),       // 11: api.public.secrets.OAuthConnectResponse
	(*LoginResponse)(nil),              // 12: api.public.secrets.LoginResponse
	(*timestamppb.Timestamp)(nil),      // 13: google.protobuf.Timestamp
}
var file_api_public_secrets_api_proto_depIdxs = []int32{
	4,  // 0: api.public.secrets.ListOrgPermissionsResponse.permissions:type_name -> api.public.secrets.OrgPermissions
	6,  // 1: api.public.secrets.ListOrgsResponse.details:type_name -> api.public.secrets.OrgDetail
	13, // 2: api.public.secrets.AuthToken.expiry:type_name -> google.protobuf.Timestamp
	8,  // 3: api.public.secrets.ListAuthTokensResponse.tokens:type_name -> api.public.secrets.AuthToken
	13, // 4: api.public.secrets.OAuthConnectResponse.expiry:type_name -> google.protobuf.Timestamp
	13, // 5: api.public.secrets.LoginResponse.expiry:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_public_secrets_api_proto_init() }
func file_api_public_secrets_api_proto_init() {
	if File_api_public_secrets_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_public_secrets_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_api_public_secrets_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthChallengeResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgPermissions); i {
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
		file_api_public_secrets_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrgPermissionsResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgDetail); i {
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
		file_api_public_secrets_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrgsResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_api_public_secrets_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthTokensResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthConnectRequest); i {
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
		file_api_public_secrets_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthConnectResponse); i {
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
		file_api_public_secrets_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_api_public_secrets_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_public_secrets_api_proto_goTypes,
		DependencyIndexes: file_api_public_secrets_api_proto_depIdxs,
		MessageInfos:      file_api_public_secrets_api_proto_msgTypes,
	}.Build()
	File_api_public_secrets_api_proto = out.File
	file_api_public_secrets_api_proto_rawDesc = nil
	file_api_public_secrets_api_proto_goTypes = nil
	file_api_public_secrets_api_proto_depIdxs = nil
}
