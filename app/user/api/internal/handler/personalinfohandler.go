package handler

import (
	"net/http"

	"github.com/wslynn/wechat-gozero/app/user/api/internal/logic"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/svc"
	"github.com/wslynn/wechat-gozero/app/user/api/internal/types"
	"github.com/wslynn/wechat-gozero/common/xresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PersonalInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonalInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewPersonalInfoLogic(r.Context(), svcCtx)
		resp, err := l.PersonalInfo(&req)
		xresp.Response(r, w, resp, err)
	}
}
