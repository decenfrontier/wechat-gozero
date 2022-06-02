package handler

import (
	"net/http"

	"wechat-gozero/common/xresp"

	"wechat-gozero/app/message/api/internal/logic"
	"wechat-gozero/app/message/api/internal/svc"
	"wechat-gozero/app/message/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func pullHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PullRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewPullLogic(r.Context(), svcCtx)
		resp, err := l.Pull(&req)
		xresp.Response(r, w, resp, err)
	}
}
