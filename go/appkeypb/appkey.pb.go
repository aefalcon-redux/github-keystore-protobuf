// Code generated by protoc-gen-go. DO NOT EDIT.
// source: appkey.proto

package appkeypb // import "github.com/aefalcon/github-keystore-protobuf/go/appkeypb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import locationpb "github.com/aefalcon/github-keystore-protobuf/go/locationpb"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type App struct {
	Id                   uint64                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keys                 map[string]*AppKeyIndexEntry `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{0}
}
func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (dst *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(dst, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetKeys() map[string]*AppKeyIndexEntry {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AppIndexEntry struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppIndexEntry) Reset()         { *m = AppIndexEntry{} }
func (m *AppIndexEntry) String() string { return proto.CompactTextString(m) }
func (*AppIndexEntry) ProtoMessage()    {}
func (*AppIndexEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{1}
}
func (m *AppIndexEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppIndexEntry.Unmarshal(m, b)
}
func (m *AppIndexEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppIndexEntry.Marshal(b, m, deterministic)
}
func (dst *AppIndexEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppIndexEntry.Merge(dst, src)
}
func (m *AppIndexEntry) XXX_Size() int {
	return xxx_messageInfo_AppIndexEntry.Size(m)
}
func (m *AppIndexEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AppIndexEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AppIndexEntry proto.InternalMessageInfo

func (m *AppIndexEntry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AppIndex struct {
	AppRefs              map[uint64]*AppIndexEntry `protobuf:"bytes,1,rep,name=app_refs,json=appRefs,proto3" json:"app_refs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AppIndex) Reset()         { *m = AppIndex{} }
func (m *AppIndex) String() string { return proto.CompactTextString(m) }
func (*AppIndex) ProtoMessage()    {}
func (*AppIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{2}
}
func (m *AppIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppIndex.Unmarshal(m, b)
}
func (m *AppIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppIndex.Marshal(b, m, deterministic)
}
func (dst *AppIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppIndex.Merge(dst, src)
}
func (m *AppIndex) XXX_Size() int {
	return xxx_messageInfo_AppIndex.Size(m)
}
func (m *AppIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_AppIndex.DiscardUnknown(m)
}

var xxx_messageInfo_AppIndex proto.InternalMessageInfo

func (m *AppIndex) GetAppRefs() map[uint64]*AppIndexEntry {
	if m != nil {
		return m.AppRefs
	}
	return nil
}

type AppKeyMeta struct {
	Fingerprint          string   `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	App                  uint64   `protobuf:"varint,2,opt,name=app,proto3" json:"app,omitempty"`
	Disabled             bool     `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
	NotBefore            string   `protobuf:"bytes,4,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NoatAfter            string   `protobuf:"bytes,5,opt,name=noat_after,json=noatAfter,proto3" json:"noat_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppKeyMeta) Reset()         { *m = AppKeyMeta{} }
func (m *AppKeyMeta) String() string { return proto.CompactTextString(m) }
func (*AppKeyMeta) ProtoMessage()    {}
func (*AppKeyMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{3}
}
func (m *AppKeyMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppKeyMeta.Unmarshal(m, b)
}
func (m *AppKeyMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppKeyMeta.Marshal(b, m, deterministic)
}
func (dst *AppKeyMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppKeyMeta.Merge(dst, src)
}
func (m *AppKeyMeta) XXX_Size() int {
	return xxx_messageInfo_AppKeyMeta.Size(m)
}
func (m *AppKeyMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_AppKeyMeta.DiscardUnknown(m)
}

var xxx_messageInfo_AppKeyMeta proto.InternalMessageInfo

func (m *AppKeyMeta) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *AppKeyMeta) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *AppKeyMeta) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *AppKeyMeta) GetNotBefore() string {
	if m != nil {
		return m.NotBefore
	}
	return ""
}

func (m *AppKeyMeta) GetNoatAfter() string {
	if m != nil {
		return m.NoatAfter
	}
	return ""
}

type AppKey struct {
	Meta                 *AppKeyMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Key                  []byte      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppKey) Reset()         { *m = AppKey{} }
func (m *AppKey) String() string { return proto.CompactTextString(m) }
func (*AppKey) ProtoMessage()    {}
func (*AppKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{4}
}
func (m *AppKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppKey.Unmarshal(m, b)
}
func (m *AppKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppKey.Marshal(b, m, deterministic)
}
func (dst *AppKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppKey.Merge(dst, src)
}
func (m *AppKey) XXX_Size() int {
	return xxx_messageInfo_AppKey.Size(m)
}
func (m *AppKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AppKey.DiscardUnknown(m)
}

