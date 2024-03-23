// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: game_service.proto

package proto

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

// 错误码
type StatusCode int32

const (
	StatusCode_OK    StatusCode = 0
	StatusCode_ERROR StatusCode = 1
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	StatusCode_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_service_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_game_service_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{0}
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.StatusCode" json:"code,omitempty"`
	Content string     `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMessage) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_OK
}

func (x *ErrorMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type FriendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FriendInfo) Reset() {
	*x = FriendInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendInfo) ProtoMessage() {}

func (x *FriendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendInfo.ProtoReflect.Descriptor instead.
func (*FriendInfo) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{1}
}

func (x *FriendInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *FriendInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FriendShipList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ErrorMessage `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	List  []*FriendInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FriendShipList) Reset() {
	*x = FriendShipList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendShipList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendShipList) ProtoMessage() {}

func (x *FriendShipList) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendShipList.ProtoReflect.Descriptor instead.
func (*FriendShipList) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{2}
}

func (x *FriendShipList) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FriendShipList) GetList() []*FriendInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type FriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *FriendListRequest) Reset() {
	*x = FriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListRequest) ProtoMessage() {}

func (x *FriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListRequest.ProtoReflect.Descriptor instead.
func (*FriendListRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{3}
}

func (x *FriendListRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid string `protobuf:"bytes,1,opt,name=senderUid,proto3" json:"senderUid,omitempty"`
	TargetUid string `protobuf:"bytes,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
}

func (x *DeleteFriendRequest) Reset() {
	*x = DeleteFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendRequest) ProtoMessage() {}

func (x *DeleteFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendRequest.ProtoReflect.Descriptor instead.
func (*DeleteFriendRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFriendRequest) GetSenderUid() string {
	if x != nil {
		return x.SenderUid
	}
	return ""
}

func (x *DeleteFriendRequest) GetTargetUid() string {
	if x != nil {
		return x.TargetUid
	}
	return ""
}

type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid string `protobuf:"bytes,1,opt,name=senderUid,proto3" json:"senderUid,omitempty"`
	TargetUid string `protobuf:"bytes,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendRequest.ProtoReflect.Descriptor instead.
func (*AddFriendRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddFriendRequest) GetSenderUid() string {
	if x != nil {
		return x.SenderUid
	}
	return ""
}

func (x *AddFriendRequest) GetTargetUid() string {
	if x != nil {
		return x.TargetUid
	}
	return ""
}

type SearchFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearcherUid string `protobuf:"bytes,1,opt,name=searcherUid,proto3" json:"searcherUid,omitempty"`
	TargetUid   string `protobuf:"bytes,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
}

func (x *SearchFriendRequest) Reset() {
	*x = SearchFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFriendRequest) ProtoMessage() {}

func (x *SearchFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFriendRequest.ProtoReflect.Descriptor instead.
func (*SearchFriendRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{6}
}

func (x *SearchFriendRequest) GetSearcherUid() string {
	if x != nil {
		return x.SearcherUid
	}
	return ""
}

func (x *SearchFriendRequest) GetTargetUid() string {
	if x != nil {
		return x.TargetUid
	}
	return ""
}

type SearchFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error      *ErrorMessage `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	TargetUid  string        `protobuf:"bytes,2,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
	TargetName string        `protobuf:"bytes,3,opt,name=targetName,proto3" json:"targetName,omitempty"`
	IsFriend   bool          `protobuf:"varint,4,opt,name=isFriend,proto3" json:"isFriend,omitempty"`
}

func (x *SearchFriendResponse) Reset() {
	*x = SearchFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFriendResponse) ProtoMessage() {}

func (x *SearchFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFriendResponse.ProtoReflect.Descriptor instead.
func (*SearchFriendResponse) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{7}
}

func (x *SearchFriendResponse) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchFriendResponse) GetTargetUid() string {
	if x != nil {
		return x.TargetUid
	}
	return ""
}

func (x *SearchFriendResponse) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *SearchFriendResponse) GetIsFriend() bool {
	if x != nil {
		return x.IsFriend
	}
	return false
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UidRequest) Reset() {
	*x = UidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidRequest) ProtoMessage() {}

func (x *UidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidRequest.ProtoReflect.Descriptor instead.
func (*UidRequest) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{9}
}

type UidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ErrorMessage `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Uid   string        `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UidResponse) Reset() {
	*x = UidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidResponse) ProtoMessage() {}

func (x *UidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidResponse.ProtoReflect.Descriptor instead.
func (*UidResponse) Descriptor() ([]byte, []int) {
	return file_game_service_proto_rawDescGZIP(), []int{10}
}

func (x *UidResponse) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UidResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_game_service_proto protoreflect.FileDescriptor

var file_game_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x0a,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x62, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x4e,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x55,
	0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x0a, 0x0a,
	0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0b, 0x55, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x2a, 0x1f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0x82, 0x03, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x69,
	0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x69, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53,
	0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_service_proto_rawDescOnce sync.Once
	file_game_service_proto_rawDescData = file_game_service_proto_rawDesc
)

func file_game_service_proto_rawDescGZIP() []byte {
	file_game_service_proto_rawDescOnce.Do(func() {
		file_game_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_service_proto_rawDescData)
	})
	return file_game_service_proto_rawDescData
}

var file_game_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_game_service_proto_goTypes = []interface{}{
	(StatusCode)(0),              // 0: proto.StatusCode
	(*ErrorMessage)(nil),         // 1: proto.ErrorMessage
	(*FriendInfo)(nil),           // 2: proto.FriendInfo
	(*FriendShipList)(nil),       // 3: proto.FriendShipList
	(*FriendListRequest)(nil),    // 4: proto.FriendListRequest
	(*DeleteFriendRequest)(nil),  // 5: proto.DeleteFriendRequest
	(*AddFriendRequest)(nil),     // 6: proto.AddFriendRequest
	(*SearchFriendRequest)(nil),  // 7: proto.SearchFriendRequest
	(*SearchFriendResponse)(nil), // 8: proto.SearchFriendResponse
	(*RegisterRequest)(nil),      // 9: proto.RegisterRequest
	(*UidRequest)(nil),           // 10: proto.UidRequest
	(*UidResponse)(nil),          // 11: proto.UidResponse
}
var file_game_service_proto_depIdxs = []int32{
	0,  // 0: proto.ErrorMessage.code:type_name -> proto.StatusCode
	1,  // 1: proto.FriendShipList.error:type_name -> proto.ErrorMessage
	2,  // 2: proto.FriendShipList.list:type_name -> proto.FriendInfo
	1,  // 3: proto.SearchFriendResponse.error:type_name -> proto.ErrorMessage
	1,  // 4: proto.UidResponse.error:type_name -> proto.ErrorMessage
	10, // 5: proto.GameService.GetUid:input_type -> proto.UidRequest
	9,  // 6: proto.GameService.RegisterUser:input_type -> proto.RegisterRequest
	4,  // 7: proto.GameService.GetFriendList:input_type -> proto.FriendListRequest
	6,  // 8: proto.GameService.AddFriend:input_type -> proto.AddFriendRequest
	5,  // 9: proto.GameService.DeleteFriend:input_type -> proto.DeleteFriendRequest
	7,  // 10: proto.GameService.SearchFriend:input_type -> proto.SearchFriendRequest
	11, // 11: proto.GameService.GetUid:output_type -> proto.UidResponse
	1,  // 12: proto.GameService.RegisterUser:output_type -> proto.ErrorMessage
	3,  // 13: proto.GameService.GetFriendList:output_type -> proto.FriendShipList
	1,  // 14: proto.GameService.AddFriend:output_type -> proto.ErrorMessage
	1,  // 15: proto.GameService.DeleteFriend:output_type -> proto.ErrorMessage
	8,  // 16: proto.GameService.SearchFriend:output_type -> proto.SearchFriendResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_game_service_proto_init() }
func file_game_service_proto_init() {
	if File_game_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_game_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendInfo); i {
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
		file_game_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendShipList); i {
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
		file_game_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendListRequest); i {
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
		file_game_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendRequest); i {
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
		file_game_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendRequest); i {
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
		file_game_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFriendRequest); i {
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
		file_game_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFriendResponse); i {
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
		file_game_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_game_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidRequest); i {
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
		file_game_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidResponse); i {
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
			RawDescriptor: file_game_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_service_proto_goTypes,
		DependencyIndexes: file_game_service_proto_depIdxs,
		EnumInfos:         file_game_service_proto_enumTypes,
		MessageInfos:      file_game_service_proto_msgTypes,
	}.Build()
	File_game_service_proto = out.File
	file_game_service_proto_rawDesc = nil
	file_game_service_proto_goTypes = nil
	file_game_service_proto_depIdxs = nil
}
