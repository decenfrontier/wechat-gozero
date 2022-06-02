package logic

import (
	"context"

	"wechat-gozero/app/group/api/internal/svc"
	"wechat-gozero/app/group/api/internal/types"
	"wechat-gozero/app/group/rpc/proto"
	"wechat-gozero/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendLogic {
	return &HandleFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleFriendLogic) HandleFriend(req *types.HandleFriendRequest) (*types.HandleFriendResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.GroupRpc.HandleFriend(l.ctx, &proto.HandleFriendRequest{
		UserId:  uid,
		GroupId: req.GroupId,
		IsAgree: req.IsAgree,
	})
	if err != nil {
		return nil, err
	}
	return &types.HandleFriendResponse{
		GroupId: resp.GroupId,
	}, nil
}
