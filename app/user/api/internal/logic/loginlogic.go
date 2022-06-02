package logic

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/user/api/internal/svc"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/types"
	"github.com/wslynn/wechat-gozero/app/user/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var resp types.LoginResponse
	copier.Copy(&resp, loginResp)
	return &resp, nil
}
