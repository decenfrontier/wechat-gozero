package handler

import (
	"net/http"

	"wechat-gozero/common/xresp"

	"wechat-gozero/app/group/api/internal/logic"
	"wechat-gozero/app/group/api/internal/svc"
	"wechat-gozero/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HandleFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HandleFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewHandleFriendLogic(r.Context(), svcCtx)
		resp, err := l.HandleFriend(&req)
		xresp.Response(r, w, resp, err)
	}
}
