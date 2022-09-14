// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Database Daemon proxy is used for privileged database ops,
// e.g. bouncing a database and the listeners. It is intended to be used by the
// agents running locally on the same database container.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: oracle/pkg/agents/oracle/dbdaemon_proxy.proto

package oracle

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

type ProxyRunNIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceDbName string   `protobuf:"bytes,1,opt,name=source_db_name,json=sourceDbName,proto3" json:"source_db_name,omitempty"`
	DestDbName   string   `protobuf:"bytes,2,opt,name=dest_db_name,json=destDbName,proto3" json:"dest_db_name,omitempty"`
	Params       []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *ProxyRunNIDRequest) Reset() {
	*x = ProxyRunNIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunNIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunNIDRequest) ProtoMessage() {}

func (x *ProxyRunNIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunNIDRequest.ProtoReflect.Descriptor instead.
func (*ProxyRunNIDRequest) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyRunNIDRequest) GetSourceDbName() string {
	if x != nil {
		return x.SourceDbName
	}
	return ""
}

func (x *ProxyRunNIDRequest) GetDestDbName() string {
	if x != nil {
		return x.DestDbName
	}
	return ""
}

func (x *ProxyRunNIDRequest) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ProxyRunNIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProxyRunNIDResponse) Reset() {
	*x = ProxyRunNIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunNIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunNIDResponse) ProtoMessage() {}

func (x *ProxyRunNIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunNIDResponse.ProtoReflect.Descriptor instead.
func (*ProxyRunNIDResponse) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{1}
}

type ProxyRunDbcaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OracleHome   string   `protobuf:"bytes,1,opt,name=oracle_home,json=oracleHome,proto3" json:"oracle_home,omitempty"`
	DatabaseName string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	Params       []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *ProxyRunDbcaRequest) Reset() {
	*x = ProxyRunDbcaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunDbcaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunDbcaRequest) ProtoMessage() {}

func (x *ProxyRunDbcaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunDbcaRequest.ProtoReflect.Descriptor instead.
func (*ProxyRunDbcaRequest) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyRunDbcaRequest) GetOracleHome() string {
	if x != nil {
		return x.OracleHome
	}
	return ""
}

func (x *ProxyRunDbcaRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *ProxyRunDbcaRequest) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ProxyRunDbcaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProxyRunDbcaResponse) Reset() {
	*x = ProxyRunDbcaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunDbcaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunDbcaResponse) ProtoMessage() {}

func (x *ProxyRunDbcaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunDbcaResponse.ProtoReflect.Descriptor instead.
func (*ProxyRunDbcaResponse) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{3}
}

type ProxyRunInitOracleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *ProxyRunInitOracleRequest) Reset() {
	*x = ProxyRunInitOracleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunInitOracleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunInitOracleRequest) ProtoMessage() {}

func (x *ProxyRunInitOracleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunInitOracleRequest.ProtoReflect.Descriptor instead.
func (*ProxyRunInitOracleRequest) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *ProxyRunInitOracleRequest) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ProxyRunInitOracleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProxyRunInitOracleResponse) Reset() {
	*x = ProxyRunInitOracleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRunInitOracleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRunInitOracleResponse) ProtoMessage() {}

func (x *ProxyRunInitOracleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRunInitOracleResponse.ProtoReflect.Descriptor instead.
func (*ProxyRunInitOracleResponse) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{5}
}

type ProxyFetchServiceImageMetaDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProxyFetchServiceImageMetaDataRequest) Reset() {
	*x = ProxyFetchServiceImageMetaDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyFetchServiceImageMetaDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyFetchServiceImageMetaDataRequest) ProtoMessage() {}

func (x *ProxyFetchServiceImageMetaDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyFetchServiceImageMetaDataRequest.ProtoReflect.Descriptor instead.
func (*ProxyFetchServiceImageMetaDataRequest) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{6}
}

type ProxyFetchServiceImageMetaDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	CdbName     string `protobuf:"bytes,2,opt,name=cdb_name,json=cdbName,proto3" json:"cdb_name,omitempty"`
	OracleHome  string `protobuf:"bytes,3,opt,name=oracle_home,json=oracleHome,proto3" json:"oracle_home,omitempty"`
	SeededImage bool   `protobuf:"varint,4,opt,name=seeded_image,json=seededImage,proto3" json:"seeded_image,omitempty"`
}

func (x *ProxyFetchServiceImageMetaDataResponse) Reset() {
	*x = ProxyFetchServiceImageMetaDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyFetchServiceImageMetaDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyFetchServiceImageMetaDataResponse) ProtoMessage() {}

func (x *ProxyFetchServiceImageMetaDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyFetchServiceImageMetaDataResponse.ProtoReflect.Descriptor instead.
func (*ProxyFetchServiceImageMetaDataResponse) Descriptor() ([]byte, []int) {
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP(), []int{7}
}

func (x *ProxyFetchServiceImageMetaDataResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProxyFetchServiceImageMetaDataResponse) GetCdbName() string {
	if x != nil {
		return x.CdbName
	}
	return ""
}

func (x *ProxyFetchServiceImageMetaDataResponse) GetOracleHome() string {
	if x != nil {
		return x.OracleHome
	}
	return ""
}

func (x *ProxyFetchServiceImageMetaDataResponse) GetSeededImage() bool {
	if x != nil {
		return x.SeededImage
	}
	return false
}

var File_oracle_pkg_agents_oracle_dbdaemon_proxy_proto protoreflect.FileDescriptor

var file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x64, 0x62, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x1a, 0x25,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75,
	0x6e, 0x4e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x74, 0x44, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x4e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x73, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x44, 0x62,
	0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x75, 0x6e, 0x44, 0x62, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x33, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x4f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e,
	0x49, 0x6e, 0x69, 0x74, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x0a, 0x25, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x26,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x65, 0x65, 0x64, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x32,
	0xdc, 0x05, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x44, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x5d, 0x0a, 0x0e, 0x42, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e,
	0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x42,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75,
	0x6e, 0x44, 0x62, 0x63, 0x61, 0x12, 0x22, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x44, 0x62,
	0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x75, 0x6e, 0x44, 0x62, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x4e, 0x49, 0x44, 0x12, 0x21, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x4e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x4e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e,
	0x49, 0x6e, 0x69, 0x74, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x75, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x69,
	0x74, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x8f, 0x01, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x44, 0x6e, 0x66, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x6e, 0x66, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x6e, 0x66, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x58,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x65, 0x6c, 0x63, 0x61, 0x72, 0x72, 0x6f, 0x2d, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x3b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescOnce sync.Once
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescData = file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDesc
)

func file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescGZIP() []byte {
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescOnce.Do(func() {
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescData)
	})
	return file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDescData
}

