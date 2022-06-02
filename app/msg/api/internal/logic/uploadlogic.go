package logic

import (
	"context"

	"wechat-gozero/app/message/api/internal/svc"
	"wechat-gozero/app/message/api/internal/types"
	"wechat-gozero/app/message/rpc/proto"
	"wechat-gozero/common/ctxdata"

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
	var pbUploadRequest proto.UploadRequest
	err := copier.Copy(&pbUploadRequest, req)
	if err != nil {
		return nil, err
	}
	userId := ctxdata.GetUidFromCtx(l.ctx)
	pbUploadRequest.SenderId = userId
	pbUploadResponse, err := l.svcCtx.MessageRpc.Upload(l.ctx, &pbUploadRequest)
	if err != nil {
		return nil, err
	}
	var resp types.UploadResponse
	copier.Copy(&resp, pbUploadResponse)
	return &resp, nil
}
