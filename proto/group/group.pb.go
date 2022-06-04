// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/group/group.proto

package group

import (
	msg "github.com/wslynn/wechat-gozero/proto/msg"
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

// 添加好友
type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUid int64 `protobuf:"varint,1,opt,name=fromUid,proto3" json:"fromUid,omitempty"`
	ToUid   int64 `protobuf:"varint,2,opt,name=toUid,proto3" json:"toUid,omitempty"`
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[0]
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
	return file_proto_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendRequest) GetFromUid() int64 {
	if x != nil {
		return x.FromUid
	}
	return 0
}

func (x *AddFriendRequest) GetToUid() int64 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

type AddFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"` // 返回创建的这个新群, 但是这个群还处于删除态, 不可发聊天消息
}

func (x *AddFriendResponse) Reset() {
	*x = AddFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendResponse) ProtoMessage() {}

func (x *AddFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendResponse.ProtoReflect.Descriptor instead.
func (*AddFriendResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{1}
}

func (x *AddFriendResponse) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// 处理好友申请
type HandleFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GroupId string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	IsAgree bool   `protobuf:"varint,3,opt,name=isAgree,proto3" json:"isAgree,omitempty"`
}

func (x *HandleFriendRequest) Reset() {
	*x = HandleFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleFriendRequest) ProtoMessage() {}

func (x *HandleFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleFriendRequest.ProtoReflect.Descriptor instead.
func (*HandleFriendRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{2}
}

func (x *HandleFriendRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *HandleFriendRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *HandleFriendRequest) GetIsAgree() bool {
	if x != nil {
		return x.IsAgree
	}
	return false
}

type HandleFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *HandleFriendResponse) Reset() {
	*x = HandleFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleFriendResponse) ProtoMessage() {}

func (x *HandleFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleFriendResponse.ProtoReflect.Descriptor instead.
func (*HandleFriendResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{3}
}

func (x *HandleFriendResponse) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// 获取群内的用户
type GroupUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GroupUserListRequest) Reset() {
	*x = GroupUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUserListRequest) ProtoMessage() {}

func (x *GroupUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUserListRequest.ProtoReflect.Descriptor instead.
func (*GroupUserListRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{4}
}

func (x *GroupUserListRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GroupUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []int64 `protobuf:"varint,1,rep,packed,name=list,proto3" json:"list,omitempty"`
}

func (x *GroupUserListResponse) Reset() {
	*x = GroupUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUserListResponse) ProtoMessage() {}

func (x *GroupUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUserListResponse.ProtoReflect.Descriptor instead.
func (*GroupUserListResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{5}
}

func (x *GroupUserListResponse) GetList() []int64 {
	if x != nil {
		return x.List
	}
	return nil
}

// 获取用户的群列表
type UserGroupListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserGroupListRequest) Reset() {
	*x = UserGroupListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupListRequest) ProtoMessage() {}

func (x *UserGroupListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupListRequest.ProtoReflect.Descriptor instead.
func (*UserGroupListRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{6}
}

func (x *UserGroupListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserGroupListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserGroupListResponse) Reset() {
	*x = UserGroupListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupListResponse) ProtoMessage() {}

func (x *UserGroupListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupListResponse.ProtoReflect.Descriptor instead.
func (*UserGroupListResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{7}
}

func (x *UserGroupListResponse) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

// 获取消息页面 群组信息列表
type MessageGroupInfoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MessageGroupInfoListRequest) Reset() {
	*x = MessageGroupInfoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageGroupInfoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroupInfoListRequest) ProtoMessage() {}

func (x *MessageGroupInfoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroupInfoListRequest.ProtoReflect.Descriptor instead.
func (*MessageGroupInfoListRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{8}
}

func (x *MessageGroupInfoListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type MessageGroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   string       `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`     // 群组id (group_user表)
	AliasName string       `protobuf:"bytes,2,opt,name=aliasName,proto3" json:"aliasName,omitempty"` // 备注 (group_user表)
	AvatarUrl string       `protobuf:"bytes,3,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"` // 头像 (group表, user表)
	LastMsg   *msg.ChatMsg `protobuf:"bytes,4,opt,name=lastMsg,proto3" json:"lastMsg,omitempty"`     // 最后一条消息 (chat_msg表)
}

func (x *MessageGroupInfo) Reset() {
	*x = MessageGroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroupInfo) ProtoMessage() {}

func (x *MessageGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroupInfo.ProtoReflect.Descriptor instead.
func (*MessageGroupInfo) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{9}
}

func (x *MessageGroupInfo) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *MessageGroupInfo) GetAliasName() string {
	if x != nil {
		return x.AliasName
	}
	return ""
}

func (x *MessageGroupInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *MessageGroupInfo) GetLastMsg() *msg.ChatMsg {
	if x != nil {
		return x.LastMsg
	}
	return nil
}

type MessageGroupInfoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*MessageGroupInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MessageGroupInfoListResponse) Reset() {
	*x = MessageGroupInfoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageGroupInfoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroupInfoListResponse) ProtoMessage() {}

func (x *MessageGroupInfoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroupInfoListResponse.ProtoReflect.Descriptor instead.
func (*MessageGroupInfoListResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{10}
}

func (x *MessageGroupInfoListResponse) GetList() []*MessageGroupInfo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_group_group_proto protoreflect.FileDescriptor

var file_proto_group_group_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x6f,
	0x6d, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d,
	0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x13, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x67, 0x72, 0x65, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x30, 0x0a,
	0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x15, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x15,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x1b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x90, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x07, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x22, 0x4b, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0x8f, 0x03, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5f, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x73, 0x6c, 0x79, 0x6e, 0x6e, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x67,
	0x6f, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_group_group_proto_rawDescOnce sync.Once
	file_proto_group_group_proto_rawDescData = file_proto_group_group_proto_rawDesc
)

func file_proto_group_group_proto_rawDescGZIP() []byte {
	file_proto_group_group_proto_rawDescOnce.Do(func() {
		file_proto_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_group_group_proto_rawDescData)
	})
	return file_proto_group_group_proto_rawDescData
}

