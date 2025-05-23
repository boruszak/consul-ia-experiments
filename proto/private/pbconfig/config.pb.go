// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: private/pbconfig/config.proto

package pbconfig

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datacenter        string       `protobuf:"bytes,1,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
	PrimaryDatacenter string       `protobuf:"bytes,2,opt,name=PrimaryDatacenter,proto3" json:"PrimaryDatacenter,omitempty"`
	NodeName          string       `protobuf:"bytes,3,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	SegmentName       string       `protobuf:"bytes,4,opt,name=SegmentName,proto3" json:"SegmentName,omitempty"`
	Partition         string       `protobuf:"bytes,9,opt,name=Partition,proto3" json:"Partition,omitempty"`
	ACL               *ACL         `protobuf:"bytes,5,opt,name=ACL,proto3" json:"ACL,omitempty"`
	AutoEncrypt       *AutoEncrypt `protobuf:"bytes,6,opt,name=AutoEncrypt,proto3" json:"AutoEncrypt,omitempty"`
	Gossip            *Gossip      `protobuf:"bytes,7,opt,name=Gossip,proto3" json:"Gossip,omitempty"`
	TLS               *TLS         `protobuf:"bytes,8,opt,name=TLS,proto3" json:"TLS,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *Config) GetPrimaryDatacenter() string {
	if x != nil {
		return x.PrimaryDatacenter
	}
	return ""
}

func (x *Config) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Config) GetSegmentName() string {
	if x != nil {
		return x.SegmentName
	}
	return ""
}

func (x *Config) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *Config) GetACL() *ACL {
	if x != nil {
		return x.ACL
	}
	return nil
}

func (x *Config) GetAutoEncrypt() *AutoEncrypt {
	if x != nil {
		return x.AutoEncrypt
	}
	return nil
}

func (x *Config) GetGossip() *Gossip {
	if x != nil {
		return x.Gossip
	}
	return nil
}

func (x *Config) GetTLS() *TLS {
	if x != nil {
		return x.TLS
	}
	return nil
}

type Gossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encryption   *GossipEncryption `protobuf:"bytes,1,opt,name=Encryption,proto3" json:"Encryption,omitempty"`
	RetryJoinLAN []string          `protobuf:"bytes,2,rep,name=RetryJoinLAN,proto3" json:"RetryJoinLAN,omitempty"`
}

func (x *Gossip) Reset() {
	*x = Gossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gossip) ProtoMessage() {}

func (x *Gossip) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gossip.ProtoReflect.Descriptor instead.
func (*Gossip) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{1}
}

func (x *Gossip) GetEncryption() *GossipEncryption {
	if x != nil {
		return x.Encryption
	}
	return nil
}

func (x *Gossip) GetRetryJoinLAN() []string {
	if x != nil {
		return x.RetryJoinLAN
	}
	return nil
}

type GossipEncryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	VerifyIncoming bool   `protobuf:"varint,2,opt,name=VerifyIncoming,proto3" json:"VerifyIncoming,omitempty"`
	VerifyOutgoing bool   `protobuf:"varint,3,opt,name=VerifyOutgoing,proto3" json:"VerifyOutgoing,omitempty"`
}

func (x *GossipEncryption) Reset() {
	*x = GossipEncryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipEncryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipEncryption) ProtoMessage() {}

func (x *GossipEncryption) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipEncryption.ProtoReflect.Descriptor instead.
func (*GossipEncryption) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{2}
}

func (x *GossipEncryption) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GossipEncryption) GetVerifyIncoming() bool {
	if x != nil {
		return x.VerifyIncoming
	}
	return false
}

func (x *GossipEncryption) GetVerifyOutgoing() bool {
	if x != nil {
		return x.VerifyOutgoing
	}
	return false
}

type TLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyOutgoing       bool   `protobuf:"varint,1,opt,name=VerifyOutgoing,proto3" json:"VerifyOutgoing,omitempty"`
	VerifyServerHostname bool   `protobuf:"varint,2,opt,name=VerifyServerHostname,proto3" json:"VerifyServerHostname,omitempty"`
	CipherSuites         string `protobuf:"bytes,3,opt,name=CipherSuites,proto3" json:"CipherSuites,omitempty"`
	MinVersion           string `protobuf:"bytes,4,opt,name=MinVersion,proto3" json:"MinVersion,omitempty"`
	// Deprecated_PreferServerCipherSuites is deprecated. It is no longer
	// populated and should be ignored by clients.
	//
	// Deprecated: Marked as deprecated in private/pbconfig/config.proto.
	Deprecated_PreferServerCipherSuites bool `protobuf:"varint,5,opt,name=Deprecated_PreferServerCipherSuites,json=DeprecatedPreferServerCipherSuites,proto3" json:"Deprecated_PreferServerCipherSuites,omitempty"`
}

func (x *TLS) Reset() {
	*x = TLS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLS) ProtoMessage() {}

func (x *TLS) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLS.ProtoReflect.Descriptor instead.
func (*TLS) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{3}
}

func (x *TLS) GetVerifyOutgoing() bool {
	if x != nil {
		return x.VerifyOutgoing
	}
	return false
}

func (x *TLS) GetVerifyServerHostname() bool {
	if x != nil {
		return x.VerifyServerHostname
	}
	return false
}

func (x *TLS) GetCipherSuites() string {
	if x != nil {
		return x.CipherSuites
	}
	return ""
}

func (x *TLS) GetMinVersion() string {
	if x != nil {
		return x.MinVersion
	}
	return ""
}

// Deprecated: Marked as deprecated in private/pbconfig/config.proto.
func (x *TLS) GetDeprecated_PreferServerCipherSuites() bool {
	if x != nil {
		return x.Deprecated_PreferServerCipherSuites
	}
	return false
}

type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled             bool       `protobuf:"varint,1,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	PolicyTTL           string     `protobuf:"bytes,2,opt,name=PolicyTTL,proto3" json:"PolicyTTL,omitempty"`
	RoleTTL             string     `protobuf:"bytes,3,opt,name=RoleTTL,proto3" json:"RoleTTL,omitempty"`
	TokenTTL            string     `protobuf:"bytes,4,opt,name=TokenTTL,proto3" json:"TokenTTL,omitempty"`
	DownPolicy          string     `protobuf:"bytes,5,opt,name=DownPolicy,proto3" json:"DownPolicy,omitempty"`
	DefaultPolicy       string     `protobuf:"bytes,6,opt,name=DefaultPolicy,proto3" json:"DefaultPolicy,omitempty"`
	EnableKeyListPolicy bool       `protobuf:"varint,7,opt,name=EnableKeyListPolicy,proto3" json:"EnableKeyListPolicy,omitempty"`
	Tokens              *ACLTokens `protobuf:"bytes,8,opt,name=Tokens,proto3" json:"Tokens,omitempty"`
	// Deprecated_DisabledTTL is deprecated. It is no longer populated and should
	// be ignored by clients.
	//
	// Deprecated: Marked as deprecated in private/pbconfig/config.proto.
	Deprecated_DisabledTTL string `protobuf:"bytes,9,opt,name=Deprecated_DisabledTTL,json=DeprecatedDisabledTTL,proto3" json:"Deprecated_DisabledTTL,omitempty"`
	EnableTokenPersistence bool   `protobuf:"varint,10,opt,name=EnableTokenPersistence,proto3" json:"EnableTokenPersistence,omitempty"`
	MSPDisableBootstrap    bool   `protobuf:"varint,11,opt,name=MSPDisableBootstrap,proto3" json:"MSPDisableBootstrap,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{4}
}

func (x *ACL) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ACL) GetPolicyTTL() string {
	if x != nil {
		return x.PolicyTTL
	}
	return ""
}

func (x *ACL) GetRoleTTL() string {
	if x != nil {
		return x.RoleTTL
	}
	return ""
}

func (x *ACL) GetTokenTTL() string {
	if x != nil {
		return x.TokenTTL
	}
	return ""
}

func (x *ACL) GetDownPolicy() string {
	if x != nil {
		return x.DownPolicy
	}
	return ""
}

func (x *ACL) GetDefaultPolicy() string {
	if x != nil {
		return x.DefaultPolicy
	}
	return ""
}

func (x *ACL) GetEnableKeyListPolicy() bool {
	if x != nil {
		return x.EnableKeyListPolicy
	}
	return false
}

func (x *ACL) GetTokens() *ACLTokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

// Deprecated: Marked as deprecated in private/pbconfig/config.proto.
func (x *ACL) GetDeprecated_DisabledTTL() string {
	if x != nil {
		return x.Deprecated_DisabledTTL
	}
	return ""
}

func (x *ACL) GetEnableTokenPersistence() bool {
	if x != nil {
		return x.EnableTokenPersistence
	}
	return false
}

func (x *ACL) GetMSPDisableBootstrap() bool {
	if x != nil {
		return x.MSPDisableBootstrap
	}
	return false
}

type ACLTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialManagement      string                     `protobuf:"bytes,1,opt,name=InitialManagement,proto3" json:"InitialManagement,omitempty"`
	Replication            string                     `protobuf:"bytes,2,opt,name=Replication,proto3" json:"Replication,omitempty"`
	AgentRecovery          string                     `protobuf:"bytes,3,opt,name=AgentRecovery,proto3" json:"AgentRecovery,omitempty"`
	Default                string                     `protobuf:"bytes,4,opt,name=Default,proto3" json:"Default,omitempty"`
	Agent                  string                     `protobuf:"bytes,5,opt,name=Agent,proto3" json:"Agent,omitempty"`
	ManagedServiceProvider []*ACLServiceProviderToken `protobuf:"bytes,6,rep,name=ManagedServiceProvider,proto3" json:"ManagedServiceProvider,omitempty"`
}

func (x *ACLTokens) Reset() {
	*x = ACLTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLTokens) ProtoMessage() {}

func (x *ACLTokens) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLTokens.ProtoReflect.Descriptor instead.
func (*ACLTokens) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{5}
}

func (x *ACLTokens) GetInitialManagement() string {
	if x != nil {
		return x.InitialManagement
	}
	return ""
}

func (x *ACLTokens) GetReplication() string {
	if x != nil {
		return x.Replication
	}
	return ""
}

func (x *ACLTokens) GetAgentRecovery() string {
	if x != nil {
		return x.AgentRecovery
	}
	return ""
}

func (x *ACLTokens) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *ACLTokens) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *ACLTokens) GetManagedServiceProvider() []*ACLServiceProviderToken {
	if x != nil {
		return x.ManagedServiceProvider
	}
	return nil
}

type ACLServiceProviderToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessorID string `protobuf:"bytes,1,opt,name=AccessorID,proto3" json:"AccessorID,omitempty"`
	SecretID   string `protobuf:"bytes,2,opt,name=SecretID,proto3" json:"SecretID,omitempty"`
}

func (x *ACLServiceProviderToken) Reset() {
	*x = ACLServiceProviderToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACLServiceProviderToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACLServiceProviderToken) ProtoMessage() {}

func (x *ACLServiceProviderToken) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACLServiceProviderToken.ProtoReflect.Descriptor instead.
func (*ACLServiceProviderToken) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{6}
}

func (x *ACLServiceProviderToken) GetAccessorID() string {
	if x != nil {
		return x.AccessorID
	}
	return ""
}

func (x *ACLServiceProviderToken) GetSecretID() string {
	if x != nil {
		return x.SecretID
	}
	return ""
}

type AutoEncrypt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TLS      bool     `protobuf:"varint,1,opt,name=TLS,proto3" json:"TLS,omitempty"`
	DNSSAN   []string `protobuf:"bytes,2,rep,name=DNSSAN,proto3" json:"DNSSAN,omitempty"`
	IPSAN    []string `protobuf:"bytes,3,rep,name=IPSAN,proto3" json:"IPSAN,omitempty"`
	AllowTLS bool     `protobuf:"varint,4,opt,name=AllowTLS,proto3" json:"AllowTLS,omitempty"`
}

func (x *AutoEncrypt) Reset() {
	*x = AutoEncrypt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbconfig_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoEncrypt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoEncrypt) ProtoMessage() {}

func (x *AutoEncrypt) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbconfig_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoEncrypt.ProtoReflect.Descriptor instead.
func (*AutoEncrypt) Descriptor() ([]byte, []int) {
	return file_private_pbconfig_config_proto_rawDescGZIP(), []int{7}
}

func (x *AutoEncrypt) GetTLS() bool {
	if x != nil {
		return x.TLS
	}
	return false
}

func (x *AutoEncrypt) GetDNSSAN() []string {
	if x != nil {
		return x.DNSSAN
	}
	return nil
}

func (x *AutoEncrypt) GetIPSAN() []string {
	if x != nil {
		return x.IPSAN
	}
	return nil
}

func (x *AutoEncrypt) GetAllowTLS() bool {
	if x != nil {
		return x.AllowTLS
	}
	return false
}

var File_private_pbconfig_config_proto protoreflect.FileDescriptor

var file_private_pbconfig_config_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xb7, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x03, 0x41, 0x43, 0x4c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x43, 0x4c, 0x52, 0x03, 0x41, 0x43, 0x4c, 0x12,
	0x4f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x12, 0x40, 0x0a, 0x06, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x52, 0x06, 0x47, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x12, 0x37, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x03, 0x54, 0x4c, 0x53, 0x22, 0x80, 0x01, 0x0a, 0x06,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x52, 0x0a, 0x0a, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x41, 0x4e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x41, 0x4e, 0x22, 0x74,
	0x0a, 0x10, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e,
	0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x67,
	0x6f, 0x69, 0x6e, 0x67, 0x22, 0xfa, 0x01, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x26, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x67,
	0x6f, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x4d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x23,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x22, 0x44,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x73, 0x22, 0xd5, 0x03, 0x0a, 0x03, 0x41, 0x43, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x54, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x54,
	0x4c, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x54, 0x4c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x54, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x54, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x54, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6f, 0x77,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x30, 0x0a,
	0x13, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x43, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x41, 0x43, 0x4c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x06, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x16, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x54, 0x54, 0x4c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x15, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x54, 0x54, 0x4c, 0x12,
	0x36, 0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x4d, 0x53, 0x50, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x4d, 0x53, 0x50, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x41, 0x43,
	0x4c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x71, 0x0a,
	0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x41, 0x43, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x22, 0x55, 0x0a, 0x17, 0x41, 0x43, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x22, 0x69, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x4e, 0x53, 0x53,
	0x41, 0x4e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x44, 0x4e, 0x53, 0x53, 0x41, 0x4e,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x50, 0x53, 0x41, 0x4e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x49, 0x50, 0x53, 0x41, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x54,
	0x4c, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x54,
	0x4c, 0x53, 0x42, 0x8b, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02,
	0x04, 0x48, 0x43, 0x49, 0x43, 0xaa, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xca, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xe2, 0x02, 0x2c, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_pbconfig_config_proto_rawDescOnce sync.Once
	file_private_pbconfig_config_proto_rawDescData = file_private_pbconfig_config_proto_rawDesc
)

func file_private_pbconfig_config_proto_rawDescGZIP() []byte {
	file_private_pbconfig_config_proto_rawDescOnce.Do(func() {
		file_private_pbconfig_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_pbconfig_config_proto_rawDescData)
	})
	return file_private_pbconfig_config_proto_rawDescData
}

var file_private_pbconfig_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_private_pbconfig_config_proto_goTypes = []any{
	(*Config)(nil),                  // 0: hashicorp.consul.internal.config.Config
	(*Gossip)(nil),                  // 1: hashicorp.consul.internal.config.Gossip
	(*GossipEncryption)(nil),        // 2: hashicorp.consul.internal.config.GossipEncryption
	(*TLS)(nil),                     // 3: hashicorp.consul.internal.config.TLS
	(*ACL)(nil),                     // 4: hashicorp.consul.internal.config.ACL
	(*ACLTokens)(nil),               // 5: hashicorp.consul.internal.config.ACLTokens
	(*ACLServiceProviderToken)(nil), // 6: hashicorp.consul.internal.config.ACLServiceProviderToken
	(*AutoEncrypt)(nil),             // 7: hashicorp.consul.internal.config.AutoEncrypt
}
var file_private_pbconfig_config_proto_depIdxs = []int32{
	4, // 0: hashicorp.consul.internal.config.Config.ACL:type_name -> hashicorp.consul.internal.config.ACL
	7, // 1: hashicorp.consul.internal.config.Config.AutoEncrypt:type_name -> hashicorp.consul.internal.config.AutoEncrypt
	1, // 2: hashicorp.consul.internal.config.Config.Gossip:type_name -> hashicorp.consul.internal.config.Gossip
	3, // 3: hashicorp.consul.internal.config.Config.TLS:type_name -> hashicorp.consul.internal.config.TLS
	2, // 4: hashicorp.consul.internal.config.Gossip.Encryption:type_name -> hashicorp.consul.internal.config.GossipEncryption
	5, // 5: hashicorp.consul.internal.config.ACL.Tokens:type_name -> hashicorp.consul.internal.config.ACLTokens
	6, // 6: hashicorp.consul.internal.config.ACLTokens.ManagedServiceProvider:type_name -> hashicorp.consul.internal.config.ACLServiceProviderToken
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_private_pbconfig_config_proto_init() }
func file_private_pbconfig_config_proto_init() {
	if File_private_pbconfig_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_pbconfig_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
		file_private_pbconfig_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Gossip); i {
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
		file_private_pbconfig_config_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GossipEncryption); i {
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
		file_private_pbconfig_config_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TLS); i {
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
		file_private_pbconfig_config_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ACL); i {
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
		file_private_pbconfig_config_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ACLTokens); i {
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
		file_private_pbconfig_config_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ACLServiceProviderToken); i {
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
		file_private_pbconfig_config_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*AutoEncrypt); i {
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
			RawDescriptor: file_private_pbconfig_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_pbconfig_config_proto_goTypes,
		DependencyIndexes: file_private_pbconfig_config_proto_depIdxs,
		MessageInfos:      file_private_pbconfig_config_proto_msgTypes,
	}.Build()
	File_private_pbconfig_config_proto = out.File
	file_private_pbconfig_config_proto_rawDesc = nil
	file_private_pbconfig_config_proto_goTypes = nil
	file_private_pbconfig_config_proto_depIdxs = nil
}
