package logic

import (
	"context"
	"database/sql"

	"github.com/wslynn/wechat-gozero/app/group/model"
	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/proto/group"
	modelMsg "github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/common/biz"
	"github.com/wslynn/wechat-gozero/common/utils"
	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type HandleFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleFriendLogic {
	return &HandleFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 处理好友请求
func (l *HandleFriendLogic) HandleFriend(in *group.HandleFriendRequest) (*group.HandleFriendResponse, error) {
	groupId := in.GroupId
	uid1 := in.UserId // 同意好友申请的用户的uid
	uid2, _ := biz.GetFriendIdFromGroupId(groupId, uid1)
	// 先查询是否有未同意的该群
	groupModel, err := l.svcCtx.GroupModel.FindOne(l.ctx, groupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "好友请求不存在"), "HandleFriend find group error, groupId:%s, err:%v", groupId, err)
	}
	if groupModel.Status != model.GroupStatusNo { // 不是未同意, 则直接返回
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CLIENT_ERROR), "HandleFriend status != 0, groupId:%s, status:%v", groupId, groupModel.Status)
	}
	// 查询两个用户的信息
	u1, _ := l.svcCtx.UserModel.FindOne(l.ctx, uid1)
	u2, _ := l.svcCtx.UserModel.FindOne(l.ctx, uid2)
	// 创建事务
	err = l.svcCtx.GroupModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 把status设为1
		groupModel.Status = model.GroupStatusYes
		err = l.svcCtx.GroupModel.TransUpdate(l.ctx, session, groupModel)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HandleFriend update group failed, groupId:%s, err:%v", groupId, err)
		}
		// 创建群内用户
		groupUser1 := &model.GroupUser{
			GroupId:   groupId,
			UserId:    uid1,
			AliasName: sql.NullString{String: u2.NickName, Valid: true},
		}
		_, err = l.svcCtx.GroupUserModel.TransInsert(l.ctx, session, groupUser1)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HandleFriend insert group user failed, groupUser:%+v, err:%v", groupUser1, err)
		}
		groupUser2 := &model.GroupUser{
			GroupId:   groupId,
			UserId:    uid2,
			AliasName: sql.NullString{String: u1.NickName, Valid: true},
		}
		_, err = l.svcCtx.GroupUserModel.TransInsert(l.ctx, session, groupUser2)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HandleFriend insert group user failed, groupUser:%+v, err:%v", groupUser2, err)
		}
		// 创建sayHello消息
		chatMsg := &modelMsg.ChatMsg{
			GroupId:  groupId,
			SenderId: uid1,
			Type:     modelMsg.MsgTypeText,
			Content:  modelMsg.MsgSayHello,
			Uuid:     utils.GenUuid(),
		}
		_, err = l.svcCtx.ChatMsgModel.TransInsert(ctx, session, chatMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"HandleFriend say hello failed, chatMsg:%+v, err: %v", chatMsg, err)
		}
		return nil // commit
	})
	if err != nil {
		return nil, err
	}

	return &group.HandleFriendResponse{
		GroupId: groupId,
	}, nil
}