var xxx_messageInfo_AppKey proto.InternalMessageInfo

func (m *AppKey) GetMeta() *AppKeyMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AppKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type AppKeyIndexEntry struct {
	Meta                 *AppKeyMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppKeyIndexEntry) Reset()         { *m = AppKeyIndexEntry{} }
func (m *AppKeyIndexEntry) String() string { return proto.CompactTextString(m) }
func (*AppKeyIndexEntry) ProtoMessage()    {}
func (*AppKeyIndexEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{5}
}
func (m *AppKeyIndexEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppKeyIndexEntry.Unmarshal(m, b)
}
func (m *AppKeyIndexEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppKeyIndexEntry.Marshal(b, m, deterministic)
}
func (dst *AppKeyIndexEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppKeyIndexEntry.Merge(dst, src)
}
func (m *AppKeyIndexEntry) XXX_Size() int {
	return xxx_messageInfo_AppKeyIndexEntry.Size(m)
}
func (m *AppKeyIndexEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AppKeyIndexEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AppKeyIndexEntry proto.InternalMessageInfo

func (m *AppKeyIndexEntry) GetMeta() *AppKeyMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type AppKeyManagerConfig struct {
	DbLoc                *locationpb.Location `protobuf:"bytes,1,opt,name=db_loc,json=dbLoc,proto3" json:"db_loc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AppKeyManagerConfig) Reset()         { *m = AppKeyManagerConfig{} }
func (m *AppKeyManagerConfig) String() string { return proto.CompactTextString(m) }
func (*AppKeyManagerConfig) ProtoMessage()    {}
func (*AppKeyManagerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{6}
}
func (m *AppKeyManagerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppKeyManagerConfig.Unmarshal(m, b)
}
func (m *AppKeyManagerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppKeyManagerConfig.Marshal(b, m, deterministic)
}
func (dst *AppKeyManagerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppKeyManagerConfig.Merge(dst, src)
}
func (m *AppKeyManagerConfig) XXX_Size() int {
	return xxx_messageInfo_AppKeyManagerConfig.Size(m)
}
func (m *AppKeyManagerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AppKeyManagerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AppKeyManagerConfig proto.InternalMessageInfo

func (m *AppKeyManagerConfig) GetDbLoc() *locationpb.Location {
	if m != nil {
		return m.DbLoc
	}
	return nil
}

type Links struct {
	AppIndex             string   `protobuf:"bytes,1,opt,name=app_index,json=appIndex,proto3" json:"app_index,omitempty"`
	App                  string   `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	KeyMeta              string   `protobuf:"bytes,4,opt,name=key_meta,json=keyMeta,proto3" json:"key_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Links) Reset()         { *m = Links{} }
func (m *Links) String() string { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()    {}
func (*Links) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{7}
}
func (m *Links) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Links.Unmarshal(m, b)
}
func (m *Links) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Links.Marshal(b, m, deterministic)
}
func (dst *Links) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Links.Merge(dst, src)
}
func (m *Links) XXX_Size() int {
	return xxx_messageInfo_Links.Size(m)
}
func (m *Links) XXX_DiscardUnknown() {
	xxx_messageInfo_Links.DiscardUnknown(m)
}

var xxx_messageInfo_Links proto.InternalMessageInfo

func (m *Links) GetAppIndex() string {
	if m != nil {
		return m.AppIndex
	}
	return ""
}

func (m *Links) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Links) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Links) GetKeyMeta() string {
	if m != nil {
		return m.KeyMeta
	}
	return ""
}

type AddAppRequest struct {
	App                  uint64    `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Keys                 []*AppKey `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddAppRequest) Reset()         { *m = AddAppRequest{} }
func (m *AddAppRequest) String() string { return proto.CompactTextString(m) }
func (*AddAppRequest) ProtoMessage()    {}
func (*AddAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{8}
}
func (m *AddAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAppRequest.Unmarshal(m, b)
}
func (m *AddAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAppRequest.Marshal(b, m, deterministic)
}
func (dst *AddAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAppRequest.Merge(dst, src)
}
func (m *AddAppRequest) XXX_Size() int {
	return xxx_messageInfo_AddAppRequest.Size(m)
}
func (m *AddAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAppRequest proto.InternalMessageInfo

func (m *AddAppRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *AddAppRequest) GetKeys() []*AppKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AddAppResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAppResponse) Reset()         { *m = AddAppResponse{} }
func (m *AddAppResponse) String() string { return proto.CompactTextString(m) }
func (*AddAppResponse) ProtoMessage()    {}
func (*AddAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{9}
}
func (m *AddAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAppResponse.Unmarshal(m, b)
}
func (m *AddAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAppResponse.Marshal(b, m, deterministic)
}
func (dst *AddAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAppResponse.Merge(dst, src)
}
func (m *AddAppResponse) XXX_Size() int {
	return xxx_messageInfo_AddAppResponse.Size(m)
}
func (m *AddAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAppResponse proto.InternalMessageInfo

type RemoveAppRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAppRequest) Reset()         { *m = RemoveAppRequest{} }
func (m *RemoveAppRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveAppRequest) ProtoMessage()    {}
func (*RemoveAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{10}
}
func (m *RemoveAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAppRequest.Unmarshal(m, b)
}
func (m *RemoveAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAppRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAppRequest.Merge(dst, src)
}
func (m *RemoveAppRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveAppRequest.Size(m)
}
func (m *RemoveAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAppRequest proto.InternalMessageInfo

func (m *RemoveAppRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type RemoveAppResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAppResponse) Reset()         { *m = RemoveAppResponse{} }
func (m *RemoveAppResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveAppResponse) ProtoMessage()    {}
func (*RemoveAppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{11}
}
func (m *RemoveAppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAppResponse.Unmarshal(m, b)
}
func (m *RemoveAppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAppResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveAppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAppResponse.Merge(dst, src)
}
func (m *RemoveAppResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveAppResponse.Size(m)
}
func (m *RemoveAppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAppResponse proto.InternalMessageInfo

type GetAppRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppRequest) Reset()         { *m = GetAppRequest{} }
func (m *GetAppRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppRequest) ProtoMessage()    {}
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{12}
}
func (m *GetAppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppRequest.Unmarshal(m, b)
}
func (m *GetAppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppRequest.Marshal(b, m, deterministic)
}
func (dst *GetAppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppRequest.Merge(dst, src)
}
func (m *GetAppRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppRequest.Size(m)
}
func (m *GetAppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppRequest proto.InternalMessageInfo

func (m *GetAppRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type ListAppsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAppsRequest) Reset()         { *m = ListAppsRequest{} }
func (m *ListAppsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAppsRequest) ProtoMessage()    {}
func (*ListAppsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{13}
}
func (m *ListAppsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAppsRequest.Unmarshal(m, b)
}
func (m *ListAppsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAppsRequest.Marshal(b, m, deterministic)
}
func (dst *ListAppsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAppsRequest.Merge(dst, src)
}
func (m *ListAppsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAppsRequest.Size(m)
}
func (m *ListAppsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAppsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAppsRequest proto.InternalMessageInfo

type AddKeyRequest struct {
	App                  uint64    `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Keys                 []*AppKey `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddKeyRequest) Reset()         { *m = AddKeyRequest{} }
func (m *AddKeyRequest) String() string { return proto.CompactTextString(m) }
func (*AddKeyRequest) ProtoMessage()    {}
func (*AddKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{14}
}
func (m *AddKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKeyRequest.Unmarshal(m, b)
}
func (m *AddKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKeyRequest.Marshal(b, m, deterministic)
}
func (dst *AddKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKeyRequest.Merge(dst, src)
}
func (m *AddKeyRequest) XXX_Size() int {
	return xxx_messageInfo_AddKeyRequest.Size(m)
}
func (m *AddKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddKeyRequest proto.InternalMessageInfo

func (m *AddKeyRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *AddKeyRequest) GetKeys() []*AppKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AddKeyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddKeyResponse) Reset()         { *m = AddKeyResponse{} }
func (m *AddKeyResponse) String() string { return proto.CompactTextString(m) }
func (*AddKeyResponse) ProtoMessage()    {}
func (*AddKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{15}
}
func (m *AddKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKeyResponse.Unmarshal(m, b)
}
func (m *AddKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKeyResponse.Marshal(b, m, deterministic)
}
func (dst *AddKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKeyResponse.Merge(dst, src)
}
func (m *AddKeyResponse) XXX_Size() int {
	return xxx_messageInfo_AddKeyResponse.Size(m)
}
func (m *AddKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddKeyResponse proto.InternalMessageInfo

type RemoveKeyRequest struct {
	App                  uint64   `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Fingerprints         []string `protobuf:"bytes,2,rep,name=fingerprints,proto3" json:"fingerprints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveKeyRequest) Reset()         { *m = RemoveKeyRequest{} }
func (m *RemoveKeyRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveKeyRequest) ProtoMessage()    {}
func (*RemoveKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{16}
}
func (m *RemoveKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveKeyRequest.Unmarshal(m, b)
}
func (m *RemoveKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveKeyRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveKeyRequest.Merge(dst, src)
}
func (m *RemoveKeyRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveKeyRequest.Size(m)
}
func (m *RemoveKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveKeyRequest proto.InternalMessageInfo

func (m *RemoveKeyRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *RemoveKeyRequest) GetFingerprints() []string {
	if m != nil {
		return m.Fingerprints
	}
	return nil
}

type RemoveKeyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveKeyResponse) Reset()         { *m = RemoveKeyResponse{} }
func (m *RemoveKeyResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveKeyResponse) ProtoMessage()    {}
func (*RemoveKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{17}
}
func (m *RemoveKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveKeyResponse.Unmarshal(m, b)
}
func (m *RemoveKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveKeyResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveKeyResponse.Merge(dst, src)
}
func (m *RemoveKeyResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveKeyResponse.Size(m)
}
func (m *RemoveKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveKeyResponse proto.InternalMessageInfo

type SignJwtRequest struct {
	App                  uint64          `protobuf:"varint,1,opt,name=app,proto3" json:"app,omitempty"`
	Claims               *_struct.Struct `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
	Algorithm            string          `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SignJwtRequest) Reset()         { *m = SignJwtRequest{} }
func (m *SignJwtRequest) String() string { return proto.CompactTextString(m) }
func (*SignJwtRequest) ProtoMessage()    {}
func (*SignJwtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{18}
}
func (m *SignJwtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignJwtRequest.Unmarshal(m, b)
}
func (m *SignJwtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignJwtRequest.Marshal(b, m, deterministic)
}
func (dst *SignJwtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignJwtRequest.Merge(dst, src)
}
func (m *SignJwtRequest) XXX_Size() int {
	return xxx_messageInfo_SignJwtRequest.Size(m)
}
func (m *SignJwtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignJwtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignJwtRequest proto.InternalMessageInfo

func (m *SignJwtRequest) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *SignJwtRequest) GetClaims() *_struct.Struct {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *SignJwtRequest) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

type SignJwtResponse struct {
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignJwtResponse) Reset()         { *m = SignJwtResponse{} }
func (m *SignJwtResponse) String() string { return proto.CompactTextString(m) }
func (*SignJwtResponse) ProtoMessage()    {}
func (*SignJwtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_appkey_d8ed3c84b84f5c73, []int{19}
}
func (m *SignJwtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignJwtResponse.Unmarshal(m, b)
}
func (m *SignJwtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignJwtResponse.Marshal(b, m, deterministic)
}
func (dst *SignJwtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignJwtResponse.Merge(dst, src)
}
func (m *SignJwtResponse) XXX_Size() int {
	return xxx_messageInfo_SignJwtResponse.Size(m)
}
func (m *SignJwtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignJwtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignJwtResponse proto.InternalMessageInfo

func (m *SignJwtResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func init() {
	proto.RegisterType((*App)(nil), "mobettersoftware.protobuf.appkey.App")
	proto.RegisterMapType((map[string]*AppKeyIndexEntry)(nil), "mobettersoftware.protobuf.appkey.App.KeysEntry")
	proto.RegisterType((*AppIndexEntry)(nil), "mobettersoftware.protobuf.appkey.AppIndexEntry")
	proto.RegisterType((*AppIndex)(nil), "mobettersoftware.protobuf.appkey.AppIndex")
	proto.RegisterMapType((map[uint64]*AppIndexEntry)(nil), "mobettersoftware.protobuf.appkey.AppIndex.AppRefsEntry")
	proto.RegisterType((*AppKeyMeta)(nil), "mobettersoftware.protobuf.appkey.AppKeyMeta")
	proto.RegisterType((*AppKey)(nil), "mobettersoftware.protobuf.appkey.AppKey")
	proto.RegisterType((*AppKeyIndexEntry)(nil), "mobettersoftware.protobuf.appkey.AppKeyIndexEntry")
	proto.RegisterType((*AppKeyManagerConfig)(nil), "mobettersoftware.protobuf.appkey.AppKeyManagerConfig")
	proto.RegisterType((*Links)(nil), "mobettersoftware.protobuf.appkey.Links")
	proto.RegisterType((*AddAppRequest)(nil), "mobettersoftware.protobuf.appkey.AddAppRequest")
	proto.RegisterType((*AddAppResponse)(nil), "mobettersoftware.protobuf.appkey.AddAppResponse")
	proto.RegisterType((*RemoveAppRequest)(nil), "mobettersoftware.protobuf.appkey.RemoveAppRequest")
	proto.RegisterType((*RemoveAppResponse)(nil), "mobettersoftware.protobuf.appkey.RemoveAppResponse")
	proto.RegisterType((*GetAppRequest)(nil), "mobettersoftware.protobuf.appkey.GetAppRequest")
	proto.RegisterType((*ListAppsRequest)(nil), "mobettersoftware.protobuf.appkey.ListAppsRequest")
	proto.RegisterType((*AddKeyRequest)(nil), "mobettersoftware.protobuf.appkey.AddKeyRequest")
	proto.RegisterType((*AddKeyResponse)(nil), "mobettersoftware.protobuf.appkey.AddKeyResponse")
	proto.RegisterType((*RemoveKeyRequest)(nil), "mobettersoftware.protobuf.appkey.RemoveKeyRequest")
	proto.RegisterType((*RemoveKeyResponse)(nil), "mobettersoftware.protobuf.appkey.RemoveKeyResponse")
	proto.RegisterType((*SignJwtRequest)(nil), "mobettersoftware.protobuf.appkey.SignJwtRequest")
	proto.RegisterType((*SignJwtResponse)(nil), "mobettersoftware.protobuf.appkey.SignJwtResponse")
}

func init() { proto.RegisterFile("appkey.proto", fileDescriptor_appkey_d8ed3c84b84f5c73) }

var fileDescriptor_appkey_d8ed3c84b84f5c73 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x9b, 0xff, 0xd3, 0x34, 0x4d, 0xb7, 0x87, 0x5f, 0x7e, 0xa6, 0x88, 0x60, 0x40, 0xaa,
	0x10, 0x8d, 0x69, 0x7a, 0xa0, 0xaa, 0x38, 0xd0, 0x56, 0x15, 0x05, 0xca, 0xc5, 0xe5, 0x54, 0x21,
	0x05, 0x3b, 0x5e, 0xbb, 0x26, 0x89, 0xd7, 0xb5, 0x37, 0x2d, 0xbd, 0xf1, 0x1c, 0xdc, 0x79, 0x16,
	0x0e, 0xbc, 0x14, 0xbb, 0xde, 0x75, 0xec, 0x04, 0xd1, 0x6e, 0x85, 0x38, 0x65, 0x3d, 0x3b, 0x33,
	0xdf, 0x7c, 0x3b, 0xdf, 0xec, 0x06, 0x9a, 0x76, 0x14, 0x8d, 0xf0, 0x75, 0x2f, 0x8a, 0x09, 0x25,
	0xa8, 0x3b, 0x21, 0x0e, 0xa6, 0x14, 0xc7, 0x09, 0xf1, 0xe8, 0x95, 0x1d, 0x63, 0x61, 0x77, 0xa6,
	0x5e, 0x4f, 0xf8, 0xe9, 0x1b, 0x3e, 0x21, 0xfe, 0x18, 0x9b, 0x99, 0xdd, 0x4c, 0x68, 0x3c, 0x1d,
	0x52, 0xe1, 0xa7, 0xb7, 0xc6, 0x64, 0x68, 0xd3, 0x80, 0x84, 0xe2, 0xdb, 0xf8, 0xa1, 0x41, 0x69,
	0x3f, 0x8a, 0x50, 0x0b, 0x96, 0x02, 0xb7, 0xa3, 0x75, 0xb5, 0xcd, 0xb2, 0xc5, 0x56, 0xe8, 0x10,
	0xca, 0x2c, 0x59, 0xd2, 0x59, 0xea, 0x96, 0x36, 0x97, 0xfb, 0x66, 0xef, 0x36, 0xd8, 0x1e, 0x4b,
	0xd2, 0x7b, 0xc7, 0x22, 0x8e, 0x42, 0x1a, 0x5f, 0x5b, 0x69, 0xb0, 0x3e, 0x82, 0xc6, 0xcc, 0x84,
	0xda, 0x50, 0x62, 0xc6, 0x14, 0xa2, 0x61, 0xf1, 0x25, 0x3a, 0x86, 0xca, 0xa5, 0x3d, 0x9e, 0x62,
	0x06, 0xa2, 0x31, 0x90, 0xbe, 0x12, 0x08, 0x4b, 0xf8, 0x26, 0x74, 0xf1, 0x17, 0x81, 0x23, 0x12,
	0xec, 0x2d, 0xed, 0x6a, 0xc6, 0x03, 0x58, 0x61, 0xdb, 0xf9, 0xde, 0x22, 0x25, 0xe3, 0xa7, 0x06,
	0xf5, 0xcc, 0x03, 0x59, 0x50, 0x67, 0x39, 0x07, 0x31, 0xf6, 0x12, 0xe6, 0xc2, 0x39, 0xbe, 0x50,
	0x82, 0x4f, 0xa3, 0xf9, 0xc2, 0x62, 0x91, 0xa2, 0x86, 0x9a, 0x2d, 0xbe, 0x18, 0xdd, 0x66, 0x71,
	0xa3, 0xc8, 0xb8, 0x2c, 0x18, 0x1f, 0xcd, 0x33, 0x36, 0xd5, 0x21, 0x7f, 0xa3, 0xfb, 0x4d, 0x03,
	0x10, 0xc7, 0xf1, 0x1e, 0x53, 0x1b, 0x75, 0x61, 0xd9, 0x0b, 0x42, 0x1f, 0xc7, 0x51, 0x1c, 0x84,
	0x54, 0x9e, 0x72, 0xd1, 0xc4, 0xab, 0x61, 0x39, 0x53, 0x64, 0x56, 0x0d, 0x5b, 0x22, 0x1d, 0xea,
	0x6e, 0x90, 0xd8, 0xce, 0x18, 0xbb, 0x9d, 0x12, 0x33, 0xd7, 0xad, 0xd9, 0x37, 0xba, 0x0f, 0x10,
	0x12, 0x3a, 0x70, 0xb0, 0x47, 0x62, 0xdc, 0x29, 0xa7, 0xe9, 0x1a, 0xcc, 0x72, 0x90, 0x1a, 0xc4,
	0xb6, 0x4d, 0x07, 0xb6, 0xc7, 0x8a, 0xef, 0x54, 0xb2, 0x6d, 0x9b, 0xee, 0x73, 0x83, 0xf1, 0x11,
	0xaa, 0xa2, 0x36, 0xf4, 0x0a, 0xca, 0x13, 0x56, 0x5f, 0x5a, 0xd0, 0x72, 0xff, 0x99, 0x6a, 0x8b,
	0x39, 0x27, 0x2b, 0x8d, 0xcc, 0x4e, 0x91, 0xd7, 0xdd, 0x4c, 0x4f, 0xd1, 0xf8, 0x00, 0xed, 0x45,
	0x21, 0xfc, 0x3d, 0x8e, 0x71, 0x06, 0xeb, 0xd2, 0x66, 0x87, 0x36, 0x3b, 0xb5, 0x43, 0x12, 0x7a,
	0x81, 0xcf, 0x06, 0xa1, 0xea, 0x3a, 0x03, 0x36, 0x35, 0x0a, 0xa9, 0x67, 0xb3, 0x75, 0x22, 0x17,
	0x56, 0xc5, 0x75, 0xd8, 0xda, 0x18, 0x42, 0xe5, 0x24, 0x08, 0x47, 0x09, 0xba, 0x07, 0x0d, 0x2e,
	0xbb, 0x80, 0x17, 0x2e, 0x9b, 0xc4, 0x75, 0x28, 0x34, 0x59, 0xe8, 0x50, 0x43, 0x74, 0x48, 0x72,
	0x2f, 0xe5, 0x33, 0xf3, 0x3f, 0xd4, 0xd9, 0xcf, 0x20, 0xe5, 0x2a, 0xba, 0x52, 0x1b, 0x09, 0x1a,
	0xc6, 0x80, 0x0d, 0x80, 0xeb, 0xa6, 0x0a, 0xbc, 0x98, 0xe2, 0x64, 0xd6, 0x71, 0x2d, 0xef, 0xf8,
	0xcb, 0xb9, 0xa9, 0xde, 0x54, 0x3d, 0x25, 0x31, 0xce, 0x46, 0x1b, 0x5a, 0x19, 0x40, 0x12, 0x91,
	0x30, 0xc1, 0xc6, 0x63, 0x68, 0x5b, 0x78, 0x42, 0x2e, 0xf1, 0x4d, 0xa8, 0xc6, 0x3a, 0xac, 0x15,
	0xbc, 0x64, 0xe8, 0x43, 0x58, 0x79, 0x8d, 0xe9, 0x8d, 0x71, 0x6b, 0xb0, 0x7a, 0x12, 0x24, 0xdc,
	0x27, 0x91, 0x4e, 0x92, 0x23, 0x2f, 0xe9, 0x9f, 0x72, 0x4c, 0x01, 0x64, 0xa1, 0xc7, 0x19, 0xc7,
	0x1b, 0x51, 0x0d, 0x68, 0x16, 0x86, 0x4d, 0xa0, 0x37, 0xac, 0x39, 0x5b, 0x7e, 0x0e, 0xc5, 0xf4,
	0x17, 0xd0, 0x3a, 0x0d, 0xfc, 0xf0, 0xed, 0x15, 0xfd, 0x73, 0x72, 0x13, 0xaa, 0xc3, 0xb1, 0x1d,
	0x4c, 0x12, 0x79, 0x6f, 0xfc, 0xd7, 0x13, 0x77, 0x7c, 0x4e, 0xe5, 0x34, 0xbd, 0xe3, 0x2d, 0xe9,
	0x86, 0x36, 0x98, 0xcc, 0xc6, 0x3e, 0x89, 0x03, 0x7a, 0x3e, 0x91, 0xea, 0xc9, 0x0d, 0xc6, 0x23,
	0x58, 0x9d, 0x41, 0x8a, 0x2a, 0x38, 0xe6, 0xe7, 0xab, 0xec, 0xda, 0xe0, 0xcb, 0xfe, 0xf7, 0x0a,
	0xac, 0xe5, 0xc3, 0x70, 0x8a, 0xe3, 0xcb, 0x60, 0x88, 0xd1, 0x88, 0x0d, 0x76, 0x2a, 0x01, 0xa4,
	0x72, 0x77, 0x15, 0xd5, 0xa8, 0x3f, 0x57, 0x0f, 0x90, 0x45, 0x51, 0x68, 0xcc, 0x74, 0x83, 0x14,
	0x5e, 0x87, 0x45, 0x29, 0xea, 0x3b, 0x77, 0x8a, 0x91, 0xa8, 0x9f, 0xa0, 0x2a, 0x84, 0xa9, 0x42,
	0x71, 0x4e, 0xc2, 0xfa, 0x13, 0x25, 0xb1, 0xa1, 0x73, 0xa8, 0x09, 0x5d, 0x27, 0x68, 0xfb, 0xf6,
	0x88, 0x85, 0x11, 0xd0, 0x9f, 0xaa, 0x3f, 0x1a, 0xb2, 0x5d, 0xfc, 0x1e, 0x56, 0x6b, 0x57, 0x2e,
	0x71, 0xc5, 0x76, 0x15, 0x94, 0x9c, 0xb7, 0x8b, 0xe3, 0x29, 0xb7, 0xab, 0x00, 0xb9, 0x73, 0xa7,
	0x18, 0x81, 0xda, 0xff, 0xaa, 0x89, 0x01, 0x62, 0x83, 0x96, 0x89, 0x34, 0x84, 0x9a, 0xd4, 0x37,
	0x52, 0x60, 0x31, 0x3f, 0x7d, 0xfa, 0xf6, 0x1d, 0x22, 0x44, 0x09, 0x07, 0x7b, 0x67, 0xbb, 0x3e,
	0x1b, 0xac, 0xa9, 0xd3, 0x1b, 0x92, 0x89, 0x69, 0x63, 0xcf, 0x1e, 0x0f, 0x49, 0x68, 0x0a, 0xdb,
	0x16, 0xbf, 0x56, 0x28, 0x7b, 0x33, 0xb7, 0x66, 0x7f, 0xc8, 0x7c, 0x62, 0x8a, 0x4c, 0x91, 0xe3,
	0x54, 0x53, 0xeb, 0xce, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xbe, 0xaf, 0xfa, 0xe6, 0x09,
	0x00, 0x00,
}
