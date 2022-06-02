package logic

import (
	"context"
	"time"

	"github.com/wslynn/wechat-gozero/app/user/model"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/svc"
	"github.com/wslynn/wechat-gozero/app/user/rpc/proto"
	"github.com/wslynn/wechat-gozero/common/xcrypt"
	"github.com/wslynn/wechat-gozero/common/xerr"
	"github.com/wslynn/wechat-gozero/common/xjwt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_DATA),
				"user login email not exist, email:%s,err:%v", in.Email, err)
		}
		return nil, errors.Wrapf(xerr.NewErrMsg("unknown"),
			"user login email unknown, email:%s,err:%v", in.Email, err)
	}
	// 判断密码是否正确
	ok := xcrypt.PasswordVerify(in.Password, user.Password)
	if !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UNAUTHORIZED),
			"user login password error, password:%s", in.Password, err)
	}
	// 生成jwt
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := xjwt.GetJwtToken(accessSecret, time.Now().Unix(), accessExpire, user.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("generate token fail"),
			"getJwtToken err userId:%d, err:%v", user.Id, err)
	}
	return &proto.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
