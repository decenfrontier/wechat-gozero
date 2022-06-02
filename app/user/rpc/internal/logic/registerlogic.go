package logic

import (
	"context"

	modelMsg "github.com/wslynn/wechat-gozero/app/msg/model"
	"github.com/wslynn/wechat-gozero/app/user/model"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/proto/user"
	"github.com/wslynn/wechat-gozero/common/biz"
	"github.com/wslynn/wechat-gozero/common/utils"
	"github.com/wslynn/wechat-gozero/common/xcrypt"
	"github.com/wslynn/wechat-gozero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err == nil { // 用户已存在
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DATA_EXIST),
			"user register exists, email:%s, err:%v", in.Email, err)
	}
	if err != nil && err != model.ErrNotFound { // 有其它错误
		return nil, errors.Wrapf(xerr.NewErrMsg("连接错误"),
			"user register fail, err:%v", err)
	}
	// 密码哈希处理
	hashedPwd, err := xcrypt.PasswordHash(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码处理失败"),
			"user register hash, password:%s, err:%v", in.Password, err)
	}
	in.Password = hashedPwd
	// 若未指定昵称, 则随机字符串
	if in.NickName == "" {
		in.NickName = biz.RandStr(8)
	}
	// 创建用户
	userModel := model.User{
		NickName: in.NickName,
		Gender:   in.Gender,
		Email:    in.Email,
		Password: in.Password,
	}
	// 创建事务
	err = l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 增加用户到db
		sqlRet, err := l.svcCtx.UserModel.TransInsert(ctx, session, &userModel)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register user create failed, user:%+v, err:%v", userModel, err)
		}
		user_id, err := sqlRet.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register user get uid failed, user:%+v, err:%v", userModel, err)
		}
		// 与 微信团队, 文件传输助手, 结为好友
		// 创建群
		_, err = l.svcCtx.GroupModel.TransInsertSystemGroup(ctx, session, user_id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register create system group failed, user:%+v, err: %v", userModel, err)
		}
		// 创建群用户
		_, err = l.svcCtx.GroupUserModel.TransInsertSystemGroupUser(ctx, session, user_id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register create system group user failed, user:%+v, err: %v", userModel, err)
		}
		// 创建 聊天问好消息
		var senderId int64 = 1
		groupId := biz.GetGroupId(senderId, user_id)
		chatMsg := &modelMsg.ChatMsg{
			GroupId:  groupId,
			SenderId: senderId,
			Type:     modelMsg.MsgTypeText,
			Content:  modelMsg.MsgSayHello,
			Uuid:     utils.GenUuid(),
		}
		_, err = l.svcCtx.ChatMsgModel.TransInsert(ctx, session, chatMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register say hello1 failed, chatMsg:%+v, err: %v", chatMsg, err)
		}
		senderId = 2
		groupId = biz.GetGroupId(senderId, user_id)
		chatMsg.GroupId = groupId
		chatMsg.SenderId = senderId
		chatMsg.Uuid = utils.GenUuid()
		_, err = l.svcCtx.ChatMsgModel.TransInsert(ctx, session, chatMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"register say hello2 failed, chatMsg:%+v, err: %v", chatMsg, err)
		}
		return nil // commit
	})
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{}, nil
}
