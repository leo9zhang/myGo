package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/status"
	model2 "mall/service/user/model"
	svc2 "mall/service/user/rpc/internal/svc"
	user2 "mall/service/user/rpc/user"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc2.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc2.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user2.UserInfoRequest) (*user2.UserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		if err == model2.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user2.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
