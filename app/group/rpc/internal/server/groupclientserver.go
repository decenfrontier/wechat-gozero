// Code generated by goctl. DO NOT EDIT!
// Source: group.proto

package server

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/logic"
	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/proto/group"
)

type GroupClientServer struct {
	svcCtx *svc.ServiceContext
	group.UnimplementedGroupClientServer
}

func NewGroupClientServer(svcCtx *svc.ServiceContext) *GroupClientServer {
	return &GroupClientServer{
		svcCtx: svcCtx,
	}
}

func (s *GroupClientServer) AddFriend(ctx context.Context, in *group.AddFriendRequest) (*group.AddFriendResponse, error) {
	l := logic.NewAddFriendLogic(ctx, s.svcCtx)
	return l.AddFriend(in)
}

func (s *GroupClientServer) HandleFriend(ctx context.Context, in *group.HandleFriendRequest) (*group.HandleFriendResponse, error) {
	l := logic.NewHandleFriendLogic(ctx, s.svcCtx)
	return l.HandleFriend(in)
}

func (s *GroupClientServer) GroupUserList(ctx context.Context, in *group.GroupUserListRequest) (*group.GroupUserListResponse, error) {
	l := logic.NewGroupUserListLogic(ctx, s.svcCtx)
	return l.GroupUserList(in)
}

func (s *GroupClientServer) UserGroupList(ctx context.Context, in *group.UserGroupListRequest) (*group.UserGroupListResponse, error) {
	l := logic.NewUserGroupListLogic(ctx, s.svcCtx)
	return l.UserGroupList(in)
}

func (s *GroupClientServer) MessageGroupInfoList(ctx context.Context, in *group.MessageGroupInfoListRequest) (*group.MessageGroupInfoListResponse, error) {
	l := logic.NewMessageGroupInfoListLogic(ctx, s.svcCtx)
	return l.MessageGroupInfoList(in)
}
