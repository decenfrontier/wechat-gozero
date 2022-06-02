package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wslynn/wechat-gozero/app/group/model"
	"github.com/wslynn/wechat-gozero/app/group/rpc/internal/svc"
	modelUser "github.com/wslynn/wechat-gozero/app/user/model"
	"github.com/wslynn/wechat-gozero/common/biz"
	"github.com/wslynn/wechat-gozero/common/xerr"
	"github.com/wslynn/wechat-gozero/proto/group"
	"github.com/wslynn/wechat-gozero/proto/msg"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageGroupInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageGroupInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageGroupInfoListLogic {
	return &MessageGroupInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录后刷新消息页
func (l *MessageGroupInfoListLogic) MessageGroupInfoList(in *group.MessageGroupInfoListRequest) (*group.MessageGroupInfoListResponse, error) {
	userId := in.UserId // 自己
	// 查询 userId的 所有群组
	groupIdList, err := l.svcCtx.GroupUserModel.FindGroupIdListByUserId(l.ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
			"MessageGroupInfoList failed, userId: %v, err: %v", userId, err)
	}
	var list []*group.MessageGroupInfo
	for _, groupId := range groupIdList {
		var avatarUrl string // 头像
		// 获取群组信息
		groupModel, err := l.svcCtx.GroupModel.FindOne(l.ctx, groupId)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			continue
		}
		// 判断 group类型
		if groupModel.Type == model.GroupTypeSingle {
			// 若为单聊, 到 user 表里查对方的avatarUrl
			friend_uid, _ := biz.GetFriendIdFromGroupId(groupId, userId)
			friend, err := l.svcCtx.UserModel.FindOne(l.ctx, friend_uid)
			if err == nil {
				avatarUrl = friend.AvatarUrl.String
			}
		} else {
			// 若为群聊, 到 group的config中取出avatarUrl
			jsonBytes := []byte(groupModel.Config.String)
			var v = make(map[string]interface{})
			json.Unmarshal(jsonBytes, &v)
			url, ok := v["avatarUrl"]
			if ok {
				avatarUrl = url.(string)
			}
		}
		if avatarUrl == "" {
			avatarUrl = modelUser.DefaultAvatarUrl
		}

		// 到 group_user表里查aliasName
		aliasName, err := l.svcCtx.GroupUserModel.FindAliasNameByGroupAndUser(l.ctx, groupId, userId)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("【RPC-SRV-ERR】 FindAliasNameByGroupAndUser failed error: %v\n", err)
			continue
		}

		// 到 chat_msg 表里查每个组的最后一条消息
		chatMsg, err := l.svcCtx.ChatMsgModel.FindLastMsgByGroupId(l.ctx, groupId)
		if err != nil {
			fmt.Printf("FindLastMsgByGroupId failed error: %v\n", err)
			continue
		}
		fmt.Printf("groupId:%s, chatMsg:%+v\n", groupId, chatMsg)
		var pbChatMsg msg.ChatMsg
		err = copier.Copy(&pbChatMsg, chatMsg)
		if err != nil {
			fmt.Printf("cope error:%v\n", err)
			continue
		}
		pbChatMsg.CreateTime = chatMsg.CreateTime.UnixMilli()
		list = append(list, &group.MessageGroupInfo{
			GroupId:   groupId,
			AliasName: aliasName,
			AvatarUrl: avatarUrl,
			LastMsg:   &pbChatMsg,
		})
	}
	return &group.MessageGroupInfoListResponse{
		List: list,
	}, nil
}
