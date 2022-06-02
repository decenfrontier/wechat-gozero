package logic

import (
	"context"

	"wechat-gozero/app/group/api/internal/svc"
	"wechat-gozero/app/group/api/internal/types"
	"wechat-gozero/app/group/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUesrListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUesrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUesrListLogic {
	return &GroupUesrListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUesrListLogic) GroupUesrList(req *types.GroupUserListRequest) (*types.GroupUserListResponse, error) {
	resp, err := l.svcCtx.GroupRpc.GroupUserList(l.ctx, &proto.GroupUserListRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GroupUserListResponse{
		List: resp.List,
	}, nil
}
