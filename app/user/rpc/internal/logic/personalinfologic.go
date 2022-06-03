package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/wslynn/wechat-gozero/app/user/model"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/common/xerr"
	"github.com/wslynn/wechat-gozero/proto/user"

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

func (l *PersonalInfoLogic) PersonalInfo(in *user.PersonalInfoRequest) (*user.PersonalInfoResponse, error) {
	// 查询用户是否存在
	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_DATA), "PersonalInfo user not found id:%d", in.Id)
		} else {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "PersonalInfo db err , id:%d , err:%v", in.Id, err)
		}
	}
	var avatarUrl string
	if userModel.AvatarUrl.String != "" {
		avatarUrl = userModel.AvatarUrl.String
	} else {
		avatarUrl = model.DefaultAvatarUrl
	}
	var resp user.PersonalInfoResponse
	copier.Copy(&resp, userModel)
	resp.UserId = userModel.Id
	resp.AvatarUrl = avatarUrl
	return &resp, nil
}
