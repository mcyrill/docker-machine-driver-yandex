// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/mongodb/v1/database_service.proto

package mongodb

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster that the database belongs to.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MongoDB database to return.
	// To get the name of the database use a [DatabaseService.List] request.
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
}

func (x *GetDatabaseRequest) Reset() {
	*x = GetDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseRequest) ProtoMessage() {}

func (x *GetDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetDatabaseRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type ListDatabasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster to list databases in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListDatabasesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListDatabasesResponse.next_page_token] returned by the previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDatabasesRequest) Reset() {
	*x = ListDatabasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatabasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesRequest) ProtoMessage() {}

func (x *ListDatabasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesRequest.ProtoReflect.Descriptor instead.
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListDatabasesRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListDatabasesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatabasesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListDatabasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of MongoDB databases.
	Databases []*Database `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListDatabasesRequest.page_size], use the [next_page_token] as the value
	// for the [ListDatabasesRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDatabasesResponse) Reset() {
	*x = ListDatabasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatabasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesResponse) ProtoMessage() {}

func (x *ListDatabasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesResponse.ProtoReflect.Descriptor instead.
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDatabasesResponse) GetDatabases() []*Database {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *ListDatabasesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster to create a database in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Configuration of the database to create.
	DatabaseSpec *DatabaseSpec `protobuf:"bytes,2,opt,name=database_spec,json=databaseSpec,proto3" json:"database_spec,omitempty"`
}

func (x *CreateDatabaseRequest) Reset() {
	*x = CreateDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseRequest) ProtoMessage() {}

func (x *CreateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateDatabaseRequest) GetDatabaseSpec() *DatabaseSpec {
	if x != nil {
		return x.DatabaseSpec
	}
	return nil
}

type CreateDatabaseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster where a database is being created.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MongoDB database that is being created.
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
}

func (x *CreateDatabaseMetadata) Reset() {
	*x = CreateDatabaseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseMetadata) ProtoMessage() {}

func (x *CreateDatabaseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseMetadata.ProtoReflect.Descriptor instead.
func (*CreateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDatabaseMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateDatabaseMetadata) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type DeleteDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster to delete a database in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the database to delete.
	// To get the name of the database, use a [DatabaseService.List] request.
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
}

func (x *DeleteDatabaseRequest) Reset() {
	*x = DeleteDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseRequest) ProtoMessage() {}

func (x *DeleteDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteDatabaseRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type DeleteDatabaseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the MongoDB cluster where a database is being deleted.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the MongoDB database that is being deleted.
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
}

func (x *DeleteDatabaseMetadata) Reset() {
	*x = DeleteDatabaseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseMetadata) ProtoMessage() {}

func (x *DeleteDatabaseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseMetadata.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDatabaseMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteDatabaseMetadata) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

var File_yandex_cloud_mdb_mongodb_v1_database_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x33, 0xf2, 0xc7, 0x31,
	0x0e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x52,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01,
	0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30,
	0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9a, 0x01,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x22, 0x5c, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x36, 0x33, 0xf2, 0xc7, 0x31, 0x0e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30,
	0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x32, 0x95, 0x06, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaa, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2f,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x12, 0x43,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0xaa, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73,
	0x12, 0xc5, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x22, 0x33, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x2d, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0xb2, 0xd2, 0x2a, 0x22, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0xdf, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x45, 0x2a, 0x43, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0xb2, 0xd2, 0x2a, 0x2f, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x6a, 0x0a, 0x1f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescData = file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDesc
)

func file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDescData
}

var file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_mdb_mongodb_v1_database_service_proto_goTypes = []interface{}{
	(*GetDatabaseRequest)(nil),     // 0: yandex.cloud.mdb.mongodb.v1.GetDatabaseRequest
	(*ListDatabasesRequest)(nil),   // 1: yandex.cloud.mdb.mongodb.v1.ListDatabasesRequest
	(*ListDatabasesResponse)(nil),  // 2: yandex.cloud.mdb.mongodb.v1.ListDatabasesResponse
	(*CreateDatabaseRequest)(nil),  // 3: yandex.cloud.mdb.mongodb.v1.CreateDatabaseRequest
	(*CreateDatabaseMetadata)(nil), // 4: yandex.cloud.mdb.mongodb.v1.CreateDatabaseMetadata
	(*DeleteDatabaseRequest)(nil),  // 5: yandex.cloud.mdb.mongodb.v1.DeleteDatabaseRequest
	(*DeleteDatabaseMetadata)(nil), // 6: yandex.cloud.mdb.mongodb.v1.DeleteDatabaseMetadata
	(*Database)(nil),               // 7: yandex.cloud.mdb.mongodb.v1.Database
	(*DatabaseSpec)(nil),           // 8: yandex.cloud.mdb.mongodb.v1.DatabaseSpec
	(*operation.Operation)(nil),    // 9: yandex.cloud.operation.Operation
}
var file_yandex_cloud_mdb_mongodb_v1_database_service_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.mdb.mongodb.v1.ListDatabasesResponse.databases:type_name -> yandex.cloud.mdb.mongodb.v1.Database
	8, // 1: yandex.cloud.mdb.mongodb.v1.CreateDatabaseRequest.database_spec:type_name -> yandex.cloud.mdb.mongodb.v1.DatabaseSpec
	0, // 2: yandex.cloud.mdb.mongodb.v1.DatabaseService.Get:input_type -> yandex.cloud.mdb.mongodb.v1.GetDatabaseRequest
	1, // 3: yandex.cloud.mdb.mongodb.v1.DatabaseService.List:input_type -> yandex.cloud.mdb.mongodb.v1.ListDatabasesRequest
	3, // 4: yandex.cloud.mdb.mongodb.v1.DatabaseService.Create:input_type -> yandex.cloud.mdb.mongodb.v1.CreateDatabaseRequest
	5, // 5: yandex.cloud.mdb.mongodb.v1.DatabaseService.Delete:input_type -> yandex.cloud.mdb.mongodb.v1.DeleteDatabaseRequest
	7, // 6: yandex.cloud.mdb.mongodb.v1.DatabaseService.Get:output_type -> yandex.cloud.mdb.mongodb.v1.Database
	2, // 7: yandex.cloud.mdb.mongodb.v1.DatabaseService.List:output_type -> yandex.cloud.mdb.mongodb.v1.ListDatabasesResponse
	9, // 8: yandex.cloud.mdb.mongodb.v1.DatabaseService.Create:output_type -> yandex.cloud.operation.Operation
	9, // 9: yandex.cloud.mdb.mongodb.v1.DatabaseService.Delete:output_type -> yandex.cloud.operation.Operation
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mongodb_v1_database_service_proto_init() }
func file_yandex_cloud_mdb_mongodb_v1_database_service_proto_init() {
	if File_yandex_cloud_mdb_mongodb_v1_database_service_proto != nil {
		return
	}
	file_yandex_cloud_mdb_mongodb_v1_database_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseRequest); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatabasesRequest); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatabasesResponse); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseRequest); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseMetadata); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseRequest); i {
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
		file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseMetadata); i {
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
			RawDescriptor: file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_mdb_mongodb_v1_database_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mongodb_v1_database_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_mongodb_v1_database_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mongodb_v1_database_service_proto = out.File
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_rawDesc = nil
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_goTypes = nil
	file_yandex_cloud_mdb_mongodb_v1_database_service_proto_depIdxs = nil
}
