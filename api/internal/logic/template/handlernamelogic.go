package template

import (
	"context"

	"myGo/api/internal/svc"
	"myGo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HandlerNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) HandlerNameLogic {
	return HandlerNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerNameLogic) HandlerName(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
