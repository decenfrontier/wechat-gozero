package logic

import (
	"context"
	"strings"

	"github.com/wslynn/wechat-gozero/app/group/model"
	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/proto/group"
	"github.com/wslynn/wechat-gozero/common/biz"
	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加好友
func (l *AddFriendLogic) AddFriend(in *group.AddFriendRequest) (*group.AddFriendResponse, error) {
	fromUid := in.FromUid
	toUid := in.ToUid
	// 生成groupId
	groupId := biz.GetGroupId(fromUid, toUid)
	// 查询这两个用户的nickName
	u1, err := l.svcCtx.UserModel.FindOne(l.ctx, fromUid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend query user failed, fromUid: %v", fromUid)
	}
	u2, err := l.svcCtx.UserModel.FindOne(l.ctx, toUid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend query user failed, toUid: %v", toUid)
	}
	groupName := strings.Join([]string{u1.NickName, u2.NickName}, ", ")
	// 创建一个group
	_, err = l.svcCtx.GroupModel.Insert(l.ctx, &model.Group{
		Id:     groupId,
		Name:   groupName,
		Type:   1,
		Status: model.GroupStatusNo,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend insert group failed: %v", err)
	}
	return &group.AddFriendResponse{
		GroupId: groupId,
	}, nil
}
