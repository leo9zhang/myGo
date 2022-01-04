package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	svc2 "mall/service/user/api/internal/svc"
	types2 "mall/service/user/api/internal/types"
	userclient2 "mall/service/user/rpc/userclient"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc2.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types2.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient2.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	return &types2.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
