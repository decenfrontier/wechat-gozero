package logic

import (
	"context"

	"wechat-gozero/app/user/model"
	"wechat-gozero/app/user/rpc/internal/svc"
	"wechat-gozero/app/user/rpc/proto"
	"wechat-gozero/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PersonalInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersonalInfoLogic {
	return &PersonalInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersonalInfoLogic) PersonalInfo(in *proto.PersonalInfoRequest) (*proto.PersonalInfoResponse, error) {
	// 查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_DATA), "PersonalInfo user not found id:%d", in.Id)
		} else {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "PersonalInfo db err , id:%d , err:%v", in.Id, err)
		}
	}
	var avatarUrl string
	if user.AvatarUrl.String != "" {
		avatarUrl = user.AvatarUrl.String
	} else {
		avatarUrl = model.DefaultAvatarUrl
	}
	return &proto.PersonalInfoResponse{
		UserId:    user.Id,
		NickName:  user.NickName,
		Gender:    user.Gender,
		Email:     user.Email,
		AvatarUrl: avatarUrl,
	}, nil
}
