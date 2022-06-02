package logic

import (
	"context"

	"wechat-gozero/app/group/api/internal/svc"
	"wechat-gozero/app/group/api/internal/types"
	"wechat-gozero/app/group/rpc/proto"
	"wechat-gozero/common/ctxdata"
	"wechat-gozero/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageGroupInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageGroupInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageGroupInfoListLogic {
	return &MessageGroupInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageGroupInfoListLogic) MessageGroupInfoList(req *types.MessageGroupInfoListRequest) (*types.MessageGroupInfoListResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.GroupRpc.MessageGroupInfoList(l.ctx, &proto.MessageGroupInfoListRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.MessageGroupInfo
	err = copier.Copy(&list, resp.List)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.MARSHAL_ERROR)
	}
	return &types.MessageGroupInfoListResponse{
		List: list,
	}, nil
}
