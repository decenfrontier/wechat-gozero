package logic

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/msg/api/internal/svc"
	"github.com/wslynn/wechat-gozero/app/msg/api/internal/types"
	"github.com/wslynn/wechat-gozero/proto/msg"
	"github.com/wslynn/wechat-gozero/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (*types.UploadResponse, error) {
	var pbUploadRequest msg.UploadRequest
	err := copier.Copy(&pbUploadRequest, req)
	if err != nil {
		return nil, err
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	pbUploadRequest.SenderId = userId
	pbUploadResponse, err := l.svcCtx.MsgRpc.Upload(l.ctx, &pbUploadRequest)
	if err != nil {
		return nil, err
	}
	var resp types.UploadResponse
	copier.Copy(&resp, pbUploadResponse)
	return &resp, nil
}
