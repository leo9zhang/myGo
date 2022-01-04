package logic

import (
	"context"
	svc2 "mall/service/user/api/internal/svc"
	types2 "mall/service/user/api/internal/types"
	userclient2 "mall/service/user/rpc/userclient"
	"time"

	"mall/common/jwtx"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc2.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types2.LoginRequest) (resp *types2.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient2.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types2.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
