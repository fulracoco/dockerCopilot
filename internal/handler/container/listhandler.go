package container

import (
	"net/http"

	"github.com/onlyLTY/dokcerCopilot/UGREEN/internal/logic/container"
	"github.com/onlyLTY/dokcerCopilot/UGREEN/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := container.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		if err != nil {
			httpx.WriteJson(w, resp.Code, resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
