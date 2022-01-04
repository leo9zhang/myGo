package logic

import (
	"context"
	"mall/common/cryptx"
	model2 "mall/service/user/model"
	svc2 "mall/service/user/rpc/internal/svc"
	user2 "mall/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc2.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc2.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user2.LoginRequest) (*user2.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model2.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user2.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
