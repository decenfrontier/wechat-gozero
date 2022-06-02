package logic

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/user/api/internal/svc"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/types"
	"github.com/wslynn/wechat-gozero/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	// 调用rpc
	_, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		NickName: req.NickName,
		Gender:   req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