var file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_goTypes = []interface{}{
	(*ProxyRunNIDRequest)(nil),                     // 0: agents.oracle.ProxyRunNIDRequest
	(*ProxyRunNIDResponse)(nil),                    // 1: agents.oracle.ProxyRunNIDResponse
	(*ProxyRunDbcaRequest)(nil),                    // 2: agents.oracle.ProxyRunDbcaRequest
	(*ProxyRunDbcaResponse)(nil),                   // 3: agents.oracle.ProxyRunDbcaResponse
	(*ProxyRunInitOracleRequest)(nil),              // 4: agents.oracle.ProxyRunInitOracleRequest
	(*ProxyRunInitOracleResponse)(nil),             // 5: agents.oracle.ProxyRunInitOracleResponse
	(*ProxyFetchServiceImageMetaDataRequest)(nil),  // 6: agents.oracle.ProxyFetchServiceImageMetaDataRequest
	(*ProxyFetchServiceImageMetaDataResponse)(nil), // 7: agents.oracle.ProxyFetchServiceImageMetaDataResponse
	(*BounceDatabaseRequest)(nil),                  // 8: agents.oracle.BounceDatabaseRequest
	(*BounceListenerRequest)(nil),                  // 9: agents.oracle.BounceListenerRequest
	(*SetDnfsStateRequest)(nil),                    // 10: agents.oracle.SetDnfsStateRequest
	(*BounceDatabaseResponse)(nil),                 // 11: agents.oracle.BounceDatabaseResponse
	(*BounceListenerResponse)(nil),                 // 12: agents.oracle.BounceListenerResponse
	(*SetDnfsStateResponse)(nil),                   // 13: agents.oracle.SetDnfsStateResponse
}
var file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_depIdxs = []int32{
	8,  // 0: agents.oracle.DatabaseDaemonProxy.BounceDatabase:input_type -> agents.oracle.BounceDatabaseRequest
	9,  // 1: agents.oracle.DatabaseDaemonProxy.BounceListener:input_type -> agents.oracle.BounceListenerRequest
	2,  // 2: agents.oracle.DatabaseDaemonProxy.ProxyRunDbca:input_type -> agents.oracle.ProxyRunDbcaRequest
	0,  // 3: agents.oracle.DatabaseDaemonProxy.ProxyRunNID:input_type -> agents.oracle.ProxyRunNIDRequest
	4,  // 4: agents.oracle.DatabaseDaemonProxy.ProxyRunInitOracle:input_type -> agents.oracle.ProxyRunInitOracleRequest
	6,  // 5: agents.oracle.DatabaseDaemonProxy.ProxyFetchServiceImageMetaData:input_type -> agents.oracle.ProxyFetchServiceImageMetaDataRequest
	10, // 6: agents.oracle.DatabaseDaemonProxy.SetDnfsState:input_type -> agents.oracle.SetDnfsStateRequest
	11, // 7: agents.oracle.DatabaseDaemonProxy.BounceDatabase:output_type -> agents.oracle.BounceDatabaseResponse
	12, // 8: agents.oracle.DatabaseDaemonProxy.BounceListener:output_type -> agents.oracle.BounceListenerResponse
	3,  // 9: agents.oracle.DatabaseDaemonProxy.ProxyRunDbca:output_type -> agents.oracle.ProxyRunDbcaResponse
	1,  // 10: agents.oracle.DatabaseDaemonProxy.ProxyRunNID:output_type -> agents.oracle.ProxyRunNIDResponse
	5,  // 11: agents.oracle.DatabaseDaemonProxy.ProxyRunInitOracle:output_type -> agents.oracle.ProxyRunInitOracleResponse
	7,  // 12: agents.oracle.DatabaseDaemonProxy.ProxyFetchServiceImageMetaData:output_type -> agents.oracle.ProxyFetchServiceImageMetaDataResponse
	13, // 13: agents.oracle.DatabaseDaemonProxy.SetDnfsState:output_type -> agents.oracle.SetDnfsStateResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_init() }
func file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_init() {
	if File_oracle_pkg_agents_oracle_dbdaemon_proxy_proto != nil {
		return
	}
	file_oracle_pkg_agents_oracle_oracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunNIDRequest); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunNIDResponse); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunDbcaRequest); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunDbcaResponse); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunInitOracleRequest); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRunInitOracleResponse); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyFetchServiceImageMetaDataRequest); i {
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
		file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyFetchServiceImageMetaDataResponse); i {
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
			RawDescriptor: file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_goTypes,
		DependencyIndexes: file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_depIdxs,
		MessageInfos:      file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_msgTypes,
	}.Build()
	File_oracle_pkg_agents_oracle_dbdaemon_proxy_proto = out.File
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_rawDesc = nil
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_goTypes = nil
	file_oracle_pkg_agents_oracle_dbdaemon_proxy_proto_depIdxs = nil
}
