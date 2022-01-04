package template

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"myGo/api/internal/logic/template"
	"myGo/api/internal/svc"
	"myGo/api/internal/types"
)

func HandlerNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := template.NewHandlerNameLogic(r.Context(), ctx)
		resp, err := l.HandlerName(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
