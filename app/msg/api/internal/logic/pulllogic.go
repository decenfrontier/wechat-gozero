package logic

import (
	"context"

	"ws_chat/app/message/api/internal/svc"
	"ws_chat/app/message/api/internal/types"
	"ws_chat/app/message/rpc/proto"
	"ws_chat/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullLogic) Pull(req *types.PullRequest) (*types.PullResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	var pbPullRequest proto.PullRequest
	copier.Copy(&pbPullRequest, req)
	pbPullRequest.UserId = uid
	pbPullResponse, err := l.svcCtx.MessageRpc.Pull(l.ctx, &pbPullRequest)
	if err != nil {
		return nil, err
	}
	var resp types.PullResponse
	copier.Copy(&resp, pbPullResponse)
	return &resp, nil
}
