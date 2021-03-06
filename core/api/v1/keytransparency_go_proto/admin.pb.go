// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/admin.proto

// Key Transparency Administration
//
// The Key Transparency API consists of a map of user names to public
// keys. Each user name also has a history of public keys that have been
// associated with it.

package keytransparency_go_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	trillian "github.com/google/trillian"
	keyspb "github.com/google/trillian/crypto/keyspb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Directory contains information on a single directory
type Directory struct {
	// DirectoryId can be any URL safe string.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,4,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// min_interval is the minimum time between revisions.
	MinInterval *duration.Duration `protobuf:"bytes,5,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	// max_interval is the maximum time between revisions.
	MaxInterval *duration.Duration `protobuf:"bytes,6,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// Deleted indicates whether the directory has been marked as deleted.
	// By its presence in a response, this directory has not been garbage
	// collected.
	Deleted              bool     `protobuf:"varint,7,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Directory) Reset()         { *m = Directory{} }
func (m *Directory) String() string { return proto.CompactTextString(m) }
func (*Directory) ProtoMessage()    {}
func (*Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{0}
}

func (m *Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Directory.Unmarshal(m, b)
}
func (m *Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Directory.Marshal(b, m, deterministic)
}
func (m *Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Directory.Merge(m, src)
}
func (m *Directory) XXX_Size() int {
	return xxx_messageInfo_Directory.Size(m)
}
func (m *Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Directory proto.InternalMessageInfo

func (m *Directory) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *Directory) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Directory) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Directory) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *Directory) GetMinInterval() *duration.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *Directory) GetMaxInterval() *duration.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func (m *Directory) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

// ListDirectories request.
// No pagination options are provided.
type ListDirectoriesRequest struct {
	// showDeleted requests directories that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted          bool     `protobuf:"varint,1,opt,name=show_deleted,json=showDeleted,proto3" json:"show_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDirectoriesRequest) Reset()         { *m = ListDirectoriesRequest{} }
func (m *ListDirectoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDirectoriesRequest) ProtoMessage()    {}
func (*ListDirectoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{1}
}

func (m *ListDirectoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirectoriesRequest.Unmarshal(m, b)
}
func (m *ListDirectoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirectoriesRequest.Marshal(b, m, deterministic)
}
func (m *ListDirectoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirectoriesRequest.Merge(m, src)
}
func (m *ListDirectoriesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDirectoriesRequest.Size(m)
}
func (m *ListDirectoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirectoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirectoriesRequest proto.InternalMessageInfo

func (m *ListDirectoriesRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// ListDirectories response contains directories.
type ListDirectoriesResponse struct {
	Directories          []*Directory `protobuf:"bytes,1,rep,name=directories,proto3" json:"directories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListDirectoriesResponse) Reset()         { *m = ListDirectoriesResponse{} }
func (m *ListDirectoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDirectoriesResponse) ProtoMessage()    {}
func (*ListDirectoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{2}
}

func (m *ListDirectoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirectoriesResponse.Unmarshal(m, b)
}
func (m *ListDirectoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirectoriesResponse.Marshal(b, m, deterministic)
}
func (m *ListDirectoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirectoriesResponse.Merge(m, src)
}
func (m *ListDirectoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDirectoriesResponse.Size(m)
}
func (m *ListDirectoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirectoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirectoriesResponse proto.InternalMessageInfo

func (m *ListDirectoriesResponse) GetDirectories() []*Directory {
	if m != nil {
		return m.Directories
	}
	return nil
}

// GetDirectoryRequest specifies the directory to retrieve information for.
type GetDirectoryRequest struct {
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// showDeleted requests directories that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted          bool     `protobuf:"varint,2,opt,name=show_deleted,json=showDeleted,proto3" json:"show_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDirectoryRequest) Reset()         { *m = GetDirectoryRequest{} }
func (m *GetDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetDirectoryRequest) ProtoMessage()    {}
func (*GetDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{3}
}

func (m *GetDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDirectoryRequest.Unmarshal(m, b)
}
func (m *GetDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *GetDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDirectoryRequest.Merge(m, src)
}
func (m *GetDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetDirectoryRequest.Size(m)
}
func (m *GetDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDirectoryRequest proto.InternalMessageInfo

func (m *GetDirectoryRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *GetDirectoryRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// CreateDirectoryRequest creates a new directory
type CreateDirectoryRequest struct {
	DirectoryId string             `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	MinInterval *duration.Duration `protobuf:"bytes,2,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	MaxInterval *duration.Duration `protobuf:"bytes,3,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// The private_key fields allows callers to set the private key.
	VrfPrivateKey        *any.Any `protobuf:"bytes,4,opt,name=vrf_private_key,json=vrfPrivateKey,proto3" json:"vrf_private_key,omitempty"`
	LogPrivateKey        *any.Any `protobuf:"bytes,5,opt,name=log_private_key,json=logPrivateKey,proto3" json:"log_private_key,omitempty"`
	MapPrivateKey        *any.Any `protobuf:"bytes,6,opt,name=map_private_key,json=mapPrivateKey,proto3" json:"map_private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDirectoryRequest) Reset()         { *m = CreateDirectoryRequest{} }
func (m *CreateDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDirectoryRequest) ProtoMessage()    {}
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{4}
}

func (m *CreateDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDirectoryRequest.Unmarshal(m, b)
}
func (m *CreateDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDirectoryRequest.Merge(m, src)
}
func (m *CreateDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDirectoryRequest.Size(m)
}
func (m *CreateDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDirectoryRequest proto.InternalMessageInfo

func (m *CreateDirectoryRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *CreateDirectoryRequest) GetMinInterval() *duration.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *CreateDirectoryRequest) GetMaxInterval() *duration.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func (m *CreateDirectoryRequest) GetVrfPrivateKey() *any.Any {
	if m != nil {
		return m.VrfPrivateKey
	}
	return nil
}

func (m *CreateDirectoryRequest) GetLogPrivateKey() *any.Any {
	if m != nil {
		return m.LogPrivateKey
	}
	return nil
}

func (m *CreateDirectoryRequest) GetMapPrivateKey() *any.Any {
	if m != nil {
		return m.MapPrivateKey
	}
	return nil
}

// DeleteDirectoryRequest deletes a directory
type DeleteDirectoryRequest struct {
	DirectoryId          string   `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDirectoryRequest) Reset()         { *m = DeleteDirectoryRequest{} }
func (m *DeleteDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDirectoryRequest) ProtoMessage()    {}
func (*DeleteDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{5}
}

func (m *DeleteDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDirectoryRequest.Unmarshal(m, b)
}
func (m *DeleteDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDirectoryRequest.Merge(m, src)
}
func (m *DeleteDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDirectoryRequest.Size(m)
}
func (m *DeleteDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDirectoryRequest proto.InternalMessageInfo

func (m *DeleteDirectoryRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

// UndeleteDirectoryRequest deletes a directory
type UndeleteDirectoryRequest struct {
	DirectoryId          string   `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UndeleteDirectoryRequest) Reset()         { *m = UndeleteDirectoryRequest{} }
func (m *UndeleteDirectoryRequest) String() string { return proto.CompactTextString(m) }
func (*UndeleteDirectoryRequest) ProtoMessage()    {}
func (*UndeleteDirectoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{6}
}

func (m *UndeleteDirectoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndeleteDirectoryRequest.Unmarshal(m, b)
}
func (m *UndeleteDirectoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndeleteDirectoryRequest.Marshal(b, m, deterministic)
}
func (m *UndeleteDirectoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndeleteDirectoryRequest.Merge(m, src)
}
func (m *UndeleteDirectoryRequest) XXX_Size() int {
	return xxx_messageInfo_UndeleteDirectoryRequest.Size(m)
}
func (m *UndeleteDirectoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UndeleteDirectoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UndeleteDirectoryRequest proto.InternalMessageInfo

func (m *UndeleteDirectoryRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

// GarbageCollect request.
type GarbageCollectRequest struct {
	// Soft-deleted directories with a deleted timestamp before this will be fully
	// deleted.
	Before               *timestamp.Timestamp `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GarbageCollectRequest) Reset()         { *m = GarbageCollectRequest{} }
func (m *GarbageCollectRequest) String() string { return proto.CompactTextString(m) }
func (*GarbageCollectRequest) ProtoMessage()    {}
func (*GarbageCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{7}
}

func (m *GarbageCollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarbageCollectRequest.Unmarshal(m, b)
}
func (m *GarbageCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarbageCollectRequest.Marshal(b, m, deterministic)
}
func (m *GarbageCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarbageCollectRequest.Merge(m, src)
}
func (m *GarbageCollectRequest) XXX_Size() int {
	return xxx_messageInfo_GarbageCollectRequest.Size(m)
}
func (m *GarbageCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GarbageCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GarbageCollectRequest proto.InternalMessageInfo

func (m *GarbageCollectRequest) GetBefore() *timestamp.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

type GarbageCollectResponse struct {
	Directories          []*Directory `protobuf:"bytes,1,rep,name=directories,proto3" json:"directories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GarbageCollectResponse) Reset()         { *m = GarbageCollectResponse{} }
func (m *GarbageCollectResponse) String() string { return proto.CompactTextString(m) }
func (*GarbageCollectResponse) ProtoMessage()    {}
func (*GarbageCollectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{8}
}

func (m *GarbageCollectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarbageCollectResponse.Unmarshal(m, b)
}
func (m *GarbageCollectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarbageCollectResponse.Marshal(b, m, deterministic)
}
func (m *GarbageCollectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarbageCollectResponse.Merge(m, src)
}
func (m *GarbageCollectResponse) XXX_Size() int {
	return xxx_messageInfo_GarbageCollectResponse.Size(m)
}
func (m *GarbageCollectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GarbageCollectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GarbageCollectResponse proto.InternalMessageInfo

func (m *GarbageCollectResponse) GetDirectories() []*Directory {
	if m != nil {
		return m.Directories
	}
	return nil
}

func init() {
	proto.RegisterType((*Directory)(nil), "google.keytransparency.v1.Directory")
	proto.RegisterType((*ListDirectoriesRequest)(nil), "google.keytransparency.v1.ListDirectoriesRequest")
	proto.RegisterType((*ListDirectoriesResponse)(nil), "google.keytransparency.v1.ListDirectoriesResponse")
	proto.RegisterType((*GetDirectoryRequest)(nil), "google.keytransparency.v1.GetDirectoryRequest")
	proto.RegisterType((*CreateDirectoryRequest)(nil), "google.keytransparency.v1.CreateDirectoryRequest")
	proto.RegisterType((*DeleteDirectoryRequest)(nil), "google.keytransparency.v1.DeleteDirectoryRequest")
	proto.RegisterType((*UndeleteDirectoryRequest)(nil), "google.keytransparency.v1.UndeleteDirectoryRequest")
	proto.RegisterType((*GarbageCollectRequest)(nil), "google.keytransparency.v1.GarbageCollectRequest")
	proto.RegisterType((*GarbageCollectResponse)(nil), "google.keytransparency.v1.GarbageCollectResponse")
}

func init() { proto.RegisterFile("v1/admin.proto", fileDescriptor_599f1e5eaea78ae3) }

var fileDescriptor_599f1e5eaea78ae3 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xe5, 0xe4, 0x36, 0xbd, 0x9d, 0xf4, 0x26, 0xea, 0xdc, 0x92, 0xba, 0x06, 0x95, 0x60,
	0x10, 0x84, 0x2e, 0x6c, 0x92, 0xee, 0x4a, 0x59, 0x94, 0x96, 0x96, 0x2a, 0x2c, 0x2a, 0xab, 0x6c,
	0x60, 0x11, 0x26, 0xf1, 0xc4, 0x1d, 0xd5, 0xf6, 0x98, 0xf1, 0xc4, 0xd4, 0x42, 0x6c, 0x10, 0x62,
	0xc5, 0xa6, 0xe2, 0x39, 0x78, 0x1a, 0x78, 0x04, 0x5e, 0x80, 0x37, 0x40, 0xb6, 0xc7, 0xf9, 0x70,
	0x52, 0xf7, 0x03, 0x56, 0xd6, 0xf8, 0x9c, 0xdf, 0x99, 0x33, 0xff, 0x73, 0xce, 0xd8, 0xa0, 0x12,
	0x34, 0x75, 0x64, 0x3a, 0xc4, 0xd5, 0x3c, 0x46, 0x39, 0x85, 0xab, 0x16, 0xa5, 0x96, 0x8d, 0xb5,
	0x13, 0x1c, 0x72, 0x86, 0x5c, 0xdf, 0x43, 0x0c, 0xbb, 0xbd, 0x50, 0x0b, 0x9a, 0x8a, 0xd2, 0x63,
	0xa1, 0xc7, 0xa9, 0x7e, 0x82, 0x43, 0xdf, 0xeb, 0x8a, 0x47, 0x82, 0x29, 0xb7, 0x12, 0x4c, 0x47,
	0x1e, 0xd1, 0x91, 0xeb, 0x52, 0x8e, 0x38, 0xa1, 0xae, 0x2f, 0xac, 0x22, 0xa8, 0x1e, 0xaf, 0xba,
	0x83, 0xbe, 0x8e, 0xdc, 0x50, 0x98, 0xd6, 0xb2, 0x26, 0x73, 0xc0, 0x62, 0x56, 0xd8, 0x6f, 0x66,
	0xed, 0xd8, 0xf1, 0x78, 0x0a, 0xdf, 0xce, 0x1a, 0x39, 0x71, 0xb0, 0xcf, 0x91, 0xe3, 0x09, 0x87,
	0x0a, 0x67, 0xc4, 0xb6, 0x09, 0x12, 0xd1, 0xd4, 0x6f, 0x05, 0xb0, 0xb0, 0x4b, 0x18, 0xee, 0x71,
	0xca, 0x42, 0x78, 0x07, 0x2c, 0x9a, 0xe9, 0xa2, 0x43, 0x4c, 0x59, 0xaa, 0x4b, 0x8d, 0x05, 0xa3,
	0x3c, 0x7c, 0x77, 0x60, 0xc2, 0x3a, 0x28, 0xda, 0xd4, 0x92, 0x0b, 0x75, 0xa9, 0x51, 0x6e, 0x55,
	0xb4, 0x61, 0xb8, 0x23, 0x86, 0xb1, 0x11, 0x99, 0x22, 0x0f, 0x07, 0x79, 0x72, 0x71, 0xb6, 0x87,
	0x83, 0x3c, 0x78, 0x17, 0x14, 0x03, 0xd6, 0x97, 0xff, 0x89, 0x3d, 0x96, 0x34, 0xa1, 0xdb, 0xe1,
	0xa0, 0x6b, 0x93, 0x5e, 0x1b, 0x87, 0x46, 0x64, 0x85, 0x5b, 0x60, 0xd1, 0x21, 0x6e, 0x87, 0xb8,
	0x1c, 0xb3, 0x00, 0xd9, 0xf2, 0x5c, 0xec, 0xbd, 0xaa, 0x89, 0x72, 0xa4, 0x27, 0xd4, 0x76, 0x85,
	0x3c, 0x46, 0xd9, 0x21, 0xee, 0x81, 0xf0, 0x8e, 0x69, 0x74, 0x3a, 0xa2, 0x4b, 0x17, 0xd3, 0xe8,
	0x74, 0x48, 0xcb, 0x60, 0xde, 0xc4, 0x36, 0xe6, 0xd8, 0x94, 0xe7, 0xeb, 0x52, 0xe3, 0x5f, 0x23,
	0x5d, 0xaa, 0x8f, 0x41, 0xed, 0x05, 0xf1, 0x79, 0x2a, 0x19, 0xc1, 0xbe, 0x81, 0xdf, 0x0e, 0xb0,
	0xcf, 0x23, 0xed, 0xfc, 0x63, 0xfa, 0xae, 0x93, 0x82, 0x52, 0x0c, 0x96, 0xa3, 0x77, 0xbb, 0x02,
	0x46, 0x60, 0x65, 0x0a, 0xf6, 0x3d, 0xea, 0xfa, 0x18, 0xee, 0x81, 0xa1, 0xca, 0x04, 0xfb, 0xb2,
	0x54, 0x2f, 0x36, 0xca, 0xad, 0x7b, 0xda, 0xb9, 0xbd, 0xa7, 0x0d, 0x8b, 0x66, 0x8c, 0x83, 0xea,
	0x6b, 0xf0, 0xff, 0x3e, 0xe6, 0x23, 0xe3, 0x28, 0xb9, 0x8b, 0x0a, 0x9b, 0xcd, 0xbf, 0x30, 0x9d,
	0xff, 0xaf, 0x02, 0xa8, 0xed, 0x30, 0x8c, 0x38, 0xbe, 0xce, 0x06, 0xd9, 0x82, 0x16, 0xfe, 0xa8,
	0xa0, 0xc5, 0x2b, 0x15, 0x74, 0x0b, 0x54, 0x03, 0xd6, 0xef, 0x78, 0x8c, 0x04, 0x88, 0xe3, 0xce,
	0x09, 0x0e, 0x45, 0xf7, 0x2d, 0x4f, 0x05, 0xd8, 0x76, 0x43, 0xe3, 0xbf, 0x80, 0xf5, 0x0f, 0x13,
	0xdf, 0x36, 0x0e, 0x23, 0xda, 0xa6, 0xd6, 0x04, 0x3d, 0x97, 0x47, 0xdb, 0xd4, 0x9a, 0xa4, 0x1d,
	0xe4, 0x4d, 0xd0, 0xa5, 0x3c, 0xda, 0x41, 0xde, 0x88, 0x8e, 0x1a, 0x2e, 0x91, 0xff, 0x1a, 0x92,
	0xab, 0x4f, 0x80, 0xfc, 0xd2, 0x35, 0xaf, 0x8d, 0xb7, 0xc1, 0x8d, 0x7d, 0xc4, 0xba, 0xc8, 0xc2,
	0x3b, 0xd4, 0xb6, 0x71, 0x8f, 0xa7, 0x6c, 0x0b, 0x94, 0xba, 0xb8, 0x4f, 0x19, 0x8e, 0xa9, 0x72,
	0x4b, 0x99, 0x3a, 0xc9, 0x51, 0x7a, 0xef, 0x18, 0xc2, 0x53, 0x7d, 0x03, 0x6a, 0xd9, 0x60, 0x7f,
	0xb7, 0xf7, 0x5b, 0x3f, 0x4a, 0x60, 0xb9, 0x8d, 0xc3, 0xa3, 0x31, 0xef, 0xed, 0xe8, 0x22, 0x87,
	0x67, 0x12, 0xa8, 0x66, 0x06, 0x0f, 0x36, 0x73, 0xe2, 0xcf, 0x9e, 0x70, 0xa5, 0x75, 0x15, 0x24,
	0x39, 0x9b, 0xba, 0xf2, 0xf1, 0xfb, 0xcf, 0xaf, 0x85, 0x25, 0x58, 0xd5, 0x83, 0xa6, 0x3e, 0x96,
	0x2c, 0xfc, 0x22, 0x81, 0xc5, 0xf1, 0x49, 0x85, 0x5a, 0x4e, 0xf4, 0x19, 0x23, 0xad, 0x5c, 0x4a,
	0x20, 0xf5, 0x7e, 0xbc, 0x7f, 0x1d, 0xae, 0x65, 0xf6, 0xd7, 0xdf, 0x8f, 0x17, 0xff, 0x03, 0xfc,
	0x2c, 0x81, 0x6a, 0x66, 0xb4, 0x73, 0x25, 0x9a, 0x7d, 0x0d, 0x5c, 0x32, 0x29, 0x25, 0x4e, 0x6a,
	0x59, 0xcd, 0x8a, 0xb2, 0x29, 0xad, 0xc3, 0x4f, 0x12, 0xa8, 0x66, 0x1a, 0x3e, 0x37, 0x91, 0xd9,
	0xc3, 0xa1, 0xd4, 0xa6, 0x3a, 0xf2, 0x59, 0xf4, 0x99, 0x4c, 0xf5, 0x58, 0xbf, 0x48, 0x8f, 0x33,
	0x09, 0x2c, 0x4d, 0x8d, 0x0e, 0xdc, 0xc8, 0x49, 0xe4, 0xbc, 0x41, 0x3b, 0x37, 0x15, 0x3d, 0x4e,
	0xe5, 0xe1, 0xfa, 0x83, 0xfc, 0x54, 0x36, 0x07, 0x22, 0x30, 0x1c, 0x80, 0xca, 0xe4, 0x04, 0xc1,
	0x47, 0x79, 0x3d, 0x33, 0x6b, 0x72, 0x95, 0xe6, 0x15, 0x88, 0xa4, 0x85, 0x9f, 0x3e, 0x7f, 0xb5,
	0x67, 0x11, 0x7e, 0x3c, 0xe8, 0x6a, 0x3d, 0xea, 0xe8, 0xe2, 0x07, 0x23, 0x83, 0xeb, 0x3d, 0xca,
	0x92, 0x7f, 0x9d, 0xa0, 0x99, 0xb5, 0x75, 0x2c, 0xda, 0x49, 0x8e, 0x5e, 0x8a, 0x1f, 0x1b, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xd9, 0x2e, 0x51, 0x63, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyTransparencyAdminClient is the client API for KeyTransparencyAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyTransparencyAdminClient interface {
	// ListDirectories returns a list of all directories this Key Transparency
	// server operates on.
	ListDirectories(ctx context.Context, in *ListDirectoriesRequest, opts ...grpc.CallOption) (*ListDirectoriesResponse, error)
	// GetDirectory returns the confiuration information for a given directory.
	GetDirectory(ctx context.Context, in *GetDirectoryRequest, opts ...grpc.CallOption) (*Directory, error)
	// CreateDirectory creates a new Trillian log/map pair.  A unique directoryId
	// must be provided.  To create a new directory with the same name as a
	// previously deleted directory, a user must wait X days until the directory
	// is garbage collected.
	CreateDirectory(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error)
	// DeleteDirectory marks a directory as deleted.  Directories will be garbage
	// collected after X days.
	DeleteDirectory(ctx context.Context, in *DeleteDirectoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// UndeleteDirectory marks a previously deleted directory as active if it has
	// not already been garbage collected.
	UndeleteDirectory(ctx context.Context, in *UndeleteDirectoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Fully delete soft-deleted directories that have been soft-deleted before
	// the specified timestamp.
	GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error)
}

type keyTransparencyAdminClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyAdminClient(cc *grpc.ClientConn) KeyTransparencyAdminClient {
	return &keyTransparencyAdminClient{cc}
}

func (c *keyTransparencyAdminClient) ListDirectories(ctx context.Context, in *ListDirectoriesRequest, opts ...grpc.CallOption) (*ListDirectoriesResponse, error) {
	out := new(ListDirectoriesResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/ListDirectories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) GetDirectory(ctx context.Context, in *GetDirectoryRequest, opts ...grpc.CallOption) (*Directory, error) {
	out := new(Directory)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/GetDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) CreateDirectory(ctx context.Context, in *CreateDirectoryRequest, opts ...grpc.CallOption) (*Directory, error) {
	out := new(Directory)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/CreateDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) DeleteDirectory(ctx context.Context, in *DeleteDirectoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/DeleteDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) UndeleteDirectory(ctx context.Context, in *UndeleteDirectoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/UndeleteDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error) {
	out := new(GarbageCollectResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/GarbageCollect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyTransparencyAdminServer is the server API for KeyTransparencyAdmin service.
type KeyTransparencyAdminServer interface {
	// ListDirectories returns a list of all directories this Key Transparency
	// server operates on.
	ListDirectories(context.Context, *ListDirectoriesRequest) (*ListDirectoriesResponse, error)
	// GetDirectory returns the confiuration information for a given directory.
	GetDirectory(context.Context, *GetDirectoryRequest) (*Directory, error)
	// CreateDirectory creates a new Trillian log/map pair.  A unique directoryId
	// must be provided.  To create a new directory with the same name as a
	// previously deleted directory, a user must wait X days until the directory
	// is garbage collected.
	CreateDirectory(context.Context, *CreateDirectoryRequest) (*Directory, error)
	// DeleteDirectory marks a directory as deleted.  Directories will be garbage
	// collected after X days.
	DeleteDirectory(context.Context, *DeleteDirectoryRequest) (*empty.Empty, error)
	// UndeleteDirectory marks a previously deleted directory as active if it has
	// not already been garbage collected.
	UndeleteDirectory(context.Context, *UndeleteDirectoryRequest) (*empty.Empty, error)
	// Fully delete soft-deleted directories that have been soft-deleted before
	// the specified timestamp.
	GarbageCollect(context.Context, *GarbageCollectRequest) (*GarbageCollectResponse, error)
}

func RegisterKeyTransparencyAdminServer(s *grpc.Server, srv KeyTransparencyAdminServer) {
	s.RegisterService(&_KeyTransparencyAdmin_serviceDesc, srv)
}

func _KeyTransparencyAdmin_ListDirectories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDirectoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).ListDirectories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/ListDirectories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).ListDirectories(ctx, req.(*ListDirectoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_GetDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).GetDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/GetDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).GetDirectory(ctx, req.(*GetDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_CreateDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).CreateDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/CreateDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).CreateDirectory(ctx, req.(*CreateDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_DeleteDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).DeleteDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/DeleteDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).DeleteDirectory(ctx, req.(*DeleteDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_UndeleteDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).UndeleteDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/UndeleteDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).UndeleteDirectory(ctx, req.(*UndeleteDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_GarbageCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarbageCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).GarbageCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/GarbageCollect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).GarbageCollect(ctx, req.(*GarbageCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.v1.KeyTransparencyAdmin",
	HandlerType: (*KeyTransparencyAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDirectories",
			Handler:    _KeyTransparencyAdmin_ListDirectories_Handler,
		},
		{
			MethodName: "GetDirectory",
			Handler:    _KeyTransparencyAdmin_GetDirectory_Handler,
		},
		{
			MethodName: "CreateDirectory",
			Handler:    _KeyTransparencyAdmin_CreateDirectory_Handler,
		},
		{
			MethodName: "DeleteDirectory",
			Handler:    _KeyTransparencyAdmin_DeleteDirectory_Handler,
		},
		{
			MethodName: "UndeleteDirectory",
			Handler:    _KeyTransparencyAdmin_UndeleteDirectory_Handler,
		},
		{
			MethodName: "GarbageCollect",
			Handler:    _KeyTransparencyAdmin_GarbageCollect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/admin.proto",
}
