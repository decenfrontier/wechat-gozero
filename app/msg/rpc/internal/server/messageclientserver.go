// Code generated by goctl. DO NOT EDIT!
// Source: msg.proto

package server

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/group/rpc/proto"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/logic"
	"github.com/wslynn/wechat-gozero/app/msg/rpc/internal/svc"
)

type MessageClientServer struct {
	svcCtx *svc.ServiceContext
	msg.UnimplementedMessageClientServer
}

func NewMessageClientServer(svcCtx *svc.ServiceContext) *MessageClientServer {
	return &MessageClientServer{
		svcCtx: svcCtx,
	}
}

func (s *MessageClientServer) Upload(ctx context.Context, in *msg.UploadRequest) (*msg.UploadResponse, error) {
	l := logic.NewUploadLogic(ctx, s.svcCtx)
	return l.Upload(in)
}

func (s *MessageClientServer) Pull(ctx context.Context, in *msg.PullRequest) (*msg.PullResponse, error) {
	l := logic.NewPullLogic(ctx, s.svcCtx)
	return l.Pull(in)
}