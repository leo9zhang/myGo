package handler

import (
	logic2 "mall/service/user/api/internal/logic"
	svc2 "mall/service/user/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic2.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
