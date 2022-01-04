package handler

import (
	logic2 "mall/service/user/api/internal/logic"
	svc2 "mall/service/user/api/internal/svc"
	types2 "mall/service/user/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types2.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic2.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