var file_proto_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_group_group_proto_goTypes = []interface{}{
	(*AddFriendRequest)(nil),             // 0: group.AddFriendRequest
	(*AddFriendResponse)(nil),            // 1: group.AddFriendResponse
	(*HandleFriendRequest)(nil),          // 2: group.HandleFriendRequest
	(*HandleFriendResponse)(nil),         // 3: group.HandleFriendResponse
	(*GroupUserListRequest)(nil),         // 4: group.GroupUserListRequest
	(*GroupUserListResponse)(nil),        // 5: group.GroupUserListResponse
	(*UserGroupListRequest)(nil),         // 6: group.UserGroupListRequest
	(*UserGroupListResponse)(nil),        // 7: group.UserGroupListResponse
	(*MessageGroupInfoListRequest)(nil),  // 8: group.MessageGroupInfoListRequest
	(*MessageGroupInfo)(nil),             // 9: group.MessageGroupInfo
	(*MessageGroupInfoListResponse)(nil), // 10: group.MessageGroupInfoListResponse
	(*msg.ChatMsg)(nil),                  // 11: msg.ChatMsg
}
var file_proto_group_group_proto_depIdxs = []int32{
	11, // 0: group.MessageGroupInfo.lastMsg:type_name -> msg.ChatMsg
	9,  // 1: group.MessageGroupInfoListResponse.list:type_name -> group.MessageGroupInfo
	0,  // 2: group.GroupClient.AddFriend:input_type -> group.AddFriendRequest
	2,  // 3: group.GroupClient.HandleFriend:input_type -> group.HandleFriendRequest
	4,  // 4: group.GroupClient.GroupUserList:input_type -> group.GroupUserListRequest
	6,  // 5: group.GroupClient.UserGroupList:input_type -> group.UserGroupListRequest
	8,  // 6: group.GroupClient.MessageGroupInfoList:input_type -> group.MessageGroupInfoListRequest
	1,  // 7: group.GroupClient.AddFriend:output_type -> group.AddFriendResponse
	3,  // 8: group.GroupClient.HandleFriend:output_type -> group.HandleFriendResponse
	5,  // 9: group.GroupClient.GroupUserList:output_type -> group.GroupUserListResponse
	7,  // 10: group.GroupClient.UserGroupList:output_type -> group.UserGroupListResponse
	10, // 11: group.GroupClient.MessageGroupInfoList:output_type -> group.MessageGroupInfoListResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_group_group_proto_init() }
func file_proto_group_group_proto_init() {
	if File_proto_group_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_group_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_group_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendResponse); i {
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
		file_proto_group_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleFriendRequest); i {
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
		file_proto_group_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleFriendResponse); i {
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
		file_proto_group_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUserListRequest); i {
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
		file_proto_group_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupUserListResponse); i {
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
		file_proto_group_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupListRequest); i {
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
		file_proto_group_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroupListResponse); i {
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
		file_proto_group_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageGroupInfoListRequest); i {
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
		file_proto_group_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageGroupInfo); i {
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
		file_proto_group_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageGroupInfoListResponse); i {
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
			RawDescriptor: file_proto_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_group_group_proto_goTypes,
		DependencyIndexes: file_proto_group_group_proto_depIdxs,
		MessageInfos:      file_proto_group_group_proto_msgTypes,
	}.Build()
	File_proto_group_group_proto = out.File
	file_proto_group_group_proto_rawDesc = nil
	file_proto_group_group_proto_goTypes = nil
	file_proto_group_group_proto_depIdxs = nil
}
